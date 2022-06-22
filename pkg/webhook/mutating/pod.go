/*
 Copyright 2021 The Hybridnet Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package mutating

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	networkingv1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	"github.com/alibaba/hybridnet/pkg/constants"
	"github.com/alibaba/hybridnet/pkg/ipam/strategy"
	ipamtypes "github.com/alibaba/hybridnet/pkg/ipam/types"
	"github.com/alibaba/hybridnet/pkg/utils"
	webhookutils "github.com/alibaba/hybridnet/pkg/webhook/utils"
)

// TaintNodeNetworkUnavailable will be added when node's network is unavailable
// and removed when network becomes ready.
const TaintNodeNetworkUnavailable = "node.kubernetes.io/network-unavailable"

var podGVK = corev1.SchemeGroupVersion.WithKind("Pod")

func init() {
	createHandlers[gvkConverter(podGVK)] = PodCreateMutation
}

func PodCreateMutation(ctx context.Context, req *admission.Request, handler *Handler) admission.Response {
	logger := log.FromContext(ctx)

	pod := &corev1.Pod{}
	err := handler.Decoder.Decode(*req, pod)
	if err != nil {
		return webhookutils.AdmissionErroredWithLog(http.StatusBadRequest, err, logger)
	}

	// special mutation for host networking pods
	if pod.Spec.HostNetwork {
		// make sure host-networking pod will not be affected by taint from hybridnet
		return generatePatchResponseFromPod(req.Object.Raw, ensureTolerationInPod(pod,
			&corev1.Toleration{
				Key:      TaintNodeNetworkUnavailable,
				Operator: corev1.TolerationOpExists,
				Effect:   corev1.TaintEffectNoSchedule,
			}), logger)
	}

	// select 4 networking configs in order as below
	var (
		networkNameStr string
		subnetNameStr  string
		networkTypeStr string
		ipFamilyStr    string

		// elected will be true iff one networking config was assigned
		elected = func() bool {
			return len(networkNameStr) > 0 || len(subnetNameStr) > 0 || len(networkTypeStr) > 0 || len(ipFamilyStr) > 0
		}

		// fetchFromObject will fetch networking configs from k8s objects
		fetchFromObject = func(obj client.Object) error {
			var err error
			if networkNameStr, subnetNameStr, err = webhookutils.SelectNetworkAndSubnetFromObject(ctx, handler.Cache, obj); err != nil {
				return fmt.Errorf("unable to select network and subnet from object %s/%s/%s: %v",
					obj.GetObjectKind().GroupVersionKind().String(), obj.GetNamespace(), obj.GetName(), err)
			}
			networkTypeStr = utils.PickFirstNonEmptyString(obj.GetAnnotations()[constants.AnnotationNetworkType],
				obj.GetLabels()[constants.LabelNetworkType])
			ipFamilyStr = obj.GetAnnotations()[constants.AnnotationIPFamily]
			return nil
		}
	)

	// priority level 1
	// if stateful pods have allocated ips and no need to be reallocated, just
	// reuse the existing network
	if strategy.OwnByStatefulWorkload(pod) {
		var shouldReuse = utils.ParseBoolOrDefault(pod.Annotations[constants.AnnotationIPRetain], strategy.DefaultIPRetain)
		if shouldReuse {
			ipList := &networkingv1.IPInstanceList{}
			if err = handler.Client.List(
				ctx,
				ipList,
				client.InNamespace(pod.Namespace),
				client.MatchingLabels{
					constants.LabelPod: pod.Name,
				}); err != nil {
				return webhookutils.AdmissionErroredWithLog(http.StatusInternalServerError, err, logger)
			}

			// ignore terminating ipInstance
			for i := range ipList.Items {
				if ipList.Items[i].DeletionTimestamp == nil {
					networkNameStr = ipList.Items[i].Spec.Network
					break
				}
			}
		}
	}

	// priority level 2
	// fetch networking configs from pod annotations/labels
	if !elected() {
		pod.Namespace, pod.Name = req.Namespace, req.Name
		if err = fetchFromObject(pod); err != nil {
			return webhookutils.AdmissionErroredWithLog(http.StatusBadRequest, err, logger)
		}
	}

	// priority level 3
	// fetch networking configs from namespace annotations/labels
	if !elected() {
		ns := &corev1.Namespace{}
		if err = handler.Cache.Get(ctx, types.NamespacedName{Name: req.Namespace}, ns); err != nil {
			return webhookutils.AdmissionErroredWithLog(http.StatusInternalServerError,
				fmt.Errorf("unable to get namespace of pod %s/%s: %v", req.Namespace, req.Name, err), logger)
		}
		if err = fetchFromObject(ns); err != nil {
			return webhookutils.AdmissionErroredWithLog(http.StatusBadRequest, err, logger)
		}
	}

	// parsing networking configs
	// TODO: validation
	var networkName = networkNameStr
	var networkType = ipamtypes.ParseNetworkTypeFromString(networkTypeStr)

	var networkNodeSelector map[string]string
	if len(networkName) > 0 {
		network := &networkingv1.Network{}
		if err = handler.Client.Get(ctx, types.NamespacedName{Name: networkName}, network); err != nil {
			return webhookutils.AdmissionErroredWithLog(http.StatusInternalServerError, err, logger)
		}

		// specified network takes higher priority than network type defaulting, if no network type specified
		// from pod, then network type should inherit from network type of specified network from pod
		if len(networkTypeStr) == 0 {
			networkType = ipamtypes.ParseNetworkTypeFromString(string(networkingv1.GetNetworkType(network)))
		}

		networkNodeSelector = network.Spec.NodeSelector
	}

	// persistent specified network and subnet in pod annotations
	patchAnnotationToPod(pod, constants.AnnotationSpecifiedNetwork, networkName)
	patchAnnotationToPod(pod, constants.AnnotationSpecifiedSubnet, subnetNameStr)
	patchAnnotationToPod(pod, constants.AnnotationNetworkType, string(networkType))
	patchAnnotationToPod(pod, constants.AnnotationIPFamily, ipFamilyStr)

	switch networkType {
	case ipamtypes.Underlay:
		if len(networkName) > 0 {
			logger.Info("patch pod with selector of network",
				"namespace", req.Namespace, "name", req.Name, "network", networkName)
			patchSelectorToPod(pod, networkNodeSelector)
		} else {
			logger.Info("patch pod with underlay attachment selector",
				"namespace", req.Namespace, "name", req.Name)
			patchSelectorToPod(pod, map[string]string{
				constants.LabelUnderlayNetworkAttachment: constants.Attached,
			})
		}
		// quota label selector to make sure pod will be scheduled on nodes
		// where capacity of network is enough
		switch ipamtypes.ParseIPFamilyFromString(pod.Annotations[constants.AnnotationIPFamily]) {
		case ipamtypes.IPv4Only:
			patchSelectorToPod(pod, map[string]string{
				constants.LabelIPv4AddressQuota: constants.QuotaNonEmpty,
			})
		case ipamtypes.IPv6Only:
			patchSelectorToPod(pod, map[string]string{
				constants.LabelIPv6AddressQuota: constants.QuotaNonEmpty,
			})
		case ipamtypes.DualStack:
			patchSelectorToPod(pod, map[string]string{
				constants.LabelDualStackAddressQuota: constants.QuotaNonEmpty,
			})
		}
	case ipamtypes.Overlay:
		logger.Info("patch pod with overlay attachment selector",
			"namespace", req.Namespace, "name", req.Name)
		patchSelectorToPod(pod, map[string]string{
			constants.LabelOverlayNetworkAttachment: constants.Attached,
		})
	case ipamtypes.GlobalBGP:
		logger.Info("patch pod with bgp attachment selector",
			"namespace", req.Namespace, "name", req.Name)
		patchSelectorToPod(pod, map[string]string{
			constants.LabelBGPNetworkAttachment: constants.Attached,
		})
	default:
		return webhookutils.AdmissionErroredWithLog(http.StatusBadRequest, fmt.Errorf("unknown network type %s", networkType), logger)
	}

	return generatePatchResponseFromPod(req.Object.Raw, pod, logger)
}

func generatePatchResponseFromPod(original []byte, pod *corev1.Pod, logger logr.Logger) admission.Response {
	marshaled, err := json.Marshal(pod)
	if err != nil {
		return webhookutils.AdmissionErroredWithLog(http.StatusInternalServerError, err, logger)
	}

	return admission.PatchResponseFromRaw(original, marshaled)
}

func patchSelectorToPod(pod *corev1.Pod, selector map[string]string) {
	if pod.Spec.NodeSelector == nil {
		pod.Spec.NodeSelector = selector
		return
	}

	for k, v := range selector {
		pod.Spec.NodeSelector[k] = v
	}
}

func patchAnnotationToPod(pod *corev1.Pod, key, value string) {
	if len(value) == 0 {
		return
	}

	if pod.Annotations == nil {
		pod.Annotations = map[string]string{
			key: value,
		}
		return
	}

	pod.Annotations[key] = value
}

func ensureTolerationInPod(pod *corev1.Pod, tolerations ...*corev1.Toleration) *corev1.Pod {
	for _, toleration := range tolerations {
		var found = false
		for i := range pod.Spec.Tolerations {
			found = found || tolerationMatch(&pod.Spec.Tolerations[i], toleration)
		}
		if !found {
			pod.Spec.Tolerations = append(pod.Spec.Tolerations, *toleration)
		}
	}

	return pod
}

func tolerationMatch(orig, diff *corev1.Toleration) bool {
	return orig.Key == diff.Key &&
		orig.Effect == diff.Effect &&
		orig.Operator == diff.Operator &&
		orig.Value == diff.Value
}

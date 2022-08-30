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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkingv1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworks implements NetworkInterface
type FakeNetworks struct {
	Fake *FakeNetworkingV1
}

var networksResource = schema.GroupVersionResource{Group: "networking", Version: "v1", Resource: "networks"}

var networksKind = schema.GroupVersionKind{Group: "networking", Version: "v1", Kind: "Network"}

// Get takes name of the network, and returns the corresponding network object, and an error if there is any.
func (c *FakeNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networksResource, name), &networkingv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Network), err
}

// List takes label and field selectors, and returns the list of Networks that match those selectors.
func (c *FakeNetworks) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.NetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networksResource, networksKind, opts), &networkingv1.NetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.NetworkList{ListMeta: obj.(*networkingv1.NetworkList).ListMeta}
	for _, item := range obj.(*networkingv1.NetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networks.
func (c *FakeNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networksResource, opts))
}

// Create takes the representation of a network and creates it.  Returns the server's representation of the network, and an error, if there is any.
func (c *FakeNetworks) Create(ctx context.Context, network *networkingv1.Network, opts v1.CreateOptions) (result *networkingv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networksResource, network), &networkingv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Network), err
}

// Update takes the representation of a network and updates it. Returns the server's representation of the network, and an error, if there is any.
func (c *FakeNetworks) Update(ctx context.Context, network *networkingv1.Network, opts v1.UpdateOptions) (result *networkingv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networksResource, network), &networkingv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Network), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworks) UpdateStatus(ctx context.Context, network *networkingv1.Network, opts v1.UpdateOptions) (*networkingv1.Network, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networksResource, "status", network), &networkingv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Network), err
}

// Delete takes name of the network and deletes it. Returns an error if one occurs.
func (c *FakeNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(networksResource, name, opts), &networkingv1.Network{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networksResource, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.NetworkList{})
	return err
}

// Patch applies the patch and returns the patched network.
func (c *FakeNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networksResource, name, pt, data, subresources...), &networkingv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Network), err
}

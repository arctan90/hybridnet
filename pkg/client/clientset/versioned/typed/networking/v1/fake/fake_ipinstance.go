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

// FakeIPInstances implements IPInstanceInterface
type FakeIPInstances struct {
	Fake *FakeNetworkingV1
	ns   string
}

var ipinstancesResource = schema.GroupVersionResource{Group: "networking", Version: "v1", Resource: "ipinstances"}

var ipinstancesKind = schema.GroupVersionKind{Group: "networking", Version: "v1", Kind: "IPInstance"}

// Get takes name of the iPInstance, and returns the corresponding iPInstance object, and an error if there is any.
func (c *FakeIPInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.IPInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipinstancesResource, c.ns, name), &networkingv1.IPInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IPInstance), err
}

// List takes label and field selectors, and returns the list of IPInstances that match those selectors.
func (c *FakeIPInstances) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.IPInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipinstancesResource, ipinstancesKind, c.ns, opts), &networkingv1.IPInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.IPInstanceList{ListMeta: obj.(*networkingv1.IPInstanceList).ListMeta}
	for _, item := range obj.(*networkingv1.IPInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPInstances.
func (c *FakeIPInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipinstancesResource, c.ns, opts))

}

// Create takes the representation of a iPInstance and creates it.  Returns the server's representation of the iPInstance, and an error, if there is any.
func (c *FakeIPInstances) Create(ctx context.Context, iPInstance *networkingv1.IPInstance, opts v1.CreateOptions) (result *networkingv1.IPInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipinstancesResource, c.ns, iPInstance), &networkingv1.IPInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IPInstance), err
}

// Update takes the representation of a iPInstance and updates it. Returns the server's representation of the iPInstance, and an error, if there is any.
func (c *FakeIPInstances) Update(ctx context.Context, iPInstance *networkingv1.IPInstance, opts v1.UpdateOptions) (result *networkingv1.IPInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipinstancesResource, c.ns, iPInstance), &networkingv1.IPInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IPInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPInstances) UpdateStatus(ctx context.Context, iPInstance *networkingv1.IPInstance, opts v1.UpdateOptions) (*networkingv1.IPInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipinstancesResource, "status", c.ns, iPInstance), &networkingv1.IPInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IPInstance), err
}

// Delete takes name of the iPInstance and deletes it. Returns an error if one occurs.
func (c *FakeIPInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipinstancesResource, c.ns, name, opts), &networkingv1.IPInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.IPInstanceList{})
	return err
}

// Patch applies the patch and returns the patched iPInstance.
func (c *FakeIPInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.IPInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipinstancesResource, c.ns, name, pt, data, subresources...), &networkingv1.IPInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.IPInstance), err
}

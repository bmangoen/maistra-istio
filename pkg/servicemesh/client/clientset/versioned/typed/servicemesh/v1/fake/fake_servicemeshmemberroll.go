// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	servicemeshv1 "istio.io/istio/pkg/servicemesh/apis/servicemesh/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceMeshMemberRolls implements ServiceMeshMemberRollInterface
type FakeServiceMeshMemberRolls struct {
	Fake *FakeMaistraV1
	ns   string
}

var servicemeshmemberrollsResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "servicemeshmemberrolls"}

var servicemeshmemberrollsKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceMeshMemberRoll"}

// Get takes name of the serviceMeshMemberRoll, and returns the corresponding serviceMeshMemberRoll object, and an error if there is any.
func (c *FakeServiceMeshMemberRolls) Get(ctx context.Context, name string, options v1.GetOptions) (result *servicemeshv1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshmemberrollsResource, c.ns, name), &servicemeshv1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicemeshv1.ServiceMeshMemberRoll), err
}

// List takes label and field selectors, and returns the list of ServiceMeshMemberRolls that match those selectors.
func (c *FakeServiceMeshMemberRolls) List(ctx context.Context, opts v1.ListOptions) (result *servicemeshv1.ServiceMeshMemberRollList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshmemberrollsResource, servicemeshmemberrollsKind, c.ns, opts), &servicemeshv1.ServiceMeshMemberRollList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicemeshv1.ServiceMeshMemberRollList{ListMeta: obj.(*servicemeshv1.ServiceMeshMemberRollList).ListMeta}
	for _, item := range obj.(*servicemeshv1.ServiceMeshMemberRollList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshMemberRolls.
func (c *FakeServiceMeshMemberRolls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshmemberrollsResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshMemberRoll and creates it.  Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Create(ctx context.Context, serviceMeshMemberRoll *servicemeshv1.ServiceMeshMemberRoll, opts v1.CreateOptions) (result *servicemeshv1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &servicemeshv1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicemeshv1.ServiceMeshMemberRoll), err
}

// Update takes the representation of a serviceMeshMemberRoll and updates it. Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Update(ctx context.Context, serviceMeshMemberRoll *servicemeshv1.ServiceMeshMemberRoll, opts v1.UpdateOptions) (result *servicemeshv1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &servicemeshv1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicemeshv1.ServiceMeshMemberRoll), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshMemberRolls) UpdateStatus(ctx context.Context, serviceMeshMemberRoll *servicemeshv1.ServiceMeshMemberRoll, opts v1.UpdateOptions) (*servicemeshv1.ServiceMeshMemberRoll, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshmemberrollsResource, "status", c.ns, serviceMeshMemberRoll), &servicemeshv1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicemeshv1.ServiceMeshMemberRoll), err
}

// Delete takes name of the serviceMeshMemberRoll and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshMemberRolls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicemeshmemberrollsResource, c.ns, name, opts), &servicemeshv1.ServiceMeshMemberRoll{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshMemberRolls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshmemberrollsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &servicemeshv1.ServiceMeshMemberRollList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshMemberRoll.
func (c *FakeServiceMeshMemberRolls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *servicemeshv1.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshmemberrollsResource, c.ns, name, pt, data, subresources...), &servicemeshv1.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicemeshv1.ServiceMeshMemberRoll), err
}
// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComponentPackages implements ComponentPackageInterface
type FakeComponentPackages struct {
	Fake *FakeBundleV1alpha1
	ns   string
}

var componentpackagesResource = schema.GroupVersionResource{Group: "bundle.gke.io", Version: "v1alpha1", Resource: "componentpackages"}

var componentpackagesKind = schema.GroupVersionKind{Group: "bundle.gke.io", Version: "v1alpha1", Kind: "ComponentPackage"}

// Get takes name of the componentPackage, and returns the corresponding componentPackage object, and an error if there is any.
func (c *FakeComponentPackages) Get(name string, options v1.GetOptions) (result *v1alpha1.ComponentPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(componentpackagesResource, c.ns, name), &v1alpha1.ComponentPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentPackage), err
}

// List takes label and field selectors, and returns the list of ComponentPackages that match those selectors.
func (c *FakeComponentPackages) List(opts v1.ListOptions) (result *v1alpha1.ComponentPackageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(componentpackagesResource, componentpackagesKind, c.ns, opts), &v1alpha1.ComponentPackageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComponentPackageList{}
	for _, item := range obj.(*v1alpha1.ComponentPackageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentPackages.
func (c *FakeComponentPackages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(componentpackagesResource, c.ns, opts))

}

// Create takes the representation of a componentPackage and creates it.  Returns the server's representation of the componentPackage, and an error, if there is any.
func (c *FakeComponentPackages) Create(componentPackage *v1alpha1.ComponentPackage) (result *v1alpha1.ComponentPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(componentpackagesResource, c.ns, componentPackage), &v1alpha1.ComponentPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentPackage), err
}

// Update takes the representation of a componentPackage and updates it. Returns the server's representation of the componentPackage, and an error, if there is any.
func (c *FakeComponentPackages) Update(componentPackage *v1alpha1.ComponentPackage) (result *v1alpha1.ComponentPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(componentpackagesResource, c.ns, componentPackage), &v1alpha1.ComponentPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentPackage), err
}

// Delete takes name of the componentPackage and deletes it. Returns an error if one occurs.
func (c *FakeComponentPackages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(componentpackagesResource, c.ns, name), &v1alpha1.ComponentPackage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponentPackages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(componentpackagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComponentPackageList{})
	return err
}

// Patch applies the patch and returns the patched componentPackage.
func (c *FakeComponentPackages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComponentPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(componentpackagesResource, c.ns, name, data, subresources...), &v1alpha1.ComponentPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentPackage), err
}
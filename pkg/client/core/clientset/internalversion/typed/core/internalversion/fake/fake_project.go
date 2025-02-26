/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeCore
}

var projectsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "", Resource: "projects"}

var projectsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "", Kind: "Project"}

// Get takes name of the project, and returns the corresponding project object, and an error if there is any.
func (c *FakeProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *core.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectsResource, name), &core.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.Project), err
}

// List takes label and field selectors, and returns the list of Projects that match those selectors.
func (c *FakeProjects) List(ctx context.Context, opts v1.ListOptions) (result *core.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectsResource, projectsKind, opts), &core.ProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.ProjectList{ListMeta: obj.(*core.ProjectList).ListMeta}
	for _, item := range obj.(*core.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectsResource, opts))
}

// Create takes the representation of a project and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Create(ctx context.Context, project *core.Project, opts v1.CreateOptions) (result *core.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectsResource, project), &core.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.Project), err
}

// Update takes the representation of a project and updates it. Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Update(ctx context.Context, project *core.Project, opts v1.UpdateOptions) (result *core.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectsResource, project), &core.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.Project), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjects) UpdateStatus(ctx context.Context, project *core.Project, opts v1.UpdateOptions) (*core.Project, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectsResource, "status", project), &core.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.Project), err
}

// Delete takes name of the project and deletes it. Returns an error if one occurs.
func (c *FakeProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(projectsResource, name, opts), &core.Project{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectsResource, listOpts)

	_, err := c.Fake.Invokes(action, &core.ProjectList{})
	return err
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *core.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectsResource, name, pt, data, subresources...), &core.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.Project), err
}

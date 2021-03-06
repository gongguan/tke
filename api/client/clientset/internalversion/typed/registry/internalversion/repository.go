/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	registry "tkestack.io/tke/api/registry"
)

// RepositoriesGetter has a method to return a RepositoryInterface.
// A group's client should implement this interface.
type RepositoriesGetter interface {
	Repositories(namespace string) RepositoryInterface
}

// RepositoryInterface has methods to work with Repository resources.
type RepositoryInterface interface {
	Create(*registry.Repository) (*registry.Repository, error)
	Update(*registry.Repository) (*registry.Repository, error)
	UpdateStatus(*registry.Repository) (*registry.Repository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*registry.Repository, error)
	List(opts v1.ListOptions) (*registry.RepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *registry.Repository, err error)
	RepositoryExpansion
}

// repositories implements RepositoryInterface
type repositories struct {
	client rest.Interface
	ns     string
}

// newRepositories returns a Repositories
func newRepositories(c *RegistryClient, namespace string) *repositories {
	return &repositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the repository, and returns the corresponding repository object, and an error if there is any.
func (c *repositories) Get(name string, options v1.GetOptions) (result *registry.Repository, err error) {
	result = &registry.Repository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Repositories that match those selectors.
func (c *repositories) List(opts v1.ListOptions) (result *registry.RepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &registry.RepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested repositories.
func (c *repositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a repository and creates it.  Returns the server's representation of the repository, and an error, if there is any.
func (c *repositories) Create(repository *registry.Repository) (result *registry.Repository, err error) {
	result = &registry.Repository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("repositories").
		Body(repository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a repository and updates it. Returns the server's representation of the repository, and an error, if there is any.
func (c *repositories) Update(repository *registry.Repository) (result *registry.Repository, err error) {
	result = &registry.Repository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("repositories").
		Name(repository.Name).
		Body(repository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *repositories) UpdateStatus(repository *registry.Repository) (result *registry.Repository, err error) {
	result = &registry.Repository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("repositories").
		Name(repository.Name).
		SubResource("status").
		Body(repository).
		Do().
		Into(result)
	return
}

// Delete takes name of the repository and deletes it. Returns an error if one occurs.
func (c *repositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("repositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *repositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched repository.
func (c *repositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *registry.Repository, err error) {
	result = &registry.Repository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("repositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

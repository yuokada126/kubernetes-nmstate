/*
Copyright 2018 The nmstate Authors.

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

package v1

import (
	v1 "github.com/nmstate/k8s-node-net-conf/pkg/apis/nmstate.io/v1"
	scheme "github.com/nmstate/k8s-node-net-conf/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeNetConfPoliciesGetter has a method to return a NodeNetConfPolicyInterface.
// A group's client should implement this interface.
type NodeNetConfPoliciesGetter interface {
	NodeNetConfPolicies(namespace string) NodeNetConfPolicyInterface
}

// NodeNetConfPolicyInterface has methods to work with NodeNetConfPolicy resources.
type NodeNetConfPolicyInterface interface {
	Create(*v1.NodeNetConfPolicy) (*v1.NodeNetConfPolicy, error)
	Update(*v1.NodeNetConfPolicy) (*v1.NodeNetConfPolicy, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.NodeNetConfPolicy, error)
	List(opts metav1.ListOptions) (*v1.NodeNetConfPolicyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NodeNetConfPolicy, err error)
	NodeNetConfPolicyExpansion
}

// nodeNetConfPolicies implements NodeNetConfPolicyInterface
type nodeNetConfPolicies struct {
	client rest.Interface
	ns     string
}

// newNodeNetConfPolicies returns a NodeNetConfPolicies
func newNodeNetConfPolicies(c *NmstateV1Client, namespace string) *nodeNetConfPolicies {
	return &nodeNetConfPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeNetConfPolicy, and returns the corresponding nodeNetConfPolicy object, and an error if there is any.
func (c *nodeNetConfPolicies) Get(name string, options metav1.GetOptions) (result *v1.NodeNetConfPolicy, err error) {
	result = &v1.NodeNetConfPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeNetConfPolicies that match those selectors.
func (c *nodeNetConfPolicies) List(opts metav1.ListOptions) (result *v1.NodeNetConfPolicyList, err error) {
	result = &v1.NodeNetConfPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeNetConfPolicies.
func (c *nodeNetConfPolicies) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nodeNetConfPolicy and creates it.  Returns the server's representation of the nodeNetConfPolicy, and an error, if there is any.
func (c *nodeNetConfPolicies) Create(nodeNetConfPolicy *v1.NodeNetConfPolicy) (result *v1.NodeNetConfPolicy, err error) {
	result = &v1.NodeNetConfPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		Body(nodeNetConfPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeNetConfPolicy and updates it. Returns the server's representation of the nodeNetConfPolicy, and an error, if there is any.
func (c *nodeNetConfPolicies) Update(nodeNetConfPolicy *v1.NodeNetConfPolicy) (result *v1.NodeNetConfPolicy, err error) {
	result = &v1.NodeNetConfPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		Name(nodeNetConfPolicy.Name).
		Body(nodeNetConfPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeNetConfPolicy and deletes it. Returns an error if one occurs.
func (c *nodeNetConfPolicies) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeNetConfPolicies) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeNetConfPolicy.
func (c *nodeNetConfPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NodeNetConfPolicy, err error) {
	result = &v1.NodeNetConfPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodenetconfpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

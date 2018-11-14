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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nmstate/k8s-node-net-conf/pkg/apis/nmstate.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeNetConfPolicyLister helps list NodeNetConfPolicies.
type NodeNetConfPolicyLister interface {
	// List lists all NodeNetConfPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.NodeNetConfPolicy, err error)
	// NodeNetConfPolicies returns an object that can list and get NodeNetConfPolicies.
	NodeNetConfPolicies(namespace string) NodeNetConfPolicyNamespaceLister
	NodeNetConfPolicyListerExpansion
}

// nodeNetConfPolicyLister implements the NodeNetConfPolicyLister interface.
type nodeNetConfPolicyLister struct {
	indexer cache.Indexer
}

// NewNodeNetConfPolicyLister returns a new NodeNetConfPolicyLister.
func NewNodeNetConfPolicyLister(indexer cache.Indexer) NodeNetConfPolicyLister {
	return &nodeNetConfPolicyLister{indexer: indexer}
}

// List lists all NodeNetConfPolicies in the indexer.
func (s *nodeNetConfPolicyLister) List(selector labels.Selector) (ret []*v1.NodeNetConfPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeNetConfPolicy))
	})
	return ret, err
}

// NodeNetConfPolicies returns an object that can list and get NodeNetConfPolicies.
func (s *nodeNetConfPolicyLister) NodeNetConfPolicies(namespace string) NodeNetConfPolicyNamespaceLister {
	return nodeNetConfPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeNetConfPolicyNamespaceLister helps list and get NodeNetConfPolicies.
type NodeNetConfPolicyNamespaceLister interface {
	// List lists all NodeNetConfPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.NodeNetConfPolicy, err error)
	// Get retrieves the NodeNetConfPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1.NodeNetConfPolicy, error)
	NodeNetConfPolicyNamespaceListerExpansion
}

// nodeNetConfPolicyNamespaceLister implements the NodeNetConfPolicyNamespaceLister
// interface.
type nodeNetConfPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeNetConfPolicies in the indexer for a given namespace.
func (s nodeNetConfPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.NodeNetConfPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeNetConfPolicy))
	})
	return ret, err
}

// Get retrieves the NodeNetConfPolicy from the indexer for a given namespace and name.
func (s nodeNetConfPolicyNamespaceLister) Get(name string) (*v1.NodeNetConfPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodenetconfpolicy"), name)
	}
	return obj.(*v1.NodeNetConfPolicy), nil
}

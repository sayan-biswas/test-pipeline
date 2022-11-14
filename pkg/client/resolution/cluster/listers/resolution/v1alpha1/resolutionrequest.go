/*
Copyright 2020 The Tekton Authors

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

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"	
	"github.com/kcp-dev/logicalcluster/v2"
	
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	resolutionv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/resolution/v1alpha1"
	resolutionv1alpha1listers "github.com/tektoncd/pipeline/pkg/client/resolution/listers/resolution/v1alpha1"
	)

// ResolutionRequestClusterLister can list ResolutionRequests across all workspaces, or scope down to a ResolutionRequestLister for one workspace.
type ResolutionRequestClusterLister interface {
	List(selector labels.Selector) (ret []*resolutionv1alpha1.ResolutionRequest, err error)
Cluster(cluster logicalcluster.Name)resolutionv1alpha1listers.ResolutionRequestLister
}

type resolutionRequestClusterLister struct {
	indexer cache.Indexer
}

// NewResolutionRequestClusterLister returns a new ResolutionRequestClusterLister.
func NewResolutionRequestClusterLister(indexer cache.Indexer) *resolutionRequestClusterLister {
	return &resolutionRequestClusterLister{indexer: indexer}
}

// List lists all ResolutionRequests in the indexer across all workspaces.
func (s *resolutionRequestClusterLister) List(selector labels.Selector) (ret []*resolutionv1alpha1.ResolutionRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*resolutionv1alpha1.ResolutionRequest))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ResolutionRequests.
func (s *resolutionRequestClusterLister) Cluster(cluster logicalcluster.Name)resolutionv1alpha1listers.ResolutionRequestLister {
return &resolutionRequestLister{indexer: s.indexer, cluster: cluster}
}

// resolutionRequestLister implements the resolutionv1alpha1listers.ResolutionRequestLister interface.
type resolutionRequestLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ResolutionRequests in the indexer for a workspace.
func (s *resolutionRequestLister) List(selector labels.Selector) (ret []*resolutionv1alpha1.ResolutionRequest, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*resolutionv1alpha1.ResolutionRequest)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// ResolutionRequests returns an object that can list and get ResolutionRequests in one namespace.
func (s *resolutionRequestLister) ResolutionRequests(namespace string) resolutionv1alpha1listers.ResolutionRequestNamespaceLister {
return &resolutionRequestNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// resolutionRequestNamespaceLister implements the resolutionv1alpha1listers.ResolutionRequestNamespaceLister interface.
type resolutionRequestNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all ResolutionRequests in the indexer for a given workspace and namespace.
func (s *resolutionRequestNamespaceLister) List(selector labels.Selector) (ret []*resolutionv1alpha1.ResolutionRequest, err error) {
	selectAll := selector == nil || selector.Empty()

	var list []interface{}
	if s.namespace == metav1.NamespaceAll {
		list, err = s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	} else {
		list, err = s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	}
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*resolutionv1alpha1.ResolutionRequest)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the ResolutionRequest from the indexer for a given workspace, namespace and name.
func (s *resolutionRequestNamespaceLister) Get(name string) (*resolutionv1alpha1.ResolutionRequest, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(resolutionv1alpha1.Resource("ResolutionRequest"), name)
	}
	return obj.(*resolutionv1alpha1.ResolutionRequest), nil
}

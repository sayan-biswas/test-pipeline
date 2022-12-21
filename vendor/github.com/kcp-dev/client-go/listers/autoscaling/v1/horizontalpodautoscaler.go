//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	autoscalingv1listers "k8s.io/client-go/listers/autoscaling/v1"
	"k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerClusterLister can list HorizontalPodAutoscalers across all workspaces, or scope down to a HorizontalPodAutoscalerLister for one workspace.
// All objects returned here must be treated as read-only.
type HorizontalPodAutoscalerClusterLister interface {
	// List lists all HorizontalPodAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error)
	// Cluster returns a lister that can list and get HorizontalPodAutoscalers in one workspace.
	Cluster(cluster logicalcluster.Name) autoscalingv1listers.HorizontalPodAutoscalerLister
	HorizontalPodAutoscalerClusterListerExpansion
}

type horizontalPodAutoscalerClusterLister struct {
	indexer cache.Indexer
}

// NewHorizontalPodAutoscalerClusterLister returns a new HorizontalPodAutoscalerClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewHorizontalPodAutoscalerClusterLister(indexer cache.Indexer) *horizontalPodAutoscalerClusterLister {
	return &horizontalPodAutoscalerClusterLister{indexer: indexer}
}

// List lists all HorizontalPodAutoscalers in the indexer across all workspaces.
func (s *horizontalPodAutoscalerClusterLister) List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*autoscalingv1.HorizontalPodAutoscaler))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get HorizontalPodAutoscalers.
func (s *horizontalPodAutoscalerClusterLister) Cluster(cluster logicalcluster.Name) autoscalingv1listers.HorizontalPodAutoscalerLister {
	return &horizontalPodAutoscalerLister{indexer: s.indexer, cluster: cluster}
}

// horizontalPodAutoscalerLister implements the autoscalingv1listers.HorizontalPodAutoscalerLister interface.
type horizontalPodAutoscalerLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all HorizontalPodAutoscalers in the indexer for a workspace.
func (s *horizontalPodAutoscalerLister) List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*autoscalingv1.HorizontalPodAutoscaler))
	})
	return ret, err
}

// HorizontalPodAutoscalers returns an object that can list and get HorizontalPodAutoscalers in one namespace.
func (s *horizontalPodAutoscalerLister) HorizontalPodAutoscalers(namespace string) autoscalingv1listers.HorizontalPodAutoscalerNamespaceLister {
	return &horizontalPodAutoscalerNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// horizontalPodAutoscalerNamespaceLister implements the autoscalingv1listers.HorizontalPodAutoscalerNamespaceLister interface.
type horizontalPodAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all HorizontalPodAutoscalers in the indexer for a given workspace and namespace.
func (s *horizontalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*autoscalingv1.HorizontalPodAutoscaler, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*autoscalingv1.HorizontalPodAutoscaler))
	})
	return ret, err
}

// Get retrieves the HorizontalPodAutoscaler from the indexer for a given workspace, namespace and name.
func (s *horizontalPodAutoscalerNamespaceLister) Get(name string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(autoscalingv1.Resource("HorizontalPodAutoscaler"), name)
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), nil
}
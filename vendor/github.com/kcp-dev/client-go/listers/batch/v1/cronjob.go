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

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	batchv1listers "k8s.io/client-go/listers/batch/v1"
	"k8s.io/client-go/tools/cache"
)

// CronJobClusterLister can list CronJobs across all workspaces, or scope down to a CronJobLister for one workspace.
// All objects returned here must be treated as read-only.
type CronJobClusterLister interface {
	// List lists all CronJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*batchv1.CronJob, err error)
	// Cluster returns a lister that can list and get CronJobs in one workspace.
	Cluster(cluster logicalcluster.Name) batchv1listers.CronJobLister
	CronJobClusterListerExpansion
}

type cronJobClusterLister struct {
	indexer cache.Indexer
}

// NewCronJobClusterLister returns a new CronJobClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewCronJobClusterLister(indexer cache.Indexer) *cronJobClusterLister {
	return &cronJobClusterLister{indexer: indexer}
}

// List lists all CronJobs in the indexer across all workspaces.
func (s *cronJobClusterLister) List(selector labels.Selector) (ret []*batchv1.CronJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*batchv1.CronJob))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get CronJobs.
func (s *cronJobClusterLister) Cluster(cluster logicalcluster.Name) batchv1listers.CronJobLister {
	return &cronJobLister{indexer: s.indexer, cluster: cluster}
}

// cronJobLister implements the batchv1listers.CronJobLister interface.
type cronJobLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all CronJobs in the indexer for a workspace.
func (s *cronJobLister) List(selector labels.Selector) (ret []*batchv1.CronJob, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.CronJob))
	})
	return ret, err
}

// CronJobs returns an object that can list and get CronJobs in one namespace.
func (s *cronJobLister) CronJobs(namespace string) batchv1listers.CronJobNamespaceLister {
	return &cronJobNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// cronJobNamespaceLister implements the batchv1listers.CronJobNamespaceLister interface.
type cronJobNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all CronJobs in the indexer for a given workspace and namespace.
func (s *cronJobNamespaceLister) List(selector labels.Selector) (ret []*batchv1.CronJob, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.CronJob))
	})
	return ret, err
}

// Get retrieves the CronJob from the indexer for a given workspace, namespace and name.
func (s *cronJobNamespaceLister) Get(name string) (*batchv1.CronJob, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batchv1.Resource("CronJob"), name)
	}
	return obj.(*batchv1.CronJob), nil
}

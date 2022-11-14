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
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamrbacv1informers "k8s.io/client-go/informers/rbac/v1"
	upstreamrbacv1listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	rbacv1listers "github.com/kcp-dev/client-go/listers/rbac/v1"
)

// RoleBindingClusterInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamrbacv1informers.RoleBindingInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() rbacv1listers.RoleBindingClusterLister
}

type roleBindingClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleBindingClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRoleBindingClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1().RoleBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1().RoleBindings().Watch(context.TODO(), options)
			},
		},
		&rbacv1.RoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *roleBindingClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *roleBindingClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&rbacv1.RoleBinding{}, f.defaultInformer)
}

func (f *roleBindingClusterInformer) Lister() rbacv1listers.RoleBindingClusterLister {
	return rbacv1listers.NewRoleBindingClusterLister(f.Informer().GetIndexer())
}

func (f *roleBindingClusterInformer) Cluster(cluster logicalcluster.Name) upstreamrbacv1informers.RoleBindingInformer {
	return &roleBindingInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type roleBindingInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamrbacv1listers.RoleBindingLister
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *roleBindingInformer) Lister() upstreamrbacv1listers.RoleBindingLister {
	return f.lister
}

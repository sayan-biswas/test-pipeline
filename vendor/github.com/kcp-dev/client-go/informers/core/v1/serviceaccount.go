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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamcorev1informers "k8s.io/client-go/informers/core/v1"
	upstreamcorev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	corev1listers "github.com/kcp-dev/client-go/listers/core/v1"
)

// ServiceAccountClusterInformer provides access to a shared informer and lister for
// ServiceAccounts.
type ServiceAccountClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamcorev1informers.ServiceAccountInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() corev1listers.ServiceAccountClusterLister
}

type serviceAccountClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewServiceAccountClusterInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceAccountClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredServiceAccountClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredServiceAccountClusterInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceAccountClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts().Watch(context.TODO(), options)
			},
		},
		&corev1.ServiceAccount{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceAccountClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredServiceAccountClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *serviceAccountClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&corev1.ServiceAccount{}, f.defaultInformer)
}

func (f *serviceAccountClusterInformer) Lister() corev1listers.ServiceAccountClusterLister {
	return corev1listers.NewServiceAccountClusterLister(f.Informer().GetIndexer())
}

func (f *serviceAccountClusterInformer) Cluster(cluster logicalcluster.Name) upstreamcorev1informers.ServiceAccountInformer {
	return &serviceAccountInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type serviceAccountInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamcorev1listers.ServiceAccountLister
}

func (f *serviceAccountInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *serviceAccountInformer) Lister() upstreamcorev1listers.ServiceAccountLister {
	return f.lister
}

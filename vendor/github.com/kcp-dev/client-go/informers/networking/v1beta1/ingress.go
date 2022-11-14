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

package v1beta1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamnetworkingv1beta1informers "k8s.io/client-go/informers/networking/v1beta1"
	upstreamnetworkingv1beta1listers "k8s.io/client-go/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	networkingv1beta1listers "github.com/kcp-dev/client-go/listers/networking/v1beta1"
)

// IngressClusterInformer provides access to a shared informer and lister for
// Ingresses.
type IngressClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamnetworkingv1beta1informers.IngressInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() networkingv1beta1listers.IngressClusterLister
}

type ingressClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIngressClusterInformer constructs a new informer for Ingress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredIngressClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIngressClusterInformer constructs a new informer for Ingress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().Ingresses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().Ingresses().Watch(context.TODO(), options)
			},
		},
		&networkingv1beta1.Ingress{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredIngressClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *ingressClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&networkingv1beta1.Ingress{}, f.defaultInformer)
}

func (f *ingressClusterInformer) Lister() networkingv1beta1listers.IngressClusterLister {
	return networkingv1beta1listers.NewIngressClusterLister(f.Informer().GetIndexer())
}

func (f *ingressClusterInformer) Cluster(cluster logicalcluster.Name) upstreamnetworkingv1beta1informers.IngressInformer {
	return &ingressInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type ingressInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamnetworkingv1beta1listers.IngressLister
}

func (f *ingressInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *ingressInformer) Lister() upstreamnetworkingv1beta1listers.IngressLister {
	return f.lister
}

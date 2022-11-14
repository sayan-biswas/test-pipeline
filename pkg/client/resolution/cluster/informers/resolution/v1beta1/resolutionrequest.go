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

package v1beta1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	resolutionv1beta1 "github.com/tektoncd/pipeline/pkg/apis/resolution/v1beta1"
	resolutionv1beta1listers "github.com/tektoncd/pipeline/pkg/client/resolution/cluster/listers/resolution/v1beta1"
	upstreamresolutionv1beta1listers "github.com/tektoncd/pipeline/pkg/client/resolution/listers/resolution/v1beta1"
	upstreamresolutionv1beta1informers "github.com/tektoncd/pipeline/pkg/client/resolution/informers/externalversions/resolution/v1beta1"
	clientset "github.com/tektoncd/pipeline/pkg/client/resolution/cluster/clientset"

	"github.com/tektoncd/pipeline/pkg/client/resolution/cluster/informers/internalinterfaces"
)

// ResolutionRequestClusterInformer provides access to a shared informer and lister for
// ResolutionRequests.
type ResolutionRequestClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamresolutionv1beta1informers.ResolutionRequestInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() resolutionv1beta1listers.ResolutionRequestClusterLister
}

type resolutionRequestClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResolutionRequestClusterInformer constructs a new informer for ResolutionRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResolutionRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResolutionRequestClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResolutionRequestClusterInformer constructs a new informer for ResolutionRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResolutionRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResolutionV1beta1().ResolutionRequests().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResolutionV1beta1().ResolutionRequests().Watch(context.TODO(), options)
			},
		},
		&resolutionv1beta1.ResolutionRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *resolutionRequestClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResolutionRequestClusterInformer(client, resyncPeriod, cache.Indexers{
			kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
			kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc,}, 
		f.tweakListOptions,
	)
}

func (f *resolutionRequestClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&resolutionv1beta1.ResolutionRequest{}, f.defaultInformer)
}

func (f *resolutionRequestClusterInformer) Lister() resolutionv1beta1listers.ResolutionRequestClusterLister {
	return resolutionv1beta1listers.NewResolutionRequestClusterLister(f.Informer().GetIndexer())
}

func (f *resolutionRequestClusterInformer) Cluster(cluster logicalcluster.Name) upstreamresolutionv1beta1informers.ResolutionRequestInformer {
	return &resolutionRequestInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type resolutionRequestInformer struct {
	informer cache.SharedIndexInformer
	lister upstreamresolutionv1beta1listers.ResolutionRequestLister
}

func (f *resolutionRequestInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *resolutionRequestInformer) Lister() upstreamresolutionv1beta1listers.ResolutionRequestLister {
	return f.lister
}

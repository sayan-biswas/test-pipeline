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
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	tektonv1alpha1listers "github.com/tektoncd/pipeline/pkg/client/resource/cluster/listers/tekton/v1alpha1"
	upstreamtektonv1alpha1listers "github.com/tektoncd/pipeline/pkg/client/resource/listers/resource/v1alpha1"
	upstreamtektonv1alpha1informers "github.com/tektoncd/pipeline/pkg/client/resource/informers/externalversions/resource/v1alpha1"
	clientset "github.com/tektoncd/pipeline/pkg/client/resource/cluster/clientset"

	"github.com/tektoncd/pipeline/pkg/client/resource/cluster/informers/internalinterfaces"
)

// PipelineResourceClusterInformer provides access to a shared informer and lister for
// PipelineResources.
type PipelineResourceClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamtektonv1alpha1informers.PipelineResourceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() tektonv1alpha1listers.PipelineResourceClusterLister
}

type pipelineResourceClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPipelineResourceClusterInformer constructs a new informer for PipelineResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineResourceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPipelineResourceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineResourceClusterInformer constructs a new informer for PipelineResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineResourceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1alpha1().PipelineResources().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1alpha1().PipelineResources().Watch(context.TODO(), options)
			},
		},
		&tektonv1alpha1.PipelineResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineResourceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPipelineResourceClusterInformer(client, resyncPeriod, cache.Indexers{
			kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
			kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc,}, 
		f.tweakListOptions,
	)
}

func (f *pipelineResourceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&tektonv1alpha1.PipelineResource{}, f.defaultInformer)
}

func (f *pipelineResourceClusterInformer) Lister() tektonv1alpha1listers.PipelineResourceClusterLister {
	return tektonv1alpha1listers.NewPipelineResourceClusterLister(f.Informer().GetIndexer())
}

func (f *pipelineResourceClusterInformer) Cluster(cluster logicalcluster.Name) upstreamtektonv1alpha1informers.PipelineResourceInformer {
	return &pipelineResourceInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type pipelineResourceInformer struct {
	informer cache.SharedIndexInformer
	lister upstreamtektonv1alpha1listers.PipelineResourceLister
}

func (f *pipelineResourceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *pipelineResourceInformer) Lister() upstreamtektonv1alpha1listers.PipelineResourceLister {
	return f.lister
}

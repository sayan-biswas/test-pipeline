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

package v1

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

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonv1listers "github.com/tektoncd/pipeline/pkg/client/cluster/listers/tekton/v1"
	upstreamtektonv1listers "github.com/tektoncd/pipeline/pkg/client/listers/pipeline/v1"
	upstreamtektonv1informers "github.com/tektoncd/pipeline/pkg/client/informers/externalversions/pipeline/v1"
	clientset "github.com/tektoncd/pipeline/pkg/client/cluster/clientset"

	"github.com/tektoncd/pipeline/pkg/client/cluster/informers/internalinterfaces"
)

// PipelineClusterInformer provides access to a shared informer and lister for
// Pipelines.
type PipelineClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamtektonv1informers.PipelineInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() tektonv1listers.PipelineClusterLister
}

type pipelineClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPipelineClusterInformer constructs a new informer for Pipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPipelineClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineClusterInformer constructs a new informer for Pipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1().Pipelines().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1().Pipelines().Watch(context.TODO(), options)
			},
		},
		&tektonv1.Pipeline{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPipelineClusterInformer(client, resyncPeriod, cache.Indexers{
			kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
			kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc,}, 
		f.tweakListOptions,
	)
}

func (f *pipelineClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&tektonv1.Pipeline{}, f.defaultInformer)
}

func (f *pipelineClusterInformer) Lister() tektonv1listers.PipelineClusterLister {
	return tektonv1listers.NewPipelineClusterLister(f.Informer().GetIndexer())
}

func (f *pipelineClusterInformer) Cluster(cluster logicalcluster.Name) upstreamtektonv1informers.PipelineInformer {
	return &pipelineInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type pipelineInformer struct {
	informer cache.SharedIndexInformer
	lister upstreamtektonv1listers.PipelineLister
}

func (f *pipelineInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *pipelineInformer) Lister() upstreamtektonv1listers.PipelineLister {
	return f.lister
}
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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"


	tektonv1alpha1client "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
)

// RunsClusterGetter has a method to return a RunClusterInterface.
// A group's cluster client should implement this interface.
type RunsClusterGetter interface {
	Runs() RunClusterInterface
}

// RunClusterInterface can operate on Runs across all clusters,
// or scope down to one cluster and return a RunsNamespacer.
type RunClusterInterface interface {
	Cluster(logicalcluster.Name) RunsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*tektonv1alpha1.RunList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type runsClusterInterface struct {
	clientCache kcpclient.Cache[*tektonv1alpha1client.TektonV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *runsClusterInterface) Cluster(name logicalcluster.Name) RunsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &runsNamespacer{clientCache: c.clientCache, name: name}
}


// List returns the entire collection of all Runs across all clusters. 
func (c *runsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1alpha1.RunList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Runs(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Runs across all clusters.
func (c *runsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Runs(metav1.NamespaceAll).Watch(ctx, opts)
}
// RunsNamespacer can scope to objects within a namespace, returning a tektonv1alpha1client.RunInterface.
type RunsNamespacer interface {
	Namespace(string) tektonv1alpha1client.RunInterface
}

type runsNamespacer struct {
	clientCache kcpclient.Cache[*tektonv1alpha1client.TektonV1alpha1Client]
	name logicalcluster.Name
}

func (n *runsNamespacer) Namespace(namespace string) tektonv1alpha1client.RunInterface {
	return n.clientCache.ClusterOrDie(n.name).Runs(namespace)
}

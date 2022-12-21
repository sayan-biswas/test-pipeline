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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"


	tektonv1beta1client "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
)

// CustomRunsClusterGetter has a method to return a CustomRunClusterInterface.
// A group's cluster client should implement this interface.
type CustomRunsClusterGetter interface {
	CustomRuns() CustomRunClusterInterface
}

// CustomRunClusterInterface can operate on CustomRuns across all clusters,
// or scope down to one cluster and return a CustomRunsNamespacer.
type CustomRunClusterInterface interface {
	Cluster(logicalcluster.Name) CustomRunsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.CustomRunList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type customRunsClusterInterface struct {
	clientCache kcpclient.Cache[*tektonv1beta1client.TektonV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *customRunsClusterInterface) Cluster(name logicalcluster.Name) CustomRunsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &customRunsNamespacer{clientCache: c.clientCache, name: name}
}


// List returns the entire collection of all CustomRuns across all clusters. 
func (c *customRunsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.CustomRunList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).CustomRuns(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all CustomRuns across all clusters.
func (c *customRunsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).CustomRuns(metav1.NamespaceAll).Watch(ctx, opts)
}
// CustomRunsNamespacer can scope to objects within a namespace, returning a tektonv1beta1client.CustomRunInterface.
type CustomRunsNamespacer interface {
	Namespace(string) tektonv1beta1client.CustomRunInterface
}

type customRunsNamespacer struct {
	clientCache kcpclient.Cache[*tektonv1beta1client.TektonV1beta1Client]
	name logicalcluster.Name
}

func (n *customRunsNamespacer) Namespace(namespace string) tektonv1beta1client.CustomRunInterface {
	return n.clientCache.ClusterOrDie(n.name).CustomRuns(namespace)
}
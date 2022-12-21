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

// ClusterTasksClusterGetter has a method to return a ClusterTaskClusterInterface.
// A group's cluster client should implement this interface.
type ClusterTasksClusterGetter interface {
	ClusterTasks() ClusterTaskClusterInterface
}

// ClusterTaskClusterInterface can operate on ClusterTasks across all clusters,
// or scope down to one cluster and return a tektonv1beta1client.ClusterTaskInterface.
type ClusterTaskClusterInterface interface {
	Cluster(logicalcluster.Name) tektonv1beta1client.ClusterTaskInterface
	List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.ClusterTaskList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type clusterTasksClusterInterface struct {
	clientCache kcpclient.Cache[*tektonv1beta1client.TektonV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterTasksClusterInterface) Cluster(name logicalcluster.Name) tektonv1beta1client.ClusterTaskInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).ClusterTasks()
}


// List returns the entire collection of all ClusterTasks across all clusters. 
func (c *clusterTasksClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.ClusterTaskList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterTasks().List(ctx, opts)
}

// Watch begins to watch all ClusterTasks across all clusters.
func (c *clusterTasksClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterTasks().Watch(ctx, opts)
}
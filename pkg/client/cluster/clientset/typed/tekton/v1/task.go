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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"


	tektonv1client "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
)

// TasksClusterGetter has a method to return a TaskClusterInterface.
// A group's cluster client should implement this interface.
type TasksClusterGetter interface {
	Tasks() TaskClusterInterface
}

// TaskClusterInterface can operate on Tasks across all clusters,
// or scope down to one cluster and return a TasksNamespacer.
type TaskClusterInterface interface {
	Cluster(logicalcluster.Name) TasksNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*tektonv1.TaskList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type tasksClusterInterface struct {
	clientCache kcpclient.Cache[*tektonv1client.TektonV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *tasksClusterInterface) Cluster(name logicalcluster.Name) TasksNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &tasksNamespacer{clientCache: c.clientCache, name: name}
}


// List returns the entire collection of all Tasks across all clusters. 
func (c *tasksClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1.TaskList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Tasks(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Tasks across all clusters.
func (c *tasksClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Tasks(metav1.NamespaceAll).Watch(ctx, opts)
}
// TasksNamespacer can scope to objects within a namespace, returning a tektonv1client.TaskInterface.
type TasksNamespacer interface {
	Namespace(string) tektonv1client.TaskInterface
}

type tasksNamespacer struct {
	clientCache kcpclient.Cache[*tektonv1client.TektonV1Client]
	name logicalcluster.Name
}

func (n *tasksNamespacer) Namespace(namespace string) tektonv1client.TaskInterface {
	return n.clientCache.ClusterOrDie(n.name).Tasks(namespace)
}

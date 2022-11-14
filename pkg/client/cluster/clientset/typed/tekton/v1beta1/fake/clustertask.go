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
	"github.com/kcp-dev/logicalcluster/v2"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	"k8s.io/apimachinery/pkg/watch"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/testing"

	"k8s.io/apimachinery/pkg/types"

	tektonv1beta1client "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
)

var clusterTasksResource = schema.GroupVersionResource{Group: "tekton.dev", Version: "v1beta1", Resource: "clustertasks"}
var clusterTasksKind = schema.GroupVersionKind{Group: "tekton.dev", Version: "v1beta1", Kind: "ClusterTask"}

type clusterTasksClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterTasksClusterClient) Cluster(cluster logicalcluster.Name) tektonv1beta1client.ClusterTaskInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterTasksClient{Fake: c.Fake, Cluster: cluster}
}


// List takes label and field selectors, and returns the list of ClusterTasks that match those selectors across all clusters.
func (c *clusterTasksClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.ClusterTaskList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTasksResource, clusterTasksKind, logicalcluster.Wildcard, opts), &tektonv1beta1.ClusterTaskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tektonv1beta1.ClusterTaskList{ListMeta: obj.(*tektonv1beta1.ClusterTaskList).ListMeta}
	for _, item := range obj.(*tektonv1beta1.ClusterTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterTasks across all clusters.
func (c *clusterTasksClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTasksResource, logicalcluster.Wildcard, opts))
}
type clusterTasksClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
	
}


func (c *clusterTasksClient) Create(ctx context.Context, clusterTask *tektonv1beta1.ClusterTask, opts metav1.CreateOptions) (*tektonv1beta1.ClusterTask, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterTasksResource, c.Cluster, clusterTask), &tektonv1beta1.ClusterTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tektonv1beta1.ClusterTask), err
}

func (c *clusterTasksClient) Update(ctx context.Context, clusterTask *tektonv1beta1.ClusterTask, opts metav1.UpdateOptions) (*tektonv1beta1.ClusterTask, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterTasksResource, c.Cluster, clusterTask), &tektonv1beta1.ClusterTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tektonv1beta1.ClusterTask), err
}

func (c *clusterTasksClient) UpdateStatus(ctx context.Context, clusterTask *tektonv1beta1.ClusterTask, opts metav1.UpdateOptions) (*tektonv1beta1.ClusterTask, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterTasksResource, c.Cluster, "status", clusterTask), &tektonv1beta1.ClusterTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tektonv1beta1.ClusterTask), err
}

func (c *clusterTasksClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterTasksResource, c.Cluster, name, opts), &tektonv1beta1.ClusterTask{})
	return err
}

func (c *clusterTasksClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterTasksResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &tektonv1beta1.ClusterTaskList{})
	return err
}

func (c *clusterTasksClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*tektonv1beta1.ClusterTask, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterTasksResource, c.Cluster, name), &tektonv1beta1.ClusterTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tektonv1beta1.ClusterTask), err
}

// List takes label and field selectors, and returns the list of ClusterTasks that match those selectors.
func (c *clusterTasksClient) List(ctx context.Context, opts metav1.ListOptions) (*tektonv1beta1.ClusterTaskList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTasksResource, clusterTasksKind, c.Cluster, opts), &tektonv1beta1.ClusterTaskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tektonv1beta1.ClusterTaskList{ListMeta: obj.(*tektonv1beta1.ClusterTaskList).ListMeta}
	for _, item := range obj.(*tektonv1beta1.ClusterTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterTasksClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTasksResource, c.Cluster, opts))
}

func (c *clusterTasksClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*tektonv1beta1.ClusterTask, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTasksResource, c.Cluster, name, pt, data, subresources...), &tektonv1beta1.ClusterTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tektonv1beta1.ClusterTask), err
}

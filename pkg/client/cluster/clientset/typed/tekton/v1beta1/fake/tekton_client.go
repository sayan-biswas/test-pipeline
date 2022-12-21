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
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/rest"
	kcptektonv1beta1 "github.com/tektoncd/pipeline/pkg/client/cluster/clientset/typed/tekton/v1beta1"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
)

var _ kcptektonv1beta1.TektonV1beta1ClusterInterface = (*TektonV1beta1ClusterClient)(nil)

type TektonV1beta1ClusterClient struct {
	*kcptesting.Fake 
}

func (c *TektonV1beta1ClusterClient) Cluster(cluster logicalcluster.Name) tektonv1beta1.TektonV1beta1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &TektonV1beta1Client{Fake: c.Fake, Cluster: cluster}
}


func (c *TektonV1beta1ClusterClient) ClusterTasks() kcptektonv1beta1.ClusterTaskClusterInterface {
	return &clusterTasksClusterClient{Fake: c.Fake}
}

func (c *TektonV1beta1ClusterClient) CustomRuns() kcptektonv1beta1.CustomRunClusterInterface {
	return &customRunsClusterClient{Fake: c.Fake}
}

func (c *TektonV1beta1ClusterClient) Pipelines() kcptektonv1beta1.PipelineClusterInterface {
	return &pipelinesClusterClient{Fake: c.Fake}
}

func (c *TektonV1beta1ClusterClient) PipelineRuns() kcptektonv1beta1.PipelineRunClusterInterface {
	return &pipelineRunsClusterClient{Fake: c.Fake}
}

func (c *TektonV1beta1ClusterClient) Tasks() kcptektonv1beta1.TaskClusterInterface {
	return &tasksClusterClient{Fake: c.Fake}
}

func (c *TektonV1beta1ClusterClient) TaskRuns() kcptektonv1beta1.TaskRunClusterInterface {
	return &taskRunsClusterClient{Fake: c.Fake}
}
var _ tektonv1beta1.TektonV1beta1Interface = (*TektonV1beta1Client)(nil)

type TektonV1beta1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *TektonV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}


func (c *TektonV1beta1Client) ClusterTasks() tektonv1beta1.ClusterTaskInterface {
	return &clusterTasksClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *TektonV1beta1Client) CustomRuns(namespace string) tektonv1beta1.CustomRunInterface {
	return &customRunsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *TektonV1beta1Client) Pipelines(namespace string) tektonv1beta1.PipelineInterface {
	return &pipelinesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *TektonV1beta1Client) PipelineRuns(namespace string) tektonv1beta1.PipelineRunInterface {
	return &pipelineRunsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *TektonV1beta1Client) Tasks(namespace string) tektonv1beta1.TaskInterface {
	return &tasksClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *TektonV1beta1Client) TaskRuns(namespace string) tektonv1beta1.TaskRunInterface {
	return &taskRunsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}
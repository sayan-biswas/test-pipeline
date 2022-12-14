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
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/rest"
	kcptektonv1alpha1 "github.com/tektoncd/pipeline/pkg/client/cluster/clientset/typed/tekton/v1alpha1"
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
)

var _ kcptektonv1alpha1.TektonV1alpha1ClusterInterface = (*TektonV1alpha1ClusterClient)(nil)

type TektonV1alpha1ClusterClient struct {
	*kcptesting.Fake 
}

func (c *TektonV1alpha1ClusterClient) Cluster(cluster logicalcluster.Name) tektonv1alpha1.TektonV1alpha1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &TektonV1alpha1Client{Fake: c.Fake, Cluster: cluster}
}


func (c *TektonV1alpha1ClusterClient) Runs() kcptektonv1alpha1.RunClusterInterface {
	return &runsClusterClient{Fake: c.Fake}
}
var _ tektonv1alpha1.TektonV1alpha1Interface = (*TektonV1alpha1Client)(nil)

type TektonV1alpha1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *TektonV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}


func (c *TektonV1alpha1Client) Runs(namespace string) tektonv1alpha1.RunInterface {
	return &runsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

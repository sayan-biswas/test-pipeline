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
	"github.com/tektoncd/pipeline/pkg/client/cluster/informers/internalinterfaces"
)

type ClusterInterface interface {
// ClusterTasks returns a ClusterTaskClusterInformer
	ClusterTasks() ClusterTaskClusterInformer
// CustomRuns returns a CustomRunClusterInformer
	CustomRuns() CustomRunClusterInformer
// Pipelines returns a PipelineClusterInformer
	Pipelines() PipelineClusterInformer
// PipelineRuns returns a PipelineRunClusterInformer
	PipelineRuns() PipelineRunClusterInformer
// Tasks returns a TaskClusterInformer
	Tasks() TaskClusterInformer
// TaskRuns returns a TaskRunClusterInformer
	TaskRuns() TaskRunClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// ClusterTasks returns a ClusterTaskClusterInformer
func (v *version) ClusterTasks() ClusterTaskClusterInformer {
	return &clusterTaskClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// CustomRuns returns a CustomRunClusterInformer
func (v *version) CustomRuns() CustomRunClusterInformer {
	return &customRunClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// Pipelines returns a PipelineClusterInformer
func (v *version) Pipelines() PipelineClusterInformer {
	return &pipelineClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// PipelineRuns returns a PipelineRunClusterInformer
func (v *version) PipelineRuns() PipelineRunClusterInformer {
	return &pipelineRunClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// Tasks returns a TaskClusterInformer
func (v *version) Tasks() TaskClusterInformer {
	return &taskClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
// TaskRuns returns a TaskRunClusterInformer
func (v *version) TaskRuns() TaskRunClusterInformer {
	return &taskRunClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
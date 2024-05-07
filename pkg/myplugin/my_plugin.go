/*
Copyright 2020 The Kubernetes Authors.

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

package myplugin

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	// Name of the plugin used in the plugin registry and configurations.
	Name = "MyPlugin"
)

// MyPlugin is a PostFilter plugin.
type MyPlugin struct {
	fh         framework.Handle
	podLister  corelisters.PodLister
	nodeLister corelisters.NodeLister
}

var _ framework.PostFilterPlugin = &MyPlugin{}

// Name returns name of the plugin. It is used in logs, etc.
func (pl *MyPlugin) Name() string {
	return Name
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, fh framework.Handle) (framework.Plugin, error) {
	pl := MyPlugin{
		fh:         fh,
		podLister:  fh.SharedInformerFactory().Core().V1().Pods().Lister(),
		nodeLister: fh.SharedInformerFactory().Core().V1().Nodes().Lister(),
	}
	return &pl, nil
}

// PostFilter invoked at the postFilter extension point.
func (pl *MyPlugin) PostFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, m framework.NodeToStatusMap) (*framework.PostFilterResult, *framework.Status) {
	// This happens when the pod is not eligible for preemption or extenders filtered all candidates.
	klog.Infof("Log di Prova")
	return &framework.PostFilterResult{}, framework.NewStatus(framework.Success)
}

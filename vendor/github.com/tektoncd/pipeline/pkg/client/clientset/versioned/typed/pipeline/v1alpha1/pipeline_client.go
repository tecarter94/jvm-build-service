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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type TektonV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterTasksGetter
	ConditionsGetter
	PipelinesGetter
	PipelineRunsGetter
	RunsGetter
	TasksGetter
	TaskRunsGetter
}

// TektonV1alpha1Client is used to interact with features provided by the tekton.dev group.
type TektonV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TektonV1alpha1Client) ClusterTasks() ClusterTaskInterface {
	return newClusterTasks(c)
}

func (c *TektonV1alpha1Client) Conditions(namespace string) ConditionInterface {
	return newConditions(c, namespace)
}

func (c *TektonV1alpha1Client) Pipelines(namespace string) PipelineInterface {
	return newPipelines(c, namespace)
}

func (c *TektonV1alpha1Client) PipelineRuns(namespace string) PipelineRunInterface {
	return newPipelineRuns(c, namespace)
}

func (c *TektonV1alpha1Client) Runs(namespace string) RunInterface {
	return newRuns(c, namespace)
}

func (c *TektonV1alpha1Client) Tasks(namespace string) TaskInterface {
	return newTasks(c, namespace)
}

func (c *TektonV1alpha1Client) TaskRuns(namespace string) TaskRunInterface {
	return newTaskRuns(c, namespace)
}

// NewForConfig creates a new TektonV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*TektonV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TektonV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TektonV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TektonV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TektonV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TektonV1alpha1Client {
	return &TektonV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TektonV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

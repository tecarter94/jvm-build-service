/*
Copyright 2021-2022 Red Hat, Inc.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/redhat-appstudio/jvm-build-service/pkg/apis/jvmbuildservice/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TektonWrapperLister helps list TektonWrappers.
// All objects returned here must be treated as read-only.
type TektonWrapperLister interface {
	// List lists all TektonWrappers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TektonWrapper, err error)
	// TektonWrappers returns an object that can list and get TektonWrappers.
	TektonWrappers(namespace string) TektonWrapperNamespaceLister
	TektonWrapperListerExpansion
}

// tektonWrapperLister implements the TektonWrapperLister interface.
type tektonWrapperLister struct {
	indexer cache.Indexer
}

// NewTektonWrapperLister returns a new TektonWrapperLister.
func NewTektonWrapperLister(indexer cache.Indexer) TektonWrapperLister {
	return &tektonWrapperLister{indexer: indexer}
}

// List lists all TektonWrappers in the indexer.
func (s *tektonWrapperLister) List(selector labels.Selector) (ret []*v1alpha1.TektonWrapper, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TektonWrapper))
	})
	return ret, err
}

// TektonWrappers returns an object that can list and get TektonWrappers.
func (s *tektonWrapperLister) TektonWrappers(namespace string) TektonWrapperNamespaceLister {
	return tektonWrapperNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TektonWrapperNamespaceLister helps list and get TektonWrappers.
// All objects returned here must be treated as read-only.
type TektonWrapperNamespaceLister interface {
	// List lists all TektonWrappers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TektonWrapper, err error)
	// Get retrieves the TektonWrapper from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TektonWrapper, error)
	TektonWrapperNamespaceListerExpansion
}

// tektonWrapperNamespaceLister implements the TektonWrapperNamespaceLister
// interface.
type tektonWrapperNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TektonWrappers in the indexer for a given namespace.
func (s tektonWrapperNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TektonWrapper, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TektonWrapper))
	})
	return ret, err
}

// Get retrieves the TektonWrapper from the indexer for a given namespace and name.
func (s tektonWrapperNamespaceLister) Get(name string) (*v1alpha1.TektonWrapper, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tektonwrapper"), name)
	}
	return obj.(*v1alpha1.TektonWrapper), nil
}
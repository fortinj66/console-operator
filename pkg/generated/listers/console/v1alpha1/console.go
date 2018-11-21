// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/console-operator/pkg/apis/console/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsoleLister helps list Consoles.
type ConsoleLister interface {
	// List lists all Consoles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Console, err error)
	// Consoles returns an object that can list and get Consoles.
	Consoles(namespace string) ConsoleNamespaceLister
	ConsoleListerExpansion
}

// consoleLister implements the ConsoleLister interface.
type consoleLister struct {
	indexer cache.Indexer
}

// NewConsoleLister returns a new ConsoleLister.
func NewConsoleLister(indexer cache.Indexer) ConsoleLister {
	return &consoleLister{indexer: indexer}
}

// List lists all Consoles in the indexer.
func (s *consoleLister) List(selector labels.Selector) (ret []*v1alpha1.Console, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Console))
	})
	return ret, err
}

// Consoles returns an object that can list and get Consoles.
func (s *consoleLister) Consoles(namespace string) ConsoleNamespaceLister {
	return consoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsoleNamespaceLister helps list and get Consoles.
type ConsoleNamespaceLister interface {
	// List lists all Consoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Console, err error)
	// Get retrieves the Console from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Console, error)
	ConsoleNamespaceListerExpansion
}

// consoleNamespaceLister implements the ConsoleNamespaceLister
// interface.
type consoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Consoles in the indexer for a given namespace.
func (s consoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Console, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Console))
	})
	return ret, err
}

// Get retrieves the Console from the indexer for a given namespace and name.
func (s consoleNamespaceLister) Get(name string) (*v1alpha1.Console, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("console"), name)
	}
	return obj.(*v1alpha1.Console), nil
}

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "kpt.dev/configsync/pkg/api/configmanagement/v1"
)

// RepoLister helps list Repos.
// All objects returned here must be treated as read-only.
type RepoLister interface {
	// List lists all Repos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Repo, err error)
	// Get retrieves the Repo from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Repo, error)
	RepoListerExpansion
}

// repoLister implements the RepoLister interface.
type repoLister struct {
	indexer cache.Indexer
}

// NewRepoLister returns a new RepoLister.
func NewRepoLister(indexer cache.Indexer) RepoLister {
	return &repoLister{indexer: indexer}
}

// List lists all Repos in the indexer.
func (s *repoLister) List(selector labels.Selector) (ret []*v1.Repo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Repo))
	})
	return ret, err
}

// Get retrieves the Repo from the index for a given name.
func (s *repoLister) Get(name string) (*v1.Repo, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("repo"), name)
	}
	return obj.(*v1.Repo), nil
}

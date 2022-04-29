// Code generated by informer-gen. DO NOT EDIT.

package informer

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1 "kpt.dev/configsync/pkg/api/configmanagement/v1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=configmanagement.gke.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("clusterconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().ClusterConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusterselectors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().ClusterSelectors().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("hierarchyconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().HierarchyConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("namespaceconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().NamespaceConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("namespaceselectors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().NamespaceSelectors().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("repos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().Repos().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("syncs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Configmanagement().V1().Syncs().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

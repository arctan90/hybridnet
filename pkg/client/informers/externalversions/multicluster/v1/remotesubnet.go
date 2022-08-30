/*
Copyright 2021 The Hybridnet Authors.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	multiclusterv1 "github.com/alibaba/hybridnet/pkg/apis/multicluster/v1"
	versioned "github.com/alibaba/hybridnet/pkg/client/clientset/versioned"
	internalinterfaces "github.com/alibaba/hybridnet/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/alibaba/hybridnet/pkg/client/listers/multicluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RemoteSubnetInformer provides access to a shared informer and lister for
// RemoteSubnets.
type RemoteSubnetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RemoteSubnetLister
}

type remoteSubnetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRemoteSubnetInformer constructs a new informer for RemoteSubnet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRemoteSubnetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRemoteSubnetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRemoteSubnetInformer constructs a new informer for RemoteSubnet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRemoteSubnetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1().RemoteSubnets().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1().RemoteSubnets().Watch(context.TODO(), options)
			},
		},
		&multiclusterv1.RemoteSubnet{},
		resyncPeriod,
		indexers,
	)
}

func (f *remoteSubnetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRemoteSubnetInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *remoteSubnetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiclusterv1.RemoteSubnet{}, f.defaultInformer)
}

func (f *remoteSubnetInformer) Lister() v1.RemoteSubnetLister {
	return v1.NewRemoteSubnetLister(f.Informer().GetIndexer())
}

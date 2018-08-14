/*
Copyright 2017 Aspen Mesh Authors.

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

// Package vetter contains interfaces which vetters should implement.
// All vetter(s) packages must export
//  func NewVetter(factory vetter.ResourceListGetter) *newVetter
// where newVetter implements the Vetter interface described below.
package vetter

import (
	"github.com/aspenmesh/istio-client-go/pkg/client/informers/externalversions"
	apiv1 "github.com/aspenmesh/istio-vet/api/v1"
	"k8s.io/client-go/informers"
)

// Vetter interface is implemented by vetters.
type Vetter interface {
	// Vet returns the list of notes generated by a vetter.
	// It takes kubernetes client interface as the argument.
	Vet() ([]*apiv1.Note, error)

	// Info returns information like name, version about the vetter.
	Info() *apiv1.Info
}

// ResourceListGetter is used by vetters to register for and list resources they are interested in.
// This currently exposes the Informer interface, but that should not be used.
// Only the "Lister()" interfaces should be considered public.
type ResourceListGetter interface {
	K8s() informers.SharedInformerFactory
	Istio() externalversions.SharedInformerFactory
}

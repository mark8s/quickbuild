/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QuickBuildSpec defines the desired state of QuickBuild
type QuickBuildSpec struct {
	// application name
	// If not specified, a random name will be used
	// +kubebuilder:validation:MaxLength:=50
	Name string `json:"name,omitempty"`

	// application namespace
	// If not specified, a random name will be used
	// +kubebuilder:validation:MaxLength:=64
	Namespace string `json:"namespace,omitempty"`

	// container image
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type:int
	Image string `json:"image,omitempty"`

	// application port
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type:int
	Port int32 `json:"port"`

	// application replicas
	// +kubebuilder:validation:Minimum:=1
	Replicas *int32 `json:"replicas,omitempty"`

	// whether to create service
	EnableService bool `json:"enableService"`
}

// QuickBuildStatus defines the observed state of QuickBuild
type QuickBuildStatus struct {
	// application running status
	Status string `json:"status"`

	// service ip
	ServiceIP string `json:"serviceIp"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=qb

// QuickBuild is the Schema for the quickbuilds API
type QuickBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuickBuildSpec   `json:"spec,omitempty"`
	Status QuickBuildStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// QuickBuildList contains a list of QuickBuild
type QuickBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuickBuild `json:"items"`
}

func init() {
	SchemeBuilder.Register(&QuickBuild{}, &QuickBuildList{})
}

// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceMeshMemberRoll is the Schema for the servicemeshmemberrolls API
// +k8s:openapi-gen=true
type ServiceMeshMemberRoll struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceMeshMemberRollSpec   `json:"spec,omitempty"`
	Status ServiceMeshMemberRollStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceMeshMemberRollList contains a list of ServiceMeshMemberRoll
type ServiceMeshMemberRollList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceMeshMemberRoll `json:"items"`
}

// ServiceMeshMemberRollSpec defines the members of the mesh
type ServiceMeshMemberRollSpec struct {
	Members []string `json:"members,omitempty"`
}

// ServiceMeshMemberRollStatus contains the state last used to reconcile the list
type ServiceMeshMemberRollStatus struct {
	ObservedGeneration    int64    `json:"observedGeneration,omitempty"`
	ServiceMeshGeneration int64    `json:"meshGeneration,omitempty"`
	ConfiguredMembers     []string `json:"configuredMembers,omitempty"`
}
/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpecAttachments struct {
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
}

type VolumeSpecSnapshotPolicies struct {
	SnapshotCount     *int64  `json:"snapshotCount" tf:"snapshot_count"`
	SnapshotFrequency *string `json:"snapshotFrequency" tf:"snapshot_frequency"`
}

type VolumeSpec struct {
	State *VolumeSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VolumeSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Attachments []VolumeSpecAttachments `json:"attachments,omitempty" tf:"attachments"`
	// +optional
	BillingCycle *string `json:"billingCycle,omitempty" tf:"billing_cycle"`
	// +optional
	Created *string `json:"created,omitempty" tf:"created"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Facility    *string `json:"facility" tf:"facility"`
	// +optional
	Locked *bool `json:"locked,omitempty" tf:"locked"`
	// +optional
	Name      *string `json:"name,omitempty" tf:"name"`
	Plan      *string `json:"plan" tf:"plan"`
	ProjectID *string `json:"projectID" tf:"project_id"`
	Size      *int64  `json:"size" tf:"size"`
	// +optional
	SnapshotPolicies []VolumeSpecSnapshotPolicies `json:"snapshotPolicies,omitempty" tf:"snapshot_policies"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Updated *string `json:"updated,omitempty" tf:"updated"`
}

type VolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeList is a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []Volume `json:"items,omitempty"`
}

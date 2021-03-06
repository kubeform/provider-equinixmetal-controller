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

type VlanAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VlanAttachmentSpec   `json:"spec,omitempty"`
	Status            VlanAttachmentStatus `json:"status,omitempty"`
}

type VlanAttachmentSpec struct {
	State *VlanAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource VlanAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VlanAttachmentSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of device to be assigned to the VLAN
	DeviceID *string `json:"deviceID" tf:"device_id"`
	// Add port back to the bond when this resource is removed. Default is false
	// +optional
	ForceBond *bool `json:"forceBond,omitempty" tf:"force_bond"`
	// Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use depends_on pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above
	// +optional
	Native *bool `json:"native,omitempty" tf:"native"`
	// UUID of device port
	// +optional
	PortID *string `json:"portID,omitempty" tf:"port_id"`
	// Name of network port to be assigned to the VLAN
	PortName *string `json:"portName" tf:"port_name"`
	// UUID of VLAN API resource
	// +optional
	VlanID *string `json:"vlanID,omitempty" tf:"vlan_id"`
	// VXLAN Network Identifier, integer
	VlanVnid *int64 `json:"vlanVnid" tf:"vlan_vnid"`
}

type VlanAttachmentStatus struct {
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

// VlanAttachmentList is a list of VlanAttachments
type VlanAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VlanAttachment CRD objects
	Items []VlanAttachment `json:"items,omitempty"`
}

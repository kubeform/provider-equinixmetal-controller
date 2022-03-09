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

type IpBlock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpBlockSpec   `json:"spec,omitempty"`
	Status            IpBlockStatus `json:"status,omitempty"`
}

type IpBlockSpec struct {
	State *IpBlockSpecResource `json:"state,omitempty" tf:"-"`

	Resource IpBlockSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IpBlockSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// Address family as integer (4 or 6)
	// +optional
	AddressFamily *int64 `json:"addressFamily,omitempty" tf:"address_family"`
	// Length of CIDR prefix of the block as integer
	// +optional
	Cidr *int64 `json:"cidr,omitempty" tf:"cidr"`
	// +optional
	CidrNotation *string `json:"cidrNotation,omitempty" tf:"cidr_notation"`
	// Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may be helpful for self-managed IPAM. The object must be valid JSON.
	// +optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data"`
	// Arbitrary description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with metro
	// +optional
	Facility *string `json:"facility,omitempty" tf:"facility"`
	// +optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway"`
	// Flag indicating whether IP block is global, i.e. assignable in any location
	// +optional
	Global *bool `json:"global,omitempty" tf:"global"`
	// +optional
	Manageable *bool `json:"manageable,omitempty" tf:"manageable"`
	// +optional
	Management *bool `json:"management,omitempty" tf:"management"`
	// Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with facility
	// +optional
	Metro *string `json:"metro,omitempty" tf:"metro"`
	// Mask in decimal notation, e.g. 255.255.255.0
	// +optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask"`
	// Network IP address portion of the block specification
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// The metal project ID where to allocate the address block
	ProjectID *string `json:"projectID" tf:"project_id"`
	// Flag indicating whether IP block is addressable from the Internet
	// +optional
	Public *bool `json:"public,omitempty" tf:"public"`
	// The number of allocated /32 addresses, a power of 2
	Quantity *int64 `json:"quantity" tf:"quantity"`
	// Tags attached to the reserved block
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Either global_ipv4 or public_ipv4, defaults to public_ipv4 for backward compatibility
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
	// +optional
	WaitForState *string `json:"waitForState,omitempty" tf:"wait_for_state"`
}

type IpBlockStatus struct {
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

// IpBlockList is a list of IpBlocks
type IpBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IpBlock CRD objects
	Items []IpBlock `json:"items,omitempty"`
}

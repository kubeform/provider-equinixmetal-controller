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

type Project struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec,omitempty"`
	Status            ProjectStatus `json:"status,omitempty"`
}

type ProjectSpecBgpConfig struct {
	// Autonomous System Number for local BGP deployment
	Asn *int64 `json:"asn" tf:"asn"`
	// "local" or "global", the local is likely to be usable immediately, the global will need to be review by Equinix Metal engineers
	DeploymentType *string `json:"deploymentType" tf:"deployment_type"`
	// The maximum number of route filters allowed per server
	// +optional
	MaxPrefix *int64 `json:"maxPrefix,omitempty" tf:"max_prefix"`
	// Password for BGP session in plaintext (not a checksum)
	// +optional
	Md5 *string `json:"-" sensitive:"true" tf:"md5"`
	// Status of BGP configuration in the project
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type ProjectSpec struct {
	State *ProjectSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProjectSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProjectSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
	// +optional
	BackendTransfer *bool `json:"backendTransfer,omitempty" tf:"backend_transfer"`
	// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/)
	// +optional
	BgpConfig *ProjectSpecBgpConfig `json:"bgpConfig,omitempty" tf:"bgp_config"`
	// The timestamp for when the project was created
	// +optional
	Created *string `json:"created,omitempty" tf:"created"`
	// The name of the project
	Name *string `json:"name" tf:"name"`
	// The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account
	// +optional
	OrganizationID *string `json:"organizationID,omitempty" tf:"organization_id"`
	// The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with organization_id, or default)
	// +optional
	PaymentMethodID *string `json:"paymentMethodID,omitempty" tf:"payment_method_id"`
	// The timestamp for the last time the project was updated
	// +optional
	Updated *string `json:"updated,omitempty" tf:"updated"`
}

type ProjectStatus struct {
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

// ProjectList is a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Project CRD objects
	Items []Project `json:"items,omitempty"`
}

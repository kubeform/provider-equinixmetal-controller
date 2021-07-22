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

type MarketRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MarketRequestSpec   `json:"spec,omitempty"`
	Status            MarketRequestStatus `json:"status,omitempty"`
}

type MarketRequestSpecInstanceParameters struct {
	// +optional
	AlwaysPxe    *bool   `json:"alwaysPxe,omitempty" tf:"always_pxe"`
	BillingCycle *string `json:"billingCycle" tf:"billing_cycle"`
	// +optional
	Customdata *string `json:"customdata,omitempty" tf:"customdata"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Features []string `json:"features,omitempty" tf:"features"`
	Hostname *string  `json:"hostname" tf:"hostname"`
	// +optional
	IpxeScriptURL *string `json:"ipxeScriptURL,omitempty" tf:"ipxe_script_url"`
	// +optional
	Locked          *bool   `json:"locked,omitempty" tf:"locked"`
	OperatingSystem *string `json:"operatingSystem" tf:"operating_system"`
	Plan            *string `json:"plan" tf:"plan"`
	// +optional
	ProjectSSHKeys []string `json:"projectSSHKeys,omitempty" tf:"project_ssh_keys"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TermintationTime *string `json:"termintationTime,omitempty" tf:"termintation_time"`
	// +optional
	UserSSHKeys []string `json:"userSSHKeys,omitempty" tf:"user_ssh_keys"`
	// +optional
	Userdata *string `json:"userdata,omitempty" tf:"userdata"`
}

type MarketRequestSpec struct {
	State *MarketRequestSpecResource `json:"state,omitempty" tf:"-"`

	Resource MarketRequestSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MarketRequestSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number devices to be created
	DevicesMax *int64 `json:"devicesMax" tf:"devices_max"`
	// Miniumum number devices to be created
	DevicesMin *int64 `json:"devicesMin" tf:"devices_min"`
	// Facility IDs where devices should be created
	// +optional
	Facilities []string `json:"facilities,omitempty" tf:"facilities"`
	// Parameters for devices provisioned from this request. You can find the parameter description from the [metal_device doc](device.md)
	InstanceParameters *MarketRequestSpecInstanceParameters `json:"instanceParameters" tf:"instance_parameters"`
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice *float64 `json:"maxBidPrice" tf:"max_bid_price"`
	// Metro where devices should be created
	// +optional
	Metro *string `json:"metro,omitempty" tf:"metro"`
	// Project ID
	ProjectID *string `json:"projectID" tf:"project_id"`
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	// +optional
	WaitForDevices *bool `json:"waitForDevices,omitempty" tf:"wait_for_devices"`
}

type MarketRequestStatus struct {
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

// MarketRequestList is a list of MarketRequests
type MarketRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MarketRequest CRD objects
	Items []MarketRequest `json:"items,omitempty"`
}
/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	bgpv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/bgp/v1alpha1"
	connectionv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/connection/v1alpha1"
	devicev1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/device/v1alpha1"
	gatewayv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/gateway/v1alpha1"
	ipv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/ip/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/organization/v1alpha1"
	portv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/port/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/project/v1alpha1"
	reservedv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/reserved/v1alpha1"
	spotv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/spot/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/ssh/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/user/v1alpha1"
	virtualv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/virtual/v1alpha1"
	vlanv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/vlan/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/volume/v1alpha1"
	"kubeform.dev/provider-equinixmetal-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var (
	allJsonIt = map[schema.GroupVersionResource]Data{
		{
			Group:    "bgp.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "sessions",
		}: {
			JsonIt:       controllers.GetJSONItr(bgpv1alpha1.GetEncoder(), bgpv1alpha1.GetDecoder()),
			ResourceType: "metal_bgp_session",
		},
		{
			Group:    "connection.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "connections",
		}: {
			JsonIt:       controllers.GetJSONItr(connectionv1alpha1.GetEncoder(), connectionv1alpha1.GetDecoder()),
			ResourceType: "metal_connection",
		},
		{
			Group:    "device.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "devices",
		}: {
			JsonIt:       controllers.GetJSONItr(devicev1alpha1.GetEncoder(), devicev1alpha1.GetDecoder()),
			ResourceType: "metal_device",
		},
		{
			Group:    "device.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "networktypes",
		}: {
			JsonIt:       controllers.GetJSONItr(devicev1alpha1.GetEncoder(), devicev1alpha1.GetDecoder()),
			ResourceType: "metal_device_network_type",
		},
		{
			Group:    "gateway.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "gateways",
		}: {
			JsonIt:       controllers.GetJSONItr(gatewayv1alpha1.GetEncoder(), gatewayv1alpha1.GetDecoder()),
			ResourceType: "metal_gateway",
		},
		{
			Group:    "ip.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "attachments",
		}: {
			JsonIt:       controllers.GetJSONItr(ipv1alpha1.GetEncoder(), ipv1alpha1.GetDecoder()),
			ResourceType: "metal_ip_attachment",
		},
		{
			Group:    "organization.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "organizations",
		}: {
			JsonIt:       controllers.GetJSONItr(organizationv1alpha1.GetEncoder(), organizationv1alpha1.GetDecoder()),
			ResourceType: "metal_organization",
		},
		{
			Group:    "port.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "vlanattachments",
		}: {
			JsonIt:       controllers.GetJSONItr(portv1alpha1.GetEncoder(), portv1alpha1.GetDecoder()),
			ResourceType: "metal_port_vlan_attachment",
		},
		{
			Group:    "project.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "projects",
		}: {
			JsonIt:       controllers.GetJSONItr(projectv1alpha1.GetEncoder(), projectv1alpha1.GetDecoder()),
			ResourceType: "metal_project",
		},
		{
			Group:    "project.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "apikeys",
		}: {
			JsonIt:       controllers.GetJSONItr(projectv1alpha1.GetEncoder(), projectv1alpha1.GetDecoder()),
			ResourceType: "metal_project_api_key",
		},
		{
			Group:    "project.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "sshkeys",
		}: {
			JsonIt:       controllers.GetJSONItr(projectv1alpha1.GetEncoder(), projectv1alpha1.GetDecoder()),
			ResourceType: "metal_project_ssh_key",
		},
		{
			Group:    "reserved.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "ipblocks",
		}: {
			JsonIt:       controllers.GetJSONItr(reservedv1alpha1.GetEncoder(), reservedv1alpha1.GetDecoder()),
			ResourceType: "metal_reserved_ip_block",
		},
		{
			Group:    "spot.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "marketrequests",
		}: {
			JsonIt:       controllers.GetJSONItr(spotv1alpha1.GetEncoder(), spotv1alpha1.GetDecoder()),
			ResourceType: "metal_spot_market_request",
		},
		{
			Group:    "ssh.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "keys",
		}: {
			JsonIt:       controllers.GetJSONItr(sshv1alpha1.GetEncoder(), sshv1alpha1.GetDecoder()),
			ResourceType: "metal_ssh_key",
		},
		{
			Group:    "user.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "apikeys",
		}: {
			JsonIt:       controllers.GetJSONItr(userv1alpha1.GetEncoder(), userv1alpha1.GetDecoder()),
			ResourceType: "metal_user_api_key",
		},
		{
			Group:    "virtual.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "circuits",
		}: {
			JsonIt:       controllers.GetJSONItr(virtualv1alpha1.GetEncoder(), virtualv1alpha1.GetDecoder()),
			ResourceType: "metal_virtual_circuit",
		},
		{
			Group:    "vlan.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "vlans",
		}: {
			JsonIt:       controllers.GetJSONItr(vlanv1alpha1.GetEncoder(), vlanv1alpha1.GetDecoder()),
			ResourceType: "metal_vlan",
		},
		{
			Group:    "volume.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "volumes",
		}: {
			JsonIt:       controllers.GetJSONItr(volumev1alpha1.GetEncoder(), volumev1alpha1.GetDecoder()),
			ResourceType: "metal_volume",
		},
		{
			Group:    "volume.equinixmetal.kubeform.com",
			Version:  "v1alpha1",
			Resource: "attachments",
		}: {
			JsonIt:       controllers.GetJSONItr(volumev1alpha1.GetEncoder(), volumev1alpha1.GetDecoder()),
			ResourceType: "metal_volume_attachment",
		},
	}
)

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}

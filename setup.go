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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	equinixmetal "github.com/equinix/terraform-provider-metal/metal"
	"github.com/gobuffalo/flect"
	auditlib "go.bytebuilders.dev/audit/lib"
	arv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
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
	controllersbgp "kubeform.dev/provider-equinixmetal-controller/controllers/bgp"
	controllersconnection "kubeform.dev/provider-equinixmetal-controller/controllers/connection"
	controllersdevice "kubeform.dev/provider-equinixmetal-controller/controllers/device"
	controllersgateway "kubeform.dev/provider-equinixmetal-controller/controllers/gateway"
	controllersip "kubeform.dev/provider-equinixmetal-controller/controllers/ip"
	controllersorganization "kubeform.dev/provider-equinixmetal-controller/controllers/organization"
	controllersport "kubeform.dev/provider-equinixmetal-controller/controllers/port"
	controllersproject "kubeform.dev/provider-equinixmetal-controller/controllers/project"
	controllersreserved "kubeform.dev/provider-equinixmetal-controller/controllers/reserved"
	controllersspot "kubeform.dev/provider-equinixmetal-controller/controllers/spot"
	controllersssh "kubeform.dev/provider-equinixmetal-controller/controllers/ssh"
	controllersuser "kubeform.dev/provider-equinixmetal-controller/controllers/user"
	controllersvirtual "kubeform.dev/provider-equinixmetal-controller/controllers/virtual"
	controllersvlan "kubeform.dev/provider-equinixmetal-controller/controllers/vlan"
	controllersvolume "kubeform.dev/provider-equinixmetal-controller/controllers/volume"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var _provider = equinixmetal.Provider()

var runningControllers = struct {
	sync.RWMutex
	mp map[schema.GroupVersionKind]bool
}{mp: make(map[schema.GroupVersionKind]bool)}

func watchCRD(ctx context.Context, crdClient *clientset.Clientset, vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, stopCh <-chan struct{}, mgr manager.Manager, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				klog.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				klog.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				klog.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "equinixmetal.kubeform.com") {
				gvk := schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: crd.Spec.Versions[0].Name,
					Kind:    crd.Spec.Names.Kind,
				}

				// check whether this gvk came before, if no then start the controller
				runningControllers.RLock()
				_, ok := runningControllers.mp[gvk]
				runningControllers.RUnlock()

				if !ok {
					runningControllers.Lock()
					runningControllers.mp[gvk] = true
					runningControllers.Unlock()

					if enableValidatingWebhook {
						// add dynamic ValidatingWebhookConfiguration

						// create empty VWC if the group has come for the first time
						err := createEmptyVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						// update
						err = updateVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						err = SetupWebhook(mgr, gvk)
						if err != nil {
							setupLog.Error(err, "unable to enable webhook")
							os.Exit(1)
						}
					}

					err = SetupManager(ctx, mgr, gvk, auditor, restrictToNamespace)
					if err != nil {
						setupLog.Error(err, "unable to start manager")
						os.Exit(1)
					}
				}
			}
		},
	})

	informerFactory.Start(stopCh)

	return nil
}

func createEmptyVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	_, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err == nil || !(errors.IsNotFound(err)) {
		return err
	}

	emptyVWC := &arv1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ValidatingWebhookConfiguration",
			APIVersion: "admissionregistration.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-"),
			Labels: map[string]string{
				"app.kubernetes.io/instance": "equinixmetal.kubeform.com",
				"app.kubernetes.io/part-of":  "kubeform.com",
			},
		},
	}
	_, err = vwcClient.ValidatingWebhookConfigurations().Create(context.TODO(), emptyVWC, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func updateVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	vwc, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	path := "/validate-" + strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-") + "-v1alpha1-" + strings.ToLower(gvk.Kind)
	fail := arv1.Fail
	sideEffects := arv1.SideEffectClassNone
	admissionReviewVersions := []string{"v1beta1"}

	rules := []arv1.RuleWithOperations{
		{
			Operations: []arv1.OperationType{
				arv1.Delete,
				arv1.Update,
			},
			Rule: arv1.Rule{
				APIGroups:   []string{strings.ToLower(gvk.Group)},
				APIVersions: []string{gvk.Version},
				Resources:   []string{strings.ToLower(flect.Pluralize(gvk.Kind))},
			},
		},
	}

	data, err := ioutil.ReadFile("/tmp/k8s-webhook-server/serving-certs/ca.crt")
	if err != nil {
		return err
	}

	name := strings.ToLower(gvk.Kind) + "." + gvk.Group
	for _, webhook := range vwc.Webhooks {
		if webhook.Name == name {
			return nil
		}
	}

	newWebhook := arv1.ValidatingWebhook{
		Name: name,
		ClientConfig: arv1.WebhookClientConfig{
			Service: &arv1.ServiceReference{
				Namespace: webhookNamespace,
				Name:      webhookName,
				Path:      &path,
			},
			CABundle: data,
		},
		Rules:                   rules,
		FailurePolicy:           &fail,
		SideEffects:             &sideEffects,
		AdmissionReviewVersions: admissionReviewVersions,
	}

	vwc.Webhooks = append(vwc.Webhooks, newWebhook)

	_, err = vwcClient.ValidatingWebhookConfigurations().Update(context.TODO(), vwc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetupManager(ctx context.Context, mgr manager.Manager, gvk schema.GroupVersionKind, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "bgp.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Session",
	}:
		if err := (&controllersbgp.SessionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Session"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_bgp_session"],
			TypeName: "metal_bgp_session",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Session")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "connection.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Connection",
	}:
		if err := (&controllersconnection.ConnectionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Connection"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_connection"],
			TypeName: "metal_connection",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Connection")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "device.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Device",
	}:
		if err := (&controllersdevice.DeviceReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Device"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_device"],
			TypeName: "metal_device",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Device")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "device.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NetworkType",
	}:
		if err := (&controllersdevice.NetworkTypeReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("NetworkType"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_device_network_type"],
			TypeName: "metal_device_network_type",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "NetworkType")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "gateway.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Gateway",
	}:
		if err := (&controllersgateway.GatewayReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Gateway"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_gateway"],
			TypeName: "metal_gateway",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Gateway")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ip.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&controllersip.AttachmentReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Attachment"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_ip_attachment"],
			TypeName: "metal_ip_attachment",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Attachment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "organization.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Organization",
	}:
		if err := (&controllersorganization.OrganizationReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Organization"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_organization"],
			TypeName: "metal_organization",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Organization")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "port.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Port",
	}:
		if err := (&controllersport.PortReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Port"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_port"],
			TypeName: "metal_port",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Port")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "port.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "VlanAttachment",
	}:
		if err := (&controllersport.VlanAttachmentReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("VlanAttachment"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_port_vlan_attachment"],
			TypeName: "metal_port_vlan_attachment",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "VlanAttachment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Project",
	}:
		if err := (&controllersproject.ProjectReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Project"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_project"],
			TypeName: "metal_project",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Project")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ApiKey",
	}:
		if err := (&controllersproject.ApiKeyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("ApiKey"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_project_api_key"],
			TypeName: "metal_project_api_key",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ApiKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "SshKey",
	}:
		if err := (&controllersproject.SshKeyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("SshKey"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_project_ssh_key"],
			TypeName: "metal_project_ssh_key",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "SshKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "reserved.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "IpBlock",
	}:
		if err := (&controllersreserved.IpBlockReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("IpBlock"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_reserved_ip_block"],
			TypeName: "metal_reserved_ip_block",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "IpBlock")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "spot.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "MarketRequest",
	}:
		if err := (&controllersspot.MarketRequestReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("MarketRequest"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_spot_market_request"],
			TypeName: "metal_spot_market_request",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "MarketRequest")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ssh.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Key",
	}:
		if err := (&controllersssh.KeyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Key"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_ssh_key"],
			TypeName: "metal_ssh_key",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Key")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ApiKey",
	}:
		if err := (&controllersuser.ApiKeyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("ApiKey"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_user_api_key"],
			TypeName: "metal_user_api_key",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ApiKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "virtual.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Circuit",
	}:
		if err := (&controllersvirtual.CircuitReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Circuit"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_virtual_circuit"],
			TypeName: "metal_virtual_circuit",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Circuit")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "vlan.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Vlan",
	}:
		if err := (&controllersvlan.VlanReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Vlan"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_vlan"],
			TypeName: "metal_vlan",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Vlan")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&controllersvolume.VolumeReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Volume"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_volume"],
			TypeName: "metal_volume",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Volume")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&controllersvolume.AttachmentReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Attachment"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["metal_volume_attachment"],
			TypeName: "metal_volume_attachment",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Attachment")
			return err
		}

	default:
		return fmt.Errorf("Invalid CRD")
	}

	return nil
}

func SetupWebhook(mgr manager.Manager, gvk schema.GroupVersionKind) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "bgp.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Session",
	}:
		if err := (&bgpv1alpha1.Session{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Session")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "connection.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Connection",
	}:
		if err := (&connectionv1alpha1.Connection{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Connection")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "device.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Device",
	}:
		if err := (&devicev1alpha1.Device{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Device")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "device.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NetworkType",
	}:
		if err := (&devicev1alpha1.NetworkType{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "NetworkType")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "gateway.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Gateway",
	}:
		if err := (&gatewayv1alpha1.Gateway{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Gateway")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ip.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&ipv1alpha1.Attachment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Attachment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "organization.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Organization",
	}:
		if err := (&organizationv1alpha1.Organization{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Organization")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "port.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Port",
	}:
		if err := (&portv1alpha1.Port{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Port")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "port.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "VlanAttachment",
	}:
		if err := (&portv1alpha1.VlanAttachment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "VlanAttachment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Project",
	}:
		if err := (&projectv1alpha1.Project{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Project")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ApiKey",
	}:
		if err := (&projectv1alpha1.ApiKey{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "ApiKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "project.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "SshKey",
	}:
		if err := (&projectv1alpha1.SshKey{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "SshKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "reserved.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "IpBlock",
	}:
		if err := (&reservedv1alpha1.IpBlock{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "IpBlock")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "spot.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "MarketRequest",
	}:
		if err := (&spotv1alpha1.MarketRequest{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "MarketRequest")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ssh.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Key",
	}:
		if err := (&sshv1alpha1.Key{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Key")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ApiKey",
	}:
		if err := (&userv1alpha1.ApiKey{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "ApiKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "virtual.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Circuit",
	}:
		if err := (&virtualv1alpha1.Circuit{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Circuit")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "vlan.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Vlan",
	}:
		if err := (&vlanv1alpha1.Vlan{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Vlan")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&volumev1alpha1.Volume{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Volume")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.equinixmetal.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&volumev1alpha1.Attachment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Attachment")
			return err
		}

	default:
		return fmt.Errorf("Invalid Webhook")
	}

	return nil
}

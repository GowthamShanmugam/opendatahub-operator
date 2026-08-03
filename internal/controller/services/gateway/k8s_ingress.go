/*
Copyright 2025.

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

package gateway

// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete

import (
	"context"

	logf "sigs.k8s.io/controller-runtime/pkg/log"

	serviceApi "github.com/opendatahub-io/opendatahub-operator/v2/api/services/v1alpha1"
	odhtypes "github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/types"
)

// createK8sIngress adds Kubernetes Ingress template when in K8sRoute mode.
func createK8sIngress(ctx context.Context, rr *odhtypes.ReconciliationRequest) error {
	l := logf.FromContext(ctx).WithName("createK8sIngress")

	gatewayConfig, err := validateGatewayConfig(rr)
	if err != nil {
		return err
	}

	if gatewayConfig.Spec.IngressMode != serviceApi.IngressModeK8sRoute {
		l.V(1).Info("IngressMode is not K8sRoute, skipping Kubernetes Ingress creation")
		return nil
	}

	l.V(1).Info("Adding Kubernetes Ingress templates for Gateway")

	rr.Templates = append(rr.Templates,
		odhtypes.TemplateInfo{
			FS:   gatewayResources,
			Path: k8sIngressTemplate,
		},
	)

	return nil
}

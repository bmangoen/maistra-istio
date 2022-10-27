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

package ior

import (
	"testing"

	networking "istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/labels"

	v1 "github.com/openshift/api/route/v1"
	routev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
)

func TestGetRouteName(t *testing.T) {
	testCases := []struct {
		name              string
		gatewayNamespace  string
		gatewayName       string
		host              string
		expectedRouteName string
	}{
		{
			name:              "standard mesh name + standard namespace + standard service name",
			gatewayNamespace:  "bookinfo-servicemesh",
			gatewayName:       "bookinfo-gateway",
			host:              "productpage",
			expectedRouteName: "bookinfo-servicemesh-bookinfo-gateway-cc5918a76ef31e4b",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			routeName := getRouteName(tc.gatewayNamespace, tc.gatewayName, tc.host, "")
			if !labels.IsDNS1123Label(routeName) {
				t.Fatalf("Not a valid RFC 1123 label. Current length of %s: %d (should be <= %d)",
					routeName, len(routeName), labels.DNS1123LabelMaxLength)
			}
			if routeName != tc.expectedRouteName {
				t.Fatalf("%s not equals to the expected resource name %s",
					routeName, tc.expectedRouteName)
			}
		})
	}
}

func TestCreateRoute(t *testing.T) {
	testCases := []struct {
		name          string
		metadata      config.Meta
		gatewayName   string
		gateway       *networking.Gateway
		host          []string
		tls           *networking.ServerTLSSettings
		expectedRoute *v1.Route
	}{
		{
			name: "long name + long namespace",
			metadata: config.Meta{
				Name:      "bookinfo-gateway",
				Namespace: "bookinfo-servicemesh",
			},
			gateway: &networking.Gateway{},
			host:    []string{"productpage"},
			expectedRoute: &v1.Route{
				Spec: v1.RouteSpec{
					Host: "",
				},
			},
		},
	}

	var stop chan struct{}
	var errorChannel chan error
	var store model.ConfigStoreController
	var k8sClient KubeClient
	var routerClient routev1.RouteV1Interface
	var mrc *fakeMemberRollController
	controlPlaneNs := "istio-system"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stop = make(chan struct{})
			defer func() { close(stop) }()
			errorChannel = make(chan error)
			mrc = newFakeMemberRollController()
			store, k8sClient, routerClient = initClients(t, stop, errorChannel, mrc, true)
			createIngressGateway(t, k8sClient.GetActualClient(), controlPlaneNs, map[string]string{"istio": "ingressgateway"})
			// createGateway(t, store, tc.metadata.Namespace, tc.metadata.Name, tc.host, map[string]string{"istio": "ingressgateway"}, false, nil)
			r := newRoute(k8sClient, routerClient, store, controlPlaneNs, mrc, stop)
			route, err := r.createRoute(tc.metadata, tc.gateway, tc.host[0], tc.tls)
			if err != nil {
				t.Fatalf(err.Error())
			}
			t.Errorf("%s - %d", route.Name, len(route.Name))
			// if !labels.IsDNS1123Label(routeName) {
			// 	t.Fatalf("Not a valid RFC 1123 label. Current length of %s: %d (should be <= %d)",
			// 		routeName, len(routeName), labels.DNS1123LabelMaxLength)
			// }
			// if routeName != tc.expectedRouteName {
			// 	t.Fatalf("%s not equals to the expected resource name %s",
			// 		routeName, tc.expectedRouteName)
			// }
		})
	}

}

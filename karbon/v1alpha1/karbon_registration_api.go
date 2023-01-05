package v1alpha1

import (
	"context"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client/karbon"
)

type K8sRegistrationService interface {
	CreateK8sRegistration(ctx context.Context, createRequest *karbon.K8sCreateClusterRegistrationRequest) (*karbon.K8sCreateClusterRegistrationResponse, error)
	DeleteK8sRegistration(ctx context.Context) (*karbon.K8sClusterRegistrationDeleteResponse, error)
	GetK8sRegistration(ctx context.Context) (*karbon.K8sClusterRegistration, error)
	GetK8sRegistrationList(ctx context.Context) (*karbon.K8sClusterRegistrationList, error)
}

func (op karbon.Client) CreateK8sRegistration(ctx context.Context, createRequest *karbon.K8sCreateClusterRegistrationRequest) (*karbon.K8sCreateClusterRegistrationResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(karbon.K8sCreateClusterRegistrationResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

func (op karbon.Client) DeleteK8sRegistration(ctx context.Context) (*karbon.K8sClusterRegistrationDeleteResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(karbon.K8sClusterRegistrationDeleteResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

func (op karbon.Client) GetK8sRegistration(ctx context.Context) (*karbon.K8sClusterRegistration, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(karbon.K8sClusterRegistration)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

func (op karbon.Client) GetK8sRegistrationList(ctx context.Context) (*karbon.K8sClusterRegistrationList, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(karbon.K8sClusterRegistrationList)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

package v1alpha1

import (
	"context"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

var _ K8sRegistrationService = &V1Alpha1{}

// V1Alpha1 defines v1alpha1 APIs
type V1Alpha1 struct {
	httpClient *internal.Client
}

// K8sRegistrationService defines the create, delete, get and get the list of the k8s registration
type K8sRegistrationService interface {
	CreateK8sRegistration(ctx context.Context, createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error)
	DeleteK8sRegistration(ctx context.Context, UUID string) (*K8sClusterRegistrationDeleteResponse, error)
	GetK8sRegistration(ctx context.Context, UUID string) (*K8sClusterRegistration, error)
	GetK8sRegistrationList(ctx context.Context) (*K8sClusterRegistrationList, error)
}

// CreateK8sRegistration creates the k8s registration
func (op V1Alpha1) CreateK8sRegistration(ctx context.Context, createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sCreateClusterRegistrationResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// DeleteK8sRegistration deletes the k8s registration with UUID
func (op V1Alpha1) DeleteK8sRegistration(ctx context.Context, UUID string) (*K8sClusterRegistrationDeleteResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + UUID
	req, err := op.httpClient.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistrationDeleteResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// GetK8sRegistration gets the k8s registration with UUID
func (op V1Alpha1) GetK8sRegistration(ctx context.Context, UUID string) (*K8sClusterRegistration, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + UUID
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistration)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// GetK8sRegistrationList gets the k8s registration list
func (op V1Alpha1) GetK8sRegistrationList(ctx context.Context) (*K8sClusterRegistrationList, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistrationList)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

package karbon

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	absolutePath = "karbon"
	userAgent    = "nutanix"
	clientName   = "karbon"
)

// Client manages the V3 API
type Client struct {
	httpClient      *internal.Client
	Cluster         ClusterService
	PrivateRegistry PrivateRegistryService
	Meta            MetaService
}

type ClientService interface {
	CreateK8sRegistration(createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error)
	DeleteK8sRegistration() (*K8sClusterRegistrationDeleteResponse, error)
	GetK8sRegistration() (*K8sClusterRegistration, error)
	GetK8sRegistrationList() (*K8sClusterRegistrationList, error)
}

func (op Client) CreateK8sRegistration(createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error) {
	ctx := context.TODO()

	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, createRequest)
	karbonClusterActionResponse := new(K8sCreateClusterRegistrationResponse)

	if err != nil {
		return nil, err
	}

	return karbonClusterActionResponse, op.httpClient.Do(ctx, req, karbonClusterActionResponse)
}

func (op Client) DeleteK8sRegistration() (*K8sClusterRegistrationDeleteResponse, error) {
	ctx := context.TODO()

	path := "v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodDelete, path, nil)
	karbonClusterActionResponse := new(K8sClusterRegistrationDeleteResponse)

	if err != nil {
		return nil, err
	}

	return karbonClusterActionResponse, op.httpClient.Do(ctx, req, karbonClusterActionResponse)
}

func (op Client) GetK8sRegistration() (*K8sClusterRegistration, error) {
	ctx := context.TODO()

	path := "/v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	karbonClusterActionResponse := new(K8sClusterRegistration)

	if err != nil {
		return nil, err
	}

	return karbonClusterActionResponse, op.httpClient.Do(ctx, req, karbonClusterActionResponse)
}

func (op Client) GetK8sRegistrationList() (*K8sClusterRegistrationList, error) {
	ctx := context.TODO()

	path := "/v1-alpha.1/k8s/cluster-registrations/eae7fe7e-34e8-4978-bb9a-e49157e858d6"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	karbonClusterActionResponse := new(K8sClusterRegistrationList)

	if err != nil {
		return nil, err
	}

	return karbonClusterActionResponse, op.httpClient.Do(ctx, req, karbonClusterActionResponse)
}

// NewKarbonAPIClient return a internal to operate Karbon resources
func NewKarbonAPIClient(credentials prismgoclient.Credentials) (*Client, error) {
	if credentials.URL == "" || credentials.Username == "" || credentials.Password == "" {
		return nil, fmt.Errorf("karbon Client is missing: %s %s",
			"Please provide required details - %s in provider configuration",
			strings.Join(credentials.RequiredFields[clientName], ", "))
	}
	c, err := internal.NewClient(
		internal.WithCredentials(&credentials),
		internal.WithUserAgent(userAgent),
		internal.WithAbsolutePath(absolutePath))
	if err != nil {
		return nil, err
	}

	f := &Client{
		httpClient: c,
		Cluster: ClusterOperations{
			httpClient: c,
		},
		PrivateRegistry: PrivateRegistryOperations{
			httpClient: c,
		},
		Meta: MetaOperations{
			httpClient: c,
		},
	}

	return f, nil
}

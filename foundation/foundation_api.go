package foundation

import (
	"fmt"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	absolutePath = "foundation"
	userAgent    = "foundation"
	clientName   = "foundation"
)

// Foundation internal with its services
type Client struct {
	// base client
	client *internal.Client

	// Service for Imaging Nodes and Cluster Creation
	NodeImaging NodeImagingService

	// Service for File Management in foundation VM
	FileManagement FileManagementService

	// Service for Networking apis in foundation VM
	Networking NetworkingService
}

// This routine returns new Foundation API Client
func NewFoundationAPIClient(credentials prismgoclient.Credentials) (*Client, error) {
	var baseClient *internal.Client
	if credentials.FoundationEndpoint != "" {
		// for foundation internal, url should be based on foundation's endpoint and port
		credentials.URL = fmt.Sprintf("%s:%s", credentials.FoundationEndpoint, credentials.FoundationPort)
		c, err := internal.NewBaseClient(&credentials, absolutePath, true)
		if err != nil {
			return nil, err
		}
		baseClient = c
	} else {
		errorMsg := fmt.Sprintf("Foundation Client is missing. "+
			"Please provide required detail - %s in provider configuration.", strings.Join(credentials.RequiredFields[clientName], ", "))
		// create empty internal if required fields are not provided
		baseClient = &internal.Client{ErrorMsg: errorMsg}
	}

	// Fill user agent details
	baseClient.UserAgent = userAgent

	foundationClient := &Client{
		client: baseClient,
		NodeImaging: NodeImagingOperations{
			client: baseClient,
		},
		FileManagement: FileManagementOperations{
			client: baseClient,
		},
		Networking: NetworkingOperations{
			client: baseClient,
		},
	}
	return foundationClient, nil
}

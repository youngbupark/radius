/*
Copyright 2023 The Radius Authors.

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

package credential

import (
	"context"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/radius-project/radius/pkg/cli/clierrors"
	ucp "github.com/radius-project/radius/pkg/ucp/api/v20231001preview"
)

//go:generate mockgen -typed -destination=./mock_aws_credential_management.go -package=credential -self_package github.com/radius-project/radius/pkg/cli/credential github.com/radius-project/radius/pkg/cli/credential AWSCredentialManagementClientInterface

// AWSCredentialManagementClient is used to interface with cloud provider configuration and credentials.
type AWSCredentialManagementClient struct {
	AWSCredentialClient ucp.AwsCredentialsClient
}

const (
	AWSCredential        = "aws"
	AWSPlaneName         = "aws"
	awsCredentialKind    = "AccessKey"
	ValidInfoTemplate    = "enter valid info for %s"
	infoRequiredTemplate = "required info %s"
)

type AWSCredentialProperties struct {
	// Kind is the credential kind (AccessKey or IRSA)
	Kind *string

	// AccessKey is the properties for an AWS access key credential
	AccessKey *AWSAccessKeyCredentialProperties

	// IRSA is the properties for an AWS IRSA credential
	IRSA *AWSIRSACredentialProperties
}

type AWSAccessKeyCredentialProperties struct {
	// Kind is the credential kind (Must be AccessKey)
	Kind *string

	// AccessKeyID is the AWS access key ID
	AccessKeyID *string
}

type AWSIRSACredentialProperties struct {
	// Kind is the credential kind (Must be IRSA)
	Kind *string

	// RoleARN is the AWS IRSA role ARN
	RoleARN *string
}

// AWSCredentialManagementClient is used to interface with cloud provider configuration and credentials.
type AWSCredentialManagementClientInterface interface {
	// Get gets the credential registered with the given ucp provider plane.
	Get(ctx context.Context, name string) (ProviderCredentialConfiguration, error)
	// List lists the credentials registered with all ucp provider planes.
	List(ctx context.Context) ([]CloudProviderStatus, error)
	// Put registers an AWS credential with the respective ucp provider plane.
	Put(ctx context.Context, credential_config ucp.AwsCredentialResource) error
	// Delete unregisters credential from the given ucp provider plane.
	Delete(ctx context.Context, name string) (bool, error)
}

// Put registers credentials with the provided credential config
//

// "Put" checks if the credential type is "AWSCredential" and if so, creates or updates the credential in the AWS plane,
// otherwise it returns an error.
func (cpm *AWSCredentialManagementClient) Put(ctx context.Context, credential ucp.AwsCredentialResource) error {
	if strings.EqualFold(*credential.Type, AWSCredential) {
		_, err := cpm.AWSCredentialClient.CreateOrUpdate(ctx, AWSPlaneName, defaultSecretName, credential, nil)
		return err
	}
	return &ErrUnsupportedCloudProvider{}
}

// Get, gets the credential from the provided ucp provider plane
//

// "Get" retrieves the credentials for the specified cloud provider from the backend and returns a ProviderCredentialConfiguration
// object containing the credentials or an error if the credentials could not be retrieved.
func (cpm *AWSCredentialManagementClient) Get(ctx context.Context, credentialName string) (ProviderCredentialConfiguration, error) {
	// We send only the name when getting credentials from backend which we already have access to
	resp, err := cpm.AWSCredentialClient.Get(ctx, AWSPlaneName, credentialName, nil)
	if err != nil {
		return ProviderCredentialConfiguration{}, err
	}

	switch *resp.AwsCredentialResource.Properties.GetAwsCredentialProperties().Kind {
	case ucp.AWSCredentialKindAccessKey:
		awsAccessKeyCredentials, ok := resp.AwsCredentialResource.Properties.(*ucp.AwsAccessKeyCredentialProperties)
		if !ok {
			return ProviderCredentialConfiguration{}, clierrors.Message("Unable to find credentials for cloud provider %s.", AWSCredential)
		}

		providerCredentialConfiguration := ProviderCredentialConfiguration{
			CloudProviderStatus: CloudProviderStatus{
				Name:    AWSCredential,
				Enabled: true,
			},
			AWSCredentials: &AWSCredentialProperties{
				Kind: (*string)(awsAccessKeyCredentials.Kind),
				AccessKey: &AWSAccessKeyCredentialProperties{
					AccessKeyID: awsAccessKeyCredentials.AccessKeyID,
				},
			},
		}
		return providerCredentialConfiguration, nil
	case ucp.AWSCredentialKindIRSA:
		awsIRSACredentials, ok := resp.AwsCredentialResource.Properties.(*ucp.AwsIRSACredentialProperties)
		if !ok {
			return ProviderCredentialConfiguration{}, clierrors.Message("Unable to find credentials for cloud provider %s.", AWSCredential)
		}

		providerCredentialConfiguration := ProviderCredentialConfiguration{
			CloudProviderStatus: CloudProviderStatus{
				Name:    AWSCredential,
				Enabled: true,
			},
			AWSCredentials: &AWSCredentialProperties{
				Kind: (*string)(awsIRSACredentials.Kind),
				IRSA: &AWSIRSACredentialProperties{
					RoleARN: awsIRSACredentials.RoleARN,
				},
			},
		}
		return providerCredentialConfiguration, nil

	default:
		return ProviderCredentialConfiguration{}, clierrors.Message("Unable to find credentials for cloud provider %s.", AWSCredential)

	}

}

// List, lists the AWS credentials registered
//

// List retrieves a list of AWS credentials and returns a slice of CloudProviderStatus objects containing the name and
// enabled status of each credential. If an error occurs, an error is returned.
func (cpm *AWSCredentialManagementClient) List(ctx context.Context) ([]CloudProviderStatus, error) {
	// list AWS credential
	var providerList []*ucp.AwsCredentialResource

	pager := cpm.AWSCredentialClient.NewListPager(AWSPlaneName, nil)
	for pager.More() {
		nextPage, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		credList := nextPage.AwsCredentialResourceListResult.Value
		providerList = append(providerList, credList...)
	}

	res := []CloudProviderStatus{}
	if len(providerList) > 0 {
		res = append(res, CloudProviderStatus{
			Name:    AWSCredential,
			Enabled: true,
		})
	}
	return res, nil
}

// Delete, deletes the credentials from the given ucp provider plane
//

// Delete checks if a credential for the provider plane is registered and if so, deletes it; if not, it returns true
// without an error. If an error occurs, it returns false and the error.
func (cpm *AWSCredentialManagementClient) Delete(ctx context.Context, name string) (bool, error) {
	var respFromCtx *http.Response
	ctxWithResp := runtime.WithCaptureResponse(ctx, &respFromCtx)
	_, err := cpm.AWSCredentialClient.Delete(ctxWithResp, AWSPlaneName, name, nil)
	if err != nil {
		return false, err
	}
	return respFromCtx.StatusCode != 204, nil
}

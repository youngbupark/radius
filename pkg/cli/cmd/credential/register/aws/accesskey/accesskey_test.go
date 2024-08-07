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

package accesskey

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/cli/connections"
	cli_credential "github.com/radius-project/radius/pkg/cli/credential"
	"github.com/radius-project/radius/pkg/cli/framework"
	"github.com/radius-project/radius/pkg/cli/output"
	"github.com/radius-project/radius/pkg/cli/workspaces"
	"github.com/radius-project/radius/pkg/to"
	ucp "github.com/radius-project/radius/pkg/ucp/api/v20231001preview"
	"github.com/radius-project/radius/test/radcli"
	"github.com/stretchr/testify/require"
)

const (
	testAccessKeyId     = "TEST-ACCESS-KEY-ID"
	testSecretAccessKey = "TEST-SECRET-ACCESS-KEY"
)

func Test_CommandValidation(t *testing.T) {
	radcli.SharedCommandValidation(t, NewCommand)
}

func Test_Validate(t *testing.T) {
	configWithWorkspace := radcli.LoadConfigWithWorkspace(t)
	testcases := []radcli.ValidateInput{
		{
			Name: "Valid AWS command",
			Input: []string{
				"--access-key-id", testAccessKeyId,
				"--secret-access-key", testSecretAccessKey,
			},
			ExpectedValid: true,
			ConfigHolder:  framework.ConfigHolder{Config: configWithWorkspace},
		},
		{
			Name: "AWS command with fallback workspace",
			Input: []string{
				"--access-key-id", testAccessKeyId,
				"--secret-access-key", testSecretAccessKey,
			},
			ExpectedValid: true,
			ConfigHolder:  framework.ConfigHolder{Config: radcli.LoadEmptyConfig(t)},
		},
		{
			Name: "AWS command with too many positional args",
			Input: []string{
				"letsgoooooo",
				"--access-key-id", testAccessKeyId,
				"--secret-access-key", testSecretAccessKey,
			},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: configWithWorkspace},
		},
		{
			Name: "AWS command without IAM access key id",
			Input: []string{
				"--access-key-id", "",
				"--secret-access-key", testSecretAccessKey,
			},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: configWithWorkspace},
		},
		{
			Name: "AWS command without IAM secret access key",
			Input: []string{
				"--access-key-id", testAccessKeyId,
				"--secret-access-key", "",
			},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: configWithWorkspace},
		},
	}
	radcli.SharedValidateValidation(t, NewCommand, testcases)
}

func Test_Run(t *testing.T) {
	t.Run("Create aws provider", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			expectedPut := ucp.AwsCredentialResource{
				Location: to.Ptr(v1.LocationGlobal),
				Type:     to.Ptr(cli_credential.AWSCredential),
				Properties: &ucp.AwsAccessKeyCredentialProperties{
					Storage: &ucp.CredentialStorageProperties{
						Kind: to.Ptr(ucp.CredentialStorageKindInternal),
					},
					AccessKeyID:     to.Ptr(testAccessKeyId),
					SecretAccessKey: to.Ptr(testSecretAccessKey),
				},
			}

			client := cli_credential.NewMockCredentialManagementClient(ctrl)
			client.EXPECT().
				PutAWS(gomock.Any(), expectedPut).
				Return(nil).
				Times(1)

			outputSink := &output.MockOutput{}

			runner := &Runner{
				ConnectionFactory: &connections.MockFactory{CredentialManagementClient: client},
				Output:            outputSink,
				Workspace: &workspaces.Workspace{
					Connection: map[string]any{
						"kind":    workspaces.KindKubernetes,
						"context": "my-context",
					},
					Source: workspaces.SourceUserConfig,
				},
				Format:          "table",
				AccessKeyID:     testAccessKeyId,
				SecretAccessKey: testSecretAccessKey,
				KubeContext:     "my-context",
			}
			err := runner.Run(context.Background())
			require.NoError(t, err)

			expected := []any{
				output.LogOutput{
					Format: "Registering credential for %q cloud provider in Radius installation %q...",
					Params: []any{"aws", "Kubernetes (context=my-context)"},
				},
				output.LogOutput{
					Format: "Successfully registered credential for %q cloud provider. Tokens may take up to 30 seconds to refresh.",
					Params: []any{"aws"},
				},
			}
			require.Equal(t, expected, outputSink.Writes)
		})
	})
}

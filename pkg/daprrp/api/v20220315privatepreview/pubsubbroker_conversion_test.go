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

package v20220315privatepreview

import (
	"encoding/json"
	"os"
	"testing"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/daprrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
	"github.com/stretchr/testify/require"
)

type fakeResource struct{}

func TestDaprPubSubBroker_ConvertVersionedToDataModel(t *testing.T) {
	testset := []string{
		"pubsubbrokerazureresource.json",
		"pubsubbrokerresource_recipe.json",
		"pubsubbrokerresource_recipe2.json",
		"pubsubbrokergenericresource.json"}

	for _, payload := range testset {
		// arrange
		rawPayload := loadTestData(payload)
		versionedResource := &DaprPubSubBrokerResource{}
		err := json.Unmarshal(rawPayload, versionedResource)
		require.NoError(t, err)

		// act
		dm, err := versionedResource.ConvertTo()

		// assert
		require.NoError(t, err)
		convertedResource := dm.(*datamodel.DaprPubSubBroker)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Dapr/pubSubBrokers/pubSub0", convertedResource.ID)
		require.Equal(t, "pubSub0", convertedResource.Name)
		require.Equal(t, linkrp.N_DaprPubSubBrokersResourceType, convertedResource.Type)
		require.Equal(t, "2022-03-15-privatepreview", convertedResource.InternalMetadata.UpdatedAPIVersion)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication", convertedResource.Properties.Application)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0", convertedResource.Properties.Environment)
		switch versionedResource.Properties.(type) {
		case *ResourceDaprPubSubProperties:
			require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testGroup/providers/Microsoft.ServiceBus/namespaces/testQueue", convertedResource.Properties.Resource)
		case *ValuesDaprPubSubProperties:
			require.Equal(t, "pubsub.kafka", convertedResource.Properties.Type)
			require.Equal(t, "v1", convertedResource.Properties.Version)
			require.Equal(t, "bar", convertedResource.Properties.Metadata["foo"])
			require.Equal(t, []rpv1.OutputResource(nil), convertedResource.Properties.Status.OutputResources)
		case *RecipeDaprPubSubProperties:
			if payload == "pubsubbrokerresource_recipe2.json" {
				parameters := map[string]any{"port": float64(6081)}
				require.Equal(t, parameters, convertedResource.Properties.Recipe.Parameters)
			} else {
				require.Equal(t, "redis-test", convertedResource.Properties.Recipe.Name)
			}
		}
	}

}

func TestDaprPubSubBroker_ConvertDataModelToVersioned(t *testing.T) {
	testset := []string{
		"pubsubbrokerazureresourcedatamodel.json",
		"pubsubbrokergenericresourcedatamodel.json",
		"pubsubbrokerresourcedatamodel_recipe.json",
		"pubsubbrokerresourcedatamodel_recipe2.json"}

	for _, payload := range testset {
		// arrange
		rawPayload := loadTestData(payload)
		resource := &datamodel.DaprPubSubBroker{}
		err := json.Unmarshal(rawPayload, resource)
		require.NoError(t, err)

		// act
		versionedResource := &DaprPubSubBrokerResource{}
		err = versionedResource.ConvertFrom(resource)

		// assert
		require.NoError(t, err)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Dapr/pubSubBrokers/pubSub0", resource.ID)
		require.Equal(t, "pubSub0", resource.Name)
		require.Equal(t, linkrp.N_DaprPubSubBrokersResourceType, resource.Type)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/applications/testApplication", resource.Properties.Application)
		require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/environments/env0", resource.Properties.Environment)
		switch v := versionedResource.Properties.(type) {
		case *ResourceDaprPubSubProperties:
			require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testGroup/providers/Microsoft.ServiceBus/namespaces/testQueue", *v.Resource)
			require.Equal(t, "Deployment", v.GetDaprPubSubBrokerProperties().Status.OutputResources[0]["LocalID"])
			require.Equal(t, "kubernetes", v.GetDaprPubSubBrokerProperties().Status.OutputResources[0]["Provider"])
		case *ValuesDaprPubSubProperties:
			require.Equal(t, "pubsub.kafka", *v.Type)
			require.Equal(t, "v1", *v.Version)
			require.Equal(t, "bar", v.Metadata["foo"])
		case *RecipeDaprPubSubProperties:
			if payload == "daprpubsubbrokerresourcedatamodel_recipe2" {
				parameters := map[string]any{"port": float64(6081)}
				require.Equal(t, parameters, resource.Properties.Recipe.Parameters)
			} else {
				require.Equal(t, "redis-test", resource.Properties.Recipe.Name)
			}
		}
	}
}

func TestDaprPubSubBroker_ConvertFromValidation(t *testing.T) {
	validationTests := []struct {
		src v1.DataModelInterface
		err error
	}{
		{&fakeResource{}, v1.ErrInvalidModelConversion},
		{nil, v1.ErrInvalidModelConversion},
	}

	for _, tc := range validationTests {
		versioned := &DaprPubSubBrokerResource{}
		err := versioned.ConvertFrom(tc.src)
		require.ErrorAs(t, tc.err, &err)
	}
}

func (f *fakeResource) ResourceTypeName() string {
	return "FakeResource"
}

func loadTestData(testfile string) []byte {
	d, err := os.ReadFile("./testdata/" + testfile)
	if err != nil {
		return nil
	}
	return d
}

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
	"fmt"
	"os"
	"time"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp"
	"github.com/project-radius/radius/pkg/to"
)

func toProvisioningStateDataModel(state *ProvisioningState) v1.ProvisioningState {
	if state == nil {
		return v1.ProvisioningStateAccepted
	}

	switch *state {
	case ProvisioningStateUpdating:
		return v1.ProvisioningStateUpdating
	case ProvisioningStateDeleting:
		return v1.ProvisioningStateDeleting
	case ProvisioningStateAccepted:
		return v1.ProvisioningStateAccepted
	case ProvisioningStateSucceeded:
		return v1.ProvisioningStateSucceeded
	case ProvisioningStateFailed:
		return v1.ProvisioningStateFailed
	case ProvisioningStateCanceled:
		return v1.ProvisioningStateCanceled
	case ProvisioningStateProvisioning:
		return v1.ProvisioningStateProvisioning
	default:
		return v1.ProvisioningStateAccepted
	}
}

func fromProvisioningStateDataModel(state v1.ProvisioningState) *ProvisioningState {
	var converted ProvisioningState
	switch state {
	case v1.ProvisioningStateUpdating:
		converted = ProvisioningStateUpdating
	case v1.ProvisioningStateDeleting:
		converted = ProvisioningStateDeleting
	case v1.ProvisioningStateAccepted:
		converted = ProvisioningStateAccepted
	case v1.ProvisioningStateSucceeded:
		converted = ProvisioningStateSucceeded
	case v1.ProvisioningStateFailed:
		converted = ProvisioningStateFailed
	case v1.ProvisioningStateCanceled:
		converted = ProvisioningStateCanceled
	default:
		converted = ProvisioningStateAccepted
	}

	return &converted
}

func toResourceProvisiongDataModel(provisioning *ResourceProvisioning) (linkrp.ResourceProvisioning, error) {
	if provisioning == nil {
		return linkrp.ResourceProvisioningRecipe, nil
	}
	switch *provisioning {
	case ResourceProvisioningManual:
		return linkrp.ResourceProvisioningManual, nil
	case ResourceProvisioningRecipe:
		return linkrp.ResourceProvisioningRecipe, nil
	default:
		return "", &v1.ErrModelConversion{PropertyName: "$.properties.resourceProvisioning", ValidValue: fmt.Sprintf("one of %s", PossibleResourceProvisioningValues())}
	}
}

func fromResourceProvisioningDataModel(provisioning linkrp.ResourceProvisioning) *ResourceProvisioning {
	var converted ResourceProvisioning
	switch provisioning {
	case linkrp.ResourceProvisioningManual:
		converted = ResourceProvisioningManual
	default:
		converted = ResourceProvisioningRecipe
	}

	return &converted
}

func fromSystemDataModel(s v1.SystemData) *SystemData {
	return &SystemData{
		CreatedBy:          to.Ptr(s.CreatedBy),
		CreatedByType:      (*CreatedByType)(to.Ptr(s.CreatedByType)),
		CreatedAt:          unmarshalTimeString(s.CreatedAt),
		LastModifiedBy:     to.Ptr(s.LastModifiedBy),
		LastModifiedByType: (*CreatedByType)(to.Ptr(s.LastModifiedByType)),
		LastModifiedAt:     unmarshalTimeString(s.LastModifiedAt),
	}
}

func toRecipeDataModel(r *Recipe) linkrp.LinkRecipe {
	if r == nil {
		return linkrp.LinkRecipe{}
	}

	recipe := linkrp.LinkRecipe{
		Name: to.String(r.Name),
	}

	if r.Parameters != nil {
		recipe.Parameters = r.Parameters
	}
	return recipe
}

func fromRecipeDataModel(r linkrp.LinkRecipe) *Recipe {
	return &Recipe{
		Name:       to.Ptr(r.Name),
		Parameters: r.Parameters,
	}
}

func loadTestData(testfile string) ([]byte, error) {
	d, err := os.ReadFile(testfile)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func unmarshalTimeString(ts string) *time.Time {
	var tt timeRFC3339
	_ = tt.UnmarshalText([]byte(ts))
	return (*time.Time)(&tt)
}
//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BasicDaprResourceProperties.
func (b BasicDaprResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "application", b.Application)
	populate(objectMap, "componentName", b.ComponentName)
	populate(objectMap, "environment", b.Environment)
	populate(objectMap, "status", b.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BasicDaprResourceProperties.
func (b *BasicDaprResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "application":
				err = unpopulate(val, "Application", &b.Application)
				delete(rawMsg, key)
		case "componentName":
				err = unpopulate(val, "ComponentName", &b.ComponentName)
				delete(rawMsg, key)
		case "environment":
				err = unpopulate(val, "Environment", &b.Environment)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &b.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BasicResourceProperties.
func (b BasicResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "application", b.Application)
	populate(objectMap, "environment", b.Environment)
	populate(objectMap, "status", b.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BasicResourceProperties.
func (b *BasicResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "application":
				err = unpopulate(val, "Application", &b.Application)
				delete(rawMsg, key)
		case "environment":
				err = unpopulate(val, "Environment", &b.Environment)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &b.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
				err = unpopulate(val, "Info", &e.Info)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &e.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
				err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
				delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "details":
				err = unpopulate(val, "Details", &e.Details)
				delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
				delete(rawMsg, key)
		case "target":
				err = unpopulate(val, "Target", &e.Target)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &e.Error)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RabbitMQListSecretsResult.
func (r RabbitMQListSecretsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectionString", r.ConnectionString)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RabbitMQListSecretsResult.
func (r *RabbitMQListSecretsResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionString":
				err = unpopulate(val, "ConnectionString", &r.ConnectionString)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RabbitMQQueueProperties.
func (r RabbitMQQueueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "application", r.Application)
	populate(objectMap, "environment", r.Environment)
	objectMap["mode"] = r.Mode
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "secrets", r.Secrets)
	populate(objectMap, "status", r.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RabbitMQQueueProperties.
func (r *RabbitMQQueueProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "application":
				err = unpopulate(val, "Application", &r.Application)
				delete(rawMsg, key)
		case "environment":
				err = unpopulate(val, "Environment", &r.Environment)
				delete(rawMsg, key)
		case "mode":
				err = unpopulate(val, "Mode", &r.Mode)
				delete(rawMsg, key)
		case "provisioningState":
				err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
				delete(rawMsg, key)
		case "secrets":
				err = unpopulate(val, "Secrets", &r.Secrets)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &r.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RabbitMQQueueResource.
func (r RabbitMQQueueResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RabbitMQQueueResource.
func (r *RabbitMQQueueResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &r.ID)
				delete(rawMsg, key)
		case "location":
				err = unpopulate(val, "Location", &r.Location)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &r.Name)
				delete(rawMsg, key)
		case "properties":
				r.Properties, err = unmarshalRabbitMQQueuePropertiesClassification(val)
				delete(rawMsg, key)
		case "systemData":
				err = unpopulate(val, "SystemData", &r.SystemData)
				delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &r.Tags)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &r.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RabbitMQQueueResourceListResult.
func (r RabbitMQQueueResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RabbitMQQueueResourceListResult.
func (r *RabbitMQQueueResourceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
				err = unpopulate(val, "NextLink", &r.NextLink)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &r.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RabbitMQSecrets.
func (r RabbitMQSecrets) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectionString", r.ConnectionString)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RabbitMQSecrets.
func (r *RabbitMQSecrets) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionString":
				err = unpopulate(val, "ConnectionString", &r.ConnectionString)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Recipe.
func (r Recipe) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", r.Name)
	populate(objectMap, "parameters", r.Parameters)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Recipe.
func (r *Recipe) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
				err = unpopulate(val, "Name", &r.Name)
				delete(rawMsg, key)
		case "parameters":
				err = unpopulate(val, "Parameters", &r.Parameters)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecipeRabbitMQQueueProperties.
func (r RecipeRabbitMQQueueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "application", r.Application)
	populate(objectMap, "environment", r.Environment)
	objectMap["mode"] = "recipe"
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "queue", r.Queue)
	populate(objectMap, "recipe", r.Recipe)
	populate(objectMap, "secrets", r.Secrets)
	populate(objectMap, "status", r.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecipeRabbitMQQueueProperties.
func (r *RecipeRabbitMQQueueProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "application":
				err = unpopulate(val, "Application", &r.Application)
				delete(rawMsg, key)
		case "environment":
				err = unpopulate(val, "Environment", &r.Environment)
				delete(rawMsg, key)
		case "mode":
				err = unpopulate(val, "Mode", &r.Mode)
				delete(rawMsg, key)
		case "provisioningState":
				err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
				delete(rawMsg, key)
		case "queue":
				err = unpopulate(val, "Queue", &r.Queue)
				delete(rawMsg, key)
		case "recipe":
				err = unpopulate(val, "Recipe", &r.Recipe)
				delete(rawMsg, key)
		case "secrets":
				err = unpopulate(val, "Secrets", &r.Secrets)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &r.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &r.ID)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &r.Name)
				delete(rawMsg, key)
		case "systemData":
				err = unpopulate(val, "SystemData", &r.SystemData)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &r.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceStatus.
func (r ResourceStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "outputResources", r.OutputResources)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceStatus.
func (r *ResourceStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "outputResources":
				err = unpopulate(val, "OutputResources", &r.OutputResources)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
				err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
				delete(rawMsg, key)
		case "createdBy":
				err = unpopulate(val, "CreatedBy", &s.CreatedBy)
				delete(rawMsg, key)
		case "createdByType":
				err = unpopulate(val, "CreatedByType", &s.CreatedByType)
				delete(rawMsg, key)
		case "lastModifiedAt":
				err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
				delete(rawMsg, key)
		case "lastModifiedBy":
				err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
				delete(rawMsg, key)
		case "lastModifiedByType":
				err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrackedResource.
func (t *TrackedResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &t.ID)
				delete(rawMsg, key)
		case "location":
				err = unpopulate(val, "Location", &t.Location)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &t.Name)
				delete(rawMsg, key)
		case "systemData":
				err = unpopulate(val, "SystemData", &t.SystemData)
				delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &t.Tags)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &t.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ValuesRabbitMQQueueProperties.
func (v ValuesRabbitMQQueueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "application", v.Application)
	populate(objectMap, "environment", v.Environment)
	objectMap["mode"] = "values"
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "queue", v.Queue)
	populate(objectMap, "secrets", v.Secrets)
	populate(objectMap, "status", v.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ValuesRabbitMQQueueProperties.
func (v *ValuesRabbitMQQueueProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "application":
				err = unpopulate(val, "Application", &v.Application)
				delete(rawMsg, key)
		case "environment":
				err = unpopulate(val, "Environment", &v.Environment)
				delete(rawMsg, key)
		case "mode":
				err = unpopulate(val, "Mode", &v.Mode)
				delete(rawMsg, key)
		case "provisioningState":
				err = unpopulate(val, "ProvisioningState", &v.ProvisioningState)
				delete(rawMsg, key)
		case "queue":
				err = unpopulate(val, "Queue", &v.Queue)
				delete(rawMsg, key)
		case "secrets":
				err = unpopulate(val, "Secrets", &v.Secrets)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &v.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}


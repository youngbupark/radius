//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import "encoding/json"

func unmarshalDaprPubSubBrokerPropertiesClassification(rawMsg json.RawMessage) (DaprPubSubBrokerPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DaprPubSubBrokerPropertiesClassification
	switch m["mode"] {
	case "recipe":
		b = &RecipeDaprPubSubProperties{}
	case "resource":
		b = &ResourceDaprPubSubProperties{}
	case "values":
		b = &ValuesDaprPubSubProperties{}
	default:
		b = &DaprPubSubBrokerProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDaprSecretStorePropertiesClassification(rawMsg json.RawMessage) (DaprSecretStorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DaprSecretStorePropertiesClassification
	switch m["mode"] {
	case "recipe":
		b = &RecipeDaprSecretStoreProperties{}
	case "values":
		b = &ValuesDaprSecretStoreProperties{}
	default:
		b = &DaprSecretStoreProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDaprStateStorePropertiesClassification(rawMsg json.RawMessage) (DaprStateStorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DaprStateStorePropertiesClassification
	switch m["mode"] {
	case "recipe":
		b = &RecipeDaprStateStoreProperties{}
	case "resource":
		b = &ResourceDaprStateStoreProperties{}
	case "values":
		b = &ValuesDaprStateStoreProperties{}
	default:
		b = &DaprStateStoreProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}


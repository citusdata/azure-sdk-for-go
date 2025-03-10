//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import "encoding/json"

func unmarshalDashboardPartMetadataClassification(rawMsg json.RawMessage) (DashboardPartMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DashboardPartMetadataClassification
	switch m["type"] {
	case "Extension/HubsExtension/PartType/MarkdownPart":
		b = &MarkdownPartMetadata{}
	default:
		b = &DashboardPartMetadata{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

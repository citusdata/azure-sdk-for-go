//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/BillingAccountDimensionsList.json
func ExampleDimensionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewDimensionsClient(cred, nil)
	res, err := client.List(ctx,
		"<scope>",
		&armcostmanagement.DimensionsClientListOptions{Filter: nil,
			Expand:    nil,
			Skiptoken: nil,
			Top:       nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DimensionsClientListResult)
}

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountsDimensions.json
func ExampleDimensionsClient_ByExternalCloudProviderType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewDimensionsClient(cred, nil)
	res, err := client.ByExternalCloudProviderType(ctx,
		armcostmanagement.ExternalCloudProviderType("externalBillingAccounts"),
		"<external-cloud-provider-id>",
		&armcostmanagement.DimensionsClientByExternalCloudProviderTypeOptions{Filter: nil,
			Expand:    nil,
			Skiptoken: nil,
			Top:       nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DimensionsClientByExternalCloudProviderTypeResult)
}

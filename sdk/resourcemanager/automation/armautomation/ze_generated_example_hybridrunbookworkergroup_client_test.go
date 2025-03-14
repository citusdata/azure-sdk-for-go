//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewHybridRunbookWorkerGroupClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<hybrid-runbook-worker-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewHybridRunbookWorkerGroupClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<hybrid-runbook-worker-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HybridRunbookWorkerGroupClientGetResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/putHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewHybridRunbookWorkerGroupClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<hybrid-runbook-worker-group-name>",
		armautomation.HybridRunbookWorkerGroupCreateOrUpdateParameters{
			Credential: &armautomation.RunAsCredentialAssociationProperty{
				Name: to.StringPtr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HybridRunbookWorkerGroupClientCreateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/updateHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewHybridRunbookWorkerGroupClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<hybrid-runbook-worker-group-name>",
		armautomation.HybridRunbookWorkerGroupCreateOrUpdateParameters{
			Credential: &armautomation.RunAsCredentialAssociationProperty{
				Name: to.StringPtr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HybridRunbookWorkerGroupClientUpdateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_ListByAutomationAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewHybridRunbookWorkerGroupClient("<subscription-id>", cred, nil)
	pager := client.ListByAutomationAccount("<resource-group-name>",
		"<automation-account-name>",
		&armautomation.HybridRunbookWorkerGroupClientListByAutomationAccountOptions{Filter: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceConnection/list.json
func ExampleWorkspaceConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspaceConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		&armmachinelearningservices.WorkspaceConnectionsClientListOptions{Target: to.StringPtr("<target>"),
			Category: to.StringPtr("<category>"),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspaceConnectionsClientListResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceConnection/create.json
func ExampleWorkspaceConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspaceConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<connection-name>",
		armmachinelearningservices.WorkspaceConnection{
			Properties: &armmachinelearningservices.WorkspaceConnectionProps{
				AuthType: to.StringPtr("<auth-type>"),
				Category: to.StringPtr("<category>"),
				Target:   to.StringPtr("<target>"),
				Value:    to.StringPtr("<value>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspaceConnectionsClientCreateResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceConnection/get.json
func ExampleWorkspaceConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspaceConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspaceConnectionsClientGetResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceConnection/delete.json
func ExampleWorkspaceConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspaceConnectionsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

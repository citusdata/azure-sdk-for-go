//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_example.json
func ExampleAssessmentsMetadataClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_example.json
func ExampleAssessmentsMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<assessment-metadata-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssessmentsMetadataClientGetResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
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

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_GetInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	res, err := client.GetInSubscription(ctx,
		"<assessment-metadata-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssessmentsMetadataClientGetInSubscriptionResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	res, err := client.CreateInSubscription(ctx,
		"<assessment-metadata-name>",
		armsecurity.AssessmentMetadataResponse{
			Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
				Description:    to.StringPtr("<description>"),
				AssessmentType: armsecurity.AssessmentType("CustomerManaged").ToPtr(),
				Categories: []*armsecurity.Categories{
					armsecurity.Categories("Compute").ToPtr()},
				DisplayName:            to.StringPtr("<display-name>"),
				ImplementationEffort:   armsecurity.ImplementationEffort("Low").ToPtr(),
				RemediationDescription: to.StringPtr("<remediation-description>"),
				Severity:               armsecurity.Severity("Medium").ToPtr(),
				Threats: []*armsecurity.Threats{
					armsecurity.Threats("dataExfiltration").ToPtr(),
					armsecurity.Threats("dataSpillage").ToPtr(),
					armsecurity.Threats("maliciousInsider").ToPtr()},
				UserImpact: armsecurity.UserImpact("Low").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssessmentsMetadataClientCreateInSubscriptionResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_DeleteInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAssessmentsMetadataClient("<subscription-id>", cred, nil)
	_, err = client.DeleteInSubscription(ctx,
		"<assessment-metadata-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

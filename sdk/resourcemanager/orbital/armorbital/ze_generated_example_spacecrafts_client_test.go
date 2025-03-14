//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
)

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftsBySubscriptionList.json
func ExampleSpacecraftsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	res, err := client.ListBySubscription(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientListBySubscriptionResult)
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftsByResourceGroupList.json
func ExampleSpacecraftsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientListResult)
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftGet.json
func ExampleSpacecraftsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientGetResult)
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftCreate.json
func ExampleSpacecraftsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.Spacecraft{
			Location: to.StringPtr("<location>"),
			Properties: &armorbital.SpacecraftsProperties{
				Links: []*armorbital.SpacecraftLink{
					{
						BandwidthMHz:       to.Float32Ptr(0.036),
						CenterFrequencyMHz: to.Float32Ptr(2106.4063),
						Direction:          armorbital.Direction("uplink").ToPtr(),
						Polarization:       armorbital.Polarization("RHCP").ToPtr(),
					},
					{
						BandwidthMHz:       to.Float32Ptr(150),
						CenterFrequencyMHz: to.Float32Ptr(8125),
						Direction:          armorbital.Direction("downlink").ToPtr(),
						Polarization:       armorbital.Polarization("RHCP").ToPtr(),
					}},
				NoradID:   to.StringPtr("<norad-id>"),
				TitleLine: to.StringPtr("<title-line>"),
				TleLine1:  to.StringPtr("<tle-line1>"),
				TleLine2:  to.StringPtr("<tle-line2>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftDelete.json
func ExampleSpacecraftsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/SpacecraftUpdateTags.json
func ExampleSpacecraftsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.TagsObject{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientUpdateTagsResult)
}

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/AvailableContactsList.json
func ExampleSpacecraftsClient_BeginListAvailableContacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginListAvailableContacts(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.ContactParameters{
			ContactProfile: &armorbital.ResourceReference{
				ID: to.StringPtr("<id>"),
			},
			EndTime:           to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-17T23:49:40.00Z"); return t }()),
			GroundStationName: to.StringPtr("<ground-station-name>"),
			StartTime:         to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T05:40:21.00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SpacecraftsClientListAvailableContactsResult)
}

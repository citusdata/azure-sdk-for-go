//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/LogAnalyticsRequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExportRequestRateByInterval(ctx,
		"<location>",
		armcompute.RequestRateByIntervalInput{
			BlobContainerSasURI: to.StringPtr("<blob-container-sas-uri>"),
			FromTime:            to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
			GroupByResourceName: to.BoolPtr(true),
			ToTime:              to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
			IntervalLength:      armcompute.IntervalInMinsFiveMins.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LogAnalyticsClientExportRequestRateByIntervalResult)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/LogAnalyticsThrottledRequests.json
func ExampleLogAnalyticsClient_BeginExportThrottledRequests() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExportThrottledRequests(ctx,
		"<location>",
		armcompute.ThrottledRequestsInput{
			BlobContainerSasURI:        to.StringPtr("<blob-container-sas-uri>"),
			FromTime:                   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
			GroupByClientApplicationID: to.BoolPtr(false),
			GroupByOperationName:       to.BoolPtr(true),
			GroupByResourceName:        to.BoolPtr(false),
			GroupByUserAgent:           to.BoolPtr(false),
			ToTime:                     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LogAnalyticsClientExportThrottledRequestsResult)
}

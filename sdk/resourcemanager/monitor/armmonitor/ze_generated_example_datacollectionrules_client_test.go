//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesListByResourceGroup.json
func ExampleDataCollectionRulesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesListBySubscription.json
func ExampleDataCollectionRulesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesGet.json
func ExampleDataCollectionRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionRulesClientGetResult)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesCreate.json
func ExampleDataCollectionRulesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		&armmonitor.DataCollectionRulesClientCreateOptions{Body: &armmonitor.DataCollectionRuleResource{
			Location: to.StringPtr("<location>"),
			Properties: &armmonitor.DataCollectionRuleResourceProperties{
				DataFlows: []*armmonitor.DataFlow{
					{
						Destinations: []*string{
							to.StringPtr("centralWorkspace")},
						Streams: []*armmonitor.KnownDataFlowStreams{
							armmonitor.KnownDataFlowStreams("Microsoft-Perf").ToPtr(),
							armmonitor.KnownDataFlowStreams("Microsoft-Syslog").ToPtr(),
							armmonitor.KnownDataFlowStreams("Microsoft-WindowsEvent").ToPtr()},
					}},
				DataSources: &armmonitor.DataCollectionRuleDataSources{
					PerformanceCounters: []*armmonitor.PerfCounterDataSource{
						{
							Name: to.StringPtr("<name>"),
							CounterSpecifiers: []*string{
								to.StringPtr("\\Processor(_Total)\\% Processor Time"),
								to.StringPtr("\\Memory\\Committed Bytes"),
								to.StringPtr("\\LogicalDisk(_Total)\\Free Megabytes"),
								to.StringPtr("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length")},
							SamplingFrequencyInSeconds: to.Int32Ptr(15),
							Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
								armmonitor.KnownPerfCounterDataSourceStreams("Microsoft-Perf").ToPtr()},
						},
						{
							Name: to.StringPtr("<name>"),
							CounterSpecifiers: []*string{
								to.StringPtr("\\Process(_Total)\\Thread Count")},
							SamplingFrequencyInSeconds: to.Int32Ptr(30),
							Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
								armmonitor.KnownPerfCounterDataSourceStreams("Microsoft-Perf").ToPtr()},
						}},
					Syslog: []*armmonitor.SyslogDataSource{
						{
							Name: to.StringPtr("<name>"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								armmonitor.KnownSyslogDataSourceFacilityNames("cron").ToPtr()},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								armmonitor.KnownSyslogDataSourceLogLevels("Debug").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Critical").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Emergency").ToPtr()},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								armmonitor.KnownSyslogDataSourceStreams("Microsoft-Syslog").ToPtr()},
						},
						{
							Name: to.StringPtr("<name>"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								armmonitor.KnownSyslogDataSourceFacilityNames("syslog").ToPtr()},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								armmonitor.KnownSyslogDataSourceLogLevels("Alert").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Critical").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Emergency").ToPtr()},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								armmonitor.KnownSyslogDataSourceStreams("Microsoft-Syslog").ToPtr()},
						}},
					WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
						{
							Name: to.StringPtr("<name>"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								armmonitor.KnownWindowsEventLogDataSourceStreams("Microsoft-WindowsEvent").ToPtr()},
							XPathQueries: []*string{
								to.StringPtr("Security!")},
						},
						{
							Name: to.StringPtr("<name>"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								armmonitor.KnownWindowsEventLogDataSourceStreams("Microsoft-WindowsEvent").ToPtr()},
							XPathQueries: []*string{
								to.StringPtr("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
								to.StringPtr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
						}},
				},
				Destinations: &armmonitor.DataCollectionRuleDestinations{
					LogAnalytics: []*armmonitor.LogAnalyticsDestination{
						{
							Name:                to.StringPtr("<name>"),
							WorkspaceResourceID: to.StringPtr("<workspace-resource-id>"),
						}},
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionRulesClientCreateResult)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesUpdate.json
func ExampleDataCollectionRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		&armmonitor.DataCollectionRulesClientUpdateOptions{Body: &armmonitor.ResourceForUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("A"),
				"tag2": to.StringPtr("B"),
				"tag3": to.StringPtr("C"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionRulesClientUpdateResult)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesDelete.json
func ExampleDataCollectionRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

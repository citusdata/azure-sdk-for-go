//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CloudServicePublicIpListAll.json
func ExamplePublicIPAddressesClient_ListCloudServicePublicIPAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.ListCloudServicePublicIPAddresses("<resource-group-name>",
		"<cloud-service-name>",
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CloudServiceRoleInstancePublicIpList.json
func ExamplePublicIPAddressesClient_ListCloudServiceRoleInstancePublicIPAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.ListCloudServiceRoleInstancePublicIPAddresses("<resource-group-name>",
		"<cloud-service-name>",
		"<role-instance-name>",
		"<network-interface-name>",
		"<ip-configuration-name>",
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CloudServicePublicIpGet.json
func ExamplePublicIPAddressesClient_GetCloudServicePublicIPAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	res, err := client.GetCloudServicePublicIPAddress(ctx,
		"<resource-group-name>",
		"<cloud-service-name>",
		"<role-instance-name>",
		"<network-interface-name>",
		"<ip-configuration-name>",
		"<public-ipaddress-name>",
		&armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientGetCloudServicePublicIPAddressResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressDelete.json
func ExamplePublicIPAddressesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<public-ipaddress-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressGet.json
func ExamplePublicIPAddressesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<public-ipaddress-name>",
		&armnetwork.PublicIPAddressesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientGetResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressCreateDns.json
func ExamplePublicIPAddressesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<public-ipaddress-name>",
		armnetwork.PublicIPAddress{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
				DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr("<domain-name-label>"),
				},
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
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressUpdateTags.json
func ExamplePublicIPAddressesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<public-ipaddress-name>",
		armnetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientUpdateTagsResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressListAll.json
func ExamplePublicIPAddressesClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.ListAll(nil)
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PublicIpAddressList.json
func ExamplePublicIPAddressesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssPublicIpListAll.json
func ExamplePublicIPAddressesClient_ListVirtualMachineScaleSetPublicIPAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.ListVirtualMachineScaleSetPublicIPAddresses("<resource-group-name>",
		"<virtual-machine-scale-set-name>",
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssVmPublicIpList.json
func ExamplePublicIPAddressesClient_ListVirtualMachineScaleSetVMPublicIPAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	pager := client.ListVirtualMachineScaleSetVMPublicIPAddresses("<resource-group-name>",
		"<virtual-machine-scale-set-name>",
		"<virtualmachine-index>",
		"<network-interface-name>",
		"<ip-configuration-name>",
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssPublicIpGet.json
func ExamplePublicIPAddressesClient_GetVirtualMachineScaleSetPublicIPAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	res, err := client.GetVirtualMachineScaleSetPublicIPAddress(ctx,
		"<resource-group-name>",
		"<virtual-machine-scale-set-name>",
		"<virtualmachine-index>",
		"<network-interface-name>",
		"<ip-configuration-name>",
		"<public-ipaddress-name>",
		&armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressResult)
}

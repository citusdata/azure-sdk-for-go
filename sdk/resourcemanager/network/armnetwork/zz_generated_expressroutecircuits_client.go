//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteCircuitsClient contains the methods for the ExpressRouteCircuits group.
// Don't use this type directly, use NewExpressRouteCircuitsClient() instead.
type ExpressRouteCircuitsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCircuitsClient creates a new instance of ExpressRouteCircuitsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCircuitsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteCircuitsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExpressRouteCircuitsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the circuit.
// parameters - Parameters supplied to the create or update express route circuit operation.
// options - ExpressRouteCircuitsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCircuitsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCircuitsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsClientBeginCreateOrUpdateOptions) (ExpressRouteCircuitsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return ExpressRouteCircuitsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteCircuitsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCircuitsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ExpressRouteCircuitsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCircuitsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCircuitsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// options - ExpressRouteCircuitsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteCircuitsClient.BeginDelete
// method.
func (client *ExpressRouteCircuitsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientBeginDeleteOptions) (ExpressRouteCircuitsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return ExpressRouteCircuitsClientDeletePollerResponse{}, err
	}
	result := ExpressRouteCircuitsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCircuitsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return ExpressRouteCircuitsClientDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCircuitsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitsClient) deleteOperation(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRouteCircuitsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of express route circuit.
// options - ExpressRouteCircuitsClientGetOptions contains the optional parameters for the ExpressRouteCircuitsClient.Get
// method.
func (client *ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientGetOptions) (ExpressRouteCircuitsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return ExpressRouteCircuitsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCircuitsClient) getHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientGetResponse, error) {
	result := ExpressRouteCircuitsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuit); err != nil {
		return ExpressRouteCircuitsClientGetResponse{}, err
	}
	return result, nil
}

// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// options - ExpressRouteCircuitsClientGetPeeringStatsOptions contains the optional parameters for the ExpressRouteCircuitsClient.GetPeeringStats
// method.
func (client *ExpressRouteCircuitsClient) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsClientGetPeeringStatsOptions) (ExpressRouteCircuitsClientGetPeeringStatsResponse, error) {
	req, err := client.getPeeringStatsCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return ExpressRouteCircuitsClientGetPeeringStatsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitsClientGetPeeringStatsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitsClientGetPeeringStatsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPeeringStatsHandleResponse(resp)
}

// getPeeringStatsCreateRequest creates the GetPeeringStats request.
func (client *ExpressRouteCircuitsClient) getPeeringStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsClientGetPeeringStatsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPeeringStatsHandleResponse handles the GetPeeringStats response.
func (client *ExpressRouteCircuitsClient) getPeeringStatsHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientGetPeeringStatsResponse, error) {
	result := ExpressRouteCircuitsClientGetPeeringStatsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitStats); err != nil {
		return ExpressRouteCircuitsClientGetPeeringStatsResponse{}, err
	}
	return result, nil
}

// GetStats - Gets all the stats from an express route circuit in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// options - ExpressRouteCircuitsClientGetStatsOptions contains the optional parameters for the ExpressRouteCircuitsClient.GetStats
// method.
func (client *ExpressRouteCircuitsClient) GetStats(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientGetStatsOptions) (ExpressRouteCircuitsClientGetStatsResponse, error) {
	req, err := client.getStatsCreateRequest(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return ExpressRouteCircuitsClientGetStatsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitsClientGetStatsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitsClientGetStatsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStatsHandleResponse(resp)
}

// getStatsCreateRequest creates the GetStats request.
func (client *ExpressRouteCircuitsClient) getStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsClientGetStatsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStatsHandleResponse handles the GetStats response.
func (client *ExpressRouteCircuitsClient) getStatsHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientGetStatsResponse, error) {
	result := ExpressRouteCircuitsClientGetStatsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitStats); err != nil {
		return ExpressRouteCircuitsClientGetStatsResponse{}, err
	}
	return result, nil
}

// List - Gets all the express route circuits in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ExpressRouteCircuitsClientListOptions contains the optional parameters for the ExpressRouteCircuitsClient.List
// method.
func (client *ExpressRouteCircuitsClient) List(resourceGroupName string, options *ExpressRouteCircuitsClientListOptions) *ExpressRouteCircuitsClientListPager {
	return &ExpressRouteCircuitsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCircuitsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCircuitsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCircuitsClient) listHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientListResponse, error) {
	result := ExpressRouteCircuitsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitListResult); err != nil {
		return ExpressRouteCircuitsClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all the express route circuits in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExpressRouteCircuitsClientListAllOptions contains the optional parameters for the ExpressRouteCircuitsClient.ListAll
// method.
func (client *ExpressRouteCircuitsClient) ListAll(options *ExpressRouteCircuitsClientListAllOptions) *ExpressRouteCircuitsClientListAllPager {
	return &ExpressRouteCircuitsClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCircuitsClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ExpressRouteCircuitsClient) listAllCreateRequest(ctx context.Context, options *ExpressRouteCircuitsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ExpressRouteCircuitsClient) listAllHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientListAllResponse, error) {
	result := ExpressRouteCircuitsClientListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitListResult); err != nil {
		return ExpressRouteCircuitsClientListAllResponse{}, err
	}
	return result, nil
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCircuitsClientBeginListArpTableOptions contains the optional parameters for the ExpressRouteCircuitsClient.BeginListArpTable
// method.
func (client *ExpressRouteCircuitsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListArpTableOptions) (ExpressRouteCircuitsClientListArpTablePollerResponse, error) {
	resp, err := client.listArpTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCircuitsClientListArpTablePollerResponse{}, err
	}
	result := ExpressRouteCircuitsClientListArpTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCircuitsClient.ListArpTable", "location", resp, client.pl)
	if err != nil {
		return ExpressRouteCircuitsClientListArpTablePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCircuitsClientListArpTablePoller{
		pt: pt,
	}
	return result, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitsClient) listArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListArpTableOptions) (*http.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCircuitsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListArpTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCircuitsClientBeginListRoutesTableOptions contains the optional parameters for the ExpressRouteCircuitsClient.BeginListRoutesTable
// method.
func (client *ExpressRouteCircuitsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableOptions) (ExpressRouteCircuitsClientListRoutesTablePollerResponse, error) {
	resp, err := client.listRoutesTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCircuitsClientListRoutesTablePollerResponse{}, err
	}
	result := ExpressRouteCircuitsClientListRoutesTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCircuitsClient.ListRoutesTable", "location", resp, client.pl)
	if err != nil {
		return ExpressRouteCircuitsClientListRoutesTablePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCircuitsClientListRoutesTablePoller{
		pt: pt,
	}
	return result, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitsClient) listRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableOptions) (*http.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCircuitsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit
// in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCircuitsClientBeginListRoutesTableSummaryOptions contains the optional parameters for the ExpressRouteCircuitsClient.BeginListRoutesTableSummary
// method.
func (client *ExpressRouteCircuitsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableSummaryOptions) (ExpressRouteCircuitsClientListRoutesTableSummaryPollerResponse, error) {
	resp, err := client.listRoutesTableSummary(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCircuitsClientListRoutesTableSummaryPollerResponse{}, err
	}
	result := ExpressRouteCircuitsClientListRoutesTableSummaryPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCircuitsClient.ListRoutesTableSummary", "location", resp, client.pl)
	if err != nil {
		return ExpressRouteCircuitsClientListRoutesTableSummaryPollerResponse{}, err
	}
	result.Poller = &ExpressRouteCircuitsClientListRoutesTableSummaryPoller{
		pt: pt,
	}
	return result, nil
}

// ListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in
// a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitsClient) listRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableSummaryOptions) (*http.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCircuitsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsClientBeginListRoutesTableSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// UpdateTags - Updates an express route circuit tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the circuit.
// parameters - Parameters supplied to update express route circuit tags.
// options - ExpressRouteCircuitsClientUpdateTagsOptions contains the optional parameters for the ExpressRouteCircuitsClient.UpdateTags
// method.
func (client *ExpressRouteCircuitsClient) UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsClientUpdateTagsOptions) (ExpressRouteCircuitsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return ExpressRouteCircuitsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteCircuitsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCircuitsClient) updateTagsHandleResponse(resp *http.Response) (ExpressRouteCircuitsClientUpdateTagsResponse, error) {
	result := ExpressRouteCircuitsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuit); err != nil {
		return ExpressRouteCircuitsClientUpdateTagsResponse{}, err
	}
	return result, nil
}

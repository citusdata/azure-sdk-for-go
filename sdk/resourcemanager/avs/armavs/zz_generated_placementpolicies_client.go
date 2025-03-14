//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// PlacementPoliciesClient contains the methods for the PlacementPolicies group.
// Don't use this type directly, use NewPlacementPoliciesClient() instead.
type PlacementPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPlacementPoliciesClient creates a new instance of PlacementPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPlacementPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PlacementPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PlacementPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// placementPolicyName - Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
// placementPolicy - A placement policy in the private cloud cluster
// options - PlacementPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the PlacementPoliciesClient.BeginCreateOrUpdate
// method.
func (client *PlacementPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, options *PlacementPoliciesClientBeginCreateOrUpdateOptions) (PlacementPoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicy, options)
	if err != nil {
		return PlacementPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := PlacementPoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PlacementPoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return PlacementPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PlacementPoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PlacementPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, options *PlacementPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicy, options)
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
func (client *PlacementPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, options *PlacementPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, placementPolicy)
}

// BeginDelete - Delete a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// placementPolicyName - Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
// options - PlacementPoliciesClientBeginDeleteOptions contains the optional parameters for the PlacementPoliciesClient.BeginDelete
// method.
func (client *PlacementPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientBeginDeleteOptions) (PlacementPoliciesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
	if err != nil {
		return PlacementPoliciesClientDeletePollerResponse{}, err
	}
	result := PlacementPoliciesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PlacementPoliciesClient.Delete", "", resp, client.pl)
	if err != nil {
		return PlacementPoliciesClientDeletePollerResponse{}, err
	}
	result.Poller = &PlacementPoliciesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PlacementPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
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
func (client *PlacementPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a placement policy by name in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// placementPolicyName - Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
// options - PlacementPoliciesClientGetOptions contains the optional parameters for the PlacementPoliciesClient.Get method.
func (client *PlacementPoliciesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientGetOptions) (PlacementPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
	if err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PlacementPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PlacementPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PlacementPoliciesClient) getHandleResponse(resp *http.Response) (PlacementPoliciesClientGetResponse, error) {
	result := PlacementPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlacementPolicy); err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// List - List placement policies in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// options - PlacementPoliciesClientListOptions contains the optional parameters for the PlacementPoliciesClient.List method.
func (client *PlacementPoliciesClient) List(resourceGroupName string, privateCloudName string, clusterName string, options *PlacementPoliciesClientListOptions) *PlacementPoliciesClientListPager {
	return &PlacementPoliciesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp PlacementPoliciesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PlacementPoliciesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PlacementPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *PlacementPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PlacementPoliciesClient) listHandleResponse(resp *http.Response) (PlacementPoliciesClientListResponse, error) {
	result := PlacementPoliciesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlacementPoliciesList); err != nil {
		return PlacementPoliciesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// placementPolicyName - Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
// placementPolicyUpdate - The placement policy properties that may be updated
// options - PlacementPoliciesClientBeginUpdateOptions contains the optional parameters for the PlacementPoliciesClient.BeginUpdate
// method.
func (client *PlacementPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, options *PlacementPoliciesClientBeginUpdateOptions) (PlacementPoliciesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicyUpdate, options)
	if err != nil {
		return PlacementPoliciesClientUpdatePollerResponse{}, err
	}
	result := PlacementPoliciesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PlacementPoliciesClient.Update", "", resp, client.pl)
	if err != nil {
		return PlacementPoliciesClientUpdatePollerResponse{}, err
	}
	result.Poller = &PlacementPoliciesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update a placement policy in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PlacementPoliciesClient) update(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, options *PlacementPoliciesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicyUpdate, options)
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

// updateCreateRequest creates the Update request.
func (client *PlacementPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, options *PlacementPoliciesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, placementPolicyUpdate)
}

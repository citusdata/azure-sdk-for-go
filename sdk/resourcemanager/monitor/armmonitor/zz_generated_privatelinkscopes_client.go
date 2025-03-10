//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// PrivateLinkScopesClient contains the methods for the PrivateLinkScopes group.
// Don't use this type directly, use NewPrivateLinkScopesClient() instead.
type PrivateLinkScopesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkScopesClient creates a new instance of PrivateLinkScopesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkScopesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkScopesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateLinkScopesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates (or updates) a Azure Monitor PrivateLinkScope. Note: You cannot specify a different value for
// InstrumentationKey nor AppId in the Put operation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// azureMonitorPrivateLinkScopePayload - Properties that need to be specified to create or update a Azure Monitor PrivateLinkScope.
// options - PrivateLinkScopesClientCreateOrUpdateOptions contains the optional parameters for the PrivateLinkScopesClient.CreateOrUpdate
// method.
func (client *PrivateLinkScopesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload AzureMonitorPrivateLinkScope, options *PrivateLinkScopesClientCreateOrUpdateOptions) (PrivateLinkScopesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, scopeName, azureMonitorPrivateLinkScopePayload, options)
	if err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkScopesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload AzureMonitorPrivateLinkScope, options *PrivateLinkScopesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, azureMonitorPrivateLinkScopePayload)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateLinkScopesClient) createOrUpdateHandleResponse(resp *http.Response) (PrivateLinkScopesClientCreateOrUpdateResponse, error) {
	result := PrivateLinkScopesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorPrivateLinkScope); err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a Azure Monitor PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// options - PrivateLinkScopesClientBeginDeleteOptions contains the optional parameters for the PrivateLinkScopesClient.BeginDelete
// method.
func (client *PrivateLinkScopesClient) BeginDelete(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (PrivateLinkScopesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return PrivateLinkScopesClientDeletePollerResponse{}, err
	}
	result := PrivateLinkScopesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateLinkScopesClient.Delete", "", resp, client.pl)
	if err != nil {
		return PrivateLinkScopesClientDeletePollerResponse{}, err
	}
	result.Poller = &PrivateLinkScopesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Azure Monitor PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkScopesClient) deleteOperation(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scopeName, options)
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
func (client *PrivateLinkScopesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Returns a Azure Monitor PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// options - PrivateLinkScopesClientGetOptions contains the optional parameters for the PrivateLinkScopesClient.Get method.
func (client *PrivateLinkScopesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientGetOptions) (PrivateLinkScopesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopesClient) getHandleResponse(resp *http.Response) (PrivateLinkScopesClientGetResponse, error) {
	result := PrivateLinkScopesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorPrivateLinkScope); err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all Azure Monitor PrivateLinkScopes within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateLinkScopesClientListOptions contains the optional parameters for the PrivateLinkScopesClient.List method.
func (client *PrivateLinkScopesClient) List(options *PrivateLinkScopesClientListOptions) *PrivateLinkScopesClientListPager {
	return &PrivateLinkScopesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkScopesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AzureMonitorPrivateLinkScopeListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateLinkScopesClient) listCreateRequest(ctx context.Context, options *PrivateLinkScopesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/privateLinkScopes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkScopesClient) listHandleResponse(resp *http.Response) (PrivateLinkScopesClientListResponse, error) {
	result := PrivateLinkScopesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorPrivateLinkScopeListResult); err != nil {
		return PrivateLinkScopesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets a list of Azure Monitor PrivateLinkScopes within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PrivateLinkScopesClientListByResourceGroupOptions contains the optional parameters for the PrivateLinkScopesClient.ListByResourceGroup
// method.
func (client *PrivateLinkScopesClient) ListByResourceGroup(resourceGroupName string, options *PrivateLinkScopesClientListByResourceGroupOptions) *PrivateLinkScopesClientListByResourceGroupPager {
	return &PrivateLinkScopesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkScopesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AzureMonitorPrivateLinkScopeListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkScopesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkScopesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes"
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
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkScopesClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkScopesClientListByResourceGroupResponse, error) {
	result := PrivateLinkScopesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorPrivateLinkScopeListResult); err != nil {
		return PrivateLinkScopesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// privateLinkScopeTags - Updated tag information to set into the PrivateLinkScope instance.
// options - PrivateLinkScopesClientUpdateTagsOptions contains the optional parameters for the PrivateLinkScopesClient.UpdateTags
// method.
func (client *PrivateLinkScopesClient) UpdateTags(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesClientUpdateTagsOptions) (PrivateLinkScopesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, scopeName, privateLinkScopeTags, options)
	if err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PrivateLinkScopesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateLinkScopeTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PrivateLinkScopesClient) updateTagsHandleResponse(resp *http.Response) (PrivateLinkScopesClientUpdateTagsResponse, error) {
	result := PrivateLinkScopesClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorPrivateLinkScope); err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	return result, nil
}

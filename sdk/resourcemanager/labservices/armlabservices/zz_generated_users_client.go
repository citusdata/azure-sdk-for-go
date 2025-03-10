//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// UsersClient contains the methods for the Users group.
// Don't use this type directly, use NewUsersClient() instead.
type UsersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUsersClient creates a new instance of UsersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UsersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &UsersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Operation to create or update a lab user.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// userName - The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
// body - The request body.
// options - UsersClientBeginCreateOrUpdateOptions contains the optional parameters for the UsersClient.BeginCreateOrUpdate
// method.
func (client *UsersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersClientBeginCreateOrUpdateOptions) (UsersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersClientCreateOrUpdatePollerResponse{}, err
	}
	result := UsersClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.CreateOrUpdate", "original-uri", resp, client.pl)
	if err != nil {
		return UsersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &UsersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Operation to create or update a lab user.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *UsersClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *UsersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Operation to delete a user resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// userName - The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
// options - UsersClientBeginDeleteOptions contains the optional parameters for the UsersClient.BeginDelete method.
func (client *UsersClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersClientBeginDeleteOptions) (UsersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, options)
	if err != nil {
		return UsersClientDeletePollerResponse{}, err
	}
	result := UsersClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Delete", "location", resp, client.pl)
	if err != nil {
		return UsersClientDeletePollerResponse{}, err
	}
	result.Poller = &UsersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a user resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *UsersClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, options)
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
func (client *UsersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of a lab user.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// userName - The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
// options - UsersClientGetOptions contains the optional parameters for the UsersClient.Get method.
func (client *UsersClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersClientGetOptions) (UsersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, options)
	if err != nil {
		return UsersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UsersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UsersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UsersClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UsersClient) getHandleResponse(resp *http.Response) (UsersClientGetResponse, error) {
	result := UsersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return UsersClientGetResponse{}, err
	}
	return result, nil
}

// BeginInvite - Operation to invite a user to a lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// userName - The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
// body - The request body.
// options - UsersClientBeginInviteOptions contains the optional parameters for the UsersClient.BeginInvite method.
func (client *UsersClient) BeginInvite(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersClientBeginInviteOptions) (UsersClientInvitePollerResponse, error) {
	resp, err := client.invite(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersClientInvitePollerResponse{}, err
	}
	result := UsersClientInvitePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Invite", "location", resp, client.pl)
	if err != nil {
		return UsersClientInvitePollerResponse{}, err
	}
	result.Poller = &UsersClientInvitePoller{
		pt: pt,
	}
	return result, nil
}

// Invite - Operation to invite a user to a lab.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *UsersClient) invite(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersClientBeginInviteOptions) (*http.Response, error) {
	req, err := client.inviteCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
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

// inviteCreateRequest creates the Invite request.
func (client *UsersClient) inviteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersClientBeginInviteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}/invite"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// ListByLab - Returns a list of all users for a lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - UsersClientListByLabOptions contains the optional parameters for the UsersClient.ListByLab method.
func (client *UsersClient) ListByLab(resourceGroupName string, labName string, options *UsersClientListByLabOptions) *UsersClientListByLabPager {
	return &UsersClientListByLabPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLabCreateRequest(ctx, resourceGroupName, labName, options)
		},
		advancer: func(ctx context.Context, resp UsersClientListByLabResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PagedUsers.NextLink)
		},
	}
}

// listByLabCreateRequest creates the ListByLab request.
func (client *UsersClient) listByLabCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *UsersClientListByLabOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLabHandleResponse handles the ListByLab response.
func (client *UsersClient) listByLabHandleResponse(resp *http.Response) (UsersClientListByLabResponse, error) {
	result := UsersClientListByLabResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUsers); err != nil {
		return UsersClientListByLabResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update a lab user.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// userName - The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
// body - The request body.
// options - UsersClientBeginUpdateOptions contains the optional parameters for the UsersClient.BeginUpdate method.
func (client *UsersClient) BeginUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersClientBeginUpdateOptions) (UsersClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersClientUpdatePollerResponse{}, err
	}
	result := UsersClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Update", "location", resp, client.pl)
	if err != nil {
		return UsersClientUpdatePollerResponse{}, err
	}
	result.Poller = &UsersClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update a lab user.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *UsersClient) update(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
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
func (client *UsersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

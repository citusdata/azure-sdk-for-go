//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// CommitmentPlansClient contains the methods for the CommitmentPlans group.
// Don't use this type directly, use NewCommitmentPlansClient() instead.
type CommitmentPlansClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCommitmentPlansClient creates a new instance of CommitmentPlansClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCommitmentPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CommitmentPlansClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &CommitmentPlansClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Update the state of specified commitmentPlans associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// commitmentPlanName - The name of the commitmentPlan associated with the Cognitive Services Account
// commitmentPlan - The commitmentPlan properties.
// options - CommitmentPlansClientCreateOrUpdateOptions contains the optional parameters for the CommitmentPlansClient.CreateOrUpdate
// method.
func (client *CommitmentPlansClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, commitmentPlan CommitmentPlan, options *CommitmentPlansClientCreateOrUpdateOptions) (CommitmentPlansClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, commitmentPlanName, commitmentPlan, options)
	if err != nil {
		return CommitmentPlansClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommitmentPlansClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CommitmentPlansClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CommitmentPlansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, commitmentPlan CommitmentPlan, options *CommitmentPlansClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if commitmentPlanName == "" {
		return nil, errors.New("parameter commitmentPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commitmentPlanName}", url.PathEscape(commitmentPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, commitmentPlan)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CommitmentPlansClient) createOrUpdateHandleResponse(resp *http.Response) (CommitmentPlansClientCreateOrUpdateResponse, error) {
	result := CommitmentPlansClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitmentPlan); err != nil {
		return CommitmentPlansClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified commitmentPlan associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// commitmentPlanName - The name of the commitmentPlan associated with the Cognitive Services Account
// options - CommitmentPlansClientBeginDeleteOptions contains the optional parameters for the CommitmentPlansClient.BeginDelete
// method.
func (client *CommitmentPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, options *CommitmentPlansClientBeginDeleteOptions) (CommitmentPlansClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, commitmentPlanName, options)
	if err != nil {
		return CommitmentPlansClientDeletePollerResponse{}, err
	}
	result := CommitmentPlansClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CommitmentPlansClient.Delete", "", resp, client.pl)
	if err != nil {
		return CommitmentPlansClientDeletePollerResponse{}, err
	}
	result.Poller = &CommitmentPlansClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified commitmentPlan associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CommitmentPlansClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, options *CommitmentPlansClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, commitmentPlanName, options)
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
func (client *CommitmentPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, options *CommitmentPlansClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if commitmentPlanName == "" {
		return nil, errors.New("parameter commitmentPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commitmentPlanName}", url.PathEscape(commitmentPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified commitmentPlans associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// commitmentPlanName - The name of the commitmentPlan associated with the Cognitive Services Account
// options - CommitmentPlansClientGetOptions contains the optional parameters for the CommitmentPlansClient.Get method.
func (client *CommitmentPlansClient) Get(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, options *CommitmentPlansClientGetOptions) (CommitmentPlansClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, commitmentPlanName, options)
	if err != nil {
		return CommitmentPlansClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommitmentPlansClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommitmentPlansClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommitmentPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, options *CommitmentPlansClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if commitmentPlanName == "" {
		return nil, errors.New("parameter commitmentPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commitmentPlanName}", url.PathEscape(commitmentPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommitmentPlansClient) getHandleResponse(resp *http.Response) (CommitmentPlansClientGetResponse, error) {
	result := CommitmentPlansClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitmentPlan); err != nil {
		return CommitmentPlansClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the commitmentPlans associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of Cognitive Services account.
// options - CommitmentPlansClientListOptions contains the optional parameters for the CommitmentPlansClient.List method.
func (client *CommitmentPlansClient) List(resourceGroupName string, accountName string, options *CommitmentPlansClientListOptions) *CommitmentPlansClientListPager {
	return &CommitmentPlansClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp CommitmentPlansClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CommitmentPlanListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *CommitmentPlansClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *CommitmentPlansClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CommitmentPlansClient) listHandleResponse(resp *http.Response) (CommitmentPlansClientListResponse, error) {
	result := CommitmentPlansClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitmentPlanListResult); err != nil {
		return CommitmentPlansClientListResponse{}, err
	}
	return result, nil
}

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsight

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

// ThreatIntelligenceIndicatorMetricsClient contains the methods for the ThreatIntelligenceIndicatorMetrics group.
// Don't use this type directly, use NewThreatIntelligenceIndicatorMetricsClient() instead.
type ThreatIntelligenceIndicatorMetricsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewThreatIntelligenceIndicatorMetricsClient creates a new instance of ThreatIntelligenceIndicatorMetricsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewThreatIntelligenceIndicatorMetricsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ThreatIntelligenceIndicatorMetricsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ThreatIntelligenceIndicatorMetricsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Get threat intelligence indicators metrics (Indicators counts by Type, Threat Type, Source).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - ThreatIntelligenceIndicatorMetricsClientListOptions contains the optional parameters for the ThreatIntelligenceIndicatorMetricsClient.List
// method.
func (client *ThreatIntelligenceIndicatorMetricsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *ThreatIntelligenceIndicatorMetricsClientListOptions) (ThreatIntelligenceIndicatorMetricsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return ThreatIntelligenceIndicatorMetricsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorMetricsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ThreatIntelligenceIndicatorMetricsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ThreatIntelligenceIndicatorMetricsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ThreatIntelligenceIndicatorMetricsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ThreatIntelligenceIndicatorMetricsClient) listHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorMetricsClientListResponse, error) {
	result := ThreatIntelligenceIndicatorMetricsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThreatIntelligenceMetricsList); err != nil {
		return ThreatIntelligenceIndicatorMetricsClientListResponse{}, err
	}
	return result, nil
}

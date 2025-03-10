//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

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

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiagnosticSettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *DiagnosticSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DiagnosticSettingsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates diagnostic settings for AadIam.
// If the operation fails it returns an *azcore.ResponseError type.
// name - The name of the diagnostic setting.
// parameters - Parameters supplied to the operation.
// options - DiagnosticSettingsClientCreateOrUpdateOptions contains the optional parameters for the DiagnosticSettingsClient.CreateOrUpdate
// method.
func (client *DiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsClientCreateOrUpdateOptions) (DiagnosticSettingsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, name, parameters, options)
	if err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticSettingsClient) createOrUpdateCreateRequest(ctx context.Context, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/microsoft.aadiam/diagnosticSettings/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DiagnosticSettingsClientCreateOrUpdateResponse, error) {
	result := DiagnosticSettingsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes existing diagnostic setting for AadIam.
// If the operation fails it returns an *azcore.ResponseError type.
// name - The name of the diagnostic setting.
// options - DiagnosticSettingsClientDeleteOptions contains the optional parameters for the DiagnosticSettingsClient.Delete
// method.
func (client *DiagnosticSettingsClient) Delete(ctx context.Context, name string, options *DiagnosticSettingsClientDeleteOptions) (DiagnosticSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, name, options)
	if err != nil {
		return DiagnosticSettingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DiagnosticSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DiagnosticSettingsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticSettingsClient) deleteCreateRequest(ctx context.Context, name string, options *DiagnosticSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/microsoft.aadiam/diagnosticSettings/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the active diagnostic setting for AadIam.
// If the operation fails it returns an *azcore.ResponseError type.
// name - The name of the diagnostic setting.
// options - DiagnosticSettingsClientGetOptions contains the optional parameters for the DiagnosticSettingsClient.Get method.
func (client *DiagnosticSettingsClient) Get(ctx context.Context, name string, options *DiagnosticSettingsClientGetOptions) (DiagnosticSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticSettingsClient) getCreateRequest(ctx context.Context, name string, options *DiagnosticSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/microsoft.aadiam/diagnosticSettings/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticSettingsClient) getHandleResponse(resp *http.Response) (DiagnosticSettingsClientGetResponse, error) {
	result := DiagnosticSettingsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the active diagnostic settings list for AadIam.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DiagnosticSettingsClientListOptions contains the optional parameters for the DiagnosticSettingsClient.List method.
func (client *DiagnosticSettingsClient) List(ctx context.Context, options *DiagnosticSettingsClientListOptions) (DiagnosticSettingsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DiagnosticSettingsClient) listCreateRequest(ctx context.Context, options *DiagnosticSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/microsoft.aadiam/diagnosticSettings"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiagnosticSettingsClient) listHandleResponse(resp *http.Response) (DiagnosticSettingsClientListResponse, error) {
	result := DiagnosticSettingsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResourceCollection); err != nil {
		return DiagnosticSettingsClientListResponse{}, err
	}
	return result, nil
}

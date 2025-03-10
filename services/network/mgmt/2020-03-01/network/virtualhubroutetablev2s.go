package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualHubRouteTableV2sClient is the network Client
type VirtualHubRouteTableV2sClient struct {
	BaseClient
}

// NewVirtualHubRouteTableV2sClient creates an instance of the VirtualHubRouteTableV2sClient client.
func NewVirtualHubRouteTableV2sClient(subscriptionID string) VirtualHubRouteTableV2sClient {
	return NewVirtualHubRouteTableV2sClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualHubRouteTableV2sClientWithBaseURI creates an instance of the VirtualHubRouteTableV2sClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVirtualHubRouteTableV2sClientWithBaseURI(baseURI string, subscriptionID string) VirtualHubRouteTableV2sClient {
	return VirtualHubRouteTableV2sClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing
// VirtualHubRouteTableV2.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the VirtualHubRouteTableV2.
// virtualHubRouteTableV2Parameters - parameters supplied to create or update VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2sClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2) (result VirtualHubRouteTableV2sCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubRouteTableV2sClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualHubRouteTableV2sClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	virtualHubRouteTableV2Parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", pathParameters),
		autorest.WithJSON(virtualHubRouteTableV2Parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubRouteTableV2sClient) CreateOrUpdateSender(req *http.Request) (future VirtualHubRouteTableV2sCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualHubRouteTableV2sClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualHubRouteTableV2, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a VirtualHubRouteTableV2.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHubRouteTableV2.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2sClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (result VirtualHubRouteTableV2sDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubRouteTableV2sClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualHubName, routeTableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualHubRouteTableV2sClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubRouteTableV2sClient) DeleteSender(req *http.Request) (future VirtualHubRouteTableV2sDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualHubRouteTableV2sClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a VirtualHubRouteTableV2.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHubRouteTableV2.
// virtualHubName - the name of the VirtualHub.
// routeTableName - the name of the VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2sClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (result VirtualHubRouteTableV2, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubRouteTableV2sClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, routeTableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualHubRouteTableV2sClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeTableName":    autorest.Encode("path", routeTableName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubRouteTableV2sClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualHubRouteTableV2sClient) GetResponder(resp *http.Response) (result VirtualHubRouteTableV2, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all VirtualHubRouteTableV2s.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client VirtualHubRouteTableV2sClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubRouteTableV2sResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubRouteTableV2sClient.List")
		defer func() {
			sc := -1
			if result.lvhrtvr.Response.Response != nil {
				sc = result.lvhrtvr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lvhrtvr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "List", resp, "Failure sending request")
		return
	}

	result.lvhrtvr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lvhrtvr.hasNextLink() && result.lvhrtvr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualHubRouteTableV2sClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubRouteTableV2sClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualHubRouteTableV2sClient) ListResponder(resp *http.Response) (result ListVirtualHubRouteTableV2sResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualHubRouteTableV2sClient) listNextResults(ctx context.Context, lastResults ListVirtualHubRouteTableV2sResult) (result ListVirtualHubRouteTableV2sResult, err error) {
	req, err := lastResults.listVirtualHubRouteTableV2sResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubRouteTableV2sClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualHubRouteTableV2sClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubRouteTableV2sResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubRouteTableV2sClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}

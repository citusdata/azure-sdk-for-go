package siterecovery

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

// ReplicationvCentersClient is the client for the ReplicationvCenters methods of the Siterecovery service.
type ReplicationvCentersClient struct {
	BaseClient
}

// NewReplicationvCentersClient creates an instance of the ReplicationvCentersClient client.
func NewReplicationvCentersClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationvCentersClient {
	return NewReplicationvCentersClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationvCentersClientWithBaseURI creates an instance of the ReplicationvCentersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewReplicationvCentersClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationvCentersClient {
	return ReplicationvCentersClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create a vCenter object..
// Parameters:
// fabricName - fabric name.
// vCenterName - vCenter name.
// addVCenterRequest - the input to the add vCenter operation.
func (client ReplicationvCentersClient) Create(ctx context.Context, fabricName string, vCenterName string, addVCenterRequest AddVCenterRequest) (result ReplicationvCentersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, fabricName, vCenterName, addVCenterRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationvCentersClient) CreatePreparer(ctx context.Context, fabricName string, vCenterName string, addVCenterRequest AddVCenterRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vCenterName":       autorest.Encode("path", vCenterName),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", pathParameters),
		autorest.WithJSON(addVCenterRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) CreateSender(req *http.Request) (future ReplicationvCentersCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationvCentersClient) CreateResponder(resp *http.Response) (result VCenter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to remove(unregister) a registered vCenter server from the vault.
// Parameters:
// fabricName - fabric name.
// vCenterName - vCenter name.
func (client ReplicationvCentersClient) Delete(ctx context.Context, fabricName string, vCenterName string) (result ReplicationvCentersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, fabricName, vCenterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationvCentersClient) DeletePreparer(ctx context.Context, fabricName string, vCenterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vCenterName":       autorest.Encode("path", vCenterName),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) DeleteSender(req *http.Request) (future ReplicationvCentersDeleteFuture, err error) {
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
func (client ReplicationvCentersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of a registered vCenter server(Add vCenter server.)
// Parameters:
// fabricName - fabric name.
// vCenterName - vCenter name.
func (client ReplicationvCentersClient) Get(ctx context.Context, fabricName string, vCenterName string) (result VCenter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, vCenterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationvCentersClient) GetPreparer(ctx context.Context, fabricName string, vCenterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vCenterName":       autorest.Encode("path", vCenterName),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationvCentersClient) GetResponder(resp *http.Response) (result VCenter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the vCenter servers registered in the vault.
func (client ReplicationvCentersClient) List(ctx context.Context) (result VCenterCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.List")
		defer func() {
			sc := -1
			if result.vcc.Response.Response != nil {
				sc = result.vcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "List", resp, "Failure sending request")
		return
	}

	result.vcc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vcc.hasNextLink() && result.vcc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationvCentersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationvCentersClient) ListResponder(resp *http.Response) (result VCenterCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationvCentersClient) listNextResults(ctx context.Context, lastResults VCenterCollection) (result VCenterCollection, err error) {
	req, err := lastResults.vCenterCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationvCentersClient) ListComplete(ctx context.Context) (result VCenterCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationFabrics lists the vCenter servers registered in a fabric.
// Parameters:
// fabricName - fabric name.
func (client ReplicationvCentersClient) ListByReplicationFabrics(ctx context.Context, fabricName string) (result VCenterCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.vcc.Response.Response != nil {
				sc = result.vcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationFabricsNextResults
	req, err := client.ListByReplicationFabricsPreparer(ctx, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.vcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result.vcc, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "ListByReplicationFabrics", resp, "Failure responding to request")
		return
	}
	if result.vcc.hasNextLink() && result.vcc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationvCentersClient) ListByReplicationFabricsPreparer(ctx context.Context, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationvCentersClient) ListByReplicationFabricsResponder(resp *http.Response) (result VCenterCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationvCentersClient) listByReplicationFabricsNextResults(ctx context.Context, lastResults VCenterCollection) (result VCenterCollection, err error) {
	req, err := lastResults.vCenterCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listByReplicationFabricsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listByReplicationFabricsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "listByReplicationFabricsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationFabricsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationvCentersClient) ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result VCenterCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationFabrics(ctx, fabricName)
	return
}

// Update the operation to update a registered vCenter.
// Parameters:
// fabricName - fabric name.
// vCenterName - vCenter name
// updateVCenterRequest - the input to the update vCenter operation.
func (client ReplicationvCentersClient) Update(ctx context.Context, fabricName string, vCenterName string, updateVCenterRequest UpdateVCenterRequest) (result ReplicationvCentersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationvCentersClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, fabricName, vCenterName, updateVCenterRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationvCentersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ReplicationvCentersClient) UpdatePreparer(ctx context.Context, fabricName string, vCenterName string, updateVCenterRequest UpdateVCenterRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vCenterName":       autorest.Encode("path", vCenterName),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", pathParameters),
		autorest.WithJSON(updateVCenterRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationvCentersClient) UpdateSender(req *http.Request) (future ReplicationvCentersUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ReplicationvCentersClient) UpdateResponder(resp *http.Response) (result VCenter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

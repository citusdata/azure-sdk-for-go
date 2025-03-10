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

// ServiceEndpointPolicyDefinitionsClient is the network Client
type ServiceEndpointPolicyDefinitionsClient struct {
	BaseClient
}

// NewServiceEndpointPolicyDefinitionsClient creates an instance of the ServiceEndpointPolicyDefinitionsClient client.
func NewServiceEndpointPolicyDefinitionsClient(subscriptionID string) ServiceEndpointPolicyDefinitionsClient {
	return NewServiceEndpointPolicyDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceEndpointPolicyDefinitionsClientWithBaseURI creates an instance of the
// ServiceEndpointPolicyDefinitionsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceEndpointPolicyDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ServiceEndpointPolicyDefinitionsClient {
	return ServiceEndpointPolicyDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a service endpoint policy definition in the specified service endpoint policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition name.
// serviceEndpointPolicyDefinitions - parameters supplied to the create or update service endpoint policy
// operation.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (result ServiceEndpointPolicyDefinitionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	serviceEndpointPolicyDefinitions.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithJSON(serviceEndpointPolicyDefinitions),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdateSender(req *http.Request) (future ServiceEndpointPolicyDefinitionsCreateOrUpdateFuture, err error) {
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
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceEndpointPolicyDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified ServiceEndpoint policy definitions.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the Service Endpoint Policy.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition.
func (client ServiceEndpointPolicyDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (result ServiceEndpointPolicyDefinitionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceEndpointPolicyDefinitionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) DeleteSender(req *http.Request) (future ServiceEndpointPolicyDefinitionsDeleteFuture, err error) {
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
func (client ServiceEndpointPolicyDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified service endpoint policy definitions from service endpoint policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy name.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition name.
func (client ServiceEndpointPolicyDefinitionsClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (result ServiceEndpointPolicyDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceEndpointPolicyDefinitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceEndpointPolicyDefinitionsClient) GetResponder(resp *http.Response) (result ServiceEndpointPolicyDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all service endpoint policy definitions in a service end point policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy name.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (result ServiceEndpointPolicyDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.sepdlr.Response.Response != nil {
				sc = result.sepdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, serviceEndpointPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.sepdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.sepdlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.sepdlr.hasNextLink() && result.sepdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyName": autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupResponder(resp *http.Response) (result ServiceEndpointPolicyDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ServiceEndpointPolicyDefinitionsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ServiceEndpointPolicyDefinitionListResult) (result ServiceEndpointPolicyDefinitionListResult, err error) {
	req, err := lastResults.serviceEndpointPolicyDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (result ServiceEndpointPolicyDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, serviceEndpointPolicyName)
	return
}

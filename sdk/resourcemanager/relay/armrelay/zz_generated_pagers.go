//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// HybridConnectionsClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type HybridConnectionsClientListAuthorizationRulesPager struct {
	client    *HybridConnectionsClient
	current   HybridConnectionsClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HybridConnectionsClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *HybridConnectionsClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *HybridConnectionsClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current HybridConnectionsClientListAuthorizationRulesResponse page.
func (p *HybridConnectionsClientListAuthorizationRulesPager) PageResponse() HybridConnectionsClientListAuthorizationRulesResponse {
	return p.current
}

// HybridConnectionsClientListByNamespacePager provides operations for iterating over paged responses.
type HybridConnectionsClientListByNamespacePager struct {
	client    *HybridConnectionsClient
	current   HybridConnectionsClientListByNamespaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HybridConnectionsClientListByNamespaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *HybridConnectionsClientListByNamespacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *HybridConnectionsClientListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HybridConnectionListResult.NextLink == nil || len(*p.current.HybridConnectionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current HybridConnectionsClientListByNamespaceResponse page.
func (p *HybridConnectionsClientListByNamespacePager) PageResponse() HybridConnectionsClientListByNamespaceResponse {
	return p.current
}

// NamespacesClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type NamespacesClientListAuthorizationRulesPager struct {
	client    *NamespacesClient
	current   NamespacesClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListAuthorizationRulesResponse page.
func (p *NamespacesClientListAuthorizationRulesPager) PageResponse() NamespacesClientListAuthorizationRulesResponse {
	return p.current
}

// NamespacesClientListByResourceGroupPager provides operations for iterating over paged responses.
type NamespacesClientListByResourceGroupPager struct {
	client    *NamespacesClient
	current   NamespacesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NamespaceListResult.NextLink == nil || len(*p.current.NamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListByResourceGroupResponse page.
func (p *NamespacesClientListByResourceGroupPager) PageResponse() NamespacesClientListByResourceGroupResponse {
	return p.current
}

// NamespacesClientListPager provides operations for iterating over paged responses.
type NamespacesClientListPager struct {
	client    *NamespacesClient
	current   NamespacesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NamespaceListResult.NextLink == nil || len(*p.current.NamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListResponse page.
func (p *NamespacesClientListPager) PageResponse() NamespacesClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// WCFRelaysClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type WCFRelaysClientListAuthorizationRulesPager struct {
	client    *WCFRelaysClient
	current   WCFRelaysClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WCFRelaysClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WCFRelaysClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WCFRelaysClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WCFRelaysClientListAuthorizationRulesResponse page.
func (p *WCFRelaysClientListAuthorizationRulesPager) PageResponse() WCFRelaysClientListAuthorizationRulesResponse {
	return p.current
}

// WCFRelaysClientListByNamespacePager provides operations for iterating over paged responses.
type WCFRelaysClientListByNamespacePager struct {
	client    *WCFRelaysClient
	current   WCFRelaysClientListByNamespaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WCFRelaysClientListByNamespaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WCFRelaysClientListByNamespacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WCFRelaysClientListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WcfRelaysListResult.NextLink == nil || len(*p.current.WcfRelaysListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WCFRelaysClientListByNamespaceResponse page.
func (p *WCFRelaysClientListByNamespacePager) PageResponse() WCFRelaysClientListByNamespaceResponse {
	return p.current
}

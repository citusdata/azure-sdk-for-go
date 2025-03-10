//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DiskPoolZonesClientListPager provides operations for iterating over paged responses.
type DiskPoolZonesClientListPager struct {
	client    *DiskPoolZonesClient
	current   DiskPoolZonesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DiskPoolZonesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DiskPoolZonesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DiskPoolZonesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DiskPoolZoneListResult.NextLink == nil || len(*p.current.DiskPoolZoneListResult.NextLink) == 0 {
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

// PageResponse returns the current DiskPoolZonesClientListResponse page.
func (p *DiskPoolZonesClientListPager) PageResponse() DiskPoolZonesClientListResponse {
	return p.current
}

// DiskPoolsClientListByResourceGroupPager provides operations for iterating over paged responses.
type DiskPoolsClientListByResourceGroupPager struct {
	client    *DiskPoolsClient
	current   DiskPoolsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DiskPoolsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DiskPoolsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DiskPoolsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DiskPoolListResult.NextLink == nil || len(*p.current.DiskPoolListResult.NextLink) == 0 {
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

// PageResponse returns the current DiskPoolsClientListByResourceGroupResponse page.
func (p *DiskPoolsClientListByResourceGroupPager) PageResponse() DiskPoolsClientListByResourceGroupResponse {
	return p.current
}

// DiskPoolsClientListBySubscriptionPager provides operations for iterating over paged responses.
type DiskPoolsClientListBySubscriptionPager struct {
	client    *DiskPoolsClient
	current   DiskPoolsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DiskPoolsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DiskPoolsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DiskPoolsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DiskPoolListResult.NextLink == nil || len(*p.current.DiskPoolListResult.NextLink) == 0 {
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
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DiskPoolsClientListBySubscriptionResponse page.
func (p *DiskPoolsClientListBySubscriptionPager) PageResponse() DiskPoolsClientListBySubscriptionResponse {
	return p.current
}

// DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager provides operations for iterating over paged responses.
type DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager struct {
	client    *DiskPoolsClient
	current   DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEnvironmentEndpointList.NextLink == nil || len(*p.current.OutboundEnvironmentEndpointList.NextLink) == 0 {
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
	result, err := p.client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse page.
func (p *DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager) PageResponse() DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse {
	return p.current
}

// IscsiTargetsClientListByDiskPoolPager provides operations for iterating over paged responses.
type IscsiTargetsClientListByDiskPoolPager struct {
	client    *IscsiTargetsClient
	current   IscsiTargetsClientListByDiskPoolResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IscsiTargetsClientListByDiskPoolResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IscsiTargetsClientListByDiskPoolPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IscsiTargetsClientListByDiskPoolPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IscsiTargetList.NextLink == nil || len(*p.current.IscsiTargetList.NextLink) == 0 {
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
	result, err := p.client.listByDiskPoolHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IscsiTargetsClientListByDiskPoolResponse page.
func (p *IscsiTargetsClientListByDiskPoolPager) PageResponse() IscsiTargetsClientListByDiskPoolResponse {
	return p.current
}

// ResourceSKUsClientListPager provides operations for iterating over paged responses.
type ResourceSKUsClientListPager struct {
	client    *ResourceSKUsClient
	current   ResourceSKUsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ResourceSKUsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ResourceSKUsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ResourceSKUsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceSKUListResult.NextLink == nil || len(*p.current.ResourceSKUListResult.NextLink) == 0 {
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

// PageResponse returns the current ResourceSKUsClientListResponse page.
func (p *ResourceSKUsClientListPager) PageResponse() ResourceSKUsClientListResponse {
	return p.current
}

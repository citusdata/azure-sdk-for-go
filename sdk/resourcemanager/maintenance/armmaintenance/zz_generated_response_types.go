//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

import "net/http"

// ApplyUpdateForResourceGroupClientListResponse contains the response from method ApplyUpdateForResourceGroupClient.List.
type ApplyUpdateForResourceGroupClientListResponse struct {
	ApplyUpdateForResourceGroupClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdateForResourceGroupClientListResult contains the result from method ApplyUpdateForResourceGroupClient.List.
type ApplyUpdateForResourceGroupClientListResult struct {
	ListApplyUpdate
}

// ApplyUpdatesClientCreateOrUpdateParentResponse contains the response from method ApplyUpdatesClient.CreateOrUpdateParent.
type ApplyUpdatesClientCreateOrUpdateParentResponse struct {
	ApplyUpdatesClientCreateOrUpdateParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdatesClientCreateOrUpdateParentResult contains the result from method ApplyUpdatesClient.CreateOrUpdateParent.
type ApplyUpdatesClientCreateOrUpdateParentResult struct {
	ApplyUpdate
}

// ApplyUpdatesClientCreateOrUpdateResponse contains the response from method ApplyUpdatesClient.CreateOrUpdate.
type ApplyUpdatesClientCreateOrUpdateResponse struct {
	ApplyUpdatesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdatesClientCreateOrUpdateResult contains the result from method ApplyUpdatesClient.CreateOrUpdate.
type ApplyUpdatesClientCreateOrUpdateResult struct {
	ApplyUpdate
}

// ApplyUpdatesClientGetParentResponse contains the response from method ApplyUpdatesClient.GetParent.
type ApplyUpdatesClientGetParentResponse struct {
	ApplyUpdatesClientGetParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdatesClientGetParentResult contains the result from method ApplyUpdatesClient.GetParent.
type ApplyUpdatesClientGetParentResult struct {
	ApplyUpdate
}

// ApplyUpdatesClientGetResponse contains the response from method ApplyUpdatesClient.Get.
type ApplyUpdatesClientGetResponse struct {
	ApplyUpdatesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdatesClientGetResult contains the result from method ApplyUpdatesClient.Get.
type ApplyUpdatesClientGetResult struct {
	ApplyUpdate
}

// ApplyUpdatesClientListResponse contains the response from method ApplyUpdatesClient.List.
type ApplyUpdatesClientListResponse struct {
	ApplyUpdatesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplyUpdatesClientListResult contains the result from method ApplyUpdatesClient.List.
type ApplyUpdatesClientListResult struct {
	ListApplyUpdate
}

// ConfigurationAssignmentsClientCreateOrUpdateParentResponse contains the response from method ConfigurationAssignmentsClient.CreateOrUpdateParent.
type ConfigurationAssignmentsClientCreateOrUpdateParentResponse struct {
	ConfigurationAssignmentsClientCreateOrUpdateParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientCreateOrUpdateParentResult contains the result from method ConfigurationAssignmentsClient.CreateOrUpdateParent.
type ConfigurationAssignmentsClientCreateOrUpdateParentResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientCreateOrUpdateResponse contains the response from method ConfigurationAssignmentsClient.CreateOrUpdate.
type ConfigurationAssignmentsClientCreateOrUpdateResponse struct {
	ConfigurationAssignmentsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientCreateOrUpdateResult contains the result from method ConfigurationAssignmentsClient.CreateOrUpdate.
type ConfigurationAssignmentsClientCreateOrUpdateResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientDeleteParentResponse contains the response from method ConfigurationAssignmentsClient.DeleteParent.
type ConfigurationAssignmentsClientDeleteParentResponse struct {
	ConfigurationAssignmentsClientDeleteParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientDeleteParentResult contains the result from method ConfigurationAssignmentsClient.DeleteParent.
type ConfigurationAssignmentsClientDeleteParentResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientDeleteResponse contains the response from method ConfigurationAssignmentsClient.Delete.
type ConfigurationAssignmentsClientDeleteResponse struct {
	ConfigurationAssignmentsClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientDeleteResult contains the result from method ConfigurationAssignmentsClient.Delete.
type ConfigurationAssignmentsClientDeleteResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientGetParentResponse contains the response from method ConfigurationAssignmentsClient.GetParent.
type ConfigurationAssignmentsClientGetParentResponse struct {
	ConfigurationAssignmentsClientGetParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientGetParentResult contains the result from method ConfigurationAssignmentsClient.GetParent.
type ConfigurationAssignmentsClientGetParentResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientGetResponse contains the response from method ConfigurationAssignmentsClient.Get.
type ConfigurationAssignmentsClientGetResponse struct {
	ConfigurationAssignmentsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientGetResult contains the result from method ConfigurationAssignmentsClient.Get.
type ConfigurationAssignmentsClientGetResult struct {
	ConfigurationAssignment
}

// ConfigurationAssignmentsClientListParentResponse contains the response from method ConfigurationAssignmentsClient.ListParent.
type ConfigurationAssignmentsClientListParentResponse struct {
	ConfigurationAssignmentsClientListParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientListParentResult contains the result from method ConfigurationAssignmentsClient.ListParent.
type ConfigurationAssignmentsClientListParentResult struct {
	ListConfigurationAssignmentsResult
}

// ConfigurationAssignmentsClientListResponse contains the response from method ConfigurationAssignmentsClient.List.
type ConfigurationAssignmentsClientListResponse struct {
	ConfigurationAssignmentsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsClientListResult contains the result from method ConfigurationAssignmentsClient.List.
type ConfigurationAssignmentsClientListResult struct {
	ListConfigurationAssignmentsResult
}

// ConfigurationAssignmentsWithinSubscriptionClientListResponse contains the response from method ConfigurationAssignmentsWithinSubscriptionClient.List.
type ConfigurationAssignmentsWithinSubscriptionClientListResponse struct {
	ConfigurationAssignmentsWithinSubscriptionClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationAssignmentsWithinSubscriptionClientListResult contains the result from method ConfigurationAssignmentsWithinSubscriptionClient.List.
type ConfigurationAssignmentsWithinSubscriptionClientListResult struct {
	ListConfigurationAssignmentsResult
}

// ConfigurationsClientCreateOrUpdateResponse contains the response from method ConfigurationsClient.CreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResponse struct {
	ConfigurationsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsClientCreateOrUpdateResult contains the result from method ConfigurationsClient.CreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResult struct {
	Configuration
}

// ConfigurationsClientDeleteResponse contains the response from method ConfigurationsClient.Delete.
type ConfigurationsClientDeleteResponse struct {
	ConfigurationsClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsClientDeleteResult contains the result from method ConfigurationsClient.Delete.
type ConfigurationsClientDeleteResult struct {
	Configuration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	ConfigurationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsClientGetResult contains the result from method ConfigurationsClient.Get.
type ConfigurationsClientGetResult struct {
	Configuration
}

// ConfigurationsClientListResponse contains the response from method ConfigurationsClient.List.
type ConfigurationsClientListResponse struct {
	ConfigurationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsClientListResult contains the result from method ConfigurationsClient.List.
type ConfigurationsClientListResult struct {
	ListMaintenanceConfigurationsResult
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.Update.
type ConfigurationsClientUpdateResponse struct {
	ConfigurationsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsClientUpdateResult contains the result from method ConfigurationsClient.Update.
type ConfigurationsClientUpdateResult struct {
	Configuration
}

// ConfigurationsForResourceGroupClientListResponse contains the response from method ConfigurationsForResourceGroupClient.List.
type ConfigurationsForResourceGroupClientListResponse struct {
	ConfigurationsForResourceGroupClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsForResourceGroupClientListResult contains the result from method ConfigurationsForResourceGroupClient.List.
type ConfigurationsForResourceGroupClientListResult struct {
	ListMaintenanceConfigurationsResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationsListResult
}

// PublicMaintenanceConfigurationsClientGetResponse contains the response from method PublicMaintenanceConfigurationsClient.Get.
type PublicMaintenanceConfigurationsClientGetResponse struct {
	PublicMaintenanceConfigurationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublicMaintenanceConfigurationsClientGetResult contains the result from method PublicMaintenanceConfigurationsClient.Get.
type PublicMaintenanceConfigurationsClientGetResult struct {
	Configuration
}

// PublicMaintenanceConfigurationsClientListResponse contains the response from method PublicMaintenanceConfigurationsClient.List.
type PublicMaintenanceConfigurationsClientListResponse struct {
	PublicMaintenanceConfigurationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublicMaintenanceConfigurationsClientListResult contains the result from method PublicMaintenanceConfigurationsClient.List.
type PublicMaintenanceConfigurationsClientListResult struct {
	ListMaintenanceConfigurationsResult
}

// UpdatesClientListParentResponse contains the response from method UpdatesClient.ListParent.
type UpdatesClientListParentResponse struct {
	UpdatesClientListParentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UpdatesClientListParentResult contains the result from method UpdatesClient.ListParent.
type UpdatesClientListParentResult struct {
	ListUpdatesResult
}

// UpdatesClientListResponse contains the response from method UpdatesClient.List.
type UpdatesClientListResponse struct {
	UpdatesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UpdatesClientListResult contains the result from method UpdatesClient.List.
type UpdatesClientListResult struct {
	ListUpdatesResult
}

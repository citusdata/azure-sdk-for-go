# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*PrivateEndpointConnectionsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *PrivateEndpointConnectionsBeginDeleteOptions)` to `(context.Context, string, string, string, *PrivateEndpointConnectionsClientBeginDeleteOptions)`
- Function `*PrivateEndpointConnectionsClient.BeginDelete` return value(s) have been changed from `(PrivateEndpointConnectionsDeletePollerResponse, error)` to `(PrivateEndpointConnectionsClientDeletePollerResponse, error)`
- Function `*PrivateLinkResourcesClient.ListByAccount` parameter(s) have been changed from `(string, string, *PrivateLinkResourcesListByAccountOptions)` to `(string, string, *PrivateLinkResourcesClientListByAccountOptions)`
- Function `*PrivateLinkResourcesClient.ListByAccount` return value(s) have been changed from `(*PrivateLinkResourcesListByAccountPager)` to `(*PrivateLinkResourcesClientListByAccountPager)`
- Function `*AccountsClient.CheckNameAvailability` parameter(s) have been changed from `(context.Context, CheckNameAvailabilityRequest, *AccountsCheckNameAvailabilityOptions)` to `(context.Context, CheckNameAvailabilityRequest, *AccountsClientCheckNameAvailabilityOptions)`
- Function `*AccountsClient.CheckNameAvailability` return value(s) have been changed from `(AccountsCheckNameAvailabilityResponse, error)` to `(AccountsClientCheckNameAvailabilityResponse, error)`
- Function `*PrivateEndpointConnectionsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *PrivateEndpointConnectionsGetOptions)` to `(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetOptions)`
- Function `*PrivateEndpointConnectionsClient.Get` return value(s) have been changed from `(PrivateEndpointConnectionsGetResponse, error)` to `(PrivateEndpointConnectionsClientGetResponse, error)`
- Function `*AccountsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *AccountsGetOptions)` to `(context.Context, string, string, *AccountsClientGetOptions)`
- Function `*AccountsClient.Get` return value(s) have been changed from `(AccountsGetResponse, error)` to `(AccountsClientGetResponse, error)`
- Function `*AccountsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, AccountUpdateParameters, *AccountsBeginUpdateOptions)` to `(context.Context, string, string, AccountUpdateParameters, *AccountsClientBeginUpdateOptions)`
- Function `*AccountsClient.BeginUpdate` return value(s) have been changed from `(AccountsUpdatePollerResponse, error)` to `(AccountsClientUpdatePollerResponse, error)`
- Function `*AccountsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *AccountsListByResourceGroupOptions)` to `(string, *AccountsClientListByResourceGroupOptions)`
- Function `*AccountsClient.ListByResourceGroup` return value(s) have been changed from `(*AccountsListByResourceGroupPager)` to `(*AccountsClientListByResourceGroupPager)`
- Function `*DefaultAccountsClient.Remove` parameter(s) have been changed from `(context.Context, string, ScopeType, *DefaultAccountsRemoveOptions)` to `(context.Context, string, ScopeType, *DefaultAccountsClientRemoveOptions)`
- Function `*DefaultAccountsClient.Remove` return value(s) have been changed from `(DefaultAccountsRemoveResponse, error)` to `(DefaultAccountsClientRemoveResponse, error)`
- Function `*PrivateLinkResourcesClient.GetByGroupID` parameter(s) have been changed from `(context.Context, string, string, string, *PrivateLinkResourcesGetByGroupIDOptions)` to `(context.Context, string, string, string, *PrivateLinkResourcesClientGetByGroupIDOptions)`
- Function `*PrivateLinkResourcesClient.GetByGroupID` return value(s) have been changed from `(PrivateLinkResourcesGetByGroupIDResponse, error)` to `(PrivateLinkResourcesClientGetByGroupIDResponse, error)`
- Function `*AccountsClient.ListKeys` parameter(s) have been changed from `(context.Context, string, string, *AccountsListKeysOptions)` to `(context.Context, string, string, *AccountsClientListKeysOptions)`
- Function `*AccountsClient.ListKeys` return value(s) have been changed from `(AccountsListKeysResponse, error)` to `(AccountsClientListKeysResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*AccountsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, Account, *AccountsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, Account, *AccountsClientBeginCreateOrUpdateOptions)`
- Function `*AccountsClient.BeginCreateOrUpdate` return value(s) have been changed from `(AccountsCreateOrUpdatePollerResponse, error)` to `(AccountsClientCreateOrUpdatePollerResponse, error)`
- Function `*DefaultAccountsClient.Set` parameter(s) have been changed from `(context.Context, DefaultAccountPayload, *DefaultAccountsSetOptions)` to `(context.Context, DefaultAccountPayload, *DefaultAccountsClientSetOptions)`
- Function `*DefaultAccountsClient.Set` return value(s) have been changed from `(DefaultAccountsSetResponse, error)` to `(DefaultAccountsClientSetResponse, error)`
- Function `*AccountsClient.AddRootCollectionAdmin` parameter(s) have been changed from `(context.Context, string, string, CollectionAdminUpdate, *AccountsAddRootCollectionAdminOptions)` to `(context.Context, string, string, CollectionAdminUpdate, *AccountsClientAddRootCollectionAdminOptions)`
- Function `*AccountsClient.AddRootCollectionAdmin` return value(s) have been changed from `(AccountsAddRootCollectionAdminResponse, error)` to `(AccountsClientAddRootCollectionAdminResponse, error)`
- Function `*AccountsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *AccountsBeginDeleteOptions)` to `(context.Context, string, string, *AccountsClientBeginDeleteOptions)`
- Function `*AccountsClient.BeginDelete` return value(s) have been changed from `(AccountsDeletePollerResponse, error)` to `(AccountsClientDeletePollerResponse, error)`
- Function `*DefaultAccountsClient.Get` parameter(s) have been changed from `(context.Context, string, ScopeType, *DefaultAccountsGetOptions)` to `(context.Context, string, ScopeType, *DefaultAccountsClientGetOptions)`
- Function `*DefaultAccountsClient.Get` return value(s) have been changed from `(DefaultAccountsGetResponse, error)` to `(DefaultAccountsClientGetResponse, error)`
- Function `*AccountsClient.ListBySubscription` parameter(s) have been changed from `(*AccountsListBySubscriptionOptions)` to `(*AccountsClientListBySubscriptionOptions)`
- Function `*AccountsClient.ListBySubscription` return value(s) have been changed from `(*AccountsListBySubscriptionPager)` to `(*AccountsClientListBySubscriptionPager)`
- Function `*PrivateEndpointConnectionsClient.ListByAccount` parameter(s) have been changed from `(string, string, *PrivateEndpointConnectionsListByAccountOptions)` to `(string, string, *PrivateEndpointConnectionsClientListByAccountOptions)`
- Function `*PrivateEndpointConnectionsClient.ListByAccount` return value(s) have been changed from `(*PrivateEndpointConnectionsListByAccountPager)` to `(*PrivateEndpointConnectionsClientListByAccountPager)`
- Function `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions)`
- Function `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate` return value(s) have been changed from `(PrivateEndpointConnectionsCreateOrUpdatePollerResponse, error)` to `(PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse, error)`
- Function `*OperationsListPager.NextPage` has been removed
- Function `*PrivateEndpointConnectionsListByAccountPager.PageResponse` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `*PrivateEndpointConnectionsListByAccountPager.NextPage` has been removed
- Function `*AccountsListBySubscriptionPager.NextPage` has been removed
- Function `AccountsDeletePollerResponse.PollUntilDone` has been removed
- Function `*AccountsDeletePollerResponse.Resume` has been removed
- Function `*AccountsUpdatePoller.Done` has been removed
- Function `*AccountsUpdatePollerResponse.Resume` has been removed
- Function `*PrivateEndpointConnectionsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*AccountsListByResourceGroupPager.NextPage` has been removed
- Function `*PrivateEndpointConnectionsDeletePoller.FinalResponse` has been removed
- Function `*AccountsListBySubscriptionPager.Err` has been removed
- Function `*AccountsUpdatePoller.FinalResponse` has been removed
- Function `*AccountsCreateOrUpdatePoller.Poll` has been removed
- Function `AccountsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*PrivateEndpointConnectionsCreateOrUpdatePoller.Done` has been removed
- Function `*AccountsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*PrivateEndpointConnectionsDeletePoller.Poll` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*AccountsDeletePoller.ResumeToken` has been removed
- Function `*AccountsListByResourceGroupPager.Err` has been removed
- Function `*AccountsDeletePoller.Poll` has been removed
- Function `*AccountsUpdatePoller.Poll` has been removed
- Function `*AccountsDeletePoller.Done` has been removed
- Function `*PrivateEndpointConnectionsDeletePollerResponse.Resume` has been removed
- Function `PrivateEndpointConnectionsDeletePollerResponse.PollUntilDone` has been removed
- Function `*AccountsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*AccountsDeletePoller.FinalResponse` has been removed
- Function `*AccountsListBySubscriptionPager.PageResponse` has been removed
- Function `AccountsUpdatePollerResponse.PollUntilDone` has been removed
- Function `ErrorResponseModel.Error` has been removed
- Function `*PrivateLinkResourcesListByAccountPager.PageResponse` has been removed
- Function `PrivateEndpointConnectionsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*AccountsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*PrivateEndpointConnectionsCreateOrUpdatePoller.Poll` has been removed
- Function `*PrivateEndpointConnectionsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*PrivateLinkResourcesListByAccountPager.NextPage` has been removed
- Function `*PrivateEndpointConnectionsDeletePoller.Done` has been removed
- Function `*PrivateEndpointConnectionsListByAccountPager.Err` has been removed
- Function `*AccountsListByResourceGroupPager.PageResponse` has been removed
- Function `*PrivateEndpointConnectionsDeletePoller.ResumeToken` has been removed
- Function `*PrivateLinkResourcesListByAccountPager.Err` has been removed
- Function `*AccountsCreateOrUpdatePoller.Done` has been removed
- Function `*PrivateEndpointConnectionsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*AccountsUpdatePoller.ResumeToken` has been removed
- Struct `AccountsAddRootCollectionAdminOptions` has been removed
- Struct `AccountsAddRootCollectionAdminResponse` has been removed
- Struct `AccountsBeginCreateOrUpdateOptions` has been removed
- Struct `AccountsBeginDeleteOptions` has been removed
- Struct `AccountsBeginUpdateOptions` has been removed
- Struct `AccountsCheckNameAvailabilityOptions` has been removed
- Struct `AccountsCheckNameAvailabilityResponse` has been removed
- Struct `AccountsCheckNameAvailabilityResult` has been removed
- Struct `AccountsCreateOrUpdatePoller` has been removed
- Struct `AccountsCreateOrUpdatePollerResponse` has been removed
- Struct `AccountsCreateOrUpdateResponse` has been removed
- Struct `AccountsCreateOrUpdateResult` has been removed
- Struct `AccountsDeletePoller` has been removed
- Struct `AccountsDeletePollerResponse` has been removed
- Struct `AccountsDeleteResponse` has been removed
- Struct `AccountsGetOptions` has been removed
- Struct `AccountsGetResponse` has been removed
- Struct `AccountsGetResult` has been removed
- Struct `AccountsListByResourceGroupOptions` has been removed
- Struct `AccountsListByResourceGroupPager` has been removed
- Struct `AccountsListByResourceGroupResponse` has been removed
- Struct `AccountsListByResourceGroupResult` has been removed
- Struct `AccountsListBySubscriptionOptions` has been removed
- Struct `AccountsListBySubscriptionPager` has been removed
- Struct `AccountsListBySubscriptionResponse` has been removed
- Struct `AccountsListBySubscriptionResult` has been removed
- Struct `AccountsListKeysOptions` has been removed
- Struct `AccountsListKeysResponse` has been removed
- Struct `AccountsListKeysResult` has been removed
- Struct `AccountsUpdatePoller` has been removed
- Struct `AccountsUpdatePollerResponse` has been removed
- Struct `AccountsUpdateResponse` has been removed
- Struct `AccountsUpdateResult` has been removed
- Struct `DefaultAccountsGetOptions` has been removed
- Struct `DefaultAccountsGetResponse` has been removed
- Struct `DefaultAccountsGetResult` has been removed
- Struct `DefaultAccountsRemoveOptions` has been removed
- Struct `DefaultAccountsRemoveResponse` has been removed
- Struct `DefaultAccountsSetOptions` has been removed
- Struct `DefaultAccountsSetResponse` has been removed
- Struct `DefaultAccountsSetResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `PrivateEndpointConnectionsBeginCreateOrUpdateOptions` has been removed
- Struct `PrivateEndpointConnectionsBeginDeleteOptions` has been removed
- Struct `PrivateEndpointConnectionsCreateOrUpdatePoller` has been removed
- Struct `PrivateEndpointConnectionsCreateOrUpdatePollerResponse` has been removed
- Struct `PrivateEndpointConnectionsCreateOrUpdateResponse` has been removed
- Struct `PrivateEndpointConnectionsCreateOrUpdateResult` has been removed
- Struct `PrivateEndpointConnectionsDeletePoller` has been removed
- Struct `PrivateEndpointConnectionsDeletePollerResponse` has been removed
- Struct `PrivateEndpointConnectionsDeleteResponse` has been removed
- Struct `PrivateEndpointConnectionsGetOptions` has been removed
- Struct `PrivateEndpointConnectionsGetResponse` has been removed
- Struct `PrivateEndpointConnectionsGetResult` has been removed
- Struct `PrivateEndpointConnectionsListByAccountOptions` has been removed
- Struct `PrivateEndpointConnectionsListByAccountPager` has been removed
- Struct `PrivateEndpointConnectionsListByAccountResponse` has been removed
- Struct `PrivateEndpointConnectionsListByAccountResult` has been removed
- Struct `PrivateLinkResourcesGetByGroupIDOptions` has been removed
- Struct `PrivateLinkResourcesGetByGroupIDResponse` has been removed
- Struct `PrivateLinkResourcesGetByGroupIDResult` has been removed
- Struct `PrivateLinkResourcesListByAccountOptions` has been removed
- Struct `PrivateLinkResourcesListByAccountPager` has been removed
- Struct `PrivateLinkResourcesListByAccountResponse` has been removed
- Struct `PrivateLinkResourcesListByAccountResult` has been removed
- Field `ProxyResource` of struct `PrivateEndpointConnection` has been removed
- Field `ErrorModel` of struct `ErrorResponseModelError` has been removed
- Field `SystemData` of struct `TrackedResourceSystemData` has been removed
- Field `AccountSKUAutoGenerated` of struct `AccountSKU` has been removed
- Field `InnerError` of struct `ErrorResponseModel` has been removed
- Field `TrackedResource` of struct `Account` has been removed
- Field `AccountEndpoints` of struct `AccountPropertiesEndpoints` has been removed
- Field `ManagedResources` of struct `AccountPropertiesManagedResources` has been removed

### Features Added

- New const `TypeNone`
- New const `TypeUserAssigned`
- New function `*AccountsClientListBySubscriptionPager.Err() error`
- New function `*AccountsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `AccountsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (AccountsClientCreateOrUpdateResponse, error)`
- New function `ErrorResponseModelError.MarshalJSON() ([]byte, error)`
- New function `TrackedResourceSystemData.MarshalJSON() ([]byte, error)`
- New function `*AccountsClientCreateOrUpdatePoller.Done() bool`
- New function `*AccountsClientDeletePollerResponse.Resume(context.Context, *AccountsClient, string) error`
- New function `*PrivateEndpointConnectionsClientListByAccountPager.NextPage(context.Context) bool`
- New function `*PrivateEndpointConnectionsClientListByAccountPager.PageResponse() PrivateEndpointConnectionsClientListByAccountResponse`
- New function `*PrivateEndpointConnectionsClientDeletePoller.FinalResponse(context.Context) (PrivateEndpointConnectionsClientDeleteResponse, error)`
- New function `*AccountsClientDeletePoller.ResumeToken() (string, error)`
- New function `*AccountsClientUpdatePollerResponse.Resume(context.Context, *AccountsClient, string) error`
- New function `*PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse.Resume(context.Context, *PrivateEndpointConnectionsClient, string) error`
- New function `*PrivateEndpointConnectionsClientDeletePoller.Done() bool`
- New function `*AccountsClientListByResourceGroupPager.PageResponse() AccountsClientListByResourceGroupResponse`
- New function `*AccountsClientCreateOrUpdatePoller.FinalResponse(context.Context) (AccountsClientCreateOrUpdateResponse, error)`
- New function `PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error)`
- New function `*PrivateLinkResourcesClientListByAccountPager.NextPage(context.Context) bool`
- New function `*AccountsClientUpdatePoller.Done() bool`
- New function `AccountsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (AccountsClientUpdateResponse, error)`
- New function `*AccountsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*AccountsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*PrivateLinkResourcesClientListByAccountPager.PageResponse() PrivateLinkResourcesClientListByAccountResponse`
- New function `*AccountsClientUpdatePoller.FinalResponse(context.Context) (AccountsClientUpdateResponse, error)`
- New function `AccountsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (AccountsClientDeleteResponse, error)`
- New function `Identity.MarshalJSON() ([]byte, error)`
- New function `*AccountsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*PrivateEndpointConnectionsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*PrivateEndpointConnectionsClientDeletePollerResponse.Resume(context.Context, *PrivateEndpointConnectionsClient, string) error`
- New function `*PrivateEndpointConnectionsClientListByAccountPager.Err() error`
- New function `*AccountsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*PrivateEndpointConnectionsClientCreateOrUpdatePoller.FinalResponse(context.Context) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error)`
- New function `*PrivateEndpointConnectionsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*TrackedResourceSystemData.UnmarshalJSON([]byte) error`
- New function `*AccountsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*AccountsClientListByResourceGroupPager.Err() error`
- New function `*AccountsClientDeletePoller.Done() bool`
- New function `*PrivateLinkResourcesClientListByAccountPager.Err() error`
- New function `*AccountsClientListBySubscriptionPager.PageResponse() AccountsClientListBySubscriptionResponse`
- New function `*PrivateEndpointConnectionsClientCreateOrUpdatePoller.Done() bool`
- New function `*AccountsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*OperationsClientListPager.Err() error`
- New function `*PrivateEndpointConnectionsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*PrivateEndpointConnectionsClientDeletePoller.ResumeToken() (string, error)`
- New function `PrivateEndpointConnectionsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (PrivateEndpointConnectionsClientDeleteResponse, error)`
- New function `*AccountsClientDeletePoller.FinalResponse(context.Context) (AccountsClientDeleteResponse, error)`
- New function `*AccountsClientCreateOrUpdatePollerResponse.Resume(context.Context, *AccountsClient, string) error`
- New struct `AccountsClientAddRootCollectionAdminOptions`
- New struct `AccountsClientAddRootCollectionAdminResponse`
- New struct `AccountsClientBeginCreateOrUpdateOptions`
- New struct `AccountsClientBeginDeleteOptions`
- New struct `AccountsClientBeginUpdateOptions`
- New struct `AccountsClientCheckNameAvailabilityOptions`
- New struct `AccountsClientCheckNameAvailabilityResponse`
- New struct `AccountsClientCheckNameAvailabilityResult`
- New struct `AccountsClientCreateOrUpdatePoller`
- New struct `AccountsClientCreateOrUpdatePollerResponse`
- New struct `AccountsClientCreateOrUpdateResponse`
- New struct `AccountsClientCreateOrUpdateResult`
- New struct `AccountsClientDeletePoller`
- New struct `AccountsClientDeletePollerResponse`
- New struct `AccountsClientDeleteResponse`
- New struct `AccountsClientGetOptions`
- New struct `AccountsClientGetResponse`
- New struct `AccountsClientGetResult`
- New struct `AccountsClientListByResourceGroupOptions`
- New struct `AccountsClientListByResourceGroupPager`
- New struct `AccountsClientListByResourceGroupResponse`
- New struct `AccountsClientListByResourceGroupResult`
- New struct `AccountsClientListBySubscriptionOptions`
- New struct `AccountsClientListBySubscriptionPager`
- New struct `AccountsClientListBySubscriptionResponse`
- New struct `AccountsClientListBySubscriptionResult`
- New struct `AccountsClientListKeysOptions`
- New struct `AccountsClientListKeysResponse`
- New struct `AccountsClientListKeysResult`
- New struct `AccountsClientUpdatePoller`
- New struct `AccountsClientUpdatePollerResponse`
- New struct `AccountsClientUpdateResponse`
- New struct `AccountsClientUpdateResult`
- New struct `DefaultAccountsClientGetOptions`
- New struct `DefaultAccountsClientGetResponse`
- New struct `DefaultAccountsClientGetResult`
- New struct `DefaultAccountsClientRemoveOptions`
- New struct `DefaultAccountsClientRemoveResponse`
- New struct `DefaultAccountsClientSetOptions`
- New struct `DefaultAccountsClientSetResponse`
- New struct `DefaultAccountsClientSetResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions`
- New struct `PrivateEndpointConnectionsClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdatePoller`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResult`
- New struct `PrivateEndpointConnectionsClientDeletePoller`
- New struct `PrivateEndpointConnectionsClientDeletePollerResponse`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateEndpointConnectionsClientGetOptions`
- New struct `PrivateEndpointConnectionsClientGetResponse`
- New struct `PrivateEndpointConnectionsClientGetResult`
- New struct `PrivateEndpointConnectionsClientListByAccountOptions`
- New struct `PrivateEndpointConnectionsClientListByAccountPager`
- New struct `PrivateEndpointConnectionsClientListByAccountResponse`
- New struct `PrivateEndpointConnectionsClientListByAccountResult`
- New struct `PrivateLinkResourcesClientGetByGroupIDOptions`
- New struct `PrivateLinkResourcesClientGetByGroupIDResponse`
- New struct `PrivateLinkResourcesClientGetByGroupIDResult`
- New struct `PrivateLinkResourcesClientListByAccountOptions`
- New struct `PrivateLinkResourcesClientListByAccountPager`
- New struct `PrivateLinkResourcesClientListByAccountResponse`
- New struct `PrivateLinkResourcesClientListByAccountResult`
- New struct `UserAssignedIdentity`
- New field `Details` in struct `ErrorResponseModelError`
- New field `Message` in struct `ErrorResponseModelError`
- New field `Target` in struct `ErrorResponseModelError`
- New field `Code` in struct `ErrorResponseModelError`
- New field `ID` in struct `PrivateEndpointConnection`
- New field `Name` in struct `PrivateEndpointConnection`
- New field `Type` in struct `PrivateEndpointConnection`
- New field `UserAssignedIdentities` in struct `Identity`
- New field `CreatedBy` in struct `TrackedResourceSystemData`
- New field `CreatedByType` in struct `TrackedResourceSystemData`
- New field `LastModifiedAt` in struct `TrackedResourceSystemData`
- New field `LastModifiedBy` in struct `TrackedResourceSystemData`
- New field `LastModifiedByType` in struct `TrackedResourceSystemData`
- New field `CreatedAt` in struct `TrackedResourceSystemData`
- New field `Catalog` in struct `AccountPropertiesEndpoints`
- New field `Guardian` in struct `AccountPropertiesEndpoints`
- New field `Scan` in struct `AccountPropertiesEndpoints`
- New field `Identity` in struct `AccountUpdateParameters`
- New field `Name` in struct `AccountSKU`
- New field `Capacity` in struct `AccountSKU`
- New field `Location` in struct `Account`
- New field `Tags` in struct `Account`
- New field `ID` in struct `Account`
- New field `Name` in struct `Account`
- New field `SystemData` in struct `Account`
- New field `Type` in struct `Account`
- New field `Identity` in struct `Account`
- New field `Error` in struct `ErrorResponseModel`
- New field `EventHubNamespace` in struct `AccountPropertiesManagedResources`
- New field `ResourceGroup` in struct `AccountPropertiesManagedResources`
- New field `StorageAccount` in struct `AccountPropertiesManagedResources`


## 0.1.0 (2021-12-16)

- Init release.

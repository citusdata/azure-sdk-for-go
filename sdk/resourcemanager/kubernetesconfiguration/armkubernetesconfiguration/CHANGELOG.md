# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*SourceControlConfigurationsClient.List` parameter(s) have been changed from `(string, Enum0, Enum1, string, *SourceControlConfigurationsListOptions)` to `(string, Enum0, Enum1, string, *SourceControlConfigurationsClientListOptions)`
- Function `*SourceControlConfigurationsClient.List` return value(s) have been changed from `(*SourceControlConfigurationsListPager)` to `(*SourceControlConfigurationsClientListPager)`
- Function `*ExtensionsClient.List` parameter(s) have been changed from `(string, Enum0, Enum1, string, *ExtensionsListOptions)` to `(string, Enum0, Enum1, string, *ExtensionsClientListOptions)`
- Function `*ExtensionsClient.List` return value(s) have been changed from `(*ExtensionsListPager)` to `(*ExtensionsClientListPager)`
- Function `*FluxConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *FluxConfigurationsGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *FluxConfigurationsClientGetOptions)`
- Function `*FluxConfigurationsClient.Get` return value(s) have been changed from `(FluxConfigurationsGetResponse, error)` to `(FluxConfigurationsClientGetResponse, error)`
- Function `*FluxConfigOperationStatusClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, string, *FluxConfigOperationStatusGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, string, *FluxConfigOperationStatusClientGetOptions)`
- Function `*FluxConfigOperationStatusClient.Get` return value(s) have been changed from `(FluxConfigOperationStatusGetResponse, error)` to `(FluxConfigOperationStatusClientGetResponse, error)`
- Function `*ExtensionsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, Extension, *ExtensionsBeginCreateOptions)` to `(context.Context, string, Enum0, Enum1, string, string, Extension, *ExtensionsClientBeginCreateOptions)`
- Function `*ExtensionsClient.BeginCreate` return value(s) have been changed from `(ExtensionsCreatePollerResponse, error)` to `(ExtensionsClientCreatePollerResponse, error)`
- Function `*OperationStatusClient.List` parameter(s) have been changed from `(string, Enum0, Enum1, string, *OperationStatusListOptions)` to `(string, Enum0, Enum1, string, *OperationStatusClientListOptions)`
- Function `*OperationStatusClient.List` return value(s) have been changed from `(*OperationStatusListPager)` to `(*OperationStatusClientListPager)`
- Function `*FluxConfigurationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *FluxConfigurationsBeginDeleteOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *FluxConfigurationsClientBeginDeleteOptions)`
- Function `*FluxConfigurationsClient.BeginDelete` return value(s) have been changed from `(FluxConfigurationsDeletePollerResponse, error)` to `(FluxConfigurationsClientDeletePollerResponse, error)`
- Function `*ClusterExtensionTypeClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *ClusterExtensionTypeGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *ClusterExtensionTypeClientGetOptions)`
- Function `*ClusterExtensionTypeClient.Get` return value(s) have been changed from `(ClusterExtensionTypeGetResponse, error)` to `(ClusterExtensionTypeClientGetResponse, error)`
- Function `*OperationStatusClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, string, *OperationStatusGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, string, *OperationStatusClientGetOptions)`
- Function `*OperationStatusClient.Get` return value(s) have been changed from `(OperationStatusGetResponse, error)` to `(OperationStatusClientGetResponse, error)`
- Function `*ExtensionTypeVersionsClient.List` parameter(s) have been changed from `(string, string, *ExtensionTypeVersionsListOptions)` to `(string, string, *ExtensionTypeVersionsClientListOptions)`
- Function `*ExtensionTypeVersionsClient.List` return value(s) have been changed from `(*ExtensionTypeVersionsListPager)` to `(*ExtensionTypeVersionsClientListPager)`
- Function `*ExtensionsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, PatchExtension, *ExtensionsBeginUpdateOptions)` to `(context.Context, string, Enum0, Enum1, string, string, PatchExtension, *ExtensionsClientBeginUpdateOptions)`
- Function `*ExtensionsClient.BeginUpdate` return value(s) have been changed from `(ExtensionsUpdatePollerResponse, error)` to `(ExtensionsClientUpdatePollerResponse, error)`
- Function `*FluxConfigurationsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, FluxConfiguration, *FluxConfigurationsBeginCreateOrUpdateOptions)` to `(context.Context, string, Enum0, Enum1, string, string, FluxConfiguration, *FluxConfigurationsClientBeginCreateOrUpdateOptions)`
- Function `*FluxConfigurationsClient.BeginCreateOrUpdate` return value(s) have been changed from `(FluxConfigurationsCreateOrUpdatePollerResponse, error)` to `(FluxConfigurationsClientCreateOrUpdatePollerResponse, error)`
- Function `*SourceControlConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *SourceControlConfigurationsGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *SourceControlConfigurationsClientGetOptions)`
- Function `*SourceControlConfigurationsClient.Get` return value(s) have been changed from `(SourceControlConfigurationsGetResponse, error)` to `(SourceControlConfigurationsClientGetResponse, error)`
- Function `*LocationExtensionTypesClient.List` parameter(s) have been changed from `(string, *LocationExtensionTypesListOptions)` to `(string, *LocationExtensionTypesClientListOptions)`
- Function `*LocationExtensionTypesClient.List` return value(s) have been changed from `(*LocationExtensionTypesListPager)` to `(*LocationExtensionTypesClientListPager)`
- Function `*SourceControlConfigurationsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, SourceControlConfiguration, *SourceControlConfigurationsCreateOrUpdateOptions)` to `(context.Context, string, Enum0, Enum1, string, string, SourceControlConfiguration, *SourceControlConfigurationsClientCreateOrUpdateOptions)`
- Function `*SourceControlConfigurationsClient.CreateOrUpdate` return value(s) have been changed from `(SourceControlConfigurationsCreateOrUpdateResponse, error)` to `(SourceControlConfigurationsClientCreateOrUpdateResponse, error)`
- Function `*ExtensionsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *ExtensionsBeginDeleteOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *ExtensionsClientBeginDeleteOptions)`
- Function `*ExtensionsClient.BeginDelete` return value(s) have been changed from `(ExtensionsDeletePollerResponse, error)` to `(ExtensionsClientDeletePollerResponse, error)`
- Function `*SourceControlConfigurationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *SourceControlConfigurationsBeginDeleteOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *SourceControlConfigurationsClientBeginDeleteOptions)`
- Function `*SourceControlConfigurationsClient.BeginDelete` return value(s) have been changed from `(SourceControlConfigurationsDeletePollerResponse, error)` to `(SourceControlConfigurationsClientDeletePollerResponse, error)`
- Function `*ExtensionsClient.Get` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, *ExtensionsGetOptions)` to `(context.Context, string, Enum0, Enum1, string, string, *ExtensionsClientGetOptions)`
- Function `*ExtensionsClient.Get` return value(s) have been changed from `(ExtensionsGetResponse, error)` to `(ExtensionsClientGetResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*ClusterExtensionTypesClient.List` parameter(s) have been changed from `(string, Enum0, Enum1, string, *ClusterExtensionTypesListOptions)` to `(string, Enum0, Enum1, string, *ClusterExtensionTypesClientListOptions)`
- Function `*ClusterExtensionTypesClient.List` return value(s) have been changed from `(*ClusterExtensionTypesListPager)` to `(*ClusterExtensionTypesClientListPager)`
- Function `*FluxConfigurationsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, Enum0, Enum1, string, string, FluxConfigurationPatch, *FluxConfigurationsBeginUpdateOptions)` to `(context.Context, string, Enum0, Enum1, string, string, FluxConfigurationPatch, *FluxConfigurationsClientBeginUpdateOptions)`
- Function `*FluxConfigurationsClient.BeginUpdate` return value(s) have been changed from `(FluxConfigurationsUpdatePollerResponse, error)` to `(FluxConfigurationsClientUpdatePollerResponse, error)`
- Function `*FluxConfigurationsClient.List` parameter(s) have been changed from `(string, Enum0, Enum1, string, *FluxConfigurationsListOptions)` to `(string, Enum0, Enum1, string, *FluxConfigurationsClientListOptions)`
- Function `*FluxConfigurationsClient.List` return value(s) have been changed from `(*FluxConfigurationsListPager)` to `(*FluxConfigurationsClientListPager)`
- Function `FluxConfigurationsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ExtensionsDeletePoller.Done` has been removed
- Function `FluxConfigurationsDeletePollerResponse.PollUntilDone` has been removed
- Function `*ExtensionsDeletePoller.ResumeToken` has been removed
- Function `*FluxConfigurationsUpdatePoller.Done` has been removed
- Function `*ExtensionsUpdatePollerResponse.Resume` has been removed
- Function `SourceControlConfigurationsDeletePollerResponse.PollUntilDone` has been removed
- Function `*ExtensionsListPager.PageResponse` has been removed
- Function `*FluxConfigurationsDeletePollerResponse.Resume` has been removed
- Function `*ExtensionsUpdatePoller.FinalResponse` has been removed
- Function `*FluxConfigurationsUpdatePoller.Poll` has been removed
- Function `*ExtensionsUpdatePoller.ResumeToken` has been removed
- Function `*LocationExtensionTypesListPager.NextPage` has been removed
- Function `*ExtensionsDeletePoller.FinalResponse` has been removed
- Function `*SourceControlConfigurationsDeletePoller.Done` has been removed
- Function `*SourceControlConfigurationsDeletePollerResponse.Resume` has been removed
- Function `*ExtensionsDeletePollerResponse.Resume` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `*ExtensionTypeVersionsListPager.PageResponse` has been removed
- Function `*FluxConfigurationsListPager.Err` has been removed
- Function `*ExtensionsUpdatePoller.Done` has been removed
- Function `*FluxConfigurationsUpdatePoller.ResumeToken` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*FluxConfigurationsCreateOrUpdatePoller.Poll` has been removed
- Function `*FluxConfigurationsListPager.PageResponse` has been removed
- Function `*SourceControlConfigurationsListPager.NextPage` has been removed
- Function `*ExtensionsListPager.NextPage` has been removed
- Function `*OperationStatusListPager.NextPage` has been removed
- Function `*SourceControlConfigurationsListPager.PageResponse` has been removed
- Function `*FluxConfigurationsDeletePoller.FinalResponse` has been removed
- Function `*FluxConfigurationsDeletePoller.ResumeToken` has been removed
- Function `ExtensionsCreatePollerResponse.PollUntilDone` has been removed
- Function `*ExtensionsListPager.Err` has been removed
- Function `ExtensionsDeletePollerResponse.PollUntilDone` has been removed
- Function `*ExtensionsCreatePollerResponse.Resume` has been removed
- Function `*ExtensionsCreatePoller.FinalResponse` has been removed
- Function `*SourceControlConfigurationsListPager.Err` has been removed
- Function `*FluxConfigurationsUpdatePoller.FinalResponse` has been removed
- Function `*SourceControlConfigurationsDeletePoller.Poll` has been removed
- Function `*LocationExtensionTypesListPager.PageResponse` has been removed
- Function `*OperationStatusListPager.Err` has been removed
- Function `*ExtensionsCreatePoller.Done` has been removed
- Function `*FluxConfigurationsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*FluxConfigurationsCreateOrUpdatePoller.Done` has been removed
- Function `*ClusterExtensionTypesListPager.NextPage` has been removed
- Function `*ClusterExtensionTypesListPager.PageResponse` has been removed
- Function `*ExtensionsCreatePoller.ResumeToken` has been removed
- Function `*ExtensionTypeVersionsListPager.NextPage` has been removed
- Function `*FluxConfigurationsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*FluxConfigurationsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*FluxConfigurationsUpdatePollerResponse.Resume` has been removed
- Function `*FluxConfigurationsDeletePoller.Poll` has been removed
- Function `*LocationExtensionTypesListPager.Err` has been removed
- Function `*ClusterExtensionTypesListPager.Err` has been removed
- Function `*ExtensionTypeVersionsListPager.Err` has been removed
- Function `*ExtensionsCreatePoller.Poll` has been removed
- Function `*FluxConfigurationsDeletePoller.Done` has been removed
- Function `*ExtensionsDeletePoller.Poll` has been removed
- Function `*SourceControlConfigurationsDeletePoller.FinalResponse` has been removed
- Function `*OperationStatusListPager.PageResponse` has been removed
- Function `ExtensionsUpdatePollerResponse.PollUntilDone` has been removed
- Function `FluxConfigurationsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*FluxConfigurationsListPager.NextPage` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `*SourceControlConfigurationsDeletePoller.ResumeToken` has been removed
- Function `*ExtensionsUpdatePoller.Poll` has been removed
- Struct `ClusterExtensionTypeGetOptions` has been removed
- Struct `ClusterExtensionTypeGetResponse` has been removed
- Struct `ClusterExtensionTypeGetResult` has been removed
- Struct `ClusterExtensionTypesListOptions` has been removed
- Struct `ClusterExtensionTypesListPager` has been removed
- Struct `ClusterExtensionTypesListResponse` has been removed
- Struct `ClusterExtensionTypesListResult` has been removed
- Struct `ExtensionTypeVersionsListOptions` has been removed
- Struct `ExtensionTypeVersionsListPager` has been removed
- Struct `ExtensionTypeVersionsListResponse` has been removed
- Struct `ExtensionTypeVersionsListResult` has been removed
- Struct `ExtensionsBeginCreateOptions` has been removed
- Struct `ExtensionsBeginDeleteOptions` has been removed
- Struct `ExtensionsBeginUpdateOptions` has been removed
- Struct `ExtensionsCreatePoller` has been removed
- Struct `ExtensionsCreatePollerResponse` has been removed
- Struct `ExtensionsCreateResponse` has been removed
- Struct `ExtensionsCreateResult` has been removed
- Struct `ExtensionsDeletePoller` has been removed
- Struct `ExtensionsDeletePollerResponse` has been removed
- Struct `ExtensionsDeleteResponse` has been removed
- Struct `ExtensionsGetOptions` has been removed
- Struct `ExtensionsGetResponse` has been removed
- Struct `ExtensionsGetResult` has been removed
- Struct `ExtensionsListOptions` has been removed
- Struct `ExtensionsListPager` has been removed
- Struct `ExtensionsListResponse` has been removed
- Struct `ExtensionsListResult` has been removed
- Struct `ExtensionsUpdatePoller` has been removed
- Struct `ExtensionsUpdatePollerResponse` has been removed
- Struct `ExtensionsUpdateResponse` has been removed
- Struct `ExtensionsUpdateResult` has been removed
- Struct `FluxConfigOperationStatusGetOptions` has been removed
- Struct `FluxConfigOperationStatusGetResponse` has been removed
- Struct `FluxConfigOperationStatusGetResult` has been removed
- Struct `FluxConfigurationsBeginCreateOrUpdateOptions` has been removed
- Struct `FluxConfigurationsBeginDeleteOptions` has been removed
- Struct `FluxConfigurationsBeginUpdateOptions` has been removed
- Struct `FluxConfigurationsCreateOrUpdatePoller` has been removed
- Struct `FluxConfigurationsCreateOrUpdatePollerResponse` has been removed
- Struct `FluxConfigurationsCreateOrUpdateResponse` has been removed
- Struct `FluxConfigurationsCreateOrUpdateResult` has been removed
- Struct `FluxConfigurationsDeletePoller` has been removed
- Struct `FluxConfigurationsDeletePollerResponse` has been removed
- Struct `FluxConfigurationsDeleteResponse` has been removed
- Struct `FluxConfigurationsGetOptions` has been removed
- Struct `FluxConfigurationsGetResponse` has been removed
- Struct `FluxConfigurationsGetResult` has been removed
- Struct `FluxConfigurationsListOptions` has been removed
- Struct `FluxConfigurationsListPager` has been removed
- Struct `FluxConfigurationsListResponse` has been removed
- Struct `FluxConfigurationsListResult` has been removed
- Struct `FluxConfigurationsUpdatePoller` has been removed
- Struct `FluxConfigurationsUpdatePollerResponse` has been removed
- Struct `FluxConfigurationsUpdateResponse` has been removed
- Struct `FluxConfigurationsUpdateResult` has been removed
- Struct `LocationExtensionTypesListOptions` has been removed
- Struct `LocationExtensionTypesListPager` has been removed
- Struct `LocationExtensionTypesListResponse` has been removed
- Struct `LocationExtensionTypesListResult` has been removed
- Struct `OperationStatusGetOptions` has been removed
- Struct `OperationStatusGetResponse` has been removed
- Struct `OperationStatusGetResult` has been removed
- Struct `OperationStatusListOptions` has been removed
- Struct `OperationStatusListPager` has been removed
- Struct `OperationStatusListResponse` has been removed
- Struct `OperationStatusListResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `SourceControlConfigurationsBeginDeleteOptions` has been removed
- Struct `SourceControlConfigurationsCreateOrUpdateOptions` has been removed
- Struct `SourceControlConfigurationsCreateOrUpdateResponse` has been removed
- Struct `SourceControlConfigurationsCreateOrUpdateResult` has been removed
- Struct `SourceControlConfigurationsDeletePoller` has been removed
- Struct `SourceControlConfigurationsDeletePollerResponse` has been removed
- Struct `SourceControlConfigurationsDeleteResponse` has been removed
- Struct `SourceControlConfigurationsGetOptions` has been removed
- Struct `SourceControlConfigurationsGetResponse` has been removed
- Struct `SourceControlConfigurationsGetResult` has been removed
- Struct `SourceControlConfigurationsListOptions` has been removed
- Struct `SourceControlConfigurationsListPager` has been removed
- Struct `SourceControlConfigurationsListResponse` has been removed
- Struct `SourceControlConfigurationsListResult` has been removed
- Field `ProxyResource` of struct `ClusterScopeSettings` has been removed
- Field `ProxyResource` of struct `FluxConfiguration` has been removed
- Field `Resource` of struct `ProxyResource` has been removed
- Field `ProxyResource` of struct `SourceControlConfiguration` has been removed
- Field `ProxyResource` of struct `Extension` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed

### Features Added

- New function `FluxConfigurationsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (FluxConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*SourceControlConfigurationsClientDeletePoller.Done() bool`
- New function `*FluxConfigurationsClientListPager.NextPage(context.Context) bool`
- New function `*FluxConfigurationsClientCreateOrUpdatePoller.Done() bool`
- New function `FluxConfigurationsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (FluxConfigurationsClientUpdateResponse, error)`
- New function `*ExtensionsClientUpdatePoller.Done() bool`
- New function `*SourceControlConfigurationsClientDeletePollerResponse.Resume(context.Context, *SourceControlConfigurationsClient, string) error`
- New function `*FluxConfigurationsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FluxConfigurationsClientDeletePoller.Done() bool`
- New function `*ExtensionsClientListPager.Err() error`
- New function `*ExtensionsClientDeletePoller.ResumeToken() (string, error)`
- New function `*ExtensionsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*FluxConfigurationsClientUpdatePoller.ResumeToken() (string, error)`
- New function `ExtensionsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (ExtensionsClientCreateResponse, error)`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*LocationExtensionTypesClientListPager.NextPage(context.Context) bool`
- New function `*FluxConfigurationsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ExtensionsClientCreatePollerResponse.Resume(context.Context, *ExtensionsClient, string) error`
- New function `*ExtensionsClientUpdatePoller.FinalResponse(context.Context) (ExtensionsClientUpdateResponse, error)`
- New function `*ExtensionsClientListPager.PageResponse() ExtensionsClientListResponse`
- New function `*FluxConfigurationsClientUpdatePoller.Done() bool`
- New function `*SourceControlConfigurationsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*SourceControlConfigurationsClientListPager.Err() error`
- New function `*SourceControlConfigurationsClientDeletePoller.FinalResponse(context.Context) (SourceControlConfigurationsClientDeleteResponse, error)`
- New function `SourceControlConfigurationsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (SourceControlConfigurationsClientDeleteResponse, error)`
- New function `*FluxConfigurationsClientCreateOrUpdatePollerResponse.Resume(context.Context, *FluxConfigurationsClient, string) error`
- New function `*ClusterExtensionTypesClientListPager.PageResponse() ClusterExtensionTypesClientListResponse`
- New function `*FluxConfigurationsClientDeletePollerResponse.Resume(context.Context, *FluxConfigurationsClient, string) error`
- New function `*LocationExtensionTypesClientListPager.PageResponse() LocationExtensionTypesClientListResponse`
- New function `*ClusterExtensionTypesClientListPager.Err() error`
- New function `*ExtensionsClientCreatePoller.Done() bool`
- New function `*FluxConfigurationsClientListPager.Err() error`
- New function `*ClusterExtensionTypesClientListPager.NextPage(context.Context) bool`
- New function `*OperationStatusClientListPager.Err() error`
- New function `*ExtensionsClientDeletePollerResponse.Resume(context.Context, *ExtensionsClient, string) error`
- New function `*FluxConfigurationsClientDeletePoller.FinalResponse(context.Context) (FluxConfigurationsClientDeleteResponse, error)`
- New function `*ExtensionsClientCreatePoller.FinalResponse(context.Context) (ExtensionsClientCreateResponse, error)`
- New function `ExtensionsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (ExtensionsClientDeleteResponse, error)`
- New function `*ExtensionsClientDeletePoller.FinalResponse(context.Context) (ExtensionsClientDeleteResponse, error)`
- New function `*OperationStatusClientListPager.NextPage(context.Context) bool`
- New function `*ExtensionTypeVersionsClientListPager.NextPage(context.Context) bool`
- New function `ExtensionsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ExtensionsClientUpdateResponse, error)`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*SourceControlConfigurationsClientListPager.NextPage(context.Context) bool`
- New function `*ExtensionsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FluxConfigurationsClientListPager.PageResponse() FluxConfigurationsClientListResponse`
- New function `*ExtensionsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ExtensionsClientUpdatePollerResponse.Resume(context.Context, *ExtensionsClient, string) error`
- New function `*FluxConfigurationsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*OperationsClientListPager.Err() error`
- New function `*OperationStatusClientListPager.PageResponse() OperationStatusClientListResponse`
- New function `*FluxConfigurationsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FluxConfigurationsClientCreateOrUpdatePoller.FinalResponse(context.Context) (FluxConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*FluxConfigurationsClientDeletePoller.ResumeToken() (string, error)`
- New function `FluxConfigurationsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (FluxConfigurationsClientDeleteResponse, error)`
- New function `*ExtensionsClientDeletePoller.Done() bool`
- New function `*LocationExtensionTypesClientListPager.Err() error`
- New function `*FluxConfigurationsClientUpdatePoller.FinalResponse(context.Context) (FluxConfigurationsClientUpdateResponse, error)`
- New function `*ExtensionsClientCreatePoller.ResumeToken() (string, error)`
- New function `*SourceControlConfigurationsClientListPager.PageResponse() SourceControlConfigurationsClientListResponse`
- New function `*ExtensionTypeVersionsClientListPager.Err() error`
- New function `*ExtensionsClientListPager.NextPage(context.Context) bool`
- New function `*ExtensionTypeVersionsClientListPager.PageResponse() ExtensionTypeVersionsClientListResponse`
- New function `*ExtensionsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FluxConfigurationsClientUpdatePollerResponse.Resume(context.Context, *FluxConfigurationsClient, string) error`
- New function `*SourceControlConfigurationsClientDeletePoller.ResumeToken() (string, error)`
- New struct `ClusterExtensionTypeClientGetOptions`
- New struct `ClusterExtensionTypeClientGetResponse`
- New struct `ClusterExtensionTypeClientGetResult`
- New struct `ClusterExtensionTypesClientListOptions`
- New struct `ClusterExtensionTypesClientListPager`
- New struct `ClusterExtensionTypesClientListResponse`
- New struct `ClusterExtensionTypesClientListResult`
- New struct `ExtensionTypeVersionsClientListOptions`
- New struct `ExtensionTypeVersionsClientListPager`
- New struct `ExtensionTypeVersionsClientListResponse`
- New struct `ExtensionTypeVersionsClientListResult`
- New struct `ExtensionsClientBeginCreateOptions`
- New struct `ExtensionsClientBeginDeleteOptions`
- New struct `ExtensionsClientBeginUpdateOptions`
- New struct `ExtensionsClientCreatePoller`
- New struct `ExtensionsClientCreatePollerResponse`
- New struct `ExtensionsClientCreateResponse`
- New struct `ExtensionsClientCreateResult`
- New struct `ExtensionsClientDeletePoller`
- New struct `ExtensionsClientDeletePollerResponse`
- New struct `ExtensionsClientDeleteResponse`
- New struct `ExtensionsClientGetOptions`
- New struct `ExtensionsClientGetResponse`
- New struct `ExtensionsClientGetResult`
- New struct `ExtensionsClientListOptions`
- New struct `ExtensionsClientListPager`
- New struct `ExtensionsClientListResponse`
- New struct `ExtensionsClientListResult`
- New struct `ExtensionsClientUpdatePoller`
- New struct `ExtensionsClientUpdatePollerResponse`
- New struct `ExtensionsClientUpdateResponse`
- New struct `ExtensionsClientUpdateResult`
- New struct `FluxConfigOperationStatusClientGetOptions`
- New struct `FluxConfigOperationStatusClientGetResponse`
- New struct `FluxConfigOperationStatusClientGetResult`
- New struct `FluxConfigurationsClientBeginCreateOrUpdateOptions`
- New struct `FluxConfigurationsClientBeginDeleteOptions`
- New struct `FluxConfigurationsClientBeginUpdateOptions`
- New struct `FluxConfigurationsClientCreateOrUpdatePoller`
- New struct `FluxConfigurationsClientCreateOrUpdatePollerResponse`
- New struct `FluxConfigurationsClientCreateOrUpdateResponse`
- New struct `FluxConfigurationsClientCreateOrUpdateResult`
- New struct `FluxConfigurationsClientDeletePoller`
- New struct `FluxConfigurationsClientDeletePollerResponse`
- New struct `FluxConfigurationsClientDeleteResponse`
- New struct `FluxConfigurationsClientGetOptions`
- New struct `FluxConfigurationsClientGetResponse`
- New struct `FluxConfigurationsClientGetResult`
- New struct `FluxConfigurationsClientListOptions`
- New struct `FluxConfigurationsClientListPager`
- New struct `FluxConfigurationsClientListResponse`
- New struct `FluxConfigurationsClientListResult`
- New struct `FluxConfigurationsClientUpdatePoller`
- New struct `FluxConfigurationsClientUpdatePollerResponse`
- New struct `FluxConfigurationsClientUpdateResponse`
- New struct `FluxConfigurationsClientUpdateResult`
- New struct `LocationExtensionTypesClientListOptions`
- New struct `LocationExtensionTypesClientListPager`
- New struct `LocationExtensionTypesClientListResponse`
- New struct `LocationExtensionTypesClientListResult`
- New struct `OperationStatusClientGetOptions`
- New struct `OperationStatusClientGetResponse`
- New struct `OperationStatusClientGetResult`
- New struct `OperationStatusClientListOptions`
- New struct `OperationStatusClientListPager`
- New struct `OperationStatusClientListResponse`
- New struct `OperationStatusClientListResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `SourceControlConfigurationsClientBeginDeleteOptions`
- New struct `SourceControlConfigurationsClientCreateOrUpdateOptions`
- New struct `SourceControlConfigurationsClientCreateOrUpdateResponse`
- New struct `SourceControlConfigurationsClientCreateOrUpdateResult`
- New struct `SourceControlConfigurationsClientDeletePoller`
- New struct `SourceControlConfigurationsClientDeletePollerResponse`
- New struct `SourceControlConfigurationsClientDeleteResponse`
- New struct `SourceControlConfigurationsClientGetOptions`
- New struct `SourceControlConfigurationsClientGetResponse`
- New struct `SourceControlConfigurationsClientGetResult`
- New struct `SourceControlConfigurationsClientListOptions`
- New struct `SourceControlConfigurationsClientListPager`
- New struct `SourceControlConfigurationsClientListResponse`
- New struct `SourceControlConfigurationsClientListResult`
- New field `Error` in struct `ErrorResponse`
- New field `ID` in struct `ClusterScopeSettings`
- New field `Name` in struct `ClusterScopeSettings`
- New field `Type` in struct `ClusterScopeSettings`
- New field `ID` in struct `FluxConfiguration`
- New field `Name` in struct `FluxConfiguration`
- New field `Type` in struct `FluxConfiguration`
- New field `ID` in struct `ProxyResource`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `ID` in struct `SourceControlConfiguration`
- New field `Name` in struct `SourceControlConfiguration`
- New field `Type` in struct `SourceControlConfiguration`
- New field `Name` in struct `Extension`
- New field `Type` in struct `Extension`
- New field `ID` in struct `Extension`


## 0.1.0 (2021-12-07)

- Init release.
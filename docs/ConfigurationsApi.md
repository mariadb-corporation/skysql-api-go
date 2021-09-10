# \ConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ConfigurationsApi.md#CreateConfiguration) | **Post** /configurations/ | Create Configuration
[**DeleteConfiguration**](ConfigurationsApi.md#DeleteConfiguration) | **Delete** /configurations/{configuration_number} | Delete Configuration
[**ListConfigurations**](ConfigurationsApi.md#ListConfigurations) | **Get** /configurations/ | List Configurations
[**ReadConfiguration**](ConfigurationsApi.md#ReadConfiguration) | **Get** /configurations/{configuration_number} | Read Configuration
[**UpdateConfiguration**](ConfigurationsApi.md#UpdateConfiguration) | **Patch** /configurations/{configuration_number} | Update Configuration



## CreateConfiguration

> NewConfigurationResponse CreateConfiguration(ctx).NewConfigurationRequest(newConfigurationRequest).Execute()

Create Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    newConfigurationRequest := *openapiclient.NewNewConfigurationRequest("Name_example", "Topology_example") // NewConfigurationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationsApi.CreateConfiguration(context.Background()).NewConfigurationRequest(newConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.CreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfiguration`: NewConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newConfigurationRequest** | [**NewConfigurationRequest**](NewConfigurationRequest.md) |  | 

### Return type

[**NewConfigurationResponse**](NewConfigurationResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> interface{} DeleteConfiguration(ctx, configurationNumber).Execute()

Delete Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationNumber := "CFG0000000" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationsApi.DeleteConfiguration(context.Background(), configurationNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfiguration`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.DeleteConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations

> []Configuration ListConfigurations(ctx).Limit(limit).Execute()

List Configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationsApi.ListConfigurations(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.ListConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigurations`: []Configuration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.ListConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 10]

### Return type

[**[]Configuration**](Configuration.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConfiguration

> ConfigurationResponse ReadConfiguration(ctx, configurationNumber).Execute()

Read Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationNumber := "CFG0000000" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationsApi.ReadConfiguration(context.Background(), configurationNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.ReadConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadConfiguration`: ConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.ReadConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ConfigurationResponse UpdateConfiguration(ctx, configurationNumber).ConfigurationUpdate(configurationUpdate).Execute()

Update Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationNumber := "CFG0000000" // string | 
    configurationUpdate := *openapiclient.NewConfigurationUpdate() // ConfigurationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationsApi.UpdateConfiguration(context.Background(), configurationNumber).ConfigurationUpdate(configurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: ConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationUpdate** | [**ConfigurationUpdate**](ConfigurationUpdate.md) |  | 

### Return type

[**ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


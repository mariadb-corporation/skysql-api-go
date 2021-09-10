# \DatabasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesApi.md#CreateDatabase) | **Post** /databases/ | Create Database
[**DeleteDatabase**](DatabasesApi.md#DeleteDatabase) | **Delete** /databases/{database_id} | Delete Database
[**ListDatabases**](DatabasesApi.md#ListDatabases) | **Get** /databases/ | List Databases
[**ReadDatabase**](DatabasesApi.md#ReadDatabase) | **Get** /databases/{database_id} | Read Database
[**UpdateDatabase**](DatabasesApi.md#UpdateDatabase) | **Patch** /databases/{database_id} | Update Database



## CreateDatabase

> Database CreateDatabase(ctx).NewDatabase(newDatabase).Execute()

Create Database



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
    newDatabase := *openapiclient.NewNewDatabase("ReleaseVersion_example", "Topology_example", "Size_example", "TxStorage_example", "MaxscaleConfig_example", "Name_example", "Region_example", "ReplRegion_example", "Provider_example", "Replicas_example", "Monitor_example", "MaxscaleProxy_example") // NewDatabase |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabasesApi.CreateDatabase(context.Background()).NewDatabase(newDatabase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDatabase** | [**NewDatabase**](NewDatabase.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> interface{} DeleteDatabase(ctx, databaseId).Execute()

Delete Database



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
    databaseId := "db00000000" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabasesApi.DeleteDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabase`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


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


## ListDatabases

> []Database ListDatabases(ctx).Limit(limit).Execute()

List Databases



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
    resp, r, err := api_client.DatabasesApi.ListDatabases(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: []Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 10]

### Return type

[**[]Database**](Database.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDatabase

> Database ReadDatabase(ctx, databaseId).Execute()

Read Database



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
    databaseId := "db00000000" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabasesApi.ReadDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.ReadDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.ReadDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Database**](Database.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabase

> Database UpdateDatabase(ctx, databaseId).DatabaseUpdate(databaseUpdate).Execute()

Update Database



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
    databaseId := "db00000000" // string | 
    databaseUpdate := *openapiclient.NewDatabaseUpdate("Name_example") // DatabaseUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabasesApi.UpdateDatabase(context.Background(), databaseId).DatabaseUpdate(databaseUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.UpdateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.UpdateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **databaseUpdate** | [**DatabaseUpdate**](DatabaseUpdate.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


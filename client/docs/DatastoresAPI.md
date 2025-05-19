# \DatastoresAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatastores**](DatastoresAPI.md#GetDatastores) | **Get** /api/data-stores/{id} | Retrieves a Datastore
[**ListDatastores**](DatastoresAPI.md#ListDatastores) | **Get** /api/data-stores | Retrieves all Datastores
[**SaveDatastore**](DatastoresAPI.md#SaveDatastore) | **Post** /api/data-stores | Create a Datastore
[**UpdateDatastores**](DatastoresAPI.md#UpdateDatastores) | **Put** /api/data-stores/{id} | Updates a Specified Datastore



## GetDatastores

> GetDatastores200Response GetDatastores(ctx, id).Execute()

Retrieves a Datastore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoresAPI.GetDatastores(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoresAPI.GetDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatastores`: GetDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `DatastoresAPI.GetDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDatastores200Response**](GetDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatastores

> ListCloudDatastores200Response ListDatastores(ctx).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()

Retrieves all Datastores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoresAPI.ListDatastores(context.Background()).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoresAPI.ListDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatastores`: ListCloudDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `DatastoresAPI.ListDatastores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]

### Return type

[**ListCloudDatastores200Response**](ListCloudDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveDatastore

> SaveCloudDatastore200Response SaveDatastore(ctx).SaveDatastoreRequest(saveDatastoreRequest).Execute()

Create a Datastore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	saveDatastoreRequest := *openapiclient.NewSaveDatastoreRequest() // SaveDatastoreRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoresAPI.SaveDatastore(context.Background()).SaveDatastoreRequest(saveDatastoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoresAPI.SaveDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveDatastore`: SaveCloudDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `DatastoresAPI.SaveDatastore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveDatastoreRequest** | [**SaveDatastoreRequest**](SaveDatastoreRequest.md) |  | 

### Return type

[**SaveCloudDatastore200Response**](SaveCloudDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatastores

> UpdateDatastores200Response UpdateDatastores(ctx, id).UpdateCloudDatastoresRequest(updateCloudDatastoresRequest).Execute()

Updates a Specified Datastore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudDatastoresRequest := *openapiclient.NewUpdateCloudDatastoresRequest(*openapiclient.NewUpdateCloudDatastoresRequestDatastore()) // UpdateCloudDatastoresRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatastoresAPI.UpdateDatastores(context.Background(), id).UpdateCloudDatastoresRequest(updateCloudDatastoresRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatastoresAPI.UpdateDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDatastores`: UpdateDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `DatastoresAPI.UpdateDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCloudDatastoresRequest** | [**UpdateCloudDatastoresRequest**](UpdateCloudDatastoresRequest.md) |  | 

### Return type

[**UpdateDatastores200Response**](UpdateDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \WikiAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWiki**](WikiAPI.md#AddWiki) | **Post** /api/wiki/pages | Create a Wiki Page
[**GetWiki**](WikiAPI.md#GetWiki) | **Get** /api/wiki/pages/{id} | Retrieves a specific Wiki page
[**GetWikiApp**](WikiAPI.md#GetWikiApp) | **Get** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**GetWikiCategories**](WikiAPI.md#GetWikiCategories) | **Get** /api/wiki/categories | Retrieves all Wiki categories associated with the account
[**GetWikiCloud**](WikiAPI.md#GetWikiCloud) | **Get** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**GetWikiCluster**](WikiAPI.md#GetWikiCluster) | **Get** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**GetWikiGroup**](WikiAPI.md#GetWikiGroup) | **Get** /api/groups/{id}/wiki | Retrieves a Group Wiki Page
[**GetWikiInstance**](WikiAPI.md#GetWikiInstance) | **Get** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**GetWikiServer**](WikiAPI.md#GetWikiServer) | **Get** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**ListWiki**](WikiAPI.md#ListWiki) | **Get** /api/wiki/pages | Retrieves wiki pages associated with the account.
[**RemoveWiki**](WikiAPI.md#RemoveWiki) | **Delete** /api/wiki/pages/{id} | Deletes a Wiki Page
[**UpdateWiki**](WikiAPI.md#UpdateWiki) | **Put** /api/wiki/pages/{id} | Updates a Wiki Page
[**UpdateWikiApp**](WikiAPI.md#UpdateWikiApp) | **Put** /api/apps/{id}/wiki | Update an App Wiki Page
[**UpdateWikiCloud**](WikiAPI.md#UpdateWikiCloud) | **Put** /api/zones/{id}/wiki | Update a Cloud Wiki Page
[**UpdateWikiCluster**](WikiAPI.md#UpdateWikiCluster) | **Put** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page
[**UpdateWikiGroup**](WikiAPI.md#UpdateWikiGroup) | **Put** /api/groups/{id}/wiki | Update a Group Wiki Page
[**UpdateWikiInstance**](WikiAPI.md#UpdateWikiInstance) | **Put** /api/instances/{id}/wiki | Update an Instance Wiki Page
[**UpdateWikiServer**](WikiAPI.md#UpdateWikiServer) | **Put** /api/servers/{id}/wiki | Update a Server Wiki Page



## AddWiki

> UpdateWikiApp200Response AddWiki(ctx).AddWikiRequest(addWikiRequest).Execute()

Create a Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	addWikiRequest := *openapiclient.NewAddWikiRequest(*openapiclient.NewAddWikiRequestPage("Sample Doc", "info", "A sample document in **markdown**.")) // AddWikiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.AddWiki(context.Background()).AddWikiRequest(addWikiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.AddWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWiki`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.AddWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addWikiRequest** | [**AddWikiRequest**](AddWikiRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWiki

> GetWikiApp200Response GetWiki(ctx, id).Execute()

Retrieves a specific Wiki page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWiki(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWiki`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiApp

> GetWikiApp200Response GetWikiApp(ctx, id).Execute()

Retrieves an App Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiApp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiApp`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiCategories

> GetWikiCategories200Response GetWikiCategories(ctx).Phrase(phrase).PagePhrase(pagePhrase).Execute()

Retrieves all Wiki categories associated with the account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	pagePhrase := "pagePhrase_example" // string | If specified will return a partial match on page name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiCategories(context.Background()).Phrase(phrase).PagePhrase(pagePhrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiCategories`: GetWikiCategories200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **pagePhrase** | **string** | If specified will return a partial match on page name | 

### Return type

[**GetWikiCategories200Response**](GetWikiCategories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiCloud

> GetWikiApp200Response GetWikiCloud(ctx, id).Execute()

Retrieves a Cloud Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiCloud(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiCloud`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiCluster

> GetWikiApp200Response GetWikiCluster(ctx, clusterId).Execute()

Retrieves a Cluster Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	clusterId := int32(5) // int32 | The ID of the cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiCluster`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiGroup

> GetWikiApp200Response GetWikiGroup(ctx, id).Execute()

Retrieves a Group Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiGroup`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiInstance

> GetWikiApp200Response GetWikiInstance(ctx, id).Execute()

Retrieves an Instance Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiInstance`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiServer

> GetWikiApp200Response GetWikiServer(ctx, id).Execute()

Retrieves a Server Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.GetWikiServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.GetWikiServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiServer`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.GetWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWiki

> GetWikiApp200Response ListWiki(ctx).Name(name).Phrase(phrase).Execute()

Retrieves wiki pages associated with the account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.ListWiki(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.ListWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWiki`: GetWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.ListWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWiki

> DeleteAlerts200Response RemoveWiki(ctx, id).Execute()

Deletes a Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.RemoveWiki(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.RemoveWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWiki`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.RemoveWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWiki

> UpdateWikiApp200Response UpdateWiki(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Updates a Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWiki(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWiki`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiApp

> UpdateWikiApp200Response UpdateWikiApp(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update an App Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiApp(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiApp`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiCloud

> UpdateWikiApp200Response UpdateWikiCloud(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update a Cloud Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiCloud(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiCloud`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiCluster

> UpdateWikiApp200Response UpdateWikiCluster(ctx, clusterId).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update a Cluster Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	clusterId := int32(5) // int32 | The ID of the cluster
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiCluster(context.Background(), clusterId).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiCluster`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiGroup

> UpdateWikiApp200Response UpdateWikiGroup(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update a Group Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiGroup(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiGroup`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiInstance

> UpdateWikiApp200Response UpdateWikiInstance(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update an Instance Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiInstance(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiInstance`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiServer

> UpdateWikiApp200Response UpdateWikiServer(ctx, id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()

Update a Server Wiki Page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateWikiAppRequest := *openapiclient.NewUpdateWikiAppRequest() // UpdateWikiAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.UpdateWikiServer(context.Background(), id).UpdateWikiAppRequest(updateWikiAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.UpdateWikiServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWikiServer`: UpdateWikiApp200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.UpdateWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWikiAppRequest** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md) |  | 

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \HostsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBaremetalHost**](HostsAPI.md#AddBaremetalHost) | **Post** /api/servers | Add a Baremetal Host
[**AssignDevice**](HostsAPI.md#AssignDevice) | **Put** /api/servers/{id}/devices/{deviceId}/assign | Assign a Host Device
[**AttachDevice**](HostsAPI.md#AttachDevice) | **Put** /api/servers/{id}/devices/{deviceId}/attach | Attach a Host Device
[**DetachDevice**](HostsAPI.md#DetachDevice) | **Put** /api/servers/{id}/devices/{deviceId}/detach | Detach a Host Device
[**EnableMaintenanceMode**](HostsAPI.md#EnableMaintenanceMode) | **Put** /api/servers/{id}/maintenance | Enable Maintenance Mode
[**GetHost**](HostsAPI.md#GetHost) | **Get** /api/servers/{id} | Get a Specific Host
[**GetHostSnpshots**](HostsAPI.md#GetHostSnpshots) | **Get** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**GetHostType**](HostsAPI.md#GetHostType) | **Get** /api/server-types/{id} | Get a Specific Host Type
[**LeaveMaintenanceMode**](HostsAPI.md#LeaveMaintenanceMode) | **Put** /api/servers/{id}/leave-maintenance | Leave Maintenance Mode
[**ListHostDevices**](HostsAPI.md#ListHostDevices) | **Get** /api/servers/{id}/devices | Get list of devices for a Host
[**ListHostTypes**](HostsAPI.md#ListHostTypes) | **Get** /api/server-types | Host Types
[**ListHosts**](HostsAPI.md#ListHosts) | **Get** /api/servers | Get All Hosts
[**ListServerServicePlans**](HostsAPI.md#ListServerServicePlans) | **Get** /api/servers/service-plans | Get Available Service Plans for a Host
[**ManageHostPlacement**](HostsAPI.md#ManageHostPlacement) | **Put** /api/servers/{id}/placement | Manage Host Placement
[**RemoveHost**](HostsAPI.md#RemoveHost) | **Delete** /api/servers/{id} | Delete a Host
[**RestartHost**](HostsAPI.md#RestartHost) | **Put** /api/servers/{id}/restart | Restart a Host
[**StartHost**](HostsAPI.md#StartHost) | **Put** /api/servers/{id}/start | Start a Host
[**StopHost**](HostsAPI.md#StopHost) | **Put** /api/servers/{id}/stop | Stop a Host
[**UpdateHost**](HostsAPI.md#UpdateHost) | **Put** /api/servers/{id} | Updating a Host
[**UpdateHostAssignTenant**](HostsAPI.md#UpdateHostAssignTenant) | **Put** /api/servers/{id}/assign-account | Assign To Tenant
[**UpdateHostCloud**](HostsAPI.md#UpdateHostCloud) | **Put** /api/servers/change-cloud | Change Server Cloud
[**UpdateHostExecuteWorkflow**](HostsAPI.md#UpdateHostExecuteWorkflow) | **Put** /api/servers/{id}/workflow | Run Workflow on a Host
[**UpdateHostInstallAgent**](HostsAPI.md#UpdateHostInstallAgent) | **Put** /api/servers/{id}/install-agent | Install Agent
[**UpdateHostManaged**](HostsAPI.md#UpdateHostManaged) | **Put** /api/servers/{id}/make-managed | Convert To Managed
[**UpdateHostResize**](HostsAPI.md#UpdateHostResize) | **Put** /api/servers/{id}/resize | Resize a Host
[**UpdateHostUpgradeAgent**](HostsAPI.md#UpdateHostUpgradeAgent) | **Put** /api/servers/{id}/upgrade | Upgrade Agent
[**UpdateServerNetworkInterface**](HostsAPI.md#UpdateServerNetworkInterface) | **Put** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network



## AddBaremetalHost

> AddBaremetalHost200Response AddBaremetalHost(ctx).AddBaremetalHostRequest(addBaremetalHostRequest).Execute()

Add a Baremetal Host



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
	addBaremetalHostRequest := *openapiclient.NewAddBaremetalHostRequest() // AddBaremetalHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.AddBaremetalHost(context.Background()).AddBaremetalHostRequest(addBaremetalHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.AddBaremetalHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBaremetalHost`: AddBaremetalHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.AddBaremetalHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBaremetalHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addBaremetalHostRequest** | [**AddBaremetalHostRequest**](AddBaremetalHostRequest.md) |  | 

### Return type

[**AddBaremetalHost200Response**](AddBaremetalHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignDevice

> RestartClusterContainer200Response AssignDevice(ctx, id, deviceId).AssignDeviceRequest(assignDeviceRequest).Execute()

Assign a Host Device



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
	id := int32(5) // int32 | The ID of the host
	deviceId := int32(99) // int32 | The ID of the device
	assignDeviceRequest := *openapiclient.NewAssignDeviceRequest(int64(123)) // AssignDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.AssignDevice(context.Background(), id, deviceId).AssignDeviceRequest(assignDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.AssignDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignDevice`: RestartClusterContainer200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.AssignDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the host | 
**deviceId** | **int32** | The ID of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignDeviceRequest** | [**AssignDeviceRequest**](AssignDeviceRequest.md) |  | 

### Return type

[**RestartClusterContainer200Response**](RestartClusterContainer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachDevice

> RestartClusterContainer200Response AttachDevice(ctx, id, deviceId).Execute()

Attach a Host Device



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
	id := int32(5) // int32 | The ID of the host
	deviceId := int32(99) // int32 | The ID of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.AttachDevice(context.Background(), id, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.AttachDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachDevice`: RestartClusterContainer200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.AttachDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the host | 
**deviceId** | **int32** | The ID of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestartClusterContainer200Response**](RestartClusterContainer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachDevice

> RestartClusterContainer200Response DetachDevice(ctx, id, deviceId).Execute()

Detach a Host Device



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
	id := int32(5) // int32 | The ID of the host
	deviceId := int32(99) // int32 | The ID of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.DetachDevice(context.Background(), id, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.DetachDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachDevice`: RestartClusterContainer200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.DetachDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the host | 
**deviceId** | **int32** | The ID of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestartClusterContainer200Response**](RestartClusterContainer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableMaintenanceMode

> DeleteAlerts200Response EnableMaintenanceMode(ctx, id).EnableMaintenanceModeRequest(enableMaintenanceModeRequest).Execute()

Enable Maintenance Mode



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
	enableMaintenanceModeRequest := *openapiclient.NewEnableMaintenanceModeRequest() // EnableMaintenanceModeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.EnableMaintenanceMode(context.Background(), id).EnableMaintenanceModeRequest(enableMaintenanceModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.EnableMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableMaintenanceMode`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.EnableMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableMaintenanceModeRequest** | [**EnableMaintenanceModeRequest**](EnableMaintenanceModeRequest.md) |  | 

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> GetHost200Response GetHost(ctx, id).Execute()

Get a Specific Host



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
	resp, r, err := apiClient.HostsAPI.GetHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: GetHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHost200Response**](GetHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostSnpshots

> GetHostSnpshots200Response GetHostSnpshots(ctx, id).Execute()

Get list of snapshots for a Host



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
	resp, r, err := apiClient.HostsAPI.GetHostSnpshots(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHostSnpshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostSnpshots`: GetHostSnpshots200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHostSnpshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostSnpshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHostSnpshots200Response**](GetHostSnpshots200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostType

> GetHostType200Response GetHostType(ctx, id).Execute()

Get a Specific Host Type



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
	resp, r, err := apiClient.HostsAPI.GetHostType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHostType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostType`: GetHostType200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHostType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHostType200Response**](GetHostType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveMaintenanceMode

> DeleteAlerts200Response LeaveMaintenanceMode(ctx, id).Execute()

Leave Maintenance Mode



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
	resp, r, err := apiClient.HostsAPI.LeaveMaintenanceMode(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.LeaveMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaveMaintenanceMode`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.LeaveMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveMaintenanceModeRequest struct via the builder pattern


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


## ListHostDevices

> ListHostDevices200Response ListHostDevices(ctx, id).Execute()

Get list of devices for a Host



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
	id := int32(5) // int32 | The ID of the host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListHostDevices(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListHostDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHostDevices`: ListHostDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListHostDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHostDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListHostDevices200Response**](ListHostDevices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostTypes

> ListHostTypes200Response ListHostTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).ZoneType(zoneType).Creatable(creatable).Execute()

Host Types



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
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	name := "example" // string | Filter by name (optional)
	code := "azr" // string | If specified will return an exact match on code (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
	zoneType := "zoneType_example" // string | Filter by Cloud Type code. (optional)
	creatable := true // bool | Filter by creatable flag. This is whether or not it can be provisioned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListHostTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).ZoneType(zoneType).Creatable(creatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListHostTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHostTypes`: ListHostTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListHostTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
 **zoneType** | **string** | Filter by Cloud Type code. | 
 **creatable** | **bool** | Filter by creatable flag. This is whether or not it can be provisioned. | 

### Return type

[**ListHostTypes200Response**](ListHostTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> ListHosts200Response ListHosts(ctx).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).TagsName(tagsName).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Execute()

Get All Hosts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
	siteId := int64(7) // int64 | The Site ID for Filtering (optional)
	clusterId := int64(135) // int64 | The Cluster ID(s) for filtering. Accepts multiple values. (optional)
	managed := false // bool | Filter by managed (true) or unmanaged (false) (optional)
	serverType := "vmwareHypervisor" // string | Filter by server type code (optional)
	powerState := "off" // string | Filter by power status (optional)
	ip := "192.168.1.45" // string | Filter by IP address (optional)
	vm := false // bool | Filter to show only Virtual Machines (true) (optional)
	vmHypervisor := false // bool | Filter to show only VM Hypervisors (true) (optional)
	bareMetalHost := false // bool | Filter to show only Baremetal Servers (optional)
	status := "status_example" // string | Filter by status (optional)
	agentInstalled := true // bool | Filter by agent installed (true) (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	createdBy := int64(63) // int64 | The User ID for Filtering (optional)
	tagsName := "value" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
	metadata := "tags.env=qa&tags.env=test" // string | Alias for tags (optional)
	uuid := "uuid_example" // string | Filter by UUID (optional)
	externalId := "externalId_example" // string | Filter by External ID (optional)
	internalId := "internalId_example" // string | Filter by Internal ID (optional)
	externalUniquelId := "externalUniquelId_example" // string | Filter by External Unique ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListHosts(context.Background()).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).TagsName(tagsName).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: ListHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **siteId** | **int64** | The Site ID for Filtering | 
 **clusterId** | **int64** | The Cluster ID(s) for filtering. Accepts multiple values. | 
 **managed** | **bool** | Filter by managed (true) or unmanaged (false) | 
 **serverType** | **string** | Filter by server type code | 
 **powerState** | **string** | Filter by power status | 
 **ip** | **string** | Filter by IP address | 
 **vm** | **bool** | Filter to show only Virtual Machines (true) | 
 **vmHypervisor** | **bool** | Filter to show only VM Hypervisors (true) | 
 **bareMetalHost** | **bool** | Filter to show only Baremetal Servers | 
 **status** | **string** | Filter by status | 
 **agentInstalled** | **bool** | Filter by agent installed (true) | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **createdBy** | **int64** | The User ID for Filtering | 
 **tagsName** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **metadata** | **string** | Alias for tags | 
 **uuid** | **string** | Filter by UUID | 
 **externalId** | **string** | Filter by External ID | 
 **internalId** | **string** | Filter by Internal ID | 
 **externalUniquelId** | **string** | Filter by External Unique ID | 

### Return type

[**ListHosts200Response**](ListHosts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerServicePlans

> ListServerServicePlans200Response ListServerServicePlans(ctx).ZoneId(zoneId).ServerTypeId(serverTypeId).SiteId(siteId).Execute()

Get Available Service Plans for a Host



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
	zoneId := int64(3) // int64 | The Zone ID for Filtering
	serverTypeId := int64(5) // int64 | The ID of the Host Type (optional)
	siteId := int64(7) // int64 | The Site ID for Filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListServerServicePlans(context.Background()).ZoneId(zoneId).ServerTypeId(serverTypeId).SiteId(siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListServerServicePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerServicePlans`: ListServerServicePlans200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListServerServicePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **serverTypeId** | **int64** | The ID of the Host Type | 
 **siteId** | **int64** | The Site ID for Filtering | 

### Return type

[**ListServerServicePlans200Response**](ListServerServicePlans200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageHostPlacement

> DeleteAlerts200Response ManageHostPlacement(ctx, id).ManageHostPlacementRequest(manageHostPlacementRequest).Execute()

Manage Host Placement



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
	manageHostPlacementRequest := *openapiclient.NewManageHostPlacementRequest() // ManageHostPlacementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ManageHostPlacement(context.Background(), id).ManageHostPlacementRequest(manageHostPlacementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ManageHostPlacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManageHostPlacement`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ManageHostPlacement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageHostPlacementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageHostPlacementRequest** | [**ManageHostPlacementRequest**](ManageHostPlacementRequest.md) |  | 

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHost

> DeleteAlerts200Response RemoveHost(ctx, id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()

Delete a Host



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
	removeResources := "off" // string | Remove Resources (optional) (default to "on")
	removeInstances := "on" // string | Remove Instances (optional) (default to "off")
	preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
	releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
	releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
	force := "on" // string | Force Delete (optional) (default to "off")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.RemoveHost(context.Background(), id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.RemoveHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveHost`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.RemoveHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **string** | Remove Resources | [default to &quot;on&quot;]
 **removeInstances** | **string** | Remove Instances | [default to &quot;off&quot;]
 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]

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


## RestartHost

> UpdateHostAssignTenant200Response RestartHost(ctx, id).Execute()

Restart a Host



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
	resp, r, err := apiClient.HostsAPI.RestartHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.RestartHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartHost`: UpdateHostAssignTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.RestartHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartHost

> DeleteAlerts200Response StartHost(ctx, id).Execute()

Start a Host



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
	resp, r, err := apiClient.HostsAPI.StartHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.StartHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartHost`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.StartHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartHostRequest struct via the builder pattern


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


## StopHost

> DeleteAlerts200Response StopHost(ctx, id).Execute()

Stop a Host



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
	resp, r, err := apiClient.HostsAPI.StopHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.StopHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopHost`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.StopHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopHostRequest struct via the builder pattern


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


## UpdateHost

> GetHost200Response UpdateHost(ctx, id).UpdateHostRequest(updateHostRequest).Execute()

Updating a Host



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
	updateHostRequest := *openapiclient.NewUpdateHostRequest() // UpdateHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHost(context.Background(), id).UpdateHostRequest(updateHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHost`: GetHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostRequest** | [**UpdateHostRequest**](UpdateHostRequest.md) |  | 

### Return type

[**GetHost200Response**](GetHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostAssignTenant

> UpdateHostAssignTenant200Response UpdateHostAssignTenant(ctx, id).AccountId(accountId).UpdateHostAssignTenantRequest(updateHostAssignTenantRequest).Execute()

Assign To Tenant



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
	accountId := int64(3) // int64 | ID of the Tenant (optional)
	updateHostAssignTenantRequest := *openapiclient.NewUpdateHostAssignTenantRequest() // UpdateHostAssignTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostAssignTenant(context.Background(), id).AccountId(accountId).UpdateHostAssignTenantRequest(updateHostAssignTenantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostAssignTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostAssignTenant`: UpdateHostAssignTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostAssignTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostAssignTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **int64** | ID of the Tenant | 
 **updateHostAssignTenantRequest** | [**UpdateHostAssignTenantRequest**](UpdateHostAssignTenantRequest.md) |  | 

### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostCloud

> UpdateHostCloud200Response UpdateHostCloud(ctx).UpdateHostCloudRequest(updateHostCloudRequest).Execute()

Change Server Cloud



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
	updateHostCloudRequest := *openapiclient.NewUpdateHostCloudRequest() // UpdateHostCloudRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostCloud(context.Background()).UpdateHostCloudRequest(updateHostCloudRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostCloud`: UpdateHostCloud200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostCloud`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateHostCloudRequest** | [**UpdateHostCloudRequest**](UpdateHostCloudRequest.md) |  | 

### Return type

[**UpdateHostCloud200Response**](UpdateHostCloud200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostExecuteWorkflow

> RunWorkflowInstance200Response UpdateHostExecuteWorkflow(ctx, id).WorkflowId(workflowId).WorkflowName(workflowName).UpdateHostExecuteWorkflowRequest(updateHostExecuteWorkflowRequest).Execute()

Run Workflow on a Host



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
	workflowId := int64(3) // int64 | ID of the workflow to execute (optional)
	workflowName := "myworkflow" // string | Name of the workflow to execute (optional)
	updateHostExecuteWorkflowRequest := *openapiclient.NewUpdateHostExecuteWorkflowRequest() // UpdateHostExecuteWorkflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostExecuteWorkflow(context.Background(), id).WorkflowId(workflowId).WorkflowName(workflowName).UpdateHostExecuteWorkflowRequest(updateHostExecuteWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostExecuteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostExecuteWorkflow`: RunWorkflowInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowId** | **int64** | ID of the workflow to execute | 
 **workflowName** | **string** | Name of the workflow to execute | 
 **updateHostExecuteWorkflowRequest** | [**UpdateHostExecuteWorkflowRequest**](UpdateHostExecuteWorkflowRequest.md) |  | 

### Return type

[**RunWorkflowInstance200Response**](RunWorkflowInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostInstallAgent

> UpdateHostInstallAgent200Response UpdateHostInstallAgent(ctx, id).UpdateHostInstallAgentRequest(updateHostInstallAgentRequest).Execute()

Install Agent



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
	updateHostInstallAgentRequest := *openapiclient.NewUpdateHostInstallAgentRequest() // UpdateHostInstallAgentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostInstallAgent(context.Background(), id).UpdateHostInstallAgentRequest(updateHostInstallAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostInstallAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostInstallAgent`: UpdateHostInstallAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostInstallAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostInstallAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostInstallAgentRequest** | [**UpdateHostInstallAgentRequest**](UpdateHostInstallAgentRequest.md) |  | 

### Return type

[**UpdateHostInstallAgent200Response**](UpdateHostInstallAgent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostManaged

> UpdateHostInstallAgent200Response UpdateHostManaged(ctx, id).UpdateHostManagedRequest(updateHostManagedRequest).Execute()

Convert To Managed



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
	updateHostManagedRequest := *openapiclient.NewUpdateHostManagedRequest() // UpdateHostManagedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostManaged(context.Background(), id).UpdateHostManagedRequest(updateHostManagedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostManaged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostManaged`: UpdateHostInstallAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostManaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostManagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostManagedRequest** | [**UpdateHostManagedRequest**](UpdateHostManagedRequest.md) |  | 

### Return type

[**UpdateHostInstallAgent200Response**](UpdateHostInstallAgent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostResize

> UpdateHostAssignTenant200Response UpdateHostResize(ctx, id).UpdateHostResizeRequest(updateHostResizeRequest).Execute()

Resize a Host



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
	updateHostResizeRequest := *openapiclient.NewUpdateHostResizeRequest() // UpdateHostResizeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHostResize(context.Background(), id).UpdateHostResizeRequest(updateHostResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostResize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostResize`: UpdateHostAssignTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostResizeRequest** | [**UpdateHostResizeRequest**](UpdateHostResizeRequest.md) |  | 

### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostUpgradeAgent

> DeleteAlerts200Response UpdateHostUpgradeAgent(ctx, id).Execute()

Upgrade Agent



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
	resp, r, err := apiClient.HostsAPI.UpdateHostUpgradeAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHostUpgradeAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostUpgradeAgent`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHostUpgradeAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostUpgradeAgentRequest struct via the builder pattern


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


## UpdateServerNetworkInterface

> UpdateServerNetworkInterface200Response UpdateServerNetworkInterface(ctx, id, networkInterfaceId).UpdateInstanceNetworkInterfaceRequest(updateInstanceNetworkInterfaceRequest).Execute()

Updating a label for a Server's Network



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
	networkInterfaceId := float32(7) // float32 | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
	updateInstanceNetworkInterfaceRequest := *openapiclient.NewUpdateInstanceNetworkInterfaceRequest() // UpdateInstanceNetworkInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateServerNetworkInterface(context.Background(), id, networkInterfaceId).UpdateInstanceNetworkInterfaceRequest(updateInstanceNetworkInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateServerNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerNetworkInterface`: UpdateServerNetworkInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateServerNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**networkInterfaceId** | **float32** | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInstanceNetworkInterfaceRequest** | [**UpdateInstanceNetworkInterfaceRequest**](UpdateInstanceNetworkInterfaceRequest.md) |  | 

### Return type

[**UpdateServerNetworkInterface200Response**](UpdateServerNetworkInterface200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


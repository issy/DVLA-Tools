# \VehicleApi

All URIs are relative to *https://driver-vehicle-licensing.api.gov.uk/vehicle-enquiry*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVehicleDetailsByRegistrationNumber**](VehicleApi.md#GetVehicleDetailsByRegistrationNumber) | **Post** /v1/vehicles | Get vehicle details by registration number



## GetVehicleDetailsByRegistrationNumber

> Vehicle GetVehicleDetailsByRegistrationNumber(ctx).XApiKey(xApiKey).VehicleRequest(vehicleRequest).XCorrelationId(xCorrelationId).Execute()

Get vehicle details by registration number



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
    xApiKey := "xApiKey_example" // string | Client Specific API Key
    vehicleRequest := *openapiclient.NewVehicleRequest() // VehicleRequest | Registration number of the vehicle to find details for
    xCorrelationId := "xCorrelationId_example" // string | Consumer Correlation ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VehicleApi.GetVehicleDetailsByRegistrationNumber(context.Background()).XApiKey(xApiKey).VehicleRequest(vehicleRequest).XCorrelationId(xCorrelationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.GetVehicleDetailsByRegistrationNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVehicleDetailsByRegistrationNumber`: Vehicle
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.GetVehicleDetailsByRegistrationNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehicleDetailsByRegistrationNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Client Specific API Key | 
 **vehicleRequest** | [**VehicleRequest**](VehicleRequest.md) | Registration number of the vehicle to find details for | 
 **xCorrelationId** | **string** | Consumer Correlation ID | 

### Return type

[**Vehicle**](Vehicle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


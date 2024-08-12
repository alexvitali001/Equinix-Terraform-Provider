# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionById**](MarketplaceSubscriptionsApi.md#GetSubscriptionById) | **Get** /fabric/v4/marketplaceSubscriptions/{subscriptionId} | Get Subscription

# **GetSubscriptionById**
> SubscriptionResponse GetSubscriptionById(ctx, subscriptionId)
Get Subscription

The API provides capability to get subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | [**string**](.md)| Subscription UUID | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


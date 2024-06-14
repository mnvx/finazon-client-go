# {{classname}}

All URIs are relative to *https://api.finazon.io/v1.2/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiUsage**](ReferenceApi.md#GetApiUsage) | **Get** /api_usage | Api usage
[**GetDatasets**](ReferenceApi.md#GetDatasets) | **Get** /datasets | List of Finazon datasets
[**GetExchangesCrypto**](ReferenceApi.md#GetExchangesCrypto) | **Get** /markets/crypto | List of crypto markets
[**GetExchangesStocks**](ReferenceApi.md#GetExchangesStocks) | **Get** /markets/stocks | List of stock markets
[**GetMarketCenter**](ReferenceApi.md#GetMarketCenter) | **Get** /sip/market_center | List of market centers
[**GetPublishers**](ReferenceApi.md#GetPublishers) | **Get** /publishers | List of Finazon publishers
[**GetSymbolsCrypto**](ReferenceApi.md#GetSymbolsCrypto) | **Get** /tickers/crypto | List of crypto pairs
[**GetSymbolsForex**](ReferenceApi.md#GetSymbolsForex) | **Get** /tickers/forex | List of forex ticker symbols
[**GetSymbolsStocks**](ReferenceApi.md#GetSymbolsStocks) | **Get** /tickers/stocks | List of stock ticker symbols
[**GetSymbolsUSStocks**](ReferenceApi.md#GetSymbolsUSStocks) | **Get** /tickers/us_stocks | List of US stock ticker symbols

# **GetApiUsage**
> ApiUsageResponseBody GetApiUsage(ctx, optional)
Api usage

Returns a list of datasets with available API calls and limits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetApiUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetApiUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **optional.String**| Product code | 

### Return type

[**ApiUsageResponseBody**](ApiUsageResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasets**
> DatasetsResponseBody GetDatasets(ctx, optional)
List of Finazon datasets

Returns a list of all datasets available at Finazon.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetDatasetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetDatasetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **optional.String**| Filter by Finazon&#x27;s dataset code | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**DatasetsResponseBody**](DatasetsResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExchangesCrypto**
> ExchangesCryptoResponseBody GetExchangesCrypto(ctx, optional)
List of crypto markets

Returns a list of supported crypto markets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetExchangesCryptoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetExchangesCryptoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**ExchangesCryptoResponseBody**](ExchangesCryptoResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExchangesStocks**
> ExchangesStocksResponseBody GetExchangesStocks(ctx, optional)
List of stock markets

Returns a list of supported stock markets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetExchangesStocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetExchangesStocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **name** | **optional.String**| Filter by market name | 
 **code** | **optional.String**| Filter by market identifier code (MIC) under ISO 10383 standard | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**ExchangesStocksResponseBody**](ExchangesStocksResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketCenter**
> MarketCenterResponseBody GetMarketCenter(ctx, )
List of market centers

Returns a list of market centers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MarketCenterResponseBody**](MarketCenterResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublishers**
> PublishersResponseBody GetPublishers(ctx, optional)
List of Finazon publishers

Returns a list of all publishers available at Finazon.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetPublishersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetPublishersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **optional.String**| Filter by Finazon&#x27;s publisher code | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**PublishersResponseBody**](PublishersResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSymbolsCrypto**
> SymbolsCryptoResponseBody GetSymbolsCrypto(ctx, optional)
List of crypto pairs

Returns a list of cryptocurrency ticker symbols (pairs). This list is updated daily.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetSymbolsCryptoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetSymbolsCryptoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataset** | **optional.String**| Filter by Finazon&#x27;s dataset code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**SymbolsCryptoResponseBody**](SymbolsCryptoResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSymbolsForex**
> SymbolsForexResponseBody GetSymbolsForex(ctx, optional)
List of forex ticker symbols

Returns a list of forex ticker symbols (pairs). This list is updated daily.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetSymbolsForexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetSymbolsForexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ticker** | **optional.String**| Filter by ticker symbol | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**SymbolsForexResponseBody**](SymbolsForexResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSymbolsStocks**
> SymbolsStocksResponseBody GetSymbolsStocks(ctx, optional)
List of stock ticker symbols

Returns a list of stock ticker symbols. This list is updated daily.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetSymbolsStocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetSymbolsStocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **dataset** | **optional.String**| Filter by Finazon&#x27;s dataset code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**SymbolsStocksResponseBody**](SymbolsStocksResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSymbolsUSStocks**
> SymbolsUsStocksResponseBody GetSymbolsUSStocks(ctx, optional)
List of US stock ticker symbols

Returns a list of US stock ticker symbols. This list is updated daily.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceApiGetSymbolsUSStocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceApiGetSymbolsUSStocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | 
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**SymbolsUsStocksResponseBody**](SymbolsUSStocksResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


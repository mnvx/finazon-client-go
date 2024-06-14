# {{classname}}

All URIs are relative to *https://api.finazon.io/v1.2/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBenzingaDividendsCalendar**](DataApi.md#GetBenzingaDividendsCalendar) | **Get** /benzinga/dividends_calendar | Dividends calendar
[**GetBenzingaEarningsCalendar**](DataApi.md#GetBenzingaEarningsCalendar) | **Get** /benzinga/earnings_calendar | Earnings calendar
[**GetBenzingaIPO**](DataApi.md#GetBenzingaIPO) | **Get** /benzinga/ipo | IPO data
[**GetBenzingaNews**](DataApi.md#GetBenzingaNews) | **Get** /benzinga/news | News articles
[**GetBinanceQuotes**](DataApi.md#GetBinanceQuotes) | **Get** /binance/time_series | Time series
[**GetCryptoQuotes**](DataApi.md#GetCryptoQuotes) | **Get** /crypto/time_series | Time series
[**GetFilings**](DataApi.md#GetFilings) | **Get** /sec/archive | Filings
[**GetForexQuotes**](DataApi.md#GetForexQuotes) | **Get** /forex/time_series | Time series
[**GetPriceBinance**](DataApi.md#GetPriceBinance) | **Get** /binance/price | Price
[**GetPriceCommon**](DataApi.md#GetPriceCommon) | **Get** /price | Price
[**GetPriceCrypto**](DataApi.md#GetPriceCrypto) | **Get** /crypto/price | Price
[**GetPriceForex**](DataApi.md#GetPriceForex) | **Get** /forex/price | Forex price
[**GetPriceSip**](DataApi.md#GetPriceSip) | **Get** /sip_non_pro/price | Price
[**GetPriceSip_0**](DataApi.md#GetPriceSip_0) | **Get** /sip_pro/price | Price
[**GetPriceUsStocks**](DataApi.md#GetPriceUsStocks) | **Get** /us_stocks_essential/price | Price
[**GetQuotes**](DataApi.md#GetQuotes) | **Get** /time_series | Time series
[**GetSipTrades**](DataApi.md#GetSipTrades) | **Get** /sip/trades | SIP trades
[**GetTechnicalIndicatorAtr**](DataApi.md#GetTechnicalIndicatorAtr) | **Get** /time_series/atr | ATR Technical indicators
[**GetTechnicalIndicatorBBands**](DataApi.md#GetTechnicalIndicatorBBands) | **Get** /time_series/bbands | Overlap Studies
[**GetTechnicalIndicatorIchimoku**](DataApi.md#GetTechnicalIndicatorIchimoku) | **Get** /time_series/ichimoku | Overlap Studies
[**GetTechnicalIndicatorMa**](DataApi.md#GetTechnicalIndicatorMa) | **Get** /time_series/ma | Overlap Studies
[**GetTechnicalIndicatorMacd**](DataApi.md#GetTechnicalIndicatorMacd) | **Get** /time_series/macd | Momentum Indicators
[**GetTechnicalIndicatorObv**](DataApi.md#GetTechnicalIndicatorObv) | **Get** /time_series/obv | Volume Indicators
[**GetTechnicalIndicatorRsi**](DataApi.md#GetTechnicalIndicatorRsi) | **Get** /time_series/rsi | Momentum Indicators
[**GetTechnicalIndicatorSar**](DataApi.md#GetTechnicalIndicatorSar) | **Get** /time_series/sar | Overlap Studies
[**GetTechnicalIndicatorStoch**](DataApi.md#GetTechnicalIndicatorStoch) | **Get** /time_series/stoch | Momentum Indicators
[**GetTickerSnapshot**](DataApi.md#GetTickerSnapshot) | **Get** /ticker/snapshot | Ticker snapshot
[**GetTrades**](DataApi.md#GetTrades) | **Get** /trades | Trades

# **GetBenzingaDividendsCalendar**
> BenzingaDividendsCalendarResponseBody GetBenzingaDividendsCalendar(ctx, ticker, optional)
Dividends calendar

Returns the dividends calendar from Benzinga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetBenzingaDividendsCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetBenzingaDividendsCalendarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Specifies the exact date to get the data for | 
 **startAt** | **optional.Int64**| Filter events by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter events by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 100]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 

### Return type

[**BenzingaDividendsCalendarResponseBody**](BenzingaDividendsCalendarResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBenzingaEarningsCalendar**
> BenzingaEarningsCalendarResponseBody GetBenzingaEarningsCalendar(ctx, ticker, optional)
Earnings calendar

Returns the earnings calendar from Benzinga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetBenzingaEarningsCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetBenzingaEarningsCalendarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Specifies the exact date to get the data for | 
 **startAt** | **optional.Int64**| Filter events by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter events by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 100]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 

### Return type

[**BenzingaEarningsCalendarResponseBody**](BenzingaEarningsCalendarResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBenzingaIPO**
> BenzingaIpoResponseBody GetBenzingaIPO(ctx, optional)
IPO data

Returns IPO data from Benzinga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetBenzingaIPOOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetBenzingaIPOOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **optional.Int64**| Filter events by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter events by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 100]
 **order** | **optional.String**| Sorting order of the output series | [default to asc]
 **exchange** | **optional.String**| Exchange where instrument is traded | 

### Return type

[**BenzingaIpoResponseBody**](BenzingaIPOResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBenzingaNews**
> BenzingaNewsResponseBody GetBenzingaNews(ctx, ticker, optional)
News articles

Returns the news articles from Benzinga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetBenzingaNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetBenzingaNewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Specifies the exact date to get the data for | 
 **startAt** | **optional.Int64**| Filter events by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter events by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 100]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 

### Return type

[**BenzingaNewsResponseBody**](BenzingaNewsResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBinanceQuotes**
> []QuoteBinanceItem GetBinanceQuotes(ctx, ticker, interval, optional)
Time series

This endpoint returns a time series of data points for any given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetBinanceQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetBinanceQuotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]

### Return type

[**[]QuoteBinanceItem**](QuoteBinanceItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCryptoQuotes**
> []QuoteItem GetCryptoQuotes(ctx, interval, ticker, optional)
Time series

This endpoint returns a time series of data points for any given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interval** | **string**| Interval between two consecutive points in time series | 
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetCryptoQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetCryptoQuotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]

### Return type

[**[]QuoteItem**](QuoteItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilings**
> []Filing GetFilings(ctx, optional)
Filings

Real-time and historical access to all forms, filings, and exhibits directly from the SEC's EDGAR system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetFilingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetFilingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cikCode** | **optional.Int64**| Filter by Central Index Key | 
 **ticker** | **optional.String**| Filter by ticker | 
 **formType** | **optional.String**| Filter by form types | 
 **filledFromTs** | **optional.Int64**| Filter by filled time from | 
 **filledToTs** | **optional.Int64**| Filter by filled time to | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 10]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 

### Return type

[**[]Filing**](Filing.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForexQuotes**
> []QuoteForexItem GetForexQuotes(ctx, interval, ticker, optional)
Time series

This endpoint returns a time series of data points for any given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interval** | **string**| Interval between two consecutive points in time series | 
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetForexQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetForexQuotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]

### Return type

[**[]QuoteForexItem**](QuoteForexItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceBinance**
> PriceResponseBody GetPriceBinance(ctx, optional)
Price

Returns price value for given binance ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceBinanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceBinanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceCommon**
> PriceResponseBody GetPriceCommon(ctx, dataset, optional)
Price

Returns price value for given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
 **optional** | ***DataApiGetPriceCommonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceCommonOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceCrypto**
> PriceResponseBody GetPriceCrypto(ctx, optional)
Price

Returns price value for given crypto ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceCryptoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceCryptoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceForex**
> PriceResponseBody GetPriceForex(ctx, optional)
Forex price

Returns price value for given forex ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceForexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceForexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceSip**
> PriceResponseBody GetPriceSip(ctx, optional)
Price

Returns price value for given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceSipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceSipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceSip_0**
> PriceResponseBody GetPriceSip_0(ctx, optional)
Price

Returns price value for given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceSip_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceSip_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceUsStocks**
> PriceResponseBody GetPriceUsStocks(ctx, optional)
Price

Returns price value for given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataApiGetPriceUsStocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetPriceUsStocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **optional.Int64**| Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at | 
 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **ticker** | **optional.String**| Filter by ticker symbol | 

### Return type

[**PriceResponseBody**](PriceResponseBody.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotes**
> []QuoteItem GetQuotes(ctx, dataset, ticker, interval, optional)
Time series

This endpoint returns a time series of data points for any given ticker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetQuotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]

### Return type

[**[]QuoteItem**](QuoteItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSipTrades**
> []SipTradesItem GetSipTrades(ctx, ticker, optional)
SIP trades

Returns detailed information on trades executed through the Securities Information Processor (SIP).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetSipTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetSipTradesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market center | 
 **startAt** | **optional.Int64**| Filter trades by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter trades by end time using a UNIX timestamp | 
 **tape** | **optional.String**| Filter by tape | 
 **page** | **optional.Int64**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int64**| Number of items displayed per page in a paginated result | [default to 10]
 **order** | **optional.String**| Sorting order of the output series | [default to DESC]

### Return type

[**[]SipTradesItem**](SipTradesItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorAtr**
> []TechnicalIndicatorResponseAtr GetTechnicalIndicatorAtr(ctx, dataset, ticker, interval, optional)
ATR Technical indicators

The Average True Range (ATR) is a volatility indicator that measures the average range of price movement over a specified period, helping traders assess market volatility.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorAtrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorAtrOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **timePeriod** | **optional.Int64**| Number of periods to average over. | [default to 14]

### Return type

[**[]TechnicalIndicatorResponseAtr**](TechnicalIndicatorResponseAtr.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorBBands**
> []TechnicalIndicatorResponseBBands GetTechnicalIndicatorBBands(ctx, dataset, ticker, interval, optional)
Overlap Studies

Bollinger Bands (BBANDS) are volatility bands placed above and below a moving average, measuring price volatility and helping traders identify potential overbought or oversold conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorBBandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorBBandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **seriesType** | **optional.String**| Specifies the price data type on which technical indicator is calculated | [default to close]
 **timePeriod** | **optional.Int64**| Number of periods to average over. | [default to 20]
 **sd** | **optional.Float64**| Number of standard deviations | [default to 2.0]
 **maType** | **optional.String**| The type of moving average used, such as SMA or EMA | [default to SMA]

### Return type

[**[]TechnicalIndicatorResponseBBands**](TechnicalIndicatorResponseBBands.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorIchimoku**
> []TechnicalIndicatorResponseIchimoku GetTechnicalIndicatorIchimoku(ctx, dataset, ticker, interval, optional)
Overlap Studies

The Ichimoku Cloud (ICHIMOKU) is a comprehensive trend-following indicator that combines multiple moving averages and support/resistance levels to help traders identify potential entry and exit points, trend direction, and momentum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorIchimokuOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorIchimokuOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **conversionLinePeriod** | **optional.Int64**| The time period used for generating the conversation line | [default to 9]
 **baseLinePeriod** | **optional.Int64**| The time period used for generating the base line | [default to 26]
 **leadingSpanBPeriod** | **optional.Int64**| The time period used for generating the leading span B line | [default to 52]
 **laggingSpanPeriod** | **optional.Int64**| The time period used for generating the lagging span line | [default to 26]
 **includeAheadSpanPeriod** | **optional.Bool**| Indicates whether to include ahead span period | [default to true]

### Return type

[**[]TechnicalIndicatorResponseIchimoku**](TechnicalIndicatorResponseIchimoku.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorMa**
> []TechnicalIndicatorResponseMa GetTechnicalIndicatorMa(ctx, dataset, ticker, interval, optional)
Overlap Studies

The Moving Average (MA) is a smoothing indicator that calculates the average price of a security over a specified period, helping traders identify trends and potential support or resistance levels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorMaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorMaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **seriesType** | **optional.String**| Specifies the price data type on which technical indicator is calculated | [default to close]
 **timePeriod** | **optional.Int64**| Number of periods to average over. | [default to 9]
 **maType** | **optional.String**| The type of moving average used, such as SMA or EMA | [default to SMA]

### Return type

[**[]TechnicalIndicatorResponseMa**](TechnicalIndicatorResponseMa.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorMacd**
> []TechnicalIndicatorResponseMacd GetTechnicalIndicatorMacd(ctx, dataset, ticker, interval, optional)
Momentum Indicators

The Moving Average Convergence Divergence (MACD) is a momentum indicator that measures the difference between two moving averages, with a signal line used to identify potential trend reversals and trading opportunities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorMacdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorMacdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **seriesType** | **optional.String**| Specifies the price data type on which technical indicator is calculated | [default to close]
 **fastPeriod** | **optional.Int64**| Number of periods for fast moving average | [default to 12]
 **slowPeriod** | **optional.Int64**| Number of periods for slow moving average | [default to 26]
 **signalPeriod** | **optional.Int64**| The time period used for generating the signal line | [default to 9]

### Return type

[**[]TechnicalIndicatorResponseMacd**](TechnicalIndicatorResponseMacd.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorObv**
> []TechnicalIndicatorResponseObv GetTechnicalIndicatorObv(ctx, dataset, ticker, interval, optional)
Volume Indicators

The On Balance Volume (OBV) indicator is a cumulative volume-based tool used to measure buying and selling pressure, helping traders identify potential price trends and reversals.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorObvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorObvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **seriesType** | **optional.String**| Specifies the price data type on which technical indicator is calculated | [default to close]

### Return type

[**[]TechnicalIndicatorResponseObv**](TechnicalIndicatorResponseObv.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorRsi**
> []TechnicalIndicatorResponseRsi GetTechnicalIndicatorRsi(ctx, dataset, ticker, interval, optional)
Momentum Indicators

The Relative Strength Index (RSI) is a momentum oscillator that measures the speed and change of price movements, helping traders identify potential overbought or oversold conditions and trend reversals.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorRsiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorRsiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **seriesType** | **optional.String**| Specifies the price data type on which technical indicator is calculated | [default to close]
 **timePeriod** | **optional.Int64**| Number of periods to average over | [default to 14]

### Return type

[**[]TechnicalIndicatorResponseRsi**](TechnicalIndicatorResponseRsi.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorSar**
> []TechnicalIndicatorResponseSar GetTechnicalIndicatorSar(ctx, dataset, ticker, interval, optional)
Overlap Studies

The Parabolic SAR (SAR) is a trend-following indicator that calculates potential support and resistance levels based on a security's price and time, helping traders identify potential entry and exit points.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorSarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorSarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **acceleration** | **optional.Float64**| Initial acceleration factor | [default to 0.02]
 **maximum** | **optional.Float64**| Maximum value for the acceleration factor | [default to 0.2]

### Return type

[**[]TechnicalIndicatorResponseSar**](TechnicalIndicatorResponseSar.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTechnicalIndicatorStoch**
> []TechnicalIndicatorResponseStoch GetTechnicalIndicatorStoch(ctx, dataset, ticker, interval, optional)
Momentum Indicators

The Stochastic Oscillator (STOCH) is a momentum indicator that compares a security's closing price to its price range over a specified period, helping traders identify potential overbought or oversold conditions and trend reversals.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
  **interval** | **string**| Interval between two consecutive points in time series | 
 **optional** | ***DataApiGetTechnicalIndicatorStochOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTechnicalIndicatorStochOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter output by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter output by end time using a UNIX timestamp | 
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 30]
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **prepost** | **optional.Bool**| Indicates whether data should be included for extended hours of trading | [default to false]
 **adjust** | **optional.String**| Apply adjusting for data (all, splits, dividends, none) | [default to all]
 **fastKPeriod** | **optional.Int64**| The time period for the fast %K line in the Stochastic Oscillator | [default to 14]
 **slowKPeriod** | **optional.Int64**| The time period for the slow %K line in the Stochastic Oscillator | [default to 1]
 **slowDPeriod** | **optional.Int64**| The time period for the slow %D line in the Stochastic Oscillator | [default to 3]
 **slowKmaType** | **optional.String**| The type of slow %K Moving Average used | [default to SMA]
 **slowDmaType** | **optional.String**| The type of slow Displaced Moving Average used | [default to SMA]

### Return type

[**[]TechnicalIndicatorResponseStoch**](TechnicalIndicatorResponseStoch.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTickerSnapshot**
> InlineResponse200 GetTickerSnapshot(ctx, dataset, ticker, optional)
Ticker snapshot

This endpoint returns a combination of different data points, such as daily performance, last quote, last trade, minute data, and previous day performance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetTickerSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTickerSnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **market** | **optional.String**| Filter by market | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrades**
> []TradesItem GetTrades(ctx, dataset, ticker, optional)
Trades

Returns general information on executed trades.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataset** | **string**| Filter by Finazon&#x27;s dataset code | 
  **ticker** | **string**| Filter by ticker symbol | 
 **optional** | ***DataApiGetTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataApiGetTradesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cqs** | **optional.String**| Filter by cqs symbol | 
 **cik** | **optional.String**| Filter by cik code | 
 **cusip** | **optional.String**| Filter by cusip code | 
 **isin** | **optional.String**| Filter by isin code | 
 **compositeFigi** | **optional.String**| Filter by composite figi code | 
 **shareFigi** | **optional.String**| Filter by share class figi code | 
 **lei** | **optional.String**| Filter by lei code | 
 **country** | **optional.String**| Filter by ISO 3166 alpha-2 code | 
 **startAt** | **optional.Int64**| Filter trades by start time using a UNIX timestamp | 
 **endAt** | **optional.Int64**| Filter trades by end time using a UNIX timestamp | 
 **order** | **optional.String**| Sorting order of the output series | [default to desc]
 **page** | **optional.Int32**| Specific page of a paginated result to be displayed | [default to 0]
 **pageSize** | **optional.Int32**| Number of items displayed per page in a paginated result | [default to 1000]

### Return type

[**[]TradesItem**](TradesItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


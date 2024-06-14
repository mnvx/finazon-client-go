# {{classname}}

All URIs are relative to *https://api.finazon.io/v1.2/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTechnicalIndicatorAtr**](TechnicalIndicatorApi.md#GetTechnicalIndicatorAtr) | **Get** /time_series/atr | ATR Technical indicators
[**GetTechnicalIndicatorBBands**](TechnicalIndicatorApi.md#GetTechnicalIndicatorBBands) | **Get** /time_series/bbands | Overlap Studies
[**GetTechnicalIndicatorIchimoku**](TechnicalIndicatorApi.md#GetTechnicalIndicatorIchimoku) | **Get** /time_series/ichimoku | Overlap Studies
[**GetTechnicalIndicatorMa**](TechnicalIndicatorApi.md#GetTechnicalIndicatorMa) | **Get** /time_series/ma | Overlap Studies
[**GetTechnicalIndicatorMacd**](TechnicalIndicatorApi.md#GetTechnicalIndicatorMacd) | **Get** /time_series/macd | Momentum Indicators
[**GetTechnicalIndicatorObv**](TechnicalIndicatorApi.md#GetTechnicalIndicatorObv) | **Get** /time_series/obv | Volume Indicators
[**GetTechnicalIndicatorRsi**](TechnicalIndicatorApi.md#GetTechnicalIndicatorRsi) | **Get** /time_series/rsi | Momentum Indicators
[**GetTechnicalIndicatorSar**](TechnicalIndicatorApi.md#GetTechnicalIndicatorSar) | **Get** /time_series/sar | Overlap Studies
[**GetTechnicalIndicatorStoch**](TechnicalIndicatorApi.md#GetTechnicalIndicatorStoch) | **Get** /time_series/stoch | Momentum Indicators

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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorAtrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorAtrOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorBBandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorBBandsOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorIchimokuOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorIchimokuOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorMaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorMaOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorMacdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorMacdOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorObvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorObvOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorRsiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorRsiOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorSarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorSarOpts struct
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
 **optional** | ***TechnicalIndicatorApiGetTechnicalIndicatorStochOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnicalIndicatorApiGetTechnicalIndicatorStochOpts struct
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


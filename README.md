# Go API client for swagger

## API reference  Finazon is a comprehensive financial data marketplace that enables developers to effortlessly integrate a wide variety of global datasets, including stocks, ETFs, cryptocurrencies, and more, all with fully customizable parameters.  The Finazon API is built around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, featuring resource-oriented URLs with predictable behavior. The API accepts [form-encoded](https://en.wikipedia.org/wiki/POST_(HTTP)#Use_for_submitting_web_forms) request bodies, returns JSON-encoded responses, and utilizes standard HTTP response codes, authentication methods, and verbs.  The Finazon API doesn't support bulk updates. You can work on only one instrument per request.  ## Authentification  To authenticate requests, the Finazon API requires [API keys](https://finazon.io/account/developers/api-keys). You can obtain, view, and manage your API keys through the [Finazon Dashboard](https://finazon.io/account/home).  Your API keys hold significant privileges, so ensure their security by not sharing your secret API keys in publicly accessible areas, such as GitHub repositories, client-side code, or any other public platforms.  All API requests must be made over [HTTPS](http://en.wikipedia.org/wiki/HTTP_Secure). Calls over plain HTTP will fail, as will API requests without authentication.  Once you have your API key, include it in the parameters as follows:  ```bash https://api.finazon.io/latest?apikey=YOUR_API_KEY ```  Alternatively, pass it as a request header:  ```bash Authorization: apikey YOUR_API_KEY ```  ## Versioning  Whenever [backwards-incompatible](https://support.finazon.io/en/articles/7792859-api-upgrades#h_1e014217bf) changes are introduced to the API, a new dated version is released. Consult our [API upgrades guide](https://support.finazon.io/en/articles/7792859-api-upgrades) for more information on backwards compatibility, and view our API changelog for all API updates.  To always use the most up-to-date version, specify it as `/latest`:  ```bash https://api.finazon.io/latest ```  To access the most recent version of `v1.*`, use the following:  ```bash https://api.finazon.io/v1  ```  Or, to retrieve a specific version, call:  ```bash  https://api.finazon.io/v1.0 ```  Finazon will provide advance notice before deprecating older API versions, giving developers ample time to migrate to the updated version.  ## Endpoints structure  The Finazon API adheres to a consistent and structured pattern for its endpoints. The base URL for all requests is:  ```bash https://api.finazon.io/ ```  API endpoints are organized by resource types, including universal resources accessible across all publishers and publisher-specific resources. For example, the `/time_series` endpoint is compatible with all publishers that support this data format. Such responses will be standardized across all datasets, facilitating rapid integration of new markets into your applications.  ```bash https://api.finazon.io/latest/{{resource}} https://api.finazon.io/latest/time_series ```  Additionally, datasets may contain unique data exclusive to that dataset. In such cases, you might want to call a separate endpoint specifying the publisher to gather more data. For instance, the [Binance](https://finazon.io/dataset/binance) dataset time series can be requested as:  ```bash  https://api.finazon.io/latest/{{publisher}}/{{resource}} https://api.finazon.io/latest/binance/time_series ```  ## Parameters  Each API request has its own set of required and optional parameters. Parameters should be separated by an ampersand. Parameter names are case-sensitive, while parameter values are not. Each API request has its own set of required and optional parameters. Parameters should be separated by an ampersand. Parameter names and parameter values are case-sensitive  ```bash https://api.finazon.io/latest/time_series?dataset=sip_non_pro&ticker=AAPL&interval=1m&apikey= ```  ### Pagination  All API resources supporting bulk fetches are retrieved via \"list\" API methods. For example, you can list time series, list trades, and list quotes. These list API methods share a common structure, accepting at least these five parameters: `page`, `page_size`, `order`, `start_at`, and `end_at`.  The response of a list API method represents a single page in a reverse chronological stream of objects. If you do not specify `start_at` or `end_at`, you will receive the first page of this list, containing the newest objects. By default, you will receive 10 objects if you do not specify an alternative value for `page_size`. You can specify `start_at` equal to the T (timestamp) value of an item to retrieve the page of older objects occurring immediately after the specified timestamp in the reverse chronological stream. Similarly, you can specify `end_at` to receive a page of newer objects occurring immediately before the named object in the stream. You can use one of `start_at` or `end_at` or both. Objects in a page always appear in reverse chronological order, unless `order` is specified.  ## Errors  Finazon employs standard HTTP response codes to signify the success or failure of an API request. Generally, the response codes can be interpreted as follows:  `2xx` range codes indicate a successful request.  `4xx` range codes signify an error resulting from the provided information (e.g., invalid API key, API rate limit exceeded, etc.).  `5xx` range codes represent errors originating from Finazon's servers (these are rare occurrences).  For all `4xx`errors that can be addressed programmatically (e.g., endpoint not found), an error message is included to succinctly explain the reported issue. This allows developers to quickly identify and resolve errors in their API requests.   status | code | message | --------|:-----|:--------|  400 |  INVALID_PARAMETER | The **{parameter_name}** parameter is missing or invalid. |  400 |  INVALID_DATE_RANGE | The requested date range is invalid or unsupported. |  400 |  UNSUPPORTED_MARKET | The requested market or exchange is not supported by the API. Please check the supported markets and try again. |  400 |  INVALID_TICKER | The provided ticker is invalid or unsupported. |  401 |  UNAUTHORIZED_ACCESS | You are not authorized to access the requested endpoint or you have insufficient permissions. |  404 |  ENDPOINT_NOT_FOUND | The requested endpoint **{endpoint_name}** does not exist or could not be found. |  429 |  API_RATE_LIMIT_EXCEEDED | You have exceeded the allowed number of API calls within the minute. Please wait and try again later. |  401 |  INVALID_API_KEY | The provided API key is invalid or has expired. Please check your API key and try again |  408 |  REQUEST_TIMEOUT | The request took too long to complete and timed out. Please try again later or reduce the complexity of your query. |  503 |  DATA_UNAVAILABLE | The requested data is temporarily unavailable or not supported. Please try again later or check the availability of the data. |  500 |  INTERNAL_SERVER_ERROR | An error occurred on the server-side while processing the request. Please try again later. If the issue persists, contact support. |

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1.2
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.finazon.io/v1.2/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DataApi* | [**GetBenzingaDividendsCalendar**](docs/DataApi.md#getbenzingadividendscalendar) | **Get** /benzinga/dividends_calendar | Dividends calendar
*DataApi* | [**GetBenzingaEarningsCalendar**](docs/DataApi.md#getbenzingaearningscalendar) | **Get** /benzinga/earnings_calendar | Earnings calendar
*DataApi* | [**GetBenzingaIPO**](docs/DataApi.md#getbenzingaipo) | **Get** /benzinga/ipo | IPO data
*DataApi* | [**GetBenzingaNews**](docs/DataApi.md#getbenzinganews) | **Get** /benzinga/news | News articles
*DataApi* | [**GetBinanceQuotes**](docs/DataApi.md#getbinancequotes) | **Get** /binance/time_series | Time series
*DataApi* | [**GetCryptoQuotes**](docs/DataApi.md#getcryptoquotes) | **Get** /crypto/time_series | Time series
*DataApi* | [**GetFilings**](docs/DataApi.md#getfilings) | **Get** /sec/archive | Filings
*DataApi* | [**GetForexQuotes**](docs/DataApi.md#getforexquotes) | **Get** /forex/time_series | Time series
*DataApi* | [**GetPriceBinance**](docs/DataApi.md#getpricebinance) | **Get** /binance/price | Price
*DataApi* | [**GetPriceCommon**](docs/DataApi.md#getpricecommon) | **Get** /price | Price
*DataApi* | [**GetPriceCrypto**](docs/DataApi.md#getpricecrypto) | **Get** /crypto/price | Price
*DataApi* | [**GetPriceForex**](docs/DataApi.md#getpriceforex) | **Get** /forex/price | Forex price
*DataApi* | [**GetPriceSip**](docs/DataApi.md#getpricesip) | **Get** /sip_non_pro/price | Price
*DataApi* | [**GetPriceSip_0**](docs/DataApi.md#getpricesip_0) | **Get** /sip_pro/price | Price
*DataApi* | [**GetPriceUsStocks**](docs/DataApi.md#getpriceusstocks) | **Get** /us_stocks_essential/price | Price
*DataApi* | [**GetQuotes**](docs/DataApi.md#getquotes) | **Get** /time_series | Time series
*DataApi* | [**GetSipTrades**](docs/DataApi.md#getsiptrades) | **Get** /sip/trades | SIP trades
*DataApi* | [**GetTechnicalIndicatorAtr**](docs/DataApi.md#gettechnicalindicatoratr) | **Get** /time_series/atr | ATR Technical indicators
*DataApi* | [**GetTechnicalIndicatorBBands**](docs/DataApi.md#gettechnicalindicatorbbands) | **Get** /time_series/bbands | Overlap Studies
*DataApi* | [**GetTechnicalIndicatorIchimoku**](docs/DataApi.md#gettechnicalindicatorichimoku) | **Get** /time_series/ichimoku | Overlap Studies
*DataApi* | [**GetTechnicalIndicatorMa**](docs/DataApi.md#gettechnicalindicatorma) | **Get** /time_series/ma | Overlap Studies
*DataApi* | [**GetTechnicalIndicatorMacd**](docs/DataApi.md#gettechnicalindicatormacd) | **Get** /time_series/macd | Momentum Indicators
*DataApi* | [**GetTechnicalIndicatorObv**](docs/DataApi.md#gettechnicalindicatorobv) | **Get** /time_series/obv | Volume Indicators
*DataApi* | [**GetTechnicalIndicatorRsi**](docs/DataApi.md#gettechnicalindicatorrsi) | **Get** /time_series/rsi | Momentum Indicators
*DataApi* | [**GetTechnicalIndicatorSar**](docs/DataApi.md#gettechnicalindicatorsar) | **Get** /time_series/sar | Overlap Studies
*DataApi* | [**GetTechnicalIndicatorStoch**](docs/DataApi.md#gettechnicalindicatorstoch) | **Get** /time_series/stoch | Momentum Indicators
*DataApi* | [**GetTickerSnapshot**](docs/DataApi.md#gettickersnapshot) | **Get** /ticker/snapshot | Ticker snapshot
*DataApi* | [**GetTrades**](docs/DataApi.md#gettrades) | **Get** /trades | Trades
*DefaultApi* | [**Reference**](docs/DefaultApi.md#reference) | **Get** /my_datasets | My Datasets
*ReferenceApi* | [**GetApiUsage**](docs/ReferenceApi.md#getapiusage) | **Get** /api_usage | Api usage
*ReferenceApi* | [**GetDatasets**](docs/ReferenceApi.md#getdatasets) | **Get** /datasets | List of Finazon datasets
*ReferenceApi* | [**GetExchangesCrypto**](docs/ReferenceApi.md#getexchangescrypto) | **Get** /markets/crypto | List of crypto markets
*ReferenceApi* | [**GetExchangesStocks**](docs/ReferenceApi.md#getexchangesstocks) | **Get** /markets/stocks | List of stock markets
*ReferenceApi* | [**GetMarketCenter**](docs/ReferenceApi.md#getmarketcenter) | **Get** /sip/market_center | List of market centers
*ReferenceApi* | [**GetPublishers**](docs/ReferenceApi.md#getpublishers) | **Get** /publishers | List of Finazon publishers
*ReferenceApi* | [**GetSymbolsCrypto**](docs/ReferenceApi.md#getsymbolscrypto) | **Get** /tickers/crypto | List of crypto pairs
*ReferenceApi* | [**GetSymbolsForex**](docs/ReferenceApi.md#getsymbolsforex) | **Get** /tickers/forex | List of forex ticker symbols
*ReferenceApi* | [**GetSymbolsStocks**](docs/ReferenceApi.md#getsymbolsstocks) | **Get** /tickers/stocks | List of stock ticker symbols
*ReferenceApi* | [**GetSymbolsUSStocks**](docs/ReferenceApi.md#getsymbolsusstocks) | **Get** /tickers/us_stocks | List of US stock ticker symbols
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorAtr**](docs/TechnicalIndicatorApi.md#gettechnicalindicatoratr) | **Get** /time_series/atr | ATR Technical indicators
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorBBands**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorbbands) | **Get** /time_series/bbands | Overlap Studies
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorIchimoku**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorichimoku) | **Get** /time_series/ichimoku | Overlap Studies
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorMa**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorma) | **Get** /time_series/ma | Overlap Studies
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorMacd**](docs/TechnicalIndicatorApi.md#gettechnicalindicatormacd) | **Get** /time_series/macd | Momentum Indicators
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorObv**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorobv) | **Get** /time_series/obv | Volume Indicators
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorRsi**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorrsi) | **Get** /time_series/rsi | Momentum Indicators
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorSar**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorsar) | **Get** /time_series/sar | Overlap Studies
*TechnicalIndicatorApi* | [**GetTechnicalIndicatorStoch**](docs/TechnicalIndicatorApi.md#gettechnicalindicatorstoch) | **Get** /time_series/stoch | Momentum Indicators

## Documentation For Models

 - [ApiCalls](docs/ApiCalls.md)
 - [ApiUsageItem](docs/ApiUsageItem.md)
 - [ApiUsageResponseBody](docs/ApiUsageResponseBody.md)
 - [BaseResponseBody](docs/BaseResponseBody.md)
 - [BenzingaDividendsCalendarResponseBody](docs/BenzingaDividendsCalendarResponseBody.md)
 - [BenzingaEarningsCalendarResponseBody](docs/BenzingaEarningsCalendarResponseBody.md)
 - [BenzingaIpoResponseBody](docs/BenzingaIpoResponseBody.md)
 - [BenzingaNewsResponseBody](docs/BenzingaNewsResponseBody.md)
 - [DatasetDetailItem](docs/DatasetDetailItem.md)
 - [DatasetItem](docs/DatasetItem.md)
 - [DatasetsResponseBody](docs/DatasetsResponseBody.md)
 - [DividendsCalendarItem](docs/DividendsCalendarItem.md)
 - [EarningsCalendarItem](docs/EarningsCalendarItem.md)
 - [EarningsCalendarItemEps](docs/EarningsCalendarItemEps.md)
 - [EarningsCalendarItemRevenue](docs/EarningsCalendarItemRevenue.md)
 - [ExchangeCryptoItem](docs/ExchangeCryptoItem.md)
 - [ExchangeStocksItem](docs/ExchangeStocksItem.md)
 - [ExchangesCryptoResponseBody](docs/ExchangesCryptoResponseBody.md)
 - [ExchangesStocksResponseBody](docs/ExchangesStocksResponseBody.md)
 - [Filing](docs/Filing.md)
 - [FilingFile](docs/FilingFile.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [IpoItem](docs/IpoItem.md)
 - [MarketCenterItem](docs/MarketCenterItem.md)
 - [MarketCenterResponseBody](docs/MarketCenterResponseBody.md)
 - [Meta](docs/Meta.md)
 - [MyDatasetItem](docs/MyDatasetItem.md)
 - [NewsItem](docs/NewsItem.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PriceResponseBody](docs/PriceResponseBody.md)
 - [PublisherItem](docs/PublisherItem.md)
 - [PublishersResponseBody](docs/PublishersResponseBody.md)
 - [QuoteBinanceItem](docs/QuoteBinanceItem.md)
 - [QuoteForexItem](docs/QuoteForexItem.md)
 - [QuoteItem](docs/QuoteItem.md)
 - [SipTradesItem](docs/SipTradesItem.md)
 - [SymbolCrypto](docs/SymbolCrypto.md)
 - [SymbolForex](docs/SymbolForex.md)
 - [SymbolStocks](docs/SymbolStocks.md)
 - [SymbolUsStocks](docs/SymbolUsStocks.md)
 - [SymbolsCryptoResponseBody](docs/SymbolsCryptoResponseBody.md)
 - [SymbolsForexResponseBody](docs/SymbolsForexResponseBody.md)
 - [SymbolsStocksResponseBody](docs/SymbolsStocksResponseBody.md)
 - [SymbolsUsStocksResponseBody](docs/SymbolsUsStocksResponseBody.md)
 - [TechnicalIndicatorResponseAtr](docs/TechnicalIndicatorResponseAtr.md)
 - [TechnicalIndicatorResponseBBands](docs/TechnicalIndicatorResponseBBands.md)
 - [TechnicalIndicatorResponseData](docs/TechnicalIndicatorResponseData.md)
 - [TechnicalIndicatorResponseIchimoku](docs/TechnicalIndicatorResponseIchimoku.md)
 - [TechnicalIndicatorResponseMa](docs/TechnicalIndicatorResponseMa.md)
 - [TechnicalIndicatorResponseMacd](docs/TechnicalIndicatorResponseMacd.md)
 - [TechnicalIndicatorResponseObv](docs/TechnicalIndicatorResponseObv.md)
 - [TechnicalIndicatorResponseRsi](docs/TechnicalIndicatorResponseRsi.md)
 - [TechnicalIndicatorResponseSar](docs/TechnicalIndicatorResponseSar.md)
 - [TechnicalIndicatorResponseStoch](docs/TechnicalIndicatorResponseStoch.md)
 - [TickerSnapshotChange](docs/TickerSnapshotChange.md)
 - [TickerSnapshotLastDay](docs/TickerSnapshotLastDay.md)
 - [TickerSnapshotLastFiftyTwoWeek](docs/TickerSnapshotLastFiftyTwoWeek.md)
 - [TickerSnapshotLastMonth](docs/TickerSnapshotLastMonth.md)
 - [TickerSnapshotLastTrade](docs/TickerSnapshotLastTrade.md)
 - [TickerSnapshotPreviousDay](docs/TickerSnapshotPreviousDay.md)
 - [TradesItem](docs/TradesItem.md)

## Documentation For Authorization

## api_key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



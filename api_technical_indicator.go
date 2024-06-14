
/*
 * Finazon API
 *
 * ## API reference  Finazon is a comprehensive financial data marketplace that enables developers to effortlessly integrate a wide variety of global datasets, including stocks, ETFs, cryptocurrencies, and more, all with fully customizable parameters.  The Finazon API is built around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, featuring resource-oriented URLs with predictable behavior. The API accepts [form-encoded](https://en.wikipedia.org/wiki/POST_(HTTP)#Use_for_submitting_web_forms) request bodies, returns JSON-encoded responses, and utilizes standard HTTP response codes, authentication methods, and verbs.  The Finazon API doesn't support bulk updates. You can work on only one instrument per request.  ## Authentification  To authenticate requests, the Finazon API requires [API keys](https://finazon.io/account/developers/api-keys). You can obtain, view, and manage your API keys through the [Finazon Dashboard](https://finazon.io/account/home).  Your API keys hold significant privileges, so ensure their security by not sharing your secret API keys in publicly accessible areas, such as GitHub repositories, client-side code, or any other public platforms.  All API requests must be made over [HTTPS](http://en.wikipedia.org/wiki/HTTP_Secure). Calls over plain HTTP will fail, as will API requests without authentication.  Once you have your API key, include it in the parameters as follows:  ```bash https://api.finazon.io/latest?apikey=YOUR_API_KEY ```  Alternatively, pass it as a request header:  ```bash Authorization: apikey YOUR_API_KEY ```  ## Versioning  Whenever [backwards-incompatible](https://support.finazon.io/en/articles/7792859-api-upgrades#h_1e014217bf) changes are introduced to the API, a new dated version is released. Consult our [API upgrades guide](https://support.finazon.io/en/articles/7792859-api-upgrades) for more information on backwards compatibility, and view our API changelog for all API updates.  To always use the most up-to-date version, specify it as `/latest`:  ```bash https://api.finazon.io/latest ```  To access the most recent version of `v1.*`, use the following:  ```bash https://api.finazon.io/v1  ```  Or, to retrieve a specific version, call:  ```bash  https://api.finazon.io/v1.0 ```  Finazon will provide advance notice before deprecating older API versions, giving developers ample time to migrate to the updated version.  ## Endpoints structure  The Finazon API adheres to a consistent and structured pattern for its endpoints. The base URL for all requests is:  ```bash https://api.finazon.io/ ```  API endpoints are organized by resource types, including universal resources accessible across all publishers and publisher-specific resources. For example, the `/time_series` endpoint is compatible with all publishers that support this data format. Such responses will be standardized across all datasets, facilitating rapid integration of new markets into your applications.  ```bash https://api.finazon.io/latest/{{resource}} https://api.finazon.io/latest/time_series ```  Additionally, datasets may contain unique data exclusive to that dataset. In such cases, you might want to call a separate endpoint specifying the publisher to gather more data. For instance, the [Binance](https://finazon.io/dataset/binance) dataset time series can be requested as:  ```bash  https://api.finazon.io/latest/{{publisher}}/{{resource}} https://api.finazon.io/latest/binance/time_series ```  ## Parameters  Each API request has its own set of required and optional parameters. Parameters should be separated by an ampersand. Parameter names are case-sensitive, while parameter values are not. Each API request has its own set of required and optional parameters. Parameters should be separated by an ampersand. Parameter names and parameter values are case-sensitive  ```bash https://api.finazon.io/latest/time_series?dataset=sip_non_pro&ticker=AAPL&interval=1m&apikey= ```  ### Pagination  All API resources supporting bulk fetches are retrieved via \"list\" API methods. For example, you can list time series, list trades, and list quotes. These list API methods share a common structure, accepting at least these five parameters: `page`, `page_size`, `order`, `start_at`, and `end_at`.  The response of a list API method represents a single page in a reverse chronological stream of objects. If you do not specify `start_at` or `end_at`, you will receive the first page of this list, containing the newest objects. By default, you will receive 10 objects if you do not specify an alternative value for `page_size`. You can specify `start_at` equal to the T (timestamp) value of an item to retrieve the page of older objects occurring immediately after the specified timestamp in the reverse chronological stream. Similarly, you can specify `end_at` to receive a page of newer objects occurring immediately before the named object in the stream. You can use one of `start_at` or `end_at` or both. Objects in a page always appear in reverse chronological order, unless `order` is specified.  ## Errors  Finazon employs standard HTTP response codes to signify the success or failure of an API request. Generally, the response codes can be interpreted as follows:  `2xx` range codes indicate a successful request.  `4xx` range codes signify an error resulting from the provided information (e.g., invalid API key, API rate limit exceeded, etc.).  `5xx` range codes represent errors originating from Finazon's servers (these are rare occurrences).  For all `4xx`errors that can be addressed programmatically (e.g., endpoint not found), an error message is included to succinctly explain the reported issue. This allows developers to quickly identify and resolve errors in their API requests.   status | code | message | --------|:-----|:--------|  400 |  INVALID_PARAMETER | The **{parameter_name}** parameter is missing or invalid. |  400 |  INVALID_DATE_RANGE | The requested date range is invalid or unsupported. |  400 |  UNSUPPORTED_MARKET | The requested market or exchange is not supported by the API. Please check the supported markets and try again. |  400 |  INVALID_TICKER | The provided ticker is invalid or unsupported. |  401 |  UNAUTHORIZED_ACCESS | You are not authorized to access the requested endpoint or you have insufficient permissions. |  404 |  ENDPOINT_NOT_FOUND | The requested endpoint **{endpoint_name}** does not exist or could not be found. |  429 |  API_RATE_LIMIT_EXCEEDED | You have exceeded the allowed number of API calls within the minute. Please wait and try again later. |  401 |  INVALID_API_KEY | The provided API key is invalid or has expired. Please check your API key and try again |  408 |  REQUEST_TIMEOUT | The request took too long to complete and timed out. Please try again later or reduce the complexity of your query. |  503 |  DATA_UNAVAILABLE | The requested data is temporarily unavailable or not supported. Please try again later or check the availability of the data. |  500 |  INTERNAL_SERVER_ERROR | An error occurred on the server-side while processing the request. Please try again later. If the issue persists, contact support. |
 *
 * API version: v1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type TechnicalIndicatorApiService service
/*
TechnicalIndicatorApiService ATR Technical indicators
The Average True Range (ATR) is a volatility indicator that measures the average range of price movement over a specified period, helping traders assess market volatility.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorAtrOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "TimePeriod" (optional.Int64) -  Number of periods to average over.
@return []TechnicalIndicatorResponseAtr
*/

type TechnicalIndicatorApiGetTechnicalIndicatorAtrOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    TimePeriod optional.Int64
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorAtr(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorAtrOpts) ([]TechnicalIndicatorResponseAtr, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseAtr
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/atr"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimePeriod.IsSet() {
		localVarQueryParams.Add("time_period", parameterToString(localVarOptionals.TimePeriod.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseAtr
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Overlap Studies
Bollinger Bands (BBANDS) are volatility bands placed above and below a moving average, measuring price volatility and helping traders identify potential overbought or oversold conditions.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorBBandsOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "SeriesType" (optional.String) -  Specifies the price data type on which technical indicator is calculated
     * @param "TimePeriod" (optional.Int64) -  Number of periods to average over.
     * @param "Sd" (optional.Float64) -  Number of standard deviations
     * @param "MaType" (optional.String) -  The type of moving average used, such as SMA or EMA
@return []TechnicalIndicatorResponseBBands
*/

type TechnicalIndicatorApiGetTechnicalIndicatorBBandsOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    SeriesType optional.String
    TimePeriod optional.Int64
    Sd optional.Float64
    MaType optional.String
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorBBands(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorBBandsOpts) ([]TechnicalIndicatorResponseBBands, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseBBands
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/bbands"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeriesType.IsSet() {
		localVarQueryParams.Add("series_type", parameterToString(localVarOptionals.SeriesType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimePeriod.IsSet() {
		localVarQueryParams.Add("time_period", parameterToString(localVarOptionals.TimePeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sd.IsSet() {
		localVarQueryParams.Add("sd", parameterToString(localVarOptionals.Sd.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaType.IsSet() {
		localVarQueryParams.Add("ma_type", parameterToString(localVarOptionals.MaType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseBBands
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Overlap Studies
The Ichimoku Cloud (ICHIMOKU) is a comprehensive trend-following indicator that combines multiple moving averages and support/resistance levels to help traders identify potential entry and exit points, trend direction, and momentum.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorIchimokuOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "ConversionLinePeriod" (optional.Int64) -  The time period used for generating the conversation line
     * @param "BaseLinePeriod" (optional.Int64) -  The time period used for generating the base line
     * @param "LeadingSpanBPeriod" (optional.Int64) -  The time period used for generating the leading span B line
     * @param "LaggingSpanPeriod" (optional.Int64) -  The time period used for generating the lagging span line
     * @param "IncludeAheadSpanPeriod" (optional.Bool) -  Indicates whether to include ahead span period
@return []TechnicalIndicatorResponseIchimoku
*/

type TechnicalIndicatorApiGetTechnicalIndicatorIchimokuOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    ConversionLinePeriod optional.Int64
    BaseLinePeriod optional.Int64
    LeadingSpanBPeriod optional.Int64
    LaggingSpanPeriod optional.Int64
    IncludeAheadSpanPeriod optional.Bool
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorIchimoku(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorIchimokuOpts) ([]TechnicalIndicatorResponseIchimoku, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseIchimoku
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/ichimoku"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConversionLinePeriod.IsSet() {
		localVarQueryParams.Add("conversion_line_period", parameterToString(localVarOptionals.ConversionLinePeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseLinePeriod.IsSet() {
		localVarQueryParams.Add("base_line_period", parameterToString(localVarOptionals.BaseLinePeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LeadingSpanBPeriod.IsSet() {
		localVarQueryParams.Add("leading_span_b_period", parameterToString(localVarOptionals.LeadingSpanBPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LaggingSpanPeriod.IsSet() {
		localVarQueryParams.Add("lagging_span_period", parameterToString(localVarOptionals.LaggingSpanPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeAheadSpanPeriod.IsSet() {
		localVarQueryParams.Add("include_ahead_span_period", parameterToString(localVarOptionals.IncludeAheadSpanPeriod.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseIchimoku
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Overlap Studies
The Moving Average (MA) is a smoothing indicator that calculates the average price of a security over a specified period, helping traders identify trends and potential support or resistance levels.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorMaOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "SeriesType" (optional.String) -  Specifies the price data type on which technical indicator is calculated
     * @param "TimePeriod" (optional.Int64) -  Number of periods to average over.
     * @param "MaType" (optional.String) -  The type of moving average used, such as SMA or EMA
@return []TechnicalIndicatorResponseMa
*/

type TechnicalIndicatorApiGetTechnicalIndicatorMaOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    SeriesType optional.String
    TimePeriod optional.Int64
    MaType optional.String
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorMa(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorMaOpts) ([]TechnicalIndicatorResponseMa, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseMa
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/ma"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeriesType.IsSet() {
		localVarQueryParams.Add("series_type", parameterToString(localVarOptionals.SeriesType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimePeriod.IsSet() {
		localVarQueryParams.Add("time_period", parameterToString(localVarOptionals.TimePeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaType.IsSet() {
		localVarQueryParams.Add("ma_type", parameterToString(localVarOptionals.MaType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseMa
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Momentum Indicators
The Moving Average Convergence Divergence (MACD) is a momentum indicator that measures the difference between two moving averages, with a signal line used to identify potential trend reversals and trading opportunities.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorMacdOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "SeriesType" (optional.String) -  Specifies the price data type on which technical indicator is calculated
     * @param "FastPeriod" (optional.Int64) -  Number of periods for fast moving average
     * @param "SlowPeriod" (optional.Int64) -  Number of periods for slow moving average
     * @param "SignalPeriod" (optional.Int64) -  The time period used for generating the signal line
@return []TechnicalIndicatorResponseMacd
*/

type TechnicalIndicatorApiGetTechnicalIndicatorMacdOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    SeriesType optional.String
    FastPeriod optional.Int64
    SlowPeriod optional.Int64
    SignalPeriod optional.Int64
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorMacd(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorMacdOpts) ([]TechnicalIndicatorResponseMacd, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseMacd
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/macd"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeriesType.IsSet() {
		localVarQueryParams.Add("series_type", parameterToString(localVarOptionals.SeriesType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FastPeriod.IsSet() {
		localVarQueryParams.Add("fast_period", parameterToString(localVarOptionals.FastPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlowPeriod.IsSet() {
		localVarQueryParams.Add("slow_period", parameterToString(localVarOptionals.SlowPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SignalPeriod.IsSet() {
		localVarQueryParams.Add("signal_period", parameterToString(localVarOptionals.SignalPeriod.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseMacd
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Volume Indicators
The On Balance Volume (OBV) indicator is a cumulative volume-based tool used to measure buying and selling pressure, helping traders identify potential price trends and reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorObvOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "SeriesType" (optional.String) -  Specifies the price data type on which technical indicator is calculated
@return []TechnicalIndicatorResponseObv
*/

type TechnicalIndicatorApiGetTechnicalIndicatorObvOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    SeriesType optional.String
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorObv(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorObvOpts) ([]TechnicalIndicatorResponseObv, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseObv
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/obv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeriesType.IsSet() {
		localVarQueryParams.Add("series_type", parameterToString(localVarOptionals.SeriesType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseObv
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Momentum Indicators
The Relative Strength Index (RSI) is a momentum oscillator that measures the speed and change of price movements, helping traders identify potential overbought or oversold conditions and trend reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorRsiOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "SeriesType" (optional.String) -  Specifies the price data type on which technical indicator is calculated
     * @param "TimePeriod" (optional.Int64) -  Number of periods to average over
@return []TechnicalIndicatorResponseRsi
*/

type TechnicalIndicatorApiGetTechnicalIndicatorRsiOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    SeriesType optional.String
    TimePeriod optional.Int64
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorRsi(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorRsiOpts) ([]TechnicalIndicatorResponseRsi, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseRsi
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/rsi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeriesType.IsSet() {
		localVarQueryParams.Add("series_type", parameterToString(localVarOptionals.SeriesType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimePeriod.IsSet() {
		localVarQueryParams.Add("time_period", parameterToString(localVarOptionals.TimePeriod.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseRsi
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Overlap Studies
The Parabolic SAR (SAR) is a trend-following indicator that calculates potential support and resistance levels based on a security&#x27;s price and time, helping traders identify potential entry and exit points.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorSarOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "Acceleration" (optional.Float64) -  Initial acceleration factor
     * @param "Maximum" (optional.Float64) -  Maximum value for the acceleration factor
@return []TechnicalIndicatorResponseSar
*/

type TechnicalIndicatorApiGetTechnicalIndicatorSarOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    Acceleration optional.Float64
    Maximum optional.Float64
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorSar(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorSarOpts) ([]TechnicalIndicatorResponseSar, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseSar
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/sar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Acceleration.IsSet() {
		localVarQueryParams.Add("acceleration", parameterToString(localVarOptionals.Acceleration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Maximum.IsSet() {
		localVarQueryParams.Add("maximum", parameterToString(localVarOptionals.Maximum.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseSar
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
TechnicalIndicatorApiService Momentum Indicators
The Stochastic Oscillator (STOCH) is a momentum indicator that compares a security&#x27;s closing price to its price range over a specified period, helping traders identify potential overbought or oversold conditions and trend reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *TechnicalIndicatorApiGetTechnicalIndicatorStochOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Adjust" (optional.String) -  Apply adjusting for data (all, splits, dividends, none)
     * @param "FastKPeriod" (optional.Int64) -  The time period for the fast %K line in the Stochastic Oscillator
     * @param "SlowKPeriod" (optional.Int64) -  The time period for the slow %K line in the Stochastic Oscillator
     * @param "SlowDPeriod" (optional.Int64) -  The time period for the slow %D line in the Stochastic Oscillator
     * @param "SlowKmaType" (optional.String) -  The type of slow %K Moving Average used
     * @param "SlowDmaType" (optional.String) -  The type of slow Displaced Moving Average used
@return []TechnicalIndicatorResponseStoch
*/

type TechnicalIndicatorApiGetTechnicalIndicatorStochOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Prepost optional.Bool
    Adjust optional.String
    FastKPeriod optional.Int64
    SlowKPeriod optional.Int64
    SlowDPeriod optional.Int64
    SlowKmaType optional.String
    SlowDmaType optional.String
}

func (a *TechnicalIndicatorApiService) GetTechnicalIndicatorStoch(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *TechnicalIndicatorApiGetTechnicalIndicatorStochOpts) ([]TechnicalIndicatorResponseStoch, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TechnicalIndicatorResponseStoch
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series/stoch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cik.IsSet() {
		localVarQueryParams.Add("cik", parameterToString(localVarOptionals.Cik.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cusip.IsSet() {
		localVarQueryParams.Add("cusip", parameterToString(localVarOptionals.Cusip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Isin.IsSet() {
		localVarQueryParams.Add("isin", parameterToString(localVarOptionals.Isin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CompositeFigi.IsSet() {
		localVarQueryParams.Add("composite_figi", parameterToString(localVarOptionals.CompositeFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShareFigi.IsSet() {
		localVarQueryParams.Add("share_figi", parameterToString(localVarOptionals.ShareFigi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lei.IsSet() {
		localVarQueryParams.Add("lei", parameterToString(localVarOptionals.Lei.Value(), ""))
	}
	localVarQueryParams.Add("dataset", parameterToString(dataset, ""))
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Adjust.IsSet() {
		localVarQueryParams.Add("adjust", parameterToString(localVarOptionals.Adjust.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FastKPeriod.IsSet() {
		localVarQueryParams.Add("fast_k_period", parameterToString(localVarOptionals.FastKPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlowKPeriod.IsSet() {
		localVarQueryParams.Add("slow_k_period", parameterToString(localVarOptionals.SlowKPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlowDPeriod.IsSet() {
		localVarQueryParams.Add("slow_d_period", parameterToString(localVarOptionals.SlowDPeriod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlowKmaType.IsSet() {
		localVarQueryParams.Add("slow_kma_type", parameterToString(localVarOptionals.SlowKmaType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SlowDmaType.IsSet() {
		localVarQueryParams.Add("slow_dma_type", parameterToString(localVarOptionals.SlowDmaType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			
			localVarQueryParams.Add("apikey", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TechnicalIndicatorResponseStoch
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

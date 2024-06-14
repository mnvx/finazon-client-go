
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

type DataApiService service
/*
DataApiService Dividends calendar
Returns the dividends calendar from Benzinga.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetBenzingaDividendsCalendarOpts - Optional Parameters:
     * @param "Date" (optional.String) -  Specifies the exact date to get the data for
     * @param "StartAt" (optional.Int64) -  Filter events by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter events by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
@return BenzingaDividendsCalendarResponseBody
*/

type DataApiGetBenzingaDividendsCalendarOpts struct {
    Date optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
}

func (a *DataApiService) GetBenzingaDividendsCalendar(ctx context.Context, ticker string, localVarOptionals *DataApiGetBenzingaDividendsCalendarOpts) (BenzingaDividendsCalendarResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BenzingaDividendsCalendarResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/benzinga/dividends_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	if localVarOptionals != nil && localVarOptionals.Date.IsSet() {
		localVarQueryParams.Add("date", parameterToString(localVarOptionals.Date.Value(), ""))
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
			var v BenzingaDividendsCalendarResponseBody
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
DataApiService Earnings calendar
Returns the earnings calendar from Benzinga.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetBenzingaEarningsCalendarOpts - Optional Parameters:
     * @param "Date" (optional.String) -  Specifies the exact date to get the data for
     * @param "StartAt" (optional.Int64) -  Filter events by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter events by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
@return BenzingaEarningsCalendarResponseBody
*/

type DataApiGetBenzingaEarningsCalendarOpts struct {
    Date optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
}

func (a *DataApiService) GetBenzingaEarningsCalendar(ctx context.Context, ticker string, localVarOptionals *DataApiGetBenzingaEarningsCalendarOpts) (BenzingaEarningsCalendarResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BenzingaEarningsCalendarResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/benzinga/earnings_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	if localVarOptionals != nil && localVarOptionals.Date.IsSet() {
		localVarQueryParams.Add("date", parameterToString(localVarOptionals.Date.Value(), ""))
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
			var v BenzingaEarningsCalendarResponseBody
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
DataApiService IPO data
Returns IPO data from Benzinga.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetBenzingaIPOOpts - Optional Parameters:
     * @param "StartAt" (optional.Int64) -  Filter events by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter events by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Exchange" (optional.String) -  Exchange where instrument is traded
@return BenzingaIpoResponseBody
*/

type DataApiGetBenzingaIPOOpts struct {
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Exchange optional.String
}

func (a *DataApiService) GetBenzingaIPO(ctx context.Context, localVarOptionals *DataApiGetBenzingaIPOOpts) (BenzingaIpoResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BenzingaIpoResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/benzinga/ipo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if localVarOptionals != nil && localVarOptionals.Exchange.IsSet() {
		localVarQueryParams.Add("exchange", parameterToString(localVarOptionals.Exchange.Value(), ""))
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
			var v BenzingaIpoResponseBody
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
DataApiService News articles
Returns the news articles from Benzinga.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetBenzingaNewsOpts - Optional Parameters:
     * @param "Date" (optional.String) -  Specifies the exact date to get the data for
     * @param "StartAt" (optional.Int64) -  Filter events by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter events by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
@return BenzingaNewsResponseBody
*/

type DataApiGetBenzingaNewsOpts struct {
    Date optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
}

func (a *DataApiService) GetBenzingaNews(ctx context.Context, ticker string, localVarOptionals *DataApiGetBenzingaNewsOpts) (BenzingaNewsResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BenzingaNewsResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/benzinga/news"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	if localVarOptionals != nil && localVarOptionals.Date.IsSet() {
		localVarQueryParams.Add("date", parameterToString(localVarOptionals.Date.Value(), ""))
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
			var v BenzingaNewsResponseBody
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
DataApiService Time series
This endpoint returns a time series of data points for any given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetBinanceQuotesOpts - Optional Parameters:
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
@return []QuoteBinanceItem
*/

type DataApiGetBinanceQuotesOpts struct {
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
}

func (a *DataApiService) GetBinanceQuotes(ctx context.Context, ticker string, interval string, localVarOptionals *DataApiGetBinanceQuotesOpts) ([]QuoteBinanceItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []QuoteBinanceItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/binance/time_series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	localVarQueryParams.Add("interval", parameterToString(interval, ""))
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
			var v []QuoteBinanceItem
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
DataApiService Time series
This endpoint returns a time series of data points for any given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param interval Interval between two consecutive points in time series
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetCryptoQuotesOpts - Optional Parameters:
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
@return []QuoteItem
*/

type DataApiGetCryptoQuotesOpts struct {
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
}

func (a *DataApiService) GetCryptoQuotes(ctx context.Context, interval string, ticker string, localVarOptionals *DataApiGetCryptoQuotesOpts) ([]QuoteItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []QuoteItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/crypto/time_series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("interval", parameterToString(interval, ""))
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
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
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
			var v []QuoteItem
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
DataApiService Filings
Real-time and historical access to all forms, filings, and exhibits directly from the SEC&#x27;s EDGAR system.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetFilingsOpts - Optional Parameters:
     * @param "CikCode" (optional.Int64) -  Filter by Central Index Key
     * @param "Ticker" (optional.String) -  Filter by ticker
     * @param "FormType" (optional.String) -  Filter by form types
     * @param "FilledFromTs" (optional.Int64) -  Filter by filled time from
     * @param "FilledToTs" (optional.Int64) -  Filter by filled time to
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
@return []Filing
*/

type DataApiGetFilingsOpts struct {
    CikCode optional.Int64
    Ticker optional.String
    FormType optional.String
    FilledFromTs optional.Int64
    FilledToTs optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Cqs optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
}

func (a *DataApiService) GetFilings(ctx context.Context, localVarOptionals *DataApiGetFilingsOpts) ([]Filing, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []Filing
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sec/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CikCode.IsSet() {
		localVarQueryParams.Add("cik_code", parameterToString(localVarOptionals.CikCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FormType.IsSet() {
		localVarQueryParams.Add("form_type", parameterToString(localVarOptionals.FormType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilledFromTs.IsSet() {
		localVarQueryParams.Add("filled_from_ts", parameterToString(localVarOptionals.FilledFromTs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilledToTs.IsSet() {
		localVarQueryParams.Add("filled_to_ts", parameterToString(localVarOptionals.FilledToTs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cqs.IsSet() {
		localVarQueryParams.Add("cqs", parameterToString(localVarOptionals.Cqs.Value(), ""))
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
			var v []Filing
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
DataApiService Time series
This endpoint returns a time series of data points for any given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param interval Interval between two consecutive points in time series
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetForexQuotesOpts - Optional Parameters:
     * @param "StartAt" (optional.Int64) -  Filter output by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter output by end time using a UNIX timestamp
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
@return []QuoteForexItem
*/

type DataApiGetForexQuotesOpts struct {
    StartAt optional.Int64
    EndAt optional.Int64
    Page optional.Int32
    PageSize optional.Int32
    Order optional.String
}

func (a *DataApiService) GetForexQuotes(ctx context.Context, interval string, ticker string, localVarOptionals *DataApiGetForexQuotesOpts) ([]QuoteForexItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []QuoteForexItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/forex/time_series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("interval", parameterToString(interval, ""))
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
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
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
			var v []QuoteForexItem
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
DataApiService Price
Returns price value for given binance ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceBinanceOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceBinanceOpts struct {
    At optional.Int64
    Ticker optional.String
}

func (a *DataApiService) GetPriceBinance(ctx context.Context, localVarOptionals *DataApiGetPriceBinanceOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/binance/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Price
Returns price value for given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param optional nil or *DataApiGetPriceCommonOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceCommonOpts struct {
    At optional.Int64
    Prepost optional.Bool
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Ticker optional.String
}

func (a *DataApiService) GetPriceCommon(ctx context.Context, dataset string, localVarOptionals *DataApiGetPriceCommonOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
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
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Price
Returns price value for given crypto ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceCryptoOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceCryptoOpts struct {
    At optional.Int64
    Ticker optional.String
}

func (a *DataApiService) GetPriceCrypto(ctx context.Context, localVarOptionals *DataApiGetPriceCryptoOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/crypto/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Forex price
Returns price value for given forex ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceForexOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceForexOpts struct {
    At optional.Int64
    Ticker optional.String
}

func (a *DataApiService) GetPriceForex(ctx context.Context, localVarOptionals *DataApiGetPriceForexOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/forex/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Price
Returns price value for given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceSipOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceSipOpts struct {
    At optional.Int64
    Prepost optional.Bool
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Ticker optional.String
}

func (a *DataApiService) GetPriceSip(ctx context.Context, localVarOptionals *DataApiGetPriceSipOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sip_non_pro/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
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
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Price
Returns price value for given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceSip_1Opts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Prepost" (optional.Bool) -  Indicates whether data should be included for extended hours of trading
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceSip_1Opts struct {
    At optional.Int64
    Prepost optional.Bool
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Ticker optional.String
}

func (a *DataApiService) GetPriceSip_1(ctx context.Context, localVarOptionals *DataApiGetPriceSip_1Opts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sip_pro/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prepost.IsSet() {
		localVarQueryParams.Add("prepost", parameterToString(localVarOptionals.Prepost.Value(), ""))
	}
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
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Price
Returns price value for given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DataApiGetPriceUsStocksOpts - Optional Parameters:
     * @param "At" (optional.Int64) -  Filter by start time using a UNIX timestamp. If not specified - last price. Else - last price from 1min interval at the event_at &lt;&#x3D; at
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Ticker" (optional.String) -  Filter by ticker symbol
@return PriceResponseBody
*/

type DataApiGetPriceUsStocksOpts struct {
    At optional.Int64
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Ticker optional.String
}

func (a *DataApiService) GetPriceUsStocks(ctx context.Context, localVarOptionals *DataApiGetPriceUsStocksOpts) (PriceResponseBody, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PriceResponseBody
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/us_stocks_essential/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.At.IsSet() {
		localVarQueryParams.Add("at", parameterToString(localVarOptionals.At.Value(), ""))
	}
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
	if localVarOptionals != nil && localVarOptionals.Ticker.IsSet() {
		localVarQueryParams.Add("ticker", parameterToString(localVarOptionals.Ticker.Value(), ""))
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
			var v PriceResponseBody
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
DataApiService Time series
This endpoint returns a time series of data points for any given ticker.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetQuotesOpts - Optional Parameters:
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
@return []QuoteItem
*/

type DataApiGetQuotesOpts struct {
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
}

func (a *DataApiService) GetQuotes(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetQuotesOpts) ([]QuoteItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []QuoteItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/time_series"

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
			var v []QuoteItem
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
DataApiService SIP trades
Returns detailed information on trades executed through the Securities Information Processor (SIP).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetSipTradesOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market center
     * @param "StartAt" (optional.Int64) -  Filter trades by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter trades by end time using a UNIX timestamp
     * @param "Tape" (optional.String) -  Filter by tape
     * @param "Page" (optional.Int64) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int64) -  Number of items displayed per page in a paginated result
     * @param "Order" (optional.String) -  Sorting order of the output series
@return []SipTradesItem
*/

type DataApiGetSipTradesOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Tape optional.String
    Page optional.Int64
    PageSize optional.Int64
    Order optional.String
}

func (a *DataApiService) GetSipTrades(ctx context.Context, ticker string, localVarOptionals *DataApiGetSipTradesOpts) ([]SipTradesItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []SipTradesItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sip/trades"

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
	localVarQueryParams.Add("ticker", parameterToString(ticker, ""))
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tape.IsSet() {
		localVarQueryParams.Add("tape", parameterToString(localVarOptionals.Tape.Value(), ""))
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
			var v []SipTradesItem
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
DataApiService ATR Technical indicators
The Average True Range (ATR) is a volatility indicator that measures the average range of price movement over a specified period, helping traders assess market volatility.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorAtrOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorAtrOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorAtr(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorAtrOpts) ([]TechnicalIndicatorResponseAtr, *http.Response, error) {
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
DataApiService Overlap Studies
Bollinger Bands (BBANDS) are volatility bands placed above and below a moving average, measuring price volatility and helping traders identify potential overbought or oversold conditions.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorBBandsOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorBBandsOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorBBands(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorBBandsOpts) ([]TechnicalIndicatorResponseBBands, *http.Response, error) {
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
DataApiService Overlap Studies
The Ichimoku Cloud (ICHIMOKU) is a comprehensive trend-following indicator that combines multiple moving averages and support/resistance levels to help traders identify potential entry and exit points, trend direction, and momentum.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorIchimokuOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorIchimokuOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorIchimoku(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorIchimokuOpts) ([]TechnicalIndicatorResponseIchimoku, *http.Response, error) {
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
DataApiService Overlap Studies
The Moving Average (MA) is a smoothing indicator that calculates the average price of a security over a specified period, helping traders identify trends and potential support or resistance levels.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorMaOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorMaOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorMa(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorMaOpts) ([]TechnicalIndicatorResponseMa, *http.Response, error) {
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
DataApiService Momentum Indicators
The Moving Average Convergence Divergence (MACD) is a momentum indicator that measures the difference between two moving averages, with a signal line used to identify potential trend reversals and trading opportunities.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorMacdOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorMacdOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorMacd(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorMacdOpts) ([]TechnicalIndicatorResponseMacd, *http.Response, error) {
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
DataApiService Volume Indicators
The On Balance Volume (OBV) indicator is a cumulative volume-based tool used to measure buying and selling pressure, helping traders identify potential price trends and reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorObvOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorObvOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorObv(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorObvOpts) ([]TechnicalIndicatorResponseObv, *http.Response, error) {
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
DataApiService Momentum Indicators
The Relative Strength Index (RSI) is a momentum oscillator that measures the speed and change of price movements, helping traders identify potential overbought or oversold conditions and trend reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorRsiOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorRsiOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorRsi(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorRsiOpts) ([]TechnicalIndicatorResponseRsi, *http.Response, error) {
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
DataApiService Overlap Studies
The Parabolic SAR (SAR) is a trend-following indicator that calculates potential support and resistance levels based on a security&#x27;s price and time, helping traders identify potential entry and exit points.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorSarOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorSarOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorSar(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorSarOpts) ([]TechnicalIndicatorResponseSar, *http.Response, error) {
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
DataApiService Momentum Indicators
The Stochastic Oscillator (STOCH) is a momentum indicator that compares a security&#x27;s closing price to its price range over a specified period, helping traders identify potential overbought or oversold conditions and trend reversals.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param interval Interval between two consecutive points in time series
 * @param optional nil or *DataApiGetTechnicalIndicatorStochOpts - Optional Parameters:
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

type DataApiGetTechnicalIndicatorStochOpts struct {
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

func (a *DataApiService) GetTechnicalIndicatorStoch(ctx context.Context, dataset string, ticker string, interval string, localVarOptionals *DataApiGetTechnicalIndicatorStochOpts) ([]TechnicalIndicatorResponseStoch, *http.Response, error) {
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
/*
DataApiService Ticker snapshot
This endpoint returns a combination of different data points, such as daily performance, last quote, last trade, minute data, and previous day performance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetTickerSnapshotOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Market" (optional.String) -  Filter by market
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
@return InlineResponse200
*/

type DataApiGetTickerSnapshotOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Market optional.String
    Country optional.String
}

func (a *DataApiService) GetTickerSnapshot(ctx context.Context, dataset string, ticker string, localVarOptionals *DataApiGetTickerSnapshotOpts) (InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ticker/snapshot"

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
	if localVarOptionals != nil && localVarOptionals.Market.IsSet() {
		localVarQueryParams.Add("market", parameterToString(localVarOptionals.Market.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
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
			var v InlineResponse200
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
DataApiService Trades
Returns general information on executed trades.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dataset Filter by Finazon&#x27;s dataset code
 * @param ticker Filter by ticker symbol
 * @param optional nil or *DataApiGetTradesOpts - Optional Parameters:
     * @param "Cqs" (optional.String) -  Filter by cqs symbol
     * @param "Cik" (optional.String) -  Filter by cik code
     * @param "Cusip" (optional.String) -  Filter by cusip code
     * @param "Isin" (optional.String) -  Filter by isin code
     * @param "CompositeFigi" (optional.String) -  Filter by composite figi code
     * @param "ShareFigi" (optional.String) -  Filter by share class figi code
     * @param "Lei" (optional.String) -  Filter by lei code
     * @param "Country" (optional.String) -  Filter by ISO 3166 alpha-2 code
     * @param "StartAt" (optional.Int64) -  Filter trades by start time using a UNIX timestamp
     * @param "EndAt" (optional.Int64) -  Filter trades by end time using a UNIX timestamp
     * @param "Order" (optional.String) -  Sorting order of the output series
     * @param "Page" (optional.Int32) -  Specific page of a paginated result to be displayed
     * @param "PageSize" (optional.Int32) -  Number of items displayed per page in a paginated result
@return []TradesItem
*/

type DataApiGetTradesOpts struct {
    Cqs optional.String
    Cik optional.String
    Cusip optional.String
    Isin optional.String
    CompositeFigi optional.String
    ShareFigi optional.String
    Lei optional.String
    Country optional.String
    StartAt optional.Int64
    EndAt optional.Int64
    Order optional.String
    Page optional.Int32
    PageSize optional.Int32
}

func (a *DataApiService) GetTrades(ctx context.Context, dataset string, ticker string, localVarOptionals *DataApiGetTradesOpts) ([]TradesItem, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TradesItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/trades"

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
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartAt.IsSet() {
		localVarQueryParams.Add("start_at", parameterToString(localVarOptionals.StartAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndAt.IsSet() {
		localVarQueryParams.Add("end_at", parameterToString(localVarOptionals.EndAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
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
			var v []TradesItem
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

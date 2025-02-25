/*
 * Aioz Stream API
 *
 * Aioz Stream Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aiozstreamsdk
import (
	"context"
	"net/http"
    "net/url"
	"strings"
)
// Linger please
var (
	_ context.Context
)

type WebhookApiListRequest struct { 
    search *string
    sortBy *string
    orderBy *string
    offset *int32
    limit *int32
    encodingFinished *bool
    encodingStarted *bool
    fileReceived *bool
}

func (r WebhookApiListRequest) Search(search string) WebhookApiListRequest {
	r.search = &search
	return r
}
func (r WebhookApiListRequest) SortBy(sortBy string) WebhookApiListRequest {
	r.sortBy = &sortBy
	return r
}
func (r WebhookApiListRequest) OrderBy(orderBy string) WebhookApiListRequest {
	r.orderBy = &orderBy
	return r
}
func (r WebhookApiListRequest) Offset(offset int32) WebhookApiListRequest {
	r.offset = &offset
	return r
}
func (r WebhookApiListRequest) Limit(limit int32) WebhookApiListRequest {
	r.limit = &limit
	return r
}
func (r WebhookApiListRequest) EncodingFinished(encodingFinished bool) WebhookApiListRequest {
	r.encodingFinished = &encodingFinished
	return r
}
func (r WebhookApiListRequest) EncodingStarted(encodingStarted bool) WebhookApiListRequest {
	r.encodingStarted = &encodingStarted
	return r
}
func (r WebhookApiListRequest) FileReceived(fileReceived bool) WebhookApiListRequest {
	r.fileReceived = &fileReceived
	return r
}



type WebhookServiceI interface {
	/*
	 * Create Create webhook
	 * @return WebhookApiCreateRequest
	 */
	
	Create(request CreateWebhookRequest) (*CreateWebhookResponse, error)

	/*
	 * Create Create webhook
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WebhookApiCreateRequest
	 */
    
    CreateWithContext(ctx context.Context, request CreateWebhookRequest) (*CreateWebhookResponse, error)


	/*
	 * Get Get user's webhook by id
	 * @param id webhook's id
	 * @return WebhookApiGetRequest
	 */
	
	Get(id string) (*GetUserWebhookResponse, error)

	/*
	 * Get Get user's webhook by id
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id webhook's id
	 * @return WebhookApiGetRequest
	 */
    
    GetWithContext(ctx context.Context, id string) (*GetUserWebhookResponse, error)


	/*
	 * Update Update event webhook
	 * @param id webhook's id
	 * @return WebhookApiUpdateRequest
	 */
	
	Update(id string, request UpdateWebhookRequest) (*ResponseSuccess, error)

	/*
	 * Update Update event webhook
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id webhook's id
	 * @return WebhookApiUpdateRequest
	 */
    
    UpdateWithContext(ctx context.Context, id string, request UpdateWebhookRequest) (*ResponseSuccess, error)


	/*
	 * Delete Delete webhook
	 * @param id Webhook ID
	 * @return WebhookApiDeleteRequest
	 */
	
	Delete(id string) (*ResponseSuccess, error)

	/*
	 * Delete Delete webhook
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Webhook ID
	 * @return WebhookApiDeleteRequest
	 */
    
    DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


	/*
	 * List Get list webhooks
	 * @return WebhookApiListRequest
	 */
	
	List(r WebhookApiListRequest) (*GetWebhooksListResponse, error)
	

	/*
	 * List Get list webhooks
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WebhookApiListRequest
	 */
    
    ListWithContext(ctx context.Context, r WebhookApiListRequest) (*GetWebhooksListResponse, error)
    


	/*
	 * Check Check webhook by id
	 * @param id webhook's id
	 * @return WebhookApiCheckRequest
	 */
	
	Check(id string) (*ResponseSuccess, error)

	/*
	 * Check Check webhook by id
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id webhook's id
	 * @return WebhookApiCheckRequest
	 */
    
    CheckWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


}

// WebhookService communicating with the Webhook
// endpoints of the AIOZ Stream API
type WebhookService struct {
    client *Client
}



/*
 * Create Create webhook
 * Webhooks can push notifications to your server, rather than polling streaming service for changes

 * @return WebhookApiCreateRequest
 */


func (s *WebhookService) Create(request CreateWebhookRequest) (*CreateWebhookResponse, error) {
	
	
	return s.CreateWithContext(context.Background(), request)

}


/*
 * Create Create webhook
 * Webhooks can push notifications to your server, rather than polling streaming service for changes
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WebhookApiCreateRequest
 */


func (s *WebhookService) CreateWithContext(ctx context.Context, request CreateWebhookRequest) (*CreateWebhookResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

	req, err := s.client.prepareRequest(ctx, http.MethodPost, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(CreateWebhookResponse)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}



/*
 * Get Get user's webhook by id
 * Retrieve webhook details by id.

 * @param id webhook's id
 * @return WebhookApiGetRequest
 */


func (s *WebhookService) Get(id string) (*GetUserWebhookResponse, error) {
	
	
	return s.GetWithContext(context.Background(), id)

}


/*
 * Get Get user's webhook by id
 * Retrieve webhook details by id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id webhook's id
 * @return WebhookApiGetRequest
 */


func (s *WebhookService) GetWithContext(ctx context.Context, id string) (*GetUserWebhookResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/webhooks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}


	req, err := s.client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(GetUserWebhookResponse)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}



/*
 * Update Update event webhook
 * This endpoint will update the indicated webhook.

 * @param id webhook's id
 * @return WebhookApiUpdateRequest
 */


func (s *WebhookService) Update(id string, request UpdateWebhookRequest) (*ResponseSuccess, error) {
	
	
	return s.UpdateWithContext(context.Background(), id, request)

}


/*
 * Update Update event webhook
 * This endpoint will update the indicated webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id webhook's id
 * @return WebhookApiUpdateRequest
 */


func (s *WebhookService) UpdateWithContext(ctx context.Context, id string, request UpdateWebhookRequest) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/webhooks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

	req, err := s.client.prepareRequest(ctx, http.MethodPatch, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(ResponseSuccess)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}



/*
 * Delete Delete webhook
 * This endpoint will delete the indicated webhook.

 * @param id Webhook ID
 * @return WebhookApiDeleteRequest
 */


func (s *WebhookService) Delete(id string) (*ResponseSuccess, error) {
	
	
	return s.DeleteWithContext(context.Background(), id)

}


/*
 * Delete Delete webhook
 * This endpoint will delete the indicated webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Webhook ID
 * @return WebhookApiDeleteRequest
 */


func (s *WebhookService) DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/webhooks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}


	req, err := s.client.prepareRequest(ctx, http.MethodDelete, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(ResponseSuccess)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}



/*
 * List Get list webhooks
 * This method returns a list of your webhooks (with all their details). 

You can filter what the webhook list that the API returns using the parameters described below.

 * @return WebhookApiListRequest
 */

func (s *WebhookService) List(r WebhookApiListRequest) (*GetWebhooksListResponse, error) { 
	
	return s.ListWithContext(context.Background(), r)

}


/*
 * List Get list webhooks
 * This method returns a list of your webhooks (with all their details). 

You can filter what the webhook list that the API returns using the parameters described below.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WebhookApiListRequest
 */

func (s *WebhookService) ListWithContext(ctx context.Context, r WebhookApiListRequest) (*GetWebhooksListResponse, error) { 
	var localVarPostBody interface{}

	localVarPath := "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.encodingFinished != nil {
		localVarQueryParams.Add("encoding_finished", parameterToString(*r.encodingFinished, ""))
	}
	if r.encodingStarted != nil {
		localVarQueryParams.Add("encoding_started", parameterToString(*r.encodingStarted, ""))
	}
	if r.fileReceived != nil {
		localVarQueryParams.Add("file_received", parameterToString(*r.fileReceived, ""))
	}

	req, err := s.client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(GetWebhooksListResponse)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}



/*
 * Check Check webhook by id
 * This endpoint will check the indicated webhook.

 * @param id webhook's id
 * @return WebhookApiCheckRequest
 */


func (s *WebhookService) Check(id string) (*ResponseSuccess, error) {
	
	
	return s.CheckWithContext(context.Background(), id)

}


/*
 * Check Check webhook by id
 * This endpoint will check the indicated webhook.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id webhook's id
 * @return WebhookApiCheckRequest
 */


func (s *WebhookService) CheckWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/webhooks/check/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}


	req, err := s.client.prepareRequest(ctx, http.MethodPost, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

    res := new(ResponseSuccess)
    _, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

    return res, nil

}

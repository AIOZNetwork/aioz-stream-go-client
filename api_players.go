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
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type PlayersApiListRequest struct {
	search  *string
	sortBy  *string
	orderBy *string
	offset  *int32
	limit   *int32
}

func (r PlayersApiListRequest) Search(search string) PlayersApiListRequest {
	r.search = &search
	return r
}
func (r PlayersApiListRequest) SortBy(sortBy string) PlayersApiListRequest {
	r.sortBy = &sortBy
	return r
}
func (r PlayersApiListRequest) OrderBy(orderBy string) PlayersApiListRequest {
	r.orderBy = &orderBy
	return r
}
func (r PlayersApiListRequest) Offset(offset int32) PlayersApiListRequest {
	r.offset = &offset
	return r
}
func (r PlayersApiListRequest) Limit(limit int32) PlayersApiListRequest {
	r.limit = &limit
	return r
}

type PlayersServiceI interface {
	/*
	 * Create Create a player theme
	 * @return PlayersApiCreateRequest
	 */

	Create(request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error)

	/*
	 * Create Create a player theme
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PlayersApiCreateRequest
	 */

	CreateWithContext(ctx context.Context, request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error)

	/*
	 * Get Get a player theme by ID
	 * @param id Player theme ID
	 * @return PlayersApiGetRequest
	 */

	Get(id string) (*GetPlayerThemeByIdResponse, error)

	/*
	 * Get Get a player theme by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Player theme ID
	 * @return PlayersApiGetRequest
	 */

	GetWithContext(ctx context.Context, id string) (*GetPlayerThemeByIdResponse, error)

	/*
	 * Update Update a player theme by ID
	 * @param id Player theme ID
	 * @return PlayersApiUpdateRequest
	 */

	Update(id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error)

	/*
	 * Update Update a player theme by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Player theme ID
	 * @return PlayersApiUpdateRequest
	 */

	UpdateWithContext(ctx context.Context, id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error)

	/*
	 * Delete Delete a player theme by ID
	 * @param id Player theme ID
	 * @return PlayersApiDeleteRequest
	 */

	Delete(id string) (*ResponseSuccess, error)

	/*
	 * Delete Delete a player theme by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Player theme ID
	 * @return PlayersApiDeleteRequest
	 */

	DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)

	/*
	 * List List all player themes
	 * @return PlayersApiListRequest
	 */

	List(r PlayersApiListRequest) (*GetPlayerThemeResponse, error)

	/*
	 * List List all player themes
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PlayersApiListRequest
	 */

	ListWithContext(ctx context.Context, r PlayersApiListRequest) (*GetPlayerThemeResponse, error)

	/*
	 * UploadLogo Upload a logo for a player theme by ID
	 * @param id Player theme ID
	 * @return PlayersApiUploadLogoRequest
	 */
	UploadLogo(id string, link string, fileName string, fileReader io.Reader) (*UploadLogoByIdResponse, error)
	/*
	 * UploadLogo Upload a logo for a player theme by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Player theme ID
	 * @return PlayersApiUploadLogoRequest
	 */
	UploadLogoWithContext(ctx context.Context, id string, link string, fileName string, fileReader io.Reader) (*UploadLogoByIdResponse, error)

	/*
	 * DeleteLogo Delete a logo for a player theme by ID
	 * @param id Player theme ID
	 * @return PlayersApiDeleteLogoRequest
	 */

	DeleteLogo(id string) (*ResponseSuccess, error)

	/*
	 * DeleteLogo Delete a logo for a player theme by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Player theme ID
	 * @return PlayersApiDeleteLogoRequest
	 */

	DeleteLogoWithContext(ctx context.Context, id string) (*ResponseSuccess, error)

	/*
	 * AddPlayer Add a player theme to a video
	 * @return PlayersApiAddPlayerRequest
	 */

	AddPlayer(request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error)

	/*
	 * AddPlayer Add a player theme to a video
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PlayersApiAddPlayerRequest
	 */

	AddPlayerWithContext(ctx context.Context, request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error)

	/*
	 * RemovePlayer Remove a player theme from a video
	 * @return PlayersApiRemovePlayerRequest
	 */

	RemovePlayer(request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error)

	/*
	 * RemovePlayer Remove a player theme from a video
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return PlayersApiRemovePlayerRequest
	 */

	RemovePlayerWithContext(ctx context.Context, request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error)
}

// PlayersService communicating with the Players
// endpoints of the AIOZ Stream API
type PlayersService struct {
	client *Client
}

/*
 * Create Create a player theme
 * Create a player for your video, and customize it.

 * @return PlayersApiCreateRequest
 */

func (s *PlayersService) Create(request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error) {

	return s.CreateWithContext(context.Background(), request)

}

/*
 * Create Create a player theme
 * Create a player for your video, and customize it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PlayersApiCreateRequest
 */

func (s *PlayersService) CreateWithContext(ctx context.Context, request CreatePlayerThemeRequest) (*CreatePlayerThemesResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/players"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

	req, err := s.client.prepareRequest(ctx, http.MethodPost, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(CreatePlayerThemesResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Get Get a player theme by ID
 * Retrieve a player theme by its ID, as well as details about it.

 * @param id Player theme ID
 * @return PlayersApiGetRequest
 */

func (s *PlayersService) Get(id string) (*GetPlayerThemeByIdResponse, error) {

	return s.GetWithContext(context.Background(), id)

}

/*
 * Get Get a player theme by ID
 * Retrieve a player theme by its ID, as well as details about it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiGetRequest
 */

func (s *PlayersService) GetWithContext(ctx context.Context, id string) (*GetPlayerThemeByIdResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := s.client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(GetPlayerThemeByIdResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Update Update a player theme by ID
 * Use a player ID to update specific details for a player.

 * @param id Player theme ID
 * @return PlayersApiUpdateRequest
 */

func (s *PlayersService) Update(id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error) {

	return s.UpdateWithContext(context.Background(), id, input)

}

/*
 * Update Update a player theme by ID
 * Use a player ID to update specific details for a player.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiUpdateRequest
 */

func (s *PlayersService) UpdateWithContext(ctx context.Context, id string, input UpdatePlayerThemeRequest) (*UpdatePlayerThemeResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = input

	req, err := s.client.prepareRequest(ctx, http.MethodPatch, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(UpdatePlayerThemeResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Delete Delete a player theme by ID
 * Delete a player if you no longer need it. You can delete any player that you have the player ID for.

 * @param id Player theme ID
 * @return PlayersApiDeleteRequest
 */

func (s *PlayersService) Delete(id string) (*ResponseSuccess, error) {

	return s.DeleteWithContext(context.Background(), id)

}

/*
 * Delete Delete a player theme by ID
 * Delete a player if you no longer need it. You can delete any player that you have the player ID for.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiDeleteRequest
 */

func (s *PlayersService) DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/{id}"
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
 * List List all player themes
 * Retrieve a list of all the player themes you created, as well as details about each one.

 * @return PlayersApiListRequest
 */

func (s *PlayersService) List(r PlayersApiListRequest) (*GetPlayerThemeResponse, error) {

	return s.ListWithContext(context.Background(), r)

}

/*
 * List List all player themes
 * Retrieve a list of all the player themes you created, as well as details about each one.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PlayersApiListRequest
 */

func (s *PlayersService) ListWithContext(ctx context.Context, r PlayersApiListRequest) (*GetPlayerThemeResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/players"

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

	req, err := s.client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(GetPlayerThemeResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * UploadLogo Upload a logo for a player theme by ID
 * Upload a logo for a player theme by its ID.

 * @param id Player theme ID
 * @return PlayersApiUploadLogoRequest
 */

func (s *PlayersService) UploadLogoFile(id string, file *os.File, link string) (*UploadLogoByIdResponse, error) {
	return s.UploadLogoFileWithContext(context.Background(), id, file, link)
}

/*
 * UploadLogo Upload a logo for a player theme by ID
 * Upload a logo for a player theme by its ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiUploadLogoRequest
 */

func (s *PlayersService) UploadLogoFileWithContext(ctx context.Context, id string, file *os.File, link string) (*UploadLogoByIdResponse, error) {
	return s.UploadLogoWithContext(ctx, id, link, file.Name(), io.Reader(file))
}

/*
* UploadLogo Upload a logo for a player theme by ID
* Upload a logo for a player theme by its ID.

* @param id Player theme ID
* @return PlayersApiUploadLogoRequest
 */
func (s *PlayersService) UploadLogo(id string, link string, fileName string, fileReader io.Reader) (*UploadLogoByIdResponse, error) {
	return s.UploadLogoWithContext(context.Background(), id, link, fileName, fileReader)
}

/*
 * UploadLogo Upload a logo for a player theme by ID
 * Upload a logo for a player theme by its ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiUploadLogoRequest
 */
func (s *PlayersService) UploadLogoWithContext(ctx context.Context, id string, link string, fileName string, fileReader io.Reader) (*UploadLogoByIdResponse, error) {
	localVarPath := "/players/{id}/logo"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)

	localVarFormParams["link"] = parameterToString(link, "")

	req, err := s.client.prepareUploadRequest(ctx, http.MethodPost, localVarPath, fileName, fileReader, localVarHeaderParams, localVarQueryParams, localVarFormParams)

	if err != nil {
		return nil, err
	}

	res := new(UploadLogoByIdResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * DeleteLogo Delete a logo for a player theme by ID
 * Delete the logo associated to a player.

 * @param id Player theme ID
 * @return PlayersApiDeleteLogoRequest
 */

func (s *PlayersService) DeleteLogo(id string) (*ResponseSuccess, error) {

	return s.DeleteLogoWithContext(context.Background(), id)

}

/*
 * DeleteLogo Delete a logo for a player theme by ID
 * Delete the logo associated to a player.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Player theme ID
 * @return PlayersApiDeleteLogoRequest
 */

func (s *PlayersService) DeleteLogoWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/{id}/logo"
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
 * AddPlayer Add a player theme to a video
 * Add a player theme to a video by Id.

 * @return PlayersApiAddPlayerRequest
 */

func (s *PlayersService) AddPlayer(request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error) {

	return s.AddPlayerWithContext(context.Background(), request)

}

/*
 * AddPlayer Add a player theme to a video
 * Add a player theme to a video by Id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PlayersApiAddPlayerRequest
 */

func (s *PlayersService) AddPlayerWithContext(ctx context.Context, request AddPlayerThemesToVideoRequest) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/add-player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

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

/*
 * RemovePlayer Remove a player theme from a video
 * Remove a player theme from a video by Id.

 * @return PlayersApiRemovePlayerRequest
 */

func (s *PlayersService) RemovePlayer(request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error) {

	return s.RemovePlayerWithContext(context.Background(), request)

}

/*
 * RemovePlayer Remove a player theme from a video
 * Remove a player theme from a video by Id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return PlayersApiRemovePlayerRequest
 */

func (s *PlayersService) RemovePlayerWithContext(ctx context.Context, request RemovePlayerThemesFromVideoRequest) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/players/remove-player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

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

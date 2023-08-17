/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
)


// RemotesMavenAPIService RemotesMavenAPI service
type RemotesMavenAPIService service

type RemotesMavenAPIRemotesMavenMavenCreateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	mavenMavenRemote *MavenMavenRemote
}

func (r RemotesMavenAPIRemotesMavenMavenCreateRequest) MavenMavenRemote(mavenMavenRemote MavenMavenRemote) RemotesMavenAPIRemotesMavenMavenCreateRequest {
	r.mavenMavenRemote = &mavenMavenRemote
	return r
}

func (r RemotesMavenAPIRemotesMavenMavenCreateRequest) Execute() (*MavenMavenRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenCreateExecute(r)
}

/*
RemotesMavenMavenCreate Create a maven remote

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesMavenAPIRemotesMavenMavenCreateRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenCreate(ctx context.Context) RemotesMavenAPIRemotesMavenMavenCreateRequest {
	return RemotesMavenAPIRemotesMavenMavenCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MavenMavenRemoteResponse
func (a *RemotesMavenAPIService) RemotesMavenMavenCreateExecute(r RemotesMavenAPIRemotesMavenMavenCreateRequest) (*MavenMavenRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/maven/maven/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("mavenMavenRemote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mavenMavenRemote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemotesMavenAPIRemotesMavenMavenDeleteRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	mavenMavenRemoteHref string
}

func (r RemotesMavenAPIRemotesMavenMavenDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenDeleteExecute(r)
}

/*
RemotesMavenMavenDelete Delete a maven remote

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenAPIRemotesMavenMavenDeleteRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenDelete(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenAPIRemotesMavenMavenDeleteRequest {
	return RemotesMavenAPIRemotesMavenMavenDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenAPIService) RemotesMavenMavenDeleteExecute(r RemotesMavenAPIRemotesMavenMavenDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemotesMavenAPIRemotesMavenMavenListRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	pulpLastUpdated *time.Time
	pulpLastUpdatedGt *time.Time
	pulpLastUpdatedGte *time.Time
	pulpLastUpdatedLt *time.Time
	pulpLastUpdatedLte *time.Time
	pulpLastUpdatedRange *[]time.Time
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) Limit(limit int32) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) Name(name string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) NameContains(nameContains string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) NameIcontains(nameIcontains string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesMavenAPIRemotesMavenMavenListRequest) NameIn(nameIn []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) NameStartswith(nameStartswith string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) Offset(offset int32) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;ca_cert&#x60; - Ca cert * &#x60;-ca_cert&#x60; - Ca cert (descending) * &#x60;client_cert&#x60; - Client cert * &#x60;-client_cert&#x60; - Client cert (descending) * &#x60;client_key&#x60; - Client key * &#x60;-client_key&#x60; - Client key (descending) * &#x60;tls_validation&#x60; - Tls validation * &#x60;-tls_validation&#x60; - Tls validation (descending) * &#x60;username&#x60; - Username * &#x60;-username&#x60; - Username (descending) * &#x60;password&#x60; - Password * &#x60;-password&#x60; - Password (descending) * &#x60;proxy_url&#x60; - Proxy url * &#x60;-proxy_url&#x60; - Proxy url (descending) * &#x60;proxy_username&#x60; - Proxy username * &#x60;-proxy_username&#x60; - Proxy username (descending) * &#x60;proxy_password&#x60; - Proxy password * &#x60;-proxy_password&#x60; - Proxy password (descending) * &#x60;download_concurrency&#x60; - Download concurrency * &#x60;-download_concurrency&#x60; - Download concurrency (descending) * &#x60;max_retries&#x60; - Max retries * &#x60;-max_retries&#x60; - Max retries (descending) * &#x60;policy&#x60; - Policy * &#x60;-policy&#x60; - Policy (descending) * &#x60;total_timeout&#x60; - Total timeout * &#x60;-total_timeout&#x60; - Total timeout (descending) * &#x60;connect_timeout&#x60; - Connect timeout * &#x60;-connect_timeout&#x60; - Connect timeout (descending) * &#x60;sock_connect_timeout&#x60; - Sock connect timeout * &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending) * &#x60;sock_read_timeout&#x60; - Sock read timeout * &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending) * &#x60;headers&#x60; - Headers * &#x60;-headers&#x60; - Headers (descending) * &#x60;rate_limit&#x60; - Rate limit * &#x60;-rate_limit&#x60; - Rate limit (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r RemotesMavenAPIRemotesMavenMavenListRequest) Ordering(ordering []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpHrefIn(pulpHrefIn []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpIdIn(pulpIdIn []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesMavenAPIRemotesMavenMavenListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// A list of fields to include in the response.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) Fields(fields []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesMavenAPIRemotesMavenMavenListRequest) ExcludeFields(excludeFields []string) RemotesMavenAPIRemotesMavenMavenListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesMavenAPIRemotesMavenMavenListRequest) Execute() (*PaginatedmavenMavenRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenListExecute(r)
}

/*
RemotesMavenMavenList List maven remotes

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesMavenAPIRemotesMavenMavenListRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenList(ctx context.Context) RemotesMavenAPIRemotesMavenMavenListRequest {
	return RemotesMavenAPIRemotesMavenMavenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedmavenMavenRemoteResponseList
func (a *RemotesMavenAPIService) RemotesMavenMavenListExecute(r RemotesMavenAPIRemotesMavenMavenListRequest) (*PaginatedmavenMavenRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedmavenMavenRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/maven/maven/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.pulpLastUpdated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated", r.pulpLastUpdated, "")
	}
	if r.pulpLastUpdatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gt", r.pulpLastUpdatedGt, "")
	}
	if r.pulpLastUpdatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gte", r.pulpLastUpdatedGte, "")
	}
	if r.pulpLastUpdatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lt", r.pulpLastUpdatedLt, "")
	}
	if r.pulpLastUpdatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lte", r.pulpLastUpdatedLte, "")
	}
	if r.pulpLastUpdatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__range", r.pulpLastUpdatedRange, "csv")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	mavenMavenRemoteHref string
	patchedmavenMavenRemote *PatchedmavenMavenRemote
}

func (r RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest) PatchedmavenMavenRemote(patchedmavenMavenRemote PatchedmavenMavenRemote) RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest {
	r.patchedmavenMavenRemote = &patchedmavenMavenRemote
	return r
}

func (r RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenPartialUpdateExecute(r)
}

/*
RemotesMavenMavenPartialUpdate Update a maven remote

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenPartialUpdate(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest {
	return RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenAPIService) RemotesMavenMavenPartialUpdateExecute(r RemotesMavenAPIRemotesMavenMavenPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedmavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("patchedmavenMavenRemote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchedmavenMavenRemote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemotesMavenAPIRemotesMavenMavenReadRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	mavenMavenRemoteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RemotesMavenAPIRemotesMavenMavenReadRequest) Fields(fields []string) RemotesMavenAPIRemotesMavenMavenReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesMavenAPIRemotesMavenMavenReadRequest) ExcludeFields(excludeFields []string) RemotesMavenAPIRemotesMavenMavenReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesMavenAPIRemotesMavenMavenReadRequest) Execute() (*MavenMavenRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenReadExecute(r)
}

/*
RemotesMavenMavenRead Inspect a maven remote

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenAPIRemotesMavenMavenReadRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenRead(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenAPIRemotesMavenMavenReadRequest {
	return RemotesMavenAPIRemotesMavenMavenReadRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return MavenMavenRemoteResponse
func (a *RemotesMavenAPIService) RemotesMavenMavenReadExecute(r RemotesMavenAPIRemotesMavenMavenReadRequest) (*MavenMavenRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemotesMavenAPIRemotesMavenMavenUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenAPIService
	mavenMavenRemoteHref string
	mavenMavenRemote *MavenMavenRemote
}

func (r RemotesMavenAPIRemotesMavenMavenUpdateRequest) MavenMavenRemote(mavenMavenRemote MavenMavenRemote) RemotesMavenAPIRemotesMavenMavenUpdateRequest {
	r.mavenMavenRemote = &mavenMavenRemote
	return r
}

func (r RemotesMavenAPIRemotesMavenMavenUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenUpdateExecute(r)
}

/*
RemotesMavenMavenUpdate Update a maven remote

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenAPIRemotesMavenMavenUpdateRequest
*/
func (a *RemotesMavenAPIService) RemotesMavenMavenUpdate(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenAPIRemotesMavenMavenUpdateRequest {
	return RemotesMavenAPIRemotesMavenMavenUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenAPIService) RemotesMavenMavenUpdateExecute(r RemotesMavenAPIRemotesMavenMavenUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenAPIService.RemotesMavenMavenUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("mavenMavenRemote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mavenMavenRemote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
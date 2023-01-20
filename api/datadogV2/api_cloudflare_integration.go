// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"bytes"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareIntegrationApi service type
type CloudflareIntegrationApi datadog.Service

type apiCreateCloudflareAccountRequest struct {
	ctx  _context.Context
	body *CloudflareAccountCreateRequest
}

func (a *CloudflareIntegrationApi) buildCreateCloudflareAccountRequest(ctx _context.Context, body CloudflareAccountCreateRequest) (apiCreateCloudflareAccountRequest, error) {
	req := apiCreateCloudflareAccountRequest{
		ctx:  ctx,
		body: &body,
	}
	return req, nil
}

// CreateCloudflareAccount Add Cloudflare account.
// Create a Cloudflare account.
func (a *CloudflareIntegrationApi) CreateCloudflareAccount(ctx _context.Context, body CloudflareAccountCreateRequest) (CloudflareAccountResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateCloudflareAccountRequest(ctx, body)
	if err != nil {
		var localVarReturnValue CloudflareAccountResponse
		return localVarReturnValue, nil, err
	}

	return a.createCloudflareAccountExecute(req)
}

// createCloudflareAccountExecute executes the request.
func (a *CloudflareIntegrationApi) createCloudflareAccountExecute(r apiCreateCloudflareAccountRequest) (CloudflareAccountResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue CloudflareAccountResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.CloudflareIntegrationApi.CreateCloudflareAccount")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integrations/cloudflare/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiDeleteCloudflareAccountRequest struct {
	ctx       _context.Context
	accountId string
}

func (a *CloudflareIntegrationApi) buildDeleteCloudflareAccountRequest(ctx _context.Context, accountId string) (apiDeleteCloudflareAccountRequest, error) {
	req := apiDeleteCloudflareAccountRequest{
		ctx:       ctx,
		accountId: accountId,
	}
	return req, nil
}

// DeleteCloudflareAccount Delete Cloudflare account.
// Delete a Cloudflare account.
func (a *CloudflareIntegrationApi) DeleteCloudflareAccount(ctx _context.Context, accountId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteCloudflareAccountRequest(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return a.deleteCloudflareAccountExecute(req)
}

// deleteCloudflareAccountExecute executes the request.
func (a *CloudflareIntegrationApi) deleteCloudflareAccountExecute(r apiDeleteCloudflareAccountRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.CloudflareIntegrationApi.DeleteCloudflareAccount")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integrations/cloudflare/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiGetCloudflareAccountRequest struct {
	ctx       _context.Context
	accountId string
}

func (a *CloudflareIntegrationApi) buildGetCloudflareAccountRequest(ctx _context.Context, accountId string) (apiGetCloudflareAccountRequest, error) {
	req := apiGetCloudflareAccountRequest{
		ctx:       ctx,
		accountId: accountId,
	}
	return req, nil
}

// GetCloudflareAccount Get Cloudflare account.
// Get a Cloudflare account.
func (a *CloudflareIntegrationApi) GetCloudflareAccount(ctx _context.Context, accountId string) (CloudflareAccountResponse, *_nethttp.Response, error) {
	req, err := a.buildGetCloudflareAccountRequest(ctx, accountId)
	if err != nil {
		var localVarReturnValue CloudflareAccountResponse
		return localVarReturnValue, nil, err
	}

	return a.getCloudflareAccountExecute(req)
}

// getCloudflareAccountExecute executes the request.
func (a *CloudflareIntegrationApi) getCloudflareAccountExecute(r apiGetCloudflareAccountRequest) (CloudflareAccountResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CloudflareAccountResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.CloudflareIntegrationApi.GetCloudflareAccount")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integrations/cloudflare/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListCloudflareAccountsRequest struct {
	ctx _context.Context
}

func (a *CloudflareIntegrationApi) buildListCloudflareAccountsRequest(ctx _context.Context) (apiListCloudflareAccountsRequest, error) {
	req := apiListCloudflareAccountsRequest{
		ctx: ctx,
	}
	return req, nil
}

// ListCloudflareAccounts List Cloudflare accounts.
// List Cloudflare accounts.
func (a *CloudflareIntegrationApi) ListCloudflareAccounts(ctx _context.Context) (CloudflareAccountsResponse, *_nethttp.Response, error) {
	req, err := a.buildListCloudflareAccountsRequest(ctx)
	if err != nil {
		var localVarReturnValue CloudflareAccountsResponse
		return localVarReturnValue, nil, err
	}

	return a.listCloudflareAccountsExecute(req)
}

// listCloudflareAccountsExecute executes the request.
func (a *CloudflareIntegrationApi) listCloudflareAccountsExecute(r apiListCloudflareAccountsRequest) (CloudflareAccountsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CloudflareAccountsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.CloudflareIntegrationApi.ListCloudflareAccounts")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integrations/cloudflare/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateCloudflareAccountRequest struct {
	ctx       _context.Context
	accountId string
	body      *CloudflareAccountUpdateRequest
}

func (a *CloudflareIntegrationApi) buildUpdateCloudflareAccountRequest(ctx _context.Context, accountId string, body CloudflareAccountUpdateRequest) (apiUpdateCloudflareAccountRequest, error) {
	req := apiUpdateCloudflareAccountRequest{
		ctx:       ctx,
		accountId: accountId,
		body:      &body,
	}
	return req, nil
}

// UpdateCloudflareAccount Update Cloudflare account.
// Update a Cloudflare account.
func (a *CloudflareIntegrationApi) UpdateCloudflareAccount(ctx _context.Context, accountId string, body CloudflareAccountUpdateRequest) (CloudflareAccountResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateCloudflareAccountRequest(ctx, accountId, body)
	if err != nil {
		var localVarReturnValue CloudflareAccountResponse
		return localVarReturnValue, nil, err
	}

	return a.updateCloudflareAccountExecute(req)
}

// updateCloudflareAccountExecute executes the request.
func (a *CloudflareIntegrationApi) updateCloudflareAccountExecute(r apiUpdateCloudflareAccountRequest) (CloudflareAccountResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue CloudflareAccountResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.CloudflareIntegrationApi.UpdateCloudflareAccount")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integrations/cloudflare/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewCloudflareIntegrationApi Returns NewCloudflareIntegrationApi.
func NewCloudflareIntegrationApi(client *datadog.APIClient) *CloudflareIntegrationApi {
	return &CloudflareIntegrationApi{
		Client: client,
	}
}
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package personalization

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
)

type Option struct {
	optionType string
	name       string
	value      string
}

func QueryParamOption(name string, val any) Option {
	return Option{
		optionType: "query",
		name:       name,
		value:      parameterToString(val),
	}
}

func HeaderParamOption(name string, val any) Option {
	return Option{
		optionType: "header",
		name:       "itemsPerPage",
		value:      parameterToString(val),
	}
}

type ApiDelRequest struct {
	path       string
	parameters map[string]interface{}
}

// Query parameters to be applied to the current query.
func (r ApiDelRequest) WithParameters(parameters map[string]interface{}) ApiDelRequest {
	r.parameters = parameters
	return r
}

// @return ApiDelRequest
func (c *APIClient) NewApiDelRequest(path string) ApiDelRequest {
	return ApiDelRequest{
		path: path,
	}
}

// @return map[string]interface{}
func (c *APIClient) Del(r ApiDelRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		queryParams.Add("parameters", parameterToString(r.parameters))
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiDeleteUserProfileRequest struct {
	userToken string
}

// @return ApiDeleteUserProfileRequest
func (c *APIClient) NewApiDeleteUserProfileRequest(userToken string) ApiDeleteUserProfileRequest {
	return ApiDeleteUserProfileRequest{
		userToken: userToken,
	}
}

// @return DeleteUserProfileResponse
func (c *APIClient) DeleteUserProfile(r ApiDeleteUserProfileRequest, opts ...Option) (*DeleteUserProfileResponse, error) {
	var (
		postBody    any
		returnValue *DeleteUserProfileResponse
	)

	requestPath := "/1/profiles/{userToken}"
	requestPath = strings.Replace(requestPath, "{"+"userToken"+"}", url.PathEscape(parameterToString(r.userToken)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetRequest struct {
	path       string
	parameters map[string]interface{}
}

// Query parameters to be applied to the current query.
func (r ApiGetRequest) WithParameters(parameters map[string]interface{}) ApiGetRequest {
	r.parameters = parameters
	return r
}

// @return ApiGetRequest
func (c *APIClient) NewApiGetRequest(path string) ApiGetRequest {
	return ApiGetRequest{
		path: path,
	}
}

// @return map[string]interface{}
func (c *APIClient) Get(r ApiGetRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		queryParams.Add("parameters", parameterToString(r.parameters))
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetPersonalizationStrategyRequest struct {
}

// @return ApiGetPersonalizationStrategyRequest
func (c *APIClient) NewApiGetPersonalizationStrategyRequest() ApiGetPersonalizationStrategyRequest {
	return ApiGetPersonalizationStrategyRequest{}
}

// @return PersonalizationStrategyParams
func (c *APIClient) GetPersonalizationStrategy(r ApiGetPersonalizationStrategyRequest, opts ...Option) (*PersonalizationStrategyParams, error) {
	var (
		postBody    any
		returnValue *PersonalizationStrategyParams
	)

	requestPath := "/1/strategies/personalization"

	headers := make(map[string]string)
	queryParams := url.Values{}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetUserTokenProfileRequest struct {
	userToken string
}

// @return ApiGetUserTokenProfileRequest
func (c *APIClient) NewApiGetUserTokenProfileRequest(userToken string) ApiGetUserTokenProfileRequest {
	return ApiGetUserTokenProfileRequest{
		userToken: userToken,
	}
}

// @return GetUserTokenResponse
func (c *APIClient) GetUserTokenProfile(r ApiGetUserTokenProfileRequest, opts ...Option) (*GetUserTokenResponse, error) {
	var (
		postBody    any
		returnValue *GetUserTokenResponse
	)

	requestPath := "/1/profiles/personalization/{userToken}"
	requestPath = strings.Replace(requestPath, "{"+"userToken"+"}", url.PathEscape(parameterToString(r.userToken)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiPostRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// Query parameters to be applied to the current query.
func (r ApiPostRequest) WithParameters(parameters map[string]interface{}) ApiPostRequest {
	r.parameters = parameters
	return r
}

// The parameters to send with the custom request.
func (r ApiPostRequest) WithBody(body map[string]interface{}) ApiPostRequest {
	r.body = body
	return r
}

// @return ApiPostRequest
func (c *APIClient) NewApiPostRequest(path string) ApiPostRequest {
	return ApiPostRequest{
		path: path,
	}
}

// @return map[string]interface{}
func (c *APIClient) Post(r ApiPostRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		queryParams.Add("parameters", parameterToString(r.parameters))
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	postBody = r.body
	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiPutRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// Query parameters to be applied to the current query.
func (r ApiPutRequest) WithParameters(parameters map[string]interface{}) ApiPutRequest {
	r.parameters = parameters
	return r
}

// The parameters to send with the custom request.
func (r ApiPutRequest) WithBody(body map[string]interface{}) ApiPutRequest {
	r.body = body
	return r
}

// @return ApiPutRequest
func (c *APIClient) NewApiPutRequest(path string) ApiPutRequest {
	return ApiPutRequest{
		path: path,
	}
}

// @return map[string]interface{}
func (c *APIClient) Put(r ApiPutRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		queryParams.Add("parameters", parameterToString(r.parameters))
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	postBody = r.body
	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodPut, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiSetPersonalizationStrategyRequest struct {
	personalizationStrategyParams *PersonalizationStrategyParams
}

func (r ApiSetPersonalizationStrategyRequest) WithPersonalizationStrategyParams(personalizationStrategyParams PersonalizationStrategyParams) ApiSetPersonalizationStrategyRequest {
	r.personalizationStrategyParams = &personalizationStrategyParams
	return r
}

// @return ApiSetPersonalizationStrategyRequest
func (c *APIClient) NewApiSetPersonalizationStrategyRequest() ApiSetPersonalizationStrategyRequest {
	return ApiSetPersonalizationStrategyRequest{}
}

// @return SetPersonalizationStrategyResponse
func (c *APIClient) SetPersonalizationStrategy(r ApiSetPersonalizationStrategyRequest, opts ...Option) (*SetPersonalizationStrategyResponse, error) {
	var (
		postBody    any
		returnValue *SetPersonalizationStrategyResponse
	)

	requestPath := "/1/strategies/personalization"

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.personalizationStrategyParams == nil {
		return returnValue, reportError("personalizationStrategyParams is required and must be specified")
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Add(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	postBody = r.personalizationStrategyParams
	req, err := c.prepareRequest(context.Background(), requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

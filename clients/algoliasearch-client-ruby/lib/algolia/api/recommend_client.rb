# Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

require 'cgi'

module Algolia
  class RecommendClient
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @return [Object]
    def del(path, opts = {})
      data, _status_code, _headers = del_with_http_info(path, opts)
      data
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def del_with_http_info(path, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.del ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling RecommendClient.del"
      end
      # resource path
      local_var_path = '/1{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'parameters'] = opts[:'parameters'] if !opts[:'parameters'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.del",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#del\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Delete a Recommend rule.
    # Delete a [Recommend rule](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param object_id [String] Unique record (object) identifier.
    # @param [Hash] opts the optional parameters
    # @return [DeletedAtResponse]
    def delete_recommend_rule(index_name, model, object_id, opts = {})
      data, _status_code, _headers = delete_recommend_rule_with_http_info(index_name, model, object_id, opts)
      data
    end

    # Delete a Recommend rule.
    # Delete a [Recommend rule](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param object_id [String] Unique record (object) identifier.
    # @param [Hash] opts the optional parameters
    # @return [Array<(DeletedAtResponse, Integer, Hash)>] DeletedAtResponse data, response status code and response headers
    def delete_recommend_rule_with_http_info(index_name, model, object_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.delete_recommend_rule ...'
      end
      # verify the required parameter 'index_name' is set
      if @api_client.config.client_side_validation && index_name.nil?
        fail ArgumentError, "Missing the required parameter 'index_name' when calling RecommendClient.delete_recommend_rule"
      end
      # verify the required parameter 'model' is set
      if @api_client.config.client_side_validation && model.nil?
        fail ArgumentError, "Missing the required parameter 'model' when calling RecommendClient.delete_recommend_rule"
      end
      # verify the required parameter 'object_id' is set
      if @api_client.config.client_side_validation && object_id.nil?
        fail ArgumentError, "Missing the required parameter 'object_id' when calling RecommendClient.delete_recommend_rule"
      end
      # resource path
      local_var_path = '/1/indexes/{indexName}/{model}/recommend/rules/{objectID}'.sub('{' + 'indexName' + '}', CGI.escape(index_name.to_s)).sub('{' + 'model' + '}', CGI.escape(model.to_s)).sub('{' + 'objectID' + '}', CGI.escape(object_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'DeletedAtResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.delete_recommend_rule",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#delete_recommend_rule\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @return [Object]
    def get(path, opts = {})
      data, _status_code, _headers = get_with_http_info(path, opts)
      data
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def get_with_http_info(path, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.get ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling RecommendClient.get"
      end
      # resource path
      local_var_path = '/1{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'parameters'] = opts[:'parameters'] if !opts[:'parameters'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get a Recommend rule.
    # Return a [Recommend rule](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param object_id [String] Unique record (object) identifier.
    # @param [Hash] opts the optional parameters
    # @return [RuleResponse]
    def get_recommend_rule(index_name, model, object_id, opts = {})
      data, _status_code, _headers = get_recommend_rule_with_http_info(index_name, model, object_id, opts)
      data
    end

    # Get a Recommend rule.
    # Return a [Recommend rule](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param object_id [String] Unique record (object) identifier.
    # @param [Hash] opts the optional parameters
    # @return [Array<(RuleResponse, Integer, Hash)>] RuleResponse data, response status code and response headers
    def get_recommend_rule_with_http_info(index_name, model, object_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.get_recommend_rule ...'
      end
      # verify the required parameter 'index_name' is set
      if @api_client.config.client_side_validation && index_name.nil?
        fail ArgumentError, "Missing the required parameter 'index_name' when calling RecommendClient.get_recommend_rule"
      end
      # verify the required parameter 'model' is set
      if @api_client.config.client_side_validation && model.nil?
        fail ArgumentError, "Missing the required parameter 'model' when calling RecommendClient.get_recommend_rule"
      end
      # verify the required parameter 'object_id' is set
      if @api_client.config.client_side_validation && object_id.nil?
        fail ArgumentError, "Missing the required parameter 'object_id' when calling RecommendClient.get_recommend_rule"
      end
      # resource path
      local_var_path = '/1/indexes/{indexName}/{model}/recommend/rules/{objectID}'.sub('{' + 'indexName' + '}', CGI.escape(index_name.to_s)).sub('{' + 'model' + '}', CGI.escape(model.to_s)).sub('{' + 'objectID' + '}', CGI.escape(object_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'RuleResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.get_recommend_rule",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#get_recommend_rule\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get a Recommend task's status.
    # Some operations, such as deleting a Recommend rule, will respond with a `taskID` value. Use this value here to check the status of that task.
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param task_id [Integer] Unique identifier of a task. Numeric value (up to 64bits).
    # @param [Hash] opts the optional parameters
    # @return [GetRecommendTaskResponse]
    def get_recommend_status(index_name, model, task_id, opts = {})
      data, _status_code, _headers = get_recommend_status_with_http_info(index_name, model, task_id, opts)
      data
    end

    # Get a Recommend task&#39;s status.
    # Some operations, such as deleting a Recommend rule, will respond with a &#x60;taskID&#x60; value. Use this value here to check the status of that task.
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param task_id [Integer] Unique identifier of a task. Numeric value (up to 64bits).
    # @param [Hash] opts the optional parameters
    # @return [Array<(GetRecommendTaskResponse, Integer, Hash)>] GetRecommendTaskResponse data, response status code and response headers
    def get_recommend_status_with_http_info(index_name, model, task_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.get_recommend_status ...'
      end
      # verify the required parameter 'index_name' is set
      if @api_client.config.client_side_validation && index_name.nil?
        fail ArgumentError, "Missing the required parameter 'index_name' when calling RecommendClient.get_recommend_status"
      end
      # verify the required parameter 'model' is set
      if @api_client.config.client_side_validation && model.nil?
        fail ArgumentError, "Missing the required parameter 'model' when calling RecommendClient.get_recommend_status"
      end
      # verify the required parameter 'task_id' is set
      if @api_client.config.client_side_validation && task_id.nil?
        fail ArgumentError, "Missing the required parameter 'task_id' when calling RecommendClient.get_recommend_status"
      end
      # resource path
      local_var_path = '/1/indexes/{indexName}/{model}/task/{taskID}'.sub('{' + 'indexName' + '}', CGI.escape(index_name.to_s)).sub('{' + 'model' + '}', CGI.escape(model.to_s)).sub('{' + 'taskID' + '}', CGI.escape(task_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'GetRecommendTaskResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.get_recommend_status",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#get_recommend_status\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get recommendations and trending items.
    # Returns results from either recommendation or trending models:    - **Recommendations** are provided by the [Related Products](https://www.algolia.com/doc/guides/algolia-recommend/overview/#related-products-and-related-content) and [Frequently Bought Together](https://www.algolia.com/doc/guides/algolia-recommend/overview/#frequently-bought-together) models   - **Trending** models are [Trending Items and Trending Facet Values](https://www.algolia.com/doc/guides/algolia-recommend/overview/#trending-items-and-trending-facet-values). 
    # @param get_recommendations_params [GetRecommendationsParams] 
    # @param [Hash] opts the optional parameters
    # @return [GetRecommendationsResponse]
    def get_recommendations(get_recommendations_params, opts = {})
      data, _status_code, _headers = get_recommendations_with_http_info(get_recommendations_params, opts)
      data
    end

    # Get recommendations and trending items.
    # Returns results from either recommendation or trending models:    - **Recommendations** are provided by the [Related Products](https://www.algolia.com/doc/guides/algolia-recommend/overview/#related-products-and-related-content) and [Frequently Bought Together](https://www.algolia.com/doc/guides/algolia-recommend/overview/#frequently-bought-together) models   - **Trending** models are [Trending Items and Trending Facet Values](https://www.algolia.com/doc/guides/algolia-recommend/overview/#trending-items-and-trending-facet-values). 
    # @param get_recommendations_params [GetRecommendationsParams] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(GetRecommendationsResponse, Integer, Hash)>] GetRecommendationsResponse data, response status code and response headers
    def get_recommendations_with_http_info(get_recommendations_params, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.get_recommendations ...'
      end
      # verify the required parameter 'get_recommendations_params' is set
      if @api_client.config.client_side_validation && get_recommendations_params.nil?
        fail ArgumentError, "Missing the required parameter 'get_recommendations_params' when calling RecommendClient.get_recommendations"
      end
      # resource path
      local_var_path = '/1/indexes/*/recommendations'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
        header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(get_recommendations_params)

      # return_type
      return_type = opts[:debug_return_type] || 'GetRecommendationsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.get_recommendations",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#get_recommendations\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @option opts [Object] :body Parameters to send with the custom request.
    # @return [Object]
    def post(path, opts = {})
      data, _status_code, _headers = post_with_http_info(path, opts)
      data
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @option opts [Object] :body Parameters to send with the custom request.
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def post_with_http_info(path, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.post ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling RecommendClient.post"
      end
      # resource path
      local_var_path = '/1{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'parameters'] = opts[:'parameters'] if !opts[:'parameters'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
        header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'body'])

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @option opts [Object] :body Parameters to send with the custom request.
    # @return [Object]
    def put(path, opts = {})
      data, _status_code, _headers = put_with_http_info(path, opts)
      data
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified.
    # @param [Hash] opts the optional parameters
    # @option opts [Hash<String, Object>] :parameters Query parameters to apply to the current query.
    # @option opts [Object] :body Parameters to send with the custom request.
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def put_with_http_info(path, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.put ...'
      end
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        fail ArgumentError, "Missing the required parameter 'path' when calling RecommendClient.put"
      end
      # resource path
      local_var_path = '/1{path}'.sub('{' + 'path' + '}', CGI.escape(path.to_s))

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'parameters'] = opts[:'parameters'] if !opts[:'parameters'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
        header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'body'])

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.put",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#put\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # List Recommend rules.
    # List [Recommend rules](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param [Hash] opts the optional parameters
    # @option opts [SearchRecommendRulesParams] :search_recommend_rules_params 
    # @return [SearchRecommendRulesResponse]
    def search_recommend_rules(index_name, model, opts = {})
      data, _status_code, _headers = search_recommend_rules_with_http_info(index_name, model, opts)
      data
    end

    # List Recommend rules.
    # List [Recommend rules](https://www.algolia.com/doc/guides/algolia-recommend/how-to/rules/).
    # @param index_name [String] Index on which to perform the request.
    # @param model [RecommendModels] [Recommend models](https://www.algolia.com/doc/guides/algolia-recommend/overview/#recommend-models). 
    # @param [Hash] opts the optional parameters
    # @option opts [SearchRecommendRulesParams] :search_recommend_rules_params 
    # @return [Array<(SearchRecommendRulesResponse, Integer, Hash)>] SearchRecommendRulesResponse data, response status code and response headers
    def search_recommend_rules_with_http_info(index_name, model, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: RecommendClient.search_recommend_rules ...'
      end
      # verify the required parameter 'index_name' is set
      if @api_client.config.client_side_validation && index_name.nil?
        fail ArgumentError, "Missing the required parameter 'index_name' when calling RecommendClient.search_recommend_rules"
      end
      # verify the required parameter 'model' is set
      if @api_client.config.client_side_validation && model.nil?
        fail ArgumentError, "Missing the required parameter 'model' when calling RecommendClient.search_recommend_rules"
      end
      # resource path
      local_var_path = '/1/indexes/{indexName}/{model}/recommend/rules/search'.sub('{' + 'indexName' + '}', CGI.escape(index_name.to_s)).sub('{' + 'model' + '}', CGI.escape(model.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
        header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'search_recommend_rules_params'])

      # return_type
      return_type = opts[:debug_return_type] || 'SearchRecommendRulesResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apiKey', 'appId']

      new_options = opts.merge(
        :operation => :"RecommendClient.search_recommend_rules",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: RecommendClient#search_recommend_rules\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
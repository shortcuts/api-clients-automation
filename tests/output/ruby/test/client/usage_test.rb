# Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
require 'algolia'
require 'test/unit'

class TestClientUsageClient < Test::Unit::TestCase
  include Algolia::Usage
  # calls api with correct read host
  def test_api0
    client = Algolia::UsageClient.create(
      'test-app-id',
      'test-api-key',
      { requester: Algolia::Transport::EchoRequester.new }
    )
    req = client.custom_get_with_http_info("test")
    assert_equal('test-app-id-dsn.algolia.net', req.host.url)
  end

  # calls api with correct write host
  def test_api1
    client = Algolia::UsageClient.create(
      'test-app-id',
      'test-api-key',
      { requester: Algolia::Transport::EchoRequester.new }
    )
    req = client.custom_post_with_http_info("test")
    assert_equal('test-app-id.algolia.net', req.host.url)
  end

  # calls api with correct user agent
  def test_common_api0
    client = Algolia::UsageClient.create(
      'APP_ID',
      'API_KEY',
      { requester: Algolia::Transport::EchoRequester.new }
    )
    req = client.custom_post_with_http_info("1/test")
    assert(req.headers['user-agent'].match(/^Algolia for Ruby \(\d+\.\d+\.\d+(-?.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-?.*)?\))?)*(; Usage (\(\d+\.\d+\.\d+(-?.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-?.*)?\))?)*$/))
  end

  # calls api with default read timeouts
  def test_common_api1
    client = Algolia::UsageClient.create(
      'APP_ID',
      'API_KEY',
      { requester: Algolia::Transport::EchoRequester.new }
    )
    req = client.custom_get_with_http_info("1/test")
    assert_equal(2000, req.connect_timeout)
    assert_equal(5000, req.timeout)
  end

  # calls api with default write timeouts
  def test_common_api2
    client = Algolia::UsageClient.create(
      'APP_ID',
      'API_KEY',
      { requester: Algolia::Transport::EchoRequester.new }
    )
    req = client.custom_post_with_http_info("1/test")
    assert_equal(2000, req.connect_timeout)
    assert_equal(30_000, req.timeout)
  end

  # client throws with invalid parameters
  def test_parameters0
    begin
      Algolia::UsageClient.create(
        '',
        '',
        { requester: Algolia::Transport::EchoRequester.new }
      )
      assert(false, 'An error should have been raised')
    rescue => e
      assert_equal('`app_id` is missing.', e.message)
    end
    begin
      Algolia::UsageClient.create(
        '',
        'my-api-key',
        { requester: Algolia::Transport::EchoRequester.new }
      )
      assert(false, 'An error should have been raised')
    rescue => e
      assert_equal('`app_id` is missing.', e.message)
    end
    begin
      Algolia::UsageClient.create(
        'my-app-id',
        '',
        { requester: Algolia::Transport::EchoRequester.new }
      )
      assert(false, 'An error should have been raised')
    rescue => e
      assert_equal('`api_key` is missing.', e.message)
    end
  end
end
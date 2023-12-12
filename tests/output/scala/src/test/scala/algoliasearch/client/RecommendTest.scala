package algoliasearch.client

import algoliasearch.api.RecommendClient
import algoliasearch.config.*
import algoliasearch.{EchoInterceptor, assertError}
import algoliasearch.recommend.*
import algoliasearch.exception.*
import org.json4s.*
import org.json4s.native.JsonParser.*
import org.scalatest.funsuite.AnyFunSuite

import scala.concurrent.duration.Duration
import scala.concurrent.{Await, ExecutionContextExecutor}

class RecommendTest extends AnyFunSuite {
  implicit val ec: ExecutionContextExecutor = scala.concurrent.ExecutionContext.global
  implicit val formats: Formats = org.json4s.DefaultFormats

  def testClient(appId: String = "appId", apiKey: String = "apiKey"): (RecommendClient, EchoInterceptor) = {
    val echo = EchoInterceptor()
    (
      RecommendClient(
        appId = appId,
        apiKey = apiKey,
        clientOptions = ClientOptions
          .builder()
          .withRequesterConfig(requester => requester.withInterceptor(echo))
          .build()
      ),
      echo
    )
  }

  test("calls api with correct read host") {

    val (client, echo) = testClient(appId = "test-app-id", apiKey = "test-api-key")

    Await.ready(
      client.get[Any](
        path = "/test"
      ),
      Duration.Inf
    )

    assert(echo.lastResponse.get.host == "test-app-id-dsn.algolia.net")
  }

  test("calls api with correct write host") {

    val (client, echo) = testClient(appId = "test-app-id", apiKey = "test-api-key")

    Await.ready(
      client.post[Any](
        path = "/test"
      ),
      Duration.Inf
    )

    assert(echo.lastResponse.get.host == "test-app-id.algolia.net")
  }

  test("calls api with correct user agent") {
    val (client, echo) = testClient()

    Await.ready(
      client.post[Any](
        path = "/test"
      ),
      Duration.Inf
    )

    val regexp =
      """^Algolia for scala \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Recommend (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$""".r
    val header = echo.lastResponse.get.headers("user-agent")
    assert(header.matches(regexp.regex), s"Expected $header to match the following regex: ${regexp.regex}")
  }

  test("calls api with default read timeouts") {
    val (client, echo) = testClient()

    Await.ready(
      client.get[Any](
        path = "/test"
      ),
      Duration.Inf
    )

    assert(echo.lastResponse.get.connectTimeout == 2000)
    assert(echo.lastResponse.get.responseTimeout == 5000)
  }

  test("calls api with default write timeouts") {
    val (client, echo) = testClient()

    Await.ready(
      client.post[Any](
        path = "/test"
      ),
      Duration.Inf
    )

    assert(echo.lastResponse.get.connectTimeout == 2000)
    assert(echo.lastResponse.get.responseTimeout == 30000)
  }
}
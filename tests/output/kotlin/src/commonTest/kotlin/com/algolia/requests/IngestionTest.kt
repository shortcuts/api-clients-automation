// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package com.algolia.methods.requests

import com.algolia.client.api.IngestionClient
import com.algolia.client.configuration.*
import com.algolia.client.model.ingestion.*
import com.algolia.client.transport.*
import com.algolia.utils.*
import io.ktor.http.*
import kotlinx.coroutines.test.*
import kotlinx.serialization.json.*
import kotlin.test.*

class IngestionTest {

  val client = IngestionClient(
    appId = "appId",
    apiKey = "apiKey",
    region = "us",
  )

  // createAuthentication

  @Test
  fun `createAuthenticationOAuth`() = runTest {
    client.runTest(
      call = {
        createAuthentication(
          authenticationCreate = AuthenticationCreate(
            type = AuthenticationType.entries.first { it.value == "oauth" },
            name = "authName",
            input = AuthOAuth(
              url = "http://test.oauth",
              clientId = "myID",
              clientSecret = "mySecret",
            ),
          ),
        )
      },
      intercept = {
        assertEquals("/1/authentications".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"type":"oauth","name":"authName","input":{"url":"http://test.oauth","client_id":"myID","client_secret":"mySecret"}}""", it.body)
      },
    )
  }

  @Test
  fun `createAuthenticationAlgolia1`() = runTest {
    client.runTest(
      call = {
        createAuthentication(
          authenticationCreate = AuthenticationCreate(
            type = AuthenticationType.entries.first { it.value == "algolia" },
            name = "authName",
            input = AuthAlgolia(
              appID = "myappID",
              apiKey = "randomApiKey",
            ),
          ),
        )
      },
      intercept = {
        assertEquals("/1/authentications".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"type":"algolia","name":"authName","input":{"appID":"myappID","apiKey":"randomApiKey"}}""", it.body)
      },
    )
  }

  // createDestination

  @Test
  fun `createDestination`() = runTest {
    client.runTest(
      call = {
        createDestination(
          destinationCreate = DestinationCreate(
            type = DestinationType.entries.first { it.value == "search" },
            name = "destinationName",
            input = DestinationIndexPrefix(
              indexPrefix = "prefix_",
            ),
            authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          ),
        )
      },
      intercept = {
        assertEquals("/1/destinations".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"type":"search","name":"destinationName","input":{"indexPrefix":"prefix_"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}""", it.body)
      },
    )
  }

  // createSource

  @Test
  fun `createSource`() = runTest {
    client.runTest(
      call = {
        createSource(
          sourceCreate = SourceCreate(
            type = SourceType.entries.first { it.value == "commercetools" },
            name = "sourceName",
            input = SourceCommercetools(
              storeKeys = listOf("myStore"),
              locales = listOf("de"),
              url = "http://commercetools.com",
              projectKey = "keyID",
            ),
            authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          ),
        )
      },
      intercept = {
        assertEquals("/1/sources".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"type":"commercetools","name":"sourceName","input":{"storeKeys":["myStore"],"locales":["de"],"url":"http://commercetools.com","projectKey":"keyID"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}""", it.body)
      },
    )
  }

  // createTask

  @Test
  fun `createTaskOnDemand`() = runTest {
    client.runTest(
      call = {
        createTask(
          taskCreate = TaskCreate(
            sourceID = "search",
            destinationID = "destinationName",
            trigger = OnDemandTriggerInput(
              type = OnDemandTriggerType.entries.first { it.value == "onDemand" },
            ),
            action = ActionType.entries.first { it.value == "replace" },
          ),
        )
      },
      intercept = {
        assertEquals("/1/tasks".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}""", it.body)
      },
    )
  }

  @Test
  fun `createTaskSchedule1`() = runTest {
    client.runTest(
      call = {
        createTask(
          taskCreate = TaskCreate(
            sourceID = "search",
            destinationID = "destinationName",
            trigger = ScheduleTriggerInput(
              type = ScheduleTriggerType.entries.first { it.value == "schedule" },
              cron = "* * * * *",
            ),
            action = ActionType.entries.first { it.value == "replace" },
          ),
        )
      },
      intercept = {
        assertEquals("/1/tasks".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"schedule","cron":"* * * * *"},"action":"replace"}""", it.body)
      },
    )
  }

  @Test
  fun `createTaskSubscription2`() = runTest {
    client.runTest(
      call = {
        createTask(
          taskCreate = TaskCreate(
            sourceID = "search",
            destinationID = "destinationName",
            trigger = OnDemandTriggerInput(
              type = OnDemandTriggerType.entries.first { it.value == "onDemand" },
            ),
            action = ActionType.entries.first { it.value == "replace" },
          ),
        )
      },
      intercept = {
        assertEquals("/1/tasks".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}""", it.body)
      },
    )
  }

  // customDelete

  @Test
  fun `allow del method for a custom path with minimal parameters`() = runTest {
    client.runTest(
      call = {
        customDelete(
          path = "test/minimal",
        )
      },
      intercept = {
        assertEquals("/test/minimal".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  @Test
  fun `allow del method for a custom path with all parameters1`() = runTest {
    client.runTest(
      call = {
        customDelete(
          path = "test/all",
          parameters = mapOf("query" to "parameters"),
        )
      },
      intercept = {
        assertEquals("/test/all".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertQueryParams("""{"query":"parameters"}""", it.url.encodedParameters)
        assertNoBody(it.body)
      },
    )
  }

  // customGet

  @Test
  fun `allow get method for a custom path with minimal parameters`() = runTest {
    client.runTest(
      call = {
        customGet(
          path = "test/minimal",
        )
      },
      intercept = {
        assertEquals("/test/minimal".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  @Test
  fun `allow get method for a custom path with all parameters1`() = runTest {
    client.runTest(
      call = {
        customGet(
          path = "test/all",
          parameters = mapOf("query" to "parameters with space"),
        )
      },
      intercept = {
        assertEquals("/test/all".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertQueryParams("""{"query":"parameters%20with%20space"}""", it.url.encodedParameters)
        assertNoBody(it.body)
      },
    )
  }

  @Test
  fun `requestOptions should be escaped too2`() = runTest {
    client.runTest(
      call = {
        customGet(
          path = "test/all",
          parameters = mapOf("query" to "to be overriden"),
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("query", "parameters with space")
              put("and an array", listOf("array", "with spaces"))
            },
            headers = buildMap {
              put("x-header-1", "spaces are left alone")
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/all".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertContainsAll("""{"x-header-1":"spaces are left alone"}""", it.headers)
        assertQueryParams("""{"query":"parameters%20with%20space","and%20an%20array":"array%2Cwith%20spaces"}""", it.url.encodedParameters)
        assertNoBody(it.body)
      },
    )
  }

  // customPost

  @Test
  fun `allow post method for a custom path with minimal parameters`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/minimal",
        )
      },
      intercept = {
        assertEquals("/test/minimal".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{}""", it.body)
      },
    )
  }

  @Test
  fun `allow post method for a custom path with all parameters1`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/all",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "body",
              JsonPrimitive("parameters"),
            )
          },
        )
      },
      intercept = {
        assertEquals("/test/all".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters"}""", it.url.encodedParameters)
        assertJsonBody("""{"body":"parameters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions can override default query parameters2`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("query", "myQueryParameter")
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"myQueryParameter"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions merges query parameters with default ones3`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("query2", "myQueryParameter")
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","query2":"myQueryParameter"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions can override default headers4`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            headers = buildMap {
              put("x-algolia-api-key", "myApiKey")
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertContainsAll("""{"x-algolia-api-key":"myApiKey"}""", it.headers)
        assertQueryParams("""{"query":"parameters"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions merges headers with default ones5`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            headers = buildMap {
              put("x-algolia-api-key", "myApiKey")
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertContainsAll("""{"x-algolia-api-key":"myApiKey"}""", it.headers)
        assertQueryParams("""{"query":"parameters"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions queryParameters accepts booleans6`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("isItWorking", true)
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","isItWorking":"true"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions queryParameters accepts integers7`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("myParam", 2)
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","myParam":"2"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions queryParameters accepts list of string8`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("myParam", listOf("b and c", "d"))
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","myParam":"b%20and%20c%2Cd"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions queryParameters accepts list of booleans9`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("myParam", listOf(true, true, false))
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","myParam":"true%2Ctrue%2Cfalse"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  @Test
  fun `requestOptions queryParameters accepts list of integers10`() = runTest {
    client.runTest(
      call = {
        customPost(
          path = "test/requestOptions",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "facet",
              JsonPrimitive("filters"),
            )
          },
          requestOptions = RequestOptions(
            urlParameters = buildMap {
              put("myParam", listOf(1, 2))
            },
          ),
        )
      },
      intercept = {
        assertEquals("/test/requestOptions".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertQueryParams("""{"query":"parameters","myParam":"1%2C2"}""", it.url.encodedParameters)
        assertJsonBody("""{"facet":"filters"}""", it.body)
      },
    )
  }

  // customPut

  @Test
  fun `allow put method for a custom path with minimal parameters`() = runTest {
    client.runTest(
      call = {
        customPut(
          path = "test/minimal",
        )
      },
      intercept = {
        assertEquals("/test/minimal".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PUT"), it.method)
        assertJsonBody("""{}""", it.body)
      },
    )
  }

  @Test
  fun `allow put method for a custom path with all parameters1`() = runTest {
    client.runTest(
      call = {
        customPut(
          path = "test/all",
          parameters = mapOf("query" to "parameters"),
          body = buildJsonObject {
            put(
              "body",
              JsonPrimitive("parameters"),
            )
          },
        )
      },
      intercept = {
        assertEquals("/test/all".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PUT"), it.method)
        assertQueryParams("""{"query":"parameters"}""", it.url.encodedParameters)
        assertJsonBody("""{"body":"parameters"}""", it.body)
      },
    )
  }

  // deleteAuthentication

  @Test
  fun `deleteAuthentication`() = runTest {
    client.runTest(
      call = {
        deleteAuthentication(
          authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // deleteDestination

  @Test
  fun `deleteDestination`() = runTest {
    client.runTest(
      call = {
        deleteDestination(
          destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // deleteSource

  @Test
  fun `deleteSource`() = runTest {
    client.runTest(
      call = {
        deleteSource(
          sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // deleteTask

  @Test
  fun `deleteTask`() = runTest {
    client.runTest(
      call = {
        deleteTask(
          taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("DELETE"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // disableTask

  @Test
  fun `disableTask`() = runTest {
    client.runTest(
      call = {
        disableTask(
          taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/disable".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PUT"), it.method)
        assertEmptyBody(it.body)
      },
    )
  }

  // enableTask

  @Test
  fun `enable task e2e`() = runTest {
    client.runTest(
      call = {
        enableTask(
          taskID = "76ab4c2a-ce17-496f-b7a6-506dc59ee498",
        )
      },
      intercept = {
        assertEquals("/1/tasks/76ab4c2a-ce17-496f-b7a6-506dc59ee498/enable".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PUT"), it.method)
        assertEmptyBody(it.body)
      },
    )
  }

  // getAuthentication

  @Test
  fun `getAuthentication`() = runTest {
    client.runTest(
      call = {
        getAuthentication(
          authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getAuthentications

  @Test
  fun `getAuthentications`() = runTest {
    client.runTest(
      call = {
        getAuthentications()
      },
      intercept = {
        assertEquals("/1/authentications".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  @Test
  fun `getAuthentications with query params1`() = runTest {
    client.runTest(
      call = {
        getAuthentications(
          itemsPerPage = 2,
          page = 1,
          type = listOf(AuthenticationType.entries.first { it.value == "basic" }, AuthenticationType.entries.first { it.value == "algolia" }),
          platform = listOf(PlatformNone.entries.first { it.value == "none" }),
          sort = AuthenticationSortKeys.entries.first { it.value == "createdAt" },
          order = OrderKeys.entries.first { it.value == "asc" },
        )
      },
      intercept = {
        assertEquals("/1/authentications".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertQueryParams("""{"itemsPerPage":"2","page":"1","type":"basic%2Calgolia","platform":"none","sort":"createdAt","order":"asc"}""", it.url.encodedParameters)
        assertNoBody(it.body)
      },
    )
  }

  // getDestination

  @Test
  fun `getDestination`() = runTest {
    client.runTest(
      call = {
        getDestination(
          destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getDestinations

  @Test
  fun `getDestinations`() = runTest {
    client.runTest(
      call = {
        getDestinations()
      },
      intercept = {
        assertEquals("/1/destinations".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getEvent

  @Test
  fun `getEvent`() = runTest {
    client.runTest(
      call = {
        getEvent(
          runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          eventID = "6c02aeb1-775e-418e-870b-1faccd4b2c0c",
        )
      },
      intercept = {
        assertEquals("/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events/6c02aeb1-775e-418e-870b-1faccd4b2c0c".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getEvents

  @Test
  fun `getEvents`() = runTest {
    client.runTest(
      call = {
        getEvents(
          runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getRun

  @Test
  fun `getRun`() = runTest {
    client.runTest(
      call = {
        getRun(
          runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getRuns

  @Test
  fun `getRuns`() = runTest {
    client.runTest(
      call = {
        getRuns()
      },
      intercept = {
        assertEquals("/1/runs".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getSource

  @Test
  fun `getSource`() = runTest {
    client.runTest(
      call = {
        getSource(
          sourceID = "75eeb306-51d3-4e5e-a279-3c92bd8893ac",
        )
      },
      intercept = {
        assertEquals("/1/sources/75eeb306-51d3-4e5e-a279-3c92bd8893ac".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getSources

  @Test
  fun `getSources`() = runTest {
    client.runTest(
      call = {
        getSources()
      },
      intercept = {
        assertEquals("/1/sources".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getTask

  @Test
  fun `getTask`() = runTest {
    client.runTest(
      call = {
        getTask(
          taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // getTasks

  @Test
  fun `getTasks`() = runTest {
    client.runTest(
      call = {
        getTasks()
      },
      intercept = {
        assertEquals("/1/tasks".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("GET"), it.method)
        assertNoBody(it.body)
      },
    )
  }

  // runTask

  @Test
  fun `runTask`() = runTest {
    client.runTest(
      call = {
        runTask(
          taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/run".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertEmptyBody(it.body)
      },
    )
  }

  // searchAuthentications

  @Test
  fun `searchAuthentications`() = runTest {
    client.runTest(
      call = {
        searchAuthentications(
          authenticationSearch = AuthenticationSearch(
            authenticationIDs = listOf("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a"),
          ),
        )
      },
      intercept = {
        assertEquals("/1/authentications/search".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"authenticationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""", it.body)
      },
    )
  }

  // searchDestinations

  @Test
  fun `searchDestinations`() = runTest {
    client.runTest(
      call = {
        searchDestinations(
          destinationSearch = DestinationSearch(
            destinationIDs = listOf("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a"),
          ),
        )
      },
      intercept = {
        assertEquals("/1/destinations/search".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"destinationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""", it.body)
      },
    )
  }

  // searchSources

  @Test
  fun `searchSources`() = runTest {
    client.runTest(
      call = {
        searchSources(
          sourceSearch = SourceSearch(
            sourceIDs = listOf("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a"),
          ),
        )
      },
      intercept = {
        assertEquals("/1/sources/search".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"sourceIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""", it.body)
      },
    )
  }

  // searchTasks

  @Test
  fun `searchTasks`() = runTest {
    client.runTest(
      call = {
        searchTasks(
          taskSearch = TaskSearch(
            taskIDs = listOf("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a", "76ab4c2a-ce17-496f-b7a6-506dc59ee498"),
          ),
        )
      },
      intercept = {
        assertEquals("/1/tasks/search".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"taskIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a","76ab4c2a-ce17-496f-b7a6-506dc59ee498"]}""", it.body)
      },
    )
  }

  // triggerDockerSourceDiscover

  @Test
  fun `triggerDockerSourceDiscover`() = runTest {
    client.runTest(
      call = {
        triggerDockerSourceDiscover(
          sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )
      },
      intercept = {
        assertEquals("/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/discover".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertEmptyBody(it.body)
      },
    )
  }

  // updateAuthentication

  @Test
  fun `updateAuthentication`() = runTest {
    client.runTest(
      call = {
        updateAuthentication(
          authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          authenticationUpdate = AuthenticationUpdate(
            name = "newName",
          ),
        )
      },
      intercept = {
        assertEquals("/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PATCH"), it.method)
        assertJsonBody("""{"name":"newName"}""", it.body)
      },
    )
  }

  // updateDestination

  @Test
  fun `updateDestination`() = runTest {
    client.runTest(
      call = {
        updateDestination(
          destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          destinationUpdate = DestinationUpdate(
            name = "newName",
          ),
        )
      },
      intercept = {
        assertEquals("/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PATCH"), it.method)
        assertJsonBody("""{"name":"newName"}""", it.body)
      },
    )
  }

  // updateSource

  @Test
  fun `updateSource`() = runTest {
    client.runTest(
      call = {
        updateSource(
          sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          sourceUpdate = SourceUpdate(
            name = "newName",
          ),
        )
      },
      intercept = {
        assertEquals("/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PATCH"), it.method)
        assertJsonBody("""{"name":"newName"}""", it.body)
      },
    )
  }

  // updateTask

  @Test
  fun `updateTask`() = runTest {
    client.runTest(
      call = {
        updateTask(
          taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          taskUpdate = TaskUpdate(
            enabled = false,
          ),
        )
      },
      intercept = {
        assertEquals("/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("PATCH"), it.method)
        assertJsonBody("""{"enabled":false}""", it.body)
      },
    )
  }

  // validateSource

  @Test
  fun `validateSource`() = runTest {
    client.runTest(
      call = {
        validateSource(
          sourceCreate = SourceCreate(
            type = SourceType.entries.first { it.value == "commercetools" },
            name = "sourceName",
            input = SourceCommercetools(
              storeKeys = listOf("myStore"),
              locales = listOf("de"),
              url = "http://commercetools.com",
              projectKey = "keyID",
            ),
            authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          ),
        )
      },
      intercept = {
        assertEquals("/1/sources/validate".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"type":"commercetools","name":"sourceName","input":{"storeKeys":["myStore"],"locales":["de"],"url":"http://commercetools.com","projectKey":"keyID"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}""", it.body)
      },
    )
  }

  // validateSourceBeforeUpdate

  @Test
  fun `validateSourceBeforeUpdate`() = runTest {
    client.runTest(
      call = {
        validateSourceBeforeUpdate(
          sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          sourceUpdate = SourceUpdate(
            name = "newName",
          ),
        )
      },
      intercept = {
        assertEquals("/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/validate".toPathSegments(), it.url.pathSegments)
        assertEquals(HttpMethod.parse("POST"), it.method)
        assertJsonBody("""{"name":"newName"}""", it.body)
      },
    )
  }
}

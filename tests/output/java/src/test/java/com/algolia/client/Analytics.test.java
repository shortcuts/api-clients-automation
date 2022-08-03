package com.algolia.client;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import com.algolia.EchoInterceptor;
import com.algolia.EchoResponse;
import com.algolia.api.AnalyticsClient;
import com.algolia.model.analytics.*;
import com.algolia.utils.ClientOptions;
import com.algolia.utils.HttpRequester;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class AnalyticsClientClientTests {

  private HttpRequester requester;
  private EchoInterceptor echo;

  @BeforeAll
  void init() {
    requester = new HttpRequester();
    echo = new EchoInterceptor();
    requester.addInterceptor(echo.getEchoInterceptor());
  }

  AnalyticsClient createClient() {
    return new AnalyticsClient("appId", "apiKey", "us", new ClientOptions().setRequester(requester));
  }

  @Test
  @DisplayName("calls api with correct user agent")
  void commonApiTest0() {
    AnalyticsClient $client = createClient();

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    {
      String regexp =
        "^Algolia for Java \\(\\d+\\.\\d+\\.\\d+(-.*)?\\)(; [a-zA-Z. ]+" +
        " (\\(\\d+((\\.\\d+)?\\.\\d+)?(-.*)?\\))?)*(; Analytics" +
        " (\\(\\d+\\.\\d+\\.\\d+(-.*)?\\)))(; [a-zA-Z. ]+" +
        " (\\(\\d+((\\.\\d+)?\\.\\d+)?(-.*)?\\))?)*$";
      assertTrue(
        result.headers.get("user-agent").matches(regexp),
        "Expected " + result.headers.get("user-agent") + " to match the following regex: " + regexp
      );
    }
  }

  @Test
  @DisplayName("calls api with default read timeouts")
  void commonApiTest1() {
    AnalyticsClient $client = createClient();

    String path0 = "/test";

    $client.get(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(5000, result.responseTimeout);
  }

  @Test
  @DisplayName("calls api with default write timeouts")
  void commonApiTest2() {
    AnalyticsClient $client = createClient();

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(30000, result.responseTimeout);
  }

  @Test
  @DisplayName("fallbacks to the alias when region is not given")
  void parametersTest0() {
    AnalyticsClient $client = new AnalyticsClient("my-app-id", "my-api-key", new ClientOptions().setRequester(requester));

    String index0 = "my-index";

    $client.getAverageClickPosition(index0);
    EchoResponse result = echo.getLastResponse();

    assertEquals("analytics.algolia.com", result.host);
  }

  @Test
  @DisplayName("uses the correct region")
  void parametersTest1() {
    AnalyticsClient $client = new AnalyticsClient("my-app-id", "my-api-key", "de", new ClientOptions().setRequester(requester));

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals("analytics.de.algolia.com", result.host);
  }

  @Test
  @DisplayName("throws when incorrect region is given")
  void parametersTest2() {
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          AnalyticsClient $client = new AnalyticsClient(
            "my-app-id",
            "my-api-key",
            "not_a_region",
            new ClientOptions().setRequester(requester)
          );
        }
      );
      assertEquals("`region` must be one of the following: de, us", exception.getMessage());
    }
  }

  @Test
  @DisplayName("getAverageClickPosition throws without index")
  void parametersTest3() {
    AnalyticsClient $client = createClient();

    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          String index0 = null;

          $client.getClickPositions(index0);
          EchoResponse result = echo.getLastResponse();
        }
      );
      assertEquals("Parameter `index` is required when calling `getClickPositions`.", exception.getMessage());
    }
  }
}
package com.algolia.utils.echo;

import com.algolia.Pair;
import com.algolia.model.abtesting.*;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import okhttp3.HttpUrl;
import okhttp3.Request;
import okio.Buffer;

public class EchoResponseAbtesting {

  private static String parseRequestBody(Request req) {
    try {
      final Request copy = req.newBuilder().build();
      final Buffer buffer = new Buffer();
      copy.body().writeTo(buffer);
      return buffer.readUtf8();
    } catch (final IOException e) {
      return "error";
    }
  }

  private static List<Pair> buildQueryParams(Request req) {
    List<Pair> params = new ArrayList<Pair>();
    HttpUrl url = req.url();
    for (String name : url.queryParameterNames()) {
      for (String value : url.queryParameterValues(name)) {
        params.add(new Pair(name, value));
      }
    }
    return params;
  }

  public static class AddABTests
    extends ABTestResponse
    implements EchoResponseInterface {

    private Request request;

    public AddABTests(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class Del extends Object implements EchoResponseInterface {

    private Request request;

    public Del(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class DeleteABTest
    extends ABTestResponse
    implements EchoResponseInterface {

    private Request request;

    public DeleteABTest(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class Get extends Object implements EchoResponseInterface {

    private Request request;

    public Get(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class GetABTest
    extends ABTest
    implements EchoResponseInterface {

    private Request request;

    public GetABTest(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class ListABTests
    extends ListABTestsResponse
    implements EchoResponseInterface {

    private Request request;

    public ListABTests(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class Post extends Object implements EchoResponseInterface {

    private Request request;

    public Post(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class Put extends Object implements EchoResponseInterface {

    private Request request;

    public Put(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }

  public static class StopABTest
    extends ABTestResponse
    implements EchoResponseInterface {

    private Request request;

    public StopABTest(Request request) {
      this.request = request;
    }

    public String getPath() {
      return request.url().encodedPath();
    }

    public String getMethod() {
      return request.method();
    }

    public String getBody() {
      return parseRequestBody(request);
    }

    public List<Pair> getQueryParams() {
      return buildQueryParams(request);
    }

    // to satisfy CompoundType in case it's a parent
    public Object getInsideValue() {
      return null;
    }
  }
}
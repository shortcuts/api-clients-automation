//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using Algolia.Search.Models;
using Algolia.Search.Models.Common;
using Algolia.Search.Serializer;

namespace Algolia.Search.Models.Recommend;

/// <summary>
/// BaseSearchResponse
/// </summary>
[DataContract(Name = "baseSearchResponse")]
[JsonObject(MemberSerialization.OptOut)]
public partial class BaseSearchResponse
{
  /// <summary>
  /// Initializes a new instance of the BaseSearchResponse class.
  /// </summary>
  [JsonConstructor]
  public BaseSearchResponse()
  {
    AdditionalProperties = new Dictionary<string, object>();
  }
  /// <summary>
  /// Initializes a new instance of the BaseSearchResponse class.
  /// </summary>
  /// <param name="hitsPerPage">Number of hits per page. (required) (default to 20).</param>
  /// <param name="nbHits">Number of hits the search query matched. (required).</param>
  /// <param name="nbPages">Number of pages of results for the current query. (required).</param>
  /// <param name="page">Page to retrieve (the first page is &#x60;0&#x60;, not &#x60;1&#x60;). (required) (default to 0).</param>
  /// <param name="processingTimeMS">Time the server took to process the request, in milliseconds. (required).</param>
  public BaseSearchResponse(int hitsPerPage, int nbHits, int nbPages, int page, int processingTimeMS)
  {
    HitsPerPage = hitsPerPage;
    NbHits = nbHits;
    NbPages = nbPages;
    Page = page;
    ProcessingTimeMS = processingTimeMS;
    AdditionalProperties = new Dictionary<string, object>();
  }

  /// <summary>
  /// A/B test ID. This is only included in the response for indices that are part of an A/B test.
  /// </summary>
  /// <value>A/B test ID. This is only included in the response for indices that are part of an A/B test.</value>
  [DataMember(Name = "abTestID")]
  public int? AbTestID { get; set; }

  /// <summary>
  /// Variant ID. This is only included in the response for indices that are part of an A/B test.
  /// </summary>
  /// <value>Variant ID. This is only included in the response for indices that are part of an A/B test.</value>
  [DataMember(Name = "abTestVariantID")]
  public int? AbTestVariantID { get; set; }

  /// <summary>
  /// Computed geographical location.
  /// </summary>
  /// <value>Computed geographical location.</value>
  [DataMember(Name = "aroundLatLng")]
  public string AroundLatLng { get; set; }

  /// <summary>
  /// Automatically-computed radius.
  /// </summary>
  /// <value>Automatically-computed radius.</value>
  [DataMember(Name = "automaticRadius")]
  public string AutomaticRadius { get; set; }

  /// <summary>
  /// Gets or Sets Exhaustive
  /// </summary>
  [DataMember(Name = "exhaustive")]
  public Exhaustive Exhaustive { get; set; }

  /// <summary>
  /// See the `facetsCount` field of the `exhaustive` object in the response.
  /// </summary>
  /// <value>See the `facetsCount` field of the `exhaustive` object in the response.</value>
  [DataMember(Name = "exhaustiveFacetsCount")]
  [Obsolete]
  public bool? ExhaustiveFacetsCount { get; set; }

  /// <summary>
  /// See the `nbHits` field of the `exhaustive` object in the response.
  /// </summary>
  /// <value>See the `nbHits` field of the `exhaustive` object in the response.</value>
  [DataMember(Name = "exhaustiveNbHits")]
  [Obsolete]
  public bool? ExhaustiveNbHits { get; set; }

  /// <summary>
  /// See the `typo` field of the `exhaustive` object in the response.
  /// </summary>
  /// <value>See the `typo` field of the `exhaustive` object in the response.</value>
  [DataMember(Name = "exhaustiveTypo")]
  [Obsolete]
  public bool? ExhaustiveTypo { get; set; }

  /// <summary>
  /// Mapping of each facet name to the corresponding facet counts.
  /// </summary>
  /// <value>Mapping of each facet name to the corresponding facet counts.</value>
  [DataMember(Name = "facets")]
  public Dictionary<string, Dictionary<string, int>> Facets { get; set; }

  /// <summary>
  /// Statistics for numerical facets.
  /// </summary>
  /// <value>Statistics for numerical facets.</value>
  [DataMember(Name = "facets_stats")]
  public Dictionary<string, FacetsStats> FacetsStats { get; set; }

  /// <summary>
  /// Number of hits per page.
  /// </summary>
  /// <value>Number of hits per page.</value>
  [DataMember(Name = "hitsPerPage")]
  public int HitsPerPage { get; set; }

  /// <summary>
  /// Index name used for the query.
  /// </summary>
  /// <value>Index name used for the query.</value>
  [DataMember(Name = "index")]
  public string Index { get; set; }

  /// <summary>
  /// Index name used for the query. During A/B testing, the targeted index isn't always the index used by the query.
  /// </summary>
  /// <value>Index name used for the query. During A/B testing, the targeted index isn't always the index used by the query.</value>
  [DataMember(Name = "indexUsed")]
  public string IndexUsed { get; set; }

  /// <summary>
  /// Warnings about the query.
  /// </summary>
  /// <value>Warnings about the query.</value>
  [DataMember(Name = "message")]
  public string Message { get; set; }

  /// <summary>
  /// Number of hits the search query matched.
  /// </summary>
  /// <value>Number of hits the search query matched.</value>
  [DataMember(Name = "nbHits")]
  public int NbHits { get; set; }

  /// <summary>
  /// Number of pages of results for the current query.
  /// </summary>
  /// <value>Number of pages of results for the current query.</value>
  [DataMember(Name = "nbPages")]
  public int NbPages { get; set; }

  /// <summary>
  /// Number of hits selected and sorted by the relevant sort algorithm.
  /// </summary>
  /// <value>Number of hits selected and sorted by the relevant sort algorithm.</value>
  [DataMember(Name = "nbSortedHits")]
  public int? NbSortedHits { get; set; }

  /// <summary>
  /// Page to retrieve (the first page is `0`, not `1`).
  /// </summary>
  /// <value>Page to retrieve (the first page is `0`, not `1`).</value>
  [DataMember(Name = "page")]
  public int Page { get; set; }

  /// <summary>
  /// Post-[normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean) query string that will be searched.
  /// </summary>
  /// <value>Post-[normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean) query string that will be searched.</value>
  [DataMember(Name = "parsedQuery")]
  public string ParsedQuery { get; set; }

  /// <summary>
  /// Time the server took to process the request, in milliseconds.
  /// </summary>
  /// <value>Time the server took to process the request, in milliseconds.</value>
  [DataMember(Name = "processingTimeMS")]
  public int ProcessingTimeMS { get; set; }

  /// <summary>
  /// Experimental. List of processing steps and their times, in milliseconds. You can use this list to investigate performance issues.
  /// </summary>
  /// <value>Experimental. List of processing steps and their times, in milliseconds. You can use this list to investigate performance issues.</value>
  [DataMember(Name = "processingTimingsMS")]
  public object ProcessingTimingsMS { get; set; }

  /// <summary>
  /// Markup text indicating which parts of the original query have been removed to retrieve a non-empty result set.
  /// </summary>
  /// <value>Markup text indicating which parts of the original query have been removed to retrieve a non-empty result set.</value>
  [DataMember(Name = "queryAfterRemoval")]
  public string QueryAfterRemoval { get; set; }

  /// <summary>
  /// Gets or Sets Redirect
  /// </summary>
  [DataMember(Name = "redirect")]
  public Redirect Redirect { get; set; }

  /// <summary>
  /// Gets or Sets RenderingContent
  /// </summary>
  [DataMember(Name = "renderingContent")]
  public RenderingContent RenderingContent { get; set; }

  /// <summary>
  /// Time the server took to process the request, in milliseconds.
  /// </summary>
  /// <value>Time the server took to process the request, in milliseconds.</value>
  [DataMember(Name = "serverTimeMS")]
  public int? ServerTimeMS { get; set; }

  /// <summary>
  /// Host name of the server that processed the request.
  /// </summary>
  /// <value>Host name of the server that processed the request.</value>
  [DataMember(Name = "serverUsed")]
  public string ServerUsed { get; set; }

  /// <summary>
  /// Lets you store custom data in your indices.
  /// </summary>
  /// <value>Lets you store custom data in your indices.</value>
  [DataMember(Name = "userData")]
  public object UserData { get; set; }

  /// <summary>
  /// Gets or Sets additional properties
  /// </summary>
  [JsonExtensionData]
  public IDictionary<string, object> AdditionalProperties { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class BaseSearchResponse {\n");
    sb.Append("  AbTestID: ").Append(AbTestID).Append("\n");
    sb.Append("  AbTestVariantID: ").Append(AbTestVariantID).Append("\n");
    sb.Append("  AroundLatLng: ").Append(AroundLatLng).Append("\n");
    sb.Append("  AutomaticRadius: ").Append(AutomaticRadius).Append("\n");
    sb.Append("  Exhaustive: ").Append(Exhaustive).Append("\n");
    sb.Append("  ExhaustiveFacetsCount: ").Append(ExhaustiveFacetsCount).Append("\n");
    sb.Append("  ExhaustiveNbHits: ").Append(ExhaustiveNbHits).Append("\n");
    sb.Append("  ExhaustiveTypo: ").Append(ExhaustiveTypo).Append("\n");
    sb.Append("  Facets: ").Append(Facets).Append("\n");
    sb.Append("  FacetsStats: ").Append(FacetsStats).Append("\n");
    sb.Append("  HitsPerPage: ").Append(HitsPerPage).Append("\n");
    sb.Append("  Index: ").Append(Index).Append("\n");
    sb.Append("  IndexUsed: ").Append(IndexUsed).Append("\n");
    sb.Append("  Message: ").Append(Message).Append("\n");
    sb.Append("  NbHits: ").Append(NbHits).Append("\n");
    sb.Append("  NbPages: ").Append(NbPages).Append("\n");
    sb.Append("  NbSortedHits: ").Append(NbSortedHits).Append("\n");
    sb.Append("  Page: ").Append(Page).Append("\n");
    sb.Append("  ParsedQuery: ").Append(ParsedQuery).Append("\n");
    sb.Append("  ProcessingTimeMS: ").Append(ProcessingTimeMS).Append("\n");
    sb.Append("  ProcessingTimingsMS: ").Append(ProcessingTimingsMS).Append("\n");
    sb.Append("  QueryAfterRemoval: ").Append(QueryAfterRemoval).Append("\n");
    sb.Append("  Redirect: ").Append(Redirect).Append("\n");
    sb.Append("  RenderingContent: ").Append(RenderingContent).Append("\n");
    sb.Append("  ServerTimeMS: ").Append(ServerTimeMS).Append("\n");
    sb.Append("  ServerUsed: ").Append(ServerUsed).Append("\n");
    sb.Append("  UserData: ").Append(UserData).Append("\n");
    sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public virtual string ToJson()
  {
    return JsonConvert.SerializeObject(this, Formatting.Indented);
  }

}


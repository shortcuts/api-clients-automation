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
/// BaseTrendingItemsQuery
/// </summary>
[DataContract(Name = "baseTrendingItemsQuery")]
[JsonObject(MemberSerialization.OptOut)]
public partial class BaseTrendingItemsQuery
{

  /// <summary>
  /// Gets or Sets Model
  /// </summary>
  [DataMember(Name = "model")]
  public TrendingItemsModel? Model { get; set; }
  /// <summary>
  /// Initializes a new instance of the BaseTrendingItemsQuery class.
  /// </summary>
  public BaseTrendingItemsQuery()
  {
  }

  /// <summary>
  /// Facet name for trending models.
  /// </summary>
  /// <value>Facet name for trending models.</value>
  [DataMember(Name = "facetName")]
  public string FacetName { get; set; }

  /// <summary>
  /// Facet value for trending models.
  /// </summary>
  /// <value>Facet value for trending models.</value>
  [DataMember(Name = "facetValue")]
  public string FacetValue { get; set; }

  /// <summary>
  /// Gets or Sets QueryParameters
  /// </summary>
  [DataMember(Name = "queryParameters")]
  public SearchParamsObject QueryParameters { get; set; }

  /// <summary>
  /// Gets or Sets FallbackParameters
  /// </summary>
  [DataMember(Name = "fallbackParameters")]
  public SearchParamsObject FallbackParameters { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class BaseTrendingItemsQuery {\n");
    sb.Append("  FacetName: ").Append(FacetName).Append("\n");
    sb.Append("  FacetValue: ").Append(FacetValue).Append("\n");
    sb.Append("  Model: ").Append(Model).Append("\n");
    sb.Append("  QueryParameters: ").Append(QueryParameters).Append("\n");
    sb.Append("  FallbackParameters: ").Append(FallbackParameters).Append("\n");
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


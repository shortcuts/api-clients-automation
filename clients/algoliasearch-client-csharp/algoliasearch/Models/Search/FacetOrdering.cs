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

namespace Algolia.Search.Models.Search;

/// <summary>
/// Defines the ordering of facets (widgets).
/// </summary>
[DataContract(Name = "facetOrdering")]
[JsonObject(MemberSerialization.OptOut)]
public partial class FacetOrdering
{
  /// <summary>
  /// Initializes a new instance of the FacetOrdering class.
  /// </summary>
  public FacetOrdering()
  {
  }

  /// <summary>
  /// Gets or Sets Facets
  /// </summary>
  [DataMember(Name = "facets")]
  public Facets Facets { get; set; }

  /// <summary>
  /// Ordering of facet values within an individual facet.
  /// </summary>
  /// <value>Ordering of facet values within an individual facet.</value>
  [DataMember(Name = "values")]
  public Dictionary<string, Value> Values { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class FacetOrdering {\n");
    sb.Append("  Facets: ").Append(Facets).Append("\n");
    sb.Append("  Values: ").Append(Values).Append("\n");
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


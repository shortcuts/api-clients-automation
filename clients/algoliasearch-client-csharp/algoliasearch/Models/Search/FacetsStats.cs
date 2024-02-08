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
/// FacetsStats
/// </summary>
[DataContract(Name = "facetsStats")]
[JsonObject(MemberSerialization.OptOut)]
public partial class FacetsStats
{
  /// <summary>
  /// Initializes a new instance of the FacetsStats class.
  /// </summary>
  public FacetsStats()
  {
  }

  /// <summary>
  /// Minimum value in the results.
  /// </summary>
  /// <value>Minimum value in the results.</value>
  [DataMember(Name = "min")]
  public double? Min { get; set; }

  /// <summary>
  /// Maximum value in the results.
  /// </summary>
  /// <value>Maximum value in the results.</value>
  [DataMember(Name = "max")]
  public double? Max { get; set; }

  /// <summary>
  /// Average facet value in the results.
  /// </summary>
  /// <value>Average facet value in the results.</value>
  [DataMember(Name = "avg")]
  public double? Avg { get; set; }

  /// <summary>
  /// Sum of all values in the results.
  /// </summary>
  /// <value>Sum of all values in the results.</value>
  [DataMember(Name = "sum")]
  public double? Sum { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class FacetsStats {\n");
    sb.Append("  Min: ").Append(Min).Append("\n");
    sb.Append("  Max: ").Append(Max).Append("\n");
    sb.Append("  Avg: ").Append(Avg).Append("\n");
    sb.Append("  Sum: ").Append(Sum).Append("\n");
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

//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Text;
using System.Linq;
using System.Text.Json.Serialization;
using System.Collections.Generic;
using Algolia.Search.Serializer;
using System.Text.Json;

namespace Algolia.Search.Models.Recommend;

/// <summary>
/// TrendingItems
/// </summary>
public partial class TrendingItems
{

  /// <summary>
  /// Gets or Sets Model
  /// </summary>
  [JsonPropertyName("model")]
  public TrendingItemsModel? Model { get; set; }
  /// <summary>
  /// Initializes a new instance of the TrendingItems class.
  /// </summary>
  [JsonConstructor]
  public TrendingItems() { }
  /// <summary>
  /// Initializes a new instance of the TrendingItems class.
  /// </summary>
  /// <param name="facetName">Facet attribute. To be used in combination with &#x60;facetValue&#x60;. If specified, only recommendations matching the facet filter will be returned.  (required).</param>
  /// <param name="facetValue">Facet value. To be used in combination with &#x60;facetName&#x60;. If specified, only recommendations matching the facet filter will be returned.  (required).</param>
  /// <param name="model">model (required).</param>
  public TrendingItems(string facetName, string facetValue, TrendingItemsModel? model)
  {
    FacetName = facetName ?? throw new ArgumentNullException(nameof(facetName));
    FacetValue = facetValue ?? throw new ArgumentNullException(nameof(facetValue));
    Model = model;
  }

  /// <summary>
  /// Facet attribute. To be used in combination with `facetValue`. If specified, only recommendations matching the facet filter will be returned. 
  /// </summary>
  /// <value>Facet attribute. To be used in combination with `facetValue`. If specified, only recommendations matching the facet filter will be returned. </value>
  [JsonPropertyName("facetName")]
  public string FacetName { get; set; }

  /// <summary>
  /// Facet value. To be used in combination with `facetName`. If specified, only recommendations matching the facet filter will be returned. 
  /// </summary>
  /// <value>Facet value. To be used in combination with `facetName`. If specified, only recommendations matching the facet filter will be returned. </value>
  [JsonPropertyName("facetValue")]
  public string FacetValue { get; set; }

  /// <summary>
  /// Gets or Sets FallbackParameters
  /// </summary>
  [JsonPropertyName("fallbackParameters")]
  public SearchParamsObject FallbackParameters { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class TrendingItems {\n");
    sb.Append("  FacetName: ").Append(FacetName).Append("\n");
    sb.Append("  FacetValue: ").Append(FacetValue).Append("\n");
    sb.Append("  Model: ").Append(Model).Append("\n");
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
    return JsonSerializer.Serialize(this, JsonConfig.Options);
  }

  /// <summary>
  /// Returns true if objects are equal
  /// </summary>
  /// <param name="obj">Object to be compared</param>
  /// <returns>Boolean</returns>
  public override bool Equals(object obj)
  {
    if (obj is not TrendingItems input)
    {
      return false;
    }

    return
        (FacetName == input.FacetName || (FacetName != null && FacetName.Equals(input.FacetName))) &&
        (FacetValue == input.FacetValue || (FacetValue != null && FacetValue.Equals(input.FacetValue))) &&
        (Model == input.Model || Model.Equals(input.Model)) &&
        (FallbackParameters == input.FallbackParameters || (FallbackParameters != null && FallbackParameters.Equals(input.FallbackParameters)));
  }

  /// <summary>
  /// Gets the hash code
  /// </summary>
  /// <returns>Hash code</returns>
  public override int GetHashCode()
  {
    unchecked // Overflow is fine, just wrap
    {
      int hashCode = 41;
      if (FacetName != null)
      {
        hashCode = (hashCode * 59) + FacetName.GetHashCode();
      }
      if (FacetValue != null)
      {
        hashCode = (hashCode * 59) + FacetValue.GetHashCode();
      }
      hashCode = (hashCode * 59) + Model.GetHashCode();
      if (FallbackParameters != null)
      {
        hashCode = (hashCode * 59) + FallbackParameters.GetHashCode();
      }
      return hashCode;
    }
  }

}

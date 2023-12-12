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
using System.ComponentModel.DataAnnotations;
using FileParameter = Algolia.Search.Analytics.Client.FileParameter;
using OpenAPIDateConverter = Algolia.Search.Analytics.Client.OpenAPIDateConverter;

namespace Algolia.Search.Analytics.Models
{
  /// <summary>
  /// SearchNoResultEvent
  /// </summary>
  [DataContract(Name = "searchNoResultEvent")]
  public partial class SearchNoResultEvent : IEquatable<SearchNoResultEvent>, IValidatableObject
  {
    /// <summary>
    /// Initializes a new instance of the <see cref="SearchNoResultEvent" /> class.
    /// </summary>
    [JsonConstructorAttribute]
    protected SearchNoResultEvent() { }
    /// <summary>
    /// Initializes a new instance of the <see cref="SearchNoResultEvent" /> class.
    /// </summary>
    /// <param name="search">User query. (required).</param>
    /// <param name="count">Number of occurrences. (required).</param>
    /// <param name="nbHits">Number of hits the search query matched. (required).</param>
    public SearchNoResultEvent(string search = default(string), int count = default(int), int nbHits = default(int))
    {
      // to ensure "search" is required (not null)
      if (search == null)
      {
        throw new ArgumentNullException("search is a required property for SearchNoResultEvent and cannot be null");
      }
      this.Search = search;
      this.Count = count;
      this.NbHits = nbHits;
    }

    /// <summary>
    /// User query.
    /// </summary>
    /// <value>User query.</value>
    [DataMember(Name = "search", IsRequired = true, EmitDefaultValue = true)]
    public string Search { get; set; }

    /// <summary>
    /// Number of occurrences.
    /// </summary>
    /// <value>Number of occurrences.</value>
    [DataMember(Name = "count", IsRequired = true, EmitDefaultValue = true)]
    public int Count { get; set; }

    /// <summary>
    /// Number of hits the search query matched.
    /// </summary>
    /// <value>Number of hits the search query matched.</value>
    [DataMember(Name = "nbHits", IsRequired = true, EmitDefaultValue = true)]
    public int NbHits { get; set; }

    /// <summary>
    /// Returns the string presentation of the object
    /// </summary>
    /// <returns>String presentation of the object</returns>
    public override string ToString()
    {
      StringBuilder sb = new StringBuilder();
      sb.Append("class SearchNoResultEvent {\n");
      sb.Append("  Search: ").Append(Search).Append("\n");
      sb.Append("  Count: ").Append(Count).Append("\n");
      sb.Append("  NbHits: ").Append(NbHits).Append("\n");
      sb.Append("}\n");
      return sb.ToString();
    }

    /// <summary>
    /// Returns the JSON string presentation of the object
    /// </summary>
    /// <returns>JSON string presentation of the object</returns>
    public virtual string ToJson()
    {
      return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
    }

    /// <summary>
    /// Returns true if objects are equal
    /// </summary>
    /// <param name="input">Object to be compared</param>
    /// <returns>Boolean</returns>
    public override bool Equals(object input)
    {
      return this.Equals(input as SearchNoResultEvent);
    }

    /// <summary>
    /// Returns true if SearchNoResultEvent instances are equal
    /// </summary>
    /// <param name="input">Instance of SearchNoResultEvent to be compared</param>
    /// <returns>Boolean</returns>
    public bool Equals(SearchNoResultEvent input)
    {
      if (input == null)
      {
        return false;
      }
      return
          (
              this.Search == input.Search ||
              (this.Search != null &&
              this.Search.Equals(input.Search))
          ) &&
          (
              this.Count == input.Count ||
              this.Count.Equals(input.Count)
          ) &&
          (
              this.NbHits == input.NbHits ||
              this.NbHits.Equals(input.NbHits)
          );
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
        if (this.Search != null)
        {
          hashCode = (hashCode * 59) + this.Search.GetHashCode();
        }
        hashCode = (hashCode * 59) + this.Count.GetHashCode();
        hashCode = (hashCode * 59) + this.NbHits.GetHashCode();
        return hashCode;
      }
    }

    /// <summary>
    /// To validate all properties of the instance
    /// </summary>
    /// <param name="validationContext">Validation context</param>
    /// <returns>Validation Result</returns>
    IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
    {
      yield break;
    }
  }

}
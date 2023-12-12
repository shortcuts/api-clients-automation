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
  /// GetUsersCountResponse
  /// </summary>
  [DataContract(Name = "getUsersCountResponse")]
  public partial class GetUsersCountResponse : IEquatable<GetUsersCountResponse>, IValidatableObject
  {
    /// <summary>
    /// Initializes a new instance of the <see cref="GetUsersCountResponse" /> class.
    /// </summary>
    [JsonConstructorAttribute]
    protected GetUsersCountResponse() { }
    /// <summary>
    /// Initializes a new instance of the <see cref="GetUsersCountResponse" /> class.
    /// </summary>
    /// <param name="count">Number of occurrences. (required).</param>
    /// <param name="dates">User count. (required).</param>
    public GetUsersCountResponse(int count = default(int), List<UserWithDate> dates = default(List<UserWithDate>))
    {
      this.Count = count;
      // to ensure "dates" is required (not null)
      if (dates == null)
      {
        throw new ArgumentNullException("dates is a required property for GetUsersCountResponse and cannot be null");
      }
      this.Dates = dates;
    }

    /// <summary>
    /// Number of occurrences.
    /// </summary>
    /// <value>Number of occurrences.</value>
    [DataMember(Name = "count", IsRequired = true, EmitDefaultValue = true)]
    public int Count { get; set; }

    /// <summary>
    /// User count.
    /// </summary>
    /// <value>User count.</value>
    [DataMember(Name = "dates", IsRequired = true, EmitDefaultValue = true)]
    public List<UserWithDate> Dates { get; set; }

    /// <summary>
    /// Returns the string presentation of the object
    /// </summary>
    /// <returns>String presentation of the object</returns>
    public override string ToString()
    {
      StringBuilder sb = new StringBuilder();
      sb.Append("class GetUsersCountResponse {\n");
      sb.Append("  Count: ").Append(Count).Append("\n");
      sb.Append("  Dates: ").Append(Dates).Append("\n");
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
      return this.Equals(input as GetUsersCountResponse);
    }

    /// <summary>
    /// Returns true if GetUsersCountResponse instances are equal
    /// </summary>
    /// <param name="input">Instance of GetUsersCountResponse to be compared</param>
    /// <returns>Boolean</returns>
    public bool Equals(GetUsersCountResponse input)
    {
      if (input == null)
      {
        return false;
      }
      return
          (
              this.Count == input.Count ||
              this.Count.Equals(input.Count)
          ) &&
          (
              this.Dates == input.Dates ||
              this.Dates != null &&
              input.Dates != null &&
              this.Dates.SequenceEqual(input.Dates)
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
        hashCode = (hashCode * 59) + this.Count.GetHashCode();
        if (this.Dates != null)
        {
          hashCode = (hashCode * 59) + this.Dates.GetHashCode();
        }
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
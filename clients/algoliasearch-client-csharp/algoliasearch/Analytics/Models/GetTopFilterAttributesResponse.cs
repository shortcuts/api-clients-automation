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
  /// GetTopFilterAttributesResponse
  /// </summary>
  [DataContract(Name = "getTopFilterAttributesResponse")]
  public partial class GetTopFilterAttributesResponse : IEquatable<GetTopFilterAttributesResponse>, IValidatableObject
  {
    /// <summary>
    /// Initializes a new instance of the <see cref="GetTopFilterAttributesResponse" /> class.
    /// </summary>
    [JsonConstructorAttribute]
    protected GetTopFilterAttributesResponse() { }
    /// <summary>
    /// Initializes a new instance of the <see cref="GetTopFilterAttributesResponse" /> class.
    /// </summary>
    /// <param name="attributes">Filterable attributes. (required).</param>
    public GetTopFilterAttributesResponse(List<GetTopFilterAttribute> attributes = default(List<GetTopFilterAttribute>))
    {
      // to ensure "attributes" is required (not null)
      if (attributes == null)
      {
        throw new ArgumentNullException("attributes is a required property for GetTopFilterAttributesResponse and cannot be null");
      }
      this.Attributes = attributes;
    }

    /// <summary>
    /// Filterable attributes.
    /// </summary>
    /// <value>Filterable attributes.</value>
    [DataMember(Name = "attributes", IsRequired = true, EmitDefaultValue = true)]
    public List<GetTopFilterAttribute> Attributes { get; set; }

    /// <summary>
    /// Returns the string presentation of the object
    /// </summary>
    /// <returns>String presentation of the object</returns>
    public override string ToString()
    {
      StringBuilder sb = new StringBuilder();
      sb.Append("class GetTopFilterAttributesResponse {\n");
      sb.Append("  Attributes: ").Append(Attributes).Append("\n");
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
      return this.Equals(input as GetTopFilterAttributesResponse);
    }

    /// <summary>
    /// Returns true if GetTopFilterAttributesResponse instances are equal
    /// </summary>
    /// <param name="input">Instance of GetTopFilterAttributesResponse to be compared</param>
    /// <returns>Boolean</returns>
    public bool Equals(GetTopFilterAttributesResponse input)
    {
      if (input == null)
      {
        return false;
      }
      return
          (
              this.Attributes == input.Attributes ||
              this.Attributes != null &&
              input.Attributes != null &&
              this.Attributes.SequenceEqual(input.Attributes)
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
        if (this.Attributes != null)
        {
          hashCode = (hashCode * 59) + this.Attributes.GetHashCode();
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
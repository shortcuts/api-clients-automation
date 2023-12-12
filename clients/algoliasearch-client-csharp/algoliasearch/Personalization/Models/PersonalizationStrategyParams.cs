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
using FileParameter = Algolia.Search.Personalization.Client.FileParameter;
using OpenAPIDateConverter = Algolia.Search.Personalization.Client.OpenAPIDateConverter;

namespace Algolia.Search.Personalization.Models
{
  /// <summary>
  /// PersonalizationStrategyParams
  /// </summary>
  [DataContract(Name = "personalizationStrategyParams")]
  public partial class PersonalizationStrategyParams : IEquatable<PersonalizationStrategyParams>, IValidatableObject
  {
    /// <summary>
    /// Initializes a new instance of the <see cref="PersonalizationStrategyParams" /> class.
    /// </summary>
    [JsonConstructorAttribute]
    protected PersonalizationStrategyParams() { }
    /// <summary>
    /// Initializes a new instance of the <see cref="PersonalizationStrategyParams" /> class.
    /// </summary>
    /// <param name="eventScoring">Scores associated with the events. (required).</param>
    /// <param name="facetScoring">Scores associated with the facets. (required).</param>
    /// <param name="personalizationImpact">The impact that personalization has on search results: a number between 0 (personalization disabled) and 100 (personalization fully enabled). (required).</param>
    public PersonalizationStrategyParams(List<EventScoring> eventScoring = default(List<EventScoring>), List<FacetScoring> facetScoring = default(List<FacetScoring>), int personalizationImpact = default(int))
    {
      // to ensure "eventScoring" is required (not null)
      if (eventScoring == null)
      {
        throw new ArgumentNullException("eventScoring is a required property for PersonalizationStrategyParams and cannot be null");
      }
      this.EventScoring = eventScoring;
      // to ensure "facetScoring" is required (not null)
      if (facetScoring == null)
      {
        throw new ArgumentNullException("facetScoring is a required property for PersonalizationStrategyParams and cannot be null");
      }
      this.FacetScoring = facetScoring;
      this.PersonalizationImpact = personalizationImpact;
    }

    /// <summary>
    /// Scores associated with the events.
    /// </summary>
    /// <value>Scores associated with the events.</value>
    [DataMember(Name = "eventScoring", IsRequired = true, EmitDefaultValue = true)]
    public List<EventScoring> EventScoring { get; set; }

    /// <summary>
    /// Scores associated with the facets.
    /// </summary>
    /// <value>Scores associated with the facets.</value>
    [DataMember(Name = "facetScoring", IsRequired = true, EmitDefaultValue = true)]
    public List<FacetScoring> FacetScoring { get; set; }

    /// <summary>
    /// The impact that personalization has on search results: a number between 0 (personalization disabled) and 100 (personalization fully enabled).
    /// </summary>
    /// <value>The impact that personalization has on search results: a number between 0 (personalization disabled) and 100 (personalization fully enabled).</value>
    [DataMember(Name = "personalizationImpact", IsRequired = true, EmitDefaultValue = true)]
    public int PersonalizationImpact { get; set; }

    /// <summary>
    /// Returns the string presentation of the object
    /// </summary>
    /// <returns>String presentation of the object</returns>
    public override string ToString()
    {
      StringBuilder sb = new StringBuilder();
      sb.Append("class PersonalizationStrategyParams {\n");
      sb.Append("  EventScoring: ").Append(EventScoring).Append("\n");
      sb.Append("  FacetScoring: ").Append(FacetScoring).Append("\n");
      sb.Append("  PersonalizationImpact: ").Append(PersonalizationImpact).Append("\n");
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
      return this.Equals(input as PersonalizationStrategyParams);
    }

    /// <summary>
    /// Returns true if PersonalizationStrategyParams instances are equal
    /// </summary>
    /// <param name="input">Instance of PersonalizationStrategyParams to be compared</param>
    /// <returns>Boolean</returns>
    public bool Equals(PersonalizationStrategyParams input)
    {
      if (input == null)
      {
        return false;
      }
      return
          (
              this.EventScoring == input.EventScoring ||
              this.EventScoring != null &&
              input.EventScoring != null &&
              this.EventScoring.SequenceEqual(input.EventScoring)
          ) &&
          (
              this.FacetScoring == input.FacetScoring ||
              this.FacetScoring != null &&
              input.FacetScoring != null &&
              this.FacetScoring.SequenceEqual(input.FacetScoring)
          ) &&
          (
              this.PersonalizationImpact == input.PersonalizationImpact ||
              this.PersonalizationImpact.Equals(input.PersonalizationImpact)
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
        if (this.EventScoring != null)
        {
          hashCode = (hashCode * 59) + this.EventScoring.GetHashCode();
        }
        if (this.FacetScoring != null)
        {
          hashCode = (hashCode * 59) + this.FacetScoring.GetHashCode();
        }
        hashCode = (hashCode * 59) + this.PersonalizationImpact.GetHashCode();
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
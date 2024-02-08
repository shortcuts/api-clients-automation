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

namespace Algolia.Search.Models.Ingestion;

/// <summary>
/// Payload to partially update an Authentication.
/// </summary>
[DataContract(Name = "AuthenticationUpdate")]
[JsonObject(MemberSerialization.OptOut)]
public partial class AuthenticationUpdate
{

  /// <summary>
  /// Gets or Sets Type
  /// </summary>
  [DataMember(Name = "type")]
  public AuthenticationType? Type { get; set; }

  /// <summary>
  /// Gets or Sets Platform
  /// </summary>
  [DataMember(Name = "platform")]
  public Platform? Platform { get; set; }
  /// <summary>
  /// Initializes a new instance of the AuthenticationUpdate class.
  /// </summary>
  public AuthenticationUpdate()
  {
  }

  /// <summary>
  /// An human readable name describing the object.
  /// </summary>
  /// <value>An human readable name describing the object.</value>
  [DataMember(Name = "name")]
  public string Name { get; set; }

  /// <summary>
  /// Gets or Sets Input
  /// </summary>
  [DataMember(Name = "input")]
  public AuthInputPartial Input { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class AuthenticationUpdate {\n");
    sb.Append("  Type: ").Append(Type).Append("\n");
    sb.Append("  Name: ").Append(Name).Append("\n");
    sb.Append("  Platform: ").Append(Platform).Append("\n");
    sb.Append("  Input: ").Append(Input).Append("\n");
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

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
/// Batch parameters.
/// </summary>
[DataContract(Name = "batchWriteParams")]
[JsonObject(MemberSerialization.OptOut)]
public partial class BatchWriteParams
{
  /// <summary>
  /// Initializes a new instance of the BatchWriteParams class.
  /// </summary>
  [JsonConstructor]
  public BatchWriteParams() { }
  /// <summary>
  /// Initializes a new instance of the BatchWriteParams class.
  /// </summary>
  /// <param name="requests">requests (required).</param>
  public BatchWriteParams(List<BatchRequest> requests)
  {
    Requests = requests ?? throw new ArgumentNullException(nameof(requests));
  }

  /// <summary>
  /// Gets or Sets Requests
  /// </summary>
  [DataMember(Name = "requests")]
  public List<BatchRequest> Requests { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class BatchWriteParams {\n");
    sb.Append("  Requests: ").Append(Requests).Append("\n");
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

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

namespace Algolia.Search.Models.Monitoring;

/// <summary>
/// Server
/// </summary>
[DataContract(Name = "Server")]
[JsonObject(MemberSerialization.OptOut)]
public partial class Server
{

  /// <summary>
  /// Gets or Sets Region
  /// </summary>
  [DataMember(Name = "region")]
  public Region? Region { get; set; }

  /// <summary>
  /// Gets or Sets Status
  /// </summary>
  [DataMember(Name = "status")]
  public ServerStatus? Status { get; set; }

  /// <summary>
  /// Gets or Sets Type
  /// </summary>
  [DataMember(Name = "type")]
  public Type? Type { get; set; }
  /// <summary>
  /// Initializes a new instance of the Server class.
  /// </summary>
  public Server()
  {
  }

  /// <summary>
  /// Server name.
  /// </summary>
  /// <value>Server name.</value>
  [DataMember(Name = "name")]
  public string Name { get; set; }

  /// <summary>
  /// Included to support legacy applications. Do not rely on this attribute being present in the response. Use `is_replica` instead. 
  /// </summary>
  /// <value>Included to support legacy applications. Do not rely on this attribute being present in the response. Use `is_replica` instead. </value>
  [DataMember(Name = "is_slave")]
  [Obsolete]
  public bool? IsSlave { get; set; }

  /// <summary>
  /// Indicates whether this server is a replica of another server.
  /// </summary>
  /// <value>Indicates whether this server is a replica of another server.</value>
  [DataMember(Name = "is_replica")]
  public bool? IsReplica { get; set; }

  /// <summary>
  /// Name of the cluster to which this server belongs.
  /// </summary>
  /// <value>Name of the cluster to which this server belongs.</value>
  [DataMember(Name = "cluster")]
  public string Cluster { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class Server {\n");
    sb.Append("  Name: ").Append(Name).Append("\n");
    sb.Append("  Region: ").Append(Region).Append("\n");
    sb.Append("  IsSlave: ").Append(IsSlave).Append("\n");
    sb.Append("  IsReplica: ").Append(IsReplica).Append("\n");
    sb.Append("  Cluster: ").Append(Cluster).Append("\n");
    sb.Append("  Status: ").Append(Status).Append("\n");
    sb.Append("  Type: ").Append(Type).Append("\n");
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


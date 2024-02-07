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
/// `searchDictionaryEntries` parameters. 
/// </summary>
[DataContract(Name = "searchDictionaryEntriesParams")]
[JsonObject(MemberSerialization.OptOut)]
public partial class SearchDictionaryEntriesParams
{
  /// <summary>
  /// Initializes a new instance of the SearchDictionaryEntriesParams class.
  /// </summary>
  [JsonConstructor]
  public SearchDictionaryEntriesParams() { }
  /// <summary>
  /// Initializes a new instance of the SearchDictionaryEntriesParams class.
  /// </summary>
  /// <param name="query">Text to search for in an index. (required) (default to &quot;&quot;).</param>
  public SearchDictionaryEntriesParams(string query)
  {
    Query = query ?? throw new ArgumentNullException(nameof(query));
  }

  /// <summary>
  /// Text to search for in an index.
  /// </summary>
  /// <value>Text to search for in an index.</value>
  [DataMember(Name = "query")]
  public string Query { get; set; }

  /// <summary>
  /// Page to retrieve (the first page is `0`, not `1`).
  /// </summary>
  /// <value>Page to retrieve (the first page is `0`, not `1`).</value>
  [DataMember(Name = "page")]
  public int? Page { get; set; }

  /// <summary>
  /// Number of hits per page.
  /// </summary>
  /// <value>Number of hits per page.</value>
  [DataMember(Name = "hitsPerPage")]
  public int? HitsPerPage { get; set; }

  /// <summary>
  /// [Supported language ISO code](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/). 
  /// </summary>
  /// <value>[Supported language ISO code](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/). </value>
  [DataMember(Name = "language")]
  public string Language { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class SearchDictionaryEntriesParams {\n");
    sb.Append("  Query: ").Append(Query).Append("\n");
    sb.Append("  Page: ").Append(Page).Append("\n");
    sb.Append("  HitsPerPage: ").Append(HitsPerPage).Append("\n");
    sb.Append("  Language: ").Append(Language).Append("\n");
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


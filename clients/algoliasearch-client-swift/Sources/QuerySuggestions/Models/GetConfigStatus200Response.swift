// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct GetConfigStatus200Response: Codable, JSONEncodable, Hashable {

  /** Query Suggestions index name. */
  public var indexName: String?
  /** Indicates whether the creation or update of the Query Suggestions is in progress. */
  public var isRunning: Bool?
  /** Timestamp in [ISO-8601](https://wikipedia.org/wiki/ISO_8601) format when the Query Suggestions index was last built. */
  public var lastBuiltAt: String?
  /** Timestamp in [ISO-8601](https://wikipedia.org/wiki/ISO_8601) format when the Query Suggestions index was last updated successfully. */
  public var lastSuccessfulBuiltAt: String?
  /** Duration of the last successful build in seconds. */
  public var lastSuccessfulBuildDuration: String?

  public init(
    indexName: String? = nil, isRunning: Bool? = nil, lastBuiltAt: String? = nil,
    lastSuccessfulBuiltAt: String? = nil, lastSuccessfulBuildDuration: String? = nil
  ) {
    self.indexName = indexName
    self.isRunning = isRunning
    self.lastBuiltAt = lastBuiltAt
    self.lastSuccessfulBuiltAt = lastSuccessfulBuiltAt
    self.lastSuccessfulBuildDuration = lastSuccessfulBuildDuration
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case indexName
    case isRunning
    case lastBuiltAt
    case lastSuccessfulBuiltAt
    case lastSuccessfulBuildDuration
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encodeIfPresent(indexName, forKey: .indexName)
    try container.encodeIfPresent(isRunning, forKey: .isRunning)
    try container.encodeIfPresent(lastBuiltAt, forKey: .lastBuiltAt)
    try container.encodeIfPresent(lastSuccessfulBuiltAt, forKey: .lastSuccessfulBuiltAt)
    try container.encodeIfPresent(lastSuccessfulBuildDuration, forKey: .lastSuccessfulBuildDuration)
  }
}
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct UpdateApiKeyResponse: Codable, JSONEncodable, Hashable {

  /** API key. */
  public var key: String
  /** Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. */
  public var updatedAt: String

  public init(key: String, updatedAt: String) {
    self.key = key
    self.updatedAt = updatedAt
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case key
    case updatedAt
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(key, forKey: .key)
    try container.encode(updatedAt, forKey: .updatedAt)
  }
}

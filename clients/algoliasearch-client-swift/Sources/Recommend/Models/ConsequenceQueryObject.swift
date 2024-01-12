// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct ConsequenceQueryObject: Codable, JSONEncodable, Hashable {

  /** Words to remove. */
  public var remove: [String]?
  /** Edits to apply. */
  public var edits: [Edit]?

  public init(remove: [String]? = nil, edits: [Edit]? = nil) {
    self.remove = remove
    self.edits = edits
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case remove
    case edits
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encodeIfPresent(remove, forKey: .remove)
    try container.encodeIfPresent(edits, forKey: .edits)
  }
}

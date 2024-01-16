// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

/// Redirect rule data.
public struct RedirectRuleIndexMetadataData: Codable, JSONEncodable, Hashable {

  public var ruleObjectID: String

  public init(ruleObjectID: String) {
    self.ruleObjectID = ruleObjectID
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case ruleObjectID
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(ruleObjectID, forKey: .ruleObjectID)
  }
}
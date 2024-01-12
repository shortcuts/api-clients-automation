// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct GetLogsResponse: Codable, JSONEncodable, Hashable {

  public var logs: [Log]

  public init(logs: [Log]) {
    self.logs = logs
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case logs
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(logs, forKey: .logs)
  }
}

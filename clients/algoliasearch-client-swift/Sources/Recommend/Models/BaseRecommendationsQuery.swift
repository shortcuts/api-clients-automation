// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct BaseRecommendationsQuery: Codable, JSONEncodable, Hashable {

  public var model: RecommendationModels
  /** Unique object identifier. */
  public var objectID: String
  public var queryParameters: SearchParamsObject?
  public var fallbackParameters: SearchParamsObject?

  public init(
    model: RecommendationModels, objectID: String, queryParameters: SearchParamsObject? = nil,
    fallbackParameters: SearchParamsObject? = nil
  ) {
    self.model = model
    self.objectID = objectID
    self.queryParameters = queryParameters
    self.fallbackParameters = fallbackParameters
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case model
    case objectID
    case queryParameters
    case fallbackParameters
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(model, forKey: .model)
    try container.encode(objectID, forKey: .objectID)
    try container.encodeIfPresent(queryParameters, forKey: .queryParameters)
    try container.encodeIfPresent(fallbackParameters, forKey: .fallbackParameters)
  }
}

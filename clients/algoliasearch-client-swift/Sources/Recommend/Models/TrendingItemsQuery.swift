// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

public struct TrendingItemsQuery: Codable, JSONEncodable, Hashable {

  static let thresholdRule = NumericRule<Int>(
    minimum: 0, exclusiveMinimum: false, maximum: 100, exclusiveMaximum: false, multipleOf: nil)
  /** Algolia index name. */
  public var indexName: String
  /** Recommendations with a confidence score lower than `threshold` won't appear in results. > **Note**: Each recommendation has a confidence score of 0 to 100. The closer the score is to 100, the more relevant the recommendations are.  */
  public var threshold: Int?
  /** Maximum number of recommendations to retrieve. If 0, all recommendations will be returned. */
  public var maxRecommendations: Int? = 0
  /** Facet name for trending models. */
  public var facetName: String?
  /** Facet value for trending models. */
  public var facetValue: String?
  public var model: TrendingItemsModel?
  public var queryParameters: SearchParamsObject?
  public var fallbackParameters: SearchParamsObject?

  public init(
    indexName: String, threshold: Int? = nil, maxRecommendations: Int? = 0,
    facetName: String? = nil, facetValue: String? = nil, model: TrendingItemsModel? = nil,
    queryParameters: SearchParamsObject? = nil, fallbackParameters: SearchParamsObject? = nil
  ) {
    self.indexName = indexName
    self.threshold = threshold
    self.maxRecommendations = maxRecommendations
    self.facetName = facetName
    self.facetValue = facetValue
    self.model = model
    self.queryParameters = queryParameters
    self.fallbackParameters = fallbackParameters
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case indexName
    case threshold
    case maxRecommendations
    case facetName
    case facetValue
    case model
    case queryParameters
    case fallbackParameters
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(indexName, forKey: .indexName)
    try container.encodeIfPresent(threshold, forKey: .threshold)
    try container.encodeIfPresent(maxRecommendations, forKey: .maxRecommendations)
    try container.encodeIfPresent(facetName, forKey: .facetName)
    try container.encodeIfPresent(facetValue, forKey: .facetValue)
    try container.encodeIfPresent(model, forKey: .model)
    try container.encodeIfPresent(queryParameters, forKey: .queryParameters)
    try container.encodeIfPresent(fallbackParameters, forKey: .fallbackParameters)
  }
}
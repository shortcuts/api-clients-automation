// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

/// Parameters to apply to this search.  You can use all search parameters, plus special
/// &#x60;automaticFacetFilters&#x60;, &#x60;automaticOptionalFacetFilters&#x60;, and &#x60;query&#x60;.
public struct RecommendParams: Codable, JSONEncodable {
    public var query: RecommendConsequenceQuery?
    public var automaticFacetFilters: RecommendAutomaticFacetFilters?
    public var automaticOptionalFacetFilters: RecommendAutomaticFacetFilters?
    public var renderingContent: RecommendRenderingContent?

    public init(
        query: RecommendConsequenceQuery? = nil,
        automaticFacetFilters: RecommendAutomaticFacetFilters? = nil,
        automaticOptionalFacetFilters: RecommendAutomaticFacetFilters? = nil,
        renderingContent: RecommendRenderingContent? = nil
    ) {
        self.query = query
        self.automaticFacetFilters = automaticFacetFilters
        self.automaticOptionalFacetFilters = automaticOptionalFacetFilters
        self.renderingContent = renderingContent
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case query
        case automaticFacetFilters
        case automaticOptionalFacetFilters
        case renderingContent
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.query, forKey: .query)
        try container.encodeIfPresent(self.automaticFacetFilters, forKey: .automaticFacetFilters)
        try container.encodeIfPresent(self.automaticOptionalFacetFilters, forKey: .automaticOptionalFacetFilters)
        try container.encodeIfPresent(self.renderingContent, forKey: .renderingContent)
    }
}

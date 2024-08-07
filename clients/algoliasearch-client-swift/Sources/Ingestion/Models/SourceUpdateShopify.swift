// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct SourceUpdateShopify: Codable, JSONEncodable {
    /// Whether to index collection IDs.   If your store has `has_collection_search_page` set to true, collection IDs
    /// will be indexed even if `collectionIDIndexing` is false.
    public var collectionIDIndexing: Bool?
    /// Whether to increase the number of indexed collections per product. If true, Algolia indexes 200 collections per
    /// product. If false, 100 collections per product are indexed.
    public var increaseProductCollectionLimit: Bool?
    /// Whether to set the default price ratio to 1 if no sale price is present.  The price ratio is determined by the
    /// ratio: `sale_price` / `regular_price`. If no sale price is present, the price ratio would be 0. If
    /// `defaultPriceRatioAsOne` is true, the price ratio is indexed as 1 instead.
    public var defaultPriceRatioAsOne: Bool?
    /// Whether to exclude out-of-stock variants when determining the `max_variant_price` and `min_variant_price`
    /// attributes.
    public var excludeOOSVariantsForPriceAtTRS: Bool?
    /// Whether to include an inventory with every variant for every product record.
    public var includeVariantsInventory: Bool?
    /// Whether to include collection IDs and handles in the product records.
    public var hasCollectionSearchPage: Bool?
    /// Whether to convert tags on products to named tags.  To learn more, see [Named
    /// tags](https://www.algolia.com/doc/integration/shopify/sending-and-managing-data/named-tags).
    public var productNamedTags: Bool?

    public init(
        collectionIDIndexing: Bool? = nil,
        increaseProductCollectionLimit: Bool? = nil,
        defaultPriceRatioAsOne: Bool? = nil,
        excludeOOSVariantsForPriceAtTRS: Bool? = nil,
        includeVariantsInventory: Bool? = nil,
        hasCollectionSearchPage: Bool? = nil,
        productNamedTags: Bool? = nil
    ) {
        self.collectionIDIndexing = collectionIDIndexing
        self.increaseProductCollectionLimit = increaseProductCollectionLimit
        self.defaultPriceRatioAsOne = defaultPriceRatioAsOne
        self.excludeOOSVariantsForPriceAtTRS = excludeOOSVariantsForPriceAtTRS
        self.includeVariantsInventory = includeVariantsInventory
        self.hasCollectionSearchPage = hasCollectionSearchPage
        self.productNamedTags = productNamedTags
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case collectionIDIndexing
        case increaseProductCollectionLimit
        case defaultPriceRatioAsOne
        case excludeOOSVariantsForPriceAtTRS
        case includeVariantsInventory
        case hasCollectionSearchPage
        case productNamedTags
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.collectionIDIndexing, forKey: .collectionIDIndexing)
        try container.encodeIfPresent(self.increaseProductCollectionLimit, forKey: .increaseProductCollectionLimit)
        try container.encodeIfPresent(self.defaultPriceRatioAsOne, forKey: .defaultPriceRatioAsOne)
        try container.encodeIfPresent(self.excludeOOSVariantsForPriceAtTRS, forKey: .excludeOOSVariantsForPriceAtTRS)
        try container.encodeIfPresent(self.includeVariantsInventory, forKey: .includeVariantsInventory)
        try container.encodeIfPresent(self.hasCollectionSearchPage, forKey: .hasCollectionSearchPage)
        try container.encodeIfPresent(self.productNamedTags, forKey: .productNamedTags)
    }
}

extension SourceUpdateShopify: Equatable {
    public static func ==(lhs: SourceUpdateShopify, rhs: SourceUpdateShopify) -> Bool {
        lhs.collectionIDIndexing == rhs.collectionIDIndexing &&
            lhs.increaseProductCollectionLimit == rhs.increaseProductCollectionLimit &&
            lhs.defaultPriceRatioAsOne == rhs.defaultPriceRatioAsOne &&
            lhs.excludeOOSVariantsForPriceAtTRS == rhs.excludeOOSVariantsForPriceAtTRS &&
            lhs.includeVariantsInventory == rhs.includeVariantsInventory &&
            lhs.hasCollectionSearchPage == rhs.hasCollectionSearchPage &&
            lhs.productNamedTags == rhs.productNamedTags
    }
}

extension SourceUpdateShopify: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.collectionIDIndexing?.hashValue)
        hasher.combine(self.increaseProductCollectionLimit?.hashValue)
        hasher.combine(self.defaultPriceRatioAsOne?.hashValue)
        hasher.combine(self.excludeOOSVariantsForPriceAtTRS?.hashValue)
        hasher.combine(self.includeVariantsInventory?.hashValue)
        hasher.combine(self.hasCollectionSearchPage?.hashValue)
        hasher.combine(self.productNamedTags?.hashValue)
    }
}

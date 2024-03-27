// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

/// Treat singular, plurals, and other forms of declensions as equivalent. You should only use this feature for the
/// languages used in your index.
public enum SearchIgnorePlurals: Codable, JSONEncodable, AbstractEncodable {
    case arrayOfSearchSupportedLanguage([SearchSupportedLanguage])
    case bool(Bool)

    public func encode(to encoder: Encoder) throws {
        var container = encoder.singleValueContainer()
        switch self {
        case let .arrayOfSearchSupportedLanguage(value):
            try container.encode(value)
        case let .bool(value):
            try container.encode(value)
        }
    }

    public init(from decoder: Decoder) throws {
        let container = try decoder.singleValueContainer()
        if let value = try? container.decode([SearchSupportedLanguage].self) {
            self = .arrayOfSearchSupportedLanguage(value)
        } else if let value = try? container.decode(Bool.self) {
            self = .bool(value)
        } else {
            throw DecodingError.typeMismatch(
                Self.Type.self,
                .init(
                    codingPath: decoder.codingPath,
                    debugDescription: "Unable to decode instance of SearchIgnorePlurals"
                )
            )
        }
    }

    public func GetActualInstance() -> Encodable {
        switch self {
        case let .arrayOfSearchSupportedLanguage(value):
            value as [SearchSupportedLanguage]
        case let .bool(value):
            value as Bool
        }
    }
}
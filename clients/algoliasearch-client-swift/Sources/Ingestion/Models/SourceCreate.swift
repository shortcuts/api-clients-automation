// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct SourceCreate: Codable, JSONEncodable {
    public var type: SourceType
    public var name: String
    public var input: SourceInput
    /// The authentication UUID.
    public var authenticationID: String?

    public init(type: SourceType, name: String, input: SourceInput, authenticationID: String? = nil) {
        self.type = type
        self.name = name
        self.input = input
        self.authenticationID = authenticationID
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case type
        case name
        case input
        case authenticationID
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.type, forKey: .type)
        try container.encode(self.name, forKey: .name)
        try container.encode(self.input, forKey: .input)
        try container.encodeIfPresent(self.authenticationID, forKey: .authenticationID)
    }
}

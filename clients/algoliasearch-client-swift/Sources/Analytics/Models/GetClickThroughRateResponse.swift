// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct GetClickThroughRateResponse: Codable, JSONEncodable {
    /// Click-through rate, calculated as number of tracked searches with at least one click event divided by the number
    /// of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true.
    public var rate: Double?
    /// Number of clicks associated with this search.
    public var clickCount: Int
    /// Number of tracked searches. Tracked searches are search requests where the `clickAnalytics` parameter is true.
    public var trackedSearchCount: Int
    /// Daily click-through rates.
    public var dates: [DailyClickThroughRates]

    public init(rate: Double?, clickCount: Int, trackedSearchCount: Int, dates: [DailyClickThroughRates]) {
        self.rate = rate
        self.clickCount = clickCount
        self.trackedSearchCount = trackedSearchCount
        self.dates = dates
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case rate
        case clickCount
        case trackedSearchCount
        case dates
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.rate, forKey: .rate)
        try container.encode(self.clickCount, forKey: .clickCount)
        try container.encode(self.trackedSearchCount, forKey: .trackedSearchCount)
        try container.encode(self.dates, forKey: .dates)
    }
}

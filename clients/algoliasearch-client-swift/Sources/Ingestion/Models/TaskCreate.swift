// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation

#if canImport(AnyCodable)
  import AnyCodable
#endif

/// The payload for a task creation.
public struct TaskCreate: Codable, JSONEncodable, Hashable {

  static let failureThresholdRule = NumericRule<Int>(
    minimum: 0, exclusiveMinimum: false, maximum: 100, exclusiveMaximum: false, multipleOf: nil)
  /** The source UUID. */
  public var sourceID: String
  /** The destination UUID. */
  public var destinationID: String
  public var trigger: TaskCreateTrigger
  public var action: ActionType
  /** Whether the task is enabled or not. */
  public var enabled: Bool?
  /** A percentage representing the accepted failure threshold to determine if a `run` succeeded or not. */
  public var failureThreshold: Int?
  public var input: TaskInput?

  public init(
    sourceID: String, destinationID: String, trigger: TaskCreateTrigger, action: ActionType,
    enabled: Bool? = nil, failureThreshold: Int? = nil, input: TaskInput? = nil
  ) {
    self.sourceID = sourceID
    self.destinationID = destinationID
    self.trigger = trigger
    self.action = action
    self.enabled = enabled
    self.failureThreshold = failureThreshold
    self.input = input
  }

  public enum CodingKeys: String, CodingKey, CaseIterable {
    case sourceID
    case destinationID
    case trigger
    case action
    case enabled
    case failureThreshold
    case input
  }

  // Encodable protocol methods

  public func encode(to encoder: Encoder) throws {
    var container = encoder.container(keyedBy: CodingKeys.self)
    try container.encode(sourceID, forKey: .sourceID)
    try container.encode(destinationID, forKey: .destinationID)
    try container.encode(trigger, forKey: .trigger)
    try container.encode(action, forKey: .action)
    try container.encodeIfPresent(enabled, forKey: .enabled)
    try container.encodeIfPresent(failureThreshold, forKey: .failureThreshold)
    try container.encodeIfPresent(input, forKey: .input)
  }
}

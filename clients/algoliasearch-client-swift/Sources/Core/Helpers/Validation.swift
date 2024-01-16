// Validation.swift
//
// Created by Algolia on 08/01/2024
//

import Foundation

public struct StringRule {
  public var minLength: Int?
  public var maxLength: Int?
  public var pattern: String?

  public init(minLength: Int? = nil, maxLength: Int? = nil, pattern: String? = nil) {
    self.minLength = minLength
    self.maxLength = maxLength
    self.pattern = pattern
  }
}

public struct NumericRule<T: Comparable & Numeric> {
  public var minimum: T?
  public var exclusiveMinimum = false
  public var maximum: T?
  public var exclusiveMaximum = false
  public var multipleOf: T?

  public init(
    minimum: T? = nil, exclusiveMinimum: Bool = false, maximum: T? = nil,
    exclusiveMaximum: Bool = false, multipleOf: T? = nil
  ) {
    self.minimum = minimum
    self.exclusiveMinimum = exclusiveMinimum
    self.maximum = maximum
    self.exclusiveMaximum = exclusiveMaximum
    self.multipleOf = multipleOf
  }
}

public enum StringValidationErrorKind: Error {
  case minLength, maxLength, pattern
}

public enum NumericValidationErrorKind: Error {
  case minimum, maximum, multipleOf
}

public struct ValidationError<T: Error & Hashable>: Error {
  public fileprivate(set) var kinds: Set<T>
}

public struct Validator {
  /// Validate a string against a rule.
  /// - Parameter string: The String you wish to validate.
  /// - Parameter rule: The StringRule you wish to use for validation.
  /// - Returns: A validated string.
  /// - Throws: `ValidationError<StringValidationErrorKind>` if the string is invalid against the rule,
  ///           `NSError` if the rule.pattern is invalid.
  public static func validate(_ string: String, against rule: StringRule) throws -> String {
    var error = ValidationError<StringValidationErrorKind>(kinds: [])
    if let minLength = rule.minLength, !(minLength <= string.count) {
      error.kinds.insert(.minLength)
    }
    if let maxLength = rule.maxLength, !(string.count <= maxLength) {
      error.kinds.insert(.maxLength)
    }
    if let pattern = rule.pattern {
      let matches = try NSRegularExpression(pattern: pattern, options: .caseInsensitive)
        .matches(in: string, range: .init(location: 0, length: string.utf16.count))
      if matches.isEmpty {
        error.kinds.insert(.pattern)
      }
    }
    guard error.kinds.isEmpty else {
      throw error
    }
    return string
  }

  /// Validate a integer against a rule.
  /// - Parameter numeric: The integer you wish to validate.
  /// - Parameter rule: The NumericRule you wish to use for validation.
  /// - Returns: A validated integer.
  /// - Throws: `ValidationError<NumericValidationErrorKind>` if the numeric is invalid against the rule.
  public static func validate<T: Comparable & BinaryInteger>(
    _ numeric: T, against rule: NumericRule<T>
  ) throws -> T {
    var error = ValidationError<NumericValidationErrorKind>(kinds: [])
    if let minium = rule.minimum {
      if !rule.exclusiveMinimum, minium > numeric {
        error.kinds.insert(.minimum)
      }
      if rule.exclusiveMinimum, minium >= numeric {
        error.kinds.insert(.minimum)
      }
    }
    if let maximum = rule.maximum {
      if !rule.exclusiveMaximum, numeric > maximum {
        error.kinds.insert(.maximum)
      }
      if rule.exclusiveMaximum, numeric >= maximum {
        error.kinds.insert(.maximum)
      }
    }
    if let multipleOf = rule.multipleOf, !numeric.isMultiple(of: multipleOf) {
      error.kinds.insert(.multipleOf)
    }
    guard error.kinds.isEmpty else {
      throw error
    }
    return numeric
  }

  /// Validate a fractional number against a rule.
  /// - Parameter numeric: The fractional number you wish to validate.
  /// - Parameter rule: The NumericRule you wish to use for validation.
  /// - Returns: A validated fractional number.
  /// - Throws: `ValidationError<NumericValidationErrorKind>` if the numeric is invalid against the rule.
  public static func validate<T: Comparable & FloatingPoint>(
    _ numeric: T, against rule: NumericRule<T>
  ) throws -> T {
    var error = ValidationError<NumericValidationErrorKind>(kinds: [])
    if let minium = rule.minimum {
      if !rule.exclusiveMinimum, minium > numeric {
        error.kinds.insert(.minimum)
      }
      if rule.exclusiveMinimum, minium >= numeric {
        error.kinds.insert(.minimum)
      }
    }
    if let maximum = rule.maximum {
      if !rule.exclusiveMaximum, numeric > maximum {
        error.kinds.insert(.maximum)
      }
      if rule.exclusiveMaximum, numeric >= maximum {
        error.kinds.insert(.maximum)
      }
    }
    if let multipleOf = rule.multipleOf, numeric.remainder(dividingBy: multipleOf) != 0 {
      error.kinds.insert(.multipleOf)
    }
    guard error.kinds.isEmpty else {
      throw error
    }
    return numeric
  }
}
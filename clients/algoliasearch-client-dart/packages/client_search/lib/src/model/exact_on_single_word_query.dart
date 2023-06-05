// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:json_annotation/json_annotation.dart';

/// Controls how the exact ranking criterion is computed when the query contains only one word.
enum ExactOnSingleWordQuery {
  /// Controls how the exact ranking criterion is computed when the query contains only one word.
  @JsonValue(r'attribute')
  attribute,

  /// Controls how the exact ranking criterion is computed when the query contains only one word.
  @JsonValue(r'none')
  none,

  /// Controls how the exact ranking criterion is computed when the query contains only one word.
  @JsonValue(r'word')
  word,
}
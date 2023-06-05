// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search_lite/src/model/match_level.dart';

import 'package:json_annotation/json_annotation.dart';

part 'highlight_result_option.g.dart';

@JsonSerializable()
final class HighlightResultOption {
  /// Returns a new [HighlightResultOption] instance.
  const HighlightResultOption({
    required this.value,
    required this.matchLevel,
    required this.matchedWords,
    this.fullyHighlighted,
  });

  /// Markup text with occurrences highlighted.
  @JsonKey(name: r'value')
  final String value;

  @JsonKey(name: r'matchLevel')
  final MatchLevel matchLevel;

  /// List of words from the query that matched the object.
  @JsonKey(name: r'matchedWords')
  final List<String> matchedWords;

  /// Whether the entire attribute value is highlighted.
  @JsonKey(name: r'fullyHighlighted')
  final bool? fullyHighlighted;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is HighlightResultOption &&
          other.value == value &&
          other.matchLevel == matchLevel &&
          other.matchedWords == matchedWords &&
          other.fullyHighlighted == fullyHighlighted;

  @override
  int get hashCode =>
      value.hashCode +
      matchLevel.hashCode +
      matchedWords.hashCode +
      fullyHighlighted.hashCode;

  factory HighlightResultOption.fromJson(Map<String, dynamic> json) =>
      _$HighlightResultOptionFromJson(json);

  Map<String, dynamic> toJson() => _$HighlightResultOptionToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
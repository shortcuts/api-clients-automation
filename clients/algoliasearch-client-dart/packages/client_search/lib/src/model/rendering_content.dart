// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search/src/model/facet_ordering.dart';

import 'package:json_annotation/json_annotation.dart';

part 'rendering_content.g.dart';

@JsonSerializable()
final class RenderingContent {
  /// Returns a new [RenderingContent] instance.
  const RenderingContent({
    this.facetOrdering,
  });

  @JsonKey(name: r'facetOrdering')
  final FacetOrdering? facetOrdering;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is RenderingContent && other.facetOrdering == facetOrdering;

  @override
  int get hashCode => facetOrdering.hashCode;

  factory RenderingContent.fromJson(Map<String, dynamic> json) =>
      _$RenderingContentFromJson(json);

  Map<String, dynamic> toJson() => _$RenderingContentToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
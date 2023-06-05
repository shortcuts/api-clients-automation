// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search/src/model/redirect_rule_index_metadata.dart';

import 'package:json_annotation/json_annotation.dart';

part 'base_search_response_redirect.g.dart';

@JsonSerializable()
final class BaseSearchResponseRedirect {
  /// Returns a new [BaseSearchResponseRedirect] instance.
  const BaseSearchResponseRedirect({
    this.index,
  });

  @JsonKey(name: r'index')
  final List<RedirectRuleIndexMetadata>? index;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is BaseSearchResponseRedirect && other.index == index;

  @override
  int get hashCode => index.hashCode;

  factory BaseSearchResponseRedirect.fromJson(Map<String, dynamic> json) =>
      _$BaseSearchResponseRedirectFromJson(json);

  Map<String, dynamic> toJson() => _$BaseSearchResponseRedirectToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
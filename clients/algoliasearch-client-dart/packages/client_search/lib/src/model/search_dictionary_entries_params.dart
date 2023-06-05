// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'search_dictionary_entries_params.g.dart';

@JsonSerializable()
final class SearchDictionaryEntriesParams {
  /// Returns a new [SearchDictionaryEntriesParams] instance.
  const SearchDictionaryEntriesParams({
    required this.query,
    this.page,
    this.hitsPerPage,
    this.language,
  });

  /// The text to search in the index.
  @JsonKey(name: r'query')
  final String query;

  /// Specify the page to retrieve.
  @JsonKey(name: r'page')
  final int? page;

  /// Set the number of hits per page.
  @JsonKey(name: r'hitsPerPage')
  final int? hitsPerPage;

  /// Language ISO code supported by the dictionary (e.g., \"en\" for English).
  @JsonKey(name: r'language')
  final String? language;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is SearchDictionaryEntriesParams &&
          other.query == query &&
          other.page == page &&
          other.hitsPerPage == hitsPerPage &&
          other.language == language;

  @override
  int get hashCode =>
      query.hashCode + page.hashCode + hitsPerPage.hashCode + language.hashCode;

  factory SearchDictionaryEntriesParams.fromJson(Map<String, dynamic> json) =>
      _$SearchDictionaryEntriesParamsFromJson(json);

  Map<String, dynamic> toJson() => _$SearchDictionaryEntriesParamsToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
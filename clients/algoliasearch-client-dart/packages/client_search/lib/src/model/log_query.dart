// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'log_query.g.dart';

@JsonSerializable()
final class LogQuery {
  /// Returns a new [LogQuery] instance.
  const LogQuery({
    this.indexName,
    this.userToken,
    this.queryId,
  });

  /// Index targeted by the query.
  @JsonKey(name: r'index_name')
  final String? indexName;

  /// User identifier.
  @JsonKey(name: r'user_token')
  final String? userToken;

  /// QueryID for the given query.
  @JsonKey(name: r'query_id')
  final String? queryId;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is LogQuery &&
          other.indexName == indexName &&
          other.userToken == userToken &&
          other.queryId == queryId;

  @override
  int get hashCode =>
      indexName.hashCode + userToken.hashCode + queryId.hashCode;

  factory LogQuery.fromJson(Map<String, dynamic> json) =>
      _$LogQueryFromJson(json);

  Map<String, dynamic> toJson() => _$LogQueryToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
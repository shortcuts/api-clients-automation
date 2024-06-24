// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/search_params.dart';

import 'package:json_annotation/json_annotation.dart';

part 'base_recommend_request.g.dart';

@JsonSerializable()
final class BaseRecommendRequest {
  /// Returns a new [BaseRecommendRequest] instance.
  const BaseRecommendRequest({
    required this.indexName,
    required this.threshold,
    this.maxRecommendations,
    this.queryParameters,
  });

  /// Index name (case-sensitive).
  @JsonKey(name: r'indexName')
  final String indexName;

  /// Minimum score a recommendation must have to be included in the response.
  // minimum: 0
  // maximum: 100
  @JsonKey(name: r'threshold')
  final double threshold;

  /// Maximum number of recommendations to retrieve. By default, all recommendations are returned and no fallback request is made. Depending on the available recommendations and the other request parameters, the actual number of recommendations may be lower than this value.
  // minimum: 1
  // maximum: 1000
  @JsonKey(name: r'maxRecommendations')
  final int? maxRecommendations;

  @JsonKey(name: r'queryParameters')
  final SearchParams? queryParameters;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is BaseRecommendRequest &&
          other.indexName == indexName &&
          other.threshold == threshold &&
          other.maxRecommendations == maxRecommendations &&
          other.queryParameters == queryParameters;

  @override
  int get hashCode =>
      indexName.hashCode +
      threshold.hashCode +
      maxRecommendations.hashCode +
      queryParameters.hashCode;

  factory BaseRecommendRequest.fromJson(Map<String, dynamic> json) =>
      _$BaseRecommendRequestFromJson(json);

  Map<String, dynamic> toJson() => _$BaseRecommendRequestToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
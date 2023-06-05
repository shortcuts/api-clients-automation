// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search/src/model/user_id.dart';

import 'package:json_annotation/json_annotation.dart';

part 'list_user_ids_response.g.dart';

@JsonSerializable()
final class ListUserIdsResponse {
  /// Returns a new [ListUserIdsResponse] instance.
  const ListUserIdsResponse({
    required this.userIDs,
  });

  /// List of userIDs.
  @JsonKey(name: r'userIDs')
  final List<UserId> userIDs;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ListUserIdsResponse && other.userIDs == userIDs;

  @override
  int get hashCode => userIDs.hashCode;

  factory ListUserIdsResponse.fromJson(Map<String, dynamic> json) =>
      _$ListUserIdsResponseFromJson(json);

  Map<String, dynamic> toJson() => _$ListUserIdsResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
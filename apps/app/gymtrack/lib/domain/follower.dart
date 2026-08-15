import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';
import 'user.dart';

part 'follower.freezed.dart';
part 'follower.g.dart';

@freezed
abstract class FollowerResponse with _$FollowerResponse {
  const factory FollowerResponse({
    required String id,
    @JsonKey(name: 'firstName') required String firstName,
    @JsonKey(name: 'lastName') required String lastName,
    String? email,
    @JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,
    @UserTypeConverter() required UserType type,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
  }) = _FollowerResponse;

  factory FollowerResponse.fromJson(Map<String, dynamic> json) =>
      _$FollowerResponseFromJson(json);
}

@freezed
abstract class CountResponse with _$CountResponse {
  const factory CountResponse({required int count}) = _CountResponse;

  factory CountResponse.fromJson(Map<String, dynamic> json) =>
      _$CountResponseFromJson(json);
}

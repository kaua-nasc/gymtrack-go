import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';

part 'user.freezed.dart';
part 'user.g.dart';

@freezed
abstract class User with _$User {
  const factory User({
    required String id,
    required String firstName,
    required String lastName,
    required String email,
    @JsonKey(name: 'emailVerifiedAt') DateTime? emailVerifiedAt,
    String? bio,
    @JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,
    @JsonKey(name: 'type') required UserType type,
    double? height,
    @JsonKey(name: 'heightUnit') HeightUnit? heightUnit,
    @JsonKey(name: 'currentWeight') double? currentWeight,
    @JsonKey(name: 'weightUnit') WeightUnit? weightUnit,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}

@freezed
abstract class UpdateProfileRequest with _$UpdateProfileRequest {
  const factory UpdateProfileRequest({
    @JsonKey(includeIfNull: false) String? firstName,
    @JsonKey(includeIfNull: false) String? lastName,
    @JsonKey(includeIfNull: false) String? bio,
    @JsonKey(includeIfNull: false) double? height,
    @JsonKey(includeIfNull: false)
    @HeightUnitConverter()
    HeightUnit? heightUnit,
    @JsonKey(includeIfNull: false)
    @WeightUnitConverter()
    WeightUnit? weightUnit,
    @JsonKey(includeIfNull: false, name: 'currentWeight') double? currentWeight,
  }) = _UpdateProfileRequest;

  factory UpdateProfileRequest.fromJson(Map<String, dynamic> json) =>
      _$UpdateProfileRequestFromJson(json);
}

class UserTypeConverter implements JsonConverter<UserType, String> {
  const UserTypeConverter();

  @override
  UserType fromJson(String json) => UserType.fromJson(json);

  @override
  String toJson(UserType object) => object.jsonValue;
}

class WeightUnitConverter implements JsonConverter<WeightUnit, String> {
  const WeightUnitConverter();

  @override
  WeightUnit fromJson(String json) =>
      WeightUnit.values.firstWhere((e) => e.name == json);

  @override
  String toJson(WeightUnit object) => object.name;
}

class HeightUnitConverter implements JsonConverter<HeightUnit, String> {
  const HeightUnitConverter();

  @override
  HeightUnit fromJson(String json) =>
      HeightUnit.values.firstWhere((e) => e.name == json);

  @override
  String toJson(HeightUnit object) => object.name;
}

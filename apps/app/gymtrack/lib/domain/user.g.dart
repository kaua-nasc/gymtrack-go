// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_User _$UserFromJson(Map<String, dynamic> json) => _User(
  id: json['id'] as String,
  firstName: json['firstName'] as String,
  lastName: json['lastName'] as String,
  email: json['email'] as String,
  emailVerifiedAt: json['emailVerifiedAt'] == null
      ? null
      : DateTime.parse(json['emailVerifiedAt'] as String),
  bio: json['bio'] as String?,
  profilePictureUrl: json['profilePictureUrl'] as String?,
  type: $enumDecode(_$UserTypeEnumMap, json['type']),
  height: (json['height'] as num?)?.toDouble(),
  heightUnit: $enumDecodeNullable(_$HeightUnitEnumMap, json['heightUnit']),
  currentWeight: (json['currentWeight'] as num?)?.toDouble(),
  weightUnit: $enumDecodeNullable(_$WeightUnitEnumMap, json['weightUnit']),
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
);

Map<String, dynamic> _$UserToJson(_User instance) => <String, dynamic>{
  'id': instance.id,
  'firstName': instance.firstName,
  'lastName': instance.lastName,
  'email': instance.email,
  'emailVerifiedAt': instance.emailVerifiedAt?.toIso8601String(),
  'bio': instance.bio,
  'profilePictureUrl': instance.profilePictureUrl,
  'type': _$UserTypeEnumMap[instance.type]!,
  'height': instance.height,
  'heightUnit': _$HeightUnitEnumMap[instance.heightUnit],
  'currentWeight': instance.currentWeight,
  'weightUnit': _$WeightUnitEnumMap[instance.weightUnit],
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
};

const _$UserTypeEnumMap = {
  UserType.client: 'client',
  UserType.personalTrainer: 'personalTrainer',
  UserType.admin: 'admin',
};

const _$HeightUnitEnumMap = {HeightUnit.cm: 'cm', HeightUnit.ft: 'ft'};

const _$WeightUnitEnumMap = {WeightUnit.kg: 'kg', WeightUnit.lbs: 'lbs'};

_UpdateProfileRequest _$UpdateProfileRequestFromJson(
  Map<String, dynamic> json,
) => _UpdateProfileRequest(
  firstName: json['firstName'] as String?,
  lastName: json['lastName'] as String?,
  bio: json['bio'] as String?,
  height: (json['height'] as num?)?.toDouble(),
  heightUnit: _$JsonConverterFromJson<String, HeightUnit>(
    json['heightUnit'],
    const HeightUnitConverter().fromJson,
  ),
  weightUnit: _$JsonConverterFromJson<String, WeightUnit>(
    json['weightUnit'],
    const WeightUnitConverter().fromJson,
  ),
  currentWeight: (json['currentWeight'] as num?)?.toDouble(),
);

Map<String, dynamic> _$UpdateProfileRequestToJson(
  _UpdateProfileRequest instance,
) => <String, dynamic>{
  'firstName': ?instance.firstName,
  'lastName': ?instance.lastName,
  'bio': ?instance.bio,
  'height': ?instance.height,
  'heightUnit': ?_$JsonConverterToJson<String, HeightUnit>(
    instance.heightUnit,
    const HeightUnitConverter().toJson,
  ),
  'weightUnit': ?_$JsonConverterToJson<String, WeightUnit>(
    instance.weightUnit,
    const WeightUnitConverter().toJson,
  ),
  'currentWeight': ?instance.currentWeight,
};

Value? _$JsonConverterFromJson<Json, Value>(
  Object? json,
  Value? Function(Json json) fromJson,
) => json == null ? null : fromJson(json as Json);

Json? _$JsonConverterToJson<Json, Value>(
  Value? value,
  Json? Function(Value value) toJson,
) => value == null ? null : toJson(value);

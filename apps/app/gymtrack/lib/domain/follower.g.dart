// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'follower.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_FollowerResponse _$FollowerResponseFromJson(Map<String, dynamic> json) =>
    _FollowerResponse(
      id: json['id'] as String,
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
      email: json['email'] as String?,
      profilePictureUrl: json['profilePictureUrl'] as String?,
      type: const UserTypeConverter().fromJson(json['type'] as String),
      createdAt: DateTime.parse(json['createdAt'] as String),
    );

Map<String, dynamic> _$FollowerResponseToJson(_FollowerResponse instance) =>
    <String, dynamic>{
      'id': instance.id,
      'firstName': instance.firstName,
      'lastName': instance.lastName,
      'email': instance.email,
      'profilePictureUrl': instance.profilePictureUrl,
      'type': const UserTypeConverter().toJson(instance.type),
      'createdAt': instance.createdAt.toIso8601String(),
    };

_CountResponse _$CountResponseFromJson(Map<String, dynamic> json) =>
    _CountResponse(count: (json['count'] as num).toInt());

Map<String, dynamic> _$CountResponseToJson(_CountResponse instance) =>
    <String, dynamic>{'count': instance.count};

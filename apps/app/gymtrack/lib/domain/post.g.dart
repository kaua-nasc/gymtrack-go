// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'post.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_Post _$PostFromJson(Map<String, dynamic> json) => _Post(
  id: json['id'] as String,
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
  authorId: json['authorId'] as String,
  content: json['content'] as String,
  mediaUrls:
      (json['mediaUrls'] as List<dynamic>?)?.map((e) => e as String).toList() ??
      const [],
  entityId: json['entityId'] as String?,
  entityType: $enumDecodeNullable(_$PostEntityTypeEnumMap, json['entityType']),
  status: const PostStatusConverter().fromJson(json['status'] as String),
  rejectedReason: json['rejectedReason'] as String?,
  author: json['author'] == null
      ? null
      : UserSummary.fromJson(json['author'] as Map<String, dynamic>),
  likeCount: (json['likeCount'] as num?)?.toInt() ?? 0,
  commentCount: (json['commentCount'] as num?)?.toInt() ?? 0,
);

Map<String, dynamic> _$PostToJson(_Post instance) => <String, dynamic>{
  'id': instance.id,
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
  'authorId': instance.authorId,
  'content': instance.content,
  'mediaUrls': instance.mediaUrls,
  'entityId': instance.entityId,
  'entityType': _$PostEntityTypeEnumMap[instance.entityType],
  'status': const PostStatusConverter().toJson(instance.status),
  'rejectedReason': instance.rejectedReason,
  'author': instance.author,
  'likeCount': instance.likeCount,
  'commentCount': instance.commentCount,
};

const _$PostEntityTypeEnumMap = {
  PostEntityType.trainingPlan: 'trainingPlan',
  PostEntityType.exercise: 'exercise',
  PostEntityType.achievement: 'achievement',
};

_UserSummary _$UserSummaryFromJson(Map<String, dynamic> json) => _UserSummary(
  id: json['id'] as String,
  firstName: json['firstName'] as String,
  lastName: json['lastName'] as String,
  email: json['email'] as String?,
  profilePictureUrl: json['profilePictureUrl'] as String?,
);

Map<String, dynamic> _$UserSummaryToJson(_UserSummary instance) =>
    <String, dynamic>{
      'id': instance.id,
      'firstName': instance.firstName,
      'lastName': instance.lastName,
      'email': instance.email,
      'profilePictureUrl': instance.profilePictureUrl,
    };

_CreatePostRequest _$CreatePostRequestFromJson(
  Map<String, dynamic> json,
) => _CreatePostRequest(
  content: json['content'] as String,
  mediaUrls:
      (json['mediaUrls'] as List<dynamic>?)?.map((e) => e as String).toList() ??
      const [],
  entityId: json['entityId'] as String?,
  entityType: $enumDecodeNullable(_$PostEntityTypeEnumMap, json['entityType']),
);

Map<String, dynamic> _$CreatePostRequestToJson(_CreatePostRequest instance) =>
    <String, dynamic>{
      'content': instance.content,
      'mediaUrls': instance.mediaUrls,
      'entityId': instance.entityId,
      'entityType': _$PostEntityTypeEnumMap[instance.entityType],
    };

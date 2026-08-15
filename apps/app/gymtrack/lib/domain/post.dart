import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';

part 'post.freezed.dart';
part 'post.g.dart';

@freezed
abstract class Post with _$Post {
  const factory Post({
    required String id,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
    @JsonKey(name: 'authorId') required String authorId,
    required String content,
    @JsonKey(name: 'mediaUrls') @Default([]) List<String> mediaUrls,
    @JsonKey(name: 'entityId') String? entityId,
    @JsonKey(name: 'entityType') PostEntityType? entityType,
    @PostStatusConverter() required PostStatus status,
    @JsonKey(name: 'rejectedReason') String? rejectedReason,
    @JsonKey(name: 'author') UserSummary? author,
    @Default(0) int likeCount,
    @Default(0) int commentCount,
  }) = _Post;

  factory Post.fromJson(Map<String, dynamic> json) => _$PostFromJson(json);
}

@freezed
abstract class UserSummary with _$UserSummary {
  const factory UserSummary({
    required String id,
    @JsonKey(name: 'firstName') required String firstName,
    @JsonKey(name: 'lastName') required String lastName,
    String? email,
    @JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,
  }) = _UserSummary;

  factory UserSummary.fromJson(Map<String, dynamic> json) =>
      _$UserSummaryFromJson(json);
}

@freezed
abstract class CreatePostRequest with _$CreatePostRequest {
  const factory CreatePostRequest({
    required String content,
    @Default([]) List<String> mediaUrls,
    @JsonKey(includeIfNull: true, name: 'entityId') String? entityId,
    @JsonKey(includeIfNull: true, name: 'entityType')
    PostEntityType? entityType,
  }) = _CreatePostRequest;

  factory CreatePostRequest.fromJson(Map<String, dynamic> json) =>
      _$CreatePostRequestFromJson(json);
}

class PostStatusConverter implements JsonConverter<PostStatus, String> {
  const PostStatusConverter();

  @override
  PostStatus fromJson(String json) =>
      PostStatus.values.firstWhere((e) => e.name == json);

  @override
  String toJson(PostStatus object) => object.name;
}

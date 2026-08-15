// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'post.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$Post {

 String get id;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;@JsonKey(name: 'authorId') String get authorId; String get content;@JsonKey(name: 'mediaUrls') List<String> get mediaUrls;@JsonKey(name: 'entityId') String? get entityId;@JsonKey(name: 'entityType') PostEntityType? get entityType;@PostStatusConverter() PostStatus get status;@JsonKey(name: 'rejectedReason') String? get rejectedReason;@JsonKey(name: 'author') UserSummary? get author; int get likeCount; int get commentCount;
/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$PostCopyWith<Post> get copyWith => _$PostCopyWithImpl<Post>(this as Post, _$identity);

  /// Serializes this Post to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is Post&&(identical(other.id, id) || other.id == id)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt)&&(identical(other.authorId, authorId) || other.authorId == authorId)&&(identical(other.content, content) || other.content == content)&&const DeepCollectionEquality().equals(other.mediaUrls, mediaUrls)&&(identical(other.entityId, entityId) || other.entityId == entityId)&&(identical(other.entityType, entityType) || other.entityType == entityType)&&(identical(other.status, status) || other.status == status)&&(identical(other.rejectedReason, rejectedReason) || other.rejectedReason == rejectedReason)&&(identical(other.author, author) || other.author == author)&&(identical(other.likeCount, likeCount) || other.likeCount == likeCount)&&(identical(other.commentCount, commentCount) || other.commentCount == commentCount));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,createdAt,updatedAt,authorId,content,const DeepCollectionEquality().hash(mediaUrls),entityId,entityType,status,rejectedReason,author,likeCount,commentCount);

@override
String toString() {
  return 'Post(id: $id, createdAt: $createdAt, updatedAt: $updatedAt, authorId: $authorId, content: $content, mediaUrls: $mediaUrls, entityId: $entityId, entityType: $entityType, status: $status, rejectedReason: $rejectedReason, author: $author, likeCount: $likeCount, commentCount: $commentCount)';
}


}

/// @nodoc
abstract mixin class $PostCopyWith<$Res>  {
  factory $PostCopyWith(Post value, $Res Function(Post) _then) = _$PostCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt,@JsonKey(name: 'authorId') String authorId, String content,@JsonKey(name: 'mediaUrls') List<String> mediaUrls,@JsonKey(name: 'entityId') String? entityId,@JsonKey(name: 'entityType') PostEntityType? entityType,@PostStatusConverter() PostStatus status,@JsonKey(name: 'rejectedReason') String? rejectedReason,@JsonKey(name: 'author') UserSummary? author, int likeCount, int commentCount
});


$UserSummaryCopyWith<$Res>? get author;

}
/// @nodoc
class _$PostCopyWithImpl<$Res>
    implements $PostCopyWith<$Res> {
  _$PostCopyWithImpl(this._self, this._then);

  final Post _self;
  final $Res Function(Post) _then;

/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? createdAt = null,Object? updatedAt = null,Object? authorId = null,Object? content = null,Object? mediaUrls = null,Object? entityId = freezed,Object? entityType = freezed,Object? status = null,Object? rejectedReason = freezed,Object? author = freezed,Object? likeCount = null,Object? commentCount = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,authorId: null == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String,content: null == content ? _self.content : content // ignore: cast_nullable_to_non_nullable
as String,mediaUrls: null == mediaUrls ? _self.mediaUrls : mediaUrls // ignore: cast_nullable_to_non_nullable
as List<String>,entityId: freezed == entityId ? _self.entityId : entityId // ignore: cast_nullable_to_non_nullable
as String?,entityType: freezed == entityType ? _self.entityType : entityType // ignore: cast_nullable_to_non_nullable
as PostEntityType?,status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PostStatus,rejectedReason: freezed == rejectedReason ? _self.rejectedReason : rejectedReason // ignore: cast_nullable_to_non_nullable
as String?,author: freezed == author ? _self.author : author // ignore: cast_nullable_to_non_nullable
as UserSummary?,likeCount: null == likeCount ? _self.likeCount : likeCount // ignore: cast_nullable_to_non_nullable
as int,commentCount: null == commentCount ? _self.commentCount : commentCount // ignore: cast_nullable_to_non_nullable
as int,
  ));
}
/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@override
@pragma('vm:prefer-inline')
$UserSummaryCopyWith<$Res>? get author {
    if (_self.author == null) {
    return null;
  }

  return $UserSummaryCopyWith<$Res>(_self.author!, (value) {
    return _then(_self.copyWith(author: value));
  });
}
}


/// Adds pattern-matching-related methods to [Post].
extension PostPatterns on Post {
/// A variant of `map` that fallback to returning `orElse`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _Post value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _Post() when $default != null:
return $default(_that);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// Callbacks receives the raw object, upcasted.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case final Subclass2 value:
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _Post value)  $default,){
final _that = this;
switch (_that) {
case _Post():
return $default(_that);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `map` that fallback to returning `null`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _Post value)?  $default,){
final _that = this;
switch (_that) {
case _Post() when $default != null:
return $default(_that);case _:
  return null;

}
}
/// A variant of `when` that fallback to an `orElse` callback.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt, @JsonKey(name: 'authorId')  String authorId,  String content, @JsonKey(name: 'mediaUrls')  List<String> mediaUrls, @JsonKey(name: 'entityId')  String? entityId, @JsonKey(name: 'entityType')  PostEntityType? entityType, @PostStatusConverter()  PostStatus status, @JsonKey(name: 'rejectedReason')  String? rejectedReason, @JsonKey(name: 'author')  UserSummary? author,  int likeCount,  int commentCount)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _Post() when $default != null:
return $default(_that.id,_that.createdAt,_that.updatedAt,_that.authorId,_that.content,_that.mediaUrls,_that.entityId,_that.entityType,_that.status,_that.rejectedReason,_that.author,_that.likeCount,_that.commentCount);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// As opposed to `map`, this offers destructuring.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case Subclass2(:final field2):
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt, @JsonKey(name: 'authorId')  String authorId,  String content, @JsonKey(name: 'mediaUrls')  List<String> mediaUrls, @JsonKey(name: 'entityId')  String? entityId, @JsonKey(name: 'entityType')  PostEntityType? entityType, @PostStatusConverter()  PostStatus status, @JsonKey(name: 'rejectedReason')  String? rejectedReason, @JsonKey(name: 'author')  UserSummary? author,  int likeCount,  int commentCount)  $default,) {final _that = this;
switch (_that) {
case _Post():
return $default(_that.id,_that.createdAt,_that.updatedAt,_that.authorId,_that.content,_that.mediaUrls,_that.entityId,_that.entityType,_that.status,_that.rejectedReason,_that.author,_that.likeCount,_that.commentCount);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `when` that fallback to returning `null`
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt, @JsonKey(name: 'authorId')  String authorId,  String content, @JsonKey(name: 'mediaUrls')  List<String> mediaUrls, @JsonKey(name: 'entityId')  String? entityId, @JsonKey(name: 'entityType')  PostEntityType? entityType, @PostStatusConverter()  PostStatus status, @JsonKey(name: 'rejectedReason')  String? rejectedReason, @JsonKey(name: 'author')  UserSummary? author,  int likeCount,  int commentCount)?  $default,) {final _that = this;
switch (_that) {
case _Post() when $default != null:
return $default(_that.id,_that.createdAt,_that.updatedAt,_that.authorId,_that.content,_that.mediaUrls,_that.entityId,_that.entityType,_that.status,_that.rejectedReason,_that.author,_that.likeCount,_that.commentCount);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _Post implements Post {
  const _Post({required this.id, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt, @JsonKey(name: 'authorId') required this.authorId, required this.content, @JsonKey(name: 'mediaUrls') final  List<String> mediaUrls = const [], @JsonKey(name: 'entityId') this.entityId, @JsonKey(name: 'entityType') this.entityType, @PostStatusConverter() required this.status, @JsonKey(name: 'rejectedReason') this.rejectedReason, @JsonKey(name: 'author') this.author, this.likeCount = 0, this.commentCount = 0}): _mediaUrls = mediaUrls;
  factory _Post.fromJson(Map<String, dynamic> json) => _$PostFromJson(json);

@override final  String id;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;
@override@JsonKey(name: 'authorId') final  String authorId;
@override final  String content;
 final  List<String> _mediaUrls;
@override@JsonKey(name: 'mediaUrls') List<String> get mediaUrls {
  if (_mediaUrls is EqualUnmodifiableListView) return _mediaUrls;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_mediaUrls);
}

@override@JsonKey(name: 'entityId') final  String? entityId;
@override@JsonKey(name: 'entityType') final  PostEntityType? entityType;
@override@PostStatusConverter() final  PostStatus status;
@override@JsonKey(name: 'rejectedReason') final  String? rejectedReason;
@override@JsonKey(name: 'author') final  UserSummary? author;
@override@JsonKey() final  int likeCount;
@override@JsonKey() final  int commentCount;

/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$PostCopyWith<_Post> get copyWith => __$PostCopyWithImpl<_Post>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$PostToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _Post&&(identical(other.id, id) || other.id == id)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt)&&(identical(other.authorId, authorId) || other.authorId == authorId)&&(identical(other.content, content) || other.content == content)&&const DeepCollectionEquality().equals(other._mediaUrls, _mediaUrls)&&(identical(other.entityId, entityId) || other.entityId == entityId)&&(identical(other.entityType, entityType) || other.entityType == entityType)&&(identical(other.status, status) || other.status == status)&&(identical(other.rejectedReason, rejectedReason) || other.rejectedReason == rejectedReason)&&(identical(other.author, author) || other.author == author)&&(identical(other.likeCount, likeCount) || other.likeCount == likeCount)&&(identical(other.commentCount, commentCount) || other.commentCount == commentCount));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,createdAt,updatedAt,authorId,content,const DeepCollectionEquality().hash(_mediaUrls),entityId,entityType,status,rejectedReason,author,likeCount,commentCount);

@override
String toString() {
  return 'Post(id: $id, createdAt: $createdAt, updatedAt: $updatedAt, authorId: $authorId, content: $content, mediaUrls: $mediaUrls, entityId: $entityId, entityType: $entityType, status: $status, rejectedReason: $rejectedReason, author: $author, likeCount: $likeCount, commentCount: $commentCount)';
}


}

/// @nodoc
abstract mixin class _$PostCopyWith<$Res> implements $PostCopyWith<$Res> {
  factory _$PostCopyWith(_Post value, $Res Function(_Post) _then) = __$PostCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt,@JsonKey(name: 'authorId') String authorId, String content,@JsonKey(name: 'mediaUrls') List<String> mediaUrls,@JsonKey(name: 'entityId') String? entityId,@JsonKey(name: 'entityType') PostEntityType? entityType,@PostStatusConverter() PostStatus status,@JsonKey(name: 'rejectedReason') String? rejectedReason,@JsonKey(name: 'author') UserSummary? author, int likeCount, int commentCount
});


@override $UserSummaryCopyWith<$Res>? get author;

}
/// @nodoc
class __$PostCopyWithImpl<$Res>
    implements _$PostCopyWith<$Res> {
  __$PostCopyWithImpl(this._self, this._then);

  final _Post _self;
  final $Res Function(_Post) _then;

/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? createdAt = null,Object? updatedAt = null,Object? authorId = null,Object? content = null,Object? mediaUrls = null,Object? entityId = freezed,Object? entityType = freezed,Object? status = null,Object? rejectedReason = freezed,Object? author = freezed,Object? likeCount = null,Object? commentCount = null,}) {
  return _then(_Post(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,authorId: null == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String,content: null == content ? _self.content : content // ignore: cast_nullable_to_non_nullable
as String,mediaUrls: null == mediaUrls ? _self._mediaUrls : mediaUrls // ignore: cast_nullable_to_non_nullable
as List<String>,entityId: freezed == entityId ? _self.entityId : entityId // ignore: cast_nullable_to_non_nullable
as String?,entityType: freezed == entityType ? _self.entityType : entityType // ignore: cast_nullable_to_non_nullable
as PostEntityType?,status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PostStatus,rejectedReason: freezed == rejectedReason ? _self.rejectedReason : rejectedReason // ignore: cast_nullable_to_non_nullable
as String?,author: freezed == author ? _self.author : author // ignore: cast_nullable_to_non_nullable
as UserSummary?,likeCount: null == likeCount ? _self.likeCount : likeCount // ignore: cast_nullable_to_non_nullable
as int,commentCount: null == commentCount ? _self.commentCount : commentCount // ignore: cast_nullable_to_non_nullable
as int,
  ));
}

/// Create a copy of Post
/// with the given fields replaced by the non-null parameter values.
@override
@pragma('vm:prefer-inline')
$UserSummaryCopyWith<$Res>? get author {
    if (_self.author == null) {
    return null;
  }

  return $UserSummaryCopyWith<$Res>(_self.author!, (value) {
    return _then(_self.copyWith(author: value));
  });
}
}


/// @nodoc
mixin _$UserSummary {

 String get id;@JsonKey(name: 'firstName') String get firstName;@JsonKey(name: 'lastName') String get lastName; String? get email;@JsonKey(name: 'profilePictureUrl') String? get profilePictureUrl;
/// Create a copy of UserSummary
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UserSummaryCopyWith<UserSummary> get copyWith => _$UserSummaryCopyWithImpl<UserSummary>(this as UserSummary, _$identity);

  /// Serializes this UserSummary to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is UserSummary&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,profilePictureUrl);

@override
String toString() {
  return 'UserSummary(id: $id, firstName: $firstName, lastName: $lastName, email: $email, profilePictureUrl: $profilePictureUrl)';
}


}

/// @nodoc
abstract mixin class $UserSummaryCopyWith<$Res>  {
  factory $UserSummaryCopyWith(UserSummary value, $Res Function(UserSummary) _then) = _$UserSummaryCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'firstName') String firstName,@JsonKey(name: 'lastName') String lastName, String? email,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl
});




}
/// @nodoc
class _$UserSummaryCopyWithImpl<$Res>
    implements $UserSummaryCopyWith<$Res> {
  _$UserSummaryCopyWithImpl(this._self, this._then);

  final UserSummary _self;
  final $Res Function(UserSummary) _then;

/// Create a copy of UserSummary
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = freezed,Object? profilePictureUrl = freezed,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: freezed == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}

}


/// Adds pattern-matching-related methods to [UserSummary].
extension UserSummaryPatterns on UserSummary {
/// A variant of `map` that fallback to returning `orElse`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _UserSummary value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _UserSummary() when $default != null:
return $default(_that);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// Callbacks receives the raw object, upcasted.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case final Subclass2 value:
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _UserSummary value)  $default,){
final _that = this;
switch (_that) {
case _UserSummary():
return $default(_that);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `map` that fallback to returning `null`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _UserSummary value)?  $default,){
final _that = this;
switch (_that) {
case _UserSummary() when $default != null:
return $default(_that);case _:
  return null;

}
}
/// A variant of `when` that fallback to an `orElse` callback.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _UserSummary() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// As opposed to `map`, this offers destructuring.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case Subclass2(:final field2):
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl)  $default,) {final _that = this;
switch (_that) {
case _UserSummary():
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `when` that fallback to returning `null`
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl)?  $default,) {final _that = this;
switch (_that) {
case _UserSummary() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _UserSummary implements UserSummary {
  const _UserSummary({required this.id, @JsonKey(name: 'firstName') required this.firstName, @JsonKey(name: 'lastName') required this.lastName, this.email, @JsonKey(name: 'profilePictureUrl') this.profilePictureUrl});
  factory _UserSummary.fromJson(Map<String, dynamic> json) => _$UserSummaryFromJson(json);

@override final  String id;
@override@JsonKey(name: 'firstName') final  String firstName;
@override@JsonKey(name: 'lastName') final  String lastName;
@override final  String? email;
@override@JsonKey(name: 'profilePictureUrl') final  String? profilePictureUrl;

/// Create a copy of UserSummary
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$UserSummaryCopyWith<_UserSummary> get copyWith => __$UserSummaryCopyWithImpl<_UserSummary>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$UserSummaryToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _UserSummary&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,profilePictureUrl);

@override
String toString() {
  return 'UserSummary(id: $id, firstName: $firstName, lastName: $lastName, email: $email, profilePictureUrl: $profilePictureUrl)';
}


}

/// @nodoc
abstract mixin class _$UserSummaryCopyWith<$Res> implements $UserSummaryCopyWith<$Res> {
  factory _$UserSummaryCopyWith(_UserSummary value, $Res Function(_UserSummary) _then) = __$UserSummaryCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'firstName') String firstName,@JsonKey(name: 'lastName') String lastName, String? email,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl
});




}
/// @nodoc
class __$UserSummaryCopyWithImpl<$Res>
    implements _$UserSummaryCopyWith<$Res> {
  __$UserSummaryCopyWithImpl(this._self, this._then);

  final _UserSummary _self;
  final $Res Function(_UserSummary) _then;

/// Create a copy of UserSummary
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = freezed,Object? profilePictureUrl = freezed,}) {
  return _then(_UserSummary(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: freezed == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}


}


/// @nodoc
mixin _$CreatePostRequest {

 String get content; List<String> get mediaUrls;@JsonKey(includeIfNull: true, name: 'entityId') String? get entityId;@JsonKey(includeIfNull: true, name: 'entityType') PostEntityType? get entityType;
/// Create a copy of CreatePostRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreatePostRequestCopyWith<CreatePostRequest> get copyWith => _$CreatePostRequestCopyWithImpl<CreatePostRequest>(this as CreatePostRequest, _$identity);

  /// Serializes this CreatePostRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreatePostRequest&&(identical(other.content, content) || other.content == content)&&const DeepCollectionEquality().equals(other.mediaUrls, mediaUrls)&&(identical(other.entityId, entityId) || other.entityId == entityId)&&(identical(other.entityType, entityType) || other.entityType == entityType));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,content,const DeepCollectionEquality().hash(mediaUrls),entityId,entityType);

@override
String toString() {
  return 'CreatePostRequest(content: $content, mediaUrls: $mediaUrls, entityId: $entityId, entityType: $entityType)';
}


}

/// @nodoc
abstract mixin class $CreatePostRequestCopyWith<$Res>  {
  factory $CreatePostRequestCopyWith(CreatePostRequest value, $Res Function(CreatePostRequest) _then) = _$CreatePostRequestCopyWithImpl;
@useResult
$Res call({
 String content, List<String> mediaUrls,@JsonKey(includeIfNull: true, name: 'entityId') String? entityId,@JsonKey(includeIfNull: true, name: 'entityType') PostEntityType? entityType
});




}
/// @nodoc
class _$CreatePostRequestCopyWithImpl<$Res>
    implements $CreatePostRequestCopyWith<$Res> {
  _$CreatePostRequestCopyWithImpl(this._self, this._then);

  final CreatePostRequest _self;
  final $Res Function(CreatePostRequest) _then;

/// Create a copy of CreatePostRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? content = null,Object? mediaUrls = null,Object? entityId = freezed,Object? entityType = freezed,}) {
  return _then(_self.copyWith(
content: null == content ? _self.content : content // ignore: cast_nullable_to_non_nullable
as String,mediaUrls: null == mediaUrls ? _self.mediaUrls : mediaUrls // ignore: cast_nullable_to_non_nullable
as List<String>,entityId: freezed == entityId ? _self.entityId : entityId // ignore: cast_nullable_to_non_nullable
as String?,entityType: freezed == entityType ? _self.entityType : entityType // ignore: cast_nullable_to_non_nullable
as PostEntityType?,
  ));
}

}


/// Adds pattern-matching-related methods to [CreatePostRequest].
extension CreatePostRequestPatterns on CreatePostRequest {
/// A variant of `map` that fallback to returning `orElse`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreatePostRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreatePostRequest() when $default != null:
return $default(_that);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// Callbacks receives the raw object, upcasted.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case final Subclass2 value:
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreatePostRequest value)  $default,){
final _that = this;
switch (_that) {
case _CreatePostRequest():
return $default(_that);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `map` that fallback to returning `null`.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case final Subclass value:
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreatePostRequest value)?  $default,){
final _that = this;
switch (_that) {
case _CreatePostRequest() when $default != null:
return $default(_that);case _:
  return null;

}
}
/// A variant of `when` that fallback to an `orElse` callback.
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return orElse();
/// }
/// ```

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String content,  List<String> mediaUrls, @JsonKey(includeIfNull: true, name: 'entityId')  String? entityId, @JsonKey(includeIfNull: true, name: 'entityType')  PostEntityType? entityType)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreatePostRequest() when $default != null:
return $default(_that.content,_that.mediaUrls,_that.entityId,_that.entityType);case _:
  return orElse();

}
}
/// A `switch`-like method, using callbacks.
///
/// As opposed to `map`, this offers destructuring.
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case Subclass2(:final field2):
///     return ...;
/// }
/// ```

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String content,  List<String> mediaUrls, @JsonKey(includeIfNull: true, name: 'entityId')  String? entityId, @JsonKey(includeIfNull: true, name: 'entityType')  PostEntityType? entityType)  $default,) {final _that = this;
switch (_that) {
case _CreatePostRequest():
return $default(_that.content,_that.mediaUrls,_that.entityId,_that.entityType);case _:
  throw StateError('Unexpected subclass');

}
}
/// A variant of `when` that fallback to returning `null`
///
/// It is equivalent to doing:
/// ```dart
/// switch (sealedClass) {
///   case Subclass(:final field):
///     return ...;
///   case _:
///     return null;
/// }
/// ```

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String content,  List<String> mediaUrls, @JsonKey(includeIfNull: true, name: 'entityId')  String? entityId, @JsonKey(includeIfNull: true, name: 'entityType')  PostEntityType? entityType)?  $default,) {final _that = this;
switch (_that) {
case _CreatePostRequest() when $default != null:
return $default(_that.content,_that.mediaUrls,_that.entityId,_that.entityType);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreatePostRequest implements CreatePostRequest {
  const _CreatePostRequest({required this.content, final  List<String> mediaUrls = const [], @JsonKey(includeIfNull: true, name: 'entityId') this.entityId, @JsonKey(includeIfNull: true, name: 'entityType') this.entityType}): _mediaUrls = mediaUrls;
  factory _CreatePostRequest.fromJson(Map<String, dynamic> json) => _$CreatePostRequestFromJson(json);

@override final  String content;
 final  List<String> _mediaUrls;
@override@JsonKey() List<String> get mediaUrls {
  if (_mediaUrls is EqualUnmodifiableListView) return _mediaUrls;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_mediaUrls);
}

@override@JsonKey(includeIfNull: true, name: 'entityId') final  String? entityId;
@override@JsonKey(includeIfNull: true, name: 'entityType') final  PostEntityType? entityType;

/// Create a copy of CreatePostRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreatePostRequestCopyWith<_CreatePostRequest> get copyWith => __$CreatePostRequestCopyWithImpl<_CreatePostRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreatePostRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreatePostRequest&&(identical(other.content, content) || other.content == content)&&const DeepCollectionEquality().equals(other._mediaUrls, _mediaUrls)&&(identical(other.entityId, entityId) || other.entityId == entityId)&&(identical(other.entityType, entityType) || other.entityType == entityType));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,content,const DeepCollectionEquality().hash(_mediaUrls),entityId,entityType);

@override
String toString() {
  return 'CreatePostRequest(content: $content, mediaUrls: $mediaUrls, entityId: $entityId, entityType: $entityType)';
}


}

/// @nodoc
abstract mixin class _$CreatePostRequestCopyWith<$Res> implements $CreatePostRequestCopyWith<$Res> {
  factory _$CreatePostRequestCopyWith(_CreatePostRequest value, $Res Function(_CreatePostRequest) _then) = __$CreatePostRequestCopyWithImpl;
@override @useResult
$Res call({
 String content, List<String> mediaUrls,@JsonKey(includeIfNull: true, name: 'entityId') String? entityId,@JsonKey(includeIfNull: true, name: 'entityType') PostEntityType? entityType
});




}
/// @nodoc
class __$CreatePostRequestCopyWithImpl<$Res>
    implements _$CreatePostRequestCopyWith<$Res> {
  __$CreatePostRequestCopyWithImpl(this._self, this._then);

  final _CreatePostRequest _self;
  final $Res Function(_CreatePostRequest) _then;

/// Create a copy of CreatePostRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? content = null,Object? mediaUrls = null,Object? entityId = freezed,Object? entityType = freezed,}) {
  return _then(_CreatePostRequest(
content: null == content ? _self.content : content // ignore: cast_nullable_to_non_nullable
as String,mediaUrls: null == mediaUrls ? _self._mediaUrls : mediaUrls // ignore: cast_nullable_to_non_nullable
as List<String>,entityId: freezed == entityId ? _self.entityId : entityId // ignore: cast_nullable_to_non_nullable
as String?,entityType: freezed == entityType ? _self.entityType : entityType // ignore: cast_nullable_to_non_nullable
as PostEntityType?,
  ));
}


}

// dart format on

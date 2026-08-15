// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'follower.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$FollowerResponse {

 String get id;@JsonKey(name: 'firstName') String get firstName;@JsonKey(name: 'lastName') String get lastName; String? get email;@JsonKey(name: 'profilePictureUrl') String? get profilePictureUrl;@UserTypeConverter() UserType get type;@JsonKey(name: 'createdAt') DateTime get createdAt;
/// Create a copy of FollowerResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$FollowerResponseCopyWith<FollowerResponse> get copyWith => _$FollowerResponseCopyWithImpl<FollowerResponse>(this as FollowerResponse, _$identity);

  /// Serializes this FollowerResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is FollowerResponse&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl)&&(identical(other.type, type) || other.type == type)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,profilePictureUrl,type,createdAt);

@override
String toString() {
  return 'FollowerResponse(id: $id, firstName: $firstName, lastName: $lastName, email: $email, profilePictureUrl: $profilePictureUrl, type: $type, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class $FollowerResponseCopyWith<$Res>  {
  factory $FollowerResponseCopyWith(FollowerResponse value, $Res Function(FollowerResponse) _then) = _$FollowerResponseCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'firstName') String firstName,@JsonKey(name: 'lastName') String lastName, String? email,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,@UserTypeConverter() UserType type,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class _$FollowerResponseCopyWithImpl<$Res>
    implements $FollowerResponseCopyWith<$Res> {
  _$FollowerResponseCopyWithImpl(this._self, this._then);

  final FollowerResponse _self;
  final $Res Function(FollowerResponse) _then;

/// Create a copy of FollowerResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = freezed,Object? profilePictureUrl = freezed,Object? type = null,Object? createdAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: freezed == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as UserType,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [FollowerResponse].
extension FollowerResponsePatterns on FollowerResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _FollowerResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _FollowerResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _FollowerResponse value)  $default,){
final _that = this;
switch (_that) {
case _FollowerResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _FollowerResponse value)?  $default,){
final _that = this;
switch (_that) {
case _FollowerResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @UserTypeConverter()  UserType type, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _FollowerResponse() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl,_that.type,_that.createdAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @UserTypeConverter()  UserType type, @JsonKey(name: 'createdAt')  DateTime createdAt)  $default,) {final _that = this;
switch (_that) {
case _FollowerResponse():
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl,_that.type,_that.createdAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'firstName')  String firstName, @JsonKey(name: 'lastName')  String lastName,  String? email, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @UserTypeConverter()  UserType type, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,) {final _that = this;
switch (_that) {
case _FollowerResponse() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.profilePictureUrl,_that.type,_that.createdAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _FollowerResponse implements FollowerResponse {
  const _FollowerResponse({required this.id, @JsonKey(name: 'firstName') required this.firstName, @JsonKey(name: 'lastName') required this.lastName, this.email, @JsonKey(name: 'profilePictureUrl') this.profilePictureUrl, @UserTypeConverter() required this.type, @JsonKey(name: 'createdAt') required this.createdAt});
  factory _FollowerResponse.fromJson(Map<String, dynamic> json) => _$FollowerResponseFromJson(json);

@override final  String id;
@override@JsonKey(name: 'firstName') final  String firstName;
@override@JsonKey(name: 'lastName') final  String lastName;
@override final  String? email;
@override@JsonKey(name: 'profilePictureUrl') final  String? profilePictureUrl;
@override@UserTypeConverter() final  UserType type;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;

/// Create a copy of FollowerResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$FollowerResponseCopyWith<_FollowerResponse> get copyWith => __$FollowerResponseCopyWithImpl<_FollowerResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$FollowerResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _FollowerResponse&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl)&&(identical(other.type, type) || other.type == type)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,profilePictureUrl,type,createdAt);

@override
String toString() {
  return 'FollowerResponse(id: $id, firstName: $firstName, lastName: $lastName, email: $email, profilePictureUrl: $profilePictureUrl, type: $type, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class _$FollowerResponseCopyWith<$Res> implements $FollowerResponseCopyWith<$Res> {
  factory _$FollowerResponseCopyWith(_FollowerResponse value, $Res Function(_FollowerResponse) _then) = __$FollowerResponseCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'firstName') String firstName,@JsonKey(name: 'lastName') String lastName, String? email,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,@UserTypeConverter() UserType type,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class __$FollowerResponseCopyWithImpl<$Res>
    implements _$FollowerResponseCopyWith<$Res> {
  __$FollowerResponseCopyWithImpl(this._self, this._then);

  final _FollowerResponse _self;
  final $Res Function(_FollowerResponse) _then;

/// Create a copy of FollowerResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = freezed,Object? profilePictureUrl = freezed,Object? type = null,Object? createdAt = null,}) {
  return _then(_FollowerResponse(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: freezed == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as UserType,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$CountResponse {

 int get count;
/// Create a copy of CountResponse
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CountResponseCopyWith<CountResponse> get copyWith => _$CountResponseCopyWithImpl<CountResponse>(this as CountResponse, _$identity);

  /// Serializes this CountResponse to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CountResponse&&(identical(other.count, count) || other.count == count));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,count);

@override
String toString() {
  return 'CountResponse(count: $count)';
}


}

/// @nodoc
abstract mixin class $CountResponseCopyWith<$Res>  {
  factory $CountResponseCopyWith(CountResponse value, $Res Function(CountResponse) _then) = _$CountResponseCopyWithImpl;
@useResult
$Res call({
 int count
});




}
/// @nodoc
class _$CountResponseCopyWithImpl<$Res>
    implements $CountResponseCopyWith<$Res> {
  _$CountResponseCopyWithImpl(this._self, this._then);

  final CountResponse _self;
  final $Res Function(CountResponse) _then;

/// Create a copy of CountResponse
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? count = null,}) {
  return _then(_self.copyWith(
count: null == count ? _self.count : count // ignore: cast_nullable_to_non_nullable
as int,
  ));
}

}


/// Adds pattern-matching-related methods to [CountResponse].
extension CountResponsePatterns on CountResponse {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CountResponse value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CountResponse() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CountResponse value)  $default,){
final _that = this;
switch (_that) {
case _CountResponse():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CountResponse value)?  $default,){
final _that = this;
switch (_that) {
case _CountResponse() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( int count)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CountResponse() when $default != null:
return $default(_that.count);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( int count)  $default,) {final _that = this;
switch (_that) {
case _CountResponse():
return $default(_that.count);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( int count)?  $default,) {final _that = this;
switch (_that) {
case _CountResponse() when $default != null:
return $default(_that.count);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CountResponse implements CountResponse {
  const _CountResponse({required this.count});
  factory _CountResponse.fromJson(Map<String, dynamic> json) => _$CountResponseFromJson(json);

@override final  int count;

/// Create a copy of CountResponse
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CountResponseCopyWith<_CountResponse> get copyWith => __$CountResponseCopyWithImpl<_CountResponse>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CountResponseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CountResponse&&(identical(other.count, count) || other.count == count));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,count);

@override
String toString() {
  return 'CountResponse(count: $count)';
}


}

/// @nodoc
abstract mixin class _$CountResponseCopyWith<$Res> implements $CountResponseCopyWith<$Res> {
  factory _$CountResponseCopyWith(_CountResponse value, $Res Function(_CountResponse) _then) = __$CountResponseCopyWithImpl;
@override @useResult
$Res call({
 int count
});




}
/// @nodoc
class __$CountResponseCopyWithImpl<$Res>
    implements _$CountResponseCopyWith<$Res> {
  __$CountResponseCopyWithImpl(this._self, this._then);

  final _CountResponse _self;
  final $Res Function(_CountResponse) _then;

/// Create a copy of CountResponse
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? count = null,}) {
  return _then(_CountResponse(
count: null == count ? _self.count : count // ignore: cast_nullable_to_non_nullable
as int,
  ));
}


}

// dart format on

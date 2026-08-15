// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$User {

 String get id; String get firstName; String get lastName; String get email;@JsonKey(name: 'emailVerifiedAt') DateTime? get emailVerifiedAt; String? get bio;@JsonKey(name: 'profilePictureUrl') String? get profilePictureUrl;@JsonKey(name: 'type') UserType get type; double? get height;@JsonKey(name: 'heightUnit') HeightUnit? get heightUnit;@JsonKey(name: 'currentWeight') double? get currentWeight;@JsonKey(name: 'weightUnit') WeightUnit? get weightUnit;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;
/// Create a copy of User
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UserCopyWith<User> get copyWith => _$UserCopyWithImpl<User>(this as User, _$identity);

  /// Serializes this User to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is User&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.emailVerifiedAt, emailVerifiedAt) || other.emailVerifiedAt == emailVerifiedAt)&&(identical(other.bio, bio) || other.bio == bio)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl)&&(identical(other.type, type) || other.type == type)&&(identical(other.height, height) || other.height == height)&&(identical(other.heightUnit, heightUnit) || other.heightUnit == heightUnit)&&(identical(other.currentWeight, currentWeight) || other.currentWeight == currentWeight)&&(identical(other.weightUnit, weightUnit) || other.weightUnit == weightUnit)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,emailVerifiedAt,bio,profilePictureUrl,type,height,heightUnit,currentWeight,weightUnit,createdAt,updatedAt);

@override
String toString() {
  return 'User(id: $id, firstName: $firstName, lastName: $lastName, email: $email, emailVerifiedAt: $emailVerifiedAt, bio: $bio, profilePictureUrl: $profilePictureUrl, type: $type, height: $height, heightUnit: $heightUnit, currentWeight: $currentWeight, weightUnit: $weightUnit, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $UserCopyWith<$Res>  {
  factory $UserCopyWith(User value, $Res Function(User) _then) = _$UserCopyWithImpl;
@useResult
$Res call({
 String id, String firstName, String lastName, String email,@JsonKey(name: 'emailVerifiedAt') DateTime? emailVerifiedAt, String? bio,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,@JsonKey(name: 'type') UserType type, double? height,@JsonKey(name: 'heightUnit') HeightUnit? heightUnit,@JsonKey(name: 'currentWeight') double? currentWeight,@JsonKey(name: 'weightUnit') WeightUnit? weightUnit,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class _$UserCopyWithImpl<$Res>
    implements $UserCopyWith<$Res> {
  _$UserCopyWithImpl(this._self, this._then);

  final User _self;
  final $Res Function(User) _then;

/// Create a copy of User
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = null,Object? emailVerifiedAt = freezed,Object? bio = freezed,Object? profilePictureUrl = freezed,Object? type = null,Object? height = freezed,Object? heightUnit = freezed,Object? currentWeight = freezed,Object? weightUnit = freezed,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: null == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String,emailVerifiedAt: freezed == emailVerifiedAt ? _self.emailVerifiedAt : emailVerifiedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,bio: freezed == bio ? _self.bio : bio // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as UserType,height: freezed == height ? _self.height : height // ignore: cast_nullable_to_non_nullable
as double?,heightUnit: freezed == heightUnit ? _self.heightUnit : heightUnit // ignore: cast_nullable_to_non_nullable
as HeightUnit?,currentWeight: freezed == currentWeight ? _self.currentWeight : currentWeight // ignore: cast_nullable_to_non_nullable
as double?,weightUnit: freezed == weightUnit ? _self.weightUnit : weightUnit // ignore: cast_nullable_to_non_nullable
as WeightUnit?,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [User].
extension UserPatterns on User {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _User value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _User() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _User value)  $default,){
final _that = this;
switch (_that) {
case _User():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _User value)?  $default,){
final _that = this;
switch (_that) {
case _User() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id,  String firstName,  String lastName,  String email, @JsonKey(name: 'emailVerifiedAt')  DateTime? emailVerifiedAt,  String? bio, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @JsonKey(name: 'type')  UserType type,  double? height, @JsonKey(name: 'heightUnit')  HeightUnit? heightUnit, @JsonKey(name: 'currentWeight')  double? currentWeight, @JsonKey(name: 'weightUnit')  WeightUnit? weightUnit, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _User() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.emailVerifiedAt,_that.bio,_that.profilePictureUrl,_that.type,_that.height,_that.heightUnit,_that.currentWeight,_that.weightUnit,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id,  String firstName,  String lastName,  String email, @JsonKey(name: 'emailVerifiedAt')  DateTime? emailVerifiedAt,  String? bio, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @JsonKey(name: 'type')  UserType type,  double? height, @JsonKey(name: 'heightUnit')  HeightUnit? heightUnit, @JsonKey(name: 'currentWeight')  double? currentWeight, @JsonKey(name: 'weightUnit')  WeightUnit? weightUnit, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _User():
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.emailVerifiedAt,_that.bio,_that.profilePictureUrl,_that.type,_that.height,_that.heightUnit,_that.currentWeight,_that.weightUnit,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id,  String firstName,  String lastName,  String email, @JsonKey(name: 'emailVerifiedAt')  DateTime? emailVerifiedAt,  String? bio, @JsonKey(name: 'profilePictureUrl')  String? profilePictureUrl, @JsonKey(name: 'type')  UserType type,  double? height, @JsonKey(name: 'heightUnit')  HeightUnit? heightUnit, @JsonKey(name: 'currentWeight')  double? currentWeight, @JsonKey(name: 'weightUnit')  WeightUnit? weightUnit, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _User() when $default != null:
return $default(_that.id,_that.firstName,_that.lastName,_that.email,_that.emailVerifiedAt,_that.bio,_that.profilePictureUrl,_that.type,_that.height,_that.heightUnit,_that.currentWeight,_that.weightUnit,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _User implements User {
  const _User({required this.id, required this.firstName, required this.lastName, required this.email, @JsonKey(name: 'emailVerifiedAt') this.emailVerifiedAt, this.bio, @JsonKey(name: 'profilePictureUrl') this.profilePictureUrl, @JsonKey(name: 'type') required this.type, this.height, @JsonKey(name: 'heightUnit') this.heightUnit, @JsonKey(name: 'currentWeight') this.currentWeight, @JsonKey(name: 'weightUnit') this.weightUnit, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt});
  factory _User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);

@override final  String id;
@override final  String firstName;
@override final  String lastName;
@override final  String email;
@override@JsonKey(name: 'emailVerifiedAt') final  DateTime? emailVerifiedAt;
@override final  String? bio;
@override@JsonKey(name: 'profilePictureUrl') final  String? profilePictureUrl;
@override@JsonKey(name: 'type') final  UserType type;
@override final  double? height;
@override@JsonKey(name: 'heightUnit') final  HeightUnit? heightUnit;
@override@JsonKey(name: 'currentWeight') final  double? currentWeight;
@override@JsonKey(name: 'weightUnit') final  WeightUnit? weightUnit;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;

/// Create a copy of User
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$UserCopyWith<_User> get copyWith => __$UserCopyWithImpl<_User>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$UserToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _User&&(identical(other.id, id) || other.id == id)&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.email, email) || other.email == email)&&(identical(other.emailVerifiedAt, emailVerifiedAt) || other.emailVerifiedAt == emailVerifiedAt)&&(identical(other.bio, bio) || other.bio == bio)&&(identical(other.profilePictureUrl, profilePictureUrl) || other.profilePictureUrl == profilePictureUrl)&&(identical(other.type, type) || other.type == type)&&(identical(other.height, height) || other.height == height)&&(identical(other.heightUnit, heightUnit) || other.heightUnit == heightUnit)&&(identical(other.currentWeight, currentWeight) || other.currentWeight == currentWeight)&&(identical(other.weightUnit, weightUnit) || other.weightUnit == weightUnit)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,firstName,lastName,email,emailVerifiedAt,bio,profilePictureUrl,type,height,heightUnit,currentWeight,weightUnit,createdAt,updatedAt);

@override
String toString() {
  return 'User(id: $id, firstName: $firstName, lastName: $lastName, email: $email, emailVerifiedAt: $emailVerifiedAt, bio: $bio, profilePictureUrl: $profilePictureUrl, type: $type, height: $height, heightUnit: $heightUnit, currentWeight: $currentWeight, weightUnit: $weightUnit, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$UserCopyWith<$Res> implements $UserCopyWith<$Res> {
  factory _$UserCopyWith(_User value, $Res Function(_User) _then) = __$UserCopyWithImpl;
@override @useResult
$Res call({
 String id, String firstName, String lastName, String email,@JsonKey(name: 'emailVerifiedAt') DateTime? emailVerifiedAt, String? bio,@JsonKey(name: 'profilePictureUrl') String? profilePictureUrl,@JsonKey(name: 'type') UserType type, double? height,@JsonKey(name: 'heightUnit') HeightUnit? heightUnit,@JsonKey(name: 'currentWeight') double? currentWeight,@JsonKey(name: 'weightUnit') WeightUnit? weightUnit,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class __$UserCopyWithImpl<$Res>
    implements _$UserCopyWith<$Res> {
  __$UserCopyWithImpl(this._self, this._then);

  final _User _self;
  final $Res Function(_User) _then;

/// Create a copy of User
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? firstName = null,Object? lastName = null,Object? email = null,Object? emailVerifiedAt = freezed,Object? bio = freezed,Object? profilePictureUrl = freezed,Object? type = null,Object? height = freezed,Object? heightUnit = freezed,Object? currentWeight = freezed,Object? weightUnit = freezed,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_User(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,firstName: null == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String,lastName: null == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String,email: null == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String,emailVerifiedAt: freezed == emailVerifiedAt ? _self.emailVerifiedAt : emailVerifiedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,bio: freezed == bio ? _self.bio : bio // ignore: cast_nullable_to_non_nullable
as String?,profilePictureUrl: freezed == profilePictureUrl ? _self.profilePictureUrl : profilePictureUrl // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as UserType,height: freezed == height ? _self.height : height // ignore: cast_nullable_to_non_nullable
as double?,heightUnit: freezed == heightUnit ? _self.heightUnit : heightUnit // ignore: cast_nullable_to_non_nullable
as HeightUnit?,currentWeight: freezed == currentWeight ? _self.currentWeight : currentWeight // ignore: cast_nullable_to_non_nullable
as double?,weightUnit: freezed == weightUnit ? _self.weightUnit : weightUnit // ignore: cast_nullable_to_non_nullable
as WeightUnit?,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$UpdateProfileRequest {

@JsonKey(includeIfNull: false) String? get firstName;@JsonKey(includeIfNull: false) String? get lastName;@JsonKey(includeIfNull: false) String? get bio;@JsonKey(includeIfNull: false) double? get height;@JsonKey(includeIfNull: false)@HeightUnitConverter() HeightUnit? get heightUnit;@JsonKey(includeIfNull: false)@WeightUnitConverter() WeightUnit? get weightUnit;@JsonKey(includeIfNull: false, name: 'currentWeight') double? get currentWeight;
/// Create a copy of UpdateProfileRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UpdateProfileRequestCopyWith<UpdateProfileRequest> get copyWith => _$UpdateProfileRequestCopyWithImpl<UpdateProfileRequest>(this as UpdateProfileRequest, _$identity);

  /// Serializes this UpdateProfileRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is UpdateProfileRequest&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.bio, bio) || other.bio == bio)&&(identical(other.height, height) || other.height == height)&&(identical(other.heightUnit, heightUnit) || other.heightUnit == heightUnit)&&(identical(other.weightUnit, weightUnit) || other.weightUnit == weightUnit)&&(identical(other.currentWeight, currentWeight) || other.currentWeight == currentWeight));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,firstName,lastName,bio,height,heightUnit,weightUnit,currentWeight);

@override
String toString() {
  return 'UpdateProfileRequest(firstName: $firstName, lastName: $lastName, bio: $bio, height: $height, heightUnit: $heightUnit, weightUnit: $weightUnit, currentWeight: $currentWeight)';
}


}

/// @nodoc
abstract mixin class $UpdateProfileRequestCopyWith<$Res>  {
  factory $UpdateProfileRequestCopyWith(UpdateProfileRequest value, $Res Function(UpdateProfileRequest) _then) = _$UpdateProfileRequestCopyWithImpl;
@useResult
$Res call({
@JsonKey(includeIfNull: false) String? firstName,@JsonKey(includeIfNull: false) String? lastName,@JsonKey(includeIfNull: false) String? bio,@JsonKey(includeIfNull: false) double? height,@JsonKey(includeIfNull: false)@HeightUnitConverter() HeightUnit? heightUnit,@JsonKey(includeIfNull: false)@WeightUnitConverter() WeightUnit? weightUnit,@JsonKey(includeIfNull: false, name: 'currentWeight') double? currentWeight
});




}
/// @nodoc
class _$UpdateProfileRequestCopyWithImpl<$Res>
    implements $UpdateProfileRequestCopyWith<$Res> {
  _$UpdateProfileRequestCopyWithImpl(this._self, this._then);

  final UpdateProfileRequest _self;
  final $Res Function(UpdateProfileRequest) _then;

/// Create a copy of UpdateProfileRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? firstName = freezed,Object? lastName = freezed,Object? bio = freezed,Object? height = freezed,Object? heightUnit = freezed,Object? weightUnit = freezed,Object? currentWeight = freezed,}) {
  return _then(_self.copyWith(
firstName: freezed == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String?,lastName: freezed == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String?,bio: freezed == bio ? _self.bio : bio // ignore: cast_nullable_to_non_nullable
as String?,height: freezed == height ? _self.height : height // ignore: cast_nullable_to_non_nullable
as double?,heightUnit: freezed == heightUnit ? _self.heightUnit : heightUnit // ignore: cast_nullable_to_non_nullable
as HeightUnit?,weightUnit: freezed == weightUnit ? _self.weightUnit : weightUnit // ignore: cast_nullable_to_non_nullable
as WeightUnit?,currentWeight: freezed == currentWeight ? _self.currentWeight : currentWeight // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}

}


/// Adds pattern-matching-related methods to [UpdateProfileRequest].
extension UpdateProfileRequestPatterns on UpdateProfileRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _UpdateProfileRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _UpdateProfileRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _UpdateProfileRequest value)  $default,){
final _that = this;
switch (_that) {
case _UpdateProfileRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _UpdateProfileRequest value)?  $default,){
final _that = this;
switch (_that) {
case _UpdateProfileRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@JsonKey(includeIfNull: false)  String? firstName, @JsonKey(includeIfNull: false)  String? lastName, @JsonKey(includeIfNull: false)  String? bio, @JsonKey(includeIfNull: false)  double? height, @JsonKey(includeIfNull: false)@HeightUnitConverter()  HeightUnit? heightUnit, @JsonKey(includeIfNull: false)@WeightUnitConverter()  WeightUnit? weightUnit, @JsonKey(includeIfNull: false, name: 'currentWeight')  double? currentWeight)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _UpdateProfileRequest() when $default != null:
return $default(_that.firstName,_that.lastName,_that.bio,_that.height,_that.heightUnit,_that.weightUnit,_that.currentWeight);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@JsonKey(includeIfNull: false)  String? firstName, @JsonKey(includeIfNull: false)  String? lastName, @JsonKey(includeIfNull: false)  String? bio, @JsonKey(includeIfNull: false)  double? height, @JsonKey(includeIfNull: false)@HeightUnitConverter()  HeightUnit? heightUnit, @JsonKey(includeIfNull: false)@WeightUnitConverter()  WeightUnit? weightUnit, @JsonKey(includeIfNull: false, name: 'currentWeight')  double? currentWeight)  $default,) {final _that = this;
switch (_that) {
case _UpdateProfileRequest():
return $default(_that.firstName,_that.lastName,_that.bio,_that.height,_that.heightUnit,_that.weightUnit,_that.currentWeight);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@JsonKey(includeIfNull: false)  String? firstName, @JsonKey(includeIfNull: false)  String? lastName, @JsonKey(includeIfNull: false)  String? bio, @JsonKey(includeIfNull: false)  double? height, @JsonKey(includeIfNull: false)@HeightUnitConverter()  HeightUnit? heightUnit, @JsonKey(includeIfNull: false)@WeightUnitConverter()  WeightUnit? weightUnit, @JsonKey(includeIfNull: false, name: 'currentWeight')  double? currentWeight)?  $default,) {final _that = this;
switch (_that) {
case _UpdateProfileRequest() when $default != null:
return $default(_that.firstName,_that.lastName,_that.bio,_that.height,_that.heightUnit,_that.weightUnit,_that.currentWeight);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _UpdateProfileRequest implements UpdateProfileRequest {
  const _UpdateProfileRequest({@JsonKey(includeIfNull: false) this.firstName, @JsonKey(includeIfNull: false) this.lastName, @JsonKey(includeIfNull: false) this.bio, @JsonKey(includeIfNull: false) this.height, @JsonKey(includeIfNull: false)@HeightUnitConverter() this.heightUnit, @JsonKey(includeIfNull: false)@WeightUnitConverter() this.weightUnit, @JsonKey(includeIfNull: false, name: 'currentWeight') this.currentWeight});
  factory _UpdateProfileRequest.fromJson(Map<String, dynamic> json) => _$UpdateProfileRequestFromJson(json);

@override@JsonKey(includeIfNull: false) final  String? firstName;
@override@JsonKey(includeIfNull: false) final  String? lastName;
@override@JsonKey(includeIfNull: false) final  String? bio;
@override@JsonKey(includeIfNull: false) final  double? height;
@override@JsonKey(includeIfNull: false)@HeightUnitConverter() final  HeightUnit? heightUnit;
@override@JsonKey(includeIfNull: false)@WeightUnitConverter() final  WeightUnit? weightUnit;
@override@JsonKey(includeIfNull: false, name: 'currentWeight') final  double? currentWeight;

/// Create a copy of UpdateProfileRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$UpdateProfileRequestCopyWith<_UpdateProfileRequest> get copyWith => __$UpdateProfileRequestCopyWithImpl<_UpdateProfileRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$UpdateProfileRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _UpdateProfileRequest&&(identical(other.firstName, firstName) || other.firstName == firstName)&&(identical(other.lastName, lastName) || other.lastName == lastName)&&(identical(other.bio, bio) || other.bio == bio)&&(identical(other.height, height) || other.height == height)&&(identical(other.heightUnit, heightUnit) || other.heightUnit == heightUnit)&&(identical(other.weightUnit, weightUnit) || other.weightUnit == weightUnit)&&(identical(other.currentWeight, currentWeight) || other.currentWeight == currentWeight));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,firstName,lastName,bio,height,heightUnit,weightUnit,currentWeight);

@override
String toString() {
  return 'UpdateProfileRequest(firstName: $firstName, lastName: $lastName, bio: $bio, height: $height, heightUnit: $heightUnit, weightUnit: $weightUnit, currentWeight: $currentWeight)';
}


}

/// @nodoc
abstract mixin class _$UpdateProfileRequestCopyWith<$Res> implements $UpdateProfileRequestCopyWith<$Res> {
  factory _$UpdateProfileRequestCopyWith(_UpdateProfileRequest value, $Res Function(_UpdateProfileRequest) _then) = __$UpdateProfileRequestCopyWithImpl;
@override @useResult
$Res call({
@JsonKey(includeIfNull: false) String? firstName,@JsonKey(includeIfNull: false) String? lastName,@JsonKey(includeIfNull: false) String? bio,@JsonKey(includeIfNull: false) double? height,@JsonKey(includeIfNull: false)@HeightUnitConverter() HeightUnit? heightUnit,@JsonKey(includeIfNull: false)@WeightUnitConverter() WeightUnit? weightUnit,@JsonKey(includeIfNull: false, name: 'currentWeight') double? currentWeight
});




}
/// @nodoc
class __$UpdateProfileRequestCopyWithImpl<$Res>
    implements _$UpdateProfileRequestCopyWith<$Res> {
  __$UpdateProfileRequestCopyWithImpl(this._self, this._then);

  final _UpdateProfileRequest _self;
  final $Res Function(_UpdateProfileRequest) _then;

/// Create a copy of UpdateProfileRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? firstName = freezed,Object? lastName = freezed,Object? bio = freezed,Object? height = freezed,Object? heightUnit = freezed,Object? weightUnit = freezed,Object? currentWeight = freezed,}) {
  return _then(_UpdateProfileRequest(
firstName: freezed == firstName ? _self.firstName : firstName // ignore: cast_nullable_to_non_nullable
as String?,lastName: freezed == lastName ? _self.lastName : lastName // ignore: cast_nullable_to_non_nullable
as String?,bio: freezed == bio ? _self.bio : bio // ignore: cast_nullable_to_non_nullable
as String?,height: freezed == height ? _self.height : height // ignore: cast_nullable_to_non_nullable
as double?,heightUnit: freezed == heightUnit ? _self.heightUnit : heightUnit // ignore: cast_nullable_to_non_nullable
as HeightUnit?,weightUnit: freezed == weightUnit ? _self.weightUnit : weightUnit // ignore: cast_nullable_to_non_nullable
as WeightUnit?,currentWeight: freezed == currentWeight ? _self.currentWeight : currentWeight // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}


}

// dart format on

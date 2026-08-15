// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'feedback.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$TrainingPlanFeedback {

 String get id;@JsonKey(name: 'planId') String get planId;@JsonKey(name: 'userId') String get userId; double get rating; String? get message;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;
/// Create a copy of TrainingPlanFeedback
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$TrainingPlanFeedbackCopyWith<TrainingPlanFeedback> get copyWith => _$TrainingPlanFeedbackCopyWithImpl<TrainingPlanFeedback>(this as TrainingPlanFeedback, _$identity);

  /// Serializes this TrainingPlanFeedback to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is TrainingPlanFeedback&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.rating, rating) || other.rating == rating)&&(identical(other.message, message) || other.message == message)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,userId,rating,message,createdAt,updatedAt);

@override
String toString() {
  return 'TrainingPlanFeedback(id: $id, planId: $planId, userId: $userId, rating: $rating, message: $message, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $TrainingPlanFeedbackCopyWith<$Res>  {
  factory $TrainingPlanFeedbackCopyWith(TrainingPlanFeedback value, $Res Function(TrainingPlanFeedback) _then) = _$TrainingPlanFeedbackCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'userId') String userId, double rating, String? message,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class _$TrainingPlanFeedbackCopyWithImpl<$Res>
    implements $TrainingPlanFeedbackCopyWith<$Res> {
  _$TrainingPlanFeedbackCopyWithImpl(this._self, this._then);

  final TrainingPlanFeedback _self;
  final $Res Function(TrainingPlanFeedback) _then;

/// Create a copy of TrainingPlanFeedback
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? planId = null,Object? userId = null,Object? rating = null,Object? message = freezed,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,rating: null == rating ? _self.rating : rating // ignore: cast_nullable_to_non_nullable
as double,message: freezed == message ? _self.message : message // ignore: cast_nullable_to_non_nullable
as String?,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [TrainingPlanFeedback].
extension TrainingPlanFeedbackPatterns on TrainingPlanFeedback {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _TrainingPlanFeedback value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _TrainingPlanFeedback() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _TrainingPlanFeedback value)  $default,){
final _that = this;
switch (_that) {
case _TrainingPlanFeedback():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _TrainingPlanFeedback value)?  $default,){
final _that = this;
switch (_that) {
case _TrainingPlanFeedback() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId,  double rating,  String? message, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _TrainingPlanFeedback() when $default != null:
return $default(_that.id,_that.planId,_that.userId,_that.rating,_that.message,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId,  double rating,  String? message, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _TrainingPlanFeedback():
return $default(_that.id,_that.planId,_that.userId,_that.rating,_that.message,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId,  double rating,  String? message, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _TrainingPlanFeedback() when $default != null:
return $default(_that.id,_that.planId,_that.userId,_that.rating,_that.message,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _TrainingPlanFeedback implements TrainingPlanFeedback {
  const _TrainingPlanFeedback({required this.id, @JsonKey(name: 'planId') required this.planId, @JsonKey(name: 'userId') required this.userId, required this.rating, this.message, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt});
  factory _TrainingPlanFeedback.fromJson(Map<String, dynamic> json) => _$TrainingPlanFeedbackFromJson(json);

@override final  String id;
@override@JsonKey(name: 'planId') final  String planId;
@override@JsonKey(name: 'userId') final  String userId;
@override final  double rating;
@override final  String? message;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;

/// Create a copy of TrainingPlanFeedback
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$TrainingPlanFeedbackCopyWith<_TrainingPlanFeedback> get copyWith => __$TrainingPlanFeedbackCopyWithImpl<_TrainingPlanFeedback>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$TrainingPlanFeedbackToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _TrainingPlanFeedback&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.rating, rating) || other.rating == rating)&&(identical(other.message, message) || other.message == message)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,userId,rating,message,createdAt,updatedAt);

@override
String toString() {
  return 'TrainingPlanFeedback(id: $id, planId: $planId, userId: $userId, rating: $rating, message: $message, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$TrainingPlanFeedbackCopyWith<$Res> implements $TrainingPlanFeedbackCopyWith<$Res> {
  factory _$TrainingPlanFeedbackCopyWith(_TrainingPlanFeedback value, $Res Function(_TrainingPlanFeedback) _then) = __$TrainingPlanFeedbackCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'userId') String userId, double rating, String? message,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class __$TrainingPlanFeedbackCopyWithImpl<$Res>
    implements _$TrainingPlanFeedbackCopyWith<$Res> {
  __$TrainingPlanFeedbackCopyWithImpl(this._self, this._then);

  final _TrainingPlanFeedback _self;
  final $Res Function(_TrainingPlanFeedback) _then;

/// Create a copy of TrainingPlanFeedback
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? planId = null,Object? userId = null,Object? rating = null,Object? message = freezed,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_TrainingPlanFeedback(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,rating: null == rating ? _self.rating : rating // ignore: cast_nullable_to_non_nullable
as double,message: freezed == message ? _self.message : message // ignore: cast_nullable_to_non_nullable
as String?,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$AddFeedbackRequest {

 double get rating;@JsonKey(includeIfNull: true) String? get message;
/// Create a copy of AddFeedbackRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$AddFeedbackRequestCopyWith<AddFeedbackRequest> get copyWith => _$AddFeedbackRequestCopyWithImpl<AddFeedbackRequest>(this as AddFeedbackRequest, _$identity);

  /// Serializes this AddFeedbackRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is AddFeedbackRequest&&(identical(other.rating, rating) || other.rating == rating)&&(identical(other.message, message) || other.message == message));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,rating,message);

@override
String toString() {
  return 'AddFeedbackRequest(rating: $rating, message: $message)';
}


}

/// @nodoc
abstract mixin class $AddFeedbackRequestCopyWith<$Res>  {
  factory $AddFeedbackRequestCopyWith(AddFeedbackRequest value, $Res Function(AddFeedbackRequest) _then) = _$AddFeedbackRequestCopyWithImpl;
@useResult
$Res call({
 double rating,@JsonKey(includeIfNull: true) String? message
});




}
/// @nodoc
class _$AddFeedbackRequestCopyWithImpl<$Res>
    implements $AddFeedbackRequestCopyWith<$Res> {
  _$AddFeedbackRequestCopyWithImpl(this._self, this._then);

  final AddFeedbackRequest _self;
  final $Res Function(AddFeedbackRequest) _then;

/// Create a copy of AddFeedbackRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? rating = null,Object? message = freezed,}) {
  return _then(_self.copyWith(
rating: null == rating ? _self.rating : rating // ignore: cast_nullable_to_non_nullable
as double,message: freezed == message ? _self.message : message // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}

}


/// Adds pattern-matching-related methods to [AddFeedbackRequest].
extension AddFeedbackRequestPatterns on AddFeedbackRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _AddFeedbackRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _AddFeedbackRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _AddFeedbackRequest value)  $default,){
final _that = this;
switch (_that) {
case _AddFeedbackRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _AddFeedbackRequest value)?  $default,){
final _that = this;
switch (_that) {
case _AddFeedbackRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( double rating, @JsonKey(includeIfNull: true)  String? message)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _AddFeedbackRequest() when $default != null:
return $default(_that.rating,_that.message);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( double rating, @JsonKey(includeIfNull: true)  String? message)  $default,) {final _that = this;
switch (_that) {
case _AddFeedbackRequest():
return $default(_that.rating,_that.message);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( double rating, @JsonKey(includeIfNull: true)  String? message)?  $default,) {final _that = this;
switch (_that) {
case _AddFeedbackRequest() when $default != null:
return $default(_that.rating,_that.message);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _AddFeedbackRequest implements AddFeedbackRequest {
  const _AddFeedbackRequest({required this.rating, @JsonKey(includeIfNull: true) this.message});
  factory _AddFeedbackRequest.fromJson(Map<String, dynamic> json) => _$AddFeedbackRequestFromJson(json);

@override final  double rating;
@override@JsonKey(includeIfNull: true) final  String? message;

/// Create a copy of AddFeedbackRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$AddFeedbackRequestCopyWith<_AddFeedbackRequest> get copyWith => __$AddFeedbackRequestCopyWithImpl<_AddFeedbackRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$AddFeedbackRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _AddFeedbackRequest&&(identical(other.rating, rating) || other.rating == rating)&&(identical(other.message, message) || other.message == message));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,rating,message);

@override
String toString() {
  return 'AddFeedbackRequest(rating: $rating, message: $message)';
}


}

/// @nodoc
abstract mixin class _$AddFeedbackRequestCopyWith<$Res> implements $AddFeedbackRequestCopyWith<$Res> {
  factory _$AddFeedbackRequestCopyWith(_AddFeedbackRequest value, $Res Function(_AddFeedbackRequest) _then) = __$AddFeedbackRequestCopyWithImpl;
@override @useResult
$Res call({
 double rating,@JsonKey(includeIfNull: true) String? message
});




}
/// @nodoc
class __$AddFeedbackRequestCopyWithImpl<$Res>
    implements _$AddFeedbackRequestCopyWith<$Res> {
  __$AddFeedbackRequestCopyWithImpl(this._self, this._then);

  final _AddFeedbackRequest _self;
  final $Res Function(_AddFeedbackRequest) _then;

/// Create a copy of AddFeedbackRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? rating = null,Object? message = freezed,}) {
  return _then(_AddFeedbackRequest(
rating: null == rating ? _self.rating : rating // ignore: cast_nullable_to_non_nullable
as double,message: freezed == message ? _self.message : message // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}


}

// dart format on

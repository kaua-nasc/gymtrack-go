// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'subscription.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$PlanSubscription {

 String get id;@JsonKey(name: 'planId') String get planId;@JsonKey(name: 'userId') String get userId;@PlanSubscriptionStatusConverter() PlanSubscriptionStatus get status;@PlanSubscriptionTypeConverter() PlanSubscriptionType get type;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;
/// Create a copy of PlanSubscription
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$PlanSubscriptionCopyWith<PlanSubscription> get copyWith => _$PlanSubscriptionCopyWithImpl<PlanSubscription>(this as PlanSubscription, _$identity);

  /// Serializes this PlanSubscription to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is PlanSubscription&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.status, status) || other.status == status)&&(identical(other.type, type) || other.type == type)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,userId,status,type,createdAt,updatedAt);

@override
String toString() {
  return 'PlanSubscription(id: $id, planId: $planId, userId: $userId, status: $status, type: $type, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $PlanSubscriptionCopyWith<$Res>  {
  factory $PlanSubscriptionCopyWith(PlanSubscription value, $Res Function(PlanSubscription) _then) = _$PlanSubscriptionCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'userId') String userId,@PlanSubscriptionStatusConverter() PlanSubscriptionStatus status,@PlanSubscriptionTypeConverter() PlanSubscriptionType type,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class _$PlanSubscriptionCopyWithImpl<$Res>
    implements $PlanSubscriptionCopyWith<$Res> {
  _$PlanSubscriptionCopyWithImpl(this._self, this._then);

  final PlanSubscription _self;
  final $Res Function(PlanSubscription) _then;

/// Create a copy of PlanSubscription
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? planId = null,Object? userId = null,Object? status = null,Object? type = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [PlanSubscription].
extension PlanSubscriptionPatterns on PlanSubscription {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _PlanSubscription value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _PlanSubscription() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _PlanSubscription value)  $default,){
final _that = this;
switch (_that) {
case _PlanSubscription():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _PlanSubscription value)?  $default,){
final _that = this;
switch (_that) {
case _PlanSubscription() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId, @PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status, @PlanSubscriptionTypeConverter()  PlanSubscriptionType type, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _PlanSubscription() when $default != null:
return $default(_that.id,_that.planId,_that.userId,_that.status,_that.type,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId, @PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status, @PlanSubscriptionTypeConverter()  PlanSubscriptionType type, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _PlanSubscription():
return $default(_that.id,_that.planId,_that.userId,_that.status,_that.type,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'userId')  String userId, @PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status, @PlanSubscriptionTypeConverter()  PlanSubscriptionType type, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _PlanSubscription() when $default != null:
return $default(_that.id,_that.planId,_that.userId,_that.status,_that.type,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _PlanSubscription implements PlanSubscription {
  const _PlanSubscription({required this.id, @JsonKey(name: 'planId') required this.planId, @JsonKey(name: 'userId') required this.userId, @PlanSubscriptionStatusConverter() required this.status, @PlanSubscriptionTypeConverter() required this.type, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt});
  factory _PlanSubscription.fromJson(Map<String, dynamic> json) => _$PlanSubscriptionFromJson(json);

@override final  String id;
@override@JsonKey(name: 'planId') final  String planId;
@override@JsonKey(name: 'userId') final  String userId;
@override@PlanSubscriptionStatusConverter() final  PlanSubscriptionStatus status;
@override@PlanSubscriptionTypeConverter() final  PlanSubscriptionType type;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;

/// Create a copy of PlanSubscription
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$PlanSubscriptionCopyWith<_PlanSubscription> get copyWith => __$PlanSubscriptionCopyWithImpl<_PlanSubscription>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$PlanSubscriptionToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _PlanSubscription&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.status, status) || other.status == status)&&(identical(other.type, type) || other.type == type)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,userId,status,type,createdAt,updatedAt);

@override
String toString() {
  return 'PlanSubscription(id: $id, planId: $planId, userId: $userId, status: $status, type: $type, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$PlanSubscriptionCopyWith<$Res> implements $PlanSubscriptionCopyWith<$Res> {
  factory _$PlanSubscriptionCopyWith(_PlanSubscription value, $Res Function(_PlanSubscription) _then) = __$PlanSubscriptionCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'userId') String userId,@PlanSubscriptionStatusConverter() PlanSubscriptionStatus status,@PlanSubscriptionTypeConverter() PlanSubscriptionType type,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class __$PlanSubscriptionCopyWithImpl<$Res>
    implements _$PlanSubscriptionCopyWith<$Res> {
  __$PlanSubscriptionCopyWithImpl(this._self, this._then);

  final _PlanSubscription _self;
  final $Res Function(_PlanSubscription) _then;

/// Create a copy of PlanSubscription
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? planId = null,Object? userId = null,Object? status = null,Object? type = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_PlanSubscription(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$SubscribeRequest {

@PlanSubscriptionTypeConverter() PlanSubscriptionType get type;
/// Create a copy of SubscribeRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$SubscribeRequestCopyWith<SubscribeRequest> get copyWith => _$SubscribeRequestCopyWithImpl<SubscribeRequest>(this as SubscribeRequest, _$identity);

  /// Serializes this SubscribeRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is SubscribeRequest&&(identical(other.type, type) || other.type == type));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type);

@override
String toString() {
  return 'SubscribeRequest(type: $type)';
}


}

/// @nodoc
abstract mixin class $SubscribeRequestCopyWith<$Res>  {
  factory $SubscribeRequestCopyWith(SubscribeRequest value, $Res Function(SubscribeRequest) _then) = _$SubscribeRequestCopyWithImpl;
@useResult
$Res call({
@PlanSubscriptionTypeConverter() PlanSubscriptionType type
});




}
/// @nodoc
class _$SubscribeRequestCopyWithImpl<$Res>
    implements $SubscribeRequestCopyWith<$Res> {
  _$SubscribeRequestCopyWithImpl(this._self, this._then);

  final SubscribeRequest _self;
  final $Res Function(SubscribeRequest) _then;

/// Create a copy of SubscribeRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? type = null,}) {
  return _then(_self.copyWith(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,
  ));
}

}


/// Adds pattern-matching-related methods to [SubscribeRequest].
extension SubscribeRequestPatterns on SubscribeRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _SubscribeRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _SubscribeRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _SubscribeRequest value)  $default,){
final _that = this;
switch (_that) {
case _SubscribeRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _SubscribeRequest value)?  $default,){
final _that = this;
switch (_that) {
case _SubscribeRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _SubscribeRequest() when $default != null:
return $default(_that.type);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)  $default,) {final _that = this;
switch (_that) {
case _SubscribeRequest():
return $default(_that.type);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)?  $default,) {final _that = this;
switch (_that) {
case _SubscribeRequest() when $default != null:
return $default(_that.type);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _SubscribeRequest implements SubscribeRequest {
  const _SubscribeRequest({@PlanSubscriptionTypeConverter() required this.type});
  factory _SubscribeRequest.fromJson(Map<String, dynamic> json) => _$SubscribeRequestFromJson(json);

@override@PlanSubscriptionTypeConverter() final  PlanSubscriptionType type;

/// Create a copy of SubscribeRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$SubscribeRequestCopyWith<_SubscribeRequest> get copyWith => __$SubscribeRequestCopyWithImpl<_SubscribeRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$SubscribeRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _SubscribeRequest&&(identical(other.type, type) || other.type == type));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type);

@override
String toString() {
  return 'SubscribeRequest(type: $type)';
}


}

/// @nodoc
abstract mixin class _$SubscribeRequestCopyWith<$Res> implements $SubscribeRequestCopyWith<$Res> {
  factory _$SubscribeRequestCopyWith(_SubscribeRequest value, $Res Function(_SubscribeRequest) _then) = __$SubscribeRequestCopyWithImpl;
@override @useResult
$Res call({
@PlanSubscriptionTypeConverter() PlanSubscriptionType type
});




}
/// @nodoc
class __$SubscribeRequestCopyWithImpl<$Res>
    implements _$SubscribeRequestCopyWith<$Res> {
  __$SubscribeRequestCopyWithImpl(this._self, this._then);

  final _SubscribeRequest _self;
  final $Res Function(_SubscribeRequest) _then;

/// Create a copy of SubscribeRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? type = null,}) {
  return _then(_SubscribeRequest(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,
  ));
}


}


/// @nodoc
mixin _$ChangeSubscriptionStatusRequest {

@PlanSubscriptionStatusConverter() PlanSubscriptionStatus get status;
/// Create a copy of ChangeSubscriptionStatusRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ChangeSubscriptionStatusRequestCopyWith<ChangeSubscriptionStatusRequest> get copyWith => _$ChangeSubscriptionStatusRequestCopyWithImpl<ChangeSubscriptionStatusRequest>(this as ChangeSubscriptionStatusRequest, _$identity);

  /// Serializes this ChangeSubscriptionStatusRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is ChangeSubscriptionStatusRequest&&(identical(other.status, status) || other.status == status));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,status);

@override
String toString() {
  return 'ChangeSubscriptionStatusRequest(status: $status)';
}


}

/// @nodoc
abstract mixin class $ChangeSubscriptionStatusRequestCopyWith<$Res>  {
  factory $ChangeSubscriptionStatusRequestCopyWith(ChangeSubscriptionStatusRequest value, $Res Function(ChangeSubscriptionStatusRequest) _then) = _$ChangeSubscriptionStatusRequestCopyWithImpl;
@useResult
$Res call({
@PlanSubscriptionStatusConverter() PlanSubscriptionStatus status
});




}
/// @nodoc
class _$ChangeSubscriptionStatusRequestCopyWithImpl<$Res>
    implements $ChangeSubscriptionStatusRequestCopyWith<$Res> {
  _$ChangeSubscriptionStatusRequestCopyWithImpl(this._self, this._then);

  final ChangeSubscriptionStatusRequest _self;
  final $Res Function(ChangeSubscriptionStatusRequest) _then;

/// Create a copy of ChangeSubscriptionStatusRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? status = null,}) {
  return _then(_self.copyWith(
status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus,
  ));
}

}


/// Adds pattern-matching-related methods to [ChangeSubscriptionStatusRequest].
extension ChangeSubscriptionStatusRequestPatterns on ChangeSubscriptionStatusRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _ChangeSubscriptionStatusRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _ChangeSubscriptionStatusRequest value)  $default,){
final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _ChangeSubscriptionStatusRequest value)?  $default,){
final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest() when $default != null:
return $default(_that.status);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status)  $default,) {final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest():
return $default(_that.status);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@PlanSubscriptionStatusConverter()  PlanSubscriptionStatus status)?  $default,) {final _that = this;
switch (_that) {
case _ChangeSubscriptionStatusRequest() when $default != null:
return $default(_that.status);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _ChangeSubscriptionStatusRequest implements ChangeSubscriptionStatusRequest {
  const _ChangeSubscriptionStatusRequest({@PlanSubscriptionStatusConverter() required this.status});
  factory _ChangeSubscriptionStatusRequest.fromJson(Map<String, dynamic> json) => _$ChangeSubscriptionStatusRequestFromJson(json);

@override@PlanSubscriptionStatusConverter() final  PlanSubscriptionStatus status;

/// Create a copy of ChangeSubscriptionStatusRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ChangeSubscriptionStatusRequestCopyWith<_ChangeSubscriptionStatusRequest> get copyWith => __$ChangeSubscriptionStatusRequestCopyWithImpl<_ChangeSubscriptionStatusRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ChangeSubscriptionStatusRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _ChangeSubscriptionStatusRequest&&(identical(other.status, status) || other.status == status));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,status);

@override
String toString() {
  return 'ChangeSubscriptionStatusRequest(status: $status)';
}


}

/// @nodoc
abstract mixin class _$ChangeSubscriptionStatusRequestCopyWith<$Res> implements $ChangeSubscriptionStatusRequestCopyWith<$Res> {
  factory _$ChangeSubscriptionStatusRequestCopyWith(_ChangeSubscriptionStatusRequest value, $Res Function(_ChangeSubscriptionStatusRequest) _then) = __$ChangeSubscriptionStatusRequestCopyWithImpl;
@override @useResult
$Res call({
@PlanSubscriptionStatusConverter() PlanSubscriptionStatus status
});




}
/// @nodoc
class __$ChangeSubscriptionStatusRequestCopyWithImpl<$Res>
    implements _$ChangeSubscriptionStatusRequestCopyWith<$Res> {
  __$ChangeSubscriptionStatusRequestCopyWithImpl(this._self, this._then);

  final _ChangeSubscriptionStatusRequest _self;
  final $Res Function(_ChangeSubscriptionStatusRequest) _then;

/// Create a copy of ChangeSubscriptionStatusRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? status = null,}) {
  return _then(_ChangeSubscriptionStatusRequest(
status: null == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus,
  ));
}


}


/// @nodoc
mixin _$ChangeSubscriptionPrivacyRequest {

@PlanSubscriptionTypeConverter() PlanSubscriptionType get type;
/// Create a copy of ChangeSubscriptionPrivacyRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ChangeSubscriptionPrivacyRequestCopyWith<ChangeSubscriptionPrivacyRequest> get copyWith => _$ChangeSubscriptionPrivacyRequestCopyWithImpl<ChangeSubscriptionPrivacyRequest>(this as ChangeSubscriptionPrivacyRequest, _$identity);

  /// Serializes this ChangeSubscriptionPrivacyRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is ChangeSubscriptionPrivacyRequest&&(identical(other.type, type) || other.type == type));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type);

@override
String toString() {
  return 'ChangeSubscriptionPrivacyRequest(type: $type)';
}


}

/// @nodoc
abstract mixin class $ChangeSubscriptionPrivacyRequestCopyWith<$Res>  {
  factory $ChangeSubscriptionPrivacyRequestCopyWith(ChangeSubscriptionPrivacyRequest value, $Res Function(ChangeSubscriptionPrivacyRequest) _then) = _$ChangeSubscriptionPrivacyRequestCopyWithImpl;
@useResult
$Res call({
@PlanSubscriptionTypeConverter() PlanSubscriptionType type
});




}
/// @nodoc
class _$ChangeSubscriptionPrivacyRequestCopyWithImpl<$Res>
    implements $ChangeSubscriptionPrivacyRequestCopyWith<$Res> {
  _$ChangeSubscriptionPrivacyRequestCopyWithImpl(this._self, this._then);

  final ChangeSubscriptionPrivacyRequest _self;
  final $Res Function(ChangeSubscriptionPrivacyRequest) _then;

/// Create a copy of ChangeSubscriptionPrivacyRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? type = null,}) {
  return _then(_self.copyWith(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,
  ));
}

}


/// Adds pattern-matching-related methods to [ChangeSubscriptionPrivacyRequest].
extension ChangeSubscriptionPrivacyRequestPatterns on ChangeSubscriptionPrivacyRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _ChangeSubscriptionPrivacyRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _ChangeSubscriptionPrivacyRequest value)  $default,){
final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _ChangeSubscriptionPrivacyRequest value)?  $default,){
final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest() when $default != null:
return $default(_that.type);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)  $default,) {final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest():
return $default(_that.type);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@PlanSubscriptionTypeConverter()  PlanSubscriptionType type)?  $default,) {final _that = this;
switch (_that) {
case _ChangeSubscriptionPrivacyRequest() when $default != null:
return $default(_that.type);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _ChangeSubscriptionPrivacyRequest implements ChangeSubscriptionPrivacyRequest {
  const _ChangeSubscriptionPrivacyRequest({@PlanSubscriptionTypeConverter() required this.type});
  factory _ChangeSubscriptionPrivacyRequest.fromJson(Map<String, dynamic> json) => _$ChangeSubscriptionPrivacyRequestFromJson(json);

@override@PlanSubscriptionTypeConverter() final  PlanSubscriptionType type;

/// Create a copy of ChangeSubscriptionPrivacyRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ChangeSubscriptionPrivacyRequestCopyWith<_ChangeSubscriptionPrivacyRequest> get copyWith => __$ChangeSubscriptionPrivacyRequestCopyWithImpl<_ChangeSubscriptionPrivacyRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ChangeSubscriptionPrivacyRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _ChangeSubscriptionPrivacyRequest&&(identical(other.type, type) || other.type == type));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type);

@override
String toString() {
  return 'ChangeSubscriptionPrivacyRequest(type: $type)';
}


}

/// @nodoc
abstract mixin class _$ChangeSubscriptionPrivacyRequestCopyWith<$Res> implements $ChangeSubscriptionPrivacyRequestCopyWith<$Res> {
  factory _$ChangeSubscriptionPrivacyRequestCopyWith(_ChangeSubscriptionPrivacyRequest value, $Res Function(_ChangeSubscriptionPrivacyRequest) _then) = __$ChangeSubscriptionPrivacyRequestCopyWithImpl;
@override @useResult
$Res call({
@PlanSubscriptionTypeConverter() PlanSubscriptionType type
});




}
/// @nodoc
class __$ChangeSubscriptionPrivacyRequestCopyWithImpl<$Res>
    implements _$ChangeSubscriptionPrivacyRequestCopyWith<$Res> {
  __$ChangeSubscriptionPrivacyRequestCopyWithImpl(this._self, this._then);

  final _ChangeSubscriptionPrivacyRequest _self;
  final $Res Function(_ChangeSubscriptionPrivacyRequest) _then;

/// Create a copy of ChangeSubscriptionPrivacyRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? type = null,}) {
  return _then(_ChangeSubscriptionPrivacyRequest(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType,
  ));
}


}


/// @nodoc
mixin _$ListSubscriptionFilters {

@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true) PlanSubscriptionStatus? get status;@PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true) PlanSubscriptionType? get type;@TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType') TrainingPlanType? get planType;@TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true) TrainingPlanVisibility? get visibility;@TrainingPlanLevelConverter()@JsonKey(includeIfNull: true) TrainingPlanLevel? get level;@JsonKey(includeIfNull: true, name: 'authorId') String? get authorId;
/// Create a copy of ListSubscriptionFilters
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ListSubscriptionFiltersCopyWith<ListSubscriptionFilters> get copyWith => _$ListSubscriptionFiltersCopyWithImpl<ListSubscriptionFilters>(this as ListSubscriptionFilters, _$identity);

  /// Serializes this ListSubscriptionFilters to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is ListSubscriptionFilters&&(identical(other.status, status) || other.status == status)&&(identical(other.type, type) || other.type == type)&&(identical(other.planType, planType) || other.planType == planType)&&(identical(other.visibility, visibility) || other.visibility == visibility)&&(identical(other.level, level) || other.level == level)&&(identical(other.authorId, authorId) || other.authorId == authorId));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,status,type,planType,visibility,level,authorId);

@override
String toString() {
  return 'ListSubscriptionFilters(status: $status, type: $type, planType: $planType, visibility: $visibility, level: $level, authorId: $authorId)';
}


}

/// @nodoc
abstract mixin class $ListSubscriptionFiltersCopyWith<$Res>  {
  factory $ListSubscriptionFiltersCopyWith(ListSubscriptionFilters value, $Res Function(ListSubscriptionFilters) _then) = _$ListSubscriptionFiltersCopyWithImpl;
@useResult
$Res call({
@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true) PlanSubscriptionStatus? status,@PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true) PlanSubscriptionType? type,@TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType') TrainingPlanType? planType,@TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true) TrainingPlanVisibility? visibility,@TrainingPlanLevelConverter()@JsonKey(includeIfNull: true) TrainingPlanLevel? level,@JsonKey(includeIfNull: true, name: 'authorId') String? authorId
});




}
/// @nodoc
class _$ListSubscriptionFiltersCopyWithImpl<$Res>
    implements $ListSubscriptionFiltersCopyWith<$Res> {
  _$ListSubscriptionFiltersCopyWithImpl(this._self, this._then);

  final ListSubscriptionFilters _self;
  final $Res Function(ListSubscriptionFilters) _then;

/// Create a copy of ListSubscriptionFilters
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? status = freezed,Object? type = freezed,Object? planType = freezed,Object? visibility = freezed,Object? level = freezed,Object? authorId = freezed,}) {
  return _then(_self.copyWith(
status: freezed == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus?,type: freezed == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType?,planType: freezed == planType ? _self.planType : planType // ignore: cast_nullable_to_non_nullable
as TrainingPlanType?,visibility: freezed == visibility ? _self.visibility : visibility // ignore: cast_nullable_to_non_nullable
as TrainingPlanVisibility?,level: freezed == level ? _self.level : level // ignore: cast_nullable_to_non_nullable
as TrainingPlanLevel?,authorId: freezed == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}

}


/// Adds pattern-matching-related methods to [ListSubscriptionFilters].
extension ListSubscriptionFiltersPatterns on ListSubscriptionFilters {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _ListSubscriptionFilters value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _ListSubscriptionFilters() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _ListSubscriptionFilters value)  $default,){
final _that = this;
switch (_that) {
case _ListSubscriptionFilters():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _ListSubscriptionFilters value)?  $default,){
final _that = this;
switch (_that) {
case _ListSubscriptionFilters() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionStatus? status, @PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionType? type, @TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType')  TrainingPlanType? planType, @TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true)  TrainingPlanVisibility? visibility, @TrainingPlanLevelConverter()@JsonKey(includeIfNull: true)  TrainingPlanLevel? level, @JsonKey(includeIfNull: true, name: 'authorId')  String? authorId)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _ListSubscriptionFilters() when $default != null:
return $default(_that.status,_that.type,_that.planType,_that.visibility,_that.level,_that.authorId);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionStatus? status, @PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionType? type, @TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType')  TrainingPlanType? planType, @TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true)  TrainingPlanVisibility? visibility, @TrainingPlanLevelConverter()@JsonKey(includeIfNull: true)  TrainingPlanLevel? level, @JsonKey(includeIfNull: true, name: 'authorId')  String? authorId)  $default,) {final _that = this;
switch (_that) {
case _ListSubscriptionFilters():
return $default(_that.status,_that.type,_that.planType,_that.visibility,_that.level,_that.authorId);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionStatus? status, @PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true)  PlanSubscriptionType? type, @TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType')  TrainingPlanType? planType, @TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true)  TrainingPlanVisibility? visibility, @TrainingPlanLevelConverter()@JsonKey(includeIfNull: true)  TrainingPlanLevel? level, @JsonKey(includeIfNull: true, name: 'authorId')  String? authorId)?  $default,) {final _that = this;
switch (_that) {
case _ListSubscriptionFilters() when $default != null:
return $default(_that.status,_that.type,_that.planType,_that.visibility,_that.level,_that.authorId);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _ListSubscriptionFilters implements ListSubscriptionFilters {
  const _ListSubscriptionFilters({@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true) this.status, @PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true) this.type, @TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType') this.planType, @TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true) this.visibility, @TrainingPlanLevelConverter()@JsonKey(includeIfNull: true) this.level, @JsonKey(includeIfNull: true, name: 'authorId') this.authorId});
  factory _ListSubscriptionFilters.fromJson(Map<String, dynamic> json) => _$ListSubscriptionFiltersFromJson(json);

@override@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true) final  PlanSubscriptionStatus? status;
@override@PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true) final  PlanSubscriptionType? type;
@override@TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType') final  TrainingPlanType? planType;
@override@TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true) final  TrainingPlanVisibility? visibility;
@override@TrainingPlanLevelConverter()@JsonKey(includeIfNull: true) final  TrainingPlanLevel? level;
@override@JsonKey(includeIfNull: true, name: 'authorId') final  String? authorId;

/// Create a copy of ListSubscriptionFilters
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ListSubscriptionFiltersCopyWith<_ListSubscriptionFilters> get copyWith => __$ListSubscriptionFiltersCopyWithImpl<_ListSubscriptionFilters>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ListSubscriptionFiltersToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _ListSubscriptionFilters&&(identical(other.status, status) || other.status == status)&&(identical(other.type, type) || other.type == type)&&(identical(other.planType, planType) || other.planType == planType)&&(identical(other.visibility, visibility) || other.visibility == visibility)&&(identical(other.level, level) || other.level == level)&&(identical(other.authorId, authorId) || other.authorId == authorId));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,status,type,planType,visibility,level,authorId);

@override
String toString() {
  return 'ListSubscriptionFilters(status: $status, type: $type, planType: $planType, visibility: $visibility, level: $level, authorId: $authorId)';
}


}

/// @nodoc
abstract mixin class _$ListSubscriptionFiltersCopyWith<$Res> implements $ListSubscriptionFiltersCopyWith<$Res> {
  factory _$ListSubscriptionFiltersCopyWith(_ListSubscriptionFilters value, $Res Function(_ListSubscriptionFilters) _then) = __$ListSubscriptionFiltersCopyWithImpl;
@override @useResult
$Res call({
@PlanSubscriptionStatusConverter()@JsonKey(includeIfNull: true) PlanSubscriptionStatus? status,@PlanSubscriptionTypeConverter()@JsonKey(includeIfNull: true) PlanSubscriptionType? type,@TrainingPlanTypeConverter()@JsonKey(includeIfNull: true, name: 'planType') TrainingPlanType? planType,@TrainingPlanVisibilityConverter()@JsonKey(includeIfNull: true) TrainingPlanVisibility? visibility,@TrainingPlanLevelConverter()@JsonKey(includeIfNull: true) TrainingPlanLevel? level,@JsonKey(includeIfNull: true, name: 'authorId') String? authorId
});




}
/// @nodoc
class __$ListSubscriptionFiltersCopyWithImpl<$Res>
    implements _$ListSubscriptionFiltersCopyWith<$Res> {
  __$ListSubscriptionFiltersCopyWithImpl(this._self, this._then);

  final _ListSubscriptionFilters _self;
  final $Res Function(_ListSubscriptionFilters) _then;

/// Create a copy of ListSubscriptionFilters
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? status = freezed,Object? type = freezed,Object? planType = freezed,Object? visibility = freezed,Object? level = freezed,Object? authorId = freezed,}) {
  return _then(_ListSubscriptionFilters(
status: freezed == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionStatus?,type: freezed == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as PlanSubscriptionType?,planType: freezed == planType ? _self.planType : planType // ignore: cast_nullable_to_non_nullable
as TrainingPlanType?,visibility: freezed == visibility ? _self.visibility : visibility // ignore: cast_nullable_to_non_nullable
as TrainingPlanVisibility?,level: freezed == level ? _self.level : level // ignore: cast_nullable_to_non_nullable
as TrainingPlanLevel?,authorId: freezed == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}


}


/// @nodoc
mixin _$WeeklyDayProgress {

 List<PlanDayProgress> get days;@JsonKey(name: 'completedDays') int get completedDays;@JsonKey(name: 'totalDays') int get totalDays;
/// Create a copy of WeeklyDayProgress
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$WeeklyDayProgressCopyWith<WeeklyDayProgress> get copyWith => _$WeeklyDayProgressCopyWithImpl<WeeklyDayProgress>(this as WeeklyDayProgress, _$identity);

  /// Serializes this WeeklyDayProgress to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is WeeklyDayProgress&&const DeepCollectionEquality().equals(other.days, days)&&(identical(other.completedDays, completedDays) || other.completedDays == completedDays)&&(identical(other.totalDays, totalDays) || other.totalDays == totalDays));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(days),completedDays,totalDays);

@override
String toString() {
  return 'WeeklyDayProgress(days: $days, completedDays: $completedDays, totalDays: $totalDays)';
}


}

/// @nodoc
abstract mixin class $WeeklyDayProgressCopyWith<$Res>  {
  factory $WeeklyDayProgressCopyWith(WeeklyDayProgress value, $Res Function(WeeklyDayProgress) _then) = _$WeeklyDayProgressCopyWithImpl;
@useResult
$Res call({
 List<PlanDayProgress> days,@JsonKey(name: 'completedDays') int completedDays,@JsonKey(name: 'totalDays') int totalDays
});




}
/// @nodoc
class _$WeeklyDayProgressCopyWithImpl<$Res>
    implements $WeeklyDayProgressCopyWith<$Res> {
  _$WeeklyDayProgressCopyWithImpl(this._self, this._then);

  final WeeklyDayProgress _self;
  final $Res Function(WeeklyDayProgress) _then;

/// Create a copy of WeeklyDayProgress
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? days = null,Object? completedDays = null,Object? totalDays = null,}) {
  return _then(_self.copyWith(
days: null == days ? _self.days : days // ignore: cast_nullable_to_non_nullable
as List<PlanDayProgress>,completedDays: null == completedDays ? _self.completedDays : completedDays // ignore: cast_nullable_to_non_nullable
as int,totalDays: null == totalDays ? _self.totalDays : totalDays // ignore: cast_nullable_to_non_nullable
as int,
  ));
}

}


/// Adds pattern-matching-related methods to [WeeklyDayProgress].
extension WeeklyDayProgressPatterns on WeeklyDayProgress {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _WeeklyDayProgress value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _WeeklyDayProgress() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _WeeklyDayProgress value)  $default,){
final _that = this;
switch (_that) {
case _WeeklyDayProgress():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _WeeklyDayProgress value)?  $default,){
final _that = this;
switch (_that) {
case _WeeklyDayProgress() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<PlanDayProgress> days, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _WeeklyDayProgress() when $default != null:
return $default(_that.days,_that.completedDays,_that.totalDays);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<PlanDayProgress> days, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)  $default,) {final _that = this;
switch (_that) {
case _WeeklyDayProgress():
return $default(_that.days,_that.completedDays,_that.totalDays);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<PlanDayProgress> days, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)?  $default,) {final _that = this;
switch (_that) {
case _WeeklyDayProgress() when $default != null:
return $default(_that.days,_that.completedDays,_that.totalDays);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _WeeklyDayProgress implements WeeklyDayProgress {
  const _WeeklyDayProgress({final  List<PlanDayProgress> days = const [], @JsonKey(name: 'completedDays') this.completedDays = 0, @JsonKey(name: 'totalDays') this.totalDays = 0}): _days = days;
  factory _WeeklyDayProgress.fromJson(Map<String, dynamic> json) => _$WeeklyDayProgressFromJson(json);

 final  List<PlanDayProgress> _days;
@override@JsonKey() List<PlanDayProgress> get days {
  if (_days is EqualUnmodifiableListView) return _days;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_days);
}

@override@JsonKey(name: 'completedDays') final  int completedDays;
@override@JsonKey(name: 'totalDays') final  int totalDays;

/// Create a copy of WeeklyDayProgress
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$WeeklyDayProgressCopyWith<_WeeklyDayProgress> get copyWith => __$WeeklyDayProgressCopyWithImpl<_WeeklyDayProgress>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$WeeklyDayProgressToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _WeeklyDayProgress&&const DeepCollectionEquality().equals(other._days, _days)&&(identical(other.completedDays, completedDays) || other.completedDays == completedDays)&&(identical(other.totalDays, totalDays) || other.totalDays == totalDays));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_days),completedDays,totalDays);

@override
String toString() {
  return 'WeeklyDayProgress(days: $days, completedDays: $completedDays, totalDays: $totalDays)';
}


}

/// @nodoc
abstract mixin class _$WeeklyDayProgressCopyWith<$Res> implements $WeeklyDayProgressCopyWith<$Res> {
  factory _$WeeklyDayProgressCopyWith(_WeeklyDayProgress value, $Res Function(_WeeklyDayProgress) _then) = __$WeeklyDayProgressCopyWithImpl;
@override @useResult
$Res call({
 List<PlanDayProgress> days,@JsonKey(name: 'completedDays') int completedDays,@JsonKey(name: 'totalDays') int totalDays
});




}
/// @nodoc
class __$WeeklyDayProgressCopyWithImpl<$Res>
    implements _$WeeklyDayProgressCopyWith<$Res> {
  __$WeeklyDayProgressCopyWithImpl(this._self, this._then);

  final _WeeklyDayProgress _self;
  final $Res Function(_WeeklyDayProgress) _then;

/// Create a copy of WeeklyDayProgress
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? days = null,Object? completedDays = null,Object? totalDays = null,}) {
  return _then(_WeeklyDayProgress(
days: null == days ? _self._days : days // ignore: cast_nullable_to_non_nullable
as List<PlanDayProgress>,completedDays: null == completedDays ? _self.completedDays : completedDays // ignore: cast_nullable_to_non_nullable
as int,totalDays: null == totalDays ? _self.totalDays : totalDays // ignore: cast_nullable_to_non_nullable
as int,
  ));
}


}


/// @nodoc
mixin _$PlanDayProgress {

@JsonKey(name: 'subscriptionId') String get subscriptionId;@JsonKey(name: 'dayId') String get dayId;@JsonKey(name: 'dayName') String get dayName;@JsonKey(name: 'planId') String get planId;@JsonKey(name: 'planTitle') String get planTitle; String? get status;@JsonKey(name: 'scheduledDate') DateTime? get scheduledDate;@JsonKey(name: 'startedAt') DateTime? get startedAt;@JsonKey(name: 'completedAt') DateTime? get completedAt;
/// Create a copy of PlanDayProgress
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$PlanDayProgressCopyWith<PlanDayProgress> get copyWith => _$PlanDayProgressCopyWithImpl<PlanDayProgress>(this as PlanDayProgress, _$identity);

  /// Serializes this PlanDayProgress to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is PlanDayProgress&&(identical(other.subscriptionId, subscriptionId) || other.subscriptionId == subscriptionId)&&(identical(other.dayId, dayId) || other.dayId == dayId)&&(identical(other.dayName, dayName) || other.dayName == dayName)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.planTitle, planTitle) || other.planTitle == planTitle)&&(identical(other.status, status) || other.status == status)&&(identical(other.scheduledDate, scheduledDate) || other.scheduledDate == scheduledDate)&&(identical(other.startedAt, startedAt) || other.startedAt == startedAt)&&(identical(other.completedAt, completedAt) || other.completedAt == completedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,subscriptionId,dayId,dayName,planId,planTitle,status,scheduledDate,startedAt,completedAt);

@override
String toString() {
  return 'PlanDayProgress(subscriptionId: $subscriptionId, dayId: $dayId, dayName: $dayName, planId: $planId, planTitle: $planTitle, status: $status, scheduledDate: $scheduledDate, startedAt: $startedAt, completedAt: $completedAt)';
}


}

/// @nodoc
abstract mixin class $PlanDayProgressCopyWith<$Res>  {
  factory $PlanDayProgressCopyWith(PlanDayProgress value, $Res Function(PlanDayProgress) _then) = _$PlanDayProgressCopyWithImpl;
@useResult
$Res call({
@JsonKey(name: 'subscriptionId') String subscriptionId,@JsonKey(name: 'dayId') String dayId,@JsonKey(name: 'dayName') String dayName,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'planTitle') String planTitle, String? status,@JsonKey(name: 'scheduledDate') DateTime? scheduledDate,@JsonKey(name: 'startedAt') DateTime? startedAt,@JsonKey(name: 'completedAt') DateTime? completedAt
});




}
/// @nodoc
class _$PlanDayProgressCopyWithImpl<$Res>
    implements $PlanDayProgressCopyWith<$Res> {
  _$PlanDayProgressCopyWithImpl(this._self, this._then);

  final PlanDayProgress _self;
  final $Res Function(PlanDayProgress) _then;

/// Create a copy of PlanDayProgress
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? subscriptionId = null,Object? dayId = null,Object? dayName = null,Object? planId = null,Object? planTitle = null,Object? status = freezed,Object? scheduledDate = freezed,Object? startedAt = freezed,Object? completedAt = freezed,}) {
  return _then(_self.copyWith(
subscriptionId: null == subscriptionId ? _self.subscriptionId : subscriptionId // ignore: cast_nullable_to_non_nullable
as String,dayId: null == dayId ? _self.dayId : dayId // ignore: cast_nullable_to_non_nullable
as String,dayName: null == dayName ? _self.dayName : dayName // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,planTitle: null == planTitle ? _self.planTitle : planTitle // ignore: cast_nullable_to_non_nullable
as String,status: freezed == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as String?,scheduledDate: freezed == scheduledDate ? _self.scheduledDate : scheduledDate // ignore: cast_nullable_to_non_nullable
as DateTime?,startedAt: freezed == startedAt ? _self.startedAt : startedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,completedAt: freezed == completedAt ? _self.completedAt : completedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,
  ));
}

}


/// Adds pattern-matching-related methods to [PlanDayProgress].
extension PlanDayProgressPatterns on PlanDayProgress {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _PlanDayProgress value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _PlanDayProgress() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _PlanDayProgress value)  $default,){
final _that = this;
switch (_that) {
case _PlanDayProgress():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _PlanDayProgress value)?  $default,){
final _that = this;
switch (_that) {
case _PlanDayProgress() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@JsonKey(name: 'subscriptionId')  String subscriptionId, @JsonKey(name: 'dayId')  String dayId, @JsonKey(name: 'dayName')  String dayName, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'planTitle')  String planTitle,  String? status, @JsonKey(name: 'scheduledDate')  DateTime? scheduledDate, @JsonKey(name: 'startedAt')  DateTime? startedAt, @JsonKey(name: 'completedAt')  DateTime? completedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _PlanDayProgress() when $default != null:
return $default(_that.subscriptionId,_that.dayId,_that.dayName,_that.planId,_that.planTitle,_that.status,_that.scheduledDate,_that.startedAt,_that.completedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@JsonKey(name: 'subscriptionId')  String subscriptionId, @JsonKey(name: 'dayId')  String dayId, @JsonKey(name: 'dayName')  String dayName, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'planTitle')  String planTitle,  String? status, @JsonKey(name: 'scheduledDate')  DateTime? scheduledDate, @JsonKey(name: 'startedAt')  DateTime? startedAt, @JsonKey(name: 'completedAt')  DateTime? completedAt)  $default,) {final _that = this;
switch (_that) {
case _PlanDayProgress():
return $default(_that.subscriptionId,_that.dayId,_that.dayName,_that.planId,_that.planTitle,_that.status,_that.scheduledDate,_that.startedAt,_that.completedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@JsonKey(name: 'subscriptionId')  String subscriptionId, @JsonKey(name: 'dayId')  String dayId, @JsonKey(name: 'dayName')  String dayName, @JsonKey(name: 'planId')  String planId, @JsonKey(name: 'planTitle')  String planTitle,  String? status, @JsonKey(name: 'scheduledDate')  DateTime? scheduledDate, @JsonKey(name: 'startedAt')  DateTime? startedAt, @JsonKey(name: 'completedAt')  DateTime? completedAt)?  $default,) {final _that = this;
switch (_that) {
case _PlanDayProgress() when $default != null:
return $default(_that.subscriptionId,_that.dayId,_that.dayName,_that.planId,_that.planTitle,_that.status,_that.scheduledDate,_that.startedAt,_that.completedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _PlanDayProgress implements PlanDayProgress {
  const _PlanDayProgress({@JsonKey(name: 'subscriptionId') required this.subscriptionId, @JsonKey(name: 'dayId') required this.dayId, @JsonKey(name: 'dayName') required this.dayName, @JsonKey(name: 'planId') required this.planId, @JsonKey(name: 'planTitle') required this.planTitle, this.status, @JsonKey(name: 'scheduledDate') this.scheduledDate, @JsonKey(name: 'startedAt') this.startedAt, @JsonKey(name: 'completedAt') this.completedAt});
  factory _PlanDayProgress.fromJson(Map<String, dynamic> json) => _$PlanDayProgressFromJson(json);

@override@JsonKey(name: 'subscriptionId') final  String subscriptionId;
@override@JsonKey(name: 'dayId') final  String dayId;
@override@JsonKey(name: 'dayName') final  String dayName;
@override@JsonKey(name: 'planId') final  String planId;
@override@JsonKey(name: 'planTitle') final  String planTitle;
@override final  String? status;
@override@JsonKey(name: 'scheduledDate') final  DateTime? scheduledDate;
@override@JsonKey(name: 'startedAt') final  DateTime? startedAt;
@override@JsonKey(name: 'completedAt') final  DateTime? completedAt;

/// Create a copy of PlanDayProgress
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$PlanDayProgressCopyWith<_PlanDayProgress> get copyWith => __$PlanDayProgressCopyWithImpl<_PlanDayProgress>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$PlanDayProgressToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _PlanDayProgress&&(identical(other.subscriptionId, subscriptionId) || other.subscriptionId == subscriptionId)&&(identical(other.dayId, dayId) || other.dayId == dayId)&&(identical(other.dayName, dayName) || other.dayName == dayName)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.planTitle, planTitle) || other.planTitle == planTitle)&&(identical(other.status, status) || other.status == status)&&(identical(other.scheduledDate, scheduledDate) || other.scheduledDate == scheduledDate)&&(identical(other.startedAt, startedAt) || other.startedAt == startedAt)&&(identical(other.completedAt, completedAt) || other.completedAt == completedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,subscriptionId,dayId,dayName,planId,planTitle,status,scheduledDate,startedAt,completedAt);

@override
String toString() {
  return 'PlanDayProgress(subscriptionId: $subscriptionId, dayId: $dayId, dayName: $dayName, planId: $planId, planTitle: $planTitle, status: $status, scheduledDate: $scheduledDate, startedAt: $startedAt, completedAt: $completedAt)';
}


}

/// @nodoc
abstract mixin class _$PlanDayProgressCopyWith<$Res> implements $PlanDayProgressCopyWith<$Res> {
  factory _$PlanDayProgressCopyWith(_PlanDayProgress value, $Res Function(_PlanDayProgress) _then) = __$PlanDayProgressCopyWithImpl;
@override @useResult
$Res call({
@JsonKey(name: 'subscriptionId') String subscriptionId,@JsonKey(name: 'dayId') String dayId,@JsonKey(name: 'dayName') String dayName,@JsonKey(name: 'planId') String planId,@JsonKey(name: 'planTitle') String planTitle, String? status,@JsonKey(name: 'scheduledDate') DateTime? scheduledDate,@JsonKey(name: 'startedAt') DateTime? startedAt,@JsonKey(name: 'completedAt') DateTime? completedAt
});




}
/// @nodoc
class __$PlanDayProgressCopyWithImpl<$Res>
    implements _$PlanDayProgressCopyWith<$Res> {
  __$PlanDayProgressCopyWithImpl(this._self, this._then);

  final _PlanDayProgress _self;
  final $Res Function(_PlanDayProgress) _then;

/// Create a copy of PlanDayProgress
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? subscriptionId = null,Object? dayId = null,Object? dayName = null,Object? planId = null,Object? planTitle = null,Object? status = freezed,Object? scheduledDate = freezed,Object? startedAt = freezed,Object? completedAt = freezed,}) {
  return _then(_PlanDayProgress(
subscriptionId: null == subscriptionId ? _self.subscriptionId : subscriptionId // ignore: cast_nullable_to_non_nullable
as String,dayId: null == dayId ? _self.dayId : dayId // ignore: cast_nullable_to_non_nullable
as String,dayName: null == dayName ? _self.dayName : dayName // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,planTitle: null == planTitle ? _self.planTitle : planTitle // ignore: cast_nullable_to_non_nullable
as String,status: freezed == status ? _self.status : status // ignore: cast_nullable_to_non_nullable
as String?,scheduledDate: freezed == scheduledDate ? _self.scheduledDate : scheduledDate // ignore: cast_nullable_to_non_nullable
as DateTime?,startedAt: freezed == startedAt ? _self.startedAt : startedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,completedAt: freezed == completedAt ? _self.completedAt : completedAt // ignore: cast_nullable_to_non_nullable
as DateTime?,
  ));
}


}

// dart format on

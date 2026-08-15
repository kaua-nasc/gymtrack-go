// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'weight_log.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$WeightLog {

 String get id; double get weight; String? get note;@JsonKey(name: 'userId') String get userId;@JsonKey(name: 'trainerId') String? get trainerId;@JsonKey(name: 'measuredAt') DateTime get measuredAt;@JsonKey(name: 'createdAt') DateTime get createdAt;
/// Create a copy of WeightLog
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$WeightLogCopyWith<WeightLog> get copyWith => _$WeightLogCopyWithImpl<WeightLog>(this as WeightLog, _$identity);

  /// Serializes this WeightLog to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is WeightLog&&(identical(other.id, id) || other.id == id)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.note, note) || other.note == note)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.trainerId, trainerId) || other.trainerId == trainerId)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,weight,note,userId,trainerId,measuredAt,createdAt);

@override
String toString() {
  return 'WeightLog(id: $id, weight: $weight, note: $note, userId: $userId, trainerId: $trainerId, measuredAt: $measuredAt, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class $WeightLogCopyWith<$Res>  {
  factory $WeightLogCopyWith(WeightLog value, $Res Function(WeightLog) _then) = _$WeightLogCopyWithImpl;
@useResult
$Res call({
 String id, double weight, String? note,@JsonKey(name: 'userId') String userId,@JsonKey(name: 'trainerId') String? trainerId,@JsonKey(name: 'measuredAt') DateTime measuredAt,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class _$WeightLogCopyWithImpl<$Res>
    implements $WeightLogCopyWith<$Res> {
  _$WeightLogCopyWithImpl(this._self, this._then);

  final WeightLog _self;
  final $Res Function(WeightLog) _then;

/// Create a copy of WeightLog
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? weight = null,Object? note = freezed,Object? userId = null,Object? trainerId = freezed,Object? measuredAt = null,Object? createdAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,trainerId: freezed == trainerId ? _self.trainerId : trainerId // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [WeightLog].
extension WeightLogPatterns on WeightLog {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _WeightLog value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _WeightLog() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _WeightLog value)  $default,){
final _that = this;
switch (_that) {
case _WeightLog():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _WeightLog value)?  $default,){
final _that = this;
switch (_that) {
case _WeightLog() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id,  double weight,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _WeightLog() when $default != null:
return $default(_that.id,_that.weight,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id,  double weight,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)  $default,) {final _that = this;
switch (_that) {
case _WeightLog():
return $default(_that.id,_that.weight,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id,  double weight,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,) {final _that = this;
switch (_that) {
case _WeightLog() when $default != null:
return $default(_that.id,_that.weight,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _WeightLog implements WeightLog {
  const _WeightLog({required this.id, required this.weight, this.note, @JsonKey(name: 'userId') required this.userId, @JsonKey(name: 'trainerId') this.trainerId, @JsonKey(name: 'measuredAt') required this.measuredAt, @JsonKey(name: 'createdAt') required this.createdAt});
  factory _WeightLog.fromJson(Map<String, dynamic> json) => _$WeightLogFromJson(json);

@override final  String id;
@override final  double weight;
@override final  String? note;
@override@JsonKey(name: 'userId') final  String userId;
@override@JsonKey(name: 'trainerId') final  String? trainerId;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;

/// Create a copy of WeightLog
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$WeightLogCopyWith<_WeightLog> get copyWith => __$WeightLogCopyWithImpl<_WeightLog>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$WeightLogToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _WeightLog&&(identical(other.id, id) || other.id == id)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.note, note) || other.note == note)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.trainerId, trainerId) || other.trainerId == trainerId)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,weight,note,userId,trainerId,measuredAt,createdAt);

@override
String toString() {
  return 'WeightLog(id: $id, weight: $weight, note: $note, userId: $userId, trainerId: $trainerId, measuredAt: $measuredAt, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class _$WeightLogCopyWith<$Res> implements $WeightLogCopyWith<$Res> {
  factory _$WeightLogCopyWith(_WeightLog value, $Res Function(_WeightLog) _then) = __$WeightLogCopyWithImpl;
@override @useResult
$Res call({
 String id, double weight, String? note,@JsonKey(name: 'userId') String userId,@JsonKey(name: 'trainerId') String? trainerId,@JsonKey(name: 'measuredAt') DateTime measuredAt,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class __$WeightLogCopyWithImpl<$Res>
    implements _$WeightLogCopyWith<$Res> {
  __$WeightLogCopyWithImpl(this._self, this._then);

  final _WeightLog _self;
  final $Res Function(_WeightLog) _then;

/// Create a copy of WeightLog
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? weight = null,Object? note = freezed,Object? userId = null,Object? trainerId = freezed,Object? measuredAt = null,Object? createdAt = null,}) {
  return _then(_WeightLog(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,trainerId: freezed == trainerId ? _self.trainerId : trainerId // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$CreateWeightLogRequest {

 double get weight;@JsonKey(name: 'measuredAt') DateTime get measuredAt;
/// Create a copy of CreateWeightLogRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateWeightLogRequestCopyWith<CreateWeightLogRequest> get copyWith => _$CreateWeightLogRequestCopyWithImpl<CreateWeightLogRequest>(this as CreateWeightLogRequest, _$identity);

  /// Serializes this CreateWeightLogRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateWeightLogRequest&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,weight,measuredAt);

@override
String toString() {
  return 'CreateWeightLogRequest(weight: $weight, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class $CreateWeightLogRequestCopyWith<$Res>  {
  factory $CreateWeightLogRequestCopyWith(CreateWeightLogRequest value, $Res Function(CreateWeightLogRequest) _then) = _$CreateWeightLogRequestCopyWithImpl;
@useResult
$Res call({
 double weight,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class _$CreateWeightLogRequestCopyWithImpl<$Res>
    implements $CreateWeightLogRequestCopyWith<$Res> {
  _$CreateWeightLogRequestCopyWithImpl(this._self, this._then);

  final CreateWeightLogRequest _self;
  final $Res Function(CreateWeightLogRequest) _then;

/// Create a copy of CreateWeightLogRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? weight = null,Object? measuredAt = null,}) {
  return _then(_self.copyWith(
weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateWeightLogRequest].
extension CreateWeightLogRequestPatterns on CreateWeightLogRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateWeightLogRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateWeightLogRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateWeightLogRequest value)  $default,){
final _that = this;
switch (_that) {
case _CreateWeightLogRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateWeightLogRequest value)?  $default,){
final _that = this;
switch (_that) {
case _CreateWeightLogRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( double weight, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateWeightLogRequest() when $default != null:
return $default(_that.weight,_that.measuredAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( double weight, @JsonKey(name: 'measuredAt')  DateTime measuredAt)  $default,) {final _that = this;
switch (_that) {
case _CreateWeightLogRequest():
return $default(_that.weight,_that.measuredAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( double weight, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,) {final _that = this;
switch (_that) {
case _CreateWeightLogRequest() when $default != null:
return $default(_that.weight,_that.measuredAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateWeightLogRequest implements CreateWeightLogRequest {
  const _CreateWeightLogRequest({required this.weight, @JsonKey(name: 'measuredAt') required this.measuredAt});
  factory _CreateWeightLogRequest.fromJson(Map<String, dynamic> json) => _$CreateWeightLogRequestFromJson(json);

@override final  double weight;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;

/// Create a copy of CreateWeightLogRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateWeightLogRequestCopyWith<_CreateWeightLogRequest> get copyWith => __$CreateWeightLogRequestCopyWithImpl<_CreateWeightLogRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateWeightLogRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateWeightLogRequest&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,weight,measuredAt);

@override
String toString() {
  return 'CreateWeightLogRequest(weight: $weight, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class _$CreateWeightLogRequestCopyWith<$Res> implements $CreateWeightLogRequestCopyWith<$Res> {
  factory _$CreateWeightLogRequestCopyWith(_CreateWeightLogRequest value, $Res Function(_CreateWeightLogRequest) _then) = __$CreateWeightLogRequestCopyWithImpl;
@override @useResult
$Res call({
 double weight,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class __$CreateWeightLogRequestCopyWithImpl<$Res>
    implements _$CreateWeightLogRequestCopyWith<$Res> {
  __$CreateWeightLogRequestCopyWithImpl(this._self, this._then);

  final _CreateWeightLogRequest _self;
  final $Res Function(_CreateWeightLogRequest) _then;

/// Create a copy of CreateWeightLogRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? weight = null,Object? measuredAt = null,}) {
  return _then(_CreateWeightLogRequest(
weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}

// dart format on

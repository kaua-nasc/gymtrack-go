// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'body_measurement.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$BodyMeasurement {

 String get id;@BodyMeasurementTypeConverter() BodyMeasurementType get type; double get value;@JsonKey(name: 'unit') String? get unit; String? get note;@JsonKey(name: 'userId') String get userId;@JsonKey(name: 'trainerId') String? get trainerId;@JsonKey(name: 'measuredAt') DateTime get measuredAt;@JsonKey(name: 'createdAt') DateTime get createdAt;
/// Create a copy of BodyMeasurement
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$BodyMeasurementCopyWith<BodyMeasurement> get copyWith => _$BodyMeasurementCopyWithImpl<BodyMeasurement>(this as BodyMeasurement, _$identity);

  /// Serializes this BodyMeasurement to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is BodyMeasurement&&(identical(other.id, id) || other.id == id)&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.unit, unit) || other.unit == unit)&&(identical(other.note, note) || other.note == note)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.trainerId, trainerId) || other.trainerId == trainerId)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,type,value,unit,note,userId,trainerId,measuredAt,createdAt);

@override
String toString() {
  return 'BodyMeasurement(id: $id, type: $type, value: $value, unit: $unit, note: $note, userId: $userId, trainerId: $trainerId, measuredAt: $measuredAt, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class $BodyMeasurementCopyWith<$Res>  {
  factory $BodyMeasurementCopyWith(BodyMeasurement value, $Res Function(BodyMeasurement) _then) = _$BodyMeasurementCopyWithImpl;
@useResult
$Res call({
 String id,@BodyMeasurementTypeConverter() BodyMeasurementType type, double value,@JsonKey(name: 'unit') String? unit, String? note,@JsonKey(name: 'userId') String userId,@JsonKey(name: 'trainerId') String? trainerId,@JsonKey(name: 'measuredAt') DateTime measuredAt,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class _$BodyMeasurementCopyWithImpl<$Res>
    implements $BodyMeasurementCopyWith<$Res> {
  _$BodyMeasurementCopyWithImpl(this._self, this._then);

  final BodyMeasurement _self;
  final $Res Function(BodyMeasurement) _then;

/// Create a copy of BodyMeasurement
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? type = null,Object? value = null,Object? unit = freezed,Object? note = freezed,Object? userId = null,Object? trainerId = freezed,Object? measuredAt = null,Object? createdAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as BodyMeasurementType,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,unit: freezed == unit ? _self.unit : unit // ignore: cast_nullable_to_non_nullable
as String?,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,trainerId: freezed == trainerId ? _self.trainerId : trainerId // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [BodyMeasurement].
extension BodyMeasurementPatterns on BodyMeasurement {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _BodyMeasurement value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _BodyMeasurement() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _BodyMeasurement value)  $default,){
final _that = this;
switch (_that) {
case _BodyMeasurement():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _BodyMeasurement value)?  $default,){
final _that = this;
switch (_that) {
case _BodyMeasurement() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'unit')  String? unit,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _BodyMeasurement() when $default != null:
return $default(_that.id,_that.type,_that.value,_that.unit,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'unit')  String? unit,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)  $default,) {final _that = this;
switch (_that) {
case _BodyMeasurement():
return $default(_that.id,_that.type,_that.value,_that.unit,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'unit')  String? unit,  String? note, @JsonKey(name: 'userId')  String userId, @JsonKey(name: 'trainerId')  String? trainerId, @JsonKey(name: 'measuredAt')  DateTime measuredAt, @JsonKey(name: 'createdAt')  DateTime createdAt)?  $default,) {final _that = this;
switch (_that) {
case _BodyMeasurement() when $default != null:
return $default(_that.id,_that.type,_that.value,_that.unit,_that.note,_that.userId,_that.trainerId,_that.measuredAt,_that.createdAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _BodyMeasurement implements BodyMeasurement {
  const _BodyMeasurement({required this.id, @BodyMeasurementTypeConverter() required this.type, required this.value, @JsonKey(name: 'unit') this.unit, this.note, @JsonKey(name: 'userId') required this.userId, @JsonKey(name: 'trainerId') this.trainerId, @JsonKey(name: 'measuredAt') required this.measuredAt, @JsonKey(name: 'createdAt') required this.createdAt});
  factory _BodyMeasurement.fromJson(Map<String, dynamic> json) => _$BodyMeasurementFromJson(json);

@override final  String id;
@override@BodyMeasurementTypeConverter() final  BodyMeasurementType type;
@override final  double value;
@override@JsonKey(name: 'unit') final  String? unit;
@override final  String? note;
@override@JsonKey(name: 'userId') final  String userId;
@override@JsonKey(name: 'trainerId') final  String? trainerId;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;

/// Create a copy of BodyMeasurement
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$BodyMeasurementCopyWith<_BodyMeasurement> get copyWith => __$BodyMeasurementCopyWithImpl<_BodyMeasurement>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$BodyMeasurementToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _BodyMeasurement&&(identical(other.id, id) || other.id == id)&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.unit, unit) || other.unit == unit)&&(identical(other.note, note) || other.note == note)&&(identical(other.userId, userId) || other.userId == userId)&&(identical(other.trainerId, trainerId) || other.trainerId == trainerId)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,type,value,unit,note,userId,trainerId,measuredAt,createdAt);

@override
String toString() {
  return 'BodyMeasurement(id: $id, type: $type, value: $value, unit: $unit, note: $note, userId: $userId, trainerId: $trainerId, measuredAt: $measuredAt, createdAt: $createdAt)';
}


}

/// @nodoc
abstract mixin class _$BodyMeasurementCopyWith<$Res> implements $BodyMeasurementCopyWith<$Res> {
  factory _$BodyMeasurementCopyWith(_BodyMeasurement value, $Res Function(_BodyMeasurement) _then) = __$BodyMeasurementCopyWithImpl;
@override @useResult
$Res call({
 String id,@BodyMeasurementTypeConverter() BodyMeasurementType type, double value,@JsonKey(name: 'unit') String? unit, String? note,@JsonKey(name: 'userId') String userId,@JsonKey(name: 'trainerId') String? trainerId,@JsonKey(name: 'measuredAt') DateTime measuredAt,@JsonKey(name: 'createdAt') DateTime createdAt
});




}
/// @nodoc
class __$BodyMeasurementCopyWithImpl<$Res>
    implements _$BodyMeasurementCopyWith<$Res> {
  __$BodyMeasurementCopyWithImpl(this._self, this._then);

  final _BodyMeasurement _self;
  final $Res Function(_BodyMeasurement) _then;

/// Create a copy of BodyMeasurement
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? type = null,Object? value = null,Object? unit = freezed,Object? note = freezed,Object? userId = null,Object? trainerId = freezed,Object? measuredAt = null,Object? createdAt = null,}) {
  return _then(_BodyMeasurement(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as BodyMeasurementType,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,unit: freezed == unit ? _self.unit : unit // ignore: cast_nullable_to_non_nullable
as String?,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,userId: null == userId ? _self.userId : userId // ignore: cast_nullable_to_non_nullable
as String,trainerId: freezed == trainerId ? _self.trainerId : trainerId // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$CreateBodyMeasurementRequest {

@BodyMeasurementTypeConverter() BodyMeasurementType get type; double get value;@JsonKey(name: 'measuredAt') DateTime get measuredAt;
/// Create a copy of CreateBodyMeasurementRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$CreateBodyMeasurementRequestCopyWith<CreateBodyMeasurementRequest> get copyWith => _$CreateBodyMeasurementRequestCopyWithImpl<CreateBodyMeasurementRequest>(this as CreateBodyMeasurementRequest, _$identity);

  /// Serializes this CreateBodyMeasurementRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is CreateBodyMeasurementRequest&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type,value,measuredAt);

@override
String toString() {
  return 'CreateBodyMeasurementRequest(type: $type, value: $value, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class $CreateBodyMeasurementRequestCopyWith<$Res>  {
  factory $CreateBodyMeasurementRequestCopyWith(CreateBodyMeasurementRequest value, $Res Function(CreateBodyMeasurementRequest) _then) = _$CreateBodyMeasurementRequestCopyWithImpl;
@useResult
$Res call({
@BodyMeasurementTypeConverter() BodyMeasurementType type, double value,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class _$CreateBodyMeasurementRequestCopyWithImpl<$Res>
    implements $CreateBodyMeasurementRequestCopyWith<$Res> {
  _$CreateBodyMeasurementRequestCopyWithImpl(this._self, this._then);

  final CreateBodyMeasurementRequest _self;
  final $Res Function(CreateBodyMeasurementRequest) _then;

/// Create a copy of CreateBodyMeasurementRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? type = null,Object? value = null,Object? measuredAt = null,}) {
  return _then(_self.copyWith(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as BodyMeasurementType,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [CreateBodyMeasurementRequest].
extension CreateBodyMeasurementRequestPatterns on CreateBodyMeasurementRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _CreateBodyMeasurementRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _CreateBodyMeasurementRequest value)  $default,){
final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _CreateBodyMeasurementRequest value)?  $default,){
final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest() when $default != null:
return $default(_that.type,_that.value,_that.measuredAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'measuredAt')  DateTime measuredAt)  $default,) {final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest():
return $default(_that.type,_that.value,_that.measuredAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@BodyMeasurementTypeConverter()  BodyMeasurementType type,  double value, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,) {final _that = this;
switch (_that) {
case _CreateBodyMeasurementRequest() when $default != null:
return $default(_that.type,_that.value,_that.measuredAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _CreateBodyMeasurementRequest implements CreateBodyMeasurementRequest {
  const _CreateBodyMeasurementRequest({@BodyMeasurementTypeConverter() required this.type, required this.value, @JsonKey(name: 'measuredAt') required this.measuredAt});
  factory _CreateBodyMeasurementRequest.fromJson(Map<String, dynamic> json) => _$CreateBodyMeasurementRequestFromJson(json);

@override@BodyMeasurementTypeConverter() final  BodyMeasurementType type;
@override final  double value;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;

/// Create a copy of CreateBodyMeasurementRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$CreateBodyMeasurementRequestCopyWith<_CreateBodyMeasurementRequest> get copyWith => __$CreateBodyMeasurementRequestCopyWithImpl<_CreateBodyMeasurementRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$CreateBodyMeasurementRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _CreateBodyMeasurementRequest&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,type,value,measuredAt);

@override
String toString() {
  return 'CreateBodyMeasurementRequest(type: $type, value: $value, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class _$CreateBodyMeasurementRequestCopyWith<$Res> implements $CreateBodyMeasurementRequestCopyWith<$Res> {
  factory _$CreateBodyMeasurementRequestCopyWith(_CreateBodyMeasurementRequest value, $Res Function(_CreateBodyMeasurementRequest) _then) = __$CreateBodyMeasurementRequestCopyWithImpl;
@override @useResult
$Res call({
@BodyMeasurementTypeConverter() BodyMeasurementType type, double value,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class __$CreateBodyMeasurementRequestCopyWithImpl<$Res>
    implements _$CreateBodyMeasurementRequestCopyWith<$Res> {
  __$CreateBodyMeasurementRequestCopyWithImpl(this._self, this._then);

  final _CreateBodyMeasurementRequest _self;
  final $Res Function(_CreateBodyMeasurementRequest) _then;

/// Create a copy of CreateBodyMeasurementRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? type = null,Object? value = null,Object? measuredAt = null,}) {
  return _then(_CreateBodyMeasurementRequest(
type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as BodyMeasurementType,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}

// dart format on

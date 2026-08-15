// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'exercise_log.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$ExerciseLogRequest {

 List<int> get reps; List<double> get weight;@JsonKey(includeIfNull: true) String? get notes;
/// Create a copy of ExerciseLogRequest
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ExerciseLogRequestCopyWith<ExerciseLogRequest> get copyWith => _$ExerciseLogRequestCopyWithImpl<ExerciseLogRequest>(this as ExerciseLogRequest, _$identity);

  /// Serializes this ExerciseLogRequest to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is ExerciseLogRequest&&const DeepCollectionEquality().equals(other.reps, reps)&&const DeepCollectionEquality().equals(other.weight, weight)&&(identical(other.notes, notes) || other.notes == notes));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(reps),const DeepCollectionEquality().hash(weight),notes);

@override
String toString() {
  return 'ExerciseLogRequest(reps: $reps, weight: $weight, notes: $notes)';
}


}

/// @nodoc
abstract mixin class $ExerciseLogRequestCopyWith<$Res>  {
  factory $ExerciseLogRequestCopyWith(ExerciseLogRequest value, $Res Function(ExerciseLogRequest) _then) = _$ExerciseLogRequestCopyWithImpl;
@useResult
$Res call({
 List<int> reps, List<double> weight,@JsonKey(includeIfNull: true) String? notes
});




}
/// @nodoc
class _$ExerciseLogRequestCopyWithImpl<$Res>
    implements $ExerciseLogRequestCopyWith<$Res> {
  _$ExerciseLogRequestCopyWithImpl(this._self, this._then);

  final ExerciseLogRequest _self;
  final $Res Function(ExerciseLogRequest) _then;

/// Create a copy of ExerciseLogRequest
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? reps = null,Object? weight = null,Object? notes = freezed,}) {
  return _then(_self.copyWith(
reps: null == reps ? _self.reps : reps // ignore: cast_nullable_to_non_nullable
as List<int>,weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as List<double>,notes: freezed == notes ? _self.notes : notes // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}

}


/// Adds pattern-matching-related methods to [ExerciseLogRequest].
extension ExerciseLogRequestPatterns on ExerciseLogRequest {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _ExerciseLogRequest value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _ExerciseLogRequest() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _ExerciseLogRequest value)  $default,){
final _that = this;
switch (_that) {
case _ExerciseLogRequest():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _ExerciseLogRequest value)?  $default,){
final _that = this;
switch (_that) {
case _ExerciseLogRequest() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<int> reps,  List<double> weight, @JsonKey(includeIfNull: true)  String? notes)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _ExerciseLogRequest() when $default != null:
return $default(_that.reps,_that.weight,_that.notes);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<int> reps,  List<double> weight, @JsonKey(includeIfNull: true)  String? notes)  $default,) {final _that = this;
switch (_that) {
case _ExerciseLogRequest():
return $default(_that.reps,_that.weight,_that.notes);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<int> reps,  List<double> weight, @JsonKey(includeIfNull: true)  String? notes)?  $default,) {final _that = this;
switch (_that) {
case _ExerciseLogRequest() when $default != null:
return $default(_that.reps,_that.weight,_that.notes);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _ExerciseLogRequest implements ExerciseLogRequest {
  const _ExerciseLogRequest({final  List<int> reps = const [], final  List<double> weight = const [], @JsonKey(includeIfNull: true) this.notes}): _reps = reps,_weight = weight;
  factory _ExerciseLogRequest.fromJson(Map<String, dynamic> json) => _$ExerciseLogRequestFromJson(json);

 final  List<int> _reps;
@override@JsonKey() List<int> get reps {
  if (_reps is EqualUnmodifiableListView) return _reps;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_reps);
}

 final  List<double> _weight;
@override@JsonKey() List<double> get weight {
  if (_weight is EqualUnmodifiableListView) return _weight;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_weight);
}

@override@JsonKey(includeIfNull: true) final  String? notes;

/// Create a copy of ExerciseLogRequest
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ExerciseLogRequestCopyWith<_ExerciseLogRequest> get copyWith => __$ExerciseLogRequestCopyWithImpl<_ExerciseLogRequest>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ExerciseLogRequestToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _ExerciseLogRequest&&const DeepCollectionEquality().equals(other._reps, _reps)&&const DeepCollectionEquality().equals(other._weight, _weight)&&(identical(other.notes, notes) || other.notes == notes));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_reps),const DeepCollectionEquality().hash(_weight),notes);

@override
String toString() {
  return 'ExerciseLogRequest(reps: $reps, weight: $weight, notes: $notes)';
}


}

/// @nodoc
abstract mixin class _$ExerciseLogRequestCopyWith<$Res> implements $ExerciseLogRequestCopyWith<$Res> {
  factory _$ExerciseLogRequestCopyWith(_ExerciseLogRequest value, $Res Function(_ExerciseLogRequest) _then) = __$ExerciseLogRequestCopyWithImpl;
@override @useResult
$Res call({
 List<int> reps, List<double> weight,@JsonKey(includeIfNull: true) String? notes
});




}
/// @nodoc
class __$ExerciseLogRequestCopyWithImpl<$Res>
    implements _$ExerciseLogRequestCopyWith<$Res> {
  __$ExerciseLogRequestCopyWithImpl(this._self, this._then);

  final _ExerciseLogRequest _self;
  final $Res Function(_ExerciseLogRequest) _then;

/// Create a copy of ExerciseLogRequest
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? reps = null,Object? weight = null,Object? notes = freezed,}) {
  return _then(_ExerciseLogRequest(
reps: null == reps ? _self._reps : reps // ignore: cast_nullable_to_non_nullable
as List<int>,weight: null == weight ? _self._weight : weight // ignore: cast_nullable_to_non_nullable
as List<double>,notes: freezed == notes ? _self.notes : notes // ignore: cast_nullable_to_non_nullable
as String?,
  ));
}


}

// dart format on

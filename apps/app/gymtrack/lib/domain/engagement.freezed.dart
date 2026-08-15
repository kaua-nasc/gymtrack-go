// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'engagement.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$EngagementSummary {

@JsonKey(name: 'totalWorkouts') int get totalWorkouts;@JsonKey(name: 'currentStreak') int get currentStreak;@JsonKey(name: 'longestStreak') int get longestStreak;@JsonKey(name: 'weeklyCompletions') int get weeklyCompletions;@JsonKey(name: 'monthlyCompletions') int get monthlyCompletions;@JsonKey(name: 'lastWorkoutDate') DateTime? get lastWorkoutDate;
/// Create a copy of EngagementSummary
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$EngagementSummaryCopyWith<EngagementSummary> get copyWith => _$EngagementSummaryCopyWithImpl<EngagementSummary>(this as EngagementSummary, _$identity);

  /// Serializes this EngagementSummary to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is EngagementSummary&&(identical(other.totalWorkouts, totalWorkouts) || other.totalWorkouts == totalWorkouts)&&(identical(other.currentStreak, currentStreak) || other.currentStreak == currentStreak)&&(identical(other.longestStreak, longestStreak) || other.longestStreak == longestStreak)&&(identical(other.weeklyCompletions, weeklyCompletions) || other.weeklyCompletions == weeklyCompletions)&&(identical(other.monthlyCompletions, monthlyCompletions) || other.monthlyCompletions == monthlyCompletions)&&(identical(other.lastWorkoutDate, lastWorkoutDate) || other.lastWorkoutDate == lastWorkoutDate));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,totalWorkouts,currentStreak,longestStreak,weeklyCompletions,monthlyCompletions,lastWorkoutDate);

@override
String toString() {
  return 'EngagementSummary(totalWorkouts: $totalWorkouts, currentStreak: $currentStreak, longestStreak: $longestStreak, weeklyCompletions: $weeklyCompletions, monthlyCompletions: $monthlyCompletions, lastWorkoutDate: $lastWorkoutDate)';
}


}

/// @nodoc
abstract mixin class $EngagementSummaryCopyWith<$Res>  {
  factory $EngagementSummaryCopyWith(EngagementSummary value, $Res Function(EngagementSummary) _then) = _$EngagementSummaryCopyWithImpl;
@useResult
$Res call({
@JsonKey(name: 'totalWorkouts') int totalWorkouts,@JsonKey(name: 'currentStreak') int currentStreak,@JsonKey(name: 'longestStreak') int longestStreak,@JsonKey(name: 'weeklyCompletions') int weeklyCompletions,@JsonKey(name: 'monthlyCompletions') int monthlyCompletions,@JsonKey(name: 'lastWorkoutDate') DateTime? lastWorkoutDate
});




}
/// @nodoc
class _$EngagementSummaryCopyWithImpl<$Res>
    implements $EngagementSummaryCopyWith<$Res> {
  _$EngagementSummaryCopyWithImpl(this._self, this._then);

  final EngagementSummary _self;
  final $Res Function(EngagementSummary) _then;

/// Create a copy of EngagementSummary
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? totalWorkouts = null,Object? currentStreak = null,Object? longestStreak = null,Object? weeklyCompletions = null,Object? monthlyCompletions = null,Object? lastWorkoutDate = freezed,}) {
  return _then(_self.copyWith(
totalWorkouts: null == totalWorkouts ? _self.totalWorkouts : totalWorkouts // ignore: cast_nullable_to_non_nullable
as int,currentStreak: null == currentStreak ? _self.currentStreak : currentStreak // ignore: cast_nullable_to_non_nullable
as int,longestStreak: null == longestStreak ? _self.longestStreak : longestStreak // ignore: cast_nullable_to_non_nullable
as int,weeklyCompletions: null == weeklyCompletions ? _self.weeklyCompletions : weeklyCompletions // ignore: cast_nullable_to_non_nullable
as int,monthlyCompletions: null == monthlyCompletions ? _self.monthlyCompletions : monthlyCompletions // ignore: cast_nullable_to_non_nullable
as int,lastWorkoutDate: freezed == lastWorkoutDate ? _self.lastWorkoutDate : lastWorkoutDate // ignore: cast_nullable_to_non_nullable
as DateTime?,
  ));
}

}


/// Adds pattern-matching-related methods to [EngagementSummary].
extension EngagementSummaryPatterns on EngagementSummary {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _EngagementSummary value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _EngagementSummary() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _EngagementSummary value)  $default,){
final _that = this;
switch (_that) {
case _EngagementSummary():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _EngagementSummary value)?  $default,){
final _that = this;
switch (_that) {
case _EngagementSummary() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@JsonKey(name: 'totalWorkouts')  int totalWorkouts, @JsonKey(name: 'currentStreak')  int currentStreak, @JsonKey(name: 'longestStreak')  int longestStreak, @JsonKey(name: 'weeklyCompletions')  int weeklyCompletions, @JsonKey(name: 'monthlyCompletions')  int monthlyCompletions, @JsonKey(name: 'lastWorkoutDate')  DateTime? lastWorkoutDate)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _EngagementSummary() when $default != null:
return $default(_that.totalWorkouts,_that.currentStreak,_that.longestStreak,_that.weeklyCompletions,_that.monthlyCompletions,_that.lastWorkoutDate);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@JsonKey(name: 'totalWorkouts')  int totalWorkouts, @JsonKey(name: 'currentStreak')  int currentStreak, @JsonKey(name: 'longestStreak')  int longestStreak, @JsonKey(name: 'weeklyCompletions')  int weeklyCompletions, @JsonKey(name: 'monthlyCompletions')  int monthlyCompletions, @JsonKey(name: 'lastWorkoutDate')  DateTime? lastWorkoutDate)  $default,) {final _that = this;
switch (_that) {
case _EngagementSummary():
return $default(_that.totalWorkouts,_that.currentStreak,_that.longestStreak,_that.weeklyCompletions,_that.monthlyCompletions,_that.lastWorkoutDate);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@JsonKey(name: 'totalWorkouts')  int totalWorkouts, @JsonKey(name: 'currentStreak')  int currentStreak, @JsonKey(name: 'longestStreak')  int longestStreak, @JsonKey(name: 'weeklyCompletions')  int weeklyCompletions, @JsonKey(name: 'monthlyCompletions')  int monthlyCompletions, @JsonKey(name: 'lastWorkoutDate')  DateTime? lastWorkoutDate)?  $default,) {final _that = this;
switch (_that) {
case _EngagementSummary() when $default != null:
return $default(_that.totalWorkouts,_that.currentStreak,_that.longestStreak,_that.weeklyCompletions,_that.monthlyCompletions,_that.lastWorkoutDate);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _EngagementSummary implements EngagementSummary {
  const _EngagementSummary({@JsonKey(name: 'totalWorkouts') this.totalWorkouts = 0, @JsonKey(name: 'currentStreak') this.currentStreak = 0, @JsonKey(name: 'longestStreak') this.longestStreak = 0, @JsonKey(name: 'weeklyCompletions') this.weeklyCompletions = 0, @JsonKey(name: 'monthlyCompletions') this.monthlyCompletions = 0, @JsonKey(name: 'lastWorkoutDate') this.lastWorkoutDate});
  factory _EngagementSummary.fromJson(Map<String, dynamic> json) => _$EngagementSummaryFromJson(json);

@override@JsonKey(name: 'totalWorkouts') final  int totalWorkouts;
@override@JsonKey(name: 'currentStreak') final  int currentStreak;
@override@JsonKey(name: 'longestStreak') final  int longestStreak;
@override@JsonKey(name: 'weeklyCompletions') final  int weeklyCompletions;
@override@JsonKey(name: 'monthlyCompletions') final  int monthlyCompletions;
@override@JsonKey(name: 'lastWorkoutDate') final  DateTime? lastWorkoutDate;

/// Create a copy of EngagementSummary
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$EngagementSummaryCopyWith<_EngagementSummary> get copyWith => __$EngagementSummaryCopyWithImpl<_EngagementSummary>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$EngagementSummaryToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _EngagementSummary&&(identical(other.totalWorkouts, totalWorkouts) || other.totalWorkouts == totalWorkouts)&&(identical(other.currentStreak, currentStreak) || other.currentStreak == currentStreak)&&(identical(other.longestStreak, longestStreak) || other.longestStreak == longestStreak)&&(identical(other.weeklyCompletions, weeklyCompletions) || other.weeklyCompletions == weeklyCompletions)&&(identical(other.monthlyCompletions, monthlyCompletions) || other.monthlyCompletions == monthlyCompletions)&&(identical(other.lastWorkoutDate, lastWorkoutDate) || other.lastWorkoutDate == lastWorkoutDate));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,totalWorkouts,currentStreak,longestStreak,weeklyCompletions,monthlyCompletions,lastWorkoutDate);

@override
String toString() {
  return 'EngagementSummary(totalWorkouts: $totalWorkouts, currentStreak: $currentStreak, longestStreak: $longestStreak, weeklyCompletions: $weeklyCompletions, monthlyCompletions: $monthlyCompletions, lastWorkoutDate: $lastWorkoutDate)';
}


}

/// @nodoc
abstract mixin class _$EngagementSummaryCopyWith<$Res> implements $EngagementSummaryCopyWith<$Res> {
  factory _$EngagementSummaryCopyWith(_EngagementSummary value, $Res Function(_EngagementSummary) _then) = __$EngagementSummaryCopyWithImpl;
@override @useResult
$Res call({
@JsonKey(name: 'totalWorkouts') int totalWorkouts,@JsonKey(name: 'currentStreak') int currentStreak,@JsonKey(name: 'longestStreak') int longestStreak,@JsonKey(name: 'weeklyCompletions') int weeklyCompletions,@JsonKey(name: 'monthlyCompletions') int monthlyCompletions,@JsonKey(name: 'lastWorkoutDate') DateTime? lastWorkoutDate
});




}
/// @nodoc
class __$EngagementSummaryCopyWithImpl<$Res>
    implements _$EngagementSummaryCopyWith<$Res> {
  __$EngagementSummaryCopyWithImpl(this._self, this._then);

  final _EngagementSummary _self;
  final $Res Function(_EngagementSummary) _then;

/// Create a copy of EngagementSummary
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? totalWorkouts = null,Object? currentStreak = null,Object? longestStreak = null,Object? weeklyCompletions = null,Object? monthlyCompletions = null,Object? lastWorkoutDate = freezed,}) {
  return _then(_EngagementSummary(
totalWorkouts: null == totalWorkouts ? _self.totalWorkouts : totalWorkouts // ignore: cast_nullable_to_non_nullable
as int,currentStreak: null == currentStreak ? _self.currentStreak : currentStreak // ignore: cast_nullable_to_non_nullable
as int,longestStreak: null == longestStreak ? _self.longestStreak : longestStreak // ignore: cast_nullable_to_non_nullable
as int,weeklyCompletions: null == weeklyCompletions ? _self.weeklyCompletions : weeklyCompletions // ignore: cast_nullable_to_non_nullable
as int,monthlyCompletions: null == monthlyCompletions ? _self.monthlyCompletions : monthlyCompletions // ignore: cast_nullable_to_non_nullable
as int,lastWorkoutDate: freezed == lastWorkoutDate ? _self.lastWorkoutDate : lastWorkoutDate // ignore: cast_nullable_to_non_nullable
as DateTime?,
  ));
}


}


/// @nodoc
mixin _$BiometricDashboard {

 List<BodyMeasurementEntry> get bodyMeasurements; List<WeightEntry> get weightLogs;
/// Create a copy of BiometricDashboard
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$BiometricDashboardCopyWith<BiometricDashboard> get copyWith => _$BiometricDashboardCopyWithImpl<BiometricDashboard>(this as BiometricDashboard, _$identity);

  /// Serializes this BiometricDashboard to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is BiometricDashboard&&const DeepCollectionEquality().equals(other.bodyMeasurements, bodyMeasurements)&&const DeepCollectionEquality().equals(other.weightLogs, weightLogs));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(bodyMeasurements),const DeepCollectionEquality().hash(weightLogs));

@override
String toString() {
  return 'BiometricDashboard(bodyMeasurements: $bodyMeasurements, weightLogs: $weightLogs)';
}


}

/// @nodoc
abstract mixin class $BiometricDashboardCopyWith<$Res>  {
  factory $BiometricDashboardCopyWith(BiometricDashboard value, $Res Function(BiometricDashboard) _then) = _$BiometricDashboardCopyWithImpl;
@useResult
$Res call({
 List<BodyMeasurementEntry> bodyMeasurements, List<WeightEntry> weightLogs
});




}
/// @nodoc
class _$BiometricDashboardCopyWithImpl<$Res>
    implements $BiometricDashboardCopyWith<$Res> {
  _$BiometricDashboardCopyWithImpl(this._self, this._then);

  final BiometricDashboard _self;
  final $Res Function(BiometricDashboard) _then;

/// Create a copy of BiometricDashboard
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? bodyMeasurements = null,Object? weightLogs = null,}) {
  return _then(_self.copyWith(
bodyMeasurements: null == bodyMeasurements ? _self.bodyMeasurements : bodyMeasurements // ignore: cast_nullable_to_non_nullable
as List<BodyMeasurementEntry>,weightLogs: null == weightLogs ? _self.weightLogs : weightLogs // ignore: cast_nullable_to_non_nullable
as List<WeightEntry>,
  ));
}

}


/// Adds pattern-matching-related methods to [BiometricDashboard].
extension BiometricDashboardPatterns on BiometricDashboard {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _BiometricDashboard value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _BiometricDashboard() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _BiometricDashboard value)  $default,){
final _that = this;
switch (_that) {
case _BiometricDashboard():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _BiometricDashboard value)?  $default,){
final _that = this;
switch (_that) {
case _BiometricDashboard() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( List<BodyMeasurementEntry> bodyMeasurements,  List<WeightEntry> weightLogs)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _BiometricDashboard() when $default != null:
return $default(_that.bodyMeasurements,_that.weightLogs);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( List<BodyMeasurementEntry> bodyMeasurements,  List<WeightEntry> weightLogs)  $default,) {final _that = this;
switch (_that) {
case _BiometricDashboard():
return $default(_that.bodyMeasurements,_that.weightLogs);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( List<BodyMeasurementEntry> bodyMeasurements,  List<WeightEntry> weightLogs)?  $default,) {final _that = this;
switch (_that) {
case _BiometricDashboard() when $default != null:
return $default(_that.bodyMeasurements,_that.weightLogs);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _BiometricDashboard implements BiometricDashboard {
  const _BiometricDashboard({final  List<BodyMeasurementEntry> bodyMeasurements = const [], final  List<WeightEntry> weightLogs = const []}): _bodyMeasurements = bodyMeasurements,_weightLogs = weightLogs;
  factory _BiometricDashboard.fromJson(Map<String, dynamic> json) => _$BiometricDashboardFromJson(json);

 final  List<BodyMeasurementEntry> _bodyMeasurements;
@override@JsonKey() List<BodyMeasurementEntry> get bodyMeasurements {
  if (_bodyMeasurements is EqualUnmodifiableListView) return _bodyMeasurements;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_bodyMeasurements);
}

 final  List<WeightEntry> _weightLogs;
@override@JsonKey() List<WeightEntry> get weightLogs {
  if (_weightLogs is EqualUnmodifiableListView) return _weightLogs;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_weightLogs);
}


/// Create a copy of BiometricDashboard
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$BiometricDashboardCopyWith<_BiometricDashboard> get copyWith => __$BiometricDashboardCopyWithImpl<_BiometricDashboard>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$BiometricDashboardToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _BiometricDashboard&&const DeepCollectionEquality().equals(other._bodyMeasurements, _bodyMeasurements)&&const DeepCollectionEquality().equals(other._weightLogs, _weightLogs));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,const DeepCollectionEquality().hash(_bodyMeasurements),const DeepCollectionEquality().hash(_weightLogs));

@override
String toString() {
  return 'BiometricDashboard(bodyMeasurements: $bodyMeasurements, weightLogs: $weightLogs)';
}


}

/// @nodoc
abstract mixin class _$BiometricDashboardCopyWith<$Res> implements $BiometricDashboardCopyWith<$Res> {
  factory _$BiometricDashboardCopyWith(_BiometricDashboard value, $Res Function(_BiometricDashboard) _then) = __$BiometricDashboardCopyWithImpl;
@override @useResult
$Res call({
 List<BodyMeasurementEntry> bodyMeasurements, List<WeightEntry> weightLogs
});




}
/// @nodoc
class __$BiometricDashboardCopyWithImpl<$Res>
    implements _$BiometricDashboardCopyWith<$Res> {
  __$BiometricDashboardCopyWithImpl(this._self, this._then);

  final _BiometricDashboard _self;
  final $Res Function(_BiometricDashboard) _then;

/// Create a copy of BiometricDashboard
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? bodyMeasurements = null,Object? weightLogs = null,}) {
  return _then(_BiometricDashboard(
bodyMeasurements: null == bodyMeasurements ? _self._bodyMeasurements : bodyMeasurements // ignore: cast_nullable_to_non_nullable
as List<BodyMeasurementEntry>,weightLogs: null == weightLogs ? _self._weightLogs : weightLogs // ignore: cast_nullable_to_non_nullable
as List<WeightEntry>,
  ));
}


}


/// @nodoc
mixin _$BodyMeasurementEntry {

 String get id; String get type; double get value; String? get note;@JsonKey(name: 'measuredAt') DateTime get measuredAt;
/// Create a copy of BodyMeasurementEntry
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$BodyMeasurementEntryCopyWith<BodyMeasurementEntry> get copyWith => _$BodyMeasurementEntryCopyWithImpl<BodyMeasurementEntry>(this as BodyMeasurementEntry, _$identity);

  /// Serializes this BodyMeasurementEntry to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is BodyMeasurementEntry&&(identical(other.id, id) || other.id == id)&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.note, note) || other.note == note)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,type,value,note,measuredAt);

@override
String toString() {
  return 'BodyMeasurementEntry(id: $id, type: $type, value: $value, note: $note, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class $BodyMeasurementEntryCopyWith<$Res>  {
  factory $BodyMeasurementEntryCopyWith(BodyMeasurementEntry value, $Res Function(BodyMeasurementEntry) _then) = _$BodyMeasurementEntryCopyWithImpl;
@useResult
$Res call({
 String id, String type, double value, String? note,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class _$BodyMeasurementEntryCopyWithImpl<$Res>
    implements $BodyMeasurementEntryCopyWith<$Res> {
  _$BodyMeasurementEntryCopyWithImpl(this._self, this._then);

  final BodyMeasurementEntry _self;
  final $Res Function(BodyMeasurementEntry) _then;

/// Create a copy of BodyMeasurementEntry
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? type = null,Object? value = null,Object? note = freezed,Object? measuredAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as String,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [BodyMeasurementEntry].
extension BodyMeasurementEntryPatterns on BodyMeasurementEntry {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _BodyMeasurementEntry value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _BodyMeasurementEntry() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _BodyMeasurementEntry value)  $default,){
final _that = this;
switch (_that) {
case _BodyMeasurementEntry():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _BodyMeasurementEntry value)?  $default,){
final _that = this;
switch (_that) {
case _BodyMeasurementEntry() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id,  String type,  double value,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _BodyMeasurementEntry() when $default != null:
return $default(_that.id,_that.type,_that.value,_that.note,_that.measuredAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id,  String type,  double value,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)  $default,) {final _that = this;
switch (_that) {
case _BodyMeasurementEntry():
return $default(_that.id,_that.type,_that.value,_that.note,_that.measuredAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id,  String type,  double value,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,) {final _that = this;
switch (_that) {
case _BodyMeasurementEntry() when $default != null:
return $default(_that.id,_that.type,_that.value,_that.note,_that.measuredAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _BodyMeasurementEntry implements BodyMeasurementEntry {
  const _BodyMeasurementEntry({required this.id, required this.type, required this.value, this.note, @JsonKey(name: 'measuredAt') required this.measuredAt});
  factory _BodyMeasurementEntry.fromJson(Map<String, dynamic> json) => _$BodyMeasurementEntryFromJson(json);

@override final  String id;
@override final  String type;
@override final  double value;
@override final  String? note;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;

/// Create a copy of BodyMeasurementEntry
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$BodyMeasurementEntryCopyWith<_BodyMeasurementEntry> get copyWith => __$BodyMeasurementEntryCopyWithImpl<_BodyMeasurementEntry>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$BodyMeasurementEntryToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _BodyMeasurementEntry&&(identical(other.id, id) || other.id == id)&&(identical(other.type, type) || other.type == type)&&(identical(other.value, value) || other.value == value)&&(identical(other.note, note) || other.note == note)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,type,value,note,measuredAt);

@override
String toString() {
  return 'BodyMeasurementEntry(id: $id, type: $type, value: $value, note: $note, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class _$BodyMeasurementEntryCopyWith<$Res> implements $BodyMeasurementEntryCopyWith<$Res> {
  factory _$BodyMeasurementEntryCopyWith(_BodyMeasurementEntry value, $Res Function(_BodyMeasurementEntry) _then) = __$BodyMeasurementEntryCopyWithImpl;
@override @useResult
$Res call({
 String id, String type, double value, String? note,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class __$BodyMeasurementEntryCopyWithImpl<$Res>
    implements _$BodyMeasurementEntryCopyWith<$Res> {
  __$BodyMeasurementEntryCopyWithImpl(this._self, this._then);

  final _BodyMeasurementEntry _self;
  final $Res Function(_BodyMeasurementEntry) _then;

/// Create a copy of BodyMeasurementEntry
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? type = null,Object? value = null,Object? note = freezed,Object? measuredAt = null,}) {
  return _then(_BodyMeasurementEntry(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as String,value: null == value ? _self.value : value // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$WeightEntry {

 String get id; double get weight; String? get note;@JsonKey(name: 'measuredAt') DateTime get measuredAt;
/// Create a copy of WeightEntry
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$WeightEntryCopyWith<WeightEntry> get copyWith => _$WeightEntryCopyWithImpl<WeightEntry>(this as WeightEntry, _$identity);

  /// Serializes this WeightEntry to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is WeightEntry&&(identical(other.id, id) || other.id == id)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.note, note) || other.note == note)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,weight,note,measuredAt);

@override
String toString() {
  return 'WeightEntry(id: $id, weight: $weight, note: $note, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class $WeightEntryCopyWith<$Res>  {
  factory $WeightEntryCopyWith(WeightEntry value, $Res Function(WeightEntry) _then) = _$WeightEntryCopyWithImpl;
@useResult
$Res call({
 String id, double weight, String? note,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class _$WeightEntryCopyWithImpl<$Res>
    implements $WeightEntryCopyWith<$Res> {
  _$WeightEntryCopyWithImpl(this._self, this._then);

  final WeightEntry _self;
  final $Res Function(WeightEntry) _then;

/// Create a copy of WeightEntry
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? weight = null,Object? note = freezed,Object? measuredAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [WeightEntry].
extension WeightEntryPatterns on WeightEntry {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _WeightEntry value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _WeightEntry() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _WeightEntry value)  $default,){
final _that = this;
switch (_that) {
case _WeightEntry():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _WeightEntry value)?  $default,){
final _that = this;
switch (_that) {
case _WeightEntry() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id,  double weight,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _WeightEntry() when $default != null:
return $default(_that.id,_that.weight,_that.note,_that.measuredAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id,  double weight,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)  $default,) {final _that = this;
switch (_that) {
case _WeightEntry():
return $default(_that.id,_that.weight,_that.note,_that.measuredAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id,  double weight,  String? note, @JsonKey(name: 'measuredAt')  DateTime measuredAt)?  $default,) {final _that = this;
switch (_that) {
case _WeightEntry() when $default != null:
return $default(_that.id,_that.weight,_that.note,_that.measuredAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _WeightEntry implements WeightEntry {
  const _WeightEntry({required this.id, required this.weight, this.note, @JsonKey(name: 'measuredAt') required this.measuredAt});
  factory _WeightEntry.fromJson(Map<String, dynamic> json) => _$WeightEntryFromJson(json);

@override final  String id;
@override final  double weight;
@override final  String? note;
@override@JsonKey(name: 'measuredAt') final  DateTime measuredAt;

/// Create a copy of WeightEntry
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$WeightEntryCopyWith<_WeightEntry> get copyWith => __$WeightEntryCopyWithImpl<_WeightEntry>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$WeightEntryToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _WeightEntry&&(identical(other.id, id) || other.id == id)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.note, note) || other.note == note)&&(identical(other.measuredAt, measuredAt) || other.measuredAt == measuredAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,weight,note,measuredAt);

@override
String toString() {
  return 'WeightEntry(id: $id, weight: $weight, note: $note, measuredAt: $measuredAt)';
}


}

/// @nodoc
abstract mixin class _$WeightEntryCopyWith<$Res> implements $WeightEntryCopyWith<$Res> {
  factory _$WeightEntryCopyWith(_WeightEntry value, $Res Function(_WeightEntry) _then) = __$WeightEntryCopyWithImpl;
@override @useResult
$Res call({
 String id, double weight, String? note,@JsonKey(name: 'measuredAt') DateTime measuredAt
});




}
/// @nodoc
class __$WeightEntryCopyWithImpl<$Res>
    implements _$WeightEntryCopyWith<$Res> {
  __$WeightEntryCopyWithImpl(this._self, this._then);

  final _WeightEntry _self;
  final $Res Function(_WeightEntry) _then;

/// Create a copy of WeightEntry
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? weight = null,Object? note = freezed,Object? measuredAt = null,}) {
  return _then(_WeightEntry(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,weight: null == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double,note: freezed == note ? _self.note : note // ignore: cast_nullable_to_non_nullable
as String?,measuredAt: null == measuredAt ? _self.measuredAt : measuredAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$InsightsDashboard {

@JsonKey(name: 'avgCompletionRate') double? get avgCompletionRate; List<MonthInsight> get monthly; List<ExerciseInsight> get topExercises;
/// Create a copy of InsightsDashboard
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$InsightsDashboardCopyWith<InsightsDashboard> get copyWith => _$InsightsDashboardCopyWithImpl<InsightsDashboard>(this as InsightsDashboard, _$identity);

  /// Serializes this InsightsDashboard to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is InsightsDashboard&&(identical(other.avgCompletionRate, avgCompletionRate) || other.avgCompletionRate == avgCompletionRate)&&const DeepCollectionEquality().equals(other.monthly, monthly)&&const DeepCollectionEquality().equals(other.topExercises, topExercises));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,avgCompletionRate,const DeepCollectionEquality().hash(monthly),const DeepCollectionEquality().hash(topExercises));

@override
String toString() {
  return 'InsightsDashboard(avgCompletionRate: $avgCompletionRate, monthly: $monthly, topExercises: $topExercises)';
}


}

/// @nodoc
abstract mixin class $InsightsDashboardCopyWith<$Res>  {
  factory $InsightsDashboardCopyWith(InsightsDashboard value, $Res Function(InsightsDashboard) _then) = _$InsightsDashboardCopyWithImpl;
@useResult
$Res call({
@JsonKey(name: 'avgCompletionRate') double? avgCompletionRate, List<MonthInsight> monthly, List<ExerciseInsight> topExercises
});




}
/// @nodoc
class _$InsightsDashboardCopyWithImpl<$Res>
    implements $InsightsDashboardCopyWith<$Res> {
  _$InsightsDashboardCopyWithImpl(this._self, this._then);

  final InsightsDashboard _self;
  final $Res Function(InsightsDashboard) _then;

/// Create a copy of InsightsDashboard
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? avgCompletionRate = freezed,Object? monthly = null,Object? topExercises = null,}) {
  return _then(_self.copyWith(
avgCompletionRate: freezed == avgCompletionRate ? _self.avgCompletionRate : avgCompletionRate // ignore: cast_nullable_to_non_nullable
as double?,monthly: null == monthly ? _self.monthly : monthly // ignore: cast_nullable_to_non_nullable
as List<MonthInsight>,topExercises: null == topExercises ? _self.topExercises : topExercises // ignore: cast_nullable_to_non_nullable
as List<ExerciseInsight>,
  ));
}

}


/// Adds pattern-matching-related methods to [InsightsDashboard].
extension InsightsDashboardPatterns on InsightsDashboard {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _InsightsDashboard value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _InsightsDashboard() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _InsightsDashboard value)  $default,){
final _that = this;
switch (_that) {
case _InsightsDashboard():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _InsightsDashboard value)?  $default,){
final _that = this;
switch (_that) {
case _InsightsDashboard() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@JsonKey(name: 'avgCompletionRate')  double? avgCompletionRate,  List<MonthInsight> monthly,  List<ExerciseInsight> topExercises)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _InsightsDashboard() when $default != null:
return $default(_that.avgCompletionRate,_that.monthly,_that.topExercises);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@JsonKey(name: 'avgCompletionRate')  double? avgCompletionRate,  List<MonthInsight> monthly,  List<ExerciseInsight> topExercises)  $default,) {final _that = this;
switch (_that) {
case _InsightsDashboard():
return $default(_that.avgCompletionRate,_that.monthly,_that.topExercises);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@JsonKey(name: 'avgCompletionRate')  double? avgCompletionRate,  List<MonthInsight> monthly,  List<ExerciseInsight> topExercises)?  $default,) {final _that = this;
switch (_that) {
case _InsightsDashboard() when $default != null:
return $default(_that.avgCompletionRate,_that.monthly,_that.topExercises);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _InsightsDashboard implements InsightsDashboard {
  const _InsightsDashboard({@JsonKey(name: 'avgCompletionRate') this.avgCompletionRate, final  List<MonthInsight> monthly = const [], final  List<ExerciseInsight> topExercises = const []}): _monthly = monthly,_topExercises = topExercises;
  factory _InsightsDashboard.fromJson(Map<String, dynamic> json) => _$InsightsDashboardFromJson(json);

@override@JsonKey(name: 'avgCompletionRate') final  double? avgCompletionRate;
 final  List<MonthInsight> _monthly;
@override@JsonKey() List<MonthInsight> get monthly {
  if (_monthly is EqualUnmodifiableListView) return _monthly;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_monthly);
}

 final  List<ExerciseInsight> _topExercises;
@override@JsonKey() List<ExerciseInsight> get topExercises {
  if (_topExercises is EqualUnmodifiableListView) return _topExercises;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_topExercises);
}


/// Create a copy of InsightsDashboard
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$InsightsDashboardCopyWith<_InsightsDashboard> get copyWith => __$InsightsDashboardCopyWithImpl<_InsightsDashboard>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$InsightsDashboardToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _InsightsDashboard&&(identical(other.avgCompletionRate, avgCompletionRate) || other.avgCompletionRate == avgCompletionRate)&&const DeepCollectionEquality().equals(other._monthly, _monthly)&&const DeepCollectionEquality().equals(other._topExercises, _topExercises));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,avgCompletionRate,const DeepCollectionEquality().hash(_monthly),const DeepCollectionEquality().hash(_topExercises));

@override
String toString() {
  return 'InsightsDashboard(avgCompletionRate: $avgCompletionRate, monthly: $monthly, topExercises: $topExercises)';
}


}

/// @nodoc
abstract mixin class _$InsightsDashboardCopyWith<$Res> implements $InsightsDashboardCopyWith<$Res> {
  factory _$InsightsDashboardCopyWith(_InsightsDashboard value, $Res Function(_InsightsDashboard) _then) = __$InsightsDashboardCopyWithImpl;
@override @useResult
$Res call({
@JsonKey(name: 'avgCompletionRate') double? avgCompletionRate, List<MonthInsight> monthly, List<ExerciseInsight> topExercises
});




}
/// @nodoc
class __$InsightsDashboardCopyWithImpl<$Res>
    implements _$InsightsDashboardCopyWith<$Res> {
  __$InsightsDashboardCopyWithImpl(this._self, this._then);

  final _InsightsDashboard _self;
  final $Res Function(_InsightsDashboard) _then;

/// Create a copy of InsightsDashboard
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? avgCompletionRate = freezed,Object? monthly = null,Object? topExercises = null,}) {
  return _then(_InsightsDashboard(
avgCompletionRate: freezed == avgCompletionRate ? _self.avgCompletionRate : avgCompletionRate // ignore: cast_nullable_to_non_nullable
as double?,monthly: null == monthly ? _self._monthly : monthly // ignore: cast_nullable_to_non_nullable
as List<MonthInsight>,topExercises: null == topExercises ? _self._topExercises : topExercises // ignore: cast_nullable_to_non_nullable
as List<ExerciseInsight>,
  ));
}


}


/// @nodoc
mixin _$MonthInsight {

 String get month;@JsonKey(name: 'completedDays') int get completedDays;@JsonKey(name: 'totalDays') int get totalDays;
/// Create a copy of MonthInsight
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$MonthInsightCopyWith<MonthInsight> get copyWith => _$MonthInsightCopyWithImpl<MonthInsight>(this as MonthInsight, _$identity);

  /// Serializes this MonthInsight to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is MonthInsight&&(identical(other.month, month) || other.month == month)&&(identical(other.completedDays, completedDays) || other.completedDays == completedDays)&&(identical(other.totalDays, totalDays) || other.totalDays == totalDays));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,month,completedDays,totalDays);

@override
String toString() {
  return 'MonthInsight(month: $month, completedDays: $completedDays, totalDays: $totalDays)';
}


}

/// @nodoc
abstract mixin class $MonthInsightCopyWith<$Res>  {
  factory $MonthInsightCopyWith(MonthInsight value, $Res Function(MonthInsight) _then) = _$MonthInsightCopyWithImpl;
@useResult
$Res call({
 String month,@JsonKey(name: 'completedDays') int completedDays,@JsonKey(name: 'totalDays') int totalDays
});




}
/// @nodoc
class _$MonthInsightCopyWithImpl<$Res>
    implements $MonthInsightCopyWith<$Res> {
  _$MonthInsightCopyWithImpl(this._self, this._then);

  final MonthInsight _self;
  final $Res Function(MonthInsight) _then;

/// Create a copy of MonthInsight
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? month = null,Object? completedDays = null,Object? totalDays = null,}) {
  return _then(_self.copyWith(
month: null == month ? _self.month : month // ignore: cast_nullable_to_non_nullable
as String,completedDays: null == completedDays ? _self.completedDays : completedDays // ignore: cast_nullable_to_non_nullable
as int,totalDays: null == totalDays ? _self.totalDays : totalDays // ignore: cast_nullable_to_non_nullable
as int,
  ));
}

}


/// Adds pattern-matching-related methods to [MonthInsight].
extension MonthInsightPatterns on MonthInsight {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _MonthInsight value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _MonthInsight() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _MonthInsight value)  $default,){
final _that = this;
switch (_that) {
case _MonthInsight():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _MonthInsight value)?  $default,){
final _that = this;
switch (_that) {
case _MonthInsight() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String month, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _MonthInsight() when $default != null:
return $default(_that.month,_that.completedDays,_that.totalDays);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String month, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)  $default,) {final _that = this;
switch (_that) {
case _MonthInsight():
return $default(_that.month,_that.completedDays,_that.totalDays);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String month, @JsonKey(name: 'completedDays')  int completedDays, @JsonKey(name: 'totalDays')  int totalDays)?  $default,) {final _that = this;
switch (_that) {
case _MonthInsight() when $default != null:
return $default(_that.month,_that.completedDays,_that.totalDays);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _MonthInsight implements MonthInsight {
  const _MonthInsight({required this.month, @JsonKey(name: 'completedDays') this.completedDays = 0, @JsonKey(name: 'totalDays') this.totalDays = 0});
  factory _MonthInsight.fromJson(Map<String, dynamic> json) => _$MonthInsightFromJson(json);

@override final  String month;
@override@JsonKey(name: 'completedDays') final  int completedDays;
@override@JsonKey(name: 'totalDays') final  int totalDays;

/// Create a copy of MonthInsight
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$MonthInsightCopyWith<_MonthInsight> get copyWith => __$MonthInsightCopyWithImpl<_MonthInsight>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$MonthInsightToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _MonthInsight&&(identical(other.month, month) || other.month == month)&&(identical(other.completedDays, completedDays) || other.completedDays == completedDays)&&(identical(other.totalDays, totalDays) || other.totalDays == totalDays));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,month,completedDays,totalDays);

@override
String toString() {
  return 'MonthInsight(month: $month, completedDays: $completedDays, totalDays: $totalDays)';
}


}

/// @nodoc
abstract mixin class _$MonthInsightCopyWith<$Res> implements $MonthInsightCopyWith<$Res> {
  factory _$MonthInsightCopyWith(_MonthInsight value, $Res Function(_MonthInsight) _then) = __$MonthInsightCopyWithImpl;
@override @useResult
$Res call({
 String month,@JsonKey(name: 'completedDays') int completedDays,@JsonKey(name: 'totalDays') int totalDays
});




}
/// @nodoc
class __$MonthInsightCopyWithImpl<$Res>
    implements _$MonthInsightCopyWith<$Res> {
  __$MonthInsightCopyWithImpl(this._self, this._then);

  final _MonthInsight _self;
  final $Res Function(_MonthInsight) _then;

/// Create a copy of MonthInsight
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? month = null,Object? completedDays = null,Object? totalDays = null,}) {
  return _then(_MonthInsight(
month: null == month ? _self.month : month // ignore: cast_nullable_to_non_nullable
as String,completedDays: null == completedDays ? _self.completedDays : completedDays // ignore: cast_nullable_to_non_nullable
as int,totalDays: null == totalDays ? _self.totalDays : totalDays // ignore: cast_nullable_to_non_nullable
as int,
  ));
}


}


/// @nodoc
mixin _$ExerciseInsight {

 String get name;@JsonKey(name: 'totalSets') int get totalSets;@JsonKey(name: 'totalReps') int get totalReps;@JsonKey(name: 'maxWeight') double? get maxWeight;
/// Create a copy of ExerciseInsight
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ExerciseInsightCopyWith<ExerciseInsight> get copyWith => _$ExerciseInsightCopyWithImpl<ExerciseInsight>(this as ExerciseInsight, _$identity);

  /// Serializes this ExerciseInsight to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is ExerciseInsight&&(identical(other.name, name) || other.name == name)&&(identical(other.totalSets, totalSets) || other.totalSets == totalSets)&&(identical(other.totalReps, totalReps) || other.totalReps == totalReps)&&(identical(other.maxWeight, maxWeight) || other.maxWeight == maxWeight));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,name,totalSets,totalReps,maxWeight);

@override
String toString() {
  return 'ExerciseInsight(name: $name, totalSets: $totalSets, totalReps: $totalReps, maxWeight: $maxWeight)';
}


}

/// @nodoc
abstract mixin class $ExerciseInsightCopyWith<$Res>  {
  factory $ExerciseInsightCopyWith(ExerciseInsight value, $Res Function(ExerciseInsight) _then) = _$ExerciseInsightCopyWithImpl;
@useResult
$Res call({
 String name,@JsonKey(name: 'totalSets') int totalSets,@JsonKey(name: 'totalReps') int totalReps,@JsonKey(name: 'maxWeight') double? maxWeight
});




}
/// @nodoc
class _$ExerciseInsightCopyWithImpl<$Res>
    implements $ExerciseInsightCopyWith<$Res> {
  _$ExerciseInsightCopyWithImpl(this._self, this._then);

  final ExerciseInsight _self;
  final $Res Function(ExerciseInsight) _then;

/// Create a copy of ExerciseInsight
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? name = null,Object? totalSets = null,Object? totalReps = null,Object? maxWeight = freezed,}) {
  return _then(_self.copyWith(
name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,totalSets: null == totalSets ? _self.totalSets : totalSets // ignore: cast_nullable_to_non_nullable
as int,totalReps: null == totalReps ? _self.totalReps : totalReps // ignore: cast_nullable_to_non_nullable
as int,maxWeight: freezed == maxWeight ? _self.maxWeight : maxWeight // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}

}


/// Adds pattern-matching-related methods to [ExerciseInsight].
extension ExerciseInsightPatterns on ExerciseInsight {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _ExerciseInsight value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _ExerciseInsight() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _ExerciseInsight value)  $default,){
final _that = this;
switch (_that) {
case _ExerciseInsight():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _ExerciseInsight value)?  $default,){
final _that = this;
switch (_that) {
case _ExerciseInsight() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String name, @JsonKey(name: 'totalSets')  int totalSets, @JsonKey(name: 'totalReps')  int totalReps, @JsonKey(name: 'maxWeight')  double? maxWeight)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _ExerciseInsight() when $default != null:
return $default(_that.name,_that.totalSets,_that.totalReps,_that.maxWeight);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String name, @JsonKey(name: 'totalSets')  int totalSets, @JsonKey(name: 'totalReps')  int totalReps, @JsonKey(name: 'maxWeight')  double? maxWeight)  $default,) {final _that = this;
switch (_that) {
case _ExerciseInsight():
return $default(_that.name,_that.totalSets,_that.totalReps,_that.maxWeight);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String name, @JsonKey(name: 'totalSets')  int totalSets, @JsonKey(name: 'totalReps')  int totalReps, @JsonKey(name: 'maxWeight')  double? maxWeight)?  $default,) {final _that = this;
switch (_that) {
case _ExerciseInsight() when $default != null:
return $default(_that.name,_that.totalSets,_that.totalReps,_that.maxWeight);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _ExerciseInsight implements ExerciseInsight {
  const _ExerciseInsight({required this.name, @JsonKey(name: 'totalSets') this.totalSets = 0, @JsonKey(name: 'totalReps') this.totalReps = 0, @JsonKey(name: 'maxWeight') this.maxWeight});
  factory _ExerciseInsight.fromJson(Map<String, dynamic> json) => _$ExerciseInsightFromJson(json);

@override final  String name;
@override@JsonKey(name: 'totalSets') final  int totalSets;
@override@JsonKey(name: 'totalReps') final  int totalReps;
@override@JsonKey(name: 'maxWeight') final  double? maxWeight;

/// Create a copy of ExerciseInsight
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ExerciseInsightCopyWith<_ExerciseInsight> get copyWith => __$ExerciseInsightCopyWithImpl<_ExerciseInsight>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ExerciseInsightToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _ExerciseInsight&&(identical(other.name, name) || other.name == name)&&(identical(other.totalSets, totalSets) || other.totalSets == totalSets)&&(identical(other.totalReps, totalReps) || other.totalReps == totalReps)&&(identical(other.maxWeight, maxWeight) || other.maxWeight == maxWeight));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,name,totalSets,totalReps,maxWeight);

@override
String toString() {
  return 'ExerciseInsight(name: $name, totalSets: $totalSets, totalReps: $totalReps, maxWeight: $maxWeight)';
}


}

/// @nodoc
abstract mixin class _$ExerciseInsightCopyWith<$Res> implements $ExerciseInsightCopyWith<$Res> {
  factory _$ExerciseInsightCopyWith(_ExerciseInsight value, $Res Function(_ExerciseInsight) _then) = __$ExerciseInsightCopyWithImpl;
@override @useResult
$Res call({
 String name,@JsonKey(name: 'totalSets') int totalSets,@JsonKey(name: 'totalReps') int totalReps,@JsonKey(name: 'maxWeight') double? maxWeight
});




}
/// @nodoc
class __$ExerciseInsightCopyWithImpl<$Res>
    implements _$ExerciseInsightCopyWith<$Res> {
  __$ExerciseInsightCopyWithImpl(this._self, this._then);

  final _ExerciseInsight _self;
  final $Res Function(_ExerciseInsight) _then;

/// Create a copy of ExerciseInsight
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? name = null,Object? totalSets = null,Object? totalReps = null,Object? maxWeight = freezed,}) {
  return _then(_ExerciseInsight(
name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,totalSets: null == totalSets ? _self.totalSets : totalSets // ignore: cast_nullable_to_non_nullable
as int,totalReps: null == totalReps ? _self.totalReps : totalReps // ignore: cast_nullable_to_non_nullable
as int,maxWeight: freezed == maxWeight ? _self.maxWeight : maxWeight // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}


}

// dart format on

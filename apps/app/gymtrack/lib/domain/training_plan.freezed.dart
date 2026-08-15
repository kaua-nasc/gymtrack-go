// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'training_plan.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$TrainingPlan {

 String get id;@JsonKey(name: 'authorId') String get authorId; String get title; String? get description;@TrainingPlanTypeConverter() TrainingPlanType get type;@TrainingPlanLevelConverter() TrainingPlanLevel get level;@TrainingPlanVisibilityConverter() TrainingPlanVisibility get visibility;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt; List<Day> get days;@JsonKey(name: 'subscriptionCount') int get subscriptionCount;@JsonKey(name: 'averageRating') double? get averageRating;
/// Create a copy of TrainingPlan
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$TrainingPlanCopyWith<TrainingPlan> get copyWith => _$TrainingPlanCopyWithImpl<TrainingPlan>(this as TrainingPlan, _$identity);

  /// Serializes this TrainingPlan to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is TrainingPlan&&(identical(other.id, id) || other.id == id)&&(identical(other.authorId, authorId) || other.authorId == authorId)&&(identical(other.title, title) || other.title == title)&&(identical(other.description, description) || other.description == description)&&(identical(other.type, type) || other.type == type)&&(identical(other.level, level) || other.level == level)&&(identical(other.visibility, visibility) || other.visibility == visibility)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt)&&const DeepCollectionEquality().equals(other.days, days)&&(identical(other.subscriptionCount, subscriptionCount) || other.subscriptionCount == subscriptionCount)&&(identical(other.averageRating, averageRating) || other.averageRating == averageRating));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,authorId,title,description,type,level,visibility,createdAt,updatedAt,const DeepCollectionEquality().hash(days),subscriptionCount,averageRating);

@override
String toString() {
  return 'TrainingPlan(id: $id, authorId: $authorId, title: $title, description: $description, type: $type, level: $level, visibility: $visibility, createdAt: $createdAt, updatedAt: $updatedAt, days: $days, subscriptionCount: $subscriptionCount, averageRating: $averageRating)';
}


}

/// @nodoc
abstract mixin class $TrainingPlanCopyWith<$Res>  {
  factory $TrainingPlanCopyWith(TrainingPlan value, $Res Function(TrainingPlan) _then) = _$TrainingPlanCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'authorId') String authorId, String title, String? description,@TrainingPlanTypeConverter() TrainingPlanType type,@TrainingPlanLevelConverter() TrainingPlanLevel level,@TrainingPlanVisibilityConverter() TrainingPlanVisibility visibility,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt, List<Day> days,@JsonKey(name: 'subscriptionCount') int subscriptionCount,@JsonKey(name: 'averageRating') double? averageRating
});




}
/// @nodoc
class _$TrainingPlanCopyWithImpl<$Res>
    implements $TrainingPlanCopyWith<$Res> {
  _$TrainingPlanCopyWithImpl(this._self, this._then);

  final TrainingPlan _self;
  final $Res Function(TrainingPlan) _then;

/// Create a copy of TrainingPlan
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? authorId = null,Object? title = null,Object? description = freezed,Object? type = null,Object? level = null,Object? visibility = null,Object? createdAt = null,Object? updatedAt = null,Object? days = null,Object? subscriptionCount = null,Object? averageRating = freezed,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,authorId: null == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String,title: null == title ? _self.title : title // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as TrainingPlanType,level: null == level ? _self.level : level // ignore: cast_nullable_to_non_nullable
as TrainingPlanLevel,visibility: null == visibility ? _self.visibility : visibility // ignore: cast_nullable_to_non_nullable
as TrainingPlanVisibility,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,days: null == days ? _self.days : days // ignore: cast_nullable_to_non_nullable
as List<Day>,subscriptionCount: null == subscriptionCount ? _self.subscriptionCount : subscriptionCount // ignore: cast_nullable_to_non_nullable
as int,averageRating: freezed == averageRating ? _self.averageRating : averageRating // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}

}


/// Adds pattern-matching-related methods to [TrainingPlan].
extension TrainingPlanPatterns on TrainingPlan {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _TrainingPlan value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _TrainingPlan() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _TrainingPlan value)  $default,){
final _that = this;
switch (_that) {
case _TrainingPlan():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _TrainingPlan value)?  $default,){
final _that = this;
switch (_that) {
case _TrainingPlan() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'authorId')  String authorId,  String title,  String? description, @TrainingPlanTypeConverter()  TrainingPlanType type, @TrainingPlanLevelConverter()  TrainingPlanLevel level, @TrainingPlanVisibilityConverter()  TrainingPlanVisibility visibility, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt,  List<Day> days, @JsonKey(name: 'subscriptionCount')  int subscriptionCount, @JsonKey(name: 'averageRating')  double? averageRating)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _TrainingPlan() when $default != null:
return $default(_that.id,_that.authorId,_that.title,_that.description,_that.type,_that.level,_that.visibility,_that.createdAt,_that.updatedAt,_that.days,_that.subscriptionCount,_that.averageRating);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'authorId')  String authorId,  String title,  String? description, @TrainingPlanTypeConverter()  TrainingPlanType type, @TrainingPlanLevelConverter()  TrainingPlanLevel level, @TrainingPlanVisibilityConverter()  TrainingPlanVisibility visibility, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt,  List<Day> days, @JsonKey(name: 'subscriptionCount')  int subscriptionCount, @JsonKey(name: 'averageRating')  double? averageRating)  $default,) {final _that = this;
switch (_that) {
case _TrainingPlan():
return $default(_that.id,_that.authorId,_that.title,_that.description,_that.type,_that.level,_that.visibility,_that.createdAt,_that.updatedAt,_that.days,_that.subscriptionCount,_that.averageRating);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'authorId')  String authorId,  String title,  String? description, @TrainingPlanTypeConverter()  TrainingPlanType type, @TrainingPlanLevelConverter()  TrainingPlanLevel level, @TrainingPlanVisibilityConverter()  TrainingPlanVisibility visibility, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt,  List<Day> days, @JsonKey(name: 'subscriptionCount')  int subscriptionCount, @JsonKey(name: 'averageRating')  double? averageRating)?  $default,) {final _that = this;
switch (_that) {
case _TrainingPlan() when $default != null:
return $default(_that.id,_that.authorId,_that.title,_that.description,_that.type,_that.level,_that.visibility,_that.createdAt,_that.updatedAt,_that.days,_that.subscriptionCount,_that.averageRating);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _TrainingPlan implements TrainingPlan {
  const _TrainingPlan({required this.id, @JsonKey(name: 'authorId') required this.authorId, required this.title, this.description, @TrainingPlanTypeConverter() required this.type, @TrainingPlanLevelConverter() required this.level, @TrainingPlanVisibilityConverter() required this.visibility, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt, final  List<Day> days = const [], @JsonKey(name: 'subscriptionCount') this.subscriptionCount = 0, @JsonKey(name: 'averageRating') this.averageRating}): _days = days;
  factory _TrainingPlan.fromJson(Map<String, dynamic> json) => _$TrainingPlanFromJson(json);

@override final  String id;
@override@JsonKey(name: 'authorId') final  String authorId;
@override final  String title;
@override final  String? description;
@override@TrainingPlanTypeConverter() final  TrainingPlanType type;
@override@TrainingPlanLevelConverter() final  TrainingPlanLevel level;
@override@TrainingPlanVisibilityConverter() final  TrainingPlanVisibility visibility;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;
 final  List<Day> _days;
@override@JsonKey() List<Day> get days {
  if (_days is EqualUnmodifiableListView) return _days;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_days);
}

@override@JsonKey(name: 'subscriptionCount') final  int subscriptionCount;
@override@JsonKey(name: 'averageRating') final  double? averageRating;

/// Create a copy of TrainingPlan
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$TrainingPlanCopyWith<_TrainingPlan> get copyWith => __$TrainingPlanCopyWithImpl<_TrainingPlan>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$TrainingPlanToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _TrainingPlan&&(identical(other.id, id) || other.id == id)&&(identical(other.authorId, authorId) || other.authorId == authorId)&&(identical(other.title, title) || other.title == title)&&(identical(other.description, description) || other.description == description)&&(identical(other.type, type) || other.type == type)&&(identical(other.level, level) || other.level == level)&&(identical(other.visibility, visibility) || other.visibility == visibility)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt)&&const DeepCollectionEquality().equals(other._days, _days)&&(identical(other.subscriptionCount, subscriptionCount) || other.subscriptionCount == subscriptionCount)&&(identical(other.averageRating, averageRating) || other.averageRating == averageRating));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,authorId,title,description,type,level,visibility,createdAt,updatedAt,const DeepCollectionEquality().hash(_days),subscriptionCount,averageRating);

@override
String toString() {
  return 'TrainingPlan(id: $id, authorId: $authorId, title: $title, description: $description, type: $type, level: $level, visibility: $visibility, createdAt: $createdAt, updatedAt: $updatedAt, days: $days, subscriptionCount: $subscriptionCount, averageRating: $averageRating)';
}


}

/// @nodoc
abstract mixin class _$TrainingPlanCopyWith<$Res> implements $TrainingPlanCopyWith<$Res> {
  factory _$TrainingPlanCopyWith(_TrainingPlan value, $Res Function(_TrainingPlan) _then) = __$TrainingPlanCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'authorId') String authorId, String title, String? description,@TrainingPlanTypeConverter() TrainingPlanType type,@TrainingPlanLevelConverter() TrainingPlanLevel level,@TrainingPlanVisibilityConverter() TrainingPlanVisibility visibility,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt, List<Day> days,@JsonKey(name: 'subscriptionCount') int subscriptionCount,@JsonKey(name: 'averageRating') double? averageRating
});




}
/// @nodoc
class __$TrainingPlanCopyWithImpl<$Res>
    implements _$TrainingPlanCopyWith<$Res> {
  __$TrainingPlanCopyWithImpl(this._self, this._then);

  final _TrainingPlan _self;
  final $Res Function(_TrainingPlan) _then;

/// Create a copy of TrainingPlan
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? authorId = null,Object? title = null,Object? description = freezed,Object? type = null,Object? level = null,Object? visibility = null,Object? createdAt = null,Object? updatedAt = null,Object? days = null,Object? subscriptionCount = null,Object? averageRating = freezed,}) {
  return _then(_TrainingPlan(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,authorId: null == authorId ? _self.authorId : authorId // ignore: cast_nullable_to_non_nullable
as String,title: null == title ? _self.title : title // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,type: null == type ? _self.type : type // ignore: cast_nullable_to_non_nullable
as TrainingPlanType,level: null == level ? _self.level : level // ignore: cast_nullable_to_non_nullable
as TrainingPlanLevel,visibility: null == visibility ? _self.visibility : visibility // ignore: cast_nullable_to_non_nullable
as TrainingPlanVisibility,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,days: null == days ? _self._days : days // ignore: cast_nullable_to_non_nullable
as List<Day>,subscriptionCount: null == subscriptionCount ? _self.subscriptionCount : subscriptionCount // ignore: cast_nullable_to_non_nullable
as int,averageRating: freezed == averageRating ? _self.averageRating : averageRating // ignore: cast_nullable_to_non_nullable
as double?,
  ));
}


}


/// @nodoc
mixin _$Day {

 String get id;@JsonKey(name: 'planId') String get planId; String get name; String? get description;@JsonKey(name: 'orderIndex') int get orderIndex; List<Exercise> get exercises;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;
/// Create a copy of Day
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$DayCopyWith<Day> get copyWith => _$DayCopyWithImpl<Day>(this as Day, _$identity);

  /// Serializes this Day to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is Day&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.name, name) || other.name == name)&&(identical(other.description, description) || other.description == description)&&(identical(other.orderIndex, orderIndex) || other.orderIndex == orderIndex)&&const DeepCollectionEquality().equals(other.exercises, exercises)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,name,description,orderIndex,const DeepCollectionEquality().hash(exercises),createdAt,updatedAt);

@override
String toString() {
  return 'Day(id: $id, planId: $planId, name: $name, description: $description, orderIndex: $orderIndex, exercises: $exercises, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $DayCopyWith<$Res>  {
  factory $DayCopyWith(Day value, $Res Function(Day) _then) = _$DayCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId, String name, String? description,@JsonKey(name: 'orderIndex') int orderIndex, List<Exercise> exercises,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class _$DayCopyWithImpl<$Res>
    implements $DayCopyWith<$Res> {
  _$DayCopyWithImpl(this._self, this._then);

  final Day _self;
  final $Res Function(Day) _then;

/// Create a copy of Day
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? planId = null,Object? name = null,Object? description = freezed,Object? orderIndex = null,Object? exercises = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,orderIndex: null == orderIndex ? _self.orderIndex : orderIndex // ignore: cast_nullable_to_non_nullable
as int,exercises: null == exercises ? _self.exercises : exercises // ignore: cast_nullable_to_non_nullable
as List<Exercise>,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [Day].
extension DayPatterns on Day {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _Day value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _Day() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _Day value)  $default,){
final _that = this;
switch (_that) {
case _Day():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _Day value)?  $default,){
final _that = this;
switch (_that) {
case _Day() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId,  String name,  String? description, @JsonKey(name: 'orderIndex')  int orderIndex,  List<Exercise> exercises, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _Day() when $default != null:
return $default(_that.id,_that.planId,_that.name,_that.description,_that.orderIndex,_that.exercises,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'planId')  String planId,  String name,  String? description, @JsonKey(name: 'orderIndex')  int orderIndex,  List<Exercise> exercises, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _Day():
return $default(_that.id,_that.planId,_that.name,_that.description,_that.orderIndex,_that.exercises,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'planId')  String planId,  String name,  String? description, @JsonKey(name: 'orderIndex')  int orderIndex,  List<Exercise> exercises, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _Day() when $default != null:
return $default(_that.id,_that.planId,_that.name,_that.description,_that.orderIndex,_that.exercises,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _Day implements Day {
  const _Day({required this.id, @JsonKey(name: 'planId') required this.planId, required this.name, this.description, @JsonKey(name: 'orderIndex') required this.orderIndex, final  List<Exercise> exercises = const [], @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt}): _exercises = exercises;
  factory _Day.fromJson(Map<String, dynamic> json) => _$DayFromJson(json);

@override final  String id;
@override@JsonKey(name: 'planId') final  String planId;
@override final  String name;
@override final  String? description;
@override@JsonKey(name: 'orderIndex') final  int orderIndex;
 final  List<Exercise> _exercises;
@override@JsonKey() List<Exercise> get exercises {
  if (_exercises is EqualUnmodifiableListView) return _exercises;
  // ignore: implicit_dynamic_type
  return EqualUnmodifiableListView(_exercises);
}

@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;

/// Create a copy of Day
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$DayCopyWith<_Day> get copyWith => __$DayCopyWithImpl<_Day>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$DayToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _Day&&(identical(other.id, id) || other.id == id)&&(identical(other.planId, planId) || other.planId == planId)&&(identical(other.name, name) || other.name == name)&&(identical(other.description, description) || other.description == description)&&(identical(other.orderIndex, orderIndex) || other.orderIndex == orderIndex)&&const DeepCollectionEquality().equals(other._exercises, _exercises)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,planId,name,description,orderIndex,const DeepCollectionEquality().hash(_exercises),createdAt,updatedAt);

@override
String toString() {
  return 'Day(id: $id, planId: $planId, name: $name, description: $description, orderIndex: $orderIndex, exercises: $exercises, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$DayCopyWith<$Res> implements $DayCopyWith<$Res> {
  factory _$DayCopyWith(_Day value, $Res Function(_Day) _then) = __$DayCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'planId') String planId, String name, String? description,@JsonKey(name: 'orderIndex') int orderIndex, List<Exercise> exercises,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class __$DayCopyWithImpl<$Res>
    implements _$DayCopyWith<$Res> {
  __$DayCopyWithImpl(this._self, this._then);

  final _Day _self;
  final $Res Function(_Day) _then;

/// Create a copy of Day
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? planId = null,Object? name = null,Object? description = freezed,Object? orderIndex = null,Object? exercises = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_Day(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,planId: null == planId ? _self.planId : planId // ignore: cast_nullable_to_non_nullable
as String,name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,orderIndex: null == orderIndex ? _self.orderIndex : orderIndex // ignore: cast_nullable_to_non_nullable
as int,exercises: null == exercises ? _self._exercises : exercises // ignore: cast_nullable_to_non_nullable
as List<Exercise>,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}


/// @nodoc
mixin _$Exercise {

 String get id;@JsonKey(name: 'dayId') String get dayId; String get name; String? get description;@JsonKey(name: 'seriesCount') int? get seriesCount;@JsonKey(name: 'repetitionCount') int? get repetitionCount; double? get weight; String? get videoUrl; String? get imageUrl;@JsonKey(name: 'restTimeSeconds') int? get restTimeSeconds;@JsonKey(name: 'orderIndex') int get orderIndex;@JsonKey(name: 'createdAt') DateTime get createdAt;@JsonKey(name: 'updatedAt') DateTime get updatedAt;
/// Create a copy of Exercise
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$ExerciseCopyWith<Exercise> get copyWith => _$ExerciseCopyWithImpl<Exercise>(this as Exercise, _$identity);

  /// Serializes this Exercise to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is Exercise&&(identical(other.id, id) || other.id == id)&&(identical(other.dayId, dayId) || other.dayId == dayId)&&(identical(other.name, name) || other.name == name)&&(identical(other.description, description) || other.description == description)&&(identical(other.seriesCount, seriesCount) || other.seriesCount == seriesCount)&&(identical(other.repetitionCount, repetitionCount) || other.repetitionCount == repetitionCount)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.videoUrl, videoUrl) || other.videoUrl == videoUrl)&&(identical(other.imageUrl, imageUrl) || other.imageUrl == imageUrl)&&(identical(other.restTimeSeconds, restTimeSeconds) || other.restTimeSeconds == restTimeSeconds)&&(identical(other.orderIndex, orderIndex) || other.orderIndex == orderIndex)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,dayId,name,description,seriesCount,repetitionCount,weight,videoUrl,imageUrl,restTimeSeconds,orderIndex,createdAt,updatedAt);

@override
String toString() {
  return 'Exercise(id: $id, dayId: $dayId, name: $name, description: $description, seriesCount: $seriesCount, repetitionCount: $repetitionCount, weight: $weight, videoUrl: $videoUrl, imageUrl: $imageUrl, restTimeSeconds: $restTimeSeconds, orderIndex: $orderIndex, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class $ExerciseCopyWith<$Res>  {
  factory $ExerciseCopyWith(Exercise value, $Res Function(Exercise) _then) = _$ExerciseCopyWithImpl;
@useResult
$Res call({
 String id,@JsonKey(name: 'dayId') String dayId, String name, String? description,@JsonKey(name: 'seriesCount') int? seriesCount,@JsonKey(name: 'repetitionCount') int? repetitionCount, double? weight, String? videoUrl, String? imageUrl,@JsonKey(name: 'restTimeSeconds') int? restTimeSeconds,@JsonKey(name: 'orderIndex') int orderIndex,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class _$ExerciseCopyWithImpl<$Res>
    implements $ExerciseCopyWith<$Res> {
  _$ExerciseCopyWithImpl(this._self, this._then);

  final Exercise _self;
  final $Res Function(Exercise) _then;

/// Create a copy of Exercise
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? dayId = null,Object? name = null,Object? description = freezed,Object? seriesCount = freezed,Object? repetitionCount = freezed,Object? weight = freezed,Object? videoUrl = freezed,Object? imageUrl = freezed,Object? restTimeSeconds = freezed,Object? orderIndex = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_self.copyWith(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,dayId: null == dayId ? _self.dayId : dayId // ignore: cast_nullable_to_non_nullable
as String,name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,seriesCount: freezed == seriesCount ? _self.seriesCount : seriesCount // ignore: cast_nullable_to_non_nullable
as int?,repetitionCount: freezed == repetitionCount ? _self.repetitionCount : repetitionCount // ignore: cast_nullable_to_non_nullable
as int?,weight: freezed == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double?,videoUrl: freezed == videoUrl ? _self.videoUrl : videoUrl // ignore: cast_nullable_to_non_nullable
as String?,imageUrl: freezed == imageUrl ? _self.imageUrl : imageUrl // ignore: cast_nullable_to_non_nullable
as String?,restTimeSeconds: freezed == restTimeSeconds ? _self.restTimeSeconds : restTimeSeconds // ignore: cast_nullable_to_non_nullable
as int?,orderIndex: null == orderIndex ? _self.orderIndex : orderIndex // ignore: cast_nullable_to_non_nullable
as int,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}

}


/// Adds pattern-matching-related methods to [Exercise].
extension ExercisePatterns on Exercise {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _Exercise value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _Exercise() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _Exercise value)  $default,){
final _that = this;
switch (_that) {
case _Exercise():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _Exercise value)?  $default,){
final _that = this;
switch (_that) {
case _Exercise() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'dayId')  String dayId,  String name,  String? description, @JsonKey(name: 'seriesCount')  int? seriesCount, @JsonKey(name: 'repetitionCount')  int? repetitionCount,  double? weight,  String? videoUrl,  String? imageUrl, @JsonKey(name: 'restTimeSeconds')  int? restTimeSeconds, @JsonKey(name: 'orderIndex')  int orderIndex, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _Exercise() when $default != null:
return $default(_that.id,_that.dayId,_that.name,_that.description,_that.seriesCount,_that.repetitionCount,_that.weight,_that.videoUrl,_that.imageUrl,_that.restTimeSeconds,_that.orderIndex,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function( String id, @JsonKey(name: 'dayId')  String dayId,  String name,  String? description, @JsonKey(name: 'seriesCount')  int? seriesCount, @JsonKey(name: 'repetitionCount')  int? repetitionCount,  double? weight,  String? videoUrl,  String? imageUrl, @JsonKey(name: 'restTimeSeconds')  int? restTimeSeconds, @JsonKey(name: 'orderIndex')  int orderIndex, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)  $default,) {final _that = this;
switch (_that) {
case _Exercise():
return $default(_that.id,_that.dayId,_that.name,_that.description,_that.seriesCount,_that.repetitionCount,_that.weight,_that.videoUrl,_that.imageUrl,_that.restTimeSeconds,_that.orderIndex,_that.createdAt,_that.updatedAt);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function( String id, @JsonKey(name: 'dayId')  String dayId,  String name,  String? description, @JsonKey(name: 'seriesCount')  int? seriesCount, @JsonKey(name: 'repetitionCount')  int? repetitionCount,  double? weight,  String? videoUrl,  String? imageUrl, @JsonKey(name: 'restTimeSeconds')  int? restTimeSeconds, @JsonKey(name: 'orderIndex')  int orderIndex, @JsonKey(name: 'createdAt')  DateTime createdAt, @JsonKey(name: 'updatedAt')  DateTime updatedAt)?  $default,) {final _that = this;
switch (_that) {
case _Exercise() when $default != null:
return $default(_that.id,_that.dayId,_that.name,_that.description,_that.seriesCount,_that.repetitionCount,_that.weight,_that.videoUrl,_that.imageUrl,_that.restTimeSeconds,_that.orderIndex,_that.createdAt,_that.updatedAt);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _Exercise implements Exercise {
  const _Exercise({required this.id, @JsonKey(name: 'dayId') required this.dayId, required this.name, this.description, @JsonKey(name: 'seriesCount') this.seriesCount, @JsonKey(name: 'repetitionCount') this.repetitionCount, this.weight, this.videoUrl, this.imageUrl, @JsonKey(name: 'restTimeSeconds') this.restTimeSeconds, @JsonKey(name: 'orderIndex') required this.orderIndex, @JsonKey(name: 'createdAt') required this.createdAt, @JsonKey(name: 'updatedAt') required this.updatedAt});
  factory _Exercise.fromJson(Map<String, dynamic> json) => _$ExerciseFromJson(json);

@override final  String id;
@override@JsonKey(name: 'dayId') final  String dayId;
@override final  String name;
@override final  String? description;
@override@JsonKey(name: 'seriesCount') final  int? seriesCount;
@override@JsonKey(name: 'repetitionCount') final  int? repetitionCount;
@override final  double? weight;
@override final  String? videoUrl;
@override final  String? imageUrl;
@override@JsonKey(name: 'restTimeSeconds') final  int? restTimeSeconds;
@override@JsonKey(name: 'orderIndex') final  int orderIndex;
@override@JsonKey(name: 'createdAt') final  DateTime createdAt;
@override@JsonKey(name: 'updatedAt') final  DateTime updatedAt;

/// Create a copy of Exercise
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$ExerciseCopyWith<_Exercise> get copyWith => __$ExerciseCopyWithImpl<_Exercise>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$ExerciseToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _Exercise&&(identical(other.id, id) || other.id == id)&&(identical(other.dayId, dayId) || other.dayId == dayId)&&(identical(other.name, name) || other.name == name)&&(identical(other.description, description) || other.description == description)&&(identical(other.seriesCount, seriesCount) || other.seriesCount == seriesCount)&&(identical(other.repetitionCount, repetitionCount) || other.repetitionCount == repetitionCount)&&(identical(other.weight, weight) || other.weight == weight)&&(identical(other.videoUrl, videoUrl) || other.videoUrl == videoUrl)&&(identical(other.imageUrl, imageUrl) || other.imageUrl == imageUrl)&&(identical(other.restTimeSeconds, restTimeSeconds) || other.restTimeSeconds == restTimeSeconds)&&(identical(other.orderIndex, orderIndex) || other.orderIndex == orderIndex)&&(identical(other.createdAt, createdAt) || other.createdAt == createdAt)&&(identical(other.updatedAt, updatedAt) || other.updatedAt == updatedAt));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,dayId,name,description,seriesCount,repetitionCount,weight,videoUrl,imageUrl,restTimeSeconds,orderIndex,createdAt,updatedAt);

@override
String toString() {
  return 'Exercise(id: $id, dayId: $dayId, name: $name, description: $description, seriesCount: $seriesCount, repetitionCount: $repetitionCount, weight: $weight, videoUrl: $videoUrl, imageUrl: $imageUrl, restTimeSeconds: $restTimeSeconds, orderIndex: $orderIndex, createdAt: $createdAt, updatedAt: $updatedAt)';
}


}

/// @nodoc
abstract mixin class _$ExerciseCopyWith<$Res> implements $ExerciseCopyWith<$Res> {
  factory _$ExerciseCopyWith(_Exercise value, $Res Function(_Exercise) _then) = __$ExerciseCopyWithImpl;
@override @useResult
$Res call({
 String id,@JsonKey(name: 'dayId') String dayId, String name, String? description,@JsonKey(name: 'seriesCount') int? seriesCount,@JsonKey(name: 'repetitionCount') int? repetitionCount, double? weight, String? videoUrl, String? imageUrl,@JsonKey(name: 'restTimeSeconds') int? restTimeSeconds,@JsonKey(name: 'orderIndex') int orderIndex,@JsonKey(name: 'createdAt') DateTime createdAt,@JsonKey(name: 'updatedAt') DateTime updatedAt
});




}
/// @nodoc
class __$ExerciseCopyWithImpl<$Res>
    implements _$ExerciseCopyWith<$Res> {
  __$ExerciseCopyWithImpl(this._self, this._then);

  final _Exercise _self;
  final $Res Function(_Exercise) _then;

/// Create a copy of Exercise
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? id = null,Object? dayId = null,Object? name = null,Object? description = freezed,Object? seriesCount = freezed,Object? repetitionCount = freezed,Object? weight = freezed,Object? videoUrl = freezed,Object? imageUrl = freezed,Object? restTimeSeconds = freezed,Object? orderIndex = null,Object? createdAt = null,Object? updatedAt = null,}) {
  return _then(_Exercise(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,dayId: null == dayId ? _self.dayId : dayId // ignore: cast_nullable_to_non_nullable
as String,name: null == name ? _self.name : name // ignore: cast_nullable_to_non_nullable
as String,description: freezed == description ? _self.description : description // ignore: cast_nullable_to_non_nullable
as String?,seriesCount: freezed == seriesCount ? _self.seriesCount : seriesCount // ignore: cast_nullable_to_non_nullable
as int?,repetitionCount: freezed == repetitionCount ? _self.repetitionCount : repetitionCount // ignore: cast_nullable_to_non_nullable
as int?,weight: freezed == weight ? _self.weight : weight // ignore: cast_nullable_to_non_nullable
as double?,videoUrl: freezed == videoUrl ? _self.videoUrl : videoUrl // ignore: cast_nullable_to_non_nullable
as String?,imageUrl: freezed == imageUrl ? _self.imageUrl : imageUrl // ignore: cast_nullable_to_non_nullable
as String?,restTimeSeconds: freezed == restTimeSeconds ? _self.restTimeSeconds : restTimeSeconds // ignore: cast_nullable_to_non_nullable
as int?,orderIndex: null == orderIndex ? _self.orderIndex : orderIndex // ignore: cast_nullable_to_non_nullable
as int,createdAt: null == createdAt ? _self.createdAt : createdAt // ignore: cast_nullable_to_non_nullable
as DateTime,updatedAt: null == updatedAt ? _self.updatedAt : updatedAt // ignore: cast_nullable_to_non_nullable
as DateTime,
  ));
}


}

// dart format on

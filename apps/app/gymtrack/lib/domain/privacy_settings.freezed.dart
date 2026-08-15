// GENERATED CODE - DO NOT MODIFY BY HAND
// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'privacy_settings.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$UserPrivacySettings {

@JsonKey(name: 'shareEmail') bool get shareEmail;@JsonKey(name: 'shareTrainingProgress') bool get shareTrainingProgress;@JsonKey(name: 'sharePastDataWithTrainer') bool get sharePastDataWithTrainer;@JsonKey(name: 'shareBodyMeasurements') bool get shareBodyMeasurements;@JsonKey(name: 'shareWeightLogs') bool get shareWeightLogs;@JsonKey(name: 'allowTrainerNotes') bool get allowTrainerNotes;
/// Create a copy of UserPrivacySettings
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UserPrivacySettingsCopyWith<UserPrivacySettings> get copyWith => _$UserPrivacySettingsCopyWithImpl<UserPrivacySettings>(this as UserPrivacySettings, _$identity);

  /// Serializes this UserPrivacySettings to a JSON map.
  Map<String, dynamic> toJson();


@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is UserPrivacySettings&&(identical(other.shareEmail, shareEmail) || other.shareEmail == shareEmail)&&(identical(other.shareTrainingProgress, shareTrainingProgress) || other.shareTrainingProgress == shareTrainingProgress)&&(identical(other.sharePastDataWithTrainer, sharePastDataWithTrainer) || other.sharePastDataWithTrainer == sharePastDataWithTrainer)&&(identical(other.shareBodyMeasurements, shareBodyMeasurements) || other.shareBodyMeasurements == shareBodyMeasurements)&&(identical(other.shareWeightLogs, shareWeightLogs) || other.shareWeightLogs == shareWeightLogs)&&(identical(other.allowTrainerNotes, allowTrainerNotes) || other.allowTrainerNotes == allowTrainerNotes));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,shareEmail,shareTrainingProgress,sharePastDataWithTrainer,shareBodyMeasurements,shareWeightLogs,allowTrainerNotes);

@override
String toString() {
  return 'UserPrivacySettings(shareEmail: $shareEmail, shareTrainingProgress: $shareTrainingProgress, sharePastDataWithTrainer: $sharePastDataWithTrainer, shareBodyMeasurements: $shareBodyMeasurements, shareWeightLogs: $shareWeightLogs, allowTrainerNotes: $allowTrainerNotes)';
}


}

/// @nodoc
abstract mixin class $UserPrivacySettingsCopyWith<$Res>  {
  factory $UserPrivacySettingsCopyWith(UserPrivacySettings value, $Res Function(UserPrivacySettings) _then) = _$UserPrivacySettingsCopyWithImpl;
@useResult
$Res call({
@JsonKey(name: 'shareEmail') bool shareEmail,@JsonKey(name: 'shareTrainingProgress') bool shareTrainingProgress,@JsonKey(name: 'sharePastDataWithTrainer') bool sharePastDataWithTrainer,@JsonKey(name: 'shareBodyMeasurements') bool shareBodyMeasurements,@JsonKey(name: 'shareWeightLogs') bool shareWeightLogs,@JsonKey(name: 'allowTrainerNotes') bool allowTrainerNotes
});




}
/// @nodoc
class _$UserPrivacySettingsCopyWithImpl<$Res>
    implements $UserPrivacySettingsCopyWith<$Res> {
  _$UserPrivacySettingsCopyWithImpl(this._self, this._then);

  final UserPrivacySettings _self;
  final $Res Function(UserPrivacySettings) _then;

/// Create a copy of UserPrivacySettings
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? shareEmail = null,Object? shareTrainingProgress = null,Object? sharePastDataWithTrainer = null,Object? shareBodyMeasurements = null,Object? shareWeightLogs = null,Object? allowTrainerNotes = null,}) {
  return _then(_self.copyWith(
shareEmail: null == shareEmail ? _self.shareEmail : shareEmail // ignore: cast_nullable_to_non_nullable
as bool,shareTrainingProgress: null == shareTrainingProgress ? _self.shareTrainingProgress : shareTrainingProgress // ignore: cast_nullable_to_non_nullable
as bool,sharePastDataWithTrainer: null == sharePastDataWithTrainer ? _self.sharePastDataWithTrainer : sharePastDataWithTrainer // ignore: cast_nullable_to_non_nullable
as bool,shareBodyMeasurements: null == shareBodyMeasurements ? _self.shareBodyMeasurements : shareBodyMeasurements // ignore: cast_nullable_to_non_nullable
as bool,shareWeightLogs: null == shareWeightLogs ? _self.shareWeightLogs : shareWeightLogs // ignore: cast_nullable_to_non_nullable
as bool,allowTrainerNotes: null == allowTrainerNotes ? _self.allowTrainerNotes : allowTrainerNotes // ignore: cast_nullable_to_non_nullable
as bool,
  ));
}

}


/// Adds pattern-matching-related methods to [UserPrivacySettings].
extension UserPrivacySettingsPatterns on UserPrivacySettings {
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

@optionalTypeArgs TResult maybeMap<TResult extends Object?>(TResult Function( _UserPrivacySettings value)?  $default,{required TResult orElse(),}){
final _that = this;
switch (_that) {
case _UserPrivacySettings() when $default != null:
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

@optionalTypeArgs TResult map<TResult extends Object?>(TResult Function( _UserPrivacySettings value)  $default,){
final _that = this;
switch (_that) {
case _UserPrivacySettings():
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

@optionalTypeArgs TResult? mapOrNull<TResult extends Object?>(TResult? Function( _UserPrivacySettings value)?  $default,){
final _that = this;
switch (_that) {
case _UserPrivacySettings() when $default != null:
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

@optionalTypeArgs TResult maybeWhen<TResult extends Object?>(TResult Function(@JsonKey(name: 'shareEmail')  bool shareEmail, @JsonKey(name: 'shareTrainingProgress')  bool shareTrainingProgress, @JsonKey(name: 'sharePastDataWithTrainer')  bool sharePastDataWithTrainer, @JsonKey(name: 'shareBodyMeasurements')  bool shareBodyMeasurements, @JsonKey(name: 'shareWeightLogs')  bool shareWeightLogs, @JsonKey(name: 'allowTrainerNotes')  bool allowTrainerNotes)?  $default,{required TResult orElse(),}) {final _that = this;
switch (_that) {
case _UserPrivacySettings() when $default != null:
return $default(_that.shareEmail,_that.shareTrainingProgress,_that.sharePastDataWithTrainer,_that.shareBodyMeasurements,_that.shareWeightLogs,_that.allowTrainerNotes);case _:
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

@optionalTypeArgs TResult when<TResult extends Object?>(TResult Function(@JsonKey(name: 'shareEmail')  bool shareEmail, @JsonKey(name: 'shareTrainingProgress')  bool shareTrainingProgress, @JsonKey(name: 'sharePastDataWithTrainer')  bool sharePastDataWithTrainer, @JsonKey(name: 'shareBodyMeasurements')  bool shareBodyMeasurements, @JsonKey(name: 'shareWeightLogs')  bool shareWeightLogs, @JsonKey(name: 'allowTrainerNotes')  bool allowTrainerNotes)  $default,) {final _that = this;
switch (_that) {
case _UserPrivacySettings():
return $default(_that.shareEmail,_that.shareTrainingProgress,_that.sharePastDataWithTrainer,_that.shareBodyMeasurements,_that.shareWeightLogs,_that.allowTrainerNotes);case _:
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

@optionalTypeArgs TResult? whenOrNull<TResult extends Object?>(TResult? Function(@JsonKey(name: 'shareEmail')  bool shareEmail, @JsonKey(name: 'shareTrainingProgress')  bool shareTrainingProgress, @JsonKey(name: 'sharePastDataWithTrainer')  bool sharePastDataWithTrainer, @JsonKey(name: 'shareBodyMeasurements')  bool shareBodyMeasurements, @JsonKey(name: 'shareWeightLogs')  bool shareWeightLogs, @JsonKey(name: 'allowTrainerNotes')  bool allowTrainerNotes)?  $default,) {final _that = this;
switch (_that) {
case _UserPrivacySettings() when $default != null:
return $default(_that.shareEmail,_that.shareTrainingProgress,_that.sharePastDataWithTrainer,_that.shareBodyMeasurements,_that.shareWeightLogs,_that.allowTrainerNotes);case _:
  return null;

}
}

}

/// @nodoc
@JsonSerializable()

class _UserPrivacySettings implements UserPrivacySettings {
  const _UserPrivacySettings({@JsonKey(name: 'shareEmail') required this.shareEmail, @JsonKey(name: 'shareTrainingProgress') required this.shareTrainingProgress, @JsonKey(name: 'sharePastDataWithTrainer') required this.sharePastDataWithTrainer, @JsonKey(name: 'shareBodyMeasurements') required this.shareBodyMeasurements, @JsonKey(name: 'shareWeightLogs') required this.shareWeightLogs, @JsonKey(name: 'allowTrainerNotes') required this.allowTrainerNotes});
  factory _UserPrivacySettings.fromJson(Map<String, dynamic> json) => _$UserPrivacySettingsFromJson(json);

@override@JsonKey(name: 'shareEmail') final  bool shareEmail;
@override@JsonKey(name: 'shareTrainingProgress') final  bool shareTrainingProgress;
@override@JsonKey(name: 'sharePastDataWithTrainer') final  bool sharePastDataWithTrainer;
@override@JsonKey(name: 'shareBodyMeasurements') final  bool shareBodyMeasurements;
@override@JsonKey(name: 'shareWeightLogs') final  bool shareWeightLogs;
@override@JsonKey(name: 'allowTrainerNotes') final  bool allowTrainerNotes;

/// Create a copy of UserPrivacySettings
/// with the given fields replaced by the non-null parameter values.
@override @JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
_$UserPrivacySettingsCopyWith<_UserPrivacySettings> get copyWith => __$UserPrivacySettingsCopyWithImpl<_UserPrivacySettings>(this, _$identity);

@override
Map<String, dynamic> toJson() {
  return _$UserPrivacySettingsToJson(this, );
}

@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is _UserPrivacySettings&&(identical(other.shareEmail, shareEmail) || other.shareEmail == shareEmail)&&(identical(other.shareTrainingProgress, shareTrainingProgress) || other.shareTrainingProgress == shareTrainingProgress)&&(identical(other.sharePastDataWithTrainer, sharePastDataWithTrainer) || other.sharePastDataWithTrainer == sharePastDataWithTrainer)&&(identical(other.shareBodyMeasurements, shareBodyMeasurements) || other.shareBodyMeasurements == shareBodyMeasurements)&&(identical(other.shareWeightLogs, shareWeightLogs) || other.shareWeightLogs == shareWeightLogs)&&(identical(other.allowTrainerNotes, allowTrainerNotes) || other.allowTrainerNotes == allowTrainerNotes));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,shareEmail,shareTrainingProgress,sharePastDataWithTrainer,shareBodyMeasurements,shareWeightLogs,allowTrainerNotes);

@override
String toString() {
  return 'UserPrivacySettings(shareEmail: $shareEmail, shareTrainingProgress: $shareTrainingProgress, sharePastDataWithTrainer: $sharePastDataWithTrainer, shareBodyMeasurements: $shareBodyMeasurements, shareWeightLogs: $shareWeightLogs, allowTrainerNotes: $allowTrainerNotes)';
}


}

/// @nodoc
abstract mixin class _$UserPrivacySettingsCopyWith<$Res> implements $UserPrivacySettingsCopyWith<$Res> {
  factory _$UserPrivacySettingsCopyWith(_UserPrivacySettings value, $Res Function(_UserPrivacySettings) _then) = __$UserPrivacySettingsCopyWithImpl;
@override @useResult
$Res call({
@JsonKey(name: 'shareEmail') bool shareEmail,@JsonKey(name: 'shareTrainingProgress') bool shareTrainingProgress,@JsonKey(name: 'sharePastDataWithTrainer') bool sharePastDataWithTrainer,@JsonKey(name: 'shareBodyMeasurements') bool shareBodyMeasurements,@JsonKey(name: 'shareWeightLogs') bool shareWeightLogs,@JsonKey(name: 'allowTrainerNotes') bool allowTrainerNotes
});




}
/// @nodoc
class __$UserPrivacySettingsCopyWithImpl<$Res>
    implements _$UserPrivacySettingsCopyWith<$Res> {
  __$UserPrivacySettingsCopyWithImpl(this._self, this._then);

  final _UserPrivacySettings _self;
  final $Res Function(_UserPrivacySettings) _then;

/// Create a copy of UserPrivacySettings
/// with the given fields replaced by the non-null parameter values.
@override @pragma('vm:prefer-inline') $Res call({Object? shareEmail = null,Object? shareTrainingProgress = null,Object? sharePastDataWithTrainer = null,Object? shareBodyMeasurements = null,Object? shareWeightLogs = null,Object? allowTrainerNotes = null,}) {
  return _then(_UserPrivacySettings(
shareEmail: null == shareEmail ? _self.shareEmail : shareEmail // ignore: cast_nullable_to_non_nullable
as bool,shareTrainingProgress: null == shareTrainingProgress ? _self.shareTrainingProgress : shareTrainingProgress // ignore: cast_nullable_to_non_nullable
as bool,sharePastDataWithTrainer: null == sharePastDataWithTrainer ? _self.sharePastDataWithTrainer : sharePastDataWithTrainer // ignore: cast_nullable_to_non_nullable
as bool,shareBodyMeasurements: null == shareBodyMeasurements ? _self.shareBodyMeasurements : shareBodyMeasurements // ignore: cast_nullable_to_non_nullable
as bool,shareWeightLogs: null == shareWeightLogs ? _self.shareWeightLogs : shareWeightLogs // ignore: cast_nullable_to_non_nullable
as bool,allowTrainerNotes: null == allowTrainerNotes ? _self.allowTrainerNotes : allowTrainerNotes // ignore: cast_nullable_to_non_nullable
as bool,
  ));
}


}

// dart format on

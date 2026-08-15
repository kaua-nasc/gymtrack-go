// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'exercise_log.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_ExerciseLogRequest _$ExerciseLogRequestFromJson(Map<String, dynamic> json) =>
    _ExerciseLogRequest(
      reps:
          (json['reps'] as List<dynamic>?)
              ?.map((e) => (e as num).toInt())
              .toList() ??
          const [],
      weight:
          (json['weight'] as List<dynamic>?)
              ?.map((e) => (e as num).toDouble())
              .toList() ??
          const [],
      notes: json['notes'] as String?,
    );

Map<String, dynamic> _$ExerciseLogRequestToJson(_ExerciseLogRequest instance) =>
    <String, dynamic>{
      'reps': instance.reps,
      'weight': instance.weight,
      'notes': instance.notes,
    };

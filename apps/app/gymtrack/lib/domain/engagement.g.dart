// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'engagement.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_EngagementSummary _$EngagementSummaryFromJson(Map<String, dynamic> json) =>
    _EngagementSummary(
      totalWorkouts: (json['totalWorkouts'] as num?)?.toInt() ?? 0,
      currentStreak: (json['currentStreak'] as num?)?.toInt() ?? 0,
      longestStreak: (json['longestStreak'] as num?)?.toInt() ?? 0,
      weeklyCompletions: (json['weeklyCompletions'] as num?)?.toInt() ?? 0,
      monthlyCompletions: (json['monthlyCompletions'] as num?)?.toInt() ?? 0,
      lastWorkoutDate: json['lastWorkoutDate'] == null
          ? null
          : DateTime.parse(json['lastWorkoutDate'] as String),
    );

Map<String, dynamic> _$EngagementSummaryToJson(_EngagementSummary instance) =>
    <String, dynamic>{
      'totalWorkouts': instance.totalWorkouts,
      'currentStreak': instance.currentStreak,
      'longestStreak': instance.longestStreak,
      'weeklyCompletions': instance.weeklyCompletions,
      'monthlyCompletions': instance.monthlyCompletions,
      'lastWorkoutDate': instance.lastWorkoutDate?.toIso8601String(),
    };

_BiometricDashboard _$BiometricDashboardFromJson(Map<String, dynamic> json) =>
    _BiometricDashboard(
      bodyMeasurements:
          (json['bodyMeasurements'] as List<dynamic>?)
              ?.map(
                (e) => BodyMeasurementEntry.fromJson(e as Map<String, dynamic>),
              )
              .toList() ??
          const [],
      weightLogs:
          (json['weightLogs'] as List<dynamic>?)
              ?.map((e) => WeightEntry.fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
    );

Map<String, dynamic> _$BiometricDashboardToJson(_BiometricDashboard instance) =>
    <String, dynamic>{
      'bodyMeasurements': instance.bodyMeasurements,
      'weightLogs': instance.weightLogs,
    };

_BodyMeasurementEntry _$BodyMeasurementEntryFromJson(
  Map<String, dynamic> json,
) => _BodyMeasurementEntry(
  id: json['id'] as String,
  type: json['type'] as String,
  value: (json['value'] as num).toDouble(),
  note: json['note'] as String?,
  measuredAt: DateTime.parse(json['measuredAt'] as String),
);

Map<String, dynamic> _$BodyMeasurementEntryToJson(
  _BodyMeasurementEntry instance,
) => <String, dynamic>{
  'id': instance.id,
  'type': instance.type,
  'value': instance.value,
  'note': instance.note,
  'measuredAt': instance.measuredAt.toIso8601String(),
};

_WeightEntry _$WeightEntryFromJson(Map<String, dynamic> json) => _WeightEntry(
  id: json['id'] as String,
  weight: (json['weight'] as num).toDouble(),
  note: json['note'] as String?,
  measuredAt: DateTime.parse(json['measuredAt'] as String),
);

Map<String, dynamic> _$WeightEntryToJson(_WeightEntry instance) =>
    <String, dynamic>{
      'id': instance.id,
      'weight': instance.weight,
      'note': instance.note,
      'measuredAt': instance.measuredAt.toIso8601String(),
    };

_InsightsDashboard _$InsightsDashboardFromJson(Map<String, dynamic> json) =>
    _InsightsDashboard(
      avgCompletionRate: (json['avgCompletionRate'] as num?)?.toDouble(),
      monthly:
          (json['monthly'] as List<dynamic>?)
              ?.map((e) => MonthInsight.fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
      topExercises:
          (json['topExercises'] as List<dynamic>?)
              ?.map((e) => ExerciseInsight.fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
    );

Map<String, dynamic> _$InsightsDashboardToJson(_InsightsDashboard instance) =>
    <String, dynamic>{
      'avgCompletionRate': instance.avgCompletionRate,
      'monthly': instance.monthly,
      'topExercises': instance.topExercises,
    };

_MonthInsight _$MonthInsightFromJson(Map<String, dynamic> json) =>
    _MonthInsight(
      month: json['month'] as String,
      completedDays: (json['completedDays'] as num?)?.toInt() ?? 0,
      totalDays: (json['totalDays'] as num?)?.toInt() ?? 0,
    );

Map<String, dynamic> _$MonthInsightToJson(_MonthInsight instance) =>
    <String, dynamic>{
      'month': instance.month,
      'completedDays': instance.completedDays,
      'totalDays': instance.totalDays,
    };

_ExerciseInsight _$ExerciseInsightFromJson(Map<String, dynamic> json) =>
    _ExerciseInsight(
      name: json['name'] as String,
      totalSets: (json['totalSets'] as num?)?.toInt() ?? 0,
      totalReps: (json['totalReps'] as num?)?.toInt() ?? 0,
      maxWeight: (json['maxWeight'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$ExerciseInsightToJson(_ExerciseInsight instance) =>
    <String, dynamic>{
      'name': instance.name,
      'totalSets': instance.totalSets,
      'totalReps': instance.totalReps,
      'maxWeight': instance.maxWeight,
    };

import 'package:freezed_annotation/freezed_annotation.dart';

part 'engagement.freezed.dart';
part 'engagement.g.dart';

@freezed
abstract class EngagementSummary with _$EngagementSummary {
  const factory EngagementSummary({
    @JsonKey(name: 'totalWorkouts') @Default(0) int totalWorkouts,
    @JsonKey(name: 'currentStreak') @Default(0) int currentStreak,
    @JsonKey(name: 'longestStreak') @Default(0) int longestStreak,
    @JsonKey(name: 'weeklyCompletions') @Default(0) int weeklyCompletions,
    @JsonKey(name: 'monthlyCompletions') @Default(0) int monthlyCompletions,
    @JsonKey(name: 'lastWorkoutDate') DateTime? lastWorkoutDate,
  }) = _EngagementSummary;

  factory EngagementSummary.fromJson(Map<String, dynamic> json) =>
      _$EngagementSummaryFromJson(json);
}

@freezed
abstract class BiometricDashboard with _$BiometricDashboard {
  const factory BiometricDashboard({
    @Default([]) List<BodyMeasurementEntry> bodyMeasurements,
    @Default([]) List<WeightEntry> weightLogs,
  }) = _BiometricDashboard;

  factory BiometricDashboard.fromJson(Map<String, dynamic> json) =>
      _$BiometricDashboardFromJson(json);
}

@freezed
abstract class BodyMeasurementEntry with _$BodyMeasurementEntry {
  const factory BodyMeasurementEntry({
    required String id,
    required String type,
    required double value,
    String? note,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
  }) = _BodyMeasurementEntry;

  factory BodyMeasurementEntry.fromJson(Map<String, dynamic> json) =>
      _$BodyMeasurementEntryFromJson(json);
}

@freezed
abstract class WeightEntry with _$WeightEntry {
  const factory WeightEntry({
    required String id,
    required double weight,
    String? note,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
  }) = _WeightEntry;

  factory WeightEntry.fromJson(Map<String, dynamic> json) =>
      _$WeightEntryFromJson(json);
}

@freezed
abstract class InsightsDashboard with _$InsightsDashboard {
  const factory InsightsDashboard({
    @JsonKey(name: 'avgCompletionRate') double? avgCompletionRate,
    @Default([]) List<MonthInsight> monthly,
    @Default([]) List<ExerciseInsight> topExercises,
  }) = _InsightsDashboard;

  factory InsightsDashboard.fromJson(Map<String, dynamic> json) =>
      _$InsightsDashboardFromJson(json);
}

@freezed
abstract class MonthInsight with _$MonthInsight {
  const factory MonthInsight({
    required String month,
    @JsonKey(name: 'completedDays') @Default(0) int completedDays,
    @JsonKey(name: 'totalDays') @Default(0) int totalDays,
  }) = _MonthInsight;

  factory MonthInsight.fromJson(Map<String, dynamic> json) =>
      _$MonthInsightFromJson(json);
}

@freezed
abstract class ExerciseInsight with _$ExerciseInsight {
  const factory ExerciseInsight({
    required String name,
    @JsonKey(name: 'totalSets') @Default(0) int totalSets,
    @JsonKey(name: 'totalReps') @Default(0) int totalReps,
    @JsonKey(name: 'maxWeight') double? maxWeight,
  }) = _ExerciseInsight;

  factory ExerciseInsight.fromJson(Map<String, dynamic> json) =>
      _$ExerciseInsightFromJson(json);
}

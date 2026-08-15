import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';

part 'training_plan.freezed.dart';
part 'training_plan.g.dart';

@freezed
abstract class TrainingPlan with _$TrainingPlan {
  const factory TrainingPlan({
    required String id,
    @JsonKey(name: 'authorId') required String authorId,
    required String title,
    String? description,
    @TrainingPlanTypeConverter() required TrainingPlanType type,
    @TrainingPlanLevelConverter() required TrainingPlanLevel level,
    @TrainingPlanVisibilityConverter()
    required TrainingPlanVisibility visibility,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
    @Default([]) List<Day> days,
    @JsonKey(name: 'subscriptionCount') @Default(0) int subscriptionCount,
    @JsonKey(name: 'averageRating') double? averageRating,
  }) = _TrainingPlan;

  factory TrainingPlan.fromJson(Map<String, dynamic> json) =>
      _$TrainingPlanFromJson(json);
}

@freezed
abstract class Day with _$Day {
  const factory Day({
    required String id,
    @JsonKey(name: 'planId') required String planId,
    required String name,
    String? description,
    @JsonKey(name: 'orderIndex') required int orderIndex,
    @Default([]) List<Exercise> exercises,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _Day;

  factory Day.fromJson(Map<String, dynamic> json) => _$DayFromJson(json);
}

@freezed
abstract class Exercise with _$Exercise {
  const factory Exercise({
    required String id,
    @JsonKey(name: 'dayId') required String dayId,
    required String name,
    String? description,
    @JsonKey(name: 'seriesCount') int? seriesCount,
    @JsonKey(name: 'repetitionCount') int? repetitionCount,
    double? weight,
    String? videoUrl,
    String? imageUrl,
    @JsonKey(name: 'restTimeSeconds') int? restTimeSeconds,
    @JsonKey(name: 'orderIndex') required int orderIndex,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _Exercise;

  factory Exercise.fromJson(Map<String, dynamic> json) =>
      _$ExerciseFromJson(json);
}

class TrainingPlanTypeConverter
    implements JsonConverter<TrainingPlanType, String> {
  const TrainingPlanTypeConverter();

  @override
  TrainingPlanType fromJson(String json) =>
      TrainingPlanType.values.firstWhere((e) => e.name == json);

  @override
  String toJson(TrainingPlanType object) => object.name;
}

class TrainingPlanLevelConverter
    implements JsonConverter<TrainingPlanLevel, String> {
  const TrainingPlanLevelConverter();

  @override
  TrainingPlanLevel fromJson(String json) =>
      TrainingPlanLevel.values.firstWhere((e) => e.name == json);

  @override
  String toJson(TrainingPlanLevel object) => object.name;
}

class TrainingPlanVisibilityConverter
    implements JsonConverter<TrainingPlanVisibility, String> {
  const TrainingPlanVisibilityConverter();

  @override
  TrainingPlanVisibility fromJson(String json) =>
      TrainingPlanVisibility.values.firstWhere((e) => e.name == json);

  @override
  String toJson(TrainingPlanVisibility object) => object.name;
}

import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';
import 'training_plan.dart';

part 'subscription.freezed.dart';
part 'subscription.g.dart';

@freezed
abstract class PlanSubscription with _$PlanSubscription {
  const factory PlanSubscription({
    required String id,
    @JsonKey(name: 'planId') required String planId,
    @JsonKey(name: 'userId') required String userId,
    @PlanSubscriptionStatusConverter() required PlanSubscriptionStatus status,
    @PlanSubscriptionTypeConverter() required PlanSubscriptionType type,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _PlanSubscription;

  factory PlanSubscription.fromJson(Map<String, dynamic> json) =>
      _$PlanSubscriptionFromJson(json);
}

@freezed
abstract class SubscribeRequest with _$SubscribeRequest {
  const factory SubscribeRequest({
    @PlanSubscriptionTypeConverter() required PlanSubscriptionType type,
  }) = _SubscribeRequest;

  factory SubscribeRequest.fromJson(Map<String, dynamic> json) =>
      _$SubscribeRequestFromJson(json);
}

@freezed
abstract class ChangeSubscriptionStatusRequest
    with _$ChangeSubscriptionStatusRequest {
  const factory ChangeSubscriptionStatusRequest({
    @PlanSubscriptionStatusConverter() required PlanSubscriptionStatus status,
  }) = _ChangeSubscriptionStatusRequest;

  factory ChangeSubscriptionStatusRequest.fromJson(Map<String, dynamic> json) =>
      _$ChangeSubscriptionStatusRequestFromJson(json);
}

@freezed
abstract class ChangeSubscriptionPrivacyRequest
    with _$ChangeSubscriptionPrivacyRequest {
  const factory ChangeSubscriptionPrivacyRequest({
    @PlanSubscriptionTypeConverter() required PlanSubscriptionType type,
  }) = _ChangeSubscriptionPrivacyRequest;

  factory ChangeSubscriptionPrivacyRequest.fromJson(
    Map<String, dynamic> json,
  ) => _$ChangeSubscriptionPrivacyRequestFromJson(json);
}

@freezed
abstract class ListSubscriptionFilters with _$ListSubscriptionFilters {
  const factory ListSubscriptionFilters({
    @PlanSubscriptionStatusConverter()
    @JsonKey(includeIfNull: true)
    PlanSubscriptionStatus? status,
    @PlanSubscriptionTypeConverter()
    @JsonKey(includeIfNull: true)
    PlanSubscriptionType? type,
    @TrainingPlanTypeConverter()
    @JsonKey(includeIfNull: true, name: 'planType')
    TrainingPlanType? planType,
    @TrainingPlanVisibilityConverter()
    @JsonKey(includeIfNull: true)
    TrainingPlanVisibility? visibility,
    @TrainingPlanLevelConverter()
    @JsonKey(includeIfNull: true)
    TrainingPlanLevel? level,
    @JsonKey(includeIfNull: true, name: 'authorId') String? authorId,
  }) = _ListSubscriptionFilters;

  factory ListSubscriptionFilters.fromJson(Map<String, dynamic> json) =>
      _$ListSubscriptionFiltersFromJson(json);
}

@freezed
abstract class WeeklyDayProgress with _$WeeklyDayProgress {
  const factory WeeklyDayProgress({
    @Default([]) List<PlanDayProgress> days,
    @JsonKey(name: 'completedDays') @Default(0) int completedDays,
    @JsonKey(name: 'totalDays') @Default(0) int totalDays,
  }) = _WeeklyDayProgress;

  factory WeeklyDayProgress.fromJson(Map<String, dynamic> json) =>
      _$WeeklyDayProgressFromJson(json);
}

@freezed
abstract class PlanDayProgress with _$PlanDayProgress {
  const factory PlanDayProgress({
    @JsonKey(name: 'subscriptionId') required String subscriptionId,
    @JsonKey(name: 'dayId') required String dayId,
    @JsonKey(name: 'dayName') required String dayName,
    @JsonKey(name: 'planId') required String planId,
    @JsonKey(name: 'planTitle') required String planTitle,
    String? status,
    @JsonKey(name: 'scheduledDate') DateTime? scheduledDate,
    @JsonKey(name: 'startedAt') DateTime? startedAt,
    @JsonKey(name: 'completedAt') DateTime? completedAt,
  }) = _PlanDayProgress;

  factory PlanDayProgress.fromJson(Map<String, dynamic> json) =>
      _$PlanDayProgressFromJson(json);
}

class PlanSubscriptionStatusConverter
    implements JsonConverter<PlanSubscriptionStatus, String> {
  const PlanSubscriptionStatusConverter();

  @override
  PlanSubscriptionStatus fromJson(String json) =>
      PlanSubscriptionStatus.values.firstWhere((e) => e.name == json);

  @override
  String toJson(PlanSubscriptionStatus object) => object.name;
}

class PlanSubscriptionTypeConverter
    implements JsonConverter<PlanSubscriptionType, String> {
  const PlanSubscriptionTypeConverter();

  @override
  PlanSubscriptionType fromJson(String json) =>
      PlanSubscriptionType.values.firstWhere((e) => e.name == json);

  @override
  String toJson(PlanSubscriptionType object) => object.name;
}

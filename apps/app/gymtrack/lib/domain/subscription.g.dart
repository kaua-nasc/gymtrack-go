// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'subscription.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_PlanSubscription _$PlanSubscriptionFromJson(Map<String, dynamic> json) =>
    _PlanSubscription(
      id: json['id'] as String,
      planId: json['planId'] as String,
      userId: json['userId'] as String,
      status: const PlanSubscriptionStatusConverter().fromJson(
        json['status'] as String,
      ),
      type: const PlanSubscriptionTypeConverter().fromJson(
        json['type'] as String,
      ),
      createdAt: DateTime.parse(json['createdAt'] as String),
      updatedAt: DateTime.parse(json['updatedAt'] as String),
    );

Map<String, dynamic> _$PlanSubscriptionToJson(_PlanSubscription instance) =>
    <String, dynamic>{
      'id': instance.id,
      'planId': instance.planId,
      'userId': instance.userId,
      'status': const PlanSubscriptionStatusConverter().toJson(instance.status),
      'type': const PlanSubscriptionTypeConverter().toJson(instance.type),
      'createdAt': instance.createdAt.toIso8601String(),
      'updatedAt': instance.updatedAt.toIso8601String(),
    };

_SubscribeRequest _$SubscribeRequestFromJson(Map<String, dynamic> json) =>
    _SubscribeRequest(
      type: const PlanSubscriptionTypeConverter().fromJson(
        json['type'] as String,
      ),
    );

Map<String, dynamic> _$SubscribeRequestToJson(_SubscribeRequest instance) =>
    <String, dynamic>{
      'type': const PlanSubscriptionTypeConverter().toJson(instance.type),
    };

_ChangeSubscriptionStatusRequest _$ChangeSubscriptionStatusRequestFromJson(
  Map<String, dynamic> json,
) => _ChangeSubscriptionStatusRequest(
  status: const PlanSubscriptionStatusConverter().fromJson(
    json['status'] as String,
  ),
);

Map<String, dynamic> _$ChangeSubscriptionStatusRequestToJson(
  _ChangeSubscriptionStatusRequest instance,
) => <String, dynamic>{
  'status': const PlanSubscriptionStatusConverter().toJson(instance.status),
};

_ChangeSubscriptionPrivacyRequest _$ChangeSubscriptionPrivacyRequestFromJson(
  Map<String, dynamic> json,
) => _ChangeSubscriptionPrivacyRequest(
  type: const PlanSubscriptionTypeConverter().fromJson(json['type'] as String),
);

Map<String, dynamic> _$ChangeSubscriptionPrivacyRequestToJson(
  _ChangeSubscriptionPrivacyRequest instance,
) => <String, dynamic>{
  'type': const PlanSubscriptionTypeConverter().toJson(instance.type),
};

_ListSubscriptionFilters _$ListSubscriptionFiltersFromJson(
  Map<String, dynamic> json,
) => _ListSubscriptionFilters(
  status: _$JsonConverterFromJson<String, PlanSubscriptionStatus>(
    json['status'],
    const PlanSubscriptionStatusConverter().fromJson,
  ),
  type: _$JsonConverterFromJson<String, PlanSubscriptionType>(
    json['type'],
    const PlanSubscriptionTypeConverter().fromJson,
  ),
  planType: _$JsonConverterFromJson<String, TrainingPlanType>(
    json['planType'],
    const TrainingPlanTypeConverter().fromJson,
  ),
  visibility: _$JsonConverterFromJson<String, TrainingPlanVisibility>(
    json['visibility'],
    const TrainingPlanVisibilityConverter().fromJson,
  ),
  level: _$JsonConverterFromJson<String, TrainingPlanLevel>(
    json['level'],
    const TrainingPlanLevelConverter().fromJson,
  ),
  authorId: json['authorId'] as String?,
);

Map<String, dynamic> _$ListSubscriptionFiltersToJson(
  _ListSubscriptionFilters instance,
) => <String, dynamic>{
  'status': _$JsonConverterToJson<String, PlanSubscriptionStatus>(
    instance.status,
    const PlanSubscriptionStatusConverter().toJson,
  ),
  'type': _$JsonConverterToJson<String, PlanSubscriptionType>(
    instance.type,
    const PlanSubscriptionTypeConverter().toJson,
  ),
  'planType': _$JsonConverterToJson<String, TrainingPlanType>(
    instance.planType,
    const TrainingPlanTypeConverter().toJson,
  ),
  'visibility': _$JsonConverterToJson<String, TrainingPlanVisibility>(
    instance.visibility,
    const TrainingPlanVisibilityConverter().toJson,
  ),
  'level': _$JsonConverterToJson<String, TrainingPlanLevel>(
    instance.level,
    const TrainingPlanLevelConverter().toJson,
  ),
  'authorId': instance.authorId,
};

Value? _$JsonConverterFromJson<Json, Value>(
  Object? json,
  Value? Function(Json json) fromJson,
) => json == null ? null : fromJson(json as Json);

Json? _$JsonConverterToJson<Json, Value>(
  Value? value,
  Json? Function(Value value) toJson,
) => value == null ? null : toJson(value);

_WeeklyDayProgress _$WeeklyDayProgressFromJson(Map<String, dynamic> json) =>
    _WeeklyDayProgress(
      days:
          (json['days'] as List<dynamic>?)
              ?.map((e) => PlanDayProgress.fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
      completedDays: (json['completedDays'] as num?)?.toInt() ?? 0,
      totalDays: (json['totalDays'] as num?)?.toInt() ?? 0,
    );

Map<String, dynamic> _$WeeklyDayProgressToJson(_WeeklyDayProgress instance) =>
    <String, dynamic>{
      'days': instance.days,
      'completedDays': instance.completedDays,
      'totalDays': instance.totalDays,
    };

_PlanDayProgress _$PlanDayProgressFromJson(Map<String, dynamic> json) =>
    _PlanDayProgress(
      subscriptionId: json['subscriptionId'] as String,
      dayId: json['dayId'] as String,
      dayName: json['dayName'] as String,
      planId: json['planId'] as String,
      planTitle: json['planTitle'] as String,
      status: json['status'] as String?,
      scheduledDate: json['scheduledDate'] == null
          ? null
          : DateTime.parse(json['scheduledDate'] as String),
      startedAt: json['startedAt'] == null
          ? null
          : DateTime.parse(json['startedAt'] as String),
      completedAt: json['completedAt'] == null
          ? null
          : DateTime.parse(json['completedAt'] as String),
    );

Map<String, dynamic> _$PlanDayProgressToJson(_PlanDayProgress instance) =>
    <String, dynamic>{
      'subscriptionId': instance.subscriptionId,
      'dayId': instance.dayId,
      'dayName': instance.dayName,
      'planId': instance.planId,
      'planTitle': instance.planTitle,
      'status': instance.status,
      'scheduledDate': instance.scheduledDate?.toIso8601String(),
      'startedAt': instance.startedAt?.toIso8601String(),
      'completedAt': instance.completedAt?.toIso8601String(),
    };

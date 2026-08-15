// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'training_plan.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_TrainingPlan _$TrainingPlanFromJson(Map<String, dynamic> json) =>
    _TrainingPlan(
      id: json['id'] as String,
      authorId: json['authorId'] as String,
      title: json['title'] as String,
      description: json['description'] as String?,
      type: const TrainingPlanTypeConverter().fromJson(json['type'] as String),
      level: const TrainingPlanLevelConverter().fromJson(
        json['level'] as String,
      ),
      visibility: const TrainingPlanVisibilityConverter().fromJson(
        json['visibility'] as String,
      ),
      createdAt: DateTime.parse(json['createdAt'] as String),
      updatedAt: DateTime.parse(json['updatedAt'] as String),
      days:
          (json['days'] as List<dynamic>?)
              ?.map((e) => Day.fromJson(e as Map<String, dynamic>))
              .toList() ??
          const [],
      subscriptionCount: (json['subscriptionCount'] as num?)?.toInt() ?? 0,
      averageRating: (json['averageRating'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$TrainingPlanToJson(_TrainingPlan instance) =>
    <String, dynamic>{
      'id': instance.id,
      'authorId': instance.authorId,
      'title': instance.title,
      'description': instance.description,
      'type': const TrainingPlanTypeConverter().toJson(instance.type),
      'level': const TrainingPlanLevelConverter().toJson(instance.level),
      'visibility': const TrainingPlanVisibilityConverter().toJson(
        instance.visibility,
      ),
      'createdAt': instance.createdAt.toIso8601String(),
      'updatedAt': instance.updatedAt.toIso8601String(),
      'days': instance.days,
      'subscriptionCount': instance.subscriptionCount,
      'averageRating': instance.averageRating,
    };

_Day _$DayFromJson(Map<String, dynamic> json) => _Day(
  id: json['id'] as String,
  planId: json['planId'] as String,
  name: json['name'] as String,
  description: json['description'] as String?,
  orderIndex: (json['orderIndex'] as num).toInt(),
  exercises:
      (json['exercises'] as List<dynamic>?)
          ?.map((e) => Exercise.fromJson(e as Map<String, dynamic>))
          .toList() ??
      const [],
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
);

Map<String, dynamic> _$DayToJson(_Day instance) => <String, dynamic>{
  'id': instance.id,
  'planId': instance.planId,
  'name': instance.name,
  'description': instance.description,
  'orderIndex': instance.orderIndex,
  'exercises': instance.exercises,
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
};

_Exercise _$ExerciseFromJson(Map<String, dynamic> json) => _Exercise(
  id: json['id'] as String,
  dayId: json['dayId'] as String,
  name: json['name'] as String,
  description: json['description'] as String?,
  seriesCount: (json['seriesCount'] as num?)?.toInt(),
  repetitionCount: (json['repetitionCount'] as num?)?.toInt(),
  weight: (json['weight'] as num?)?.toDouble(),
  videoUrl: json['videoUrl'] as String?,
  imageUrl: json['imageUrl'] as String?,
  restTimeSeconds: (json['restTimeSeconds'] as num?)?.toInt(),
  orderIndex: (json['orderIndex'] as num).toInt(),
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
);

Map<String, dynamic> _$ExerciseToJson(_Exercise instance) => <String, dynamic>{
  'id': instance.id,
  'dayId': instance.dayId,
  'name': instance.name,
  'description': instance.description,
  'seriesCount': instance.seriesCount,
  'repetitionCount': instance.repetitionCount,
  'weight': instance.weight,
  'videoUrl': instance.videoUrl,
  'imageUrl': instance.imageUrl,
  'restTimeSeconds': instance.restTimeSeconds,
  'orderIndex': instance.orderIndex,
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
};

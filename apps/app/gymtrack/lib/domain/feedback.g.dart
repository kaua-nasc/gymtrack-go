// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'feedback.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_TrainingPlanFeedback _$TrainingPlanFeedbackFromJson(
  Map<String, dynamic> json,
) => _TrainingPlanFeedback(
  id: json['id'] as String,
  planId: json['planId'] as String,
  userId: json['userId'] as String,
  rating: (json['rating'] as num).toDouble(),
  message: json['message'] as String?,
  createdAt: DateTime.parse(json['createdAt'] as String),
  updatedAt: DateTime.parse(json['updatedAt'] as String),
);

Map<String, dynamic> _$TrainingPlanFeedbackToJson(
  _TrainingPlanFeedback instance,
) => <String, dynamic>{
  'id': instance.id,
  'planId': instance.planId,
  'userId': instance.userId,
  'rating': instance.rating,
  'message': instance.message,
  'createdAt': instance.createdAt.toIso8601String(),
  'updatedAt': instance.updatedAt.toIso8601String(),
};

_AddFeedbackRequest _$AddFeedbackRequestFromJson(Map<String, dynamic> json) =>
    _AddFeedbackRequest(
      rating: (json['rating'] as num).toDouble(),
      message: json['message'] as String?,
    );

Map<String, dynamic> _$AddFeedbackRequestToJson(_AddFeedbackRequest instance) =>
    <String, dynamic>{'rating': instance.rating, 'message': instance.message};

// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'weight_log.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_WeightLog _$WeightLogFromJson(Map<String, dynamic> json) => _WeightLog(
  id: json['id'] as String,
  weight: (json['weight'] as num).toDouble(),
  note: json['note'] as String?,
  userId: json['userId'] as String,
  trainerId: json['trainerId'] as String?,
  measuredAt: DateTime.parse(json['measuredAt'] as String),
  createdAt: DateTime.parse(json['createdAt'] as String),
);

Map<String, dynamic> _$WeightLogToJson(_WeightLog instance) =>
    <String, dynamic>{
      'id': instance.id,
      'weight': instance.weight,
      'note': instance.note,
      'userId': instance.userId,
      'trainerId': instance.trainerId,
      'measuredAt': instance.measuredAt.toIso8601String(),
      'createdAt': instance.createdAt.toIso8601String(),
    };

_CreateWeightLogRequest _$CreateWeightLogRequestFromJson(
  Map<String, dynamic> json,
) => _CreateWeightLogRequest(
  weight: (json['weight'] as num).toDouble(),
  measuredAt: DateTime.parse(json['measuredAt'] as String),
);

Map<String, dynamic> _$CreateWeightLogRequestToJson(
  _CreateWeightLogRequest instance,
) => <String, dynamic>{
  'weight': instance.weight,
  'measuredAt': instance.measuredAt.toIso8601String(),
};

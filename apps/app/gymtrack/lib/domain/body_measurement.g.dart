// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'body_measurement.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_BodyMeasurement _$BodyMeasurementFromJson(Map<String, dynamic> json) =>
    _BodyMeasurement(
      id: json['id'] as String,
      type: const BodyMeasurementTypeConverter().fromJson(
        json['type'] as String,
      ),
      value: (json['value'] as num).toDouble(),
      unit: json['unit'] as String?,
      note: json['note'] as String?,
      userId: json['userId'] as String,
      trainerId: json['trainerId'] as String?,
      measuredAt: DateTime.parse(json['measuredAt'] as String),
      createdAt: DateTime.parse(json['createdAt'] as String),
    );

Map<String, dynamic> _$BodyMeasurementToJson(_BodyMeasurement instance) =>
    <String, dynamic>{
      'id': instance.id,
      'type': const BodyMeasurementTypeConverter().toJson(instance.type),
      'value': instance.value,
      'unit': instance.unit,
      'note': instance.note,
      'userId': instance.userId,
      'trainerId': instance.trainerId,
      'measuredAt': instance.measuredAt.toIso8601String(),
      'createdAt': instance.createdAt.toIso8601String(),
    };

_CreateBodyMeasurementRequest _$CreateBodyMeasurementRequestFromJson(
  Map<String, dynamic> json,
) => _CreateBodyMeasurementRequest(
  type: const BodyMeasurementTypeConverter().fromJson(json['type'] as String),
  value: (json['value'] as num).toDouble(),
  measuredAt: DateTime.parse(json['measuredAt'] as String),
);

Map<String, dynamic> _$CreateBodyMeasurementRequestToJson(
  _CreateBodyMeasurementRequest instance,
) => <String, dynamic>{
  'type': const BodyMeasurementTypeConverter().toJson(instance.type),
  'value': instance.value,
  'measuredAt': instance.measuredAt.toIso8601String(),
};

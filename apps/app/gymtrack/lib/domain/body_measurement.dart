import 'package:freezed_annotation/freezed_annotation.dart';

import 'enums.dart';

part 'body_measurement.freezed.dart';
part 'body_measurement.g.dart';

@freezed
abstract class BodyMeasurement with _$BodyMeasurement {
  const factory BodyMeasurement({
    required String id,
    @BodyMeasurementTypeConverter() required BodyMeasurementType type,
    required double value,
    @JsonKey(name: 'unit') String? unit,
    String? note,
    @JsonKey(name: 'userId') required String userId,
    @JsonKey(name: 'trainerId') String? trainerId,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
  }) = _BodyMeasurement;

  factory BodyMeasurement.fromJson(Map<String, dynamic> json) =>
      _$BodyMeasurementFromJson(json);
}

@freezed
abstract class CreateBodyMeasurementRequest
    with _$CreateBodyMeasurementRequest {
  const factory CreateBodyMeasurementRequest({
    @BodyMeasurementTypeConverter() required BodyMeasurementType type,
    required double value,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
  }) = _CreateBodyMeasurementRequest;

  factory CreateBodyMeasurementRequest.fromJson(Map<String, dynamic> json) =>
      _$CreateBodyMeasurementRequestFromJson(json);
}

class BodyMeasurementTypeConverter
    implements JsonConverter<BodyMeasurementType, String> {
  const BodyMeasurementTypeConverter();

  @override
  BodyMeasurementType fromJson(String json) =>
      BodyMeasurementType.values.firstWhere((e) => e.name == json);

  @override
  String toJson(BodyMeasurementType object) => object.name;
}

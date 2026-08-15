import 'package:freezed_annotation/freezed_annotation.dart';

part 'weight_log.freezed.dart';
part 'weight_log.g.dart';

@freezed
abstract class WeightLog with _$WeightLog {
  const factory WeightLog({
    required String id,
    required double weight,
    String? note,
    @JsonKey(name: 'userId') required String userId,
    @JsonKey(name: 'trainerId') String? trainerId,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
  }) = _WeightLog;

  factory WeightLog.fromJson(Map<String, dynamic> json) =>
      _$WeightLogFromJson(json);
}

@freezed
abstract class CreateWeightLogRequest with _$CreateWeightLogRequest {
  const factory CreateWeightLogRequest({
    required double weight,
    @JsonKey(name: 'measuredAt') required DateTime measuredAt,
  }) = _CreateWeightLogRequest;

  factory CreateWeightLogRequest.fromJson(Map<String, dynamic> json) =>
      _$CreateWeightLogRequestFromJson(json);
}

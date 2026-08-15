import 'package:freezed_annotation/freezed_annotation.dart';

part 'feedback.freezed.dart';
part 'feedback.g.dart';

@freezed
abstract class TrainingPlanFeedback with _$TrainingPlanFeedback {
  const factory TrainingPlanFeedback({
    required String id,
    @JsonKey(name: 'planId') required String planId,
    @JsonKey(name: 'userId') required String userId,
    required double rating,
    String? message,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _TrainingPlanFeedback;

  factory TrainingPlanFeedback.fromJson(Map<String, dynamic> json) =>
      _$TrainingPlanFeedbackFromJson(json);
}

@freezed
abstract class AddFeedbackRequest with _$AddFeedbackRequest {
  const factory AddFeedbackRequest({
    required double rating,
    @JsonKey(includeIfNull: true) String? message,
  }) = _AddFeedbackRequest;

  factory AddFeedbackRequest.fromJson(Map<String, dynamic> json) =>
      _$AddFeedbackRequestFromJson(json);
}

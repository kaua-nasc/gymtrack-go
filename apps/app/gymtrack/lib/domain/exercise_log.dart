import 'package:freezed_annotation/freezed_annotation.dart';

part 'exercise_log.freezed.dart';
part 'exercise_log.g.dart';

@freezed
abstract class ExerciseLogRequest with _$ExerciseLogRequest {
  const factory ExerciseLogRequest({
    @Default([]) List<int> reps,
    @Default([]) List<double> weight,
    @JsonKey(includeIfNull: true) String? notes,
  }) = _ExerciseLogRequest;

  factory ExerciseLogRequest.fromJson(Map<String, dynamic> json) =>
      _$ExerciseLogRequestFromJson(json);
}

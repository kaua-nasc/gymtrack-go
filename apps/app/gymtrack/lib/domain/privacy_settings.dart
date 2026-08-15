import 'package:freezed_annotation/freezed_annotation.dart';

part 'privacy_settings.freezed.dart';
part 'privacy_settings.g.dart';

@freezed
abstract class UserPrivacySettings with _$UserPrivacySettings {
  const factory UserPrivacySettings({
    @JsonKey(name: 'shareEmail') required bool shareEmail,
    @JsonKey(name: 'shareTrainingProgress') required bool shareTrainingProgress,
    @JsonKey(name: 'sharePastDataWithTrainer')
    required bool sharePastDataWithTrainer,
    @JsonKey(name: 'shareBodyMeasurements') required bool shareBodyMeasurements,
    @JsonKey(name: 'shareWeightLogs') required bool shareWeightLogs,
    @JsonKey(name: 'allowTrainerNotes') required bool allowTrainerNotes,
  }) = _UserPrivacySettings;

  factory UserPrivacySettings.fromJson(Map<String, dynamic> json) =>
      _$UserPrivacySettingsFromJson(json);
}

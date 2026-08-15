// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'privacy_settings.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_UserPrivacySettings _$UserPrivacySettingsFromJson(Map<String, dynamic> json) =>
    _UserPrivacySettings(
      shareEmail: json['shareEmail'] as bool,
      shareTrainingProgress: json['shareTrainingProgress'] as bool,
      sharePastDataWithTrainer: json['sharePastDataWithTrainer'] as bool,
      shareBodyMeasurements: json['shareBodyMeasurements'] as bool,
      shareWeightLogs: json['shareWeightLogs'] as bool,
      allowTrainerNotes: json['allowTrainerNotes'] as bool,
    );

Map<String, dynamic> _$UserPrivacySettingsToJson(
  _UserPrivacySettings instance,
) => <String, dynamic>{
  'shareEmail': instance.shareEmail,
  'shareTrainingProgress': instance.shareTrainingProgress,
  'sharePastDataWithTrainer': instance.sharePastDataWithTrainer,
  'shareBodyMeasurements': instance.shareBodyMeasurements,
  'shareWeightLogs': instance.shareWeightLogs,
  'allowTrainerNotes': instance.allowTrainerNotes,
};

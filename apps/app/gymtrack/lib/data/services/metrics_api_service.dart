import '../../domain/body_measurement.dart';
import '../../domain/weight_log.dart';
import 'api/api_client.dart';

class MetricsApiService {
  MetricsApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<void>> createBodyMeasurement(
    CreateBodyMeasurementRequest request,
  ) => _apiClient.post(
    '/identity/users/body-measurements',
    body: request.toJson(),
  );

  Future<Result<BodyMeasurement>> getLatestBodyMeasurement() => _apiClient.get(
    '/identity/users/body-measurements/latest',
    fromJson: (json) => BodyMeasurement.fromJson(json),
  );

  Future<Result<void>> createWeightLog(CreateWeightLogRequest request) =>
      _apiClient.post('/identity/users/weight-logs', body: request.toJson());

  Future<Result<WeightLog>> getLatestWeightLog() => _apiClient.get(
    '/identity/users/weight-logs/latest',
    fromJson: (json) => WeightLog.fromJson(json),
  );
}

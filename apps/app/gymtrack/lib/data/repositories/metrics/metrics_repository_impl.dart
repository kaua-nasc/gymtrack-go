import '../../../domain/body_measurement.dart';
import '../../../domain/weight_log.dart';
import '../../services/api/api_client.dart';
import '../../services/metrics_api_service.dart';
import 'metrics_repository.dart';

class MetricsRepositoryImpl implements MetricsRepository {
  MetricsRepositoryImpl({required MetricsApiService metricsApiService})
    : _metricsApiService = metricsApiService;

  final MetricsApiService _metricsApiService;

  @override
  Future<void> createBodyMeasurement(
    CreateBodyMeasurementRequest request,
  ) async {
    final result = await _metricsApiService.createBodyMeasurement(request);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<BodyMeasurement> getLatestBodyMeasurement() async {
    final result = await _metricsApiService.getLatestBodyMeasurement();
    return switch (result) {
      Success<BodyMeasurement>() => result.data,
      Failure<BodyMeasurement>() => throw Exception(result.message),
    };
  }

  @override
  Future<void> createWeightLog(CreateWeightLogRequest request) async {
    final result = await _metricsApiService.createWeightLog(request);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<WeightLog> getLatestWeightLog() async {
    final result = await _metricsApiService.getLatestWeightLog();
    return switch (result) {
      Success<WeightLog>() => result.data,
      Failure<WeightLog>() => throw Exception(result.message),
    };
  }
}

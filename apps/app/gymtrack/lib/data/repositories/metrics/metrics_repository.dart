import '../../../domain/body_measurement.dart';
import '../../../domain/weight_log.dart';

abstract class MetricsRepository {
  Future<void> createBodyMeasurement(CreateBodyMeasurementRequest request);
  Future<BodyMeasurement> getLatestBodyMeasurement();
  Future<void> createWeightLog(CreateWeightLogRequest request);
  Future<WeightLog> getLatestWeightLog();
}

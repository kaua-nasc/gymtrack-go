import '../../domain/engagement.dart';
import 'api/api_client.dart';

class DashboardApiService {
  DashboardApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<EngagementSummary>> getStudentEngagement(String studentId) =>
      _apiClient.get(
        '/identity/users/trainers/students/$studentId/dashboard/engagement',
        fromJson: (json) => EngagementSummary.fromJson(json),
      );

  Future<Result<BiometricDashboard>> getStudentBiometrics(
    String studentId, {
    required DateTime start,
    required DateTime end,
  }) => _apiClient.get(
    '/identity/users/trainers/students/$studentId/dashboard/biometrics',
    queryParams: {
      'start': start.toIso8601String(),
      'end': end.toIso8601String(),
    },
    fromJson: (json) => BiometricDashboard.fromJson(json),
  );

  Future<Result<InsightsDashboard>> getStudentInsights(String studentId) =>
      _apiClient.get(
        '/identity/users/trainers/students/$studentId/dashboard/insights',
        fromJson: (json) => InsightsDashboard.fromJson(json),
      );
}

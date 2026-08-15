import '../../../domain/engagement.dart';
import '../../services/api/api_client.dart';
import '../../services/dashboard_api_service.dart';
import 'dashboard_repository.dart';

class DashboardRepositoryImpl implements DashboardRepository {
  DashboardRepositoryImpl({required DashboardApiService dashboardApiService})
    : _dashboardApiService = dashboardApiService;

  final DashboardApiService _dashboardApiService;

  @override
  Future<EngagementSummary> getStudentEngagement(String studentId) async {
    final result = await _dashboardApiService.getStudentEngagement(studentId);
    return switch (result) {
      Success<EngagementSummary>() => result.data,
      Failure<EngagementSummary>() => throw Exception(result.message),
    };
  }

  @override
  Future<BiometricDashboard> getStudentBiometrics(
    String studentId, {
    required DateTime start,
    required DateTime end,
  }) async {
    final result = await _dashboardApiService.getStudentBiometrics(
      studentId,
      start: start,
      end: end,
    );
    return switch (result) {
      Success<BiometricDashboard>() => result.data,
      Failure<BiometricDashboard>() => throw Exception(result.message),
    };
  }

  @override
  Future<InsightsDashboard> getStudentInsights(String studentId) async {
    final result = await _dashboardApiService.getStudentInsights(studentId);
    return switch (result) {
      Success<InsightsDashboard>() => result.data,
      Failure<InsightsDashboard>() => throw Exception(result.message),
    };
  }
}

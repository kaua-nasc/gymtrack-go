import '../../../domain/engagement.dart';

abstract class DashboardRepository {
  Future<EngagementSummary> getStudentEngagement(String studentId);
  Future<BiometricDashboard> getStudentBiometrics(
    String studentId, {
    required DateTime start,
    required DateTime end,
  });
  Future<InsightsDashboard> getStudentInsights(String studentId);
}

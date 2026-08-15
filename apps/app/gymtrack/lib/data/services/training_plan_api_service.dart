import '../../domain/exercise_log.dart';
import '../../domain/feedback.dart';
import '../../domain/subscription.dart';
import '../../domain/training_plan.dart';
import 'api/api_client.dart';

class TrainingPlanApiService {
  TrainingPlanApiService({required ApiClient apiClient})
    : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<void>> createPlan(TrainingPlan plan) =>
      _apiClient.post('/training-plans', body: plan.toJson());

  Future<Result<void>> createPlanForStudent(
    String studentId,
    TrainingPlan plan,
  ) => _apiClient.post(
    '/training-plans/student/$studentId',
    body: plan.toJson(),
  );

  Future<Result<void>> updatePlan(String id, TrainingPlan plan) =>
      _apiClient.put('/training-plans/$id', body: plan.toJson());

  Future<Result<void>> deletePlan(String id) =>
      _apiClient.delete('/training-plans/$id');

  Future<Result<void>> createDay(String planId) =>
      _apiClient.post('/training-plans/$planId/days');

  Future<Result<void>> deleteDay(String planId, String dayId) =>
      _apiClient.delete('/training-plans/$planId/days/$dayId');

  Future<Result<void>> createExercise(String planId, String dayId) =>
      _apiClient.post('/training-plans/$planId/days/$dayId/exercises');

  Future<Result<void>> deleteExercise(
    String planId,
    String dayId,
    String exerciseId,
  ) => _apiClient.delete(
    '/training-plans/$planId/days/$dayId/exercises/$exerciseId',
  );

  Future<Result<void>> logExercise(
    String planId,
    String dayId,
    String exerciseId,
    ExerciseLogRequest request,
  ) => _apiClient.post(
    '/training-plans/$planId/days/$dayId/exercises/$exerciseId/logs',
    body: request.toJson(),
  );

  Future<Result<void>> subscribe(String planId, SubscribeRequest request) =>
      _apiClient.post(
        '/training-plans/$planId/subscriptions',
        body: request.toJson(),
      );

  Future<Result<void>> unsubscribe(String planId) =>
      _apiClient.delete('/training-plans/$planId/subscriptions');

  Future<Result<void>> changeSubscriptionStatus(
    String planId,
    ChangeSubscriptionStatusRequest request,
  ) => _apiClient.post(
    '/training-plans/$planId/subscriptions/send',
    body: request.toJson(),
  );

  Future<Result<void>> changeSubscriptionPrivacy(
    String planId,
    ChangeSubscriptionPrivacyRequest request,
  ) => _apiClient.post(
    '/training-plans/$planId/subscriptions/privacy',
    body: request.toJson(),
  );

  Future<Result<void>> completeDay(String subscriptionId, String dayId) =>
      _apiClient.post(
        '/training-plans/subscriptions/$subscriptionId/days/$dayId/complete',
      );

  Future<Result<void>> cancelDay(String subscriptionId, String dayId) =>
      _apiClient.post(
        '/training-plans/subscriptions/$subscriptionId/days/$dayId/cancel',
      );

  Future<Result<void>> startDay(String subscriptionId, String dayId) =>
      _apiClient.post(
        '/training-plans/subscriptions/$subscriptionId/days/$dayId/start',
      );

  Future<Result<void>> addFeedback(String planId, AddFeedbackRequest request) =>
      _apiClient.post(
        '/training-plans/$planId/feedback',
        body: request.toJson(),
      );

  Future<Result<void>> deleteFeedback(String planId, String feedbackId) =>
      _apiClient.delete('/training-plans/$planId/feedback/$feedbackId');
}

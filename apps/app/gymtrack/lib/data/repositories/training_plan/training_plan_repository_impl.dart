import '../../../domain/exercise_log.dart';
import '../../../domain/feedback.dart';
import '../../../domain/subscription.dart';
import '../../../domain/training_plan.dart';
import '../../services/api/api_client.dart';
import '../../services/training_plan_api_service.dart';
import 'training_plan_repository.dart';

class TrainingPlanRepositoryImpl implements TrainingPlanRepository {
  TrainingPlanRepositoryImpl({
    required TrainingPlanApiService trainingPlanApiService,
  }) : _trainingPlanApiService = trainingPlanApiService;

  final TrainingPlanApiService _trainingPlanApiService;

  @override
  Future<void> createPlan(TrainingPlan plan) async {
    final result = await _trainingPlanApiService.createPlan(plan);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> createPlanForStudent(String studentId, TrainingPlan plan) async {
    final result = await _trainingPlanApiService.createPlanForStudent(
      studentId,
      plan,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> updatePlan(String id, TrainingPlan plan) async {
    final result = await _trainingPlanApiService.updatePlan(id, plan);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deletePlan(String id) async {
    final result = await _trainingPlanApiService.deletePlan(id);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> createDay(String planId) async {
    final result = await _trainingPlanApiService.createDay(planId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deleteDay(String planId, String dayId) async {
    final result = await _trainingPlanApiService.deleteDay(planId, dayId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> createExercise(String planId, String dayId) async {
    final result = await _trainingPlanApiService.createExercise(planId, dayId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deleteExercise(
    String planId,
    String dayId,
    String exerciseId,
  ) async {
    final result = await _trainingPlanApiService.deleteExercise(
      planId,
      dayId,
      exerciseId,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> logExercise(
    String planId,
    String dayId,
    String exerciseId,
    ExerciseLogRequest request,
  ) async {
    final result = await _trainingPlanApiService.logExercise(
      planId,
      dayId,
      exerciseId,
      request,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> subscribe(String planId, SubscribeRequest request) async {
    final result = await _trainingPlanApiService.subscribe(planId, request);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> unsubscribe(String planId) async {
    final result = await _trainingPlanApiService.unsubscribe(planId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> changeSubscriptionStatus(
    String planId,
    ChangeSubscriptionStatusRequest request,
  ) async {
    final result = await _trainingPlanApiService.changeSubscriptionStatus(
      planId,
      request,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> changeSubscriptionPrivacy(
    String planId,
    ChangeSubscriptionPrivacyRequest request,
  ) async {
    final result = await _trainingPlanApiService.changeSubscriptionPrivacy(
      planId,
      request,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> completeDay(String subscriptionId, String dayId) async {
    final result = await _trainingPlanApiService.completeDay(
      subscriptionId,
      dayId,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> cancelDay(String subscriptionId, String dayId) async {
    final result = await _trainingPlanApiService.cancelDay(
      subscriptionId,
      dayId,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> startDay(String subscriptionId, String dayId) async {
    final result = await _trainingPlanApiService.startDay(
      subscriptionId,
      dayId,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> addFeedback(String planId, AddFeedbackRequest request) async {
    final result = await _trainingPlanApiService.addFeedback(planId, request);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deleteFeedback(String planId, String feedbackId) async {
    final result = await _trainingPlanApiService.deleteFeedback(
      planId,
      feedbackId,
    );
    if (result case Failure<void>()) throw Exception(result.message);
  }
}

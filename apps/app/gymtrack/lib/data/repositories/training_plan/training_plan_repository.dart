import '../../../domain/exercise_log.dart';
import '../../../domain/feedback.dart';
import '../../../domain/subscription.dart';
import '../../../domain/training_plan.dart';

abstract class TrainingPlanRepository {
  Future<void> createPlan(TrainingPlan plan);
  Future<void> createPlanForStudent(String studentId, TrainingPlan plan);
  Future<void> updatePlan(String id, TrainingPlan plan);
  Future<void> deletePlan(String id);
  Future<void> createDay(String planId);
  Future<void> deleteDay(String planId, String dayId);
  Future<void> createExercise(String planId, String dayId);
  Future<void> deleteExercise(String planId, String dayId, String exerciseId);
  Future<void> logExercise(
    String planId,
    String dayId,
    String exerciseId,
    ExerciseLogRequest request,
  );
  Future<void> subscribe(String planId, SubscribeRequest request);
  Future<void> unsubscribe(String planId);
  Future<void> changeSubscriptionStatus(
    String planId,
    ChangeSubscriptionStatusRequest request,
  );
  Future<void> changeSubscriptionPrivacy(
    String planId,
    ChangeSubscriptionPrivacyRequest request,
  );
  Future<void> completeDay(String subscriptionId, String dayId);
  Future<void> cancelDay(String subscriptionId, String dayId);
  Future<void> startDay(String subscriptionId, String dayId);
  Future<void> addFeedback(String planId, AddFeedbackRequest request);
  Future<void> deleteFeedback(String planId, String feedbackId);
}

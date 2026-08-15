import '../../services/api/api_client.dart';
import '../../services/trainer_api_service.dart';
import 'trainer_repository.dart';

class TrainerRepositoryImpl implements TrainerRepository {
  TrainerRepositoryImpl({required TrainerApiService trainerApiService})
    : _trainerApiService = trainerApiService;

  final TrainerApiService _trainerApiService;

  @override
  Future<void> createTrainerCode(String code) async {
    final result = await _trainerApiService.createTrainerCode(code);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> linkTrainer(String code) async {
    final result = await _trainerApiService.linkTrainer(code);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> unlinkTrainer() async {
    final result = await _trainerApiService.unlinkTrainer();
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> unlinkStudent(String studentId) async {
    final result = await _trainerApiService.unlinkStudent(studentId);
    if (result case Failure<void>()) throw Exception(result.message);
  }
}

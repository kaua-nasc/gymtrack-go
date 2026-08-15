import 'api/api_client.dart';

class TrainerApiService {
  TrainerApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<void>> createTrainerCode(String code) => _apiClient.patch(
    '/identity/users/trainers/profile/code',
    body: {'code': code},
  );

  Future<Result<void>> linkTrainer(String code) => _apiClient.post(
    '/identity/users/trainers/profile/link',
    body: {'code': code},
  );

  Future<Result<void>> unlinkTrainer() =>
      _apiClient.post('/identity/users/trainers/profile/unlink');

  Future<Result<void>> unlinkStudent(String studentId) => _apiClient.post(
    '/identity/users/trainers/students/$studentId/profile/unlink',
  );
}

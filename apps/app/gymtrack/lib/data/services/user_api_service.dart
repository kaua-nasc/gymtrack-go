import '../../domain/auth_tokens.dart';
import '../../domain/follower.dart';
import '../../domain/privacy_settings.dart';
import '../../domain/user.dart';
import 'api/api_client.dart';

class UserApiService {
  UserApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<User>> getMe() => _apiClient.get(
    '/identity/users/me',
    fromJson: (json) => User.fromJson(json),
  );

  Future<Result<User>> getUser(String id) => _apiClient.get(
    '/identity/users/$id',
    fromJson: (json) => User.fromJson(json),
  );

  Future<Result<List<User>>> listUsers(List<String> ids) => _apiClient.post(
    '/identity/users',
    body: ids,
    fromJson: (json) => (json as List)
        .map((e) => User.fromJson(e as Map<String, dynamic>))
        .toList(),
  );

  Future<Result<void>> updateProfile(String id, UpdateProfileRequest request) =>
      _apiClient.put('/identity/users/$id', body: request.toJson());

  Future<Result<void>> changePassword(ChangePasswordRequest request) =>
      _apiClient.put(
        '/identity/users/profile/password',
        body: request.toJson(),
      );

  Future<Result<UserPrivacySettings>> getPrivacySettings() => _apiClient.get(
    '/identity/users/profile/privacy',
    fromJson: (json) => UserPrivacySettings.fromJson(json),
  );

  Future<Result<void>> updatePrivacySettings(UserPrivacySettings settings) =>
      _apiClient.put(
        '/identity/users/profile/privacy',
        body: settings.toJson(),
      );

  Future<Result<void>> changeToTrainer(String cref) =>
      _apiClient.post('/identity/users/profile/upgrade', body: {'cref': cref});

  Future<Result<void>> changeToClient() =>
      _apiClient.post('/identity/users/profile/downgrade');

  Future<Result<FollowerResponse>> searchUsers(String query) => _apiClient.get(
    '/identity/users/search',
    queryParams: {'q': query},
    fromJson: (json) => FollowerResponse.fromJson(json),
  );
}

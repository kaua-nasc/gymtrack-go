import '../../domain/follower.dart';
import 'api/api_client.dart';

class FollowersApiService {
  FollowersApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<void>> followUser(String userId) =>
      _apiClient.post('/identity/users/$userId/follows');

  Future<Result<void>> unfollowUser(String userId) =>
      _apiClient.post('/identity/users/$userId/unfollows');

  Future<Result<CountResponse>> countFollowers(String userId) => _apiClient.get(
    '/identity/users/$userId/followers/count',
    fromJson: (json) => CountResponse.fromJson(json),
  );

  Future<Result<CountResponse>> countFollowing(String userId) => _apiClient.get(
    '/identity/users/$userId/following/count',
    fromJson: (json) => CountResponse.fromJson(json),
  );
}

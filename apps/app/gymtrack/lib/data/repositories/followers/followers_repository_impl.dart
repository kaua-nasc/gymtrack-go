import '../../../domain/follower.dart';
import '../../services/api/api_client.dart';
import '../../services/followers_api_service.dart';
import 'followers_repository.dart';

class FollowersRepositoryImpl implements FollowersRepository {
  FollowersRepositoryImpl({required FollowersApiService followersApiService})
    : _followersApiService = followersApiService;

  final FollowersApiService _followersApiService;

  @override
  Future<void> followUser(String userId) async {
    final result = await _followersApiService.followUser(userId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> unfollowUser(String userId) async {
    final result = await _followersApiService.unfollowUser(userId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<CountResponse> countFollowers(String userId) async {
    final result = await _followersApiService.countFollowers(userId);
    return switch (result) {
      Success<CountResponse>() => result.data,
      Failure<CountResponse>() => throw Exception(result.message),
    };
  }

  @override
  Future<CountResponse> countFollowing(String userId) async {
    final result = await _followersApiService.countFollowing(userId);
    return switch (result) {
      Success<CountResponse>() => result.data,
      Failure<CountResponse>() => throw Exception(result.message),
    };
  }
}

import '../../../domain/follower.dart';

abstract class FollowersRepository {
  Future<void> followUser(String userId);
  Future<void> unfollowUser(String userId);
  Future<CountResponse> countFollowers(String userId);
  Future<CountResponse> countFollowing(String userId);
}

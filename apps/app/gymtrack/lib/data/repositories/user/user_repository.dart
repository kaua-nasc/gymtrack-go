import 'package:flutter/foundation.dart';

import '../../../domain/follower.dart';
import '../../../domain/privacy_settings.dart';
import '../../../domain/user.dart';

abstract class UserRepository extends ChangeNotifier {
  User? get currentUser;
  Future<User> getCurrentUser();
  Future<User> getUser(String id);
  Future<List<User>> listUsers(List<String> ids);
  Future<User> updateProfile(UpdateProfileRequest request);
  Future<void> changePassword(String currentPassword, String newPassword);
  Future<UserPrivacySettings> getPrivacySettings();
  Future<UserPrivacySettings> updatePrivacySettings(
    UserPrivacySettings settings,
  );
  Future<void> changeToTrainer(String cref);
  Future<void> changeToClient();
  Future<FollowerResponse> searchUsers(String query);
}

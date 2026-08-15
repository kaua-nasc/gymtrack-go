import '../../../domain/auth_tokens.dart';
import '../../../domain/follower.dart';
import '../../../domain/privacy_settings.dart';
import '../../../domain/user.dart';
import '../../services/api/api_client.dart';
import '../../services/user_api_service.dart';
import 'user_repository.dart';

class UserRepositoryImpl extends UserRepository {
  UserRepositoryImpl({required UserApiService userApiService})
    : _userApiService = userApiService;

  final UserApiService _userApiService;

  User? _currentUser;

  @override
  User? get currentUser => _currentUser;

  @override
  Future<User> getCurrentUser() async {
    if (_currentUser != null) return _currentUser!;
    final result = await _userApiService.getMe();
    return switch (result) {
      Success<User>() => _currentUser = result.data,
      Failure<User>() => throw Exception(result.message),
    };
  }

  @override
  Future<User> getUser(String id) async {
    final result = await _userApiService.getUser(id);
    return switch (result) {
      Success<User>() => result.data,
      Failure<User>() => throw Exception(result.message),
    };
  }

  @override
  Future<List<User>> listUsers(List<String> ids) async {
    final result = await _userApiService.listUsers(ids);
    return switch (result) {
      Success<List<User>>() => result.data,
      Failure<List<User>>() => throw Exception(result.message),
    };
  }

  @override
  Future<User> updateProfile(UpdateProfileRequest request) async {
    final id = _currentUser?.id;
    if (id == null) throw Exception('User not loaded');
    final result = await _userApiService.updateProfile(id, request);
    switch (result) {
      case Success<void>():
        _currentUser = await getCurrentUser();
        notifyListeners();
        return _currentUser!;
      case Failure<void>():
        throw Exception(result.message);
    }
  }

  @override
  Future<void> changePassword(
    String currentPassword,
    String newPassword,
  ) async {
    final request = ChangePasswordRequest(
      currentPassword: currentPassword,
      newPassword: newPassword,
    );
    final result = await _userApiService.changePassword(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<UserPrivacySettings> getPrivacySettings() async {
    final result = await _userApiService.getPrivacySettings();
    return switch (result) {
      Success<UserPrivacySettings>() => result.data,
      Failure<UserPrivacySettings>() => throw Exception(result.message),
    };
  }

  @override
  Future<UserPrivacySettings> updatePrivacySettings(
    UserPrivacySettings settings,
  ) async {
    final result = await _userApiService.updatePrivacySettings(settings);
    return switch (result) {
      Success<void>() => settings,
      Failure<void>() => throw Exception(result.message),
    };
  }

  @override
  Future<void> changeToTrainer(String cref) async {
    final result = await _userApiService.changeToTrainer(cref);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<void> changeToClient() async {
    final result = await _userApiService.changeToClient();
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<FollowerResponse> searchUsers(String query) async {
    final result = await _userApiService.searchUsers(query);
    return switch (result) {
      Success<FollowerResponse>() => result.data,
      Failure<FollowerResponse>() => throw Exception(result.message),
    };
  }
}

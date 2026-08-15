import '../../../domain/auth_tokens.dart';
import '../../services/api/api_client.dart';
import '../../services/api/token_storage.dart';
import '../../services/auth_api_service.dart';
import 'auth_repository.dart';

class AuthRepositoryImpl extends AuthRepository {
  AuthRepositoryImpl({
    required AuthApiService authApiService,
    required TokenStorage tokenStorage,
  }) : _authApiService = authApiService,
       _tokenStorage = tokenStorage;

  final AuthApiService _authApiService;
  final TokenStorage _tokenStorage;

  bool _isAuthenticated = false;
  bool _isInitialized = false;

  @override
  bool get isAuthenticated => _isAuthenticated;

  @override
  bool get isInitialized => _isInitialized;

  @override
  Future<void> initialize() async {
    final token = await _tokenStorage.getAccessToken();
    _isAuthenticated = token != null;
    _isInitialized = true;
    notifyListeners();
  }

  @override
  Future<void> login(String email, String password) async {
    final request = LoginRequest(email: email, password: password);
    final result = await _authApiService.login(request);
    switch (result) {
      case Success<LoginResponse>():
        await _tokenStorage.saveAccessToken(result.data.accessToken);
        _isAuthenticated = true;
        notifyListeners();
      case Failure<LoginResponse>():
        throw Exception(result.message);
    }
  }

  @override
  Future<void> register(RegisterRequest request) async {
    final result = await _authApiService.register(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<void> logout() async {
    await _tokenStorage.clearAccessToken();
    _isAuthenticated = false;
    notifyListeners();
  }

  @override
  Future<void> sendVerificationEmail(String email) async {
    final request = SendVerificationRequest(email: email);
    final result = await _authApiService.sendVerificationEmail(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<void> verifyEmail(String email, String code) async {
    final request = VerifyEmailRequest(email: email, code: code);
    final result = await _authApiService.verifyEmail(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<void> sendResetPasswordToken(String email) async {
    final request = ResetPasswordSendTokenRequest(email: email);
    final result = await _authApiService.sendResetPasswordToken(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }

  @override
  Future<bool> verifyResetPasswordToken(String email, String code) async {
    final request = ResetPasswordVerifyTokenRequest(email: email, code: code);
    final result = await _authApiService.verifyResetPasswordToken(request);
    return switch (result) {
      Success<ResetPasswordVerifyTokenResponse>() => result.data.valid,
      Failure<ResetPasswordVerifyTokenResponse>() => throw Exception(
        result.message,
      ),
    };
  }

  @override
  Future<void> resetPassword(String email, String code, String password) async {
    final request = ResetPasswordRequest(
      email: email,
      code: code,
      password: password,
    );
    final result = await _authApiService.resetPassword(request);
    if (result case Failure<void>()) {
      throw Exception(result.message);
    }
  }
}

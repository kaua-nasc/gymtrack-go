import '../../domain/auth_tokens.dart';
import 'api/api_client.dart';

class AuthApiService {
  AuthApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<LoginResponse>> login(LoginRequest request) => _apiClient.post(
    '/identity/auth/login',
    body: request.toJson(),
    fromJson: (json) => LoginResponse.fromJson(json),
  );

  Future<Result<LoginResponse>> adminLogin(LoginRequest request) =>
      _apiClient.post(
        '/identity/auth/admin/login',
        body: request.toJson(),
        fromJson: (json) => LoginResponse.fromJson(json),
      );

  Future<Result<void>> register(RegisterRequest request) =>
      _apiClient.post('/identity/auth/register', body: request.toJson());

  Future<Result<void>> sendVerificationEmail(SendVerificationRequest request) =>
      _apiClient.post(
        '/identity/auth/verify/send-token',
        body: request.toJson(),
      );

  Future<Result<void>> verifyEmail(VerifyEmailRequest request) =>
      _apiClient.post('/identity/auth/verify', body: request.toJson());

  Future<Result<void>> sendResetPasswordToken(
    ResetPasswordSendTokenRequest request,
  ) => _apiClient.post(
    '/identity/auth/reset-password/send-token',
    body: request.toJson(),
  );

  Future<Result<ResetPasswordVerifyTokenResponse>> verifyResetPasswordToken(
    ResetPasswordVerifyTokenRequest request,
  ) => _apiClient.post(
    '/identity/auth/reset-password/verify-token',
    body: request.toJson(),
    fromJson: (json) => ResetPasswordVerifyTokenResponse.fromJson(json),
  );

  Future<Result<void>> resetPassword(ResetPasswordRequest request) =>
      _apiClient.post('/identity/auth/reset-password', body: request.toJson());
}

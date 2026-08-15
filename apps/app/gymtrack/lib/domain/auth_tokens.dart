import 'package:freezed_annotation/freezed_annotation.dart';

part 'auth_tokens.freezed.dart';
part 'auth_tokens.g.dart';

@freezed
abstract class LoginRequest with _$LoginRequest {
  const factory LoginRequest({
    required String email,
    required String password,
  }) = _LoginRequest;

  factory LoginRequest.fromJson(Map<String, dynamic> json) =>
      _$LoginRequestFromJson(json);
}

@freezed
abstract class LoginResponse with _$LoginResponse {
  const factory LoginResponse({
    @JsonKey(name: 'accessToken') required String accessToken,
  }) = _LoginResponse;

  factory LoginResponse.fromJson(Map<String, dynamic> json) =>
      _$LoginResponseFromJson(json);
}

@freezed
abstract class RegisterRequest with _$RegisterRequest {
  const factory RegisterRequest({
    @JsonKey(name: 'firstName') required String firstName,
    @JsonKey(name: 'lastName') required String lastName,
    required String email,
    required String password,
  }) = _RegisterRequest;

  factory RegisterRequest.fromJson(Map<String, dynamic> json) =>
      _$RegisterRequestFromJson(json);
}

@freezed
abstract class SendVerificationRequest with _$SendVerificationRequest {
  const factory SendVerificationRequest({required String email}) =
      _SendVerificationRequest;

  factory SendVerificationRequest.fromJson(Map<String, dynamic> json) =>
      _$SendVerificationRequestFromJson(json);
}

@freezed
abstract class VerifyEmailRequest with _$VerifyEmailRequest {
  const factory VerifyEmailRequest({
    required String email,
    required String code,
  }) = _VerifyEmailRequest;

  factory VerifyEmailRequest.fromJson(Map<String, dynamic> json) =>
      _$VerifyEmailRequestFromJson(json);
}

@freezed
abstract class ResetPasswordSendTokenRequest
    with _$ResetPasswordSendTokenRequest {
  const factory ResetPasswordSendTokenRequest({required String email}) =
      _ResetPasswordSendTokenRequest;

  factory ResetPasswordSendTokenRequest.fromJson(Map<String, dynamic> json) =>
      _$ResetPasswordSendTokenRequestFromJson(json);
}

@freezed
abstract class ResetPasswordVerifyTokenRequest
    with _$ResetPasswordVerifyTokenRequest {
  const factory ResetPasswordVerifyTokenRequest({
    required String email,
    required String code,
  }) = _ResetPasswordVerifyTokenRequest;

  factory ResetPasswordVerifyTokenRequest.fromJson(Map<String, dynamic> json) =>
      _$ResetPasswordVerifyTokenRequestFromJson(json);
}

@freezed
abstract class ResetPasswordVerifyTokenResponse
    with _$ResetPasswordVerifyTokenResponse {
  const factory ResetPasswordVerifyTokenResponse({required bool valid}) =
      _ResetPasswordVerifyTokenResponse;

  factory ResetPasswordVerifyTokenResponse.fromJson(
    Map<String, dynamic> json,
  ) => _$ResetPasswordVerifyTokenResponseFromJson(json);
}

@freezed
abstract class ResetPasswordRequest with _$ResetPasswordRequest {
  const factory ResetPasswordRequest({
    required String email,
    required String code,
    required String password,
  }) = _ResetPasswordRequest;

  factory ResetPasswordRequest.fromJson(Map<String, dynamic> json) =>
      _$ResetPasswordRequestFromJson(json);
}

@freezed
abstract class ChangePasswordRequest with _$ChangePasswordRequest {
  const factory ChangePasswordRequest({
    @JsonKey(name: 'currentPassword') required String currentPassword,
    @JsonKey(name: 'newPassword') required String newPassword,
  }) = _ChangePasswordRequest;

  factory ChangePasswordRequest.fromJson(Map<String, dynamic> json) =>
      _$ChangePasswordRequestFromJson(json);
}

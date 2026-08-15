// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'auth_tokens.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_LoginRequest _$LoginRequestFromJson(Map<String, dynamic> json) =>
    _LoginRequest(
      email: json['email'] as String,
      password: json['password'] as String,
    );

Map<String, dynamic> _$LoginRequestToJson(_LoginRequest instance) =>
    <String, dynamic>{'email': instance.email, 'password': instance.password};

_LoginResponse _$LoginResponseFromJson(Map<String, dynamic> json) =>
    _LoginResponse(accessToken: json['accessToken'] as String);

Map<String, dynamic> _$LoginResponseToJson(_LoginResponse instance) =>
    <String, dynamic>{'accessToken': instance.accessToken};

_RegisterRequest _$RegisterRequestFromJson(Map<String, dynamic> json) =>
    _RegisterRequest(
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
      email: json['email'] as String,
      password: json['password'] as String,
    );

Map<String, dynamic> _$RegisterRequestToJson(_RegisterRequest instance) =>
    <String, dynamic>{
      'firstName': instance.firstName,
      'lastName': instance.lastName,
      'email': instance.email,
      'password': instance.password,
    };

_SendVerificationRequest _$SendVerificationRequestFromJson(
  Map<String, dynamic> json,
) => _SendVerificationRequest(email: json['email'] as String);

Map<String, dynamic> _$SendVerificationRequestToJson(
  _SendVerificationRequest instance,
) => <String, dynamic>{'email': instance.email};

_VerifyEmailRequest _$VerifyEmailRequestFromJson(Map<String, dynamic> json) =>
    _VerifyEmailRequest(
      email: json['email'] as String,
      code: json['code'] as String,
    );

Map<String, dynamic> _$VerifyEmailRequestToJson(_VerifyEmailRequest instance) =>
    <String, dynamic>{'email': instance.email, 'code': instance.code};

_ResetPasswordSendTokenRequest _$ResetPasswordSendTokenRequestFromJson(
  Map<String, dynamic> json,
) => _ResetPasswordSendTokenRequest(email: json['email'] as String);

Map<String, dynamic> _$ResetPasswordSendTokenRequestToJson(
  _ResetPasswordSendTokenRequest instance,
) => <String, dynamic>{'email': instance.email};

_ResetPasswordVerifyTokenRequest _$ResetPasswordVerifyTokenRequestFromJson(
  Map<String, dynamic> json,
) => _ResetPasswordVerifyTokenRequest(
  email: json['email'] as String,
  code: json['code'] as String,
);

Map<String, dynamic> _$ResetPasswordVerifyTokenRequestToJson(
  _ResetPasswordVerifyTokenRequest instance,
) => <String, dynamic>{'email': instance.email, 'code': instance.code};

_ResetPasswordVerifyTokenResponse _$ResetPasswordVerifyTokenResponseFromJson(
  Map<String, dynamic> json,
) => _ResetPasswordVerifyTokenResponse(valid: json['valid'] as bool);

Map<String, dynamic> _$ResetPasswordVerifyTokenResponseToJson(
  _ResetPasswordVerifyTokenResponse instance,
) => <String, dynamic>{'valid': instance.valid};

_ResetPasswordRequest _$ResetPasswordRequestFromJson(
  Map<String, dynamic> json,
) => _ResetPasswordRequest(
  email: json['email'] as String,
  code: json['code'] as String,
  password: json['password'] as String,
);

Map<String, dynamic> _$ResetPasswordRequestToJson(
  _ResetPasswordRequest instance,
) => <String, dynamic>{
  'email': instance.email,
  'code': instance.code,
  'password': instance.password,
};

_ChangePasswordRequest _$ChangePasswordRequestFromJson(
  Map<String, dynamic> json,
) => _ChangePasswordRequest(
  currentPassword: json['currentPassword'] as String,
  newPassword: json['newPassword'] as String,
);

Map<String, dynamic> _$ChangePasswordRequestToJson(
  _ChangePasswordRequest instance,
) => <String, dynamic>{
  'currentPassword': instance.currentPassword,
  'newPassword': instance.newPassword,
};

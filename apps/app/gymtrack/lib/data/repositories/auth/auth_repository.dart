import 'package:flutter/foundation.dart';

import '../../../domain/auth_tokens.dart';

abstract class AuthRepository extends ChangeNotifier {
  bool get isAuthenticated;
  bool get isInitialized;
  Future<void> initialize();
  Future<void> login(String email, String password);
  Future<void> register(RegisterRequest request);
  Future<void> logout();
  Future<void> sendVerificationEmail(String email);
  Future<void> verifyEmail(String email, String code);
  Future<void> sendResetPasswordToken(String email);
  Future<bool> verifyResetPasswordToken(String email, String code);
  Future<void> resetPassword(String email, String code, String password);
}

import 'package:flutter/foundation.dart';

import '../../../data/repositories/auth/auth_repository.dart';

enum ForgotPasswordStep { email, code, reset }

class ForgotPasswordViewModel extends ChangeNotifier {
  ForgotPasswordViewModel({required AuthRepository authRepository})
    : _authRepository = authRepository;

  final AuthRepository _authRepository;

  ForgotPasswordStep currentStep = ForgotPasswordStep.email;

  String email = '';
  String code = '';
  String newPassword = '';
  String confirmPassword = '';

  bool isLoading = false;
  String? errorMessage;
  String? successMessage;
  bool obscurePassword = true;

  String? emailError;
  String? codeError;
  String? newPasswordError;
  String? confirmPasswordError;

  void reset() {
    currentStep = ForgotPasswordStep.email;
    email = '';
    code = '';
    newPassword = '';
    confirmPassword = '';
    isLoading = false;
    errorMessage = null;
    successMessage = null;
    emailError = null;
    codeError = null;
    newPasswordError = null;
    confirmPasswordError = null;
    notifyListeners();
  }

  void onEmailChanged(String value) {
    email = value;
    emailError = null;
  }

  void onCodeChanged(String value) {
    code = value;
    codeError = null;
  }

  void onNewPasswordChanged(String value) {
    newPassword = value;
    newPasswordError = null;
  }

  void onConfirmPasswordChanged(String value) {
    confirmPassword = value;
    confirmPasswordError = null;
  }

  void togglePasswordVisibility() {
    obscurePassword = !obscurePassword;
    notifyListeners();
  }

  bool _isValidEmail(String value) {
    return RegExp(r'^[^@]+@[^@]+\.[^@]+$').hasMatch(value);
  }

  bool _validateEmail() {
    emailError = null;
    if (email.isEmpty) {
      emailError = 'E-mail é obrigatório';
    } else if (!_isValidEmail(email)) {
      emailError = 'E-mail inválido';
    }
    return emailError == null;
  }

  bool _validateCode() {
    codeError = null;
    if (code.isEmpty) {
      codeError = 'Código é obrigatório';
    }
    return codeError == null;
  }

  bool _validateReset() {
    newPasswordError = null;
    confirmPasswordError = null;
    if (newPassword.isEmpty) {
      newPasswordError = 'Nova senha é obrigatória';
    } else if (newPassword.length < 6) {
      newPasswordError = 'Mínimo 6 caracteres';
    }
    if (confirmPassword.isEmpty) {
      confirmPasswordError = 'Confirmação é obrigatória';
    } else if (newPassword != confirmPassword) {
      confirmPasswordError = 'Senhas não conferem';
    }
    return newPasswordError == null && confirmPasswordError == null;
  }

  Future<void> sendToken() async {
    if (!_validateEmail()) {
      notifyListeners();
      return;
    }

    isLoading = true;
    errorMessage = null;
    notifyListeners();

    try {
      await _authRepository.sendResetPasswordToken(email);
      currentStep = ForgotPasswordStep.code;
    } on Exception catch (e) {
      errorMessage = e.toString().replaceFirst('Exception: ', '');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> verifyToken() async {
    if (!_validateCode()) {
      notifyListeners();
      return;
    }

    isLoading = true;
    errorMessage = null;
    notifyListeners();

    try {
      final valid = await _authRepository.verifyResetPasswordToken(email, code);
      if (valid) {
        currentStep = ForgotPasswordStep.reset;
      } else {
        errorMessage = 'Código inválido';
      }
    } on Exception catch (e) {
      errorMessage = e.toString().replaceFirst('Exception: ', '');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> resetPassword() async {
    if (!_validateReset()) {
      notifyListeners();
      return;
    }

    isLoading = true;
    errorMessage = null;
    notifyListeners();

    try {
      await _authRepository.resetPassword(email, code, newPassword);
      successMessage = 'Senha alterada com sucesso!';
    } on Exception catch (e) {
      errorMessage = e.toString().replaceFirst('Exception: ', '');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }
}

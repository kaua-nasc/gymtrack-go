import 'package:flutter/foundation.dart';

import '../../../data/repositories/auth/auth_repository.dart';

class LoginViewModel extends ChangeNotifier {
  LoginViewModel({required AuthRepository authRepository})
    : _authRepository = authRepository;

  final AuthRepository _authRepository;

  String email = '';
  String password = '';

  bool isLoading = false;
  String? errorMessage;
  bool obscurePassword = true;

  String? emailError;
  String? passwordError;

  void onEmailChanged(String value) {
    email = value;
    emailError = null;
  }

  void onPasswordChanged(String value) {
    password = value;
    passwordError = null;
  }

  void togglePasswordVisibility() {
    obscurePassword = !obscurePassword;
    notifyListeners();
  }

  bool _validate() {
    emailError = null;
    passwordError = null;

    if (email.isEmpty) {
      emailError = 'E-mail é obrigatório';
    } else if (!_isValidEmail(email)) {
      emailError = 'E-mail inválido';
    }

    if (password.isEmpty) {
      passwordError = 'Senha é obrigatória';
    }

    return emailError == null && passwordError == null;
  }

  bool _isValidEmail(String value) {
    return RegExp(r'^[^@]+@[^@]+\.[^@]+$').hasMatch(value);
  }

  Future<void> login() async {
    if (!_validate()) {
      notifyListeners();
      return;
    }

    isLoading = true;
    errorMessage = null;
    notifyListeners();

    try {
      await _authRepository.login(email, password);
    } on Exception catch (e) {
      errorMessage = _extractErrorMessage(e);
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  String _extractErrorMessage(Exception e) {
    final msg = e.toString().replaceFirst('Exception: ', '');
    return msg;
  }
}

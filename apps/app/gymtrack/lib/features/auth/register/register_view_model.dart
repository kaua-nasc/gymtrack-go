import 'package:flutter/foundation.dart';

import '../../../data/repositories/auth/auth_repository.dart';
import '../../../domain/auth_tokens.dart';

class RegisterViewModel extends ChangeNotifier {
  RegisterViewModel({required AuthRepository authRepository})
    : _authRepository = authRepository;

  final AuthRepository _authRepository;

  String firstName = '';
  String lastName = '';
  String email = '';
  String password = '';
  String confirmPassword = '';

  bool isLoading = false;
  String? errorMessage;
  bool obscurePassword = true;

  String? firstNameError;
  String? lastNameError;
  String? emailError;
  String? passwordError;
  String? confirmPasswordError;

  void onFirstNameChanged(String value) {
    firstName = value;
    firstNameError = null;
  }

  void onLastNameChanged(String value) {
    lastName = value;
    lastNameError = null;
  }

  void onEmailChanged(String value) {
    email = value;
    emailError = null;
  }

  void onPasswordChanged(String value) {
    password = value;
    passwordError = null;
  }

  void onConfirmPasswordChanged(String value) {
    confirmPassword = value;
    confirmPasswordError = null;
  }

  void togglePasswordVisibility() {
    obscurePassword = !obscurePassword;
    notifyListeners();
  }

  bool _validate() {
    firstNameError = null;
    lastNameError = null;
    emailError = null;
    passwordError = null;
    confirmPasswordError = null;

    if (firstName.isEmpty) {
      firstNameError = 'Nome é obrigatório';
    }
    if (lastName.isEmpty) {
      lastNameError = 'Sobrenome é obrigatório';
    }
    if (email.isEmpty) {
      emailError = 'E-mail é obrigatório';
    } else if (!_isValidEmail(email)) {
      emailError = 'E-mail inválido';
    }
    if (password.isEmpty) {
      passwordError = 'Senha é obrigatória';
    } else if (password.length < 6) {
      passwordError = 'Mínimo 6 caracteres';
    }
    if (confirmPassword.isEmpty) {
      confirmPasswordError = 'Confirmação é obrigatória';
    } else if (password != confirmPassword) {
      confirmPasswordError = 'Senhas não conferem';
    }

    return firstNameError == null &&
        lastNameError == null &&
        emailError == null &&
        passwordError == null &&
        confirmPasswordError == null;
  }

  bool _isValidEmail(String value) {
    return RegExp(r'^[^@]+@[^@]+\.[^@]+$').hasMatch(value);
  }

  Future<void> register() async {
    if (!_validate()) {
      notifyListeners();
      return;
    }

    isLoading = true;
    errorMessage = null;
    notifyListeners();

    try {
      final request = RegisterRequest(
        firstName: firstName,
        lastName: lastName,
        email: email,
        password: password,
      );
      await _authRepository.register(request);
    } on Exception catch (e) {
      errorMessage = e.toString().replaceFirst('Exception: ', '');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }
}

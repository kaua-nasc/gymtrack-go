import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

import 'package:gymtrack/data/repositories/auth/auth_repository.dart';
import 'package:gymtrack/features/auth/forgot_password/forgot_password_view_model.dart';

class _MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late _MockAuthRepository authRepository;
  late ForgotPasswordViewModel sut;

  setUp(() {
    authRepository = _MockAuthRepository();
    sut = ForgotPasswordViewModel(authRepository: authRepository);
  });

  group('sendToken', () {
    test('fails when email is empty', () async {
      sut.onEmailChanged('');

      await sut.sendToken();

      expect(sut.emailError, isNotNull);
      expect(sut.currentStep, ForgotPasswordStep.email);
    });

    test('advances to code step on success', () async {
      when(
        () => authRepository.sendResetPasswordToken(any()),
      ).thenAnswer((_) async {});

      sut.onEmailChanged('teste@email.com');
      await sut.sendToken();

      expect(sut.currentStep, ForgotPasswordStep.code);
      expect(sut.errorMessage, isNull);
    });

    test('sets error on failure', () async {
      when(
        () => authRepository.sendResetPasswordToken(any()),
      ).thenThrow(Exception('E-mail não encontrado'));

      sut.onEmailChanged('teste@email.com');
      await sut.sendToken();

      expect(sut.errorMessage, 'E-mail não encontrado');
      expect(sut.currentStep, ForgotPasswordStep.email);
    });
  });

  group('verifyToken', () {
    test('fails when code is empty', () async {
      sut.onCodeChanged('');

      await sut.verifyToken();

      expect(sut.codeError, isNotNull);
    });

    test('advances to reset step when valid', () async {
      sut.currentStep = ForgotPasswordStep.code;
      when(
        () => authRepository.verifyResetPasswordToken(any(), any()),
      ).thenAnswer((_) async => true);

      sut.onEmailChanged('teste@email.com');
      sut.onCodeChanged('123456');
      await sut.verifyToken();

      expect(sut.currentStep, ForgotPasswordStep.reset);
    });

    test('sets error when invalid', () async {
      sut.currentStep = ForgotPasswordStep.code;
      when(
        () => authRepository.verifyResetPasswordToken(any(), any()),
      ).thenAnswer((_) async => false);

      sut.onEmailChanged('teste@email.com');
      sut.onCodeChanged('000000');
      await sut.verifyToken();

      expect(sut.errorMessage, 'Código inválido');
      expect(sut.currentStep, ForgotPasswordStep.code);
    });
  });

  group('resetPassword', () {
    test('fails when passwords do not match', () async {
      sut.onNewPasswordChanged('123456');
      sut.onConfirmPasswordChanged('654321');

      await sut.resetPassword();

      expect(sut.confirmPasswordError, isNotNull);
    });

    test('sets success message on completion', () async {
      sut.currentStep = ForgotPasswordStep.reset;
      when(
        () => authRepository.resetPassword(any(), any(), any()),
      ).thenAnswer((_) async {});

      sut.onEmailChanged('teste@email.com');
      sut.onCodeChanged('123456');
      sut.onNewPasswordChanged('novaSenha123');
      sut.onConfirmPasswordChanged('novaSenha123');
      await sut.resetPassword();

      expect(sut.successMessage, 'Senha alterada com sucesso!');
    });

    test('sets error on failure', () async {
      sut.currentStep = ForgotPasswordStep.reset;
      when(
        () => authRepository.resetPassword(any(), any(), any()),
      ).thenThrow(Exception('Token expirado'));

      sut.onEmailChanged('teste@email.com');
      sut.onCodeChanged('123456');
      sut.onNewPasswordChanged('novaSenha123');
      sut.onConfirmPasswordChanged('novaSenha123');
      await sut.resetPassword();

      expect(sut.errorMessage, 'Token expirado');
    });
  });

  group('reset', () {
    test('clears all fields and goes to first step', () {
      sut.currentStep = ForgotPasswordStep.reset;
      sut.email = 'teste@email.com';
      sut.errorMessage = 'erro';

      sut.reset();

      expect(sut.currentStep, ForgotPasswordStep.email);
      expect(sut.email, isEmpty);
      expect(sut.errorMessage, isNull);
    });
  });
}

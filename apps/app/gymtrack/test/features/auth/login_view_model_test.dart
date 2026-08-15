import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

import 'package:gymtrack/data/repositories/auth/auth_repository.dart';
import 'package:gymtrack/features/auth/login/login_view_model.dart';

class _MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late _MockAuthRepository authRepository;
  late LoginViewModel sut;

  setUp(() {
    authRepository = _MockAuthRepository();
    sut = LoginViewModel(authRepository: authRepository);
  });

  group('validate', () {
    test('fails when email is empty', () {
      sut.onEmailChanged('');

      sut.login();

      expect(sut.emailError, isNotNull);
    });

    test('fails when email is invalid', () {
      sut.onEmailChanged('invalido');
      sut.onPasswordChanged('123456');

      sut.login();

      expect(sut.emailError, 'E-mail inválido');
    });

    test('fails when password is empty', () {
      sut.onEmailChanged('teste@email.com');
      sut.onPasswordChanged('');

      sut.login();

      expect(sut.passwordError, isNotNull);
    });

    test('clears errors when fields change', () {
      sut.onEmailChanged('');
      sut.login();
      expect(sut.emailError, isNotNull);

      sut.onEmailChanged('teste@email.com');
      expect(sut.emailError, isNull);
    });
  });

  group('login', () {
    test('calls repository on success', () async {
      when(() => authRepository.login(any(), any())).thenAnswer((_) async {});

      sut.onEmailChanged('teste@email.com');
      sut.onPasswordChanged('123456');
      await sut.login();

      verify(() => authRepository.login('teste@email.com', '123456')).called(1);
      expect(sut.isLoading, false);
      expect(sut.errorMessage, isNull);
    });

    test('sets error on failure', () async {
      when(
        () => authRepository.login(any(), any()),
      ).thenThrow(Exception('Credenciais inválidas'));

      sut.onEmailChanged('teste@email.com');
      sut.onPasswordChanged('123456');
      await sut.login();

      expect(sut.errorMessage, 'Credenciais inválidas');
      expect(sut.isLoading, false);
    });
  });

  group('togglePasswordVisibility', () {
    test('toggles obscurePassword', () {
      expect(sut.obscurePassword, true);

      sut.togglePasswordVisibility();
      expect(sut.obscurePassword, false);

      sut.togglePasswordVisibility();
      expect(sut.obscurePassword, true);
    });
  });
}

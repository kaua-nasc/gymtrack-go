import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

import 'package:gymtrack/data/repositories/auth/auth_repository.dart';
import 'package:gymtrack/domain/auth_tokens.dart';
import 'package:gymtrack/features/auth/register/register_view_model.dart';

class _MockAuthRepository extends Mock implements AuthRepository {}

class _RegisterRequestFake extends Fake implements RegisterRequest {}

void main() {
  late _MockAuthRepository authRepository;
  late RegisterViewModel sut;

  setUpAll(() {
    registerFallbackValue(_RegisterRequestFake());
  });

  setUp(() {
    authRepository = _MockAuthRepository();
    sut = RegisterViewModel(authRepository: authRepository);
  });

  group('validate', () {
    test('fails when required fields are empty', () {
      sut.register();

      expect(sut.firstNameError, isNotNull);
      expect(sut.lastNameError, isNotNull);
      expect(sut.emailError, isNotNull);
      expect(sut.passwordError, isNotNull);
      expect(sut.confirmPasswordError, isNotNull);
    });

    test('fails when password is too short', () {
      sut.onFirstNameChanged('João');
      sut.onLastNameChanged('Silva');
      sut.onEmailChanged('joao@email.com');
      sut.onPasswordChanged('123');
      sut.onConfirmPasswordChanged('123');

      sut.register();

      expect(sut.passwordError, 'Mínimo 6 caracteres');
    });

    test('fails when passwords do not match', () {
      sut.onFirstNameChanged('João');
      sut.onLastNameChanged('Silva');
      sut.onEmailChanged('joao@email.com');
      sut.onPasswordChanged('123456');
      sut.onConfirmPasswordChanged('654321');

      sut.register();

      expect(sut.confirmPasswordError, 'Senhas não conferem');
    });

    test('fails when email is invalid', () {
      sut.onFirstNameChanged('João');
      sut.onLastNameChanged('Silva');
      sut.onEmailChanged('invalido');
      sut.onPasswordChanged('123456');
      sut.onConfirmPasswordChanged('123456');

      sut.register();

      expect(sut.emailError, 'E-mail inválido');
    });
  });

  group('register', () {
    test('calls repository on success', () async {
      when(() => authRepository.register(any())).thenAnswer((_) async {});

      sut.onFirstNameChanged('João');
      sut.onLastNameChanged('Silva');
      sut.onEmailChanged('joao@email.com');
      sut.onPasswordChanged('123456');
      sut.onConfirmPasswordChanged('123456');
      await sut.register();

      verify(() => authRepository.register(any<RegisterRequest>())).called(1);
      expect(sut.isLoading, false);
      expect(sut.errorMessage, isNull);
    });

    test('sets error on failure', () async {
      when(
        () => authRepository.register(any()),
      ).thenThrow(Exception('E-mail já cadastrado'));

      sut.onFirstNameChanged('João');
      sut.onLastNameChanged('Silva');
      sut.onEmailChanged('joao@email.com');
      sut.onPasswordChanged('123456');
      sut.onConfirmPasswordChanged('123456');
      await sut.register();

      expect(sut.errorMessage, 'E-mail já cadastrado');
      expect(sut.isLoading, false);
    });
  });
}

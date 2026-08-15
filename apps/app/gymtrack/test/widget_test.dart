import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:provider/provider.dart';

import 'package:gymtrack/data/repositories/auth/auth_repository.dart';
import 'package:gymtrack/features/auth/login/login_screen.dart';

class _MockAuthRepository extends AuthRepository {
  @override
  bool get isAuthenticated => false;

  @override
  bool get isInitialized => true;

  @override
  Future<void> initialize() async {}

  @override
  Future<void> login(String email, String password) async {}

  @override
  Future<void> register(covariant dynamic request) async {}

  @override
  Future<void> logout() async {}

  @override
  Future<void> sendVerificationEmail(String email) async {}

  @override
  Future<void> verifyEmail(String email, String code) async {}

  @override
  Future<void> sendResetPasswordToken(String email) async {}

  @override
  Future<bool> verifyResetPasswordToken(String email, String code) async =>
      true;

  @override
  Future<void> resetPassword(
    String email,
    String code,
    String password,
  ) async {}
}

void main() {
  testWidgets('Login screen renders', (WidgetTester tester) async {
    await tester.pumpWidget(
      MaterialApp(
        home: ChangeNotifierProvider<AuthRepository>(
          create: (_) => _MockAuthRepository(),
          child: const LoginScreen(),
        ),
      ),
    );

    expect(find.text('Gymtrack'), findsOneWidget);
    expect(find.text('Entrar'), findsAtLeast(1));
    expect(find.text('E-mail'), findsOneWidget);
    expect(find.text('Senha'), findsOneWidget);
    expect(find.text('Esqueceu a senha?'), findsOneWidget);
    expect(find.text('Não tem conta?'), findsOneWidget);
  });

  testWidgets('Login screen shows validation errors', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(
      MaterialApp(
        home: ChangeNotifierProvider<AuthRepository>(
          create: (_) => _MockAuthRepository(),
          child: const LoginScreen(),
        ),
      ),
    );

    await tester.tap(find.widgetWithText(FilledButton, 'Entrar'));
    await tester.pump();

    expect(find.text('E-mail é obrigatório'), findsOneWidget);
    expect(find.text('Senha é obrigatória'), findsOneWidget);
  });
}

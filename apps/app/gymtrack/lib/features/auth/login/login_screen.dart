import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';

import '../widgets/auth_layout.dart';
import '../widgets/auth_text_field.dart';
import 'login_view_model.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => LoginViewModel(authRepository: context.read()),
      child: const _LoginBody(),
    );
  }
}

class _LoginBody extends StatelessWidget {
  const _LoginBody();

  @override
  Widget build(BuildContext context) {
    final vm = context.watch<LoginViewModel>();

    return AuthLayout(
      title: 'Entrar',
      footer: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text('Não tem conta?'),
          TextButton(
            onPressed: () => context.goNamed('register'),
            child: const Text('Cadastre-se'),
          ),
        ],
      ),
      children: [
        if (vm.errorMessage != null)
          Container(
            width: double.infinity,
            padding: const EdgeInsets.all(12),
            margin: const EdgeInsets.only(bottom: 16),
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.errorContainer,
              borderRadius: BorderRadius.circular(8),
            ),
            child: Text(
              vm.errorMessage!,
              style: TextStyle(
                color: Theme.of(context).colorScheme.onErrorContainer,
              ),
            ),
          ),
        AuthTextField(
          label: 'E-mail',
          onChanged: vm.onEmailChanged,
          keyboardType: TextInputType.emailAddress,
          textInputAction: TextInputAction.next,
          errorText: vm.emailError,
          autofocus: true,
        ),
        const SizedBox(height: 16),
        AuthTextField(
          label: 'Senha',
          onChanged: vm.onPasswordChanged,
          obscureText: vm.obscurePassword,
          textInputAction: TextInputAction.done,
          errorText: vm.passwordError,
          suffixIcon: IconButton(
            icon: Icon(
              vm.obscurePassword ? Icons.visibility_off : Icons.visibility,
            ),
            onPressed: vm.togglePasswordVisibility,
          ),
        ),
        const SizedBox(height: 8),
        Align(
          alignment: Alignment.centerRight,
          child: TextButton(
            onPressed: () => context.goNamed('forgotPassword'),
            child: const Text('Esqueceu a senha?'),
          ),
        ),
        const SizedBox(height: 16),
        SizedBox(
          width: double.infinity,
          height: 48,
          child: FilledButton(
            onPressed: vm.isLoading ? null : vm.login,
            child: vm.isLoading
                ? const SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(strokeWidth: 2),
                  )
                : const Text('Entrar'),
          ),
        ),
      ],
    );
  }
}

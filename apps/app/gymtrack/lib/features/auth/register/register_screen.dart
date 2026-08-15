import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';

import '../widgets/auth_layout.dart';
import '../widgets/auth_text_field.dart';
import 'register_view_model.dart';

class RegisterScreen extends StatelessWidget {
  const RegisterScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => RegisterViewModel(authRepository: context.read()),
      child: const _RegisterBody(),
    );
  }
}

class _RegisterBody extends StatelessWidget {
  const _RegisterBody();

  @override
  Widget build(BuildContext context) {
    final vm = context.watch<RegisterViewModel>();

    return AuthLayout(
      title: 'Criar Conta',
      footer: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text('Já tem conta?'),
          TextButton(
            onPressed: () => context.goNamed('login'),
            child: const Text('Entrar'),
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
          label: 'Nome',
          onChanged: vm.onFirstNameChanged,
          textInputAction: TextInputAction.next,
          errorText: vm.firstNameError,
          autofocus: true,
        ),
        const SizedBox(height: 16),
        AuthTextField(
          label: 'Sobrenome',
          onChanged: vm.onLastNameChanged,
          textInputAction: TextInputAction.next,
          errorText: vm.lastNameError,
        ),
        const SizedBox(height: 16),
        AuthTextField(
          label: 'E-mail',
          onChanged: vm.onEmailChanged,
          keyboardType: TextInputType.emailAddress,
          textInputAction: TextInputAction.next,
          errorText: vm.emailError,
        ),
        const SizedBox(height: 16),
        AuthTextField(
          label: 'Senha',
          onChanged: vm.onPasswordChanged,
          obscureText: vm.obscurePassword,
          textInputAction: TextInputAction.next,
          errorText: vm.passwordError,
          suffixIcon: IconButton(
            icon: Icon(
              vm.obscurePassword ? Icons.visibility_off : Icons.visibility,
            ),
            onPressed: vm.togglePasswordVisibility,
          ),
        ),
        const SizedBox(height: 16),
        AuthTextField(
          label: 'Confirmar Senha',
          onChanged: vm.onConfirmPasswordChanged,
          obscureText: vm.obscurePassword,
          textInputAction: TextInputAction.done,
          errorText: vm.confirmPasswordError,
        ),
        const SizedBox(height: 24),
        SizedBox(
          width: double.infinity,
          height: 48,
          child: FilledButton(
            onPressed: vm.isLoading ? null : vm.register,
            child: vm.isLoading
                ? const SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(strokeWidth: 2),
                  )
                : const Text('Cadastrar'),
          ),
        ),
      ],
    );
  }
}

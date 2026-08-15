import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';

import '../widgets/auth_layout.dart';
import '../widgets/auth_text_field.dart';
import 'forgot_password_view_model.dart';

class ForgotPasswordScreen extends StatelessWidget {
  const ForgotPasswordScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => ForgotPasswordViewModel(authRepository: context.read()),
      child: const _ForgotPasswordBody(),
    );
  }
}

class _ForgotPasswordBody extends StatelessWidget {
  const _ForgotPasswordBody();

  @override
  Widget build(BuildContext context) {
    final vm = context.watch<ForgotPasswordViewModel>();

    return AuthLayout(
      title: switch (vm.currentStep) {
        ForgotPasswordStep.email => 'Recuperar Senha',
        ForgotPasswordStep.code => 'Verificar Código',
        ForgotPasswordStep.reset => 'Nova Senha',
      },
      footer: TextButton(
        onPressed: () => context.goNamed('login'),
        child: const Text('Voltar ao login'),
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
        if (vm.successMessage != null)
          Container(
            width: double.infinity,
            padding: const EdgeInsets.all(12),
            margin: const EdgeInsets.only(bottom: 16),
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.primaryContainer,
              borderRadius: BorderRadius.circular(8),
            ),
            child: Text(
              vm.successMessage!,
              style: TextStyle(
                color: Theme.of(context).colorScheme.onPrimaryContainer,
              ),
            ),
          ),
        ..._buildStepContent(vm, context),
        const SizedBox(height: 24),
        SizedBox(
          width: double.infinity,
          height: 48,
          child: FilledButton(
            onPressed: vm.isLoading ? null : _actionForStep(vm),
            child: vm.isLoading
                ? const SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(strokeWidth: 2),
                  )
                : Text(_buttonLabel(vm.currentStep)),
          ),
        ),
      ],
    );
  }

  List<Widget> _buildStepContent(
    ForgotPasswordViewModel vm,
    BuildContext context,
  ) {
    return switch (vm.currentStep) {
      ForgotPasswordStep.email => [
        const Text(
          'Digite seu e-mail para receber um código de recuperação.',
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 24),
        AuthTextField(
          label: 'E-mail',
          onChanged: vm.onEmailChanged,
          keyboardType: TextInputType.emailAddress,
          textInputAction: TextInputAction.done,
          errorText: vm.emailError,
          autofocus: true,
        ),
      ],
      ForgotPasswordStep.code => [
        Text(
          'Enviamos um código para ${vm.email}',
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 24),
        AuthTextField(
          label: 'Código',
          onChanged: vm.onCodeChanged,
          keyboardType: TextInputType.number,
          textInputAction: TextInputAction.done,
          errorText: vm.codeError,
          autofocus: true,
        ),
      ],
      ForgotPasswordStep.reset => [
        const Text('Escolha sua nova senha.', textAlign: TextAlign.center),
        const SizedBox(height: 24),
        AuthTextField(
          label: 'Nova Senha',
          onChanged: vm.onNewPasswordChanged,
          obscureText: vm.obscurePassword,
          textInputAction: TextInputAction.next,
          errorText: vm.newPasswordError,
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
      ],
    };
  }

  VoidCallback? _actionForStep(ForgotPasswordViewModel vm) {
    return switch (vm.currentStep) {
      ForgotPasswordStep.email => vm.sendToken,
      ForgotPasswordStep.code => vm.verifyToken,
      ForgotPasswordStep.reset => vm.resetPassword,
    };
  }

  String _buttonLabel(ForgotPasswordStep step) {
    return switch (step) {
      ForgotPasswordStep.email => 'Enviar Código',
      ForgotPasswordStep.code => 'Verificar',
      ForgotPasswordStep.reset => 'Redefinir Senha',
    };
  }
}

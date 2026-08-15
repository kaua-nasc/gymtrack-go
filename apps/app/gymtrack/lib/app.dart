import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'bootstrap.dart';
import 'config/environment.dart';
import 'data/repositories/auth/auth_repository.dart';
import 'routing/app_router.dart';

class GymtrackApp extends StatelessWidget {
  const GymtrackApp({super.key, required this.environment});

  final Environment environment;

  @override
  Widget build(BuildContext context) {
    return Bootstrap(
      environment: environment,
      child: _AppBody(environment: environment),
    );
  }
}

class _AppBody extends StatelessWidget {
  const _AppBody({required this.environment});

  final Environment environment;

  @override
  Widget build(BuildContext context) {
    final authRepository = context.watch<AuthRepository>();
    final router = appRouter(authRepository);

    return MaterialApp.router(
      title: 'Gymtrack',
      debugShowCheckedModeBanner: environment == Environment.development,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.green,
          brightness: Brightness.light,
        ),
        useMaterial3: true,
      ),
      routerConfig: router,
    );
  }
}

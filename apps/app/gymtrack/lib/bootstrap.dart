import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'config/environment.dart';
import 'data/repositories/auth/auth_repository.dart';
import 'data/repositories/auth/auth_repository_impl.dart';
import 'data/repositories/dashboard/dashboard_repository.dart';
import 'data/repositories/dashboard/dashboard_repository_impl.dart';
import 'data/repositories/followers/followers_repository.dart';
import 'data/repositories/followers/followers_repository_impl.dart';
import 'data/repositories/metrics/metrics_repository.dart';
import 'data/repositories/metrics/metrics_repository_impl.dart';
import 'data/repositories/social/social_repository.dart';
import 'data/repositories/social/social_repository_impl.dart';
import 'data/repositories/trainer/trainer_repository.dart';
import 'data/repositories/trainer/trainer_repository_impl.dart';
import 'data/repositories/training_plan/training_plan_repository.dart';
import 'data/repositories/training_plan/training_plan_repository_impl.dart';
import 'data/repositories/user/user_repository.dart';
import 'data/repositories/user/user_repository_impl.dart';
import 'data/services/api/api_client.dart';
import 'data/services/api/token_storage.dart';
import 'data/services/auth_api_service.dart';
import 'data/services/dashboard_api_service.dart';
import 'data/services/followers_api_service.dart';
import 'data/services/metrics_api_service.dart';
import 'data/services/social_api_service.dart';
import 'data/services/trainer_api_service.dart';
import 'data/services/training_plan_api_service.dart';
import 'data/services/user_api_service.dart';

class Bootstrap extends StatelessWidget {
  const Bootstrap({super.key, required this.environment, required this.child});

  final Environment environment;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider<Environment>.value(value: environment),
        Provider<TokenStorage>(create: (_) => TokenStorage()),
        Provider<ApiClient>(
          create: (context) => ApiClient(
            tokenStorage: context.read<TokenStorage>(),
            environment: context.read<Environment>(),
          ),
        ),
        Provider<AuthApiService>(
          create: (context) =>
              AuthApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<UserApiService>(
          create: (context) =>
              UserApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<TrainingPlanApiService>(
          create: (context) =>
              TrainingPlanApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<SocialApiService>(
          create: (context) =>
              SocialApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<MetricsApiService>(
          create: (context) =>
              MetricsApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<FollowersApiService>(
          create: (context) =>
              FollowersApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<DashboardApiService>(
          create: (context) =>
              DashboardApiService(apiClient: context.read<ApiClient>()),
        ),
        Provider<TrainerApiService>(
          create: (context) =>
              TrainerApiService(apiClient: context.read<ApiClient>()),
        ),
        ChangeNotifierProvider<AuthRepository>(
          create: (context) => AuthRepositoryImpl(
            authApiService: context.read<AuthApiService>(),
            tokenStorage: context.read<TokenStorage>(),
          ),
        ),
        ChangeNotifierProvider<UserRepository>(
          create: (context) => UserRepositoryImpl(
            userApiService: context.read<UserApiService>(),
          ),
        ),
        Provider<TrainingPlanRepository>(
          create: (context) => TrainingPlanRepositoryImpl(
            trainingPlanApiService: context.read<TrainingPlanApiService>(),
          ),
        ),
        Provider<SocialRepository>(
          create: (context) => SocialRepositoryImpl(
            socialApiService: context.read<SocialApiService>(),
          ),
        ),
        Provider<MetricsRepository>(
          create: (context) => MetricsRepositoryImpl(
            metricsApiService: context.read<MetricsApiService>(),
          ),
        ),
        Provider<FollowersRepository>(
          create: (context) => FollowersRepositoryImpl(
            followersApiService: context.read<FollowersApiService>(),
          ),
        ),
        Provider<DashboardRepository>(
          create: (context) => DashboardRepositoryImpl(
            dashboardApiService: context.read<DashboardApiService>(),
          ),
        ),
        Provider<TrainerRepository>(
          create: (context) => TrainerRepositoryImpl(
            trainerApiService: context.read<TrainerApiService>(),
          ),
        ),
      ],
      child: child,
    );
  }
}

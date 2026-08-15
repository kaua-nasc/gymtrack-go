import 'package:go_router/go_router.dart';

import '../data/repositories/auth/auth_repository.dart';
import '../features/auth/forgot_password/forgot_password_screen.dart';
import '../features/auth/login/login_screen.dart';
import '../features/auth/register/register_screen.dart';
import '../features/home/home_screen.dart';

GoRouter appRouter(AuthRepository authRepository) {
  return GoRouter(
    refreshListenable: authRepository,
    initialLocation: '/',
    debugLogDiagnostics: true,
    redirect: (context, state) {
      final isLoggedIn = authRepository.isAuthenticated;
      final isOnAuth = state.matchedLocation.startsWith('/auth');
      if (!isLoggedIn && !isOnAuth) return '/auth/login';
      if (isLoggedIn && isOnAuth) return '/';
      return null;
    },
    routes: [
      GoRoute(
        path: '/auth/login',
        name: 'login',
        builder: (context, state) => const LoginScreen(),
      ),
      GoRoute(
        path: '/auth/register',
        name: 'register',
        builder: (context, state) => const RegisterScreen(),
      ),
      GoRoute(
        path: '/auth/forgot-password',
        name: 'forgotPassword',
        builder: (context, state) => const ForgotPasswordScreen(),
      ),
      GoRoute(
        path: '/',
        name: 'home',
        builder: (context, state) => const HomeScreen(),
      ),
    ],
  );
}

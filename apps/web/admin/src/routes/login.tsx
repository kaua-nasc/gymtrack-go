import { createFileRoute, redirect } from '@tanstack/react-router';
import { LoginForm } from '../features/auth/LoginForm';
import { useAuthStore } from '../store/auth';

export const Route = createFileRoute('/login')({
  beforeLoad: () => {
    if (useAuthStore.getState().token) {
      throw redirect({
        to: '/posts/pending',
      });
    }
  },
  component: LoginComponent,
});

function LoginComponent() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-4 bg-bg">
      <div className="mb-8 text-center">
        <h1 className="text-4xl font-black tracking-tight text-primary">GYMTRACK</h1>
        <p className="text-text-muted font-medium">Portal de Administração</p>
      </div>
      <LoginForm />
    </div>
  );
}

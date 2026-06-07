import { useForm } from '@tanstack/react-form';
import { useLogin } from '../../api/queries/useLogin';

export function LoginForm() {
  const loginMutation = useLogin();

  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      loginMutation.mutate(value);
    },
  });

  return (
    <div className="w-full max-w-md p-8 bg-surface rounded-xl shadow-lg border border-border">
      <h2 className="text-2xl font-bold text-center mb-6 text-text-strong">Acesso Administrativo</h2>

      <form
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          form.handleSubmit();
        }}
        className="space-y-4"
      >
        <div>
          <form.Field
            name="email"
            validators={{
              onChange: ({ value }) =>
                !value ? 'Email é obrigatório' :
                  !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value) ? 'Email inválido' : undefined,
            }}
            children={(field) => (
              <>
                <label className="block text-sm font-medium mb-1 text-text-muted">Email</label>
                <input
                  value={field.state.value}
                  onBlur={field.handleBlur}
                  onChange={(e) => field.handleChange(e.target.value)}
                  type="email"
                  className="w-full px-4 py-2 rounded-lg border border-border bg-input text-text-strong focus:ring-2 focus:ring-primary outline-none transition-all placeholder:text-placeholder"
                  placeholder="admin@gymtrack.com"
                />
                {field.state.meta.errors ? (
                  <em className="text-error text-xs mt-1">{field.state.meta.errors.join(', ')}</em>
                ) : null}
              </>
            )}
          />
        </div>

        <div>
          <form.Field
            name="password"
            validators={{
              onChange: ({ value }) => !value ? 'Senha é obrigatória' : value.length < 8 ? 'Mínimo 8 caracteres' : undefined,
            }}
            children={(field) => (
              <>
                <label className="block text-sm font-medium mb-1 text-text-muted">Senha</label>
                <input
                  value={field.state.value}
                  onBlur={field.handleBlur}
                  onChange={(e) => field.handleChange(e.target.value)}
                  type="password"
                  className="w-full px-4 py-2 rounded-lg border border-border bg-input text-text-strong focus:ring-2 focus:ring-primary outline-none transition-all placeholder:text-placeholder"
                  placeholder="••••••••"
                />
                {field.state.meta.errors ? (
                  <em className="text-error text-xs mt-1">{field.state.meta.errors.join(', ')}</em>
                ) : null}
              </>
            )}
          />
        </div>

        {loginMutation.isError && (
          <div className="p-3 text-sm bg-error/10 text-error rounded-lg">
            {loginMutation.error instanceof Error ? loginMutation.error.message : 'Falha ao entrar'}
          </div>
        )}

        <button
          type="submit"
          disabled={loginMutation.isPending}
          className="w-full py-2 px-4 bg-primary hover:bg-primary-hover text-white font-semibold rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {loginMutation.isPending ? 'Entrando...' : 'Entrar'}
        </button>
      </form>
    </div>
  );
}

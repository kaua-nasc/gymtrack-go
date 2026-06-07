import { useMutation } from '@tanstack/react-query';
import { apiFetch } from '../client';
import { useAuthStore } from '../../store/auth';
import { useNavigate } from '@tanstack/react-router';
import { decodeJwtPayload } from '../../lib/jwt';
import type { UserResponse } from './useProfile';

interface LoginResponse {
  accessToken: string;
}

interface LoginCredentials {
  email: string;
  password: string;
}

export function useLogin() {
  const setToken = useAuthStore((state) => state.setToken);
  const setUser = useAuthStore((state) => state.setUser);
  const navigate = useNavigate();

  return useMutation({
    mutationFn: (credentials: LoginCredentials) =>
      apiFetch<LoginResponse>('/identity/auth/admin/login', {
        method: 'POST',
        body: JSON.stringify(credentials),
      }),
    onSuccess: async (data) => {
      setToken(data.accessToken);

      const payload = decodeJwtPayload(data.accessToken);
      if (payload?.sub) {
        try {
          const profile = await apiFetch<UserResponse>(`/identity/users/${payload.sub}`);
          setUser({
            id: profile.id,
            firstName: profile.firstName,
            lastName: profile.lastName,
            email: profile.email || '',
            profilePictureUrl: profile.profilePictureUrl,
            type: profile.type,
          });
        } catch {
          // Continua para a página mesmo se falhar buscar o perfil
        }
      }

      navigate({ to: '/posts/pending' });
    },
  });
}

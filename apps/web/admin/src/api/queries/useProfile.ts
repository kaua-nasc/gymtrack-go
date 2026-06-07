import { useQuery } from '@tanstack/react-query';
import { apiFetch } from '../client';
import { useAuthStore } from '../../store/auth';
import { decodeJwtPayload } from '../../lib/jwt';

export interface UserResponse {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  profilePictureUrl?: string;
  type: string;
}

export function useProfile() {
  const token = useAuthStore((state) => state.token);
  const setUser = useAuthStore((state) => state.setUser);

  const userId = token ? decodeJwtPayload(token)?.sub : null;

  return useQuery({
    queryKey: ['profile', userId],
    queryFn: async () => {
      const data = await apiFetch<UserResponse>(`/identity/users/${userId}`);
      setUser({
        id: data.id,
        firstName: data.firstName,
        lastName: data.lastName,
        email: data.email || '',
        profilePictureUrl: data.profilePictureUrl,
        type: data.type,
      });
      return data;
    },
    enabled: !!userId,
    retry: false,
    staleTime: 5 * 60 * 1000,
  });
}

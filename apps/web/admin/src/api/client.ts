import { useAuthStore } from '../store/auth';

const IDENTITY_URL = import.meta.env.VITE_IDENTITY_API_URL ?? '__VITE_IDENTITY_API_URL__';
const SOCIAL_URL = import.meta.env.VITE_SOCIAL_API_URL ?? '__VITE_SOCIAL_API_URL__';

export function resolveBaseUrl(endpoint: string): string {
  if (endpoint.startsWith('/identity')) return IDENTITY_URL;
  if (endpoint.startsWith('/social')) return SOCIAL_URL;
  return IDENTITY_URL;
}

export async function apiFetch<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = useAuthStore.getState().token;
  const baseUrl = resolveBaseUrl(endpoint);

  const headers = new Headers(options.headers);
  if (token) {
    headers.set('Authorization', `Bearer ${token}`);
  }
  if (options.body && !headers.has('Content-Type')) {
    headers.set('Content-Type', 'application/json');
  }

  const response = await fetch(`${baseUrl}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
  }

  if (response.status === 204) {
    return {} as T;
  }

  return response.json();
}

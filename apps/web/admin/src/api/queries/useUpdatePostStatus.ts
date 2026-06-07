import { useMutation, useQueryClient } from '@tanstack/react-query'
import { apiFetch } from '../client'

export function useUpdatePostStatus() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ postId, status, reason }: { postId: string; status: 'APPROVED' | 'REJECTED'; reason?: string }) =>
      apiFetch(`/social/admin/posts/${postId}/status`, {
        method: 'PATCH',
        body: JSON.stringify({ status, reason }),
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts', 'pending'] })
    },
  })
}

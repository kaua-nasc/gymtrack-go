import { useMutation, useQueryClient } from '@tanstack/react-query'
import { apiFetch } from '../client'

export function useUpdatePostStatus() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ postId, status }: { postId: string; status: 'APPROVED' | 'REJECTED' }) =>
      apiFetch(`/social/admin/posts/${postId}/status`, {
        method: 'PATCH',
        body: JSON.stringify({ status }),
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts', 'pending'] })
    },
  })
}

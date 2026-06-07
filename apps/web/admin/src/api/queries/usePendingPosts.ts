import { useQuery } from '@tanstack/react-query'
import { apiFetch } from '../client'

export interface Post {
  id: string
  createdAt: string
  updatedAt: string
  authorId: string
  content: string
  mediaUrls: string[]
  status: 'PENDING' | 'APPROVED' | 'REJECTED'
  entityId?: string
  entityType?: 'TRAINING_PLAN'
  trainingPlan?: {
    id: string
    name: string
    timeInDays: number
    type: string
    level: string
    imageUrl?: string
    description?: string
  }
  author?: {
    firstName: string
    lastName: string
    profilePictureUrl?: string
  }
}

interface PendingPostsResponse {
  data: Post[]
  nextCursor?: string
}

export function usePendingPosts(cursor?: string) {
  return useQuery({
    queryKey: ['posts', 'pending', cursor],
    queryFn: () => {
      const url = `/social/admin/posts/pending${cursor ? `?cursor=${cursor}` : ''}`
      return apiFetch<PendingPostsResponse>(url)
    },
  })
}

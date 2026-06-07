import { useQuery } from '@tanstack/react-query'
import { apiFetch } from '../client'

export interface AuditedPost {
  id: string
  createdAt: string
  updatedAt: string
  authorId: string
  content: string
  mediaUrls: string[]
  status: 'APPROVED' | 'REJECTED'
  rejectedReason?: string
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

interface AuditHistoryResponse {
  data: AuditedPost[]
  nextCursor?: string
}

interface AuditFilters {
  status?: string
  startDate?: string
  endDate?: string
}

export function useAuditHistory(filters: AuditFilters, cursor?: string) {
  const params = new URLSearchParams()
  if (cursor) params.set('cursor', cursor)
  if (filters.status) params.set('status', filters.status)
  if (filters.startDate) params.set('startDate', filters.startDate)
  if (filters.endDate) params.set('endDate', filters.endDate)

  const qs = params.toString()

  return useQuery({
    queryKey: ['posts', 'audit-history', filters, cursor],
    queryFn: () => {
      const url = `/social/admin/posts/history${qs ? `?${qs}` : ''}`
      return apiFetch<AuditHistoryResponse>(url)
    },
  })
}

import { useState } from 'react'
import type { Post } from '../../api/queries/usePendingPosts'
import { useUpdatePostStatus } from '../../api/queries/useUpdatePostStatus'

interface PostCardProps {
  post: Post
}

const videoExts = new Set(['mp4', 'webm', 'mov', 'avi', 'mkv', 'm4v'])

function isVideoUrl(url: string): boolean {
  const ext = url.split('?')[0].split('.').pop()?.toLowerCase()
  return ext ? videoExts.has(ext) : false
}

function MediaItem({ url, className }: { url: string; className?: string }) {
  if (isVideoUrl(url)) {
    return (
      <video
        src={url}
        controls
        playsInline
        preload="metadata"
        className={className}
      />
    )
  }
  return <img src={url} alt="Mídia do post" className={className} />
}

const levelLabel: Record<string, string> = {
  BEGINNER: 'Iniciante',
  INTERMEDIATE: 'Intermediário',
  ADVANCED: 'Avançado',
}

const typeLabel: Record<string, string> = {
  WEIGHT_LOSS: 'Emagrecimento',
  HYPERTROPHY: 'Hipertrofia',
  FUNCTIONAL: 'Funcional',
  CARDIO: 'Cardio',
  GENERAL: 'Geral',
}

export function PostCard({ post }: PostCardProps) {
  const [showRejectForm, setShowRejectForm] = useState(false)
  const [rejectReason, setRejectReason] = useState('')
  const statusMutation = useUpdatePostStatus()

  const handleApprove = () => {
    statusMutation.mutate({ postId: post.id, status: 'APPROVED' })
  }

  const handleReject = () => {
    if (!rejectReason.trim()) return
    statusMutation.mutate({ postId: post.id, status: 'REJECTED', reason: rejectReason.trim() })
  }

  const cancelReject = () => {
    setShowRejectForm(false)
    setRejectReason('')
  }

  const authorName = post.author
    ? `${post.author.firstName} ${post.author.lastName}`
    : 'Autor Desconhecido'

  const dateStr = new Date(post.createdAt).toLocaleDateString('pt-BR', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })

  return (
    <article className="bg-surface border border-border rounded-2xl overflow-hidden shadow-sm transition-all hover:shadow-md">
      <div className="p-6">
        <header className="flex items-center gap-4 mb-4">
          <div className="w-12 h-12 rounded-full bg-card flex items-center justify-center font-bold text-text-muted overflow-hidden">
            {post.author?.profilePictureUrl ? (
              <img src={post.author.profilePictureUrl} alt={authorName} className="w-full h-full object-cover" />
            ) : (
              authorName.charAt(0)
            )}
          </div>
          <div>
            <h4 className="font-bold text-text-strong">{authorName}</h4>
            <time className="text-xs text-text-muted font-medium">{dateStr}</time>
          </div>
        </header>

        <div className="space-y-4">
          <p className="text-text-gray whitespace-pre-wrap leading-relaxed">
            {post.content}
          </p>

          {post.trainingPlan && (
            <div className="bg-card border border-border rounded-xl overflow-hidden">
              <div className="flex gap-4">
                {post.trainingPlan.imageUrl && (
                  <div className="w-28 shrink-0">
                    <img
                      src={post.trainingPlan.imageUrl}
                      alt={post.trainingPlan.name}
                      className="w-full h-full object-cover"
                    />
                  </div>
                )}
                <div className="flex-1 p-4">
                  <div className="flex items-center gap-2 mb-1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="text-primary"><path d="M6 2L3 6v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V6l-3-4H6z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 0 1-8 0"/></svg>
                    <span className="text-xs uppercase tracking-wider font-bold text-primary">Plano de Treino</span>
                  </div>
                  <h5 className="text-lg font-bold text-text-strong">
                    {post.trainingPlan.name}
                  </h5>
                  <div className="flex flex-wrap gap-2 mt-2">
                    {post.trainingPlan.level && (
                      <span className="text-[11px] px-2 py-0.5 rounded-full bg-primary/10 text-primary font-semibold">
                        {levelLabel[post.trainingPlan.level] || post.trainingPlan.level}
                      </span>
                    )}
                    {post.trainingPlan.type && (
                      <span className="text-[11px] px-2 py-0.5 rounded-full bg-border text-text-muted font-semibold">
                        {typeLabel[post.trainingPlan.type] || post.trainingPlan.type}
                      </span>
                    )}
                    <span className="text-[11px] px-2 py-0.5 rounded-full bg-border text-text-muted font-semibold">
                      {post.trainingPlan.timeInDays} dias
                    </span>
                  </div>
                </div>
              </div>
            </div>
          )}

          {post.mediaUrls && post.mediaUrls.length > 0 && (() => {
            const count = post.mediaUrls.length
            const base = "rounded-xl overflow-hidden bg-black/20"

            if (count === 1) {
              return (
                <MediaItem
                  url={post.mediaUrls[0]}
                  className={`${base} w-full max-h-[600px] object-contain`}
                />
              )
            }

            if (count === 3) {
              return (
                <div className="grid grid-cols-2 gap-2">
                  <div className="row-span-2">
                    <MediaItem
                      url={post.mediaUrls[0]}
                      className={`${base} w-full h-full object-cover`}
                    />
                  </div>
                  <MediaItem
                    url={post.mediaUrls[1]}
                    className={`${base} w-full aspect-square object-cover`}
                  />
                  <MediaItem
                    url={post.mediaUrls[2]}
                    className={`${base} w-full aspect-square object-cover`}
                  />
                </div>
              )
            }

            return (
              <div className="grid grid-cols-2 gap-2">
                {post.mediaUrls.map((url, i) => (
                  <MediaItem
                    key={i}
                    url={url}
                    className={`${base} w-full aspect-square object-cover`}
                  />
                ))}
              </div>
            )
          })()}
        </div>
      </div>

      <footer className="bg-card/50 p-4 border-t border-border">
        {showRejectForm ? (
          <div className="space-y-3">
            <div>
              <label className="block text-sm font-medium text-text-muted mb-1">
                Motivo da rejeição
              </label>
              <textarea
                value={rejectReason}
                onChange={(e) => setRejectReason(e.target.value)}
                placeholder="Explique ao autor por que o post foi rejeitado..."
                rows={3}
                maxLength={1024}
                className="w-full px-4 py-2 rounded-xl border border-border bg-input text-text-strong text-sm focus:ring-2 focus:ring-primary outline-none transition-all resize-none placeholder:text-placeholder"
                autoFocus
              />
              <div className="flex justify-end mt-1">
                <span className={`text-xs font-medium ${rejectReason.length >= 1024 ? 'text-error' : 'text-text-muted'}`}>
                  {rejectReason.length}/1024
                </span>
              </div>
            </div>
            <div className="flex items-center justify-end gap-3">
              <button
                onClick={cancelReject}
                disabled={statusMutation.isPending}
                className="px-6 py-2 rounded-xl font-bold text-text-muted hover:bg-surface-hover transition-colors disabled:opacity-50"
              >
                Cancelar
              </button>
              <button
                onClick={handleReject}
                disabled={statusMutation.isPending || !rejectReason.trim()}
                className="px-6 py-2 rounded-xl font-bold bg-error text-white hover:opacity-90 shadow-lg shadow-error/20 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {statusMutation.isPending ? 'Rejeitando...' : 'Confirmar Rejeição'}
              </button>
            </div>
          </div>
        ) : (
          <div className="flex items-center justify-end gap-3">
            <button
              onClick={() => setShowRejectForm(true)}
              disabled={statusMutation.isPending}
              className="px-6 py-2 rounded-xl font-bold text-error hover:bg-error/10 transition-colors disabled:opacity-50"
            >
              Rejeitar
            </button>
            <button
              onClick={handleApprove}
              disabled={statusMutation.isPending}
              className="px-6 py-2 rounded-xl font-bold bg-primary text-white hover:bg-primary-hover shadow-lg shadow-primary/20 transition-all disabled:opacity-50"
            >
              {statusMutation.isPending ? 'Aprovando...' : 'Aprovar'}
            </button>
          </div>
        )}
      </footer>
    </article>
  )
}

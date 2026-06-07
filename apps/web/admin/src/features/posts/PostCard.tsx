import type { Post } from '../../api/queries/usePendingPosts'
import { useUpdatePostStatus } from '../../api/queries/useUpdatePostStatus'

interface PostCardProps {
  post: Post
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
  const statusMutation = useUpdatePostStatus()

  const handleAction = (status: 'APPROVED' | 'REJECTED') => {
    statusMutation.mutate({ postId: post.id, status })
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
    <article className="bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl overflow-hidden shadow-sm transition-all hover:shadow-md">
      <div className="p-6">
        <header className="flex items-center gap-4 mb-4">
          <div className="w-12 h-12 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center font-bold text-slate-400 overflow-hidden">
            {post.author?.profilePictureUrl ? (
              <img src={post.author.profilePictureUrl} alt={authorName} className="w-full h-full object-cover" />
            ) : (
              authorName.charAt(0)
            )}
          </div>
          <div>
            <h4 className="font-bold text-slate-900 dark:text-slate-50">{authorName}</h4>
            <time className="text-xs text-slate-500 font-medium">{dateStr}</time>
          </div>
        </header>

        <div className="space-y-4">
          <p className="text-slate-700 dark:text-slate-300 whitespace-pre-wrap leading-relaxed">
            {post.content}
          </p>

          {post.trainingPlan && (
            <div className="bg-slate-50 dark:bg-slate-800/40 border border-slate-200 dark:border-slate-700 rounded-xl overflow-hidden">
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
                  <h5 className="text-lg font-bold text-slate-900 dark:text-slate-50">
                    {post.trainingPlan.name}
                  </h5>
                  <div className="flex flex-wrap gap-2 mt-2">
                    {post.trainingPlan.level && (
                      <span className="text-[11px] px-2 py-0.5 rounded-full bg-primary/10 text-primary font-semibold">
                        {levelLabel[post.trainingPlan.level] || post.trainingPlan.level}
                      </span>
                    )}
                    {post.trainingPlan.type && (
                      <span className="text-[11px] px-2 py-0.5 rounded-full bg-slate-200 dark:bg-slate-700 text-slate-600 dark:text-slate-300 font-semibold">
                        {typeLabel[post.trainingPlan.type] || post.trainingPlan.type}
                      </span>
                    )}
                    <span className="text-[11px] px-2 py-0.5 rounded-full bg-slate-200 dark:bg-slate-700 text-slate-600 dark:text-slate-300 font-semibold">
                      {post.trainingPlan.timeInDays} dias
                    </span>
                  </div>
                </div>
              </div>
            </div>
          )}

          {post.mediaUrls && post.mediaUrls.length > 0 && (
            <div className={`grid gap-2 ${post.mediaUrls.length > 1 ? 'grid-cols-2' : 'grid-cols-1'}`}>
              {post.mediaUrls.map((url, i) => (
                <img
                  key={i}
                  src={url}
                  alt="Mídia do post"
                  className="rounded-xl w-full h-64 object-cover bg-slate-100 dark:bg-slate-800"
                />
              ))}
            </div>
          )}
        </div>
      </div>

      <footer className="bg-slate-50/50 dark:bg-slate-800/30 p-4 border-t border-slate-200 dark:border-slate-800 flex items-center justify-end gap-3">
        <button
          onClick={() => handleAction('REJECTED')}
          disabled={statusMutation.isPending}
          className="px-6 py-2 rounded-xl font-bold text-red-500 hover:bg-red-50 dark:hover:bg-red-900/10 transition-colors disabled:opacity-50"
        >
          {statusMutation.isPending && statusMutation.variables?.status === 'REJECTED' ? 'Rejeitando...' : 'Rejeitar'}
        </button>
        <button
          onClick={() => handleAction('APPROVED')}
          disabled={statusMutation.isPending}
          className="px-6 py-2 rounded-xl font-bold bg-primary text-white hover:bg-primary/90 shadow-lg shadow-primary/20 transition-all disabled:opacity-50"
        >
          {statusMutation.isPending && statusMutation.variables?.status === 'APPROVED' ? 'Aprovando...' : 'Aprovar'}
        </button>
      </footer>
    </article>
  )
}

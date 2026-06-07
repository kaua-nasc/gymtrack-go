import type { AuditLog } from '../../api/queries/useAuditHistory'

interface AuditedPostCardProps {
  log: AuditLog
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

export function AuditedPostCard({ log }: AuditedPostCardProps) {
  const post = log.post
  const authorName = post?.author
    ? `${post.author.firstName} ${post.author.lastName}`
    : 'Autor Desconhecido'

  const adminName = log.admin
    ? `${log.admin.firstName} ${log.admin.lastName}`
    : 'Administrador Desconhecido'

  const dateStr = post?.createdAt
    ? new Date(post.createdAt).toLocaleDateString('pt-BR', {
        day: 'numeric',
        month: 'long',
        year: 'numeric',
      })
    : '-'

  const auditedDateStr = new Date(log.createdAt).toLocaleDateString('pt-BR', {
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
            {post?.author?.profilePictureUrl ? (
              <img src={post.author.profilePictureUrl} alt={authorName} className="w-full h-full object-cover" />
            ) : (
              authorName.charAt(0)
            )}
          </div>
          <div>
            <h4 className="font-bold text-text-strong">{authorName}</h4>
            <time className="text-xs text-text-muted font-medium">Criado em {dateStr}</time>
          </div>
        </header>

        <div className="space-y-4">
          <p className="text-text-gray whitespace-pre-wrap leading-relaxed">
            {post?.content}
          </p>

          {post?.trainingPlan && (
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
                  <h5 className="text-lg font-bold text-text-strong">{post.trainingPlan.name}</h5>
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

          {post?.mediaUrls && post.mediaUrls.length > 0 && (
            <div className="grid gap-2 grid-cols-2">
              {post.mediaUrls.map((url, i) => (
                <img
                  key={i}
                  src={url}
                  alt="Mídia do post"
                  className="rounded-xl w-full aspect-square object-cover bg-card"
                />
              ))}
            </div>
          )}
        </div>
      </div>

      <footer className="bg-card/50 p-4 border-t border-border">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <span
              className={`text-xs uppercase tracking-wider font-bold px-3 py-1 rounded-full ${
                log.newStatus === 'APPROVED'
                  ? 'bg-success/15 text-success'
                  : 'bg-error/15 text-error'
              }`}
            >
              {log.newStatus === 'APPROVED' ? 'Aprovado' : 'Rejeitado'}
            </span>
            <span className="text-xs text-text-muted">em {auditedDateStr}</span>
          </div>
          <div className="text-xs text-text-muted">
            por <span className="font-semibold text-text-gray">{adminName}</span>
          </div>
        </div>
        {log.newStatus === 'REJECTED' && log.reason && (
          <div className="mt-3 p-3 rounded-xl bg-error/5 border border-error/20">
            <p className="text-xs font-semibold text-error mb-1">Motivo da rejeição:</p>
            <p className="text-sm text-text-gray">{log.reason}</p>
          </div>
        )}
      </footer>
    </article>
  )
}
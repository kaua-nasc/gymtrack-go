import { createFileRoute, redirect } from '@tanstack/react-router'
import { useState } from 'react'
import { useAuthStore } from '../../store/auth'
import { Sidebar } from '../../components/Sidebar'
import { useAuditHistory } from '../../api/queries/useAuditHistory'
import { AuditedPostCard } from '../../features/posts/AuditedPostCard'

export const Route = createFileRoute('/posts/history')({
  beforeLoad: () => {
    if (!useAuthStore.getState().token) {
      throw redirect({
        to: '/login',
      })
    }
  },
  component: AuditHistoryPage,
})

function AuditHistoryPage() {
  const [statusFilter, setStatusFilter] = useState('')
  const [startDate, setStartDate] = useState('')
  const [endDate, setEndDate] = useState('')
  const [filterError, setFilterError] = useState('')
  const [appliedFilters, setAppliedFilters] = useState({ status: '', startDate: '', endDate: '' })
  const [cursor, setCursor] = useState<string | undefined>(undefined)
  const [allPosts, setAllPosts] = useState<any[]>([])

  const today = new Date().toISOString().split('T')[0]

  const filters = appliedFilters.status || appliedFilters.startDate || appliedFilters.endDate
    ? appliedFilters
    : {}

  const { data, isLoading, isError, error } = useAuditHistory(filters, cursor)

  const posts = data?.data || []
  const nextCursor = data?.nextCursor

  const clearError = () => { if (filterError) setFilterError('') }

  const handleFilter = () => {
    setFilterError('')

    if (startDate && startDate > today) {
      setFilterError('A data inicial não pode ser futura.')
      return
    }
    if (endDate && endDate > today) {
      setFilterError('A data final não pode ser futura.')
      return
    }
    if (startDate && endDate && startDate > endDate) {
      setFilterError('A data inicial não pode ser posterior à data final.')
      return
    }

    const start = startDate ? new Date(startDate).toISOString() : ''
    const end = endDate ? new Date(endDate + 'T23:59:59').toISOString() : ''
    setAppliedFilters({ status: statusFilter, startDate: start, endDate: end })
    setCursor(undefined)
    setAllPosts([])
  }

  const loadMore = () => {
    if (nextCursor) setCursor(nextCursor)
  }

  const canFilter = statusFilter || startDate || endDate

  return (
    <div className="flex h-screen bg-bg text-text-strong">
      <Sidebar />
      <main className="flex-1 overflow-hidden p-8 flex flex-col">
        <header className="mb-8 max-w-3xl mx-auto w-full">
          <h1 className="text-3xl font-bold tracking-tight">Histórico de Auditoria</h1>
          <p className="text-text-muted mt-2">Consulte posts aprovados e rejeitados em um período.</p>
        </header>

        <div className="max-w-3xl w-full mx-auto mb-8">
          <div className="bg-surface border border-border rounded-2xl p-6">
            <div className="grid grid-cols-4 gap-4">
              <div>
                <label className="block text-sm font-medium text-text-muted mb-1">Status</label>
                <select
                  value={statusFilter}
                  onChange={(e) => setStatusFilter(e.target.value)}
                  className="w-full px-3 py-2 rounded-xl border border-border bg-input text-text-strong text-sm focus:ring-2 focus:ring-primary outline-none"
                >
                  <option value="">Todos</option>
                  <option value="APPROVED">Aprovados</option>
                  <option value="REJECTED">Rejeitados</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-text-muted mb-1">Data inicial</label>
                <input
                  type="date"
                  value={startDate}
                  onChange={(e) => { setStartDate(e.target.value); clearError() }}
                  max={today}
                  className="w-full px-3 py-2 rounded-xl border border-border bg-input text-text-strong text-sm focus:ring-2 focus:ring-primary outline-none"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-text-muted mb-1">Data final</label>
                <input
                  type="date"
                  value={endDate}
                  onChange={(e) => { setEndDate(e.target.value); clearError() }}
                  max={today}
                  className="w-full px-3 py-2 rounded-xl border border-border bg-input text-text-strong text-sm focus:ring-2 focus:ring-primary outline-none"
                />
              </div>
              <div className="flex items-end">
                <button
                  onClick={handleFilter}
                  disabled={!canFilter}
                  className="w-full px-4 py-2 rounded-xl font-bold bg-primary text-white hover:bg-primary-hover transition-colors disabled:opacity-50"
                >
                  Filtrar
                </button>
              </div>
            </div>
          </div>
          {filterError && (
            <div className="mt-3 px-4 py-3 bg-error/10 border border-error/30 rounded-xl text-error text-sm font-medium">
              {filterError}
            </div>
          )}
        </div>

        <div className="max-w-3xl w-full mx-auto flex-1 overflow-y-auto space-y-6 pr-2">
          {isLoading && cursor === undefined && (
            <div className="space-y-6">
              {[1, 2, 3].map((i) => (
                <div key={i} className="bg-surface h-64 rounded-2xl animate-pulse border border-border" />
              ))}
            </div>
          )}

          {isError && (
            <div className="p-8 bg-error/10 border border-error/30 rounded-2xl text-error">
              <p className="font-bold">Erro ao carregar histórico</p>
              <p className="text-sm">{(error as Error).message}</p>
            </div>
          )}

          {!isLoading && !isError && posts.length === 0 && cursor === undefined && (
            <div className="flex flex-col items-center justify-center p-20 bg-surface border border-border rounded-3xl text-center">
              <div className="w-20 h-20 bg-text-muted/10 text-text-muted rounded-full flex items-center justify-center mb-6 text-3xl">
                📋
              </div>
              <h3 className="text-xl font-bold text-text-strong">Nenhum resultado</h3>
              <p className="text-text-muted mt-2">Nenhum post encontrado para os filtros selecionados.</p>
            </div>
          )}

          {posts.map((post) => (
            <AuditedPostCard key={post.id} post={post} />
          ))}

          {nextCursor && (
            <div className="flex justify-center pb-8">
              <button
                onClick={loadMore}
                disabled={isLoading}
                className="px-8 py-3 rounded-xl font-bold bg-surface border border-border text-text-muted hover:bg-surface-hover transition-colors disabled:opacity-50"
              >
                {isLoading ? 'Carregando...' : 'Carregar mais'}
              </button>
            </div>
          )}
        </div>
      </main>
    </div>
  )
}

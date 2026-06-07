import { createFileRoute, redirect } from '@tanstack/react-router'
import { useAuthStore } from '../../store/auth'
import { Sidebar } from '../../components/Sidebar'
import { PendingPostsList } from '../../features/posts/PendingPostsList'

export const Route = createFileRoute('/posts/pending')({
  beforeLoad: () => {
    if (!useAuthStore.getState().token) {
      throw redirect({
        to: '/login',
      })
    }
  },
  component: PendingPostsPage,
})

function PendingPostsPage() {
  return (
    <div className="flex h-screen bg-bg text-text-strong">
      <Sidebar />
      <main className="flex-1 overflow-hidden p-8 flex flex-col">
        <header className="mb-8 text-center max-w-3xl mx-auto w-full">
          <h1 className="text-3xl font-bold tracking-tight">Fila de Auditoria</h1>
          <p className="text-text-muted mt-2">Revise e gerencie os posts pendentes da rede social.</p>
        </header>
        
        <div className="max-w-3xl w-full mx-auto flex-1 overflow-hidden">
          <PendingPostsList />
        </div>
      </main>
    </div>
  )
}

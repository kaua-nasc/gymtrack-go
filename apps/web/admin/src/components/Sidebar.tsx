import { Link, useNavigate } from '@tanstack/react-router'
import { useAuthStore } from '../store/auth'

export function Sidebar() {
  const user = useAuthStore((state) => state.user)
  const logout = useAuthStore((state) => state.logout)
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate({ to: '/login' })
  }

  const initials = user
    ? `${user.firstName.charAt(0)}${user.lastName.charAt(0)}`
    : 'AD'

  const displayName = user
    ? `${user.firstName} ${user.lastName}`
    : 'Usuário Admin'

  return (
    <aside className="w-72 flex flex-col border-r border-border bg-surface shadow-sm">
      <div className="p-8">
        <h2 className="text-2xl font-black tracking-tighter text-primary">GYMTRACK</h2>
        <span className="text-[10px] uppercase tracking-widest font-bold text-text-muted">Painel Administrativo</span>
      </div>

      <nav className="flex-1 px-4 space-y-1">
        <Link
          to="/posts/pending"
          activeProps={{ className: 'bg-primary/10 text-primary' }}
          className="flex items-center gap-3 px-4 py-3 rounded-xl transition-all font-medium hover:bg-surface-hover text-text-gray"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="M9 21V9"/></svg>
          Posts Pendentes
        </Link>
        <Link
          to="/posts/history"
          activeProps={{ className: 'bg-primary/10 text-primary' }}
          className="flex items-center gap-3 px-4 py-3 rounded-xl transition-all font-medium hover:bg-surface-hover text-text-gray"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
          Histórico
        </Link>
      </nav>

      <div className="p-4 border-t border-border">
        <div className="flex items-center gap-3 px-4 py-3 rounded-2xl bg-card mb-4">
          <div className="w-10 h-10 rounded-full bg-primary/20 flex items-center justify-center text-primary font-bold shrink-0">
            {initials}
          </div>
          <div className="flex-1 min-w-0">
            <p className="text-sm font-semibold truncate text-text-strong">{displayName}</p>
            <p className="text-[10px] text-text-muted uppercase tracking-wider">Administrador</p>
          </div>
        </div>
        
        <button
          onClick={handleLogout}
          className="flex items-center gap-3 w-full px-4 py-3 rounded-xl text-error font-medium transition-all hover:bg-error/10"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" x2="9" y1="12" y2="12"/></svg>
          Sair
        </button>
      </div>
    </aside>
  )
}

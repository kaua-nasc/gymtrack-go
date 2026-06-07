import { usePendingPosts } from '../../api/queries/usePendingPosts'
import { PostCard } from './PostCard'
import { useVirtualizer } from '@tanstack/react-virtual'
import { useRef } from 'react'

export function PendingPostsList() {
  const { data, isLoading, isError, error } = usePendingPosts()
  const parentRef = useRef<HTMLDivElement>(null)

  const posts = data?.data || []

  const rowVirtualizer = useVirtualizer({
    count: posts.length,
    getScrollElement: () => parentRef.current,
    estimateSize: () => 400,
    overscan: 5,
  })

  if (isLoading) {
    return (
      <div className="space-y-6">
        {[1, 2, 3].map((i) => (
          <div key={i} className="bg-surface h-64 rounded-2xl animate-pulse border border-border" />
        ))}
      </div>
    )
  }

  if (isError) {
    return (
      <div className="p-8 bg-error/10 border border-error/30 rounded-2xl text-error">
        <p className="font-bold">Erro ao carregar posts</p>
        <p className="text-sm">{(error as Error).message}</p>
      </div>
    )
  }

  if (posts.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center p-20 bg-surface border border-border rounded-3xl text-center">
        <div className="w-20 h-20 bg-success/20 text-success rounded-full flex items-center justify-center mb-6 text-3xl">
          🎉
        </div>
        <h3 className="text-xl font-bold text-text-strong">Tudo em dia!</h3>
        <p className="text-text-muted mt-2">Nenhum post pendente para revisão no momento.</p>
      </div>
    )
  }

  return (
    <div
      ref={parentRef}
      className="h-[calc(100vh-200px)] overflow-y-auto scrollbar-hide pr-2"
    >
      <div
        className="relative w-full"
        style={{ height: `${rowVirtualizer.getTotalSize()}px` }}
      >
        {rowVirtualizer.getVirtualItems().map((virtualRow) => (
          <div
            key={virtualRow.key}
            data-index={virtualRow.index}
            ref={rowVirtualizer.measureElement}
            className="absolute top-0 left-0 w-full pb-8"
            style={{
              transform: `translateY(${virtualRow.start}px)`,
            }}
          >
            <PostCard post={posts[virtualRow.index]} />
          </div>
        ))}
      </div>
    </div>
  )
}

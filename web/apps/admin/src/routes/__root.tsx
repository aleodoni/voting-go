import { createRootRoute, Outlet } from '@tanstack/react-router'
import { useSSE } from '@voting/shared'
import { useQueryClient } from '@tanstack/react-query'

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  const queryClient = useQueryClient()

  useSSE({
    onConnect: () => {
  		console.log('[SSE] conectado')
  		queryClient.invalidateQueries({ queryKey: ['connected-users'] })
		},
    onEvent: (event) => {
			switch (event.type) {
				case 'votacao_fechada':
					queryClient.invalidateQueries({ queryKey: ['voting-stats'] })
      		break
				case 'votacao_cancelada':
					queryClient.invalidateQueries({ queryKey: ['voting-stats'] })
					break
			}
      console.log('[SSE]', event.type, event.payload)
    },
    onError: (e) => {
      console.error('[SSE] erro:', e)
    },
  })

  return (
    <div className="flex w-full h-full">
      <Outlet />
    </div>
  )
}
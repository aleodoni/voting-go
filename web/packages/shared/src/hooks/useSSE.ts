import { useEffect } from 'react'
import { getKeycloak } from '../keycloak'

export type SSEEventType =
  | 'votacao_aberta'
  | 'votacao_fechada'
  | 'votacao_cancelada'
  | 'voto_registrado'

export type SSEEvent = {
  type: SSEEventType
  payload?: unknown
}

type UseSSEOptions = {
  onConnect?: () => void
  onEvent: (event: SSEEvent) => void
  onError?: (error: Event) => void
}

const SSE_EVENTS: SSEEventType[] = [
  'votacao_aberta',
  'votacao_fechada',
  'votacao_cancelada',
  'voto_registrado',
]

export function useSSE({ onConnect, onEvent, onError }: UseSSEOptions) {
  useEffect(() => {
    const token = getKeycloak().token
    const url = `${import.meta.env.VITE_API_URL}/eventos?token=${token}`
    const eventSource = new EventSource(url)

    eventSource.onopen = () => {
      onConnect?.()
    }

    SSE_EVENTS.forEach((type) => {
      eventSource.addEventListener(type, (e: MessageEvent) => {
        try {
          const payload = JSON.parse(e.data)
          onEvent({ type, payload })
        } catch {
          onEvent({ type })
        }
      })
    })

    eventSource.onerror = (e) => {
      onError?.(e)
    }

    return () => {
      eventSource.close()
    }
  }, [])
}
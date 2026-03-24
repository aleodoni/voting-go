import { createContext, useContext, useEffect, useRef, useState, ReactNode } from 'react'
import { initKeycloak, getKeycloak } from './keycloak'
import { initApi } from './api-client'
import type { User } from './types'
import { Button } from './components'

interface AuthConfig {
  apiUrl: string
  keycloak: {
    url: string
    realm: string
    clientId: string
  }
  authorize: (user: User) => boolean
}

interface AuthContextValue {
  user: User | null
  logout: () => void
}

const AuthContext = createContext<AuthContextValue | null>(null)

export function AuthProvider({ config, children }: { config: AuthConfig; children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [ready, setReady] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const initialized = useRef(false)

  useEffect(() => {
    if (initialized.current) return
    initialized.current = true

    const keycloak = initKeycloak(config.keycloak)
    const api = initApi(config.apiUrl)

    keycloak
      .init({ onLoad: 'login-required', checkLoginIframe: false })
      .then(async (authenticated) => {
        if (!authenticated) {
          keycloak.login()
          return
        }

        try {
          const { data } = await api.get<User>('/me')

          if (!data.credencial.ativo) {
            setError('Usuário inativo. Entre em contato com o administrador.')
            return
          }

          if (!config.authorize(data)) {
            setError('Você não tem permissão para acessar este sistema.')
            return
          }

          setUser(data)
          setReady(true)
        } catch {
          setError('Erro ao buscar dados do usuário.')
        }
      })
      .catch(() => {
        setError('Erro ao inicializar autenticação.')
      })
  }, [])

  const logout = () => {
    getKeycloak().logout({ redirectUri: window.location.origin })
  }

  if (error) {
    return (
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100vh', flexDirection: 'column', gap: '1rem' }}>
        <p style={{ color: 'red' }}>{error}</p>
        <Button variant={'outline'} onClick={logout}>Sair</Button>
      </div>
    )
  }

  if (!ready) {
    return (
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
        <p>Carregando...</p>
      </div>
    )
  }

  return (
    <AuthContext.Provider value={{ user, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth(): AuthContextValue {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth deve ser usado dentro de AuthProvider')
  return ctx
}
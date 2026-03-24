import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { AuthProvider, Container, ContainerWrapper, ThemeProvider, ClientWrapper } from '@voting/shared'
import App from './App'
import './index.css' 

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
    <ClientWrapper>
    <AuthProvider
      config={{
        apiUrl: import.meta.env.VITE_API_URL,
        keycloak: {
          url: import.meta.env.VITE_KEYCLOAK_URL,
          realm: import.meta.env.VITE_KEYCLOAK_REALM,
          clientId: import.meta.env.VITE_KEYCLOAK_CLIENT_ID,
        },
        authorize: (user) => user.credencial.pode_votar,
      }}
    >
      <ContainerWrapper>
        <Container>
          <App />
        </Container>
      </ContainerWrapper>
    </AuthProvider>
    </ClientWrapper>
    </ThemeProvider>
  </StrictMode>,
)
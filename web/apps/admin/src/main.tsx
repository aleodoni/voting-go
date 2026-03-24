import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createRouter, RouterProvider } from '@tanstack/react-router';
import { AuthProvider, Container, ContainerWrapper, ThemeProvider, ClientWrapper } from '@voting/shared'
import { routeTree } from './routeTree.gen'
import './index.css' 

const router = createRouter({ routeTree })


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
        authorize: (user) => user.credencial.pode_administrar,
      }}
    >
      <ContainerWrapper>
        <Container>
          <RouterProvider router={router} />
        </Container>
      </ContainerWrapper>
    </AuthProvider>
    </ClientWrapper>
    </ThemeProvider>
  </StrictMode>,
)
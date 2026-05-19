import { createRouter, RouterProvider } from '@tanstack/react-router';
import {
	AuthProvider,
	ClientWrapper,
	Container,
	ContainerWrapper,
	ThemeProvider,
	User,
} from '@voting/shared';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { routeTree } from './routeTree.gen';
import './index.css';
import { Toaster } from 'react-hot-toast';

const router = createRouter({ routeTree });

const authConfig = {
	apiUrl: import.meta.env.VITE_API_URL,
	keycloak: {
		url: import.meta.env.VITE_KEYCLOAK_URL,
		realm: import.meta.env.VITE_KEYCLOAK_REALM,
		clientId: import.meta.env.VITE_KEYCLOAK_CLIENT_ID,
	},
	authorize: (user: User) => user.credencial.pode_votar,
};

createRoot(document.getElementById('root')!).render(
	<StrictMode>
		<ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
			<Toaster position="bottom-right" />
			<ClientWrapper>
				<App />
			</ClientWrapper>
		</ThemeProvider>
	</StrictMode>,
);

function App() {
	return (
		<AuthProvider config={authConfig}>
			<ContainerWrapper>
				<Container>
					<RouterProvider router={router} />
				</Container>
			</ContainerWrapper>
		</AuthProvider>
	);
}

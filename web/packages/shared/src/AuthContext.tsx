import {
	createContext,
	ReactNode,
	useContext,
	useEffect,
	useRef,
	useState,
} from 'react';
import { getApi, initApi } from './api-client';
import { Button } from './components';
import { getKeycloak, initKeycloak } from './keycloak';
import type { User } from './types';

interface AuthConfig {
	apiUrl: string;
	keycloak: {
		url: string;
		realm: string;
		clientId: string;
	};
	authorize: (user: User) => boolean;
}

interface AuthContextValue {
	user: User | null;
	logout: () => void;
	refreshUser: () => Promise<void>;
}

const AuthContext = createContext<AuthContextValue | null>(null);

export function AuthProvider({
	config,
	children,
}: {
	config: AuthConfig;
	children: ReactNode;
}) {
	const [user, setUser] = useState<User | null>(null);
	const [ready, setReady] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const initialized = useRef(false);

	useEffect(() => {
		if (initialized.current) return;
		initialized.current = true;

		const keycloak = initKeycloak(config.keycloak);
		const api = initApi(config.apiUrl);

		let interval: ReturnType<typeof setInterval>;

		keycloak
			.init({ onLoad: 'login-required', checkLoginIframe: false })
			.then(async (authenticated) => {
				if (!authenticated) {
					keycloak.login();
					return;
				}

				// Atualiza token a cada 1 hora, renova se faltar menos de 5 min
				// CORRETO: verifica frequentemente, renova só quando necessário
				interval = setInterval(() => {
					keycloak
						.updateToken(300) // renova se faltar menos de 5 min para expirar
						.then((refreshed) => {
							if (refreshed) console.log('Token renovado automaticamente.');
						})
						.catch(() => {
							console.warn(
								'Falha ao renovar token, redirecionando para login...',
							);
							keycloak.login();
						});
				}, 30 * 1000); // ← verifica a cada 30s

				try {
					const { data } = await api.get<User>('/me');

					if (!data.credencial.ativo) {
						setError('Usuário inativo. Entre em contato com o administrador.');
						return;
					}

					if (!config.authorize(data)) {
						setError('Você não tem permissão para acessar este sistema.');
						return;
					}

					setUser(data);
					setReady(true);
				} catch (err) {
					console.error('Erro ao buscar dados do usuário:', err);
					setError('Erro ao buscar dados do usuário.');
				}
			})
			.catch((err) => {
				console.error('Erro ao inicializar Keycloak:', err);
				setError('Erro ao inicializar autenticação.');
			});

		return () => {
			if (interval) clearInterval(interval);
		};
	}, [config.apiUrl, config.keycloak, config.authorize]);

	const logout = () => {
		getKeycloak().logout({ redirectUri: window.location.origin });
	};

	const refreshUser = async () => {
		try {
			const { data } = await getApi().get<User>('/me');
			setUser(data);
		} catch (err) {
			console.error('Erro ao atualizar usuário:', err);
			setError('Erro ao buscar dados do usuário.');
		}
	};

	if (error) {
		return (
			<div
				style={{
					display: 'flex',
					alignItems: 'center',
					justifyContent: 'center',
					height: '100vh',
					flexDirection: 'column',
					gap: '1rem',
				}}
			>
				<p style={{ color: 'red' }}>{error}</p>
				<Button variant="outline" onClick={logout}>
					Sair
				</Button>
			</div>
		);
	}

	if (!ready) {
		return (
			<div
				style={{
					display: 'flex',
					alignItems: 'center',
					justifyContent: 'center',
					height: '100vh',
				}}
			>
				<p>Carregando...</p>
			</div>
		);
	}

	return (
		<AuthContext.Provider value={{ user, logout, refreshUser }}>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth(): AuthContextValue {
	const ctx = useContext(AuthContext);
	if (!ctx) throw new Error('useAuth deve ser usado dentro de AuthProvider');
	return ctx;
}

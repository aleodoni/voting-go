import axios from 'axios';
import { getKeycloak } from './keycloak';

let apiInstance: ReturnType<typeof axios.create> | null = null;

export function initApi(baseURL: string) {
	apiInstance = axios.create({ baseURL });

	// Interceptor adiciona token em todas as requisições
	apiInstance.interceptors.request.use(async (config) => {
		const keycloak = getKeycloak();

		if (!keycloak.token) {
			throw new Error('Token do Keycloak não disponível');
		}

		await keycloak.updateToken(30); // renova se faltar menos de 30s
		config.headers.Authorization = `Bearer ${keycloak.token}`;
		return config;
	});

	// Interceptor de resposta
	apiInstance.interceptors.response.use(
		(response) => response,
		(error) => {
			if (error.response?.status === 401) {
				getKeycloak().login();
			}
			return Promise.reject(error);
		},
	);

	return apiInstance;
}

export function getApi() {
	if (!apiInstance)
		throw new Error('API não inicializada. Chame initApi() primeiro.');
	return apiInstance;
}

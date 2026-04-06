import Keycloak from 'keycloak-js';

let keycloakInstance: Keycloak | null = null;

export function initKeycloak(config: {
	url: string;
	realm: string;
	clientId: string;
}) {
	if (!keycloakInstance) {
		keycloakInstance = new Keycloak({
			url: config.url,
			realm: config.realm,
			clientId: config.clientId,
		});
	}
	return keycloakInstance;
}

export function getKeycloak(): Keycloak {
	if (!keycloakInstance) throw new Error('Keycloak não inicializado');
	return keycloakInstance;
}

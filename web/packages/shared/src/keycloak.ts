import Keycloak from 'keycloak-js'

let keycloakInstance: Keycloak | null = null

export function getKeycloak(): Keycloak {
  if (!keycloakInstance) {
    throw new Error('Keycloak não foi inicializado. Chame initKeycloak() primeiro.')
  }
  return keycloakInstance
}

export function initKeycloak(config: { url: string; realm: string; clientId: string }): Keycloak {
  keycloakInstance = new Keycloak(config)
  return keycloakInstance
}
export interface User {
  id: string
  keycloak_id: string
  username: string
  email: string
  nome: string
  nome_fantasia?: string
  credencial: Credencial
}

export interface Credencial {
  ativo: boolean
  pode_administrar: boolean
  pode_votar: boolean
}
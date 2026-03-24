import axios from 'axios'
import { getKeycloak } from './keycloak'

let apiInstance: ReturnType<typeof axios.create> | null = null

export function getApi() {
  if (!apiInstance) {
    throw new Error('API não foi inicializada. Chame initApi() primeiro.')
  }
  return apiInstance
}

export function initApi(baseURL: string) {
  apiInstance = axios.create({ baseURL })

  apiInstance.interceptors.request.use(async (config) => {
    const keycloak = getKeycloak()
    await keycloak.updateToken(30)
    if (keycloak.token) {
      config.headers.Authorization = `Bearer ${keycloak.token}`
    }
    return config
  })

  apiInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response?.status === 401) {
        getKeycloak().login()
      }
      return Promise.reject(error)
    },
  )

  return apiInstance
}
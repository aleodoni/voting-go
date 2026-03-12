# Voting API

API de gerenciamento de votações desenvolvida em **Go**.
Permite criar reuniões, registrar projetos em votação e contabilizar votos dos participantes.

## Tecnologias

* Go
* Gin (HTTP framework)
* PostgreSQL
* Keycloak (autenticação)
* Docker
* k6 (testes de carga)

## Funcionalidades

* Autenticação via Keycloak
* Gerenciamento de usuários
* Criação de reuniões
* Cadastro de projetos em votação
* Registro de votos
* Apuração de resultados

---

# Como rodar o projeto

## 1. Clonar o repositório

```bash
git clone https://github.com/seu-usuario/voting-go.git
cd voting-go
```

## 2. Configurar variáveis de ambiente

A aplicação utiliza variáveis de ambiente para configuração.

Crie um arquivo `.env` na raiz do projeto:

```env
APPNAME=Voting API
APPVERSION=1.0.0
APPPORT=8080
APPENV=development

DBHOST=localhost
DBPORT=15432
DBUSER=postgres
DBPASSWORD=postgres
DBNAME=voting_db
DBSSLMODE=disable

KEYCLOAK_ISSUER=http://localhost:8081/realms/voting-realm
JWKSURL=http://localhost:8081/realms/voting-realm/protocol/openid-connect/certs
```

### Descrição das variáveis

| Variável          | Descrição                                               |
| ----------------- | ------------------------------------------------------- |
| `APPNAME`         | Nome da aplicação                                       |
| `APPVERSION`      | Versão da API                                           |
| `APPPORT`         | Porta em que a API será executada                       |
| `APPENV`          | Ambiente da aplicação (`development`, `production`)     |
| `DBHOST`          | Host do banco de dados                                  |
| `DBPORT`          | Porta do banco                                          |
| `DBUSER`          | Usuário do banco                                        |
| `DBPASSWORD`      | Senha do banco                                          |
| `DBNAME`          | Nome do banco                                           |
| `DBSSLMODE`       | Configuração SSL do PostgreSQL                          |
| `KEYCLOAK_ISSUER` | URL do realm do Keycloak usado para autenticação        |
| `JWKSURL`         | Endpoint de chaves públicas usado para validação do JWT |

## Rodando a aplicação

```bash
go run cmd/api/main.go
```

A API ficará disponível em:

```
http://localhost:8080
```


## 3. Rodar dependências (Docker)

```bash
docker-compose up -d
```

## 4. Rodar a aplicação

```bash
go run cmd/api/main.go
```

ou

```bash
go build -o voting-api cmd/api/main.go
./voting-api
```

A API estará disponível em:

```
http://localhost:8080
```

---

# Autenticação

A API utiliza **JWT emitido pelo Keycloak**.

Exemplo de header:

```
Authorization: Bearer <token>
```

---

# Endpoints principais

## Health check

```
GET /api/v1/health
```

Resposta:

```json
{
  "status": "ok"
}
```

---

## Listar reuniões do dia

```
GET /api/v1/reunioes-dia
```

Resposta:

```json
[
  {
    "id": "123",
    "data": "2026-03-12",
    "descricao": "Reunião ordinária"
  }
]
```

---

## Listar projetos de uma reunião

```
GET /api/v1/reunioes/{reuniaoId}/projetos
```

---

## Registrar voto

```
POST /api/v1/projetos/{projetoId}/votos
```

Body:

```json
{
  "voto": "SIM"
}
```

---

# Testes de carga

Os testes de carga são feitos com **k6**.

Executar:

```bash
k6 run tests/load/reunioes.js
```

Passando token:

```bash
TOKEN=<jwt> k6 run tests/load/reunioes.js
```

---

# Estrutura do projeto

```
cmd/
  api/
    main.go

internal/
  application/
  domain/
  infrastructure/
  interfaces/

pkg/

tests/
  load/
```

Arquitetura baseada em **Clean Architecture**.

---

# Roadmap

* [ ] Apuração automática de votos
* [ ] Websocket para resultado em tempo real
* [ ] Dashboard de votação
* [ ] Auditoria de votos

---

# Licença

MIT

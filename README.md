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
* Swagger (documentação)
* go-pdf/fpdf (geração de PDF)

## Funcionalidades

* Autenticação via Keycloak
* Gerenciamento de usuários
* Criação de reuniões
* Cadastro de projetos em votação
* Registro de votos
* Apuração de resultados
* Geração de relatório PDF por reunião

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

## 3. Rodar dependências (Docker)
```bash
docker-compose up -d
```

## 4. Gerar documentação Swagger
```bash
make swagger
```

> A pasta `docs/` é gerada automaticamente e não é versionada. Execute este comando sempre que alterar as anotações dos handlers.

## 5. Rodar a aplicação
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

# Documentação

A documentação interativa da API está disponível via Swagger UI após subir a aplicação:
```
http://localhost:8080/swagger/index.html
```

---

# Autenticação

A API utiliza **JWT emitido pelo Keycloak**.

Exemplo de header:
```
Authorization: Bearer <token>
```

No Swagger UI, clique em **Authorize** e informe o token no formato `Bearer <token>`.

---

# Endpoints

| Método   | Rota                                         | Descrição                                        | Auth  |
| -------- | -------------------------------------------- | ------------------------------------------------ | ----- |
| `GET`    | `/api/v1/health`                             | Health check                                     | ❌    |
| `GET`    | `/api/v1/me`                                 | Retorna o usuário autenticado                    | ✅    |
| `GET`    | `/api/v1/usuarios`                           | Pesquisa usuários (admin)                        | ✅    |
| `PUT`    | `/api/v1/usuarios/fantasia-credenciais`      | Atualiza nome fantasia e permissões              | ✅    |
| `PATCH`  | `/api/v1/usuarios/{id}/credencial`           | Atualiza credencial de um usuário                | ✅    |
| `GET`    | `/api/v1/reunioes-dia`                       | Retorna reuniões do dia                          | ✅    |
| `GET`    | `/api/v1/reunioes/{reuniaoId}/projetos`      | Retorna projetos de uma reunião (admin)          | ✅    |
| `GET`    | `/api/v1/reunioes/{reuniaoId}/relatorio`     | Gera relatório PDF da reunião                    | ✅    |
| `POST`   | `/api/v1/projetos/{projetoId}/votacao/abrir` | Abre uma votação (admin)                         | ✅    |
| `POST`   | `/api/v1/projetos/{projetoId}/votacao/fechar`| Fecha uma votação (admin)                        | ✅    |
| `DELETE` | `/api/v1/projetos/{projetoId}/votacao`       | Cancela uma votação (admin)                      | ✅    |
| `POST`   | `/api/v1/votacao/{votacaoId}/voto`           | Registra um voto                                 | ✅    |
| `GET`    | `/api/v1/votacao/aberta`                     | Retorna o projeto com votação aberta             | ✅    |
| `GET`    | `/api/v1/votacao/stats`                      | Retorna estatísticas de votação do dia (admin)   | ✅    |

Para detalhes completos de request/response, consulte o Swagger UI.

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
    main.go           # Entrypoint e anotações gerais do Swagger

internal/
  application/        # Casos de uso
  domain/             # Entidades e interfaces de repositório
  infrastructure/     # Implementações de persistência e mappers
  handler/            # Handlers HTTP, DTOs e mappers de response
  middleware/         # JWT e outros middlewares
  router/             # Configuração de rotas
  platform/           # Utilitários (JWT, ID, transações)

pkg/
  logger/

tests/
  api/                # Testes de carga com k6
```

Arquitetura baseada em **Clean Architecture**.

---

# Roadmap

* [x] Apuração automática de votos
* [x] Votação em tempo real via SSE
* [ ] Dashboard de votação
* [ ] Auditoria de votos

---

# Licença

MIT
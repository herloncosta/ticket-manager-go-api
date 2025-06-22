# 📞 API de Gestão de Chamados - Backend em Go

API RESTful para gestão de chamados técnicos, desenvolvida em **Go (Golang)**, seguindo os princípios da **Arquitetura Limpa**, com foco em **escalabilidade, segurança, performance e manutenibilidade**.

> 🔥 Backend de um sistema para empresas de TI que gerenciam chamados de múltiplos clientes.

---

## 🚀 Tecnologias e Ferramentas

- [Go (Golang)](https://golang.org/) — Backend
- [Gin](https://gin-gonic.com/) — Web Framework HTTP
- [SQLite](https://www.sqlite.org/) — Banco de dados leve (ideal para desenvolvimento e testes)
- Arquitetura Limpa (Clean Architecture)
- RESTful API
- Docker (Opcional)
- SQL Migrations (Manual ou scripts)

---

## 🏗️ Estrutura de Pastas

```bash
.
├── cmd/                # Entry point (main.go)
├── internal/           # Código interno (Domínio, Handlers, Serviços, Repositórios, Router, DB, Config)
│   ├── app/
│   │   ├── handler/    # Camada de entrega (HTTP handlers)
│   │   ├── model/      # Entidades (Domain Models)
│   │   ├── repository/ # Acesso ao banco de dados
│   │   └── service/    # Regras de negócio (Use Cases)
│   ├── config/         # Configurações da aplicação
│   ├── db/             # Conexão com o banco de dados
│   └── router/         # Definição das rotas
├── pkg/                # Pacotes utilitários
├── go.mod
└── go.sum
```

---

## 🔥 Funcionalidades

- ✅ CRUD de Chamados (`Tickets`)
- 🔍 Filtragem e consulta de chamados
- 🏢 Suporte a múltiplos clientes (multi-tenancy simples via `client_id`)
- 🛡️ Segurança com prepared statements (contra SQL Injection)
- 📜 Logs, validações e tratamento de erros
- 🚀 Pronto para escalar: fácil adaptação para PostgreSQL, MySQL, MongoDB, etc.

---

## 📦 Instalação e Execução Local

### Pré-requisitos:

- [Go](https://golang.org/dl/) 1.21 ou superior
- [SQLite3](https://sqlite.org/download.html) instalado (opcional, mas recomendado para inspeção do banco)

### Clone o projeto

```bash
git clone https://github.com/herloncosta/ticket-manager-go-api.git
cd nome-do-projeto
```

### Instale as dependências

```bash
go mod tidy
```

> ⚙️ Você pode rodar esse script via DB Browser for SQLite, SQLite CLI, ou adicionar uma automação futura.

### Execute o projeto

```bash
go run cmd/api/main.go
```

> A API estará rodando em `http://localhost:8080`

---

## 🔗 Endpoints

### Health Check

```http
GET /health
```

### Tickets

| Método | Endpoint       | Descrição            |
| ------ | -------------- | -------------------- |
| GET    | `/tickets`     | Listar chamados      |
| GET    | `/tickets/:id` | Obter chamado por ID |
| POST   | `/tickets`     | Criar novo chamado   |
| PUT    | `/tickets/:id` | Atualizar chamado    |
| DELETE | `/tickets/:id` | Deletar chamado      |

### 📦 Payload Exemplo (POST/PUT)

```json
{
  "title": "Internet Instável",
  "description": "Cliente relata quedas constantes na conexão.",
  "client_id": 1
}
```

---

## 🔐 Segurança

- ✅ SQL Injection protegido (prepared statements)
- 🚧 Futuro: autenticação via JWT
- 🚧 Rate Limiting e CORS configuráveis

---

## 🐳 Docker (Opcional)

### Build da imagem

```bash
docker build -t tickets-api .
```

### Rodar container

```bash
docker run -p 8080:8080 tickets-api
```

> ⚠️ Banco SQLite será armazenado dentro do container (persistência temporária).

---

## 📈 Melhorias Futuras (Roadmap)

- 🔐 Implementação de autenticação e autorização com JWT e RBAC
- 📊 Logs estruturados
- 🐘 Suporte nativo para PostgreSQL, MySQL e MongoDB
- 🚦 Middleware de rate limiting e segurança
- 📝 Integração com Swagger/OpenAPI para documentação automática
- 📤 Deploy com Docker Compose, Kubernetes ou Serverless
- 🧪 Testes unitários e testes de integração

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se livre para abrir uma issue, criar um fork e enviar um Pull Request.

---

## 🧑‍💻 Autor

Feito com 💻 por **Herlon Costa**

- 🔗 [LinkedIn](https://www.linkedin.com/in/herloncosta)

---

## ⚖️ Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

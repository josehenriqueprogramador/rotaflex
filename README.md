# 🚗 RotaFlex PRO

Plataforma backend inspirada em apps de mobilidade urbana, desenvolvida com Go + Gin + Docker + Clean Architecture.

---

## 📌 Visão Geral

O RotaFlex PRO é uma API moderna para gerenciamento de corridas urbanas, preparada para evoluir para:

- Uber / 99
- Delivery
- Logística
- Rastreamento em tempo real

---

## 🧠 Tecnologias

- Go 1.25
- Gin Framework
- Docker
- PostgreSQL
- Redis
- Clean Architecture

---

## 📁 Estrutura

backend/
├── cmd/api/main.go
├── internal/domain
├── internal/handler
├── internal/middleware
├── internal/repository/postgres
├── internal/service
├── internal/websocket
├── go.mod
└── go.sum

docker/go/Dockerfile
docker-compose.yml

---

## 🧱 Arquitetura

HTTP Request
↓
Handler
↓
Service
↓
Repository
↓
PostgreSQL / Redis

---

## 🚀 Rodar Localmente

cd backend

go run cmd/api/main.go

Servidor:

http://localhost:8000

---

## 🐳 Rodar com Docker

docker compose up -d --build

---

## 📡 Endpoints

GET /api/health

Resposta:

{
  "status": "ok"
}

POST /api/rides

Resposta:

{
  "message": "corrida criada com sucesso"
}

---

## 🧪 Testes

curl http://localhost:8000/api/health

curl -X POST http://localhost:8000/api/rides

---

## 🔐 Próximas Features

- JWT Authentication
- Cadastro Passageiro
- Cadastro Motorista
- PostgreSQL real
- Redis cache
- WebSocket tempo real
- Matching motoristas
- Painel Admin
- Deploy Cloud

---

## 📈 Roadmap

Fase 1
- Estrutura base
- API inicial
- Docker

Fase 2
- JWT

Fase 3
- WebSocket

Fase 4
- Matching

---

## 👨‍💻 Autor

José Henrique Programador

https://github.com/josehenriqueprogramador

---
## 📺 Processo de Desenvolvimento

[![Assista no YouTube](https://img.youtube.com/vi/_w1NCeMMqn8/0.jpg)](https://youtu.be/_w1NCeMMqn8)

---

## 📜 Licença

MIT License


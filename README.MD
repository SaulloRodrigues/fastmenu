# 🚀 Menu Delivery - VittaTech - Sistema de Menu Digital

[![Go](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org/)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.5-4FC08D.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.8-blue.svg)](https://www.typescriptlang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue.svg)](https://www.postgresql.org/)

> **Sistema completo para gerenciamento de menus digitais**  
> Ideal para restaurantes, lanchonetes e estabelecimentos que desejam modernizar o atendimento ao cliente.

---

## 📋 Índice

- [Sobre o Projeto](#-sobre-o-projeto)
- [Funcionalidades](#-funcionalidades)
- [Tecnologias](#-tecnologias)
- [Arquitetura](#-arquitetura)
- [Pré-requisitos](#-pré-requisitos)
- [Instalação](#-instalação)
- [Configuração](#-configuração)
- [Como Rodar](#-como-rodar)
- [API Endpoints](#-api-endpoints)
- [Exemplos de Uso](#-exemplos-de-uso)
- [Contribuição](#-contribuição)

---

## 🎯 Sobre o Projeto

O **Menu Delivery** é uma solução moderna para gerenciamento de menus digitais, permitindo que estabelecimentos alimentícios ofereçam uma experiência digital completa aos seus clientes. O sistema permite cadastrar produtos, organizá-los por categorias e gerenciar usuários de forma simples e intuitiva.

---

## ✨ Funcionalidades

### 👤 **Gerenciamento de Usuários**
- ✅ Cadastro e listagem de usuários
- ✅ Dados completos (nome, email, endereço, telefone, CPF)

### 🍕 **Gerenciamento de Produtos**
- ✅ CRUD completo de produtos
- ✅ Preços precisos com decimal
- ✅ Upload de imagens
- ✅ Descrições detalhadas
- ✅ Associação com múltiplas categorias

### 📁 **Gerenciamento de Categorias**
- ✅ CRUD completo de categorias
- ✅ Organização de produtos

### 🔍 **Recursos Especiais**
- ✅ Busca de produtos por categoria
- ✅ API RESTful completa
- ✅ Health check endpoint
- ✅ Arquitetura escalável

---

## 🛠 Tecnologias

### **Backend**
- **[Go 1.24](https://golang.org/)** - Linguagem de programação
- **[Gin Framework](https://gin-gonic.com/)** - Framework web
- **[GORM](https://gorm.io/)** - ORM para Go
- **[PostgreSQL](https://www.postgresql.org/)** - Banco de dados relacional
- **[Shopspring Decimal](https://github.com/shopspring/decimal)** - Precisão decimal para preços

### **Frontend**
- **[Vue.js 3.5](https://vuejs.org/)** - Framework JavaScript
- **[TypeScript](https://www.typescriptlang.org/)** - Linguagem tipada
- **[Vite](https://vitejs.dev/)** - Build tool
- **[Tailwind CSS](https://tailwindcss.com/)** - Framework CSS
- **[Lucide Icons](https://lucide.dev/)** - Biblioteca de ícones

### **DevOps**
- **[Docker](https://www.docker.com/)** - Containerização
- **[Docker Compose](https://docs.docker.com/compose/)** - Orquestração

---

## 🏗 Arquitetura

```
fastmenu/
├── backend/              # API em Go
│   ├── controllers/      # Handlers HTTP
│   ├── models/           # Modelos de dados
│   ├── routes/           # Definição de rotas
│   ├── services/         # Lógica de negócio
│   └── config/           # Configurações
├── frontend/             # Interface em Vue.js
│   ├── src/
│   │   ├── components/   # Componentes reutilizáveis
│   │   ├── pages/        # Páginas da aplicação
│   │   └── router/       # Configuração de rotas
└── docker-compose.yml    # Orquestração Docker
```

---

## 📦 Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- **[Docker](https://www.docker.com/get-started)** (versão 20.10 ou superior)
- **[Docker Compose](https://docs.docker.com/compose/install/)** (versão 2.0 ou superior)
- **[Git](https://git-scm.com/downloads)**

---

## 🚀 Instalação

### **1. Clone o repositório**

```bash
git clone https://github.com/seuusuario/fastmenu.git
cd fastmenu
```

### **2. Copie o arquivo de ambiente**

```bash
cp example.env .env
```

---

## ⚙️ Configuração

Edite o arquivo `.env` com suas configurações:

```env
# Backend
APP_PORT=8080

# Frontend
VITE_BASE_URL=http://localhost:8080

# Database
POSTGRES_USER=fastmenu_user
POSTGRES_PASSWORD=fastmenu_pass
POSTGRES_DB=fastmenu_db
```

---

## 🏃 Como Rodar

### **Opção 1: Usando Docker Compose (Recomendado)**

```bash
# Inicia todos os serviços
docker-compose up -d

# Verifica o status
docker-compose ps

# Visualiza os logs
docker-compose logs -f backend
```

**Serviços disponíveis:**
- 🌐 **API Backend**: `http://localhost:8080`
- 🎨 **Frontend**: `http://localhost:5173`
- 🗄️ **PostgreSQL**: `localhost:5432`

### **Opção 2: Executar localmente**

#### **Backend**

```bash
cd backend
go mod download
go run main.go
```

#### **Frontend**

```bash
cd frontend
yarn install
yarn dev
```

---

## 📡 API Endpoints

### **Health Check**
```
GET /health
```

### **Usuários**
```
GET    /api/users           # Lista todos os usuários
POST   /api/users           # Cria um novo usuário
```

### **Produtos**
```
GET    /api/products                        # Lista todos os produtos
GET    /api/products/:id                    # Busca produto por ID
GET    /api/products/categories/:category_id # Produtos por categoria
POST   /api/products                        # Cria um novo produto
PATCH  /api/products/:id                    # Atualiza um produto
DELETE /api/products/:id                    # Remove um produto
```

### **Categorias**
```
GET    /api/categories   # Lista todas as categorias
GET    /api/categories/:id # Busca categoria por ID
POST   /api/categories   # Cria uma nova categoria
PUT    /api/categories/:id # Atualiza uma categoria
DELETE /api/categories/:id # Remove uma categoria
```

---

## 💡 Exemplos de Uso

### **1. Criar uma Categoria**

```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pizzas"
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "name": "Pizzas"
}
```

### **2. Criar um Produto**

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pizza Margherita",
    "description": "Pizza clássica com molho de tomate, mussarela e manjericão",
    "image_url": "https://example.com/pizza.jpg",
    "price": "45.90",
    "category_ids": [1]
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "name": "Pizza Margherita",
  "description": "Pizza clássica com molho de tomate, mussarela e manjericão",
  "image_url": "https://example.com/pizza.jpg",
  "price": "45.90",
  "categories": [
    {
      "id": 1,
      "name": "Pizzas"
    }
  ]
}
```

### **3. Listar Produtos por Categoria**

```bash
curl http://localhost:8080/api/products/categories/1
```

### **4. Criar um Usuário**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "senha123",
    "address": "Rua Exemplo, 123",
    "phone": "11999999999",
    "cpf": "12345678900"
  }'
```

---

## 🧪 Testando a API

Você pode usar ferramentas como:
- **[Postman](https://www.postman.com/)**
- **[Insomnia](https://insomnia.rest/)**
- **[Thunder Client](https://www.thunderclient.com/)** (extensão VS Code)
- **[curl](https://curl.se/)** (linha de comando)

Exemplo de collection para Postman:
```
[
  {
    "name": "Menu Delivery API",
    "requests": [
      {
        "name": "Health Check",
        "method": "GET",
        "url": "http://localhost:8080/health"
      }
    ]
  }
]
```

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

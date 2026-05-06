# ⚙️ Backend - Loquor Sign

API responsável pelo gerenciamento de símbolos para a aplicação de Comunicação Alternativa.

---

## 🚀 Como rodar o backend

### Pré-requisitos
- Go instalado
- MongoDB instalado

---

### 🟢 1. Iniciar o MongoDB

```bash
mongod
```

### 🔵 2. Rodar a API

#### Opção 1 - go run (pode exigir permissão no Windows)
```bash
cd backend
go run .
```

#### Opção 2 - go build (recomendado para Windows)
```bash
cd backend
go build -o app.exe
.\app.exe
```

- A API estará disponível em:

http://localhost:8080

---

## 🧠 Banco de Dados

- Banco: MongoDB

- URL: mongodb://localhost:27017

- Database: loquor-sign

---

## 📁 Estrutura

backend
 ├── handlers
 ├── models
 ├── routes
 ├── database
 ├── middleware
 └── main.go

---

## 🔐 Autenticação

A API utiliza autenticação via JWT.

### 🔐 Login

POST /login

### 📥 Body

JSON
{
  "email": "usuario@gmail.com",
  "password": "123456"
}

### 📥 Resposta

JSON
{
  "user":{
    "id":"INCLUIR_ID_AQUI",
    "name":"Usuário",
    "email":"usuário@gmail.com"
  },
  "token":"JWT_TOKEN"

}

### 🔒 Rotas Protegidas

As rotas de phrases exigem autenticação.

Enviar no header:

`Authorization: Bearer SEU_TOKEN`

---

## 📡 Rotas da API

### 👤 USERS

🔹 POST /users
- Cria um novo usuário.

📥 Body
{
  "name": "Bartolomeu",
  "email": "Barto@gmail.com",
  "password": "123456"
}

### 🗂 CATEGORIES

🔹 GET /categories
- Listar todas as categorias

🔹 POST /categories
- Criar uma categoria
JSON
{
  "name": "Ações"
}

🔹 PUT /categories/:id
- Atualiza uma categoria

🔹 DELETE /categories
- Remove um categoria

### 🔤 SYMBOLS

🔹 GET /symbols
- Retorna todos os símbolos.

✅ Resposta
[
  {
    "id": "1",
    "name": "Água",
    "image": "agua.png",
    "category_id": "1"
  }
]

🔹 POST /symbols
- Cria um novo símbolo.

📥 Body
{
  "id": "4",
  "name": "Beber",
  "image": "beber.png",
  "category_id": "1"
}

✅ Resposta
{
  "id": "4",
  "name": "Beber",
  "image": "beber.png",
  "category_id": "1"
}

🔹 PUT /symbols/:id
- Atualiza um símbolo existente.

📥 Body
{
  "id": "1",
  "name": "Água (editado)",
  "image": "agua.png",
  "category_id": "1"
}
✅ Resposta
{
  "id": "1",
  "name": "Água (editado)",
  "image": "agua.png",
  "category_id": "1"
}

🔹 DELETE /symbols/:id
- Remove um símbolo.

✅ Resposta
{
  "message": "símbolo deletado com sucesso"
}

### 🧩 PHRASES (🔒 protegidas)

🔹 GET /phrases
- Lista frases do usuário autenticado

🔹 POST /phrases
JSON
{
  "symbols": ["ID_SYMBOLS_1", "ID_SYMBOLS_2"]
}

🔹 PUT /phrases/:id
- Atualizar uma frase

🔹 DELETE /phrases/:id
- Remove uma frase

---

## 🧪 Observações importantes
- IDs são gerados automaticamente pelo 
- MongoDB (ObjectID)
- Não é necessário enviar id no body
- Senhas são armazenadas com hash (bcrypt)
- Cada usuário acessa apenas suas próprias frases

---

## 📌 Status do Projeto

🚧 Protótipo funcional
✅ CRUD completo
✅ Autenticação com JWT
✅ Integração com MongoDB

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
```bash
cd backend
go run .
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
 └── main.go

---

## 📡 Rotas da API

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

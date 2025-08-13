# API de Frases

API RESTful desenvolvida em **Go**, utilizando **Echo**, **GORM**, **JWT** e **SQLite**.  
Permite cadastro de usuários, login com autenticação JWT e gerenciamento de frases pessoais.

---

## Tecnologias Usadas

- Go
- Echo
- GORM
- JWT
- SQLite

---

## Rodando Localmente

1. Clone o repositório:

```bash
git clone https://github.com/miguellaig/api-phrases.git
```

2. Entre na pasta do projeto:

```bash
cd api-phrases
```

3. Instale dependências:

```bash
go mod tidy
```

4. Rode a API:

```bash
go run main.go
```

> Observação: testes foram realizados pelo Postman.

---

## Endpoints

### Usuários

- `POST /register` – Cadastra um usuário
- `POST /login` – Autentica usuário e retorna token JWT

### Frases (autenticadas)

- `POST /phrase` – Cria uma nova frase
- `GET /phrase` – Lista as frases de um usuário
- `PUT /phrase/:id` – Atualiza uma frase existente
- `DELETE /phrase/:id` – Deleta uma frase

> Observação: todas as rotas de frase requerem token JWT.

---

## Detalhes do Projeto

- Autenticação com JWT (chave secreta definida diretamente no código).
- Estrutura modular com separação de **handlers**, **services** e **models**.
- Código organizado para facilitar futuras melhorias, como testes automatizados e documentação Swagger.

---

## Próximos Passos

- Adicionar testes automatizados
- Implementar Swagger/OpenAPI para documentação
- Remover chave secreta do código e colocar em variável de ambiente

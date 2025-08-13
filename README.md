# API Alemão - Projeto Finalizado

Este projeto é uma API simples para gerenciar frases, com autenticação básica via JWT.

## Como rodar

1. Tenha o Go instalado (versão 1.18+ recomendada)  
2. Configure seu banco de dados PostgreSQL (ou outro suportado pelo GORM)  
3. Ajuste as variáveis de ambiente (ex: conexão com banco, chave JWT)  
4. Rode o comando para iniciar a API:

   go run main.go

## Rotas principais

| Método | Rota               | Descrição                                                |
|--------|--------------------|---------------------------------------------------------|
| POST   | /phrases           | Criar uma nova frase                                     |
| GET    | /phrases           | Listar frases do usuário (opcional filtro por linguagem e texto) |
| PUT    | /phrases/:id       | Atualizar uma frase existente                            |
| DELETE | /phrases/:id       | Deletar uma frase existente                              |
| POST   | /login             | Autenticar usuário e gerar token                         |

## Requisitos

- Banco de dados configurado e acessível  
- Usuário autenticado via token JWT para rotas protegidas

## Melhorias futuras

As melhorias planejadas para evoluir este projeto estão anotadas no arquivo `MELHORIAS_FUTURAS.md`.

---

Qualquer dúvida ou sugestão, só avisar!

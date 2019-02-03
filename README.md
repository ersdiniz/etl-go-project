# ETL Project in GoLang

Para rodar o projeto, apenas dois passos são necessários:

1 - Criar o database (apenas caso ainda não exista):

`CREATE DATABASE postgres
     WITH 
     OWNER = postgres
     ENCODING = 'UTF8'
     CONNECTION LIMIT = -1;`

Obs.: O usuário para conexão com o banco deve ser o padrão do Postgres:
usuário: **postgres**
senha: **postgres**

2 - Rodar o projeto:

`go run main.go`

Ao rodar o projeto, todas as tabelas utilizadas são criadas/recriadas.
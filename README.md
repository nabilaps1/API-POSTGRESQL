# API-POSTGRESQL
Implementação de uma API do zero utilizando Golang + PostgreSQL. Inclui levantar um container PostgreSQL no Docker, criar tabela e setar permissões. Testes feitos no Postman. Desenvolvimento focado na aprendizagem da linguagem Go para criação de APIs.

## Estrutura da aplicação (packages)
  * db: Responsável pela conexão com o Bando de Dados (Abrir/fechar conexão).
  * configs: Responsável por ler arquivo de configuração, como dados do Host, Porta, Endereço de conexão.
  * handlers: Recebe e trata as chamadas da API.
  * models: Responsável pelas transações com o banco de dados.

## Comandos Docker e Postgres
  *  Levantar container com Postgres
    ```
    docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
    ```
  * Conectar no container
    ```
    docker exec -it api-todo psql -U postgres
    ```
  * Criar usuario e senha
    ```
    postgres=# create user user_todo;
    postgres=# alter user user_todo with encrypted password '1122';
    ```
  * Permitir que o usuário acesse o database e altere tabelas
    ```
    postgres=# grant all privileges on database api-todo to user_todo;
    postgres=# grant all privileges on all tables in schema public to user_todo;
    ```
  * Conectar usuario e criar tabela todos
    ```
    postgres=# \c api_todos;
    postgres=# create table todos(id serial primary key, title varchar, description text, done bool default FALSE);
    ```

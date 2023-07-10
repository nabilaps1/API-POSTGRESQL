package models

import (
	"github.com/nabilaps1/API-POSTGRESQL/db"
)

func Insert(todo Todo) (id int64, err error) { // chamada em handlers
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close() // independente do que aconteça abaixo (se erro ou nao), ao final é fechada a conexao

	// cria tabela no bd. Faz o INSERT dos dados, e ao final retorna o numero de id que gerou dps da insercao
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// faz a transacao no banco de dados
	// scan retorna o id e atribui o valor para a variavel id
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}

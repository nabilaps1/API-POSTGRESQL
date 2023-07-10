package models

import "github.com/nabilaps1/API-POSTGRESQL/db"

func Get(id int64) (todo Todo, err error) { // pega 1 todo por vez
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close() // fecha a conexao quando a execucao do get for finalizada

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id) // $1 = id

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}

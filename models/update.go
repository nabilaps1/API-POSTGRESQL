package models

import "github.com/nabilaps1/API-POSTGRESQL/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	// O retorno serve como uma 2a verificacao, ja que se atualizar mais do que o numero passado no parametro, algo esta errado
	return res.RowsAffected()
}

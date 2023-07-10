package models

import "github.com/nabilaps1/API-POSTGRESQL/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close() // fecha a conexao quando a execucao do get for finalizada

	rows, err := conn.Query(`SELECT * FROM todos`) // Sem filtro, pega todos os valores que tem no BD
	if err != nil {
		return
	}

	// se nao houver erros
	for rows.Next() { // escaneia todos os itens retornados e passa o valor deles para o objeto todo
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue // aqui, uma boa pratica seria logar os erros, para analise dps
		}

		todos = append(todos, todo) // retorna uma lista de todos. vai ser usada em handleres
	}
	return
}

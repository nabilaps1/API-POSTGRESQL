package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // drive de conexcao do postgre
	"github.com/nabilaps1/API-POSTGRESQL/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// string de conexao
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc) // abre conexao com Banco de Dados

	if err != nil {
		panic(err)
	}

	err = conn.Ping() // ping no banco para verificar se conexao esta estabelecida

	return conn, err
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

//nome do servidor
var server = "localhost"

//porta do servidor
var port = 1433

//nome do usuário
var user = "user"

//senha do usuário
var password = "password"

//base a se conectar
var database = "MeuApp"

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Erro ao criar conexão: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Conectado!")

	//Loop para ficar bombardeando o banco com requisições, caso queira apenas 1 vez comentar o for
	var i int = 0
	var teste string
	for {
		fmt.Println("Requisição no.", i)
		result, err := getDadosProcedure()
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
			fmt.Scanf(teste)
		}
		i++
	}
}

func getDadosProcedure() ([]string, error) {
	var references []string

	ctx := context.Background()

	// Verifica se o banco está respondendo
	err := db.PingContext(ctx)
	if err != nil {
		return references, err
	}

	tsql := fmt.Sprintf("NomeStorageProcedure 'parametro(s)' ou query que queira fazer ex.: Select * from Table where id = 1")

	// Executa a query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return references, err
	}

	defer rows.Close()

	// Traz os resultados
	for rows.Next() {
		var reference string

		// Pega os valores das linhas
		err := rows.Scan(&reference)
		if err != nil {
			return references, err
		}
		references = append(references, reference)
	}

	return references, nil
}

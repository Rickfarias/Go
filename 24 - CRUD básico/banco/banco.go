package banco

import(
	 "database/sql"

	 _ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang_3506@@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
}
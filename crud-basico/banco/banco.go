package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	strConexao := "root:secret@tcp(localhost:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}

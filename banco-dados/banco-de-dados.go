package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	strConexao := "root:secret@tcp(localhost:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexão bem-sucedida!")

	linhas, err := db.Query("select * from usuarios ")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}

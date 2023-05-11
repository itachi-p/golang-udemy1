package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//CRUD のD、削除処理
//Create, Read, Update, Delete

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}
}

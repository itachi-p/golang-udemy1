package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//CRUD処理

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データの更新
	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	_, err := Db.Exec(cmd, 25, "tarou")
	if err != nil {
		log.Fatalln(err)
	}
}

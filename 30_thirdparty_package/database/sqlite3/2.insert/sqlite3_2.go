package main

import (
	"database/sql"
	"log"

	//コードの中では使用しないが必要なインポート
	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//CRUD処理

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データの追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := Db.Exec(cmd, "tarou", 20)
	if err != nil {
		log.Fatalln(err)
	}
}

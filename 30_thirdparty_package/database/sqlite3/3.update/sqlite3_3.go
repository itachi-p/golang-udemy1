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

	//データの更新
	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	//VALUESの中の"?"を置換して実際にデータを挿入する
	_, err := Db.Exec(cmd, 25, "tarou")
	if err != nil {
		log.Fatalln(err)
	}
}

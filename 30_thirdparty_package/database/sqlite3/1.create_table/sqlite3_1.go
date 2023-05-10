package main

import (
	"database/sql"
	"log"

	//コードの中では使用しないが必要なインポート
	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
// テーブル作成

var Db *sql.DB

func main() {

	Db, _ := sql.Open("sqlite3", "./example.sql")

	//処理が終了したらSqlite3との接続を切断する
	defer Db.Close()

	//新規テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS persons(
		name STRING,
		age INT)`

	_, err := Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}

}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//CRUD処理

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データベース（特定のテーブル）の全データを取得
	cmd := "SELECT * FROM persons"
	//Query 条件に合う全てのデータを取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err2 := rows.Scan(&p.Name, &p.Age)

		if err2 != nil {
			log.Println(err2)
		}
		pp = append(pp, p)

	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

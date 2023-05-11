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
	//Query : 条件に合う全てのデータを取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	//構造体PersonのSliceとしてデータベースの各行を格納
	var pp []Person

	//読み取ったデータの行数分だけループ
	for rows.Next() {
		//1行ごとにデータを受け取る入れ物にする
		var p Person
		err2 := rows.Scan(&p.Name, &p.Age)

		if err2 != nil {
			log.Println(err2)
		}
		//NameとAgeのセットを1行づつSliceに追加
		pp = append(pp, p)

	}
	//取得した全データを行毎に表示
	/* ここではindexを破棄してNameとAgeのデータだけを列挙しているが、
	そもそもDBのテーブル側（最初のCREATE TABLE時）に
	UNIQUEなIDをPRIMARY KEYとして設定した方が設計的に望ましい*/
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

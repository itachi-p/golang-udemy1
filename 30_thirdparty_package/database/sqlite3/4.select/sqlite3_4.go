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

	//CREATEしたTABLEにUNIQUE KEYの設定が無い場合、全く同じデータも追加可能
	//通常は一意な整数であるIDなど、重複登録ができない項目(カラム)を設定する
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := Db.Exec(cmd, "hanako", 19)
	if err != nil {
		log.Fatalln(err)
	}

	//特定のデータ（1行のみ）を取得
	cmd = "SELECT * FROM persons WHERE age = ?"
	//QueryRow : 1レコード取得
	row := Db.QueryRow(cmd, 25)
	var p Person
	err2 := row.Scan(&p.Name, &p.Age)
	if err2 != nil {
		//指定の条件に合うデータが存在しない場合
		if err2 == sql.ErrNoRows {
			log.Println("No row.")
		} else {
			log.Println(err2)
		}
	}
	fmt.Println(p.Name, p.Age)
}

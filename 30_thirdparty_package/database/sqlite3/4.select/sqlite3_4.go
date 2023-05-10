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

	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := Db.Exec(cmd, "hanako", 19)
	if err != nil {
		log.Fatalln(err)
	}

	//特定のデータを取得
	cmd = "SELECT * FROM persons WHERE age = ?"
	//QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 25)
	var p Person
	err2 := row.Scan(&p.Name, &p.Age)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			log.Println("No row.")
		} else {
			log.Println(err2)
		}
	}
	fmt.Println(p.Name, p.Age)
}

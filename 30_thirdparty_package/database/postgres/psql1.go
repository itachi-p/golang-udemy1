package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Database
//Postgres

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	defer Db.Close()

	/*
			//C
			//CREATE TABLEはpsqlではSQL文を記述した.sqlファイルにて実行済み
			//データの追加 sqlite3のコードとほぼ同じ(違いはValueの中の?→$xだけ)
			//VALUES()の"$"がSQLインジェクション(悪意あるコマンド)をエスケープしてくれる
			cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
			//VALUESの中の"$x"を置換して実際にデータを挿入する
			_, err := Db.Exec(cmd, "Takeo", 18)
			if err != nil {
				log.Fatalln(err)
			}

			//R
			//特定のデータ（1行のみ）を取得
			cmd = "SELECT * FROM persons WHERE age = $1"
			//QueryRow : 1レコード取得
			row := Db.QueryRow(cmd, 20)
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

			//データベース（特定のテーブル）の全データを取得
			cmd = "SELECT * FROM persons"
			//Query : 条件に合う全てのデータを取得
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
			//取得した全データを行毎に表示
			for _, p := range pp {
				fmt.Println(p.Name, p.Age)
			}

		//U
		//データの更新
		cmd := "UPDATE persons SET age = $1 WHERE name = $2"
		_, err := Db.Exec(cmd, 25, "Nancy")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	//D
	//データの削除
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json
//エンコーダーとデコーダーの機能を提供

//構造体からJSONテキストへの変換

// Userに入れ子にする構造体（テストのため簡易設定）
type StructA struct{}

type User struct {
	//3番目のフィールドには変換後のJSONでの別名を指定できる
	//省略した場合は左のフィールド名がそのまま入る
	//omitemptyを指定し、かつ値が型の初期値の場合JSON側では非表示になる
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Email   string
	Created time.Time
	A       *StructA `json:"aaa"`
}

func main() {

	u := new(User)
	u.ID = 0
	u.Name = "testuser"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshl: JSONに変換してbyteのsliceとして受け取る
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	//JSON形式の文字列型に変換して表示
	fmt.Println(string(bs))

	//JSON形式から構造体への変換

	//先ほど生成されたbyteのsliceの型を確認
	fmt.Printf("%T\n", bs)

	//別のUserを追加で生成
	u2 := new(User)
	//（確認）new()の戻り値はポインタ型(アドレス)で返されている
	fmt.Printf("%T\n", u2)

	//Unmarshal: JSONを構造体に変換
	//第二引数には構造体のアドレスを渡す
	//new()の結果が最初からポインタ型なので&は不要(付けても結果は同じ)
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println("err:", err)
	}

	//構造体に戻っているか確認
	fmt.Println("u2:", u2)

}

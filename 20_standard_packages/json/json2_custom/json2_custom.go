package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

//json

//Marshal, Unmarshalのカスタム

type StructA struct{}

type User struct {
	//3番目のフィールドには変換後のJSONでの別名を指定できる
	//別名を省略した場合は左のフィールド名がそのまま入る
	//int→stringなど構造体とJSON側で型変換もできる
	//omitemptyを指定し、かつ値が型の初期値の場合JSON側では非表示になる
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string
	Created time.Time
	A       *StructA `json:"aaa,omitempty"`
}

// Marshalのカスタマイズ
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		//カスタマイズしないフィールドも列挙する必要がある
		//今回はテストのため一部省略
		ID    int    `json:"id,string"` //元のintをJSON側では文字に変換
		Name  string `json:"name"`
		Email string `json:"-"` //JSON側では非表示にする
	}{
		//カスタマイズしたい内容
		ID:    1000 + u.ID,     //int型で計算後に文字列に変換される
		Name:  "Mr. " + u.Name, //名前の先頭にMr. と付加する
		Email: u.Email,
	})
	return v, err
}

// Unmarshalのカスタマイズ
// こちらは引数を取り値の変更を行う為にレシーバのUserにポインタ型を指定
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		//省略したフィールドは全て型の初期値が入る
		//JSON側に無いフィールド(上で消したEmail)は受け取れない
		//JSON側の型に合わせて受け取る必要がある(上で文字列に変えたID)
		ID   string //一旦stringで受け取り、intにカスタマイズする
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}

	//レシーバのUserがポインタ型でないと、以下のフィールドの変更が反映されない
	i, _ := strconv.Atoi(u2.ID) //JSON側で文字列のIDをint型に変換
	u.ID = i
	u.Name = u2.Name + "様"
	//上記以外は構造体の各フィールドの初期値が入る
	//EmailはMarshalの際に消したので、そもそもJSON側にはデータが無い
	return err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "testuser1"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshl:JSONに変換してbyteのsliceとして受け取る
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	//JSON形式の文字列型に変換して表示
	fmt.Println(string(bs))

	//------------------------------

	//JSON形式から構造体への変換

	//先ほど生成されたbyteのsliceの型を確認
	fmt.Printf("%T\n", bs)

	//別のUserを追加で生成(u2は最初からポインタ型のアドレス)
	u2 := new(User)

	//Unmarshal: JSONを構造体に変換
	//第二引数には構造体のアドレスを渡す
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println("err:", err)
	}

	//構造体に戻っているか確認
	fmt.Println("u2:", u2)

}

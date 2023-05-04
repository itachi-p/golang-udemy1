package main

import "fmt"

// 構造体(structure)

type User struct {
	Name string
	Age  int
	// X, Y int // 複数を一度に定義も可能
}

func UpdateUser(user User) {
	user.Name = "Siegfried Kircheis"
	user.Age = 10000
}

func UpdateUser_v2(user *User) {
	user.Name = "Siegfried Kircheis"
	user.Age = 999
}

func main() {
	//明示的な定義
	var user1 User
	//宣言しただけだと、各型の初期値が入っている
	fmt.Println(user1)
	user1.Name = "Gopher"
	user1.Age = 14
	fmt.Println(user1)

	//暗黙的な定義
	user2 := User{}
	fmt.Println(user2) //同様に初期値が入る
	user2.Name = "Reinhard von Müsel"
	user2.Age = 20
	fmt.Println(user2)

	user3 := User{Name: "Yang Wen-li", Age: 29}
	fmt.Println(user3)

	//構造体の順番通りならフィールド名の省略も可能
	user4 := User{"Julian Mintz", 14}
	fmt.Println(user4)

	//user5 := User{30, "user5"}
	//fmt.Println(user5)

	user6 := User{Name: "Dusty Attemborough"}
	fmt.Println(user6) // 設定しなかったAgeは初期値のまま

	user7 := new(User) //構造体のポインタ型を返す
	fmt.Println(user7)

	//アドレス演算子（&）を使って宣言
	user8 := &User{} //上記のnew()と同じくポインタ型
	fmt.Println(user8)
	//明示的にポインタ型の構造体を宣言し、関数の引数としてよく使われる

	UpdateUser(user1)
	UpdateUser_v2(user8)
	fmt.Println(user1) //値渡しの為に元の本体は更新されない
	fmt.Println(user8) //参照渡しにより更新される
}

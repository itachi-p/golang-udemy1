package main

import "fmt"

// 構造体のメソッド
// 任意の型に特化した関数定義の仕組み

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	//レシーバーとして宣言した変数を内部で使える
	fmt.Println(u.Name)
}

func (u User) setName(name string) {
	u.Name = name
}

// 原則メソッドのレシーバはポインタ型にしておくのが推奨
// 呼び出し側はポインタ型にする必要はない
func (u *User) setName_v2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "Wilibard Joachim von Merkatz"}
	user1.SayName()

	user1.setName("Neithardt Muller")
	user1.SayName() //値渡しでは本体が更新されない

	user1.setName_v2("Neithardt Muller")
	user1.SayName()

	//呼び出し側は値型、ポインタ型どちらで引数を渡してもどちらでもよい
	user2 := &User{Name: "user2"}
	user2.setName_v2("Paul von Oberstein")
	user2.SayName()

}

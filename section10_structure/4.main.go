package main

import "fmt"

// 構造体：コンストラクタ(関数)

type User struct {
	Name string
	Age  int
}

// User型のポインタを生成するコンストラクタ関数
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("Walter von Schonkopf", 32)
	fmt.Println(user1)  //ポインタ型
	fmt.Println(*user1) //実体にアクセス
}

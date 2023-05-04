package main

import "fmt"

// 構造体：slice

type User struct {
	Name string
	Age  int
}

// 構造体Userのポインタをsliceとして格納できる構造体
type Users []*User

//下記のようにも定義できるが、上記の方が望ましい。
/*
type Users struct {
	Users []*Users
}
*/

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	//構造体のスライス作成
	users := Users{}
	//要素の追加
	users = append(users, &user1)
	users = append(users, &user2)
	//可変長引数なので複数一度に追加もできる
	users = append(users, &user3, &user4)

	fmt.Println(users) //中身はポインタ型なのでアドレスの列挙になる
	for _, u := range users {
		fmt.Println(*u)
	}

	//make関数を使って構造体のスライスを生成することもできる
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)
	for _, u := range users2 {
		fmt.Println("users2:", *u)
	}
}

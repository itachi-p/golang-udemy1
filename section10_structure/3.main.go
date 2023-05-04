package main

import "fmt"

// 構造体の埋め込み

type T struct {
	//User   User //フィールド名:型
	//埋め込み型の場合は型名を省略することもできる
	User   // 型名を省略したフィールド名
	Number int
}

type User struct {
	Name string
	Age  int
}

func (u *User) setName() {
	u.Name = "Wolfgang Mittermeier"
}

func main() {
	t := T{User: User{Name: "Reuentahl", Age: 29}, Number: 2}

	fmt.Println(t) //二重の入れ子データ構造として表示される
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	//型名を省略した場合のみ直接のアクセスも可能になる
	fmt.Println(t.Name)

	//内側のメソッドも使用可能
	t.User.setName()
	fmt.Println(t)

}

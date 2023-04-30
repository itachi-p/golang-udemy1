package main

import "fmt"

// 関数を引数に取る関数

func CallFunction(f func()) {
	f()
}

func main() {
	// 関数の呼び出し側で無名関数の中身を定義
	CallFunction(func() {
		fmt.Println("I'm a function.")
	})
}

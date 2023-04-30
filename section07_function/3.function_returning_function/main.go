package main

import "fmt"

// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}

func main() {
	// 戻り値として関数が返ってくるのを変数で受け取る
	f := ReturnFunc()
	f()
	// 戻り値として関数を受け取った変数の型を調べるとfunc()
	fmt.Printf("%T\n", f)
}

package main

import "fmt"

// クロージャー(関数閉包)の実装

// 関数を返す関数を定義
func Later() func(string) string {
	var store string
	// 内部の無名関数(これがクロージャー)
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	// 関数を呼び出し、その戻り値である内部の無名関数を変数で受け取る
	f := Later()
	// 呼び出した関数の戻り値として受け取った内部無名関数に引数を渡す
	fmt.Println(f("Hello")) // 最初の実行では""(空文字)だけが戻される
	fmt.Println(f("My"))    //この時点で"Hello"と表示される
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Gopher")) // この値は代入はされるが戻り値は1つ手前
}

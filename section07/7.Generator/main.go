package main

import "fmt"

// ジェネレーター
/* 何らかのルールに従い、連続した値を返し続ける仕組みの事。
PythonやPHPには機能として備わっている。
Goではクロージャーの応用によりジェネレーターのように振る舞う機能を実現できる。*/

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
func main() {
	ints := integers()

	// 実行される度に+1インクリメントされるジェネレーターが生成される
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

}

package main

import "fmt"

// 無名関数 (簡易的に関数定義したい場合に使われる)

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	// 無名関数を定義し、そのまま実行する場合
	i2 := func(x, y int) int {
		return x + y
	}(3, 2)
	fmt.Println(i2)
}

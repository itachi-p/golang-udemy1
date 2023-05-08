package main

import "fmt"

// slice : 可変長引数
func Sum(s ...int) int { // 引数の数が限定されない
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(Sum())

	// スライスを作って引数に渡すこともできる
	sl := []int{1, 1, 2, 3, 5, 8, 13}
	fmt.Println(Sum(sl...))
}

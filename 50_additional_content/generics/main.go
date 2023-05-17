package main

import "fmt"

//Generics
/* 型を抽象化してコードの再利用性を向上させつつ、
静的型付け言語の持つ型安全性を維持する機能 */

func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func main() {
	PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

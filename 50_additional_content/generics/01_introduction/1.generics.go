package main

import (
	"fmt"
)

//Generics(Golang1.18〜)
/* 型を抽象化してコードの再利用性を向上させつつ、
静的型付け言語の持つ型安全性を維持する機能 */

// 従来の方法では同じような処理でも型ごとに別の関数が必要
/*
func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}
func PrintSliceStrings(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
*/

// Genericsを用いて型を抽象化することにより、コードを再利用
// any型: interface{}型のalias
func PrintSlice[T any](s []T) {
	//同じ中身で複数の型の処理が可能
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSliceStrings([]string{"a", "b", "c", "d", "e"})

	//同じ関数に違う型を渡して実行できる
	// PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSlice[string]([]string{"a", "b", "c", "d", "e"})
	//型推論により、[int]や[string]の部分は省略も可能
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice([]string{"a", "b", "c", "d", "e"})

}

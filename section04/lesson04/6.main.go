package main

import "fmt"

// 配列型
/* 他の言語の配列型と違って、後で要素数の追加ができない。
　　それを可能にする仕組みがスライス。
*/

func main() {
	var arr1 [3]int
	// 初期化していないので全ての要素がintの初期値0
	fmt.Println(arr1)
	// 型を表示すると[n]intとnの値が不変の固定長配列として表示される。
	fmt.Printf("%T\n", arr1)
}

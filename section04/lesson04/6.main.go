package main

import "fmt"

// 配列型
/* 他の言語の配列型と違って、後で要素数の追加ができない（固定長配列）。
　　それを可能にする仕組みがスライス型（=可変長配列）。
*/

func main() {
	var arr1 [3]int
	// 初期化していないので全ての要素がintの初期値0
	fmt.Println(arr1)
	// 型を表示すると[n]intのように、nの値が不変の固定長配列として表示される。
	fmt.Printf("%T\n", arr1)

	// 初期値を与えて明示的に配列を宣言する
	// (要素が2つなので、3つ目の要素にはstringの初期値""(空文字)が入る)
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的な宣言と同時に初期値を与える
	arr3 := [3]int{3, 5, 8}
	fmt.Println(arr3)

	// 配列の要素数を省略した定義
	arr4 := [...]string{"X", "Y", "Z"}
	fmt.Println(arr4) // 初期値によって自動で要素数が決まり、型も確定する。
	fmt.Printf("%T\n", arr4)

	// 配列の要素の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	// 配列の要素の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	//arr5 = arr1 // 同じint型の配列でも要素数が違えば型が違うので代入不可
	//fmt.Println(arr5)

	// 配列の要素数を数える関数
	fmt.Println(len(arr2))

}

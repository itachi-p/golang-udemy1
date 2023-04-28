package main

import "fmt"

// interface型 & nil
// 詳しくはsection6の型スイッチ及びsection12を参照

func main() {
	var x interface{} // 全ての型と互換性のある特殊な型
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "X"
	fmt.Println(x)

	x = 3.14159
	fmt.Println(x)

	x = [...]int{1, 1, 2, 3, 5, 8}
	fmt.Println(x)

	x = 2
	//fmt.Println(x + 3) // 全ての型を代入・表示はできるが、計算はできない

}

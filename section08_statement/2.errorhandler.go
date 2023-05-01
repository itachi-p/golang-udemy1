package main

import (
	"fmt"
	"strconv"
)

// if 条件分岐
// エラーハンドリング

func main() {
	var s string = "100"
	var s2 string = "ABC"

	// 2つ目の引数を破棄する場合
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T\n", i)

	// エラーメッセージを受け取る場合
	i2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i2)
}

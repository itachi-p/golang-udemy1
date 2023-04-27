package main

// 変数

import (
	"fmt"
	"time"
)

func main() {
	var nowMonth = time.Now().Month()
	var today int = time.Now().Day()
	fmt.Printf("今日は%d月%d日です。\n", nowMonth, today)

	// 複数同時定義
	var x, y, z int = 3, 5, 13
	var t, f bool = true, false

	fmt.Print(x, y, z)
	fmt.Println(t, f)

	// 異なる型を一度に定義
	var (
		age  int    = 88
		name string = "八十八"
	)
	fmt.Printf("彼は%sさんで、年齢は%d歳です。\n", name, age)

	// 暗黙的な型指定
	num := 100
	str := "009"
	fmt.Println(str, num)
}

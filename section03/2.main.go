package main

// 変数の宣言と初期化、代入

import (
	"fmt"
	"time"
)

func main() {
	var nowMonth = time.Now().Month()
	var today int // 宣言だけで同時に初期化しない場合、型の明示は必須
	today = time.Now().Day()
	fmt.Printf("今日は%d月%d日です。\n", nowMonth, today)

	// 複数同時定義
	var x, y, z int = 3, 5, 13
	var t, f bool = true, false

	fmt.Print(x, y, z, "\n") // ダブルクォートで囲んだ\nは改行コードとして処理される
	fmt.Println(t, f, '\n')  // シングルクォートで囲んだ\nは対応するASCIIコード番号が表示される

	// 異なる型を一度に定義
	var (
		age  int    = 88
		name string = "八十八"
	)
	fmt.Printf("彼は%sさんで、年齢は%d歳です。\n", name, age)

	// 暗黙的な型指定(宣言と同時に初期値を与える場合のみ)
	num := 100
	str := "009"
	fmt.Println(str, num)
}

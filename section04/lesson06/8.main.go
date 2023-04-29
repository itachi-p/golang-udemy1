package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	var i int = 1
	fl64 := float64(i) // intからfloatへの変換
	fmt.Println(fl64)
	fmt.Printf("iの型: %T\n", i)
	fmt.Printf("fl64の型: %T\n", fl64)

	fl := 3.14159
	yutori := int(fl)
	fmt.Printf("\ni2の型: %T\n", yutori)
	// 浮動小数点型から整数型へのキャストは四捨五入ではなく小数点以下切り捨て
	fmt.Printf("flの値: %f\ni2の値: %d\n\n", fl, yutori)

	// 文字列型から数値型への変換
	var str_num string = "100"
	fmt.Println(str_num)
	fmt.Printf("str_numの型: %T\n", str_num)

	// 標準関数からの戻り値が2つある場合に2つ目を破棄する書き方
	// Golangでは定義した変数は必ず使うルールをこれで回避できる
	int_num, _ := strconv.Atoi(str_num)
	fmt.Println(int_num)
	fmt.Printf("int_numの型: %T\n\n", int_num)

	// 上記の型変換関数Atoi()の2つ目の戻り値を破棄しない場合
	int_num2, err_msg := strconv.Atoi("1２3") // 数値に変換できない文字列
	if (err_msg) != nil {
		fmt.Println(err_msg)
	}
	fmt.Println(int_num2) //エラー発生の場合、int_num2にはint型の初期値0が入る

}

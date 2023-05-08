package main

import (
	"fmt"
	"strconv"
)

//strconv : 文字列変換

func main() {
	//真偽値の型を文字列("true"|"false")に変換
	bt := true
	fmt.Printf("%T\n", bt)
	fmt.Printf("%T\n\n", strconv.FormatBool(bt))

	//整数を文字列に変換
	i := strconv.FormatInt(-100, 10) //第2引数:基数(n進数)
	fmt.Printf("%v, %T\n", i, i)
	// 上記を簡易的に変換（10進数固定）
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n\n", i2, i2)

	//浮動小数店型を文字列に変換
	//第2引数:書式指定子(fmtパッケージの%f等のfと同じ)
	//第3引数:桁数の制限ルール -1の場合必要な桁数が自動的に設定される
	//第4引数:文字列化する浮動小数点数の精度（ビット数）
	//実数表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	//小数点以下2位まで
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	//指数表現による文字列化（小数点以下2位まで）
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))

	//文字列から基本型への変換
	//文字列を真偽値に変換
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("\n%v, %T\n", bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6) //全て通る(true)
	//2番目の戻り値はerror型なので、エラーハンドリングも可能
	bf1, ok := strconv.ParseBool("false")
	if ok != nil {
		fmt.Println("Convert Error")
	}
	fmt.Printf("%v, %T\n", bf1, bf1)
	bf2, _ := strconv.ParseBool("0")
	bf3, _ := strconv.ParseBool("f")
	bf4, _ := strconv.ParseBool("F")
	bf5, _ := strconv.ParseBool("FALSE")
	bf6, _ := strconv.ParseBool("False")
	fmt.Println(bf2, bf3, bf4, bf5, bf6)

	//文字列を整数型に変換
	//第2引数:基数(n進数)
	//第3引数:精度（ビット数）0を指定するとGoのInt型の精度
	i3, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("\n%v, %T\n", i3, i3)
	i4, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Printf("%v, %T\n", i4, i4)

	//簡易的に変換可能
	i5, _ := strconv.Atoi("-2147483648")
	fmt.Println(i5)

	//文字列を浮動小数点型に変換
	fl1, _ := strconv.ParseFloat("3.14159", 64)
	fl2, _ := strconv.ParseFloat(".2", 64)
	fl3, _ := strconv.ParseFloat("-2", 64)
	fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	fl5, _ := strconv.ParseFloat("0.1234E8", 64)

	fmt.Println(fl1, fl2, fl3, fl4, fl5)
}

package main

import "fmt"

// 独自型
// 構造体ではないが、特有の機能を持った独自の型を定義できる

type MyInt int

// MyInt型用のメソッドも作れる
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi) //MyInt型になっている

	//i := 100
	//fmt.Println(mi + i) //通常のintとは型が違うのでエラーになる

	mi.Print()
}

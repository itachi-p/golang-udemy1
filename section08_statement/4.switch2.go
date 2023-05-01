package main

import "fmt"

// switch文
// 型switch

// 全ての型と互換性のあるinterface型を使って型を調べる
func anything(a interface{}) {
	/* どんな型でも受け取れるが、全てinterface型になる為に
	元の型は失われ、そのままでは計算などはできなくなる。
	そこで型アサーションを用いて型を復元する。 */
	fmt.Println(a)
}

func main() {
	anything("aaa") // 文字列型
	anything(1)     // 整数型

	// 型アサーション
	// 動的に変数の型をチェックする機能
	var x interface{} = 3
	// 書式は　変数名.(復元したい型)
	i := x.(int) // int(x)では型変換できない
	// fmt.Println(x + 5) // interface型とint型なので計算できない
	fmt.Println(i + 5)

	/* 以下は本来int型のところをfloat型に復元しようとしているので
	ランタイムエラーが発生し、プログラムの実行が強制終了される。
	f := x.(float64)
	fmt.Println(f)
	*/

	// 以下のようにするとランタイムエラーにならず実行可能
	f, isFloat64 := x.(float64)
	// 型の変換には失敗してもエラー終了はしない
	fmt.Println(f, isFloat64)

}

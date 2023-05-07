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

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt { // 右辺はisInt == trueの省略形
		fmt.Println(i, "x is integer")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know.")
	}

	// 型によるswitch文
	// if文による型アサーションよりも簡単に書けるので推奨
	var y interface{} = "A"

	// 変数yの型を調べ、処理を分ける
	switch y.(type) {
	case int:
		fmt.Println("type: int")
	case string:
		fmt.Println("type: string")
	default:
		fmt.Println("I don't know.")
	}

	// 復元した値を使いたい場合
	switch v := y.(type) {
	case bool:
		fmt.Println(v, "is boolean")
	case int:
		fmt.Println(v, "is integer")
	case string:
		fmt.Println(v, "is string")
	}
	/* 上記をそのまま関数anything()の中身にすれば、
	受け取った引数の型によって処理を分ける関数が書ける。 */
}

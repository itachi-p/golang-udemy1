package main

import "fmt"

//Generics: typesets
//インターフェースを満たす任意の型の「集合」を制約として持たせる

type Number interface {
	//いずれかの数値型であればNumberのインターフェースを満たす
	~int | int32 | int64 | float32 | float64
}

func Max[T Number](x, y T) T {
	//整数型及び浮動小数点型であれば同じ処理をする
	if x >= y {
		return x
	}
	return y
}

// 独自型を設定した場合、元となる型がintであっても
// 「Number型ではない」というエラーになるが、~(チルダ)を付ければ通る
type MyInt int

func main() {
	//型推論が働くため、実際は[int]や[float]の部分は省略も可能
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[float64](1.1, 1.3))

	var x, y MyInt = 3, 4
	fmt.Println(Max[MyInt](x, y))
}

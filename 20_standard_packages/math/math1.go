package main

import (
	"fmt"
	"math"
)

//math

func main() {
	//数学的な定数
	//円周率
	fmt.Println(math.Pi)
	//2の平方根
	fmt.Println(math.Sqrt2)

	//数値型に関する定数
	//int16で表現可能な最大値
	fmt.Println(math.MaxInt16)
	//int32で表現可能な最小値
	fmt.Println(math.MinInt32)
	//float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)
	//float64で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64)

	//mathの関数
	//絶対値
	fmt.Println("\n", math.Abs(-256))
	//累乗(XのY乗)
	fmt.Println(math.Pow(2, 16))
	//平方根、立法根
	fmt.Println(math.Sqrt(3))
	fmt.Println(math.Cbrt(8))
	//最大値、最小値
	fmt.Println(math.Max(5, 8))
	fmt.Println(math.Min(5, 8))

	//小数点以下の切り捨て、切り上げ
	//数値の正負を問わず小数点以下を単純に切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))
	//引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))
	//引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}

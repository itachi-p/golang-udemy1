package main

import (
	"fmt"
	"math"
)

func main() {

	//any: type any = interface{}
	//Go1.18以降は従来のinterface型の別名としてanyが定義されている
	var a any
	var b interface{} //同じ意味になる

	//どのような型でも代入できる
	a, b = true, -255
	a, b = 3.14159, nil
	a, b = 3e-3+4i, math.MinInt32
	fmt.Println(a, b)
}

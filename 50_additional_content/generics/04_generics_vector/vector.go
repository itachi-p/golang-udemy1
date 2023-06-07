package main

import "fmt"

//Generics: vector
//型定義でも型パラメータを使ってみる
//C++のVector型のようなものをGoのSliceとして実装

type Vector[T any] []T

type IntVector = Vector[int]

func main() {
	//任意の型を指定し、Sliceを生成できる
	var v Vector[int] = Vector[int]{10, 20, 30}
	fmt.Println(v)

	var v2 Vector[float64] = Vector[float64]{1.2, 3.4, 5.6}
	fmt.Println(v2)

	v3 := IntVector{1, 2, 3}
	fmt.Println(v3)
}

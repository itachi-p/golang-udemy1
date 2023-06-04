package main

import (
	"fmt"
	"strconv"
)

//Generics(Golang1.18〜)
//interface型パラメータに制約を持たせる

// any型は任意の型をなんでも受け取れるが、他にも様々なinterfaceが指定可能
// fmt.Stringer(): string型を返すString()というメソッドを実装すればよい
func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		//未定義のT型のスライスxsの要素である変数xは必ずString()を持つ
		//そのString()メソッドの中身の実装は後でよい（これこそが利点）
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
}

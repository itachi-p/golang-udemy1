package main

import (
	"fmt"
	"strconv"
)

//Generics
//型パラメータに制約を持たせることもできる

// any型は任意の型をなんでも受け取れるが、他にも様々なinterfaceが指定可能
// fmt.Stringer(): string型を返すString()という関数を実装すればよい
func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
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

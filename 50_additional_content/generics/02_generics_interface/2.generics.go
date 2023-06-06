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

/* 独自のインターフェースを追記してみる
//この部分はあってもなくても同様に動作する
type Stringer interface {
	String() string
}
*/

// 独自の型を定義(int型のAlias)
type MyInt int

// MyInt型に対してString()メソッドを実装する
func (i MyInt) String() string {
	//中身はなんでもいいが、とりあえずMyInt型の値を文字列に変換する
	return strconv.Itoa(int(i))
}

func main() {
	//実行結果はパッと見は区別付かないが、String型のSliceになっている
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
	fmt.Printf("%T", f([]MyInt{1, 2, 3, 4}))
}

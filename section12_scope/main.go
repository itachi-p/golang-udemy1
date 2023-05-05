package main

//パッケージ名は省略可能
//ただし、同名メソッドが含まれている場合はエラーになるので注意
//いずれもあまり推奨されないやり方で、省略しないのがベター。
import (
	"fmt"
	. "fmt" //いきなりPrintln()だけで使用できる。
	f "fmt" // fmtをf.Println()のように省略できる。
	"golang_udemy1/section12_scope/foo"
)

// スコープ
// public, private, パッケージ分割
/* publicやprivateという識別子を明示的に付けるのでなく、
パッケージ外の定数の頭文字が大文字であればpublic,
小文字であれば同一パッケージ内でのみ参照できるprivateになる。
後者もメンバ変数に直接ではなく、関数を介すればアクセス可能(getterとsetter)*/

// 関数のスコープ
func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//上で戻り値として変数bを定義しているので、変数bの再定義でエラーとなる
	//var b string = s
	b = s
	{
		b := "BBBB"
		//変数bはこの深いブロック内でのみBBBBと表示される
		fmt.Println(b)
	}
	//ここではbの値はAAAに戻る。
	//敢えてこのような重複、同名変数再定義をやる必要性はない。
	fmt.Println(b)
	return b
}

func main() {
	f.Println(foo.MAX)
	//頭文字を小文字で定義した定数はパッケージ外からは直接アクセス不可
	//fmt.Println(foo.min)
	Println(foo.ReturnMin()) //関数を介してなら内部の変数にもアクセス可能

	Println(appName())
	//関数の中で定義された定数や変数に直接アクセスはできない
	//fmt.Println(AppName, Version)

	fmt.Println(Do("AAA"))
}

package main

import "fmt"

// interface型
// カスタムエラー

//goの組み込み型エラーは以下のようにinterface型として纏められている
//これをカスタマイズして独自のエラー型として定義できる
/*
type error interface {
	Error() string
}
*/

// 独自のエラー型を構造体として宣言
type MyError struct {
	Message string
	ErrCode int
}

/*
上記の独自エラー型をレシーバーに持つメソッドとしてError()を定義することで、
goのソースコードのerrorインターフェース型と共通の性質を持つ型としてMyErrorが認識される
*/
func (e *MyError) Error() string {
	return e.Message
}

// エラーを発生させる関数
func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	//カスタマイズしたエラーメッセージが表示される
	fmt.Println(err.Error())

	/* ただし、errはあくまでもerror型の変数なので、
	MyError型で定義されている各フィールドには直接アクセスできない */
	//fmt.Println(err.Message)

	/* 上記のようにMyError型の各フィールドにアクセスするには、
	型アサーションでやったように元の型に復元する必要がある。*/
	e, ok := err.(*MyError)
	//型アサーションの結果、戻り値がMyError型の実体ならok==true
	if ok {
		fmt.Println(e.ErrCode, e.Message)
	}

}

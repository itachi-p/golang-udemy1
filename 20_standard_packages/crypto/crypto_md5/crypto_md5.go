package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//crypto : 暗号化
//MD5

func main() {
	//MDSハッシュ値を生成
	//任意の文字列からMD5ハッシュ値を生成する処理例
	h := md5.New()

	//文字列を書き込む
	io.WriteString(h, "ABCDE")

	//ハッシュ値のバイト配列を表示
	fmt.Println(h.Sum(nil))

	//ハッシュ値のバイト配列を16進数の文字列に変換
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}

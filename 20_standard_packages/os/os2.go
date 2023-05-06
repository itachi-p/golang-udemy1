package main

import (
	"fmt"
	"log"
	"os"
)

//os

func main() {
	//ファイル操作
	//OpenFile
	/*フラグ一覧
	O_RDONLY 読み込み専用
	O_WRONLY 書き込み専用
	O_RDWR 読み書き可能
	O_APPEND ファイル末尾に追記
	O_CREATE ファイルが存在しなければ作成
	O_TRUNC 可能であればファイルの内容をオープン時に空にする
	*/

	//os.Openfile(<対象ファイル>, <フラグ(複数指定可能)>, <パーミッション>)
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	//確保するバイト数は小さ過ぎず大き過ぎない適切範囲が望ましい
	bs := make([]byte, 16)
	n, err := f.Read(bs)
	fmt.Println(n) //タブ文字(\t)も改行と同じく1文字で数えられる
	fmt.Println(bs)
	fmt.Println(string(bs))

}

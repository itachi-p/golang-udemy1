package main

import (
	"fmt"
	"log"
	"os"
)

//log

func main() {
	//ログの出力先を変更
	log.SetOutput(os.Stdout)

	//任意のログメッセージと、デフォルトで日付時刻情報が出力される
	log.Printf("Log%4d\n", 3)   //桁数指定&半角スペース埋め
	log.Printf("Log%04d\n", 12) //桁数指定&0埋め

	//log.Fatalf("Fatal error%d\n", 003)

	//log.Panicln("Panic!")

	//任意のファイルを作成し、出力先に指定
	f, err := os.Create("log/log_output/test.log")
	if err != nil {
		log.Fatalln(err)
	}
	//作成したファイルを出力先に設定
	fmt.Printf("%T\n", f)
	log.SetOutput(f)
	log.Println("ログをファイルに書き込み")

	//ログのフォーマット指定
	//再度出力先を標準出力(コマンドプロンプト)に設定
	log.SetOutput(os.Stdout)

	//標準のログフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("標準フラグ")

	//マイクロ秒(μs)を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("マイクロ秒まで表示")

	//時刻とファイルの行番号を表示（フルパス）
	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("ファイルの実行行番号を出力（フルパス）")

	// ログのプリフィックスを指定（時刻より前の先頭に表示される）
	log.SetPrefix("[LOG]")
	//時刻とファイルの行番号を表示（短縮形）
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("プリフィックス表示+時刻+行番号（短縮形）")
}

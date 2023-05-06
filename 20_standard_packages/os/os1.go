package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//標準パッケージ
//os

func main() {
	//Exit
	//任意で終了ステータスを指定して終了
	/*
		os.Exit(1)
		//以下は実行されない
		fmt.Println("Start")
	*/

	//defer文も破棄され実行されない点に注意
	/*
		defer func() {
			fmt.Println("defer")
		}()
		os.Exit(0)
	*/

	//log.Fatal
	/*
		//ファイル不存在の際にエラー終了させる
		_, err := os.Open("A.txt") //存在しないファイル
		if err != nil {
			log.Fatalln(err)
		}
	*/

	//Args
	//C言語と同じく、実行時のコマンドライン引数の0番目は実行ファイル名が自動で入る
	/*
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])

		//os.Argsの要素数を表示
		fmt.Printf("length=%d\n", len(os.Args))
		//os.Argsの中身を全て表示
		for i, v := range os.Args {
			fmt.Println(i, v)
		}
	*/

	//ファイル操作
	//Create
	//新規ファイル作成
	f, _ := os.Create("test.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6) //offset位置を6文字目に指定
	//Go1.7以降os.SEEK_ENDは非推奨。推奨はio.SeekEnd
	//f.Seek(0, os.SEEK_END) //ファイルの末尾にoffsetを移動
	f.Seek(0, io.SeekEnd)
	f.WriteString("Yaah")

	//Open, Read
	//ファイルが存在すれば何も起きない
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//リソースの解放
	//deferで確実にファイルがクローズされる。
	defer f.Close()

	//make関数でbyteのsliceとして文字列の入れ物を作成
	bs := make([]byte, 128)

	//ファイルを読み込んだ内容をbsに書き込み
	n, err := f.Read(bs)
	//読み込んだバイト数(とエラーハンドリングの結果)を表示
	//改行(\n)もバイト数1の6文字として計算される
	fmt.Println(n, err)
	//ファイルを読み込んだ内容を表示
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	//offset位置を10文字目に指定して読み込む
	nn, _ := f.ReadAt(bs2, 10)
	//offset位置10から読み込んだので結果は6文字
	fmt.Println(nn)
	fmt.Println(string(bs2))

}

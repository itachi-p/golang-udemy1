package main

import (
	"log"
	"os"
)

//log2

func main() {
	//ロガーの生成
	/*logパッケージの関数はデフォルト設定の1つのロガーを全体に適用しているので、
	一部のログの出力方法を変えたい場合、新規で生成する。
	通常のログは標準出力、特別なエラーはファイルを残すなど。*/
	logger := log.New(os.Stdout, "{prefix}", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message1")
	log.Println("message2")

	//条件分岐。エラーで終了させる。
	_, err := os.Open("nonexistent.file")
	if err != nil {
		//ログ出力
		//log.Fatalln("[Exit]", err)
		logger.Fatalln("[Exit]", err)
	}

}

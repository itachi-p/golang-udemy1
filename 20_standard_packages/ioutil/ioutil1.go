package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//ioutil
//osパッケージで代用可能
//io/ioutil はGo1.19から非推奨

func main() {

	//入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	//読み込んだ内容をそのままファイルに書き込み(要はコピー)
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}

}

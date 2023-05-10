package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

//net/http
//クライアント

func main() {
	//GETメソッド
	//戻り値の*http.Response型は様々なHTML情報の構造体になっている
	res, _ := http.Get("https://example.com")
	fmt.Printf("%T\n", res)
	fmt.Println(res)

	fmt.Println("\nStatusCode:", res.StatusCode)
	fmt.Println("Protocol:", res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])
	fmt.Println(res.Request.Method) //GETメソッド
	fmt.Println(res.Request.URL)

	//リソース解放処理
	defer res.Body.Close()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//HTML全体を読み込み
	/*"io/ioutil" は Go 1.19 から非推奨
	ioutilの既存の型と関数は全てioパッケージとosパッケージに分割
	body, _ := ioutil.ReadAll(res.Body) */
	body, _ := io.ReadAll(res.Body)
	fmt.Print(string(body))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//-----------
	//POSTメソッド(フォームの送信)

	//POSTパラメータを組み立てる
	vs := url.Values{}

	//POSTに送信するデータを生成main
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode()) //日本語を含むパラメータをエンコード

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body2, _ := io.ReadAll(res.Body)
	fmt.Print(string(body2))
}

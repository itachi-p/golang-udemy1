package main

import (
	"html/template"
	"log"
	"net/http"
)

//net/http
//サーバー

type MyHandler struct{}

/*
golangにおけるハンドラーとは、ServeHTTPメソッドを持ったinterface
http.ResponseWriter, *http.Requestを引数に取るメソッドを持つもの全て

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
*/

func home(w http.ResponseWriter, r *http.Request) {
	//ローカルに用意したHTMLファイルのテンプレートを解析し、その構造体を生成
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	//テンプレートファイルの{{ . }}の中に渡したデータが組み込まれ表示される
	t.Execute(w, "ホーム画面")
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "ログイン画面(テスト用ダミー)")
}

func main() {

	//ハンドラー関数: URLごとに違うページ処理をさせる
	//第1引数にURL、第2引数に上で設定したハンドラー関数
	//http://localhost:8080/top などのURLでアクセス可能になる
	//"/"と指定すればデフォルトページとなり、404 page not foundにならない
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)

	//http.ListenAndServe
	/* この関数が実行されている間はサーバーが立っている。
	ブラウザでURLに http://localhost:8080/ でアクセス可能
	第2引数にはハンドラーを指定、nilの場合はデフォルトのマルチプレクサを使用
	nilの場合 http://localhost:8080/ だけでは404 page not foundになるが
	ハンドラーを指定すると、URLによらず全て同じ処理になってしまう
	通常はURL毎に処理を変えたいので、nilにしてマルチプレクサを使用する*/
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", &MyHandler{})
}

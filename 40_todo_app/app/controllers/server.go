package controllers

import (
	"fmt"
	"golang_udemy1/40_todo_app/config"
	"html/template"
	"net/http"
)

// 様々なハンドラー関数でテンプレートとしてHTML表示される部分を共通化する関数
// 第3引数は可変長引数を取れるように...<型>とする。実際は/views/templates/%s.htmlの%s部分
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	//第3引数で渡されたfilenamesの中身を取り出し、ファイルパスを加えfiles[]に格納
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	//可変長引数で受け取ったファイルパス群からテンプレートのキャッシュを生成
	templates := template.Must(template.ParseFiles(files...))
	//上記で生成されたテンプレートのキャッシュに対し、適用するレイアウトを指定し実行
	//第2引数に実行するテンプレートを明示的に渡す(HTML中の{{define "templatename"}})
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	//CSSファイルとJavaScriptファイルを静的(static)ファイルとして読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	/* URLを/staticの下の階層とするが、実際のディレクトリ構成では
	/viewsと/css,/jsの間に/staticは無いのでプリフィックスを取り除く*/
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//ハンドラー(topやsignup)と、それに対応させるURLの登録
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	//サーバ起動: 第2引数にnilを渡すことで、デフォルトのマルチプレクサを使用
	//登録されていないURLへのアクセスはデフォルトで"404 page not found"を返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

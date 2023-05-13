package controllers

import (
	"golang_udemy1/40_todo_app/config"
	"net/http"
)

func StartMainServer() error {
	//CSSファイルとJavaScriptファイルを静的(static)ファイルとして読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	/* URLを/staticの下の階層とするが、実際のディレクトリ構成では
	/viewsと/css,/jsの間に/staticは無いのでプリフィックスを取り除く*/
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//ハンドラー(top)と、それに対応させるURLの登録
	http.HandleFunc("/", top)
	//第2引数はnilを渡すことで、デフォルトのマルチプレクサを使用
	//登録されていないURLへのアクセスはデフォルトで"404 page not found"を返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// 固定の2つの引数を取るハンドラー関数
// http.HandleFuncにて指定するURLで開かれるページの設定
func top(w http.ResponseWriter, r *http.Request) {
	//テンプレートファイルの解析
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	//テンプレートファイルの実行
	t.Execute(w, "Hello")
}
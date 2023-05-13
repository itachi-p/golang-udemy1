package controllers

import (
	"golang_udemy1/40_todo_app/config"
	"net/http"
)

func StartMainServer() error {
	//ハンドラー(top)と、それに対応させるURLの登録
	http.HandleFunc("/", top)
	//第2引数はnilを渡すことで、デフォルトのマルチプレクサを使用
	//登録されていないURLへのアクセスはデフォルトで"404 page not found"を返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

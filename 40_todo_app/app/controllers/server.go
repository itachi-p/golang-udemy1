package controllers

import (
	"golang_udemy1/40_todo_app/config"
	"net/http"
)

func StartMainServer() error {
	//第2引数はnilを渡すことで、デフォルトのマルチプレクサを使用
	//登録されていないURLへのアクセスはデフォルトで"404 page not found"を返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

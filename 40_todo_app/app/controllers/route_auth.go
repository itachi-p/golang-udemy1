package controllers

import (
	"golang_udemy1/40_todo_app/app/models"
	"log"
	"net/http"
)

/* /signup/でアクセスするページのハンドラー関数 */
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		//入力フォームの値が送信された場合の処理を記述(ユーザー登録)
		err := r.ParseForm() //入力フォームの解析
		if err != nil {
			log.Fatalln(err)
		}
		//入力フォームの値からユーザーの構造体を新規作成
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		//データベースに登録(INSERT)
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		//ユーザー登録が成功したらトップページにリダイレクトさせる
		//第3引数にリダイレクト先、第4引数にステータスコードを指定
		//http.Redirect(w, r, "/", 302)
		//ステータスコードを302のように数値リテラルで指定せず、定数の使用が推奨
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

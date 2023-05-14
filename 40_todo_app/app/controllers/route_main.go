package controllers

import (
	"log"
	"net/http"
)

// 固定の2つの引数を取るハンドラー関数
// http.HandleFuncにて指定するURLで開かれるページの設定
func top(w http.ResponseWriter, r *http.Request) {
	//Cookieを取得し、ログイン状態では表示されないようにする
	_, err := session(w, r)
	if err != nil {
		//様々なハンドラー関数で使うHTMLページ生成の共通関数でトップページを表示
		generateHTML(w, "Hello, Gopher!", "layout", "public_navbar", "top")
	} else {
		//既にログインしている場合はindex.htmlに飛ばす
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

// ログイン成功後のホーム画面用ハンドラー
func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		//エラーがあればログイン状態ではないのでトップページに飛ばす
		http.Redirect(w, r, "/", http.StatusFound)
	} else { //セッションが存在すれば、indexページを表示する

		//セッションを使ってユーザー情報を取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//このユーザーが作成したToDo一覧を取得
		todos, _ := user.GetTodosByUser()
		user.Todos = todos

		//ログイン状態であればtop,signup,login画面とそのリンクは表示しない
		//第2引数をnilではなくuserを渡し、画面にユーザーの持つ情報を表示する
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

// 新規ToDo作成画面
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r) //sessinの値は使わないのでエラーハンドリングのみ
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// 新規ToDo保存（画面表示は必ずしも不要）
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//POSTメソッドで送信されたフォームの値を取得
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		//ToDoの新規登録が完了したらToDo一覧画面に戻る
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

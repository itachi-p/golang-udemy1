package controllers

import "net/http"

/* /signup/でアクセスするページのハンドラー関数 */
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "signup")
}

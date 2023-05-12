package models

import (
	"log"
	"time"
)

//usersテーブルにuserを追加

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// User構造体のポインタをレシーバーとして受け取るメソッドを定義
func (u *User) CreateUser() (err error) {
	//idは自動で一意な値としてインクリメントされるので指定しない
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// メソッドの名前はCreateではなくIncertTodoでもいいかもしれない
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (
		content,
		user_id,
		created_at) VALUES (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Todoデータ1件の取得関数
func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos
	WHERE id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

// 全件取得関数
func GetTodos() (todos []Todo, err error) {
	cmd := "SELECT id, content, user_id, created_at FROM todos"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		//1行づつスキャンしてTodo構造体に格納
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// 特定のユーザーに紐付いた全Todoの取得メソッド
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos
	WHERE user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// 更新: Todo構造体のポインタをレシーバに持つメソッドとして定義
func (t *Todo) UpdateTodo() error {
	cmd := `UPDATE todos SET content = ?, user_id = ?
	WHERE id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

package main

import (
	"fmt"
	"golang_udemy1/40_todo_app/app/models"
)

func main() {
	/*
		fmt.Println("Port:", config.Config.Port)
		fmt.Println("SQLDriver:", config.Config.SQLDriver)
		fmt.Println("DbName:", config.Config.DbName)
		fmt.Println("LogFlie:", config.Config.LogFile)

		log.Println("test")
	*/

	//特に意味はないが、init()関数を呼び出す為に適当に記述
	fmt.Println(models.Db)

	/*
		//テストユーザーをUser構造体のポインタ型として生成
		u := &models.User{}
		u.Name = "テスト三郎"
		u.Email = "test3@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/
	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "TestName2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/
	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("First Todo")
			td, _ := models.GetTodo(1)
			fmt.Println(td)
		user, _ := models.GetUser(4)
		user.CreateTodo("5th Todo")
		   todos, _ := models.GetTodos()

		   	for _, v := range todos {
		   		fmt.Println(v)
		   	}
	*/
	/*
		//特定のuserだけの全Todosを列挙
		user2, _ := models.GetUser(4)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/
	td, _ := models.GetTodo(1)
	td.Content = "Update Todo"
	td.UpdateTodo()
}

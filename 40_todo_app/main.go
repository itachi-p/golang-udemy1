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
		u.Name = "testuser"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)
}

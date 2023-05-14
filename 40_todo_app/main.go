package main

import (
	"fmt"
	"golang_udemy1/40_todo_app/app/models"
	"log"
)

func main() {
	//特に意味はないが、init()関数を呼び出す為に適当に記述
	fmt.Println(models.Db)

	//controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)
}

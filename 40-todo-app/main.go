package main

import (
	"fmt"

	"/Users/itachi-p/go/src/golang-udemy1/40-todo-app/app/controllers"
	"/Users/itachi-p/go/src/golang-udemy1/40-todo-app/app/models"
)

func main() {
	//特に意味はないが、init()関数を呼び出す為に適当に記述
	fmt.Println(models.Db)

	controllers.StartMainServer()

}

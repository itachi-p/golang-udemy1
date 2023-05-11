package main

import (
	"fmt"
	"golang_udemy1/40_todo_app/config"
	"log"
)

func main() {
	fmt.Println("Port:", config.Config.Port)
	fmt.Println("SQLDriver:", config.Config.SQLDriver)
	fmt.Println("DbName:", config.Config.DbName)
	fmt.Println("LogFlie:", config.Config.LogFile)

	log.Println("test")
}

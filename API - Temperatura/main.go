package main

import (
	"API/database"
	"API/server"
	"fmt"
)

func main() {
	fmt.Println("API Rodando")
	database.Conectabanco()

	s := server.NewServer()
	s.Run()
}

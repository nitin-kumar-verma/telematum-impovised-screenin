package main

import (
	"fmt"
	"screening/internal/db"
	"screening/internal/routes"
	"screening/internal/startup"
)

func main() {
	startup.Init()
	//Below line will ensure sqldb conn is closed when app shutsdown
	defer db.SQLDB.Close()

	app := routes.SetupRoutes()
	//Below port can also be moved to envs
	err := app.Listen(":8080")
	if err != nil {
		panic(fmt.Sprintf("Error starting server :%s", err.Error()))
	}
}

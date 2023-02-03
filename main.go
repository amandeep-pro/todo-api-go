
package main

import (
    "github.com/gin-gonic/gin"
    "todo-api/database"
    "todo-api/handlers"
	"todo-api/migrations"
)

func main() {
    database.InitDB()
	migrations.Migrate()
    migrations.Seed()
    r := gin.Default()
    r.GET("/todos", handlers.GetTodos)
    r.Run(":4500")
}

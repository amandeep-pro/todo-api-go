
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
    r.POST("/todos", handlers.CreateTodo)
    r.GET("/todos", handlers.GetTodos)
    r.GET("/todos/:id", handlers.GetTodoByID)
    r.PUT("/todos/:id", handlers.UpdateTodo)
    r.DELETE("/todos/:id", handlers.DeleteTodo)
    r.Run(":4500")
}

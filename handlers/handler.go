
package handlers

import (
	"net/http"
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)


func GetTodos(c *gin.Context) {
    var todos []models.Todo
    database.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, todo)
}


func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if todo.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }
    database.DB.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if todo.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }
    database.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}


func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.Status(http.StatusNoContent)
}
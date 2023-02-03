
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

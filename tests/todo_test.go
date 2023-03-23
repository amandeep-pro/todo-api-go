package tests

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "todo-api/handlers" 
    "todo-api/database"
)

func TestCreateTodo(t *testing.T) {
    // Initialize the database
    database.InitDB()
    router := gin.Default()
    router.POST("/todos", handlers.CreateTodo)

    req, _ := http.NewRequest("POST", "/todos", strings.NewReader(`{"title":"Test Todo","status":"pending"}`))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTodos(t *testing.T) {
    // Initialize the database
    database.InitDB()
    router := gin.Default()
    router.GET("/todos", handlers.GetTodos)

    req, _ := http.NewRequest("GET", "/todos", nil)
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
   
}

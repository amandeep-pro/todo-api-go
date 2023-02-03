

package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "todo-api/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    DB.AutoMigrate(&models.Todo{})
}

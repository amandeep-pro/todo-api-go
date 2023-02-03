
package migrations

import (
    "todo-api/database"
    "todo-api/models"
)

func Migrate() {
    database.DB.AutoMigrate(&models.Todo{})
}


func Seed() {
    database.DB.Create(&models.Todo{Title: "Sample Todo", Status: "pending"})
}


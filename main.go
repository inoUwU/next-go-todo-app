package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB

// DBを初期化する
func initDB() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// テーブルが無ければ自動で生成される
	db.AutoMigrate(&Todo{})
}

// TODO 作成
func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.Copy().ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// Todo一覧を取得
func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()

	initDB()

	// 新規追加
	r.POST("/todos", createTodo)

	// 一覧取得
	r.GET("/todos", getTodos)

	// サーバー起動
	r.Run(":8080")
}

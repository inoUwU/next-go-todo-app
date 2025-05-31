package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
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

var err error

// DBを初期化する
func initDB() {
	db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
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

// TODO削除
func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func updateTodo(c *gin.Context) {
	var req Todo
	if err := c.Copy().ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var todo Todo
	if err := db.Where(&Todo{ID: req.ID}).Find(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todo.Completed = req.Completed
	todo.Title = req.Title
	db.Save(&todo)

	// 更新後の値を返す
	c.JSON(http.StatusOK, todo)
}

func main() {
	r := gin.Default()

	initDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
			"PUT",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	// 新規追加
	r.POST("/todos", createTodo)

	// 一覧取得
	r.GET("/todos", getTodos)

	// 削除
	r.DELETE("/todos/:id", deleteTodo)

	// 更新
	r.PUT("/todos/:id", updateTodo)

	// サーバー起動
	r.Run(":8080")
}

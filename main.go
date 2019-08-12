package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go!!!"
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}

type Todo struct {
	gorm.Model
	Text string
	Status string
}

// 初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースをひらけなかった！")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// 追加
func dbInsert(text string, status string){
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// 更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// 全取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}
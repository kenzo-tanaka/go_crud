package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title       string
	Description string
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open db! (init)")
	}

	defer db.Close()
	db.AutoMigrate(&Book{})
}

func dbInsert(title string, description string) {
    db, err := gorm.Open("sqlite3", "book.sqlite3")
    if err != nil {
        panic("You can't open db! (insert)")
    }

    defer db.Close()
    db.Create(&Book{Title: title, Description: description})
}

func dbGetAll() []Book {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	defer db.Close()
	var books []Book
	db.Order("created_at desc").Find(&books)
	return books
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		books := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			// "data": data,
			// "book": []Book{{"Sample book", "Lorem ipsum dolor..."}, {"Sample book2", "Lorem ipsum ..."}},
			"books": books,
		})
    })
    
    // 追加
    router.POST("/new", func(c *gin.Context) {
        title := c.PostForm("title")
        description := c.PostForm("description")
        dbInsert(title,description)
        c.Redirect(302, "/")
    })

	router.Run()
}

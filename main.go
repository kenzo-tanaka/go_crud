package main

import (
    "github.com/gin-gonic/gin"
)

type Book struct {
    Title string
    Description string
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    data := "Go/Gin!!"
    
    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{
            "data": data,
            "book": []Book{{"Sample book", "Lorem ipsum dolor..."}, {"Sample book2", "Lorem ipsum ..."}},
        })
    })

    router.Run()
}
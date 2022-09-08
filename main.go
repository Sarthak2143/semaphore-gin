package main

import (
    //"net/http"
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    // creating a router with default config
    router = gin.Default()
    // loading our html templates
    router.LoadHTMLGlob("templates/*")
    // adding routes:
    router.GET("/", showIndexPage)
    router.GET("/article/view/:article_id", getArticle)
    // running our web server
    router.Run() // running on localhost:8080
}

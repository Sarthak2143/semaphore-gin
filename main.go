package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // creating a router with default config
    router := gin.Default()
    // loading our html templates
    router.LoadHTMLGlob("templates/*")
    // adding routes:
    router.GET("/", home)
    // running our web server
    router.Run() // running on localhost:8080
}

func home(c *gin.Context) {
    // rendering HTML template
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Home Page",
    })
}

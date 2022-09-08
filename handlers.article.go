package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
    articles := getAllArticles()
    render(c, gin.H{
        "title": "Home Page",
        "payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
    // check if the article id is valid
    if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
        // check if the id exist
        if article, err := getArticleByID(articleID); err == nil {
            // render html
            c.HTML(http.StatusOK, "article.html", gin.H{
                "title": article.Title,
                "payload": article,
            })
        } else {
            c.AbortWithError(http.StatusNotFound, err)
        }
    } else {
        c.AbortWithStatus(http.StatusNotFound)
    }
}

func render(c *gin.Context, data gin.H, tmp string) {
    switch c.Request.Header.Get("Accept") {
    case "application/json":
        // respond with json
        c.JSON(http.StatusOK, data["payload"])
    case "application/xml":
        // respond with xml
        c.XML(http.StatusOK, data["payload"])
    default:
        // respond with html
        c.HTML(http.StatusOK, tmp, data)
    }
}

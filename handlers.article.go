package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
    articles := getAllArticles()
    // rendering template
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Home Page",
        "payload": articles,
    })
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

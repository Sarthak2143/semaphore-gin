package main

import (
    "encoding/json"
)

type article struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var articleList = []article{
    // dummy data
    article{1, "Article I", "Artcle I Content"},
    article{2, "Article II", "Artcle II Content"},
}

func getAllArticles() []article {
    // returning our dummy data
    return articleList
}

package main

type article struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var articleList = []article{
    // dummy data
    article{ID: 1, Title: "Article I", Content: "Artcle I Content"},
    article{ID: 2, Title: "Article II", Content: "Artcle II Content"},
}

func getAllArticles() []article {
    // returning our dummy data
    return articleList
}

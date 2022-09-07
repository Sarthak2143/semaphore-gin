package main

import (
    "testing"
)

func TestGetAllArticles(t *testing.T) {
    list := getAllArticles()
    // checking length of returned and global list
    if len(alist) != len(articleList) {
        t.Fail()
    }
    // checking for each member is identical
    for i, v := range alist {
        if v.ID != articleList[i].ID ||
        v.Title != articleList[i].Title ||
        v.Content != articleList[i].Content {
            t.Fail()
            break
        }
    }
}

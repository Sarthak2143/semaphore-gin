package main

import (
    "strings"
    "testing"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
)

func TestShowPageUnauth(t *testing.T) {
    // getting router
    r := getRouter(true)
    // adding route for home page
    r.GET("/", showIndexPage)
    // create a req to send to above route
    req, _ := http.NewRequest("GET", "/", nil)
    // test http response
    testHTTPResponse(t, r, req, responseRec)
}

func responseRec(w *httptest.ResponseRecorder) bool {
    // test that the http status code is 200 OK
    statusOK := w.Code == http.StatusOK
    // test that page title is "Home Page"
    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && 
    strings.Index(string(p),
        "<title>Home Page</title>") >0
    return statusOK && pageOK
}

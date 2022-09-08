package main

import (
    "os"
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/gin-gonic/gin"
)

var tmpArticleList []article

// setup before testing
func TestMain(m *testing.M) {
    // set gin to test mode
    gin.SetMode(gin.TestMode)
    // run the other tests
    os.Exit(m.Run())
}

// helper func to create router during testing
func getRouter(withTemplate bool) *gin.Engine {
    r := gin.Default()
    if withTemplate {
        r.LoadHTMLGlob("templates/*")
    }
    return r
}

// helper func to process a req and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
    // create a response recorder
    w := httptest.NewRecorder()
    // create a service and process the above req.
    r.ServeHTTP(w, req)
    
    if !f(w) {
        t.Fail()
    }
}

func saveLists() {
    tmpArticleList = articleList
}

func restoreLists() {
    articleList = tmpArticleList
}

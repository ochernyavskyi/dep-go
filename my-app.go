package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hellow world"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")

	router.GET("/", HelloWorld)
	err := router.Run(":80")
	if err != nil {
		fmt.Print("fk u linter hehe")
	}

}


// package main
// import (
//     "net/http"
//     "fmt"
//     )
//
// type msg string
// func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//    fmt.Fprint(resp, m)
// }
// func main() {
//     msgHandler := msg("Hello World 2 !")
//     http.ListenAndServe("0.0.0.0:80", msgHandler)
// }
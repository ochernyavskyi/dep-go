package main
import (
    "net/http"
    "fmt"
    )

type msg string
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
   fmt.Fprint(resp, m)
}
func main() {
    msgHandler := msg("Hello World 2 !")
    http.ListenAndServe("0.0.0.0:80", msgHandler)
}
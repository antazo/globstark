package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./web")))
    fmt.Println("Server started at http://localhost:7979")
    http.ListenAndServe(":7979", nil)
}
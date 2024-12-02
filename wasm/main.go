package main

import (
    "syscall/js"
)

func main() {
    document := js.Global().Get("document")
    body := document.Call("createElement", "div")
    body.Set("innerHTML", "Hello, Globstark!")
    document.Get("body").Call("appendChild", body)
    select {}
}
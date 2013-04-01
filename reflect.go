package main

import (
    "fmt"
    "reflect"
)

var urls = map[string]*Func {
    "/" : test,
}

func main() {
    fmt.Printf("Im main%s\n", reflect.ValueOf("test"))
}

func test() {
    fmt.Printf("Im test func\n")
}
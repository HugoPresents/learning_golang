package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
    var f interface{}
    err := json.Unmarshal(b, &f)
    if err != nil {
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("%v\n", f)
    }
}
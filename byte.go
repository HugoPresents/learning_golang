package main

import (
    "fmt"
    "os"
    "encoding/json"
)

func main() {
    settings := loadSettings()
    fmt.Printf("%v\n", settings)
}

func loadSettings() interface{} {
    f, _ := os.Open("./settings.json")
    buf := make([]byte, 1024)
    n, _ := f.Read(buf)
    if n == 0 {
        fmt.Printf("settings file is null")
    }
    var j interface{}
    json.Unmarshal(buf[:n], &j)
    return j
}
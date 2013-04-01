package main

import (
    "fmt"
    "bytes"
)

func main() {
    a := [][]byte{
        
    }
    b := []byte{1, 2, 3, 4, 5}
    c := []byte{1, 2}
    d := bytes.Join(a, b, c)
    fmt.Printf("%v\n", d)
}
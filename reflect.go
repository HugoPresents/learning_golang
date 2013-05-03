package main

import (
	"fmt"
	"reflect"
)

type Index struct {
	Name string
}

func (index *Index) DisplayName() {
	fmt.Printf("%s", index.Name)
}

func DisplayName() {
	fmt.Printf("%s", "my name")
}

func main() {
	url := "Index.DisplayName"
	index := Index{Name: "my name"}
	funcs := map[string]interface{}{
		"Index": index.DisplayName,
	}
	f := reflect.ValueOf(funcs[url])
	in := make([]reflect.Value, 0)
	f.Call(in)
}

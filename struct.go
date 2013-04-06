package main

import "fmt"

type Model struct {
    TableName string
}

type PostModel struct {
    Model
    PostName string
}

func (model *PostModel) getName() string {
    return model.PostName
}

func (model *Model) info() string {
    return model.TableName
}

func main() {
    post_model := newPostModel()
    fmt.Printf("%v\n", post_model)
    // 取值的时候可以直接去隐藏字段名，也可以加上所包含的类型名（但是这两种访问方式的意义不同，你懂得~）
    fmt.Printf("Table name is \"%s\"\n post name is \"%s\"\n", post_model.TableName, post_model.PostName)
    fmt.Printf("Table name is \"%s\"\n post name is \"%s\"\n", post_model.Model.TableName, post_model.PostName)
}

func newPostModel() *PostModel {
    // new的时候隐藏字段要注明所包含的类型名
    return &PostModel{Model{"post"}, "mypost"}
}

package main

import (
    "database/sql"
    "fmt"
    "github.com/astaxie/beedb"
    _ "github.com/ziutek/mymysql/godrv"
)

type Category struct {
    Id int
    ParentId int
    OrderId int
    NavDisplay int
    Name string
    DisplayName string
    Status int
}

func main() {
    db, err := sql.Open("mymysql", "golog/root/rabbit")
    if err != nil {
        panic(err)
    }
    orm := beedb.New(db)
    var cat Category
    err = orm.Where("id=?", 1).Find(&cat)
    if err != nil {
        fmt.Printf("%v", err)
    }
    fmt.Printf("%v", cat)
}
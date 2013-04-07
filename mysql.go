package main

import (
    "fmt"
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql"
)

//var db

func main () {
    //var db_err error
    db, err := sql.Open("mysql", "rabbit:rabbit@tcp(192.168.0.102:3306)/golog?charset=utf8")
    checkErr(err)
    /*
    data := map[string]string {
        "name" : "cat1",
        "display_name" : "分类1",
    }
    */
    stmt, err := db.Prepare("INSERT category SET name=?, display_name=?")
    checkErr(err)
    res, err := stmt.Exec("cat1", "分类1")
    checkErr(err)
    fmt.Printf("%v\n", res)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
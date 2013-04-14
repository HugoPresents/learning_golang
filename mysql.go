package main

import (
    "fmt"
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql"
    "strings"
)

//var db

func main () {
    //var db_err error
    //db, err := sql.Open("mysql", "rabbit:rabbit@tcp(192.168.0.102:3306)/golog?charset=utf8")
    //checkErr(err)
    
    data := map[string]string {
        "name" : "cat2",
        "display_name" : "分类1",
    }
    
    /*
    stmt, err := db.Prepare("INSERT category SET name=?, display_name=?")
    checkErr(err)
    res, err := stmt.Exec("cat1", "分类1")
    checkErr(err)
    fmt.Printf("%v\n", res)
    */
    insert(data)
    /*
    my_s := []int{1, 2, 3}
    test(1, 2, 3)
    test(my_s...)
    */
}

func test(params ...interface{}) {
    fmt.Println(params)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func insert(data map[string]string) {
    db, err := sql.Open("mysql", "rabbit:rabbit@tcp(192.168.0.102:3306)/golog?charset=utf8")
    sql := "INSERT category SET "
    fmt.Printf("%v\n", data)
    if len(data) < 1 {
        fmt.Print("no data!")
        return
    }
    params := make([]interface{}, len(data))
    i := 0
    for column, value := range(data) {
        params[i] = value
        sql += column + "=?, "
        i ++
    }
    fmt.Printf("sql: %s\nparams: %v(%T)\n", sql, params, params)
    sql = strings.Trim(sql, ", ")
    fmt.Printf("sql: %s\nparams: %v(%T)\n", sql, params, params)
    stmt, err := db.Prepare(sql)
    res, err := stmt.Exec(params...)
    checkErr(err)
    fmt.Printf("%v\n", res)
}
package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:Ram_3007@tcp(127.0.0.1:3306)/student")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    insert, err := db.Query("INSERT INTO student_info VALUES('Ram Nivas',39130030,9494921767)")
    if err !=nil {
        fmt.Println(err)
    }
    defer insert.Close()
    fmt.Println("Values added!")
}
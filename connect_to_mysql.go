package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

func main() {
    db, err := sql.Open("mysql", "root:password@/go_test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    rows, err := db.Query("SELECT * FROM tt")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s\n", name)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}

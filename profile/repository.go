package profile

import (
    "database/sql"
    "net/http"
    "fmt"
    "runtime"
)


func qAllProfile(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT user_id,username,email FROM users")
    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    fmt.Println(" qAllProfile goroutine..", runtime.NumGoroutine())
    return rows,err
}
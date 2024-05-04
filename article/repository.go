package article

import (
    "database/sql"
    "net/http"
    "fmt"
    "runtime"
)


func qArt(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT id, title, description, img, owner, completed, created_at, updated_at FROM article WHERE Completed=$1", true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    fmt.Println(" qArt goroutine..", runtime.NumGoroutine())

    return rows,err
}


func qUsArt(w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM article WHERE owner=$1", owner)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }

    return rows,err
}

package schedule


import (
    "database/sql"
    "net/http"
    "fmt"
    "time"
)


func qSchSelect(
    w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM schedule WHERE en_hour >= $1", time.Now().UTC())

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    defer conn.Close()
    return rows,err
}


func qSch(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM schedule ORDER BY id DESC")

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    defer conn.Close()
    return rows,err
}

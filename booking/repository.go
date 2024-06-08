package booking

import (
    "database/sql"
    "fmt"
    "net/http"
    "time"

    // "github.com/lib/pq"
)


func qSearch(
    w http.ResponseWriter, conn *sql.DB, start time.Time,end time.Time) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_d WHERE st_date <= $1 AND en_date >= $2 AND NOT daterange($1, $2, '[]') @> ANY(dates) OR dates IS NULL", start,end)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}


func prvSearcHours(
    w http.ResponseWriter, conn *sql.DB, start time.Time,end time.Time, s time.Time,e time.Time) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_h WHERE st_hour <= $1 AND en_hour >= $2 AND NOT tsrange($3, $4, '[]') @> ANY(hours) OR hours IS NULL", start,end,s,e)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}


/*func qSearch(
    w http.ResponseWriter, conn *sql.DB, start time.Time,end time.Time) (rows *sql.Rows, err error) {

    fmt.Println("start..", start)
    fmt.Println("end..", end)

    rows,err = conn.Query("SELECT * FROM provision_d WHERE st_date <= $1 AND en_date >= $2 AND NOT daterange($1, $2, '[]') @> ANY(s_dates) AND NOT daterange($1, $2, '[]') @> ANY(e_dates)", start,end)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}*/


func qBkg(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM booking")

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}
func qBkgBool(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM booking WHERE Completed=$1", true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}


func qUsBkg(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM booking WHERE owner=$1", owner)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}

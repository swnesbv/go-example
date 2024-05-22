package slider

import (
    "database/sql"
    "net/http"
    "fmt"
)

func qCollection(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM collection WHERE owner=$1 ORDER BY id", owner)

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


func qArt(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT id FROM article WHERE owner=$1 ORDER BY id", owner)

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
func qSch(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT id FROM schedule WHERE owner=$1 ORDER BY id", owner)

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
func qPrvD(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT id FROM provision_d WHERE owner=$1 ORDER BY id", owner)

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
func qPrvH(
    w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT id FROM provision_h WHERE owner=$1 ORDER BY id", owner)

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
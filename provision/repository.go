package provision


import (
    "database/sql"
    "net/http"
    "fmt"
)


func qPrvD(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_d")

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
func qPrvH(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_h")

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


func qPrvDBool(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_d WHERE Completed=$1", true)

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


func qUsPrvD(w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM provision_d WHERE owner=$1", owner)

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

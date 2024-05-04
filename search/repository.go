package search

import (
    "database/sql"
    "net/http"
    "fmt"
)


func qSearchArt(w http.ResponseWriter, r *http.Request, conn *sql.DB) (rows *sql.Rows, err error) {

    q := r.URL.Query().Get("q")

    rows,err = conn.Query("SELECT id,title,completed FROM article WHERE completed=$1 AND LOWER(title) LIKE '%' || $2 || '%'", true, q)

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

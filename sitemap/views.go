package sitemap

import (
    "database/sql"
    "fmt"
    "net/http"
)


func allSpArt(w http.ResponseWriter, rows *sql.Rows) (list []*Article, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Article)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Img,
            &i.Owner,
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list, err
}

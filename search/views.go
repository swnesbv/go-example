package search

import (
    "database/sql"
    "fmt"
    "net/http"

    "go_authentication/article"
)


func searcArt(w http.ResponseWriter, rows *sql.Rows) (list []*article.Article, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(article.Article)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Completed,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}

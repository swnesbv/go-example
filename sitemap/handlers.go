package sitemap

import (
    "time"
    "fmt"
    "net/http"
    "text/template"

    "go_authentication/connect"
)


 func SitemapHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        start := time.Now()

        conn := connect.ConnSql()
        rows,err := qSpArt(w, conn)
        if err != nil {
            return
        }
        list,err := allSpArt(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        elapsed := time.Since(start)
        fmt.Printf(" sitemap art time.. :  %s \n", elapsed)
        
        tpl := template.Must(template.ParseFiles("./sitemap.xml"))
        tpl.Execute(w, list)
    }
 }

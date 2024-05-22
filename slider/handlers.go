package slider

import (
    //"fmt"
    //"time"
    "net/http"
    "html/template"

    //"github.com/lib/pq"

    // "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func CollectionAll(w http.ResponseWriter, r *http.Request) {

    cls, err := authtoken.OnToken(w, r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        conn := connect.ConnSql()
        owner := cls.User_id
        rows,err := qCollection(w, conn,owner)
        if err != nil {
            return
        }
        list,err := allCollection(w, rows)
        if err != nil {
            return
        }

        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/all_collection.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}
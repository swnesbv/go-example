package owner_ssc

import (
    // "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func OwrAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        conn := connect.ConnSql()
        rows,err := qsOwSsc(w, conn,cls.User_id)
        if err != nil {
            return
        }
        list,err := owSsc(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}


func DetOwrSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }

        cls,err := authtoken.ListToken(w,r)
        if err != nil {
            return
        }
        
        owner := cls.User_id
        conn := connect.ConnSql()
        i,err := ownerIdSsc(w, conn,id,owner)
        if err != nil {
            return
        }
        defer conn.Close()

        type ListData struct {
            Auth string
            I Subscription
        }
        data := ListData {
            Auth: cls.Email,
            I: i,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/detail.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }
}
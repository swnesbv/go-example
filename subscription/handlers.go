package subscription

import (
    // "fmt"
    "net/http"
    "html/template"

    // "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func AllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := qsAllSsc(w, conn)
        if err != nil {
            return
        }
        list,err := allSsc(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/all.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)
    }
}


func ToUsAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        to_user := cls.User_id

        conn := connect.ConnSql()
        rows,err := qsUserAllSsc(w, conn,to_user)
        if err != nil {
            return
        }
        list,err := userSsc(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/user.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)
    }
}

func ToGroupAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id

        conn := connect.ConnSql()
        rows,err := qsGroupAllSsc(w, conn,owner)
        if err != nil {
            return
        }
        list,err := roomSsc(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)
    }
}
package booking

import (
    // "fmt"
    // "time"
    "net/http"
    "html/template"

    // "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func SearchDays(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        i,err := authtoken.TokenSearch(w,r)
        if i == nil {
            return
        }
        if err != nil {
            return
        }

        conn := connect.ConnSql()
        rows,err := qSearch(w, conn,i.D_start,i.D_end)
        if err != nil {
            return
        }
        list,err := allSearchD(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/all_days.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}


func SearcHours(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        i,err := authtoken.TokenSearch(w,r)
        if i == nil {
            return
        }
        if err != nil {
            return
        }

        conn := connect.ConnSql()
        rows,err := prvSearcHours(w,conn, i.D_start,i.D_end,i.Start,i.End)
        if err != nil {
            return
        }
        list,err := allSearcH(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/all_hours.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}


func BkgAll(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := qBkg(w, conn)
        if err != nil {
            return
        }
        list,err := allBkg(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/booking/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}

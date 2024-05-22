package chat

import (
    "time"
    "fmt"
    //"os"
    //"io"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func Creat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/creativity.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        owner := cls.User_id
        title := r.FormValue("title")
        description := r.FormValue("description")

        sqlstr := "INSERT INTO groups (title, description, owner, created_at) VALUES ($1,$2,$3,$4)"
        _, err := db.Exec(sqlstr, title,description,owner,time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w,r, "/all-group", http.StatusFound)
    }
}


func UpGr(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.SqlToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    art,err := idGroup(w, id)
    if err != nil {
        return
    }

    flag,err := options.ParseBool(r.FormValue("completed"))
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, " Error: ParseBool()..  : %+v\n", err)
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/update.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", art)
    }


    if r.Method == "POST" {

        owner := cls.User_id
        title := r.FormValue("title")
        description := r.FormValue("description")

        sqlstr := "UPDATE groups SET title=$3, description=$4, completed=$5, updated_at=$6 WHERE id=$1 AND owner=$2"
        
        _, err := db.Exec(sqlstr, id,owner,title,description,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, " Error: db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w,r, "/all-or-owner-group", http.StatusFound)
    }
}
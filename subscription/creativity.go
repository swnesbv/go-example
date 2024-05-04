package subscription

import (
    "time"
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/connect"
    "go_authentication/options"
    "go_authentication/authtoken"
)


func ToUpUsSsc(w http.ResponseWriter, r *http.Request) {

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

    flag,err := options.ParseBool(r.FormValue("completed"))
    if err != nil {
        fmt.Fprintf(w, "err ParseBool()..  : %+v\n", err)
        return
    }

    conn := connect.ConnSql()
    i,err := userIdSsc(w, conn,id)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/update_ssc.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        conn := connect.ConnSql()
        sqlstr := "UPDATE subscription SET completed=$2, updated_at=$3 WHERE id=$1"
        
        _, err := conn.Exec(sqlstr, id,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }
        defer conn.Close()
        http.Redirect(w,r, "/all-touser-ssc", http.StatusFound)
    }
}


func ToUpGroupSsc(w http.ResponseWriter, r *http.Request) {

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

    flag,err := options.ParseBool(r.FormValue("completed"))

    if err != nil {
        fmt.Fprintf(w, "err ParseBool()..  : %+v\n", err)
        return
    }

    owner := cls.User_id
    conn := connect.ConnSql()
    i,err := roomIdSsc(w, conn,id,owner)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/update_ssc.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        conn := connect.ConnSql()
        sqlstr := "UPDATE subscription SET completed=$2, updated_at=$3 WHERE id=$1"
        
        _, err := conn.Exec(sqlstr, id,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }
        defer conn.Close()
        http.Redirect(w,r, "/all-toroom-ssc", http.StatusFound)
    }
}
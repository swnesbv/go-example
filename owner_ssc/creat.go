package owner_ssc

import (
    "time"
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/connect"
    "go_authentication/options"
    "go_authentication/authtoken"
)


func AddSscUs(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/creativity.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        conn := connect.ConnSql()
        sqlstr := "INSERT INTO subscription (title, description, owner, to_user, created_at) VALUES ($1,$2,$3,$4,$5)"

        _, err := conn.Exec(sqlstr, title,description,owner,id,time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-ssc", http.StatusFound)
    }
}

func AddSscGr(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/creativity.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        conn := connect.ConnSql()
        sqlstr := "INSERT INTO subscription (title, description, owner, to_group, created_at) VALUES ($1,$2,$3,$4,$5)"

        _, err := conn.Exec(sqlstr, title,description,owner,id,time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-ssc", http.StatusFound)
    }
}


func OwrUpSsc(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.SqlToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    owner := cls.User_id
    conn := connect.ConnSql()
    i,err := ownerIdSsc(w, conn,id,owner)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/update.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        sqlstr := "UPDATE subscription SET title=$3,description=$4,updated_at=$5 WHERE id=$1 AND owner=$2;"

        _, err := conn.Exec(sqlstr, id,owner,title,description,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-ssc", http.StatusFound)
    }
}

func OwrDelSsc(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        data := struct {
            Items string
        }{
            Items: cls.Email,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/delete.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }


    if r.Method == "POST" {

        owner := cls.User_id
        conn := connect.ConnSql()
        sqlstr := "DELETE FROM subscription WHERE id=$1 AND owner=$2"
        
        _, err := conn.Exec(sqlstr, id,owner)
        
        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }
        
        defer conn.Close()
        http.Redirect(w,r, "/all-ssc", http.StatusFound)
    }
}
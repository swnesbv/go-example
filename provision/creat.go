package provision


import (
    "time"
    "fmt"
    "net/http"
    "html/template"
    // "github.com/lib/pq"

    "go_authentication/connect"
    // "go_authentication/options"
    "go_authentication/authtoken"
)


func CreatDays(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/creat.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        title       := r.FormValue("title")
        description := r.FormValue("description")

        start,terr := time.Parse(time.DateOnly, r.FormValue("start"))
        end,terr := time.Parse(time.DateOnly, r.FormValue("end"))
        if terr != nil {
            fmt.Fprintf(w, "err time Parse..! : %+v\n", terr)
            return
        }

        conn := connect.ConnSql()
        sqlstr := "INSERT INTO provision_d (title,description,owner,st_date,en_date,created_at) VALUES ($1,$2,$3,$4,$5,$6)"

        _, err := conn.Exec(sqlstr, title,description,cls.User_id,start,end,time.Now())

        if err != nil {
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-prv-days", http.StatusFound)
    }
}

func CreatHours(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/creat_hours.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {
        
        title       := r.FormValue("title")
        description := r.FormValue("description")

        loc, _ := time.LoadLocation("UTC")
        start,terr := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("start") + ":00", loc)
        end,terr := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("end") + ":00", loc)

        if terr != nil {
            fmt.Fprintf(w, "err time Parse..! : %+v\n", terr)
            return
        }

        conn := connect.ConnSql()
        sqlstr := "INSERT INTO provision_h (title,description,owner,st_hour,en_hour,created_at) VALUES ($1,$2,$3,$4,$5,$6)"

        _, err := conn.Exec(sqlstr, title,description,cls.User_id,start,end,time.Now())

        if err != nil {
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-prv-hours", http.StatusFound)
    }
}

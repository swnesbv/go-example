package schedule

import (
    "time"
    "fmt"
    "net/http"
    "html/template"
    "github.com/lib/pq"

    "go_authentication/connect"
    // "go_authentication/options"
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

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/schedule/creat.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {
        
        title       := r.FormValue("title")
        description := r.FormValue("description")

        loc, _ := time.LoadLocation("UTC")
        start,derr := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("start") + ":00", loc)
        end,derr := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("end") + ":00", loc)
        if derr != nil {
            fmt.Fprintf(w, "err ParseInLocation..! : %+v\n", derr)
            return
        }

        list,pserr := psForm(w,r)
        if pserr != nil {
            return
        }

        conn := connect.ConnSql()
        owner := cls.User_id
        str := "INSERT INTO schedule (title,description,owner,st_hour,en_hour,hours,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"

        _, err := conn.Exec(
            str, title,description,owner,start,end,pq.Array(list),time.Now())

        if err != nil {
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-schedule", http.StatusFound)
    }
}

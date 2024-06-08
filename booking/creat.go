package booking


import (
    "time"
    "fmt"
    "net/http"
    "html/template"
    "encoding/json"
    "encoding/base64"

    "go_authentication/connect"
    "go_authentication/options"
    "go_authentication/authtoken"
)


func Creat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil { return }
    if err != nil { return }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/booking/creat.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {

        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        conn := connect.ConnSql()
        sqlstr := "INSERT INTO booking (title, description, owner, created_at) VALUES ($1,$2,$3,$4)"

        _, err := conn.Exec(sqlstr, title,description,owner,time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-bkg", http.StatusFound)
    }
}


func Period(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/booking/period.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {

        start,err := time.Parse(time.DateOnly, r.FormValue("start"))
        end,err   := time.Parse(time.DateOnly, r.FormValue("end"))
        if err != nil {
            fmt.Fprintf(w, " Error: time Parse..! : %+v\n", err)
            return
        }

        // if start.After(end) || start.Before(time.Now())  {
        //     fmt.Fprintf(w, " incorrect time interval points..!")
        //     return 
        // }

        fmt.Println( "start", start)
        fmt.Println( "end", end)

        data := authtoken.ListData {
            D_start: start,
            D_end: end,
        }
        list,err := json.Marshal(data)
        if err != nil {
            fmt.Fprintf(w, " Error: json Marshal..! : %+v\n", err)
            return
        }

        tokenstr := base64.StdEncoding.EncodeToString([]byte(list))

        cookie := http.Cookie{
            Name:     "Period",
            Value:    tokenstr,
            Path:     "/search-period-days",
            MaxAge:   3600,
            HttpOnly: true,
            Secure:   false,
            SameSite: http.SameSiteLaxMode,
        }
        http.SetCookie(w, &cookie)

        http.Redirect(w,r, "/search-period-days", http.StatusFound)
        return
    }
}


func PeriodHours(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/booking/period_hours.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {

        // if start.After(end) || start.Before(time.Now())  {
        //     fmt.Fprintf(w, " incorrect time interval points..!")
        //     return 
        // }

        loc, _ := time.LoadLocation("UTC")
        start,err := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("start") + ":00", loc)
        end,err := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("end") + ":00", loc)

        // layout := "2006-01-02T15:04:05.000Z"
        // start,err := time.Parse(
        //     layout, r.FormValue("start") + ":00.000Z")
        // end,err := time.Parse(
        //     layout, r.FormValue("end") + ":00.000Z")

        if err != nil {
            fmt.Fprintf(w, " Error: time Parse..! : %+v\n", err)
            return
        }

        s := start.Format(time.TimeOnly)
        e := end.Format(time.TimeOnly)
        ss,err := time.Parse(time.TimeOnly, s)
        ee,err := time.Parse(time.TimeOnly, e)

        data := authtoken.ListData {
            D_start: start,
            D_end: end,
            Start: ss,
            End: ee,
        }
        list,err := json.Marshal(data)
        if err != nil {
            fmt.Fprintf(w, " Error: json Marshal..! : %+v\n", err)
            return
        }

        tokenstr := base64.StdEncoding.EncodeToString([]byte(list))

        cookie := http.Cookie{
            Name:     "Period",
            Value:    tokenstr,
            Path:     "/search-period-hours",
            MaxAge:   3600,
            HttpOnly: true,
            Secure:   false,
            SameSite: http.SameSiteLaxMode,
        }
        http.SetCookie(w, &cookie)

        http.Redirect(w,r, "/search-period-hours", http.StatusFound)
        return
    }
}


func UpBkg(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil { return }
    cls,err := authtoken.SqlToken(w,r)
    if cls == nil { return }
    if err != nil { return }

    conn := connect.ConnSql()

    i,err := idBkg(w, conn,id,cls)
    if err != nil { return }

    flag,err := options.ParseBool(r.FormValue("completed"))
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, " Error: ParseBool()..  : %+v\n", err)
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles( "./tpl/navbar.html", "./tpl/booking/update.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        str := "UPDATE booking SET title=$3, description=$4, completed=$5, updated_at=$6 WHERE id=$1 AND owner=$2"
        
        _, err := conn.Exec(str, id,owner,title,description,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w, r, "/all-bkg", http.StatusFound)
    }
}

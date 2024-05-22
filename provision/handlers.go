package provision


import (
    "fmt"
    "time"
    "net/http"
    "html/template"

    "github.com/lib/pq"

    "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func PrvAllD(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := qPrvD(w, conn)
        if err != nil {
            return
        }
        list,err := allPrvD(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/all_days.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}
func PrvAllH(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := qPrvH(w, conn)
        if err != nil {
            return
        }
        list,err := allPrvH(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/all_hours.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}


func IdPrvDays(w http.ResponseWriter, r *http.Request) {

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
    
    conn := connect.ConnSql()

    if r.Method == "GET" {

        i,err := idPrvD(w, conn,id)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/detail_days.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }

    if r.Method == "POST" {

        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        start,err := time.Parse(time.DateOnly, r.FormValue("start"))
        end,err := time.Parse(time.DateOnly, r.FormValue("end"))
        if err != nil {
            fmt.Fprintf(w, " Error: time Parse..! : %+v\n", err)
            return
        }

        var s_list []time.Time
        var e_list []time.Time

        var list []time.Time

        list = append(list, start,end)
        fmt.Println(" list..", list)

        s_list = append(s_list, start)
        fmt.Println(" start list..", s_list)
        e_list = append(e_list, end)
        fmt.Println(" end list..", e_list)

        bkg := "INSERT INTO booking (title,description,owner,to_prv_d,st_date,en_date,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"
        _, berr := conn.Exec(bkg, title,description,owner,id,start,end,time.Now())
        if berr != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", berr)
            return
        }

        prv := "UPDATE provision_d SET s_dates=array_cat(s_dates, $2),e_dates=array_cat(e_dates, $3),dates=array_cat(dates, $4),updated_at=$5 WHERE id=$1"
        _, perr := conn.Exec(prv, id,pq.Array(s_list),pq.Array(e_list),pq.Array(list),time.Now())
        if perr != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", perr)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-bkg", http.StatusFound)
    }
}


func IdPrvHours(w http.ResponseWriter, r *http.Request) {

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
    
    conn := connect.ConnSql()

    if r.Method == "GET" {

        i,err := idPrvH(w, conn,id)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/provision/detail_hours.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }

    if r.Method == "POST" {

        owner := cls.User_id
        title       := r.FormValue("title")
        description := r.FormValue("description")

        loc, _ := time.LoadLocation("UTC")
        start,err := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("start") + ":00", loc)
        end,err := time.ParseInLocation(
            "2006-01-02T15:04:05", r.FormValue("end") + ":00", loc)
            
        if err != nil {
            fmt.Fprintf(w, " Error: time Parse..! : %+v\n", err)
            return
        }

        var s_list []time.Time
        var e_list []time.Time
        var list []time.Time

        s := start.Format(time.TimeOnly)
        e := end.Format(time.TimeOnly)
        ss,err := time.Parse(time.TimeOnly, s)
        ee,err := time.Parse(time.TimeOnly, e)

        list   = append(list, ss,ee)
        s_list = append(s_list, start)
        e_list = append(e_list, end)

        bkg := "INSERT INTO booking (title,description,owner,to_prv_h,st_hour,en_hour,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"
        _, berr := conn.Exec(bkg, title,description,owner,id,start,end,time.Now())
        
        if berr != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", berr)
            return
        }

        prv := "UPDATE provision_h SET s_hours=array_cat(s_hours, $2),e_hours=array_cat(e_hours, $3),hours=array_cat(hours, $4),updated_at=$5 WHERE id=$1"
        _, perr := conn.Exec(prv, id,pq.Array(s_list),pq.Array(e_list),pq.Array(list),time.Now())

        if perr != nil {
            fmt.Fprintf(w, " Error: Exec..! : %+v\n", perr)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/all-bkg", http.StatusFound)
    }
}


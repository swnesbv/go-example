package profile

import (
    "fmt"
    "net/http"
    "html/template"
    "runtime"

    "go_authentication/connect"
    "go_authentication/authtoken"
)


func Home(w http.ResponseWriter, r *http.Request) {

    tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/index.html", "./tpl/base.html" ))
    tpl.ExecuteTemplate(w, "base", nil)
}


func Alluser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/all.html", "./tpl/base.html" ))

        conn := connect.ConnSql()
        rows,err := qAllProfile(w, conn)

        if err != nil {
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            return
        }
        defer rows.Close()

        cls := authtoken.WhoisWho(w,r)

        type ListData struct {
            Auth string
            I []*UserList
        }

        if cls != nil {
            
        var list []*UserList
        for rows.Next() {
            i := new(UserList)

            err := rows.Scan(&i.User_id,&i.Username,&i.Email)
            if err != nil {
                fmt.Fprintf(w, "Error: Scan..! : %+v\n", err)
                return
            }
            list = append(list, i)
        }

        data := ListData {
            Auth: cls.Email,
            I: list,
        }
        tpl.ExecuteTemplate(w, "base", data)

        } else {

        var list []*UserList
        for rows.Next() {
            i := new(UserList)
            err := rows.Scan(&i.User_id,&i.Username,&i.Email)

            if err != nil {
                fmt.Fprintf(w, "Error: Scan..! : %+v\n", err)
                return
            }
            list = append(list, i)
        }
        data := ListData {
            I: list,
        }
        tpl.ExecuteTemplate(w, "base", data)
        }

        if err = rows.Close(); err != nil {
            fmt.Fprintf(w, "Error: sql..! : %+v\n", err)
        }
        defer conn.Close()
        fmt.Println(" All user goroutine..", runtime.NumGoroutine())
    }
}


/*func Alluser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := conn.Query("SELECT username,email FROM users")

        if err != nil {
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            return
        }
        defer rows.Close()

        var list []*UserList
        
        for rows.Next() {
            data := new(UserList)
            err := rows.Scan(&data.Username, &data.Email)

            if err != nil {
                fmt.Fprintf(w, "Error: Scan..! : %+v\n", err)
                return
            }
            list = append(list, data)
        }

        if err = rows.Close(); err != nil {
            fmt.Fprintf(w, "Error: sql..! : %+v\n", err)
        }

        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
}*/
package sqlcsv

import (
    "os"
    "fmt"
    "time"
    "strings"
    "net/http"
    "encoding/csv"
    "html/template"

    "github.com/lib/pq"

    "go_authentication/connect"
    "go_authentication/authtoken"
)


func ImpSch(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        
        rows,err := conn.Query("SELECT * FROM schedule WHERE owner=$1", owner)

        if err != nil {
            switch {
                case true:
                fmt.Fprintf(w, " Error: rows..! : %+v\n", err)
                break
            }
            return
        }
        defer rows.Close()

        mkdirerr := os.MkdirAll("./sfl/static/csv/" + cls.Email, 0750)
        if mkdirerr != nil {
            fmt.Println(" Error os.MkdirAll..", mkdirerr)
        }
        file,err := os.Create("./sfl/static/csv/" + cls.Email + "/schedule.csv")
        if err != nil {
            fmt.Println(" Error os.Create..", err)
        }
        defer file.Close()

        wr := csv.NewWriter(file)
        defer wr.Flush()

        for rows.Next() {
            i := new(Schedule)
            err := rows.Scan(&i.Id,&i.Title,&i.Description,&i.Owner,&i.St_hour,&i.En_hour,pq.Array(&i.Hours),pq.Array(&i.Occupied),&i.Completed,&i.Created_at,&i.Updated_at)
            if err != nil {
                fmt.Println(" Error Scan row..", err)
                return
            }
            list := []*Schedule{}
            list = append(list, i)

            var row string
            for _, val := range list {
                row += fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", val.Id,val.Title,val.Description,val.Owner,val.St_hour,val.En_hour,val.Hours,val.Occupied,val.Completed,val.Created_at,val.Updated_at)
            }
            file.WriteString(fmt.Sprintf("%s\n", row))
        }

        fmt.Println(" CSV successfully..!")

        defer conn.Close()

        http.Redirect(w,r, "/static/csv/" + cls.Email + "/schedule.csv", http.StatusFound)

    }
}


func ExpSch(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/csv/export.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", nil)

    }
    if r.Method == "POST" {

        conn := connect.ConnSql()
        owner := cls.User_id

        file, _, err := r.FormFile("file")
        if err != nil{
            fmt.Println(" err data retrieving..", err)
            return
        }

        reader := csv.NewReader(file)
        rows,err := reader.ReadAll()
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, " Error ReadAll..! : %+v\n", err)
            return
        }

        for _, row := range rows {

            t  := row[1]
            d  := row[2]
            sh := row[4]
            eh := row[5]
            h  := row[6]
            o  := row[7]

            var list []time.Time
            th := strings.Fields(h)
            for x := range th {
                tht,_ := time.Parse(time.TimeOnly, th[x])
                list = append(list, tht)
            }
            
            var list2 []time.Time
            to := strings.Fields(o)
            for x := range to {
                tot,_ := time.Parse(time.TimeOnly, to[x])
                list2 = append(list2, tot)
            }

            sqlst := "INSERT INTO schedule (title,description,owner,st_hour,en_hour,hours,occupied,completed,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

            _, err := conn.Exec(sqlst, t,d,owner,sh,eh,pq.Array(list),pq.Array(list2),false,time.Now(),nil)

            if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, " Error Exec..! : %+v\n", err)
                return
            }
        }

        defer conn.Close()

        http.Redirect(w,r, "/all-schedule", http.StatusFound)
    }
}

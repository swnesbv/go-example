package schedule

import (
    "database/sql"
    "fmt"
    "time"
    "net/http"
    
    "github.com/lib/pq"

    // "go_authentication/authtoken"
)


func Converter (w http.ResponseWriter, v string) time.Time {

    loc, _ := time.LoadLocation("UTC")
    to, _ := time.ParseInLocation("2006-01-02T15:04:05",v + ":00", loc)

    s := to.Format(time.TimeOnly)
    ss, _ := time.Parse(time.TimeOnly, s)
    return ss
}
func psForm (w http.ResponseWriter, r *http.Request) (list []time.Time, err error) {

    pserr := r.ParseForm()
    if pserr != nil {
        fmt.Fprintf(w, "Error ParseForm..! : %+v\n", pserr)
        return
    }

    ps := r.Form["list"]

    var v string
    var ss time.Time
    for _, v = range ps {
        ss = Converter(w,v)
        list = append(list, ss)
    }
    return list,err
}


func Select(hours,occupied []string) (slt []string) {
    m := make(map[string]bool)

    for _, i := range occupied {
        m[i] = true
    }

    for _, i := range hours {
        if _, ok := m[i]; !ok {
            slt = append(slt, i)
        }
    }
    fmt.Println(" slt..", slt)
    return slt
}
func allSelect(w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Schedule)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_hour,
            &i.En_hour,
            pq.Array(&i.Hours),
            pq.Array(&i.Occupied),
            &i.Completed,
            &i.Created_at, 
            &i.Updated_at, 
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        free := Select(i.Hours,i.Occupied)
        i.Hours = free
        list = append(list, i)
    }
    return list,err
}


func allSch(w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Schedule)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_hour,
            &i.En_hour,
            pq.Array(&i.Hours),
            pq.Array(&i.Occupied),
            &i.Completed,
            &i.Created_at, 
            &i.Updated_at, 
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
        // fmt.Printf("Array..! :  %T", pq.Array(&i.Hours))
    }
    return list,err
}

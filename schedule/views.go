package schedule

import (
    "database/sql"
    "fmt"
    "net/http"
    
    "github.com/lib/pq"

    // "go_authentication/authtoken"
)


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


package provision

import (
    "database/sql"
    "fmt"
    "net/http"

    "github.com/lib/pq"

    // "go_authentication/authtoken"
)


func allPrvD(w http.ResponseWriter, rows *sql.Rows) (list []*ProvisionD, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(ProvisionD)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_date,
            &i.En_date,
            pq.Array(&i.S_dates),
            pq.Array(&i.E_dates),
            pq.Array(&i.Dates),
            &i.Completed,
            &i.Created_at, 
            &i.Updated_at, 
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}
func allPrvH(w http.ResponseWriter, rows *sql.Rows) (list []*ProvisionH, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(ProvisionH)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_hour,
            &i.En_hour,
            pq.Array(&i.S_hours),
            pq.Array(&i.E_hours),
            pq.Array(&i.Hours),
            &i.Completed,
            &i.Created_at, 
            &i.Updated_at, 
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}


func idPrvD(w http.ResponseWriter, conn *sql.DB, id int) (i ProvisionD, err error) {

    row := conn.QueryRow("SELECT * FROM provision_d WHERE id=$1", id)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.St_date,
        &i.En_date,
        pq.Array(&i.S_dates),
        pq.Array(&i.E_dates),
        pq.Array(&i.Dates),
        &i.Completed,
        &i.Created_at, 
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "err sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "err sql..! : %+v\n", err)
        return
    }

    return i,err
}
func idPrvH(w http.ResponseWriter, conn *sql.DB, id int) (i ProvisionH, err error) {

    row := conn.QueryRow("SELECT * FROM provision_h WHERE id=$1", id)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.St_hour,
        &i.En_hour,
        pq.Array(&i.S_hours),
        pq.Array(&i.E_hours),
        pq.Array(&i.Hours),
        &i.Completed,
        &i.Created_at, 
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "err sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "err sql..! : %+v\n", err)
        return
    }

    return i,err
}

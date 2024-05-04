package booking

import (
    "database/sql"
    "fmt"
    "net/http"

    "go_authentication/authtoken"
)


func allSearchD(w http.ResponseWriter, rows *sql.Rows) (list []*ProvisionD, err error) {

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
            &i.S_dates,
            &i.E_dates,
            &i.Dates,
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
func allSearcH(w http.ResponseWriter, rows *sql.Rows) (list []*ProvisionH, err error) {

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
            &i.S_hours,
            &i.E_hours,
            &i.Hours,
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


func allBkg(w http.ResponseWriter, rows *sql.Rows) (list []*Booking, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Booking)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.To_prv,
            &i.St_date,
            &i.En_date,
            &i.St_hour,
            &i.En_hour,
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


func idBkg(w http.ResponseWriter, conn *sql.DB, id int,cls *authtoken.Claims) (i Booking, err error) {

    owner := cls.User_id
    row := conn.QueryRow("SELECT * FROM booking WHERE id=$1 AND owner=$2", id,owner)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.St_date,
        &i.En_date,
        &i.St_hour,
        &i.En_hour,
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


func userBkg(w http.ResponseWriter, rows *sql.Rows) (list []*Booking, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Booking)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_date,
            &i.En_date,
            &i.St_hour,
            &i.En_hour,
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
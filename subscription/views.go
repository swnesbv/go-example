package subscription

import (
    "database/sql"
    "fmt"
    "net/http"

    "github.com/lib/pq"
)

// creat
func userIdSsc(w http.ResponseWriter, conn *sql.DB, id int) (i Subscription, err error) {

    row := conn.QueryRow("SELECT * FROM subscription WHERE id=$1", id)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.To_user,
        &i.To_group,
        &i.Completed,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, " Error: sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, " Error: sql..! : %+v\n", err)
        return
    }

    return i,err
}

// creat
func roomIdSsc(w http.ResponseWriter, conn *sql.DB, id int, owner int) (i Subscription, err error) {

    // list,err := qsAdminGroupSsc(w, conn,owner)
    // if err != nil {
    //     return
    // }

    admin,err := conn.Query("SELECT id FROM groups WHERE owner=$1", owner)
    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    defer admin.Close()

    var list []int
    for admin.Next() {
        i := new(Subscription)
        err = admin.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
        }
        list = append(list, i.Id)
    }
    
    row := conn.QueryRow("SELECT * FROM subscription WHERE id=$1 AND to_group = ANY($2)", id,pq.Array(list))
    
    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.To_user,
        &i.To_group,
        &i.Completed,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, " Error: sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, " Error: sql..! : %+v\n", err)
        return
    }

    return i,err
}
// ..


// list
func allSsc(w http.ResponseWriter, rows *sql.Rows) (list []*Subscription, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.To_user,
            &i.To_group,
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

func userSsc(w http.ResponseWriter, rows *sql.Rows) (list []*Subscription, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.To_user,
            &i.To_group,
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

func roomSsc(w http.ResponseWriter, rows *sql.Rows) (list []*Subscription, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.To_user,
            &i.To_group,
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
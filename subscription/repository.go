package subscription

import (
    "database/sql"
    "net/http"
    "fmt"
    
    "github.com/lib/pq"
)


func qsAllSsc(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM subscription")

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    
    return rows,err
}


func qsUserAllSsc(w http.ResponseWriter, conn *sql.DB, to_user int) (rows *sql.Rows, err error) {

    rows,err = conn.Query("SELECT * FROM subscription WHERE to_user=$1", to_user)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }

    return rows,err
}


func qsAdminGroupSsc(w http.ResponseWriter, conn *sql.DB, owner int) (list []int, err error) {

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

    for admin.Next() {
        i := new(Subscription)
        err = admin.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i.Id)
    }

    return list,err
}


func qsGroupAllSsc(w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

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
            return
        }
        list = append(list, i.Id)
    }
    
    rows,err = conn.Query("SELECT * FROM subscription WHERE to_group = ANY($1);", pq.Array(list))

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }

    return rows,err
}


    // rows, err = db.Query("SELECT s.id, s.title, s.description, s.owner, s.to_user, s.to_group, s.completed, s.created_at, s.updated_at FROM subscription AS s JOIN groups AS r ON s.to_group=r.owner WHERE r.owner=$1", owner)
    // rows, err = db.Query("SELECT s.* FROM subscription AS s JOIN groups AS r ON s.to_group=r.owner WHERE r.owner=$1", owner)
package chat

import (
    "database/sql"
    "net/http"
    "fmt"

    "github.com/lib/pq"
)


func qSscChUs(w http.ResponseWriter, id int) (list []int,err error) {

    row := db.QueryRow("SELECT id,owner,to_user,completed FROM subscription WHERE id=$1 AND completed=$2", id,true)

    var i Subscription
    err = row.Scan(&i.Id,&i.Owner,&i.To_user,&i.Completed)

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "user ErrNoRows..! : %+v\n", err)
    } else if err != nil {
        fmt.Fprintf(w, "user err sql..! : %+v\n", err)
    }

    list = append(list, i.Owner,i.To_user)
    return list,err
}


func qUsCh(w http.ResponseWriter, id int) (rows *sql.Rows,err error) {

    list,err := qSscChUs(w,id)
    if err != nil {
        return
    }

    rows,err = db.Query("SELECT id,coming,img,owner,to_user,completed,created_at,updated_at FROM msguser WHERE owner = ANY($1) AND to_user=$2 AND completed=$3", pq.Array(list),id,true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: qUsCh Query..! : %+v\n", err)
            break
        }
        return
    }

    return rows,err
}


func qSscGrChUs(w http.ResponseWriter, id int) (list []int,err error) {

    rows,err := db.Query("SELECT owner,to_group,completed FROM subscription WHERE to_group=$1 AND completed=$2", id,true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: qSscGrChUs Query..! : %+v\n", err)
            break
        }
        return
    }

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Owner,
            &i.To_group,
            &i.Completed,
        )
        if err != nil {
            fmt.Fprintf(w, "Error qSscGrChUs Scan..! : %+v\n", err)
            return
        }
        list = append(list, i.Owner)
    }
    fmt.Println(" user to group list..", list)
    return list,err
}

func qAdmin(w http.ResponseWriter, id int, owner int) (i Group) {

    row := db.QueryRow("SELECT id,owner,completed FROM groups WHERE id=$1 AND owner=$2 AND completed=$3", id,owner,true)

    _ = row.Scan(
        &i.Id,
        &i.Owner,
        &i.Completed,
    )
    return i
}



func qGroup(w http.ResponseWriter) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM groups WHERE completed=$1", true)

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


func qUsGroup(w http.ResponseWriter, owner int) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM groups WHERE owner=$1", owner)

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


func qGrChat(w http.ResponseWriter, to_group int) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM msggroups WHERE to_group=$1 AND completed=$2", to_group,true)

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


            // w.Header().Set("Content-Type", "text/html")
            // w.WriteHeader(http.StatusInternalServerError)
            // w.Write([]byte("500 - Something bad happened!"))
            // http.Error(w, "my own error message", 400)
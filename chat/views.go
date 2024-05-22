package chat

import (
    "database/sql"
    "fmt"
    // "os"
    "net/http"
    // "path/filepath"

    "github.com/lib/pq"
    // "go_authentication/authtoken"
)


func idUs(w http.ResponseWriter, list []int) (names []*ChUser, err error) {

    rows,err := db.Query("SELECT user_id,username,email FROM users WHERE user_id = ANY($1)", pq.Array(list))

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
    }

    defer rows.Close()
    for rows.Next() {
        i := new(ChUser)
        err = rows.Scan(
            &i.User_id,
            &i.Username,
            &i.Email,
        )
        if err != nil {
            fmt.Fprintf(w, "Error idUs Scan..! : %+v\n", err)
        }
        names = append(names, i)
    }
    return names,err
}


func usChat(w http.ResponseWriter, rows *sql.Rows) (list []*MsgUser, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(MsgUser)
        err = rows.Scan(
            &i.Id,
            &i.Coming,
            &i.Img,
            &i.Owner,
            &i.To_user,
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, "Error usChat Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}


func allGroup(w http.ResponseWriter, rows *sql.Rows) (list []*Group, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Group)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Img,
            &i.Owner,
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


func userGroup(w http.ResponseWriter, rows *sql.Rows) (list []*Group, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Group)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Img,
            &i.Owner,
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


func idGroup(w http.ResponseWriter, id int) (i Group, err error) {
    
    row := db.QueryRow("SELECT * FROM groups WHERE id=$1", id)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Img,
        &i.Owner,
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


func groupChat(w http.ResponseWriter, rows *sql.Rows,to_group int) (list []*MsgGroup,err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(MsgGroup)
        err = rows.Scan(
            &i.Id,
            &i.Coming,
            &i.Img,
            &i.Owner,
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
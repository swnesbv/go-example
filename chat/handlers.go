package chat

import (
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func UsChat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }
    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    // detail..
    detail,err := qSscChUs(w, id)
    if err != nil {
        return
    }
    in_users,err := idUs(w, detail)
    if err != nil {
        return
    }
    // ..

    // list..
    rows,err := qUsCh(w, id)
    if err != nil {
        return
    }
    list,err := usChat(w, rows)
    if err != nil {
        return
    }

    type ListData struct {
        Ssc int
        I []*MsgUser
        Uid int
        Uemail string
        T []*ChUser
    }
    data := ListData {
        Ssc: id,
        I: list,
        Uid: cls.User_id,
        Uemail: cls.Email,
        T: in_users,
    }

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/user.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }
}


func GrChat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }
    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    type ListData struct {
        I []*MsgGroup
        D Group
    }

    idname,err := idGroup(w, id)
    if err != nil {
        return
    }

    owner := cls.User_id
    user,err := qSscGrChUs(w, id)
    admin := qAdmin(w, id,owner)
    if err != nil {
        return
    }

    u := options.InSlice(owner,user)
    a := options.OkBool(owner,admin.Owner)
    if u || a {

        // msg..
        rows,err := qGrChat(w, id)
        if err != nil {
            return
        }
        list,err := groupChat(w, rows,id)
        if err != nil {
            return
        }
        // ..

        data := ListData {
            I: list,
            D: idname,
        }

        if r.Method == "GET" {
            tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/group.html", "./tpl/base.html" ))
            tpl.ExecuteTemplate(w, "base", data)
        }
    } else {
        fmt.Fprintf(w, " User No Group..! : %+v\n", err)
    }
}


func GrAll(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)

        if cls != nil && err == nil {

        rows,err := qGroup(w)
        if err != nil {
            return
        }
        list,err := allGroup(w, rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all_group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)

        }
    }
}
func GrOwr(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.ListToken(w,r)

        if cls != nil && err == nil {

        owner := cls.User_id

        rows,err := qUsGroup(w, owner)
        if err != nil {
            return
        }
        list,err := userGroup(w, rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all_owr_group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)

        } else {

        rows,err := qGroup(w)
        if err != nil {
            return
        }
        list,err := allGroup(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", list)

        }
    }
}


func DetGr(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/detail.html", "./tpl/base.html" ))

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }
        
        i,err := idGroup(w,id)
        if err != nil {
            return
        }

        cls := authtoken.WhoisWho(w,r)
        
        type ListData struct {
            Auth string
            I Group
        }
        if cls != nil {
            data := ListData {
                Auth: cls.Email,
                I: i,
            }
            fmt.Println("is cls")
            tpl.ExecuteTemplate(w, "base", data)
        } else {
            data := ListData {
                I: i,
            }
            fmt.Println("ne cls", cls)
            tpl.ExecuteTemplate(w, "base", data)
        }
    }
}


func HomeChat(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/index.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }
}
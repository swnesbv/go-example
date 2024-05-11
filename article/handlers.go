package article

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "runtime"

    "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
    "go_authentication/pagination"
)


func HomeArticle(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/index.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }
}


func Allarticle(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()

        lt,err := qArtCount(w, conn)
        if err != nil {
            return
        }
        count,err := ArtCount(w, lt)
        if err != nil {
            return
        }

        p,err := pagination.PageNumber(r)
        if err != nil {
            pagination.InvalidArgument(w)
            return
        }

        a := pagination.Args{
            Max:     5,
            Pos:     1,
            Page:    p,
            Records: 5,
            Total:   len(count),
            Size:    5,
        }

        limit := 5
        offset := limit * (a.Page - 1)
        rows,err := qArt(w, conn, limit,offset)
        if err != nil {
            return
        }
        list,err := allArt(w, rows)
        if err != nil {
            return
        }
        fmt.Println(" offset..", offset)
        defer conn.Close()

        pgn,err := pagination.NewPn(a)
        if err != nil {
            if err.Error() == pagination.ErrPageNo {
                log.Println(" NewPn..", err.Error())
                return
            }
            fmt.Println(err.Error())
            return
        }

        type ListPage struct {
            Articles []*Article
            Pgn *pagination.Pagination
        }
        view := ListPage{
            Articles: list,
            Pgn: pgn,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/pagination/pagination.html","./tpl/art/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", view)
    }

    fmt.Println(" All article goroutine..", runtime.NumGoroutine())
}


func UsAllArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        conn := connect.ConnSql()
        rows,err := qUsArt(w, conn,owner)
        if err != nil {
            return
        }
        list,err := userArt(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/author_id_article.html", "./tpl/base.html" ))
        
        tpl.ExecuteTemplate(w, "base", list)
    }
}


func DetArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }
        
        conn := connect.ConnSql()
        i,err := idArt(w, conn,id)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/detail.html", "./tpl/base.html" ))
        
        tpl.ExecuteTemplate(w, "base", i)
    }
}
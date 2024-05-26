package slider

import (
    "fmt"
    "time"
    "net/http"
    "html/template"

    "github.com/lib/pq"

    "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func CollectionAll(w http.ResponseWriter, r *http.Request) {

    cls, err := authtoken.OnToken(w, r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }
    owner := cls.User_id

    conn := connect.ConnSql()

    rowsArt,err := qArt(w, conn,owner)
    if err != nil {
        return
    }
    listArt,err := allArt(w, rowsArt)
    if err != nil {
        return
    }
    rowsSch,err := qSch(w, conn,owner)
    if err != nil {
        return
    }
    listSch,err := allSch(w, rowsSch)
    if err != nil {
        return
    }
    rowsPd,err := qPrvD(w, conn,owner)
    if err != nil {
        return
    }
    listPd,err := allPrvD(w, rowsPd)
    if err != nil {
        return
    }
    rowsPh,err := qPrvH(w, conn,owner)
    if err != nil {
        return
    }
    listPh,err := allPrvH(w, rowsPh)
    if err != nil {
        return
    }

    rows,err := qCollection(w, conn,owner)
    if err != nil {
        return
    }
    list,err := allCollection(w, rows)
    if err != nil {
        return
    }

    type ListSelect struct {
        List []*Collection
        Art  []*Article
        Sch  []*Schedule
        PrvD []*Provision_d
        PrvH []*Provision_h
    }
    view := ListSelect{
        List: list,
        Art:  listArt,
        Sch:  listSch,
        PrvD: listPd,
        PrvH: listPh,
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/all_collection.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", view)
    }


    if r.Method == "POST" {

        flag, _ := options.ParseBool(r.FormValue("img"))
        path := r.FormValue("path")
        if flag == true {
            path = r.FormValue("path")
        }

        var pfile []string
        pfile = append(pfile, path)
        fmt.Printf("list img.. %+v\n: ", pfile)
        
        title         := r.FormValue("title")
        description   := r.FormValue("description")
        collection_id := r.FormValue("collection_id")

        to_art   := ToNullInt64(r.FormValue("to_art"))
        to_sch   := ToNullInt64(r.FormValue("to_sch"))
        to_prv_d := ToNullInt64(r.FormValue("to_prv_d"))
        to_prv_h := ToNullInt64(r.FormValue("to_prv_h"))

        lt_t,pserr := psFormT(w,r)
        if pserr != nil {
            return
        }
        lt_d,pserr := psFormD(w,r)
        if pserr != nil {
            return
        }

        conn := connect.ConnSql()

        str := "INSERT INTO slider (collection_id, title, description, owner, to_art, to_sch, to_prv_d, to_prv_h, lt_t, lt_d, pfile, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"

        _, err := conn.Exec(
            str, collection_id,title,description,owner,to_art,to_sch,to_prv_d,to_prv_h,pq.Array(lt_t),pq.Array(lt_d),pq.Array(pfile),time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: CreatSlider Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/", http.StatusFound)
    }
}
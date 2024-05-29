package slider

import (
    "database/sql"
    "math/rand"
    "net/http"
    "strconv"
    "fmt"
    "os"
    "io"

    "github.com/lib/pq"

    "go_authentication/authtoken"
)


func psFormI (
    w http.ResponseWriter, r *http.Request, cls *authtoken.Claims, sid string) (list []string, err error) {

    pserr := r.ParseMultipartForm(10 * 1024 * 1024)
    if pserr != nil {
        fmt.Fprintf(w, " Error ParseMultipartForm..! : %+v\n", pserr)
        return
    }
    files := r.MultipartForm.File["file"]
    for _, i := range files {

        flname := i.Filename
        fpath := "./sfl/static/slider/" + cls.Email + "/" + sid + "/"
        fname := "./sfl/static/slider/" + cls.Email + "/" + sid + "/"  + flname
        fle := "/static/slider/" + cls.Email + "/" + sid + "/" + flname

        mkdirerr := os.MkdirAll(fpath, 0750)
        if mkdirerr != nil {
            fmt.Fprintf(w, " Error MkdirAll..! : %+v\n", mkdirerr)
        }
        img,err := os.Create(fname)
        if err != nil {
            fmt.Fprintf(w, " Error Create..! : %+v\n", err)
        }
        defer img.Close()

        readerFile, _ := i.Open()
        if _, err := io.Copy(img, readerFile); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        list = append(list, fle)
    }
    return list,err
}


 func ToNullInt64(s string) sql.NullInt64 {
   i, err := strconv.Atoi(s)
   return sql.NullInt64{Int64 : int64(i), Valid : err == nil}
 }

func randomString(len int) string {

    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randInt(97, 122))
    }
    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}


func allCollection(
    w http.ResponseWriter, rows *sql.Rows) (list []*Collection, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Collection)
        err = rows.Scan(
            &i.Id,
            &i.Collection_id,
            &i.Owner,
            pq.Array(&i.Pfile),
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}


func allArt(
    w http.ResponseWriter, rows *sql.Rows) (list []*Article, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Article)
        err = rows.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}
func allSch(
    w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Schedule)
        err = rows.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}
func allPrvD(
    w http.ResponseWriter, rows *sql.Rows) (list []*Provision_d, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Provision_d)
        err = rows.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}
func allPrvH(
    w http.ResponseWriter, rows *sql.Rows) (list []*Provision_h, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Provision_h)
        err = rows.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        list = append(list, i)
    }
    return list,err
}
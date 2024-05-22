package slider

import (
	"archive/zip"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/lib/pq"

	"go_authentication/connect"
	// "go_authentication/options"
	"go_authentication/authtoken"
)

func CreatCollection(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/collection.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", nil)
	}

	if r.Method == "POST" {

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(" err Data retrieving", err)
			return
		}

		rand.Seed(time.Now().UTC().UnixNano())
		sid := randomString(8)
		fpath := "./sfl/static/collection/" + cls.Email + "/" + sid + "/"

		archive, err := zip.NewReader(file, handler.Size)

		if err != nil {
			fmt.Fprintf(w, " Error: OpenReader..! : %+v\n", err)
			return
		}

		var list []string

		for _, f := range archive.File {
			fle := "/static/collection/" + cls.Email + "/" + sid + "/" + f.Name
			filePath := filepath.Join(fpath, f.Name)

			if f.FileInfo().IsDir() {
				fmt.Println("creating directory...")
				if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
					fmt.Fprintf(w, " Error: FileInfo..! : %+v\n", err)
				}
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				fmt.Fprintf(w, " Error: MkdirAll..! : %+v\n", err)
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				fmt.Fprintf(w, " Error: Mode..! : %+v\n", err)
			}

			srcFile, err := f.Open()
			if err != nil {
				fmt.Fprintf(w, " Error: Open..! : %+v\n", err)
			}
			if _, err := io.Copy(dstFile, srcFile); err != nil {
				fmt.Fprintf(w, " Error: Copy..! : %+v\n", err)
			}

			list = append(list, fle)
			dstFile.Close()
			srcFile.Close()
		}

		conn := connect.ConnSql()
		owner := cls.User_id
		str := "INSERT INTO collection (collection_id, owner, pfile, created_at) VALUES ($1,$2,$3,$4,$5)"

		_, adderr := conn.Exec(str, sid,owner,pq.Array(list),time.Now())

		if adderr != nil {
			fmt.Fprintf(w, " Error: adderr Exec..! : %+v\n", adderr)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-collection", http.StatusFound)
	}
}

func CreatSlider(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
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

    if r.Method == "GET" {

        type ListSelect struct {
            Art  []*Article
            Sch  []*Schedule
            PrvD []*Provision_d
            PrvH []*Provision_h
        }
        view := ListSelect{
            Art:  listArt,
            Sch:  listSch,
            PrvD: listPd,
            PrvH: listPh,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/slider.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", view)
    }

    if r.Method == "POST" {

        title         := r.FormValue("title")
        description   := r.FormValue("description")

        to_art   := ToNullInt64(r.FormValue("to_art"))
        to_sch	 := ToNullInt64(r.FormValue("to_sch"))
        to_prv_d := ToNullInt64(r.FormValue("to_prv_d"))
        to_prv_h := ToNullInt64(r.FormValue("to_prv_h"))

		rand.Seed(time.Now().UTC().UnixNano())
		sid := randomString(8)

        list,pserr := psForm(w,r, cls,sid)
        if pserr != nil {
            return
        }
        fmt.Printf("list %+v\n: ", list)

        conn := connect.ConnSql()
        
        str := "INSERT INTO slider (collection_id, title, description, owner, to_art, to_sch, to_prv_d, to_prv_h, pfile, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

        _, err := conn.Exec(
            str, sid,title,description,owner,to_art,to_sch,to_prv_d,to_prv_h,pq.Array(list),time.Now())

        if err != nil {
            fmt.Fprintf(w, " Error: CreatSlider Exec..! : %+v\n", err)
            return
        }

        defer conn.Close()
        http.Redirect(w,r, "/", http.StatusFound)
    }
}

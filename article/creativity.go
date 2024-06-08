package article

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"go_authentication/authtoken"
	"go_authentication/connect"
	"go_authentication/options"
)

func Creativity(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/creativity.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", nil)
	}

	if r.Method == "POST" {

		owner := cls.User_id
		title := r.FormValue("title")
		description := r.FormValue("description")

		conn := connect.ConnSql()
		str := "INSERT INTO article (title, description, owner, created_at) VALUES ($1,$2,$3,$4)"

		_, err := conn.Exec(str, title, description, owner, time.Now())
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/author-id-article", http.StatusFound)
	}
}

func UpArt(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.SqlToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	conn := connect.ConnSql()
	art, err := authorArt(w, conn, id, cls)
	if err != nil { return }

	flag, err := options.ParseBool(r.FormValue("completed"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, " Error: ParseBool..  : %+v\n", err)
		return
	}

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/update.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", art)
	}

	if r.Method == "POST" {

		owner := cls.User_id
		title := r.FormValue("title")
		description := r.FormValue("description")

		str := "UPDATE article SET title=$3, description=$4, completed=$5, updated_at=$6 WHERE id=$1 AND owner=$2"

		_, err := conn.Exec(str, id, owner, title, description, flag, time.Now())
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/author-id-article", http.StatusFound)
	}
}

func DelArt(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	if r.Method == "GET" {

		data := struct {
			Items string
		}{
			Items: cls.Email,
		}

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/delete.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", data)
	}

	if r.Method == "POST" {

		sid := r.URL.Query().Get("id")
		fpath := "./static/img/art/" + cls.Email + "/" + sid

		p := options.DelFolder(fpath)
		if p != nil {
			fmt.Fprintf(w, "del path img..! : %+v\n", p)
			return
		}

		owner := cls.User_id
		conn := connect.ConnSql()
		str := "DELETE FROM article WHERE id=$1 AND owner=$2"

		_, err := conn.Exec(str, id, owner)
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/author-id-article", http.StatusFound)
	}
}

func ImgArt(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.SqlToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	conn := connect.ConnSql()
	owner := cls.User_id

	i, err := authorArt(w, conn, id, cls)
	if err != nil { return }

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/img.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", i)
	}

	if r.Method == "POST" {

		if i.Img != nil {
			err := os.Remove("." + *i.Img)
			if err != nil {
				fmt.Fprintf(w, " Error: Remove..! : %+v\n", err)
			}
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Fprintf(w, "Error Data retrieving..! : %+v\n", err)
			return
		}
		defer file.Close()

		flname := handler.Filename
		sid := r.URL.Query().Get("id")

		fpath := "./sfl/static/img/art/" + cls.Email + "/" + sid + "/"
		fname := "./sfl/static/img/art/" + cls.Email + "/" + sid + "/" + flname
		fle := "/static/img/art/" + cls.Email + "/" + sid + "/" + flname

		fmt.Printf("fle %+v\n: ", fle)
		fmt.Printf("Uploaded File : %+v\n", flname)
		fmt.Printf("File Size : %+v\n", handler.Size)
		fmt.Printf("MIME Header : %+v\n", handler.Header)

		mkdirerr := os.MkdirAll(fpath, 0750)
		if mkdirerr != nil {
			fmt.Fprintf(w, "Error MkdirAll..! : %+v\n", mkdirerr)
		}
		img, err := os.Create(fname)
		if err != nil {
			fmt.Fprintf(w, "Error Create..! : %+v\n", err)
		}
		defer img.Close()

		if _, err := io.Copy(img, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		str := "UPDATE article SET img=$3, updated_at=$4 WHERE id=$1 AND owner=$2"

		_, err = conn.Exec(str, id, owner, fle, time.Now())
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, fname, http.StatusFound)
	}
}

func DelImgArt(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	conn := connect.ConnSql()
	owner := cls.User_id

	i, err := authorArt(w, conn, id, cls)
	if err != nil { return }

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/imgdel.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", i)
	}

	if r.Method == "POST" {

		e := os.Remove("." + *i.Img)
		if e != nil {
			fmt.Println("e.. ", e)
		}

		str := "UPDATE article SET img=$3, updated_at=$4 WHERE id=$1 AND owner=$2"
		
		_, err = conn.Exec(str, id, owner, nil, time.Now())
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/author-id-article", http.StatusFound)
	}
}

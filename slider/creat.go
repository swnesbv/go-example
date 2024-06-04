package slider

import (
	"archive/zip"
	"fmt"
	"github.com/lib/pq"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"go_authentication/authtoken"
	"go_authentication/connect"
	"go_authentication/options"
)

func CreatCollection(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/add_collection.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", nil)
	}

	if r.Method == "POST" {

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(" err Data retrieving", err)
			return
		}

		rand.Seed(time.Now().UTC().UnixNano())
		sid := randomString(4)
		fpath := "./sfl/static/collection/" + cls.Email + "/" + sid + "/"

		archive, err := zip.NewReader(file, handler.Size)

		if err != nil {
			fmt.Fprintf(w, " Error: OpenReader..! : %+v\n", err)
			return
		}

		list := make([]string, len(archive.File))

		for k, f := range archive.File {
			fle := "/static/collection/" + cls.Email + "/" + sid + "/" + f.Name
			filePath := filepath.Join(fpath, f.Name)

			if f.FileInfo().IsDir() {
				fmt.Println("creating directory..")
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

			list[k] = fle
			dstFile.Close()
			srcFile.Close()
		}

		conn := connect.ConnSql()
		owner := cls.User_id
		str := "INSERT INTO collection (collection_id, owner, pfile, created_at) VALUES ($1,$2,$3,$4)"

		_, adderr := conn.Exec(str, sid, owner, pq.Array(list), time.Now())

		if adderr != nil {
			fmt.Fprintf(w, " Error: adderr Exec..! : %+v\n", adderr)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-collection", http.StatusFound)
	}
}

func CreatSlider(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	owner := cls.User_id

	conn := connect.ConnSql()

	rowsArt, err := qArt(w, conn, owner)
	if err != nil { return }
	listArt, err := allArt(w, rowsArt)
	if err != nil { return }
	rowsSch, err := qSch(w, conn, owner)
	if err != nil { return }
	listSch, err := allSch(w, rowsSch)
	if err != nil { return }
	rowsPd, err := qPrvD(w, conn, owner)
	if err != nil { return }
	listPd, err := allPrvD(w, rowsPd)
	if err != nil { return }
	rowsPh, err := qPrvH(w, conn, owner)
	if err != nil { return }
	listPh, err := allPrvH(w, rowsPh)
	if err != nil { return }

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

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/slider.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {

		title := r.FormValue("title")
		description := r.FormValue("description")

		to_art := ToNullInt64(r.FormValue("to_art"))
		to_sch := ToNullInt64(r.FormValue("to_sch"))
		to_prv_d := ToNullInt64(r.FormValue("to_prv_d"))
		to_prv_h := ToNullInt64(r.FormValue("to_prv_h"))

		rand.Seed(time.Now().UTC().UnixNano())
		sid := randomString(8)

		lt_t, pserr := options.PsFormString(w, r, "lt_t")
		lt_d, pserr := options.PsFormString(w, r, "lt_d")
		pfile, pserr := psFormI(w, r, cls, sid)
		if pserr != nil { return }

		str := "INSERT INTO slider (collection_id, title, description, owner, to_art, to_sch, to_prv_d, to_prv_h, lt_t, lt_d, pfile, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"

		_, err := conn.Exec(
			str, sid, title, description, owner, to_art, to_sch, to_prv_d, to_prv_h, pq.Array(lt_t), pq.Array(lt_d), pq.Array(pfile), time.Now())

		if err != nil {
			fmt.Fprintf(w, " Error: CreatSlider Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func UpSlKey(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.SqlToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	owner := cls.User_id

	conn := connect.ConnSql()

	rowsArt, err := qArt(w, conn, owner)
	if err != nil { return }
	listArt, err := allArt(w, rowsArt)
	if err != nil { return }
	rowsSch, err := qSch(w, conn, owner)
	if err != nil { return }
	listSch, err := allSch(w, rowsSch)
	if err != nil { return }
	rowsPd, err := qPrvD(w, conn, owner)
	if err != nil { return }
	listPd, err := allPrvD(w, rowsPd)
	if err != nil { return }
	rowsPh, err := qPrvH(w, conn, owner)
	if err != nil { return }
	listPh, err := allPrvH(w, rowsPh)
	if err != nil { return }

	i, err := authorSl(w, conn, id, cls)
	if err != nil { return }
	type ListSelect struct {
		I    *Slider
		Art  []*Article
		Sch  []*Schedule
		PrvD []*Provision_d
		PrvH []*Provision_h
	}
	view := ListSelect{
		I:    i,
		Art:  listArt,
		Sch:  listSch,
		PrvD: listPd,
		PrvH: listPh,
	}

	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/update_key.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {

		to_art := ToNullInt64(r.FormValue("to_art"))
		to_sch := ToNullInt64(r.FormValue("to_sch"))
		to_prv_d := ToNullInt64(r.FormValue("to_prv_d"))
		to_prv_h := ToNullInt64(r.FormValue("to_prv_h"))

		flag, err := options.ParseBool(r.FormValue("completed"))
		if err != nil {
			fmt.Fprintf(w, " Error: ParseBool()..  : %+v\n", err)
			return
		}
		str := "UPDATE slider SET to_art=$3, to_sch=$4, to_prv_d=$5, to_prv_h=$6, completed=$7, updated_at=$8 WHERE id=$1 AND owner=$2"

		switch {
		case r.FormValue("to_art") == "":
			_, err := conn.Exec(str, id, owner, i.To_art, to_sch, to_prv_d, to_prv_h, flag, time.Now())
			fmt.Println("  to_art..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		case r.FormValue("to_sch") == "":
			_, err := conn.Exec(str, id, owner, to_art, i.To_sch, to_prv_d, to_prv_h, flag, time.Now())
			fmt.Println("  to_sch..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		case r.FormValue("to_prv_d") == "":
			_, err := conn.Exec(str, id, owner, to_art, to_sch, i.To_prv_d, to_prv_h, flag, time.Now())
			fmt.Println("  to_prv_d..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		case r.FormValue("to_prv_h") == "":
			_, err := conn.Exec(str, id, owner, to_art, to_sch, to_prv_d, i.To_prv_d, flag, time.Now())
			fmt.Println("  to_prv_h..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		default:
			_, err := conn.Exec(str, id, owner, to_art, to_sch, to_prv_d, to_prv_h, flag, time.Now())
			fmt.Println("  default..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		}

		defer conn.Close()
		http.Redirect(w, r, "/detail-slider?id="+strconv.Itoa(i.Id), http.StatusFound)
	}
}

func UpSlText(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.SqlToken(w, r)
	if cls == nil { return }
	if err != nil { return }
	owner := cls.User_id

	conn := connect.ConnSql()

	rowsArt, err := qArt(w, conn, owner)
	if err != nil { return }
	listArt, err := allArt(w, rowsArt)
	if err != nil { return }
	rowsSch, err := qSch(w, conn, owner)
	if err != nil { return }
	listSch, err := allSch(w, rowsSch)
	if err != nil { return }
	rowsPd, err := qPrvD(w, conn, owner)
	if err != nil { return }
	listPd, err := allPrvD(w, rowsPd)
	if err != nil { return }
	rowsPh, err := qPrvH(w, conn, owner)
	if err != nil { return }
	listPh, err := allPrvH(w, rowsPh)
	if err != nil { return }

	i, err := authorSl(w, conn, id, cls)
	if err != nil { return }

	if r.Method == "GET" {
		type ListSelect struct {
			I    *Slider
			Art  []*Article
			Sch  []*Schedule
			PrvD []*Provision_d
			PrvH []*Provision_h
		}
		view := ListSelect{
			I:    i,
			Art:  listArt,
			Sch:  listSch,
			PrvD: listPd,
			PrvH: listPh,
		}
		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/update_text.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")

		lt_t, err := options.PsFormString(w, r, "lt_t")
		lt_d, err := options.PsFormString(w, r, "lt_d")
		if err != nil { return }

		pserr := r.ParseForm()
		if pserr != nil {
			fmt.Println(" err ParseForm..", pserr)
		}
		on_off_t := make([]bool, len(r.Form["del_t"]))
		for k, v := range r.Form["del_t"] {
			p, _ := options.ParseBool(v)
			on_off_t[k] = p
		}
		on_off_d := make([]bool, len(r.Form["del_d"]))
		for k, v := range r.Form["del_d"] {
			p, _ := options.ParseBool(v)
			on_off_d[k] = p
		}
		ps_t := delOk(on_off_t)
		ps_d := delOk(on_off_d)

		flag, err := options.ParseBool(r.FormValue("completed"))
		if err != nil {
			fmt.Fprintf(w, " Error: ParseBool()..  : %+v\n", err)
			return
		}

		str := "UPDATE slider SET title=$3, description=$4, lt_t=$5, lt_d=$6, completed=$7, updated_at=$8 WHERE id=$1 AND owner=$2"

		switch {
		case r.FormValue("lt_t") == "" && r.FormValue("lt_d") == "":
			_, err := conn.Exec(str, id, owner, title, description, pq.Array(i.Lt_t), pq.Array(i.Lt_d), flag, time.Now())
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}

		case ps_t == true && ps_d == true:
			del_t,err := psDelStr(w,r, lt_t,"del_t")
			if err != nil { return }
			t := delList(lt_t, del_t)
			del_d,err := psDelStr(w,r, lt_d,"del_d")
			if err != nil { return }
			d := delList(lt_d, del_d)
			_, crash := conn.Exec(str, id, owner, title, description, pq.Array(t), pq.Array(d), flag, time.Now())
			if crash != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", crash)
				return
			}

		case ps_t == true:
			del_t,err := psDelStr(w,r, lt_t,"del_t")
			if err != nil { return}
			t := delList(lt_t, del_t)
			_, crash := conn.Exec(str, id, owner, title, description, pq.Array(t), pq.Array(lt_d), flag, time.Now())
			if crash != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", crash)
				return
			}
		case ps_d == true:
			del_d,err := psDelStr(w,r, lt_d,"del_d")
			if err != nil { return }
			d := delList(lt_d, del_d)
			_, crash := conn.Exec(str, id, owner, title, description, pq.Array(lt_t), pq.Array(d), flag, time.Now())
			if crash != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", crash)
				return
			}

		default:
			_, crash := conn.Exec(str, id, owner, title, description, pq.Array(lt_t), pq.Array(lt_d), flag, time.Now())
			if crash != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", crash)
				return
			}
		}
		defer conn.Close()
		http.Redirect(w, r, "/detail-slider?id="+strconv.Itoa(i.Id), http.StatusFound)
	}
}

func UpSlImg(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.SqlToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	owner := cls.User_id

	conn := connect.ConnSql()

	rowsArt, err := qArt(w, conn, owner)
	if err != nil { return }
	listArt, err := allArt(w, rowsArt)
	if err != nil { return }
	rowsSch, err := qSch(w, conn, owner)
	if err != nil { return }
	listSch, err := allSch(w, rowsSch)
	if err != nil { return }
	rowsPd, err := qPrvD(w, conn, owner)
	if err != nil { return }
	listPd, err := allPrvD(w, rowsPd)
	if err != nil { return }
	rowsPh, err := qPrvH(w, conn, owner)
	if err != nil { return }
	listPh, err := allPrvH(w, rowsPh)
	if err != nil { return }

	i, err := authorSl(w, conn, id, cls)
	if err != nil { return }

	if r.Method == "GET" {
		type ListSelect struct {
			I    *Slider
			Art  []*Article
			Sch  []*Schedule
			PrvD []*Provision_d
			PrvH []*Provision_h
		}
		view := ListSelect{
			I:    i,
			Art:  listArt,
			Sch:  listSch,
			PrvD: listPd,
			PrvH: listPh,
		}
		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/update_img.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {

		delfl, err := psDelImg(w, r)
		if err != nil { return }

		for _, v := range delfl {
			err := os.Remove("./sfl" + v)
			if err != nil {
				fmt.Println(" err when deleting file..", err)
			}
		}
		obj := delList(i.Pfile, delfl)

		pfile, err := psFormI(w, r, cls, i.Collection_id)
		if err != nil { return }

		flag, err := options.ParseBool(r.FormValue("completed"))
		if err != nil {
			fmt.Fprintf(w, " Error: ParseBool..  : %+v\n", err)
			return
		}

		dbstr := "UPDATE slider SET pfile=$3, completed=$4, updated_at=$5 WHERE id=$1 AND owner=$2"
		str := "UPDATE slider SET pfile=array_cat(pfile, $3), completed=$4, updated_at=$5 WHERE id=$1 AND owner=$2"

		switch {
		case pfile == nil && delfl == nil:
			_, err := conn.Exec(dbstr, id, owner, pq.Array(i.Pfile), flag, time.Now())
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		case pfile != nil && delfl != nil:
			_, err := conn.Exec(dbstr, id, owner, pq.Array(obj), flag, time.Now())
			_, err = conn.Exec(str, id, owner, pq.Array(pfile), flag, time.Now())
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		case pfile == nil && delfl != nil:
			_, err := conn.Exec(dbstr, id, owner, pq.Array(obj), flag, time.Now())
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}

		default:
			_, err := conn.Exec(str, id, owner, pq.Array(pfile), flag, time.Now())
			fmt.Println("  default..!")
			if err != nil {
				fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
				return
			}
		}

		defer conn.Close()
		http.Redirect(w, r, "/detail-slider?id="+strconv.Itoa(i.Id), http.StatusFound)
	}
}

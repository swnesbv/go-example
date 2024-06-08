package slider

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/lib/pq"

	"go_authentication/authtoken"
	"go_authentication/connect"
	"go_authentication/options"
)

func CollectionAll(w http.ResponseWriter, r *http.Request) {

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

	rows, err := qCollection(w, conn, owner)
	if err != nil { return }
	list, err := allCollection(w, rows)
	if err != nil { return }

	if r.Method == "GET" {

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
		tpl := template.Must(
			template.ParseFiles(
				"./tpl/navbar.html",
				"./tpl/slider/all_collection.html",
				"./tpl/base.html",
			),
		)
		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {

		crash := r.ParseForm()
		if crash != nil {
			fmt.Fprintf(w, " Error ParseForm..! : %+v\n", crash)
			return
		}

		on_off := make([]bool, len(r.Form["img"]))
		pfile  := make([]string, 0, len(r.Form["path"]))

		for k, v := range r.Form["img"] {
			flag, _ := options.ParseBool(v)
			on_off[k] = flag
		}
		ps_img := delOk(on_off)
		if ps_img == true {
			file := r.Form["path"]
			for k := range on_off {
				if on_off[k] == true {
					pfile = append(pfile, file[k])
				}
			}
		} else {
			pfile = nil
		}

		title         := r.FormValue("title")
		description   := r.FormValue("description")
		collection_id := r.FormValue("collection_id")
		to_collection := r.FormValue("to_collection")

		to_art   := ToNullInt64(r.FormValue("to_art"))
		to_sch   := ToNullInt64(r.FormValue("to_sch"))
		to_prv_d := ToNullInt64(r.FormValue("to_prv_d"))
		to_prv_h := ToNullInt64(r.FormValue("to_prv_h"))

		lt_t, crash := options.PsFormString(w, r, "lt_t")
		if crash != nil { return }
		lt_d, crash := options.PsFormString(w, r, "lt_d")
		if crash != nil { return }

		conn := connect.ConnSql()

		str := "INSERT INTO slider (collection_id, title, description, owner, to_collection, to_art, to_sch, to_prv_d, to_prv_h, lt_t, lt_d, pfile, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"

		_, err := conn.Exec(
			str,
			collection_id,
			title,
			description,
			owner,
			to_collection,
			to_art,
			to_sch,
			to_prv_d,
			to_prv_h,
			pq.Array(lt_t),
			pq.Array(lt_d),
			pq.Array(pfile),
			time.Now(),
		)

		if err != nil {
			fmt.Fprintf(w, " Error: CreatSlider Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-slider", http.StatusFound)
	}
}

func UpColl(w http.ResponseWriter, r *http.Request) {

	id, err := options.IdUrl(w, r)
	if err != nil { return }
	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	owner := cls.User_id
	conn := connect.ConnSql()

	s, err := authorSl(w, conn, id, owner)
	if err != nil { return }
	c, err := authorColl(w, conn, s.To_collection, owner)
	if err != nil { return }

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
			S *Slider
			C *Collection
			Art  []*Article
			Sch  []*Schedule
			PrvD []*Provision_d
			PrvH []*Provision_h
		}
		view := ListSelect{
			S: s,
			C: c,
			Art:  listArt,
			Sch:  listSch,
			PrvD: listPd,
			PrvH: listPh,
		}
		tpl := template.Must(
			template.ParseFiles(
				"./tpl/navbar.html",
				"./tpl/slider/update_collection.html",
				"./tpl/base.html",
			),
		)
		tpl.ExecuteTemplate(w, "base", view)
	}

	if r.Method == "POST" {

		crash := r.ParseForm()
		if crash != nil {
			fmt.Fprintf(w, " Error ParseForm..! : %+v\n", crash)
			return
		}

		on_off := make([]bool, len(r.Form["img"]))
		pfile  := make([]string, 0, len(r.Form["path"]))

		for k, v := range r.Form["img"] {
			flag, _ := options.ParseBool(v)
			on_off[k] = flag
		}
		ps_img := delOk(on_off)
		if ps_img == true {
			file := r.Form["path"]
			for k := range on_off {
				if on_off[k] == true {
					pfile = append(pfile, file[k])
				}
			}
		} else {
			pfile = s.Pfile
		}

		flag, crash := options.ParseBool(r.FormValue("completed"))
		if crash != nil {
			fmt.Fprintf(w, " Error: ParseBool..  : %+v\n", crash)
			return
		}

		str := "UPDATE slider SET pfile=$3, completed=$4, updated_at=$5 WHERE id=$1 AND owner=$2"

		err := selectImg(
			w,
			r,
			conn,
			s,
			str,
			pfile,
			id,
			owner,
			flag,
		)

		if err != nil {
			fmt.Fprintf(w, " Error: CreatSlider Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-slider", http.StatusFound)
	}
}


func AllSlider(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		cls, err := authtoken.OnToken(w, r)
		if cls == nil { return }
		if err != nil { return }
		
		owner := cls.User_id
		conn := connect.ConnSql()

		rows, err := qSlider(w, conn, owner)
		if err != nil { return }
		list, err := allSl(w, rows)
		if err != nil { return }

		defer conn.Close()

		tpl := template.Must(
			template.ParseFiles("./tpl/navbar.html", "./tpl/slider/all_slider.html", "./tpl/base.html"),
		)
		tpl.ExecuteTemplate(w, "base", list)
	}
}

func DetSlider(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		id, err := options.IdUrl(w, r)
		if err != nil { return }
		cls, err := authtoken.OnToken(w, r)
		if cls == nil { return }
		if err != nil { return }

		owner := cls.User_id
		conn := connect.ConnSql()
		
		i, err := authorSl(w, conn, id, owner)
		if err != nil { return }

		defer conn.Close()

		tpl := template.Must(
			template.ParseFiles("./tpl/navbar.html", "./tpl/slider/detail.html", "./tpl/base.html"),
		)
		tpl.ExecuteTemplate(w, "base", i)
	}
}

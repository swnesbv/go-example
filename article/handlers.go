package article

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go_authentication/authtoken"
	"go_authentication/connect"
	"go_authentication/options"
	"go_authentication/pagination"
)

func HomeArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/index.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", nil)
	}
}

func Allarticle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		conn := connect.ConnSql()

		count, err := qArtCount(w, conn)
		if err != nil {
			return
		}
		fmt.Println(" count..", count)

		p, err := pagination.PageNumber(r)
		if err != nil {
			pagination.InvalidArgument(w)
			return
		}

		a := pagination.Args{
			Max:     5,
			Pos:     1,
			Page:    p,
			Records: 5,
			Total:   count,
			Size:    5,
		}
		limit := 5
		offset := limit * (a.Page - 1)

		rows, err := qArt(w, conn, limit, offset)
		if err != nil {
			return
		}
		list, err := allArt(w, rows)
		if err != nil {
			return
		}
		// fmt.Printf("%s\n", list)
		fmt.Println(" offset..", offset)

		defer conn.Close()

		pgn, _ := pagination.NewPn(a)
		type ListPgn struct {
			Articles []*Article
			Pgn      *pagination.Pagination
		}
		view := ListPgn{
			Articles: list,
			Pgn:      pgn,
		}
		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/pagination/pagination.html", "./tpl/art/all.html", "./tpl/base.html"))
		tpl.ExecuteTemplate(w, "base", view)

	}
}

func UsAllArt(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		cls, err := authtoken.OnToken(w, r)
		if cls == nil {
			return
		}
		if err != nil {
			return
		}

		owner := cls.User_id
		conn := connect.ConnSql()
		rows, err := qUsArt(w, conn, owner)
		if err != nil {
			return
		}
		list, err := userArt(w, rows)
		if err != nil {
			return
		}
		defer conn.Close()

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/author_id_article.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", list)
	}
}

func (u *LSl) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &[]any{
		&u.Id,
		&u.T,
		&u.D,
		&u.F,
	})
}
func DetArt(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		id, err := options.IdUrl(w, r)
		if err != nil {
			return
		}

		conn := connect.ConnSql()
		i, err := idArt(w, conn, id)
		if err != nil {
			return
		}
		s := idSlider(w, conn, id)

		var s1 []string = s.Lt_t
		var s2 []string = s.Lt_d
		var s3 []string = s.Pfile

		list  := make([]string, 0, 40)
		if s.Pfile == nil {
			for k := range s1 {
				list = append(list, strconv.Itoa(k), s1[k], s2[k], "")
			}
		} else {
			for k := range s1 {
				list = append(list, strconv.Itoa(k), s1[k], s2[k], s3[k])
			}
		}

		// fmt.Println(" list..", list)

		section := Section(list, 4)
		// fmt.Println(" section..", section)
		// fmt.Printf(" f section : %#v\n", section)

		encjson, _ := json.Marshal(section)
		// fmt.Println(" M..", string(encjson))

		u := []LSl{}
		if err := json.Unmarshal([]byte(encjson), &u); err != nil {
			fmt.Fprintf(w, " Error Unmarshal..! : %+v\n", err)
		}
		fmt.Printf(" u.. %+v\n", u)

		type ListSelect struct {
			Art Article
			Sl  []LSl
		}
		view := ListSelect{
			Art: i,
			Sl:  u,
		}

		defer conn.Close()

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/slider/slider_css.html", "./tpl/slider/slider_js.html", "./tpl/slider/index.html", "./tpl/art/detail.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", view)
	}
}

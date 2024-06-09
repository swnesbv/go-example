package schedule

import (
	"fmt"
	"github.com/lib/pq"
	"html/template"
	"net/http"
	"time"

	"go_authentication/connect"
	// "go_authentication/options"
	"go_authentication/authtoken"
)

func Creat(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	if r.Method == "GET" {

		tpl := template.Must(
			template.ParseFiles(
				"./tpl/navbar.html",
				"./tpl/schedule/creat.html",
				"./tpl/base.html",
			),
		)
		tpl.ExecuteTemplate(w, "base", nil)
	}

	if r.Method == "POST" {

		title := r.FormValue("title")
		description := r.FormValue("description")

		loc, _ := time.LoadLocation("UTC")
		start, crash := time.ParseInLocation(
			"2006-01-02T15:04:05", r.FormValue("start")+":00", loc)
		end, crash := time.ParseInLocation(
			"2006-01-02T15:04:05", r.FormValue("end")+":00", loc)
		if crash != nil {
			fmt.Fprintf(w, " Error: ParseInLocation..! : %+v\n", crash)
			return
		}

		list, crash := psForm(w, r)
		if crash != nil { return }

		conn := connect.ConnSql()
		owner := cls.User_id
		str := "INSERT INTO schedule (title,description,owner,st_hour,en_hour,hours,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"

		_, err := conn.Exec(
			str, title, description, owner, start, end, pq.Array(list), time.Now())

		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-schedule", http.StatusFound)
	}
}

package schedule

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/lib/pq"

	// "go_authentication/options"
	"go_authentication/authtoken"
	"go_authentication/connect"
)

func SchSelect(w http.ResponseWriter, r *http.Request) {

	cls, err := authtoken.OnToken(w, r)
	if cls == nil { return }
	if err != nil { return }

	conn := connect.ConnSql()

	rows, err := qSchSelect(w, conn)
	if err != nil { return }
	list, err := allSelect(w, rows)
	if err != nil { return }

	if r.Method == "GET" {
		tpl := template.Must(
			template.ParseFiles(
				"./tpl/navbar.html",
				"./tpl/schedule/select.html",
				"./tpl/base.html",
			),
		)
		tpl.ExecuteTemplate(w, "base", list)
	}

	if r.Method == "POST" {

		id := r.FormValue("to_schedule")

		fmt.Println(" FormValue..", r.FormValue("hour"))
		date, crash := time.Parse(
			time.DateTime, r.FormValue("date"))
        if crash != nil {
            fmt.Fprintf(w, " Error: D Parse..! : %+v\n", crash)
            return
        }
		hour, crash := time.Parse(
			time.TimeOnly, r.FormValue("hour"))
		if crash != nil {
			fmt.Fprintf(w, " Error: H Parse..! : %+v\n", crash)
			return
		}

		var list []time.Time
		list = append(list, hour)

		owner := cls.User_id
		str := "INSERT INTO recording (owner,to_schedule,record_d,record_h,created_at) VALUES ($1,$2,$3,$4,$5)"
		_, err := conn.Exec(
			str, owner, id, date, hour, time.Now())
		if err != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", err)
			return
		}

		prv := "UPDATE schedule SET occupied=array_cat(occupied, $2),completed=$3,updated_at=$4 WHERE id=$1"
		_, perr := conn.Exec(prv, id, pq.Array(list), false, time.Now())
		if perr != nil {
			fmt.Fprintf(w, " Error: Exec..! : %+v\n", perr)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/all-recording", http.StatusFound)
	}

}

func SchAll(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		conn := connect.ConnSql()
		rows, err := qSch(w, conn)
		if err != nil {
			return
		}
		list, err := allSch(w, rows)
		if err != nil {
			return
		}

		defer conn.Close()

		tpl := template.Must(
			template.ParseFiles("./tpl/navbar.html", "./tpl/schedule/all.html", "./tpl/base.html"),
		)

		tpl.ExecuteTemplate(w, "base", list)
	}
}

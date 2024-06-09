package schedule

import (
	"database/sql"
	"fmt"
	"time"
	//"sync"
	"net/http"
	"runtime"

	"github.com/lib/pq"
	// "go_authentication/authtoken"
)

func Converter(w http.ResponseWriter, v string) time.Time {

	loc, _ := time.LoadLocation("UTC")
	to, _ := time.ParseInLocation("2006-01-02T15:04:05", v+":00", loc)

	s := to.Format(time.TimeOnly)
	ss, _ := time.Parse(time.TimeOnly, s)
	return ss
}
func psForm(w http.ResponseWriter, r *http.Request) ([]time.Time, error) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error ParseForm..! : %+v\n", err)
	}

	ps := r.Form["list"]

	on_off := make([]string, len(ps))
	for k, v := range ps {
		on_off[k] = v
	}

	list := make([]time.Time, 0, len(ps))
	for k, v := range ps {
		if on_off[k] != "" {
			ss := Converter(w, v)
			list = append(list, ss)
		}
	}
	return list, err
}

func selection(hours, occupied []string) (slt []string) {
	m := make(map[string]bool)

	for _, i := range occupied {
		m[i] = true
	}

	for _, i := range hours {
		if _, ok := m[i]; !ok {
			slt = append(slt, i)
		}
	}
	fmt.Println(" slt..", slt)
	return slt
}

/*func allSelect (
    w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

    start := time.Now()
    defer rows.Close()
    for rows.Next() {
        i := new(Schedule)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.St_hour,
            &i.En_hour,
            pq.Array(&i.Hours),
            pq.Array(&i.Occupied),
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
            return
        }
        free := selection(i.Hours,i.Occupied)
        i.Hours = free
        list = append(list, i)
    }

    fmt.Println(" sel goroutine..", runtime.NumGoroutine())
    elapsed := time.Since(start)
    fmt.Println(" sel time..", elapsed)

    return list,err
}*/

func scanning(w http.ResponseWriter, rows *sql.Rows, ch chan *Schedule) {
	i := new(Schedule)
	err := rows.Scan(
		&i.Id,
		&i.Title,
		&i.Description,
		&i.Owner,
		&i.St_hour,
		&i.En_hour,
		pq.Array(&i.Hours),
		pq.Array(&i.Occupied),
		&i.Completed,
		&i.Created_at,
		&i.Updated_at,
	)
	if err != nil {
		fmt.Fprintf(w, " Error Scan..! : %+v\n", err)
		return
	}
	ch <- i
}
func allSelect(w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

	start := time.Now()

	defer rows.Close()

	for rows.Next() {
		ch := make(chan *Schedule)
		go scanning(w, rows, ch)
		c := <-ch
		free := selection(c.Hours, c.Occupied)
		c.Hours = free
		list = append(list, c)
	}

	fmt.Println(" sel goroutine..", runtime.NumGoroutine())
	elapsed := time.Since(start)
	fmt.Println(" sel time..", elapsed)

	return list, err
}

func allSch(w http.ResponseWriter, rows *sql.Rows) (list []*Schedule, err error) {

	defer rows.Close()
	for rows.Next() {
		i := new(Schedule)
		err = rows.Scan(
			&i.Id,
			&i.Title,
			&i.Description,
			&i.Owner,
			&i.St_hour,
			&i.En_hour,
			pq.Array(&i.Hours),
			pq.Array(&i.Occupied),
			&i.Completed,
			&i.Created_at,
			&i.Updated_at,
		)
		if err != nil {
			fmt.Fprintf(w, "Error Scan..! : %+v\n", err)
			return
		}
		list = append(list, i)
	}
	fmt.Println(" allSch goroutine..", runtime.NumGoroutine())
	return list, err
}

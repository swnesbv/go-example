package slider

import (
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/lib/pq"

	"go_authentication/authtoken"
	"go_authentication/options"
)

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
	return list, err
}

func allSl(w http.ResponseWriter, rows *sql.Rows) (list []*Slider, err error) {

	defer rows.Close()
	for rows.Next() {
		i := new(Slider)
		err = rows.Scan(
			&i.Id,
			&i.Collection_id,
			&i.Title,
			&i.Description,
			&i.Owner,
			&i.To_art,
			&i.To_sch,
			&i.To_prv_d,
			&i.To_prv_h,
			pq.Array(&i.Lt_t),
			pq.Array(&i.Lt_d),
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
	return list, err
}

func authorSl(w http.ResponseWriter, conn *sql.DB, id int, cls *authtoken.Claims) (i *Slider, err error) {

	i = &Slider{}
	i = new(Slider)
	owner := cls.User_id

	row := conn.QueryRow("SELECT * FROM slider WHERE id=$1 AND owner=$2", id, owner)

	err = row.Scan(
		&i.Id,
		&i.Collection_id,
		&i.Title,
		&i.Description,
		&i.Owner,
		&i.To_art,
		&i.To_sch,
		&i.To_prv_d,
		&i.To_prv_h,
		pq.Array(&i.Lt_t),
		pq.Array(&i.Lt_d),
		pq.Array(&i.Pfile),
		&i.Completed,
		&i.Created_at,
		&i.Updated_at,
	)

	if err == sql.ErrNoRows {
		fmt.Fprintf(w, " Error: sql.ErrNoRows..! : %+v\n", err)
		return
	} else if err != nil {
		fmt.Fprintf(w, " Error: sql..! : %+v\n", err)
		return
	}
	return i, err
}

func idSl(w http.ResponseWriter, conn *sql.DB, id int) (i *Slider, err error) {

	i = &Slider{}
	i = new(Slider)

	row := conn.QueryRow("SELECT * FROM slider WHERE id=$1", id)

	err = row.Scan(
		&i.Id,
		&i.Collection_id,
		&i.Title,
		&i.Description,
		&i.Owner,
		&i.To_art,
		&i.To_sch,
		&i.To_prv_d,
		&i.To_prv_h,
		pq.Array(&i.Lt_t),
		pq.Array(&i.Lt_d),
		pq.Array(&i.Pfile),
		&i.Completed,
		&i.Created_at,
		&i.Updated_at,
	)

	if err == sql.ErrNoRows {
		fmt.Fprintf(w, " Error: sql.ErrNoRows..! : %+v\n", err)
		return
	} else if err != nil {
		fmt.Fprintf(w, " Error: sql..! : %+v\n", err)
		return
	}
	return i, err
}

// ..
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
	return list, err
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
	return list, err
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
	return list, err
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
	return list, err
}


func delOk(src []bool) bool {
	for i := range src {
	    if src[i] == true {
	        return true
	    }
	}
    return false
}
func delList(src []string, del []string) []string {
	list := make([]string, 0, len(src))
	for _, v := range src {
		exist := true
		for _, w := range del {
			if v == w {
				exist = false
				break
			}
		}
		if exist {
			list = append(list, v)
		}
	}
	return list
}
func psDelStr(w http.ResponseWriter, r *http.Request, del []string, str string) ([]string, error) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error ParseForm..! : %+v\n", err)
	}
	on_off := make([]bool, len(r.Form[str]))
	for k, v := range r.Form[str] {
		flag, _ := options.ParseBool(v)
		on_off[k] = flag
	}
	fmt.Println(" on_off..", on_off)
	list := make([]string, 0, len(del))
	for k := range on_off {
		if on_off[k] == true {
			list = append(list, del[k])
		}
	}
	return list,err
}

func psDelImg(w http.ResponseWriter, r *http.Request) ([]string, error) {

	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		fmt.Fprintf(w, "Error MultipartForm..! : %+v\n", err)
	}

	on_off := make([]bool, len(r.Form["del"]))
	fmt.Println(" len on_off..", len(r.Form["del"]))
	for k, v := range r.Form["del"] {
		flag, _ := options.ParseBool(v)
		on_off[k] = flag
	}

	file := r.Form["path"]
	list := make([]string, 0, len(file))
	for k := range on_off {
		if on_off[k] == true {
			list = append(list, file[k])
		}
	}
	return list, err
}


func psFormI(
	w http.ResponseWriter, r *http.Request, cls *authtoken.Claims, sid string) ([]string, error) {

	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		fmt.Fprintf(w, " Error ParseMultipartForm..! : %+v\n", err)
	}
	files := r.MultipartForm.File["file"]
	list := make([]string, len(files))
	for k, i := range files {

		flname := i.Filename
		fpath := "./sfl/static/slider/" + cls.Email + "/" + sid + "/"
		fname := "./sfl/static/slider/" + cls.Email + "/" + sid + "/" + flname
		fle := "/static/slider/" + cls.Email + "/" + sid + "/" + flname

		mkdirerr := os.MkdirAll(fpath, 0750)
		if mkdirerr != nil {
			fmt.Fprintf(w, " Error MkdirAll..! : %+v\n", mkdirerr)
		}
		img, err := os.Create(fname)
		if err != nil {
			fmt.Fprintf(w, " Error Create..! : %+v\n", err)
		}
		defer img.Close()

		readerFile, _ := i.Open()
		if _, err := io.Copy(img, readerFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		list[k] = fle
	}
	return list, err
}

func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
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

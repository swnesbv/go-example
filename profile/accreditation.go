package profile

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"go_authentication/connect"
	"go_authentication/authtoken"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/signup.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", nil)
	}

	if r.Method == "POST" {

		username := r.FormValue("username")
		email 	 := r.FormValue("email")
		password := r.FormValue("password")
		hash, _  := hashPassword(password)

		conn := connect.ConnSql()
		sqlstr := "INSERT INTO users (username,email,password,created_at) VALUES ($1,$2,$3,$4)"

		_, err := conn.Exec(sqlstr, username,email,hash,time.Now())

		if err != nil {
			fmt.Fprintf(w, "err Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w,r, "/alluser", http.StatusFound)
	}
}

func UpName(w http.ResponseWriter, r *http.Request) {

	cls,err := authtoken.SqlToken(w,r)

	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	conn := connect.ConnSql()
	i,err := profilUser(w, conn,cls)
	if err != nil {
		return
	}

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/update_name.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", i)
	}

	if r.Method == "POST" {

		user := User{
			Username: r.FormValue("username"),
		}

		sqlstr := "UPDATE users SET username=$2, updated_at=$3 WHERE user_id=$1"

		_, err := conn.Exec(sqlstr, cls.User_id, user.Username, time.Now())

		if err != nil {
			fmt.Fprintf(w, "err Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w,r, "/alluser", http.StatusFound)
	}
}

func UpPass(w http.ResponseWriter, r *http.Request) {

	cls,err := authtoken.SqlToken(w,r)
	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	conn := connect.ConnSql()
	i,err := profilUser(w, conn,cls)
	if err != nil {
		return
	}

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/update_password.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", i)
	}

	if r.Method == "POST" {

		password := r.FormValue("password")
		hash, _  := hashPassword(password)

		sqlstr := "UPDATE users SET password=$2, updated_at=$3 WHERE user_id=$1"

		_, err := conn.Exec(sqlstr, cls.User_id,hash,time.Now())

		if err != nil {
			fmt.Fprintf(w, "err Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w,r, "/alluser", http.StatusFound)
	}
}

func EmailSend(w http.ResponseWriter, r *http.Request) {

	cls,err := authtoken.SqlToken(w,r)
	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	conn := connect.ConnSql()
	i,err := profilUser(w, conn,cls)
	if err != nil {
		return
	}

	if r.Method == "GET" {

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/send_email.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", i)
	}

	if r.Method == "POST" {

		email := r.FormValue("email")

		token,err := authtoken.BuildSend(w, email)
		if err != nil {
			return
		}
		emailerr := emailSend(r, token,email)
		if emailerr != nil {
			return
		}

		defer conn.Close()
		http.Redirect(w,r, "/alluser", http.StatusFound)
	}
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {

	cls,err := authtoken.SqlToken(w,r)
	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	veri := r.URL.Query().Get("veri")

	veri_email,err := authtoken.VerifySendToken(w, veri)
	if veri_email == nil {
		return
	}
	if err != nil {
		return
	}

	if r.Method == "GET" {

		conn := connect.ConnSql()
		sqlst := "UPDATE users SET email=$2, updated_at=$3 WHERE user_id=$1"

		_, err := conn.Exec(sqlst, cls.User_id,veri_email.Email,time.Now())

		if err != nil {
			fmt.Fprintf(w, "err Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w, r, "/alluser", http.StatusFound)
	}
}

func DelUs(w http.ResponseWriter, r *http.Request) {

	cls,err := authtoken.SqlToken(w,r)
	if cls == nil {
		return
	}
	if err != nil {
		return
	}

	if r.Method == "GET" {

		data := struct {
			Items string
		}{
			Items: cls.Email,
		}

		tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/delete.html", "./tpl/base.html"))

		tpl.ExecuteTemplate(w, "base", data)
	}

	if r.Method == "POST" {

		conn := connect.ConnSql()
		sqlstr := "DELETE FROM users WHERE user_id=$1"

		_, err := conn.Exec(sqlstr, cls.User_id)

		if err != nil {
			fmt.Fprintf(w, "err Exec..! : %+v\n", err)
			return
		}

		defer conn.Close()
		http.Redirect(w,r, "/alluser", http.StatusFound)
	}
}

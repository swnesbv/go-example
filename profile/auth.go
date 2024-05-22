package profile

import (
    // "database/sql"
    // "context"
    "time"
    "fmt"
    "os"
    "net/http"
    "html/template"
    
    "github.com/joho/godotenv"
    "github.com/golang-jwt/jwt/v5"

    "go_authentication/connect"
)


func Login(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/auth/login.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {

        email    := r.FormValue("email")
        password := r.FormValue("password")

        start := time.Now()
        start_1 := time.Now()

        conn := connect.ConnSql()
        saved_password,err := qPass(w, conn,email)
        if err != nil {
            fmt.Fprintf(w, " No user password : %+v\n", err)
            return
        }
        defer conn.Close()

        elapsed1 := time.Since(start_1)
        fmt.Printf(" 1 time.. :  %s \n", elapsed1)

        start_2 := time.Now()
        match := checkPass(password, saved_password)
        elapsed2 := time.Since(start_2)
        fmt.Printf(" 2 time.. :  %s \n", elapsed2)

        if match == false {
            fmt.Fprintf(w, " Match matching passwords..! : %+v\n", match)
            return 
        }

        start_3 := time.Now()

        if err := godotenv.Load(); err != nil {
            fmt.Fprintf(w, " No .env file found : %+v\n", err)
            return
        }

        user_id,err := userId(w, conn,email)
        if err != nil {
            fmt.Fprintf(w, " No user email : %+v\n", err)
            return
        }

        token := jwt.New(jwt.SigningMethodHS256)
        cls := token.Claims.(jwt.MapClaims)

        cls["authorized"] = true
        cls["user_id"] = user_id
        cls["email"] = email
        cls["exp"] = time.Now().Add(time.Minute * 60).Unix()

        tokenstr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, " err SignedString..! : %+v\n", err)
            return
        }

        cookie := http.Cookie{
            Name:     "Visitor",
            Value:    tokenstr,
            Path:     "/",
            MaxAge:   3600,
            HttpOnly: true,
            Secure:   false,
            SameSite: http.SameSiteLaxMode,
        }
        http.SetCookie(w, &cookie)

        fmt.Fprintf(w, " OK : token..!")
        fmt.Fprintf(w, " OK : login successful..!")


        elapsed3 := time.Since(start_3)
        fmt.Printf(" 3 time.. :  %s \n", elapsed3)

        elapsed := time.Since(start)
        fmt.Printf(" all time.. :  %s \n", elapsed)
        return
    }
}


/*func Login(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/auth/login.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {

        email := r.FormValue("email")
        password := r.FormValue("password")

        var saved_password string

        start_1 := time.Now()
        conn := connect.ConnSql()
        err := conn.QueryRow(`SELECT password FROM users WHERE email=$1`, email).Scan(&saved_password)

        elapsed1 := time.Since(start_1)
        fmt.Printf(" 1 time.. :  %s \n", elapsed1)

        switch {
        case err == sql.ErrNoRows:
            fmt.Fprintf(w, "Error: login failed email..! : %+v\n", email)
            fmt.Fprintf(w, "err: login failed email err..! : %+v\n", err)
            return
        case err != nil:
            fmt.Fprintf(w, "Error: QueryRow..! : %+v\n", err)
            break
        default:
            fmt.Println("email is : ", email)
        }


        start_2 := time.Now()
        match := checkPasswordHash(password, saved_password)
        if match == false {
            fmt.Fprintf(w, "Match matching passwords..! : %+v\n", match)
            return 
        }
        elapsed2 := time.Since(start_2)
        fmt.Printf(" 2 time.. :  %s \n", elapsed2)

        if match {

            start_3 := time.Now()

            if err := godotenv.Load(); err != nil {
                fmt.Fprintf(w, "No .env file found : %+v\n", err)
                return
            }

            // var ctx context.Context
            var user_id int

            err := conn.QueryRow(`SELECT user_id FROM users WHERE email=$1`, email).Scan(&user_id)
            // err := conn.QueryRowContext(ctx, "SELECT user_id FROM users WHERE email=$1", email).Scan(&user_id)

            switch {
            case err == sql.ErrNoRows:
                fmt.Fprintf(w, "no user with user_id : %+v\n", user_id)
                fmt.Fprintf(w, "no user with user_id err : %+v\n", err)
                return
            case err != nil:
                break
            default:
                fmt.Println("user_id is : ", user_id)
            }

            token := jwt.New(jwt.SigningMethodHS256)
            cls := token.Claims.(jwt.MapClaims)

            cls["authorized"] = true
            cls["user_id"] = user_id
            cls["email"] = email
            cls["exp"] = time.Now().Add(time.Minute * 60).Unix()

            tokenstr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

            if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, " Error: SignedString..! : %+v\n", err)
                return
            }

            cookie := http.Cookie{
                Name:     "Visitor",
                Value:    tokenstr,
                Path:     "/",
                MaxAge:   3600,
                HttpOnly: true,
                Secure:   false,
                SameSite: http.SameSiteLaxMode,
            }
            http.SetCookie(w, &cookie)

            fmt.Fprintf(w, "OK : token..!")
            fmt.Fprintf(w, " OK : login successful..!")

            elapsed3 := time.Since(start_3)
            fmt.Printf("3 time.. :  %s \n", elapsed3)

            defer conn.Close()
            return
        }

        defer conn.Close()
        fmt.Println("Error: conn.Close..!")
        fmt.Fprintf(w, "Error: login failed..!")

        return
    }
}*/


func AuthToken(w http.ResponseWriter, r *http.Request) {

    tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/auth/auth.html", "./tpl/base.html" ))

    c,err := r.Cookie("Visitor")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, " Error: http.ErrNoCookie..! : %+v\n", err)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, " Error: Cookie..! : %+v\n", err)
        return
    }

    tkstr := c.Value
    cls := &Claims{}

    token, err := jwt.ParseWithClaims(tkstr, cls, func(token *jwt.Token) (any, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, " Error: ErrSignatureInvalid..! : %+v\n", err)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, " Error: ParseWithClaims()..! : %+v\n", err)
        return
    }
    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    tpl.ExecuteTemplate(w, "base", cls)
}
package profile

import (
    "database/sql"
    "os"
    "fmt"
    "net/http"
    // "time"
    "net/smtp"

    "golang.org/x/crypto/bcrypt"

    "go_authentication/authtoken"
)

func hashPassword(password string) (string, error) {
    b,err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(b),err
}

func checkPass(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func qPass(w http.ResponseWriter, conn *sql.DB, email string) (saved_password string, err error) {

    row := conn.QueryRow("SELECT password FROM users WHERE email=$1", email).Scan(&saved_password)

    switch {
    case row == sql.ErrNoRows:
        fmt.Fprintf(w, " Error: login email..! : %+v\n", email)
        fmt.Fprintf(w, " err: login email..! : %+v\n", row)
        return
    case row != nil:
        fmt.Fprintf(w, " Error: QueryRow..! : %+v\n", row)
        break
    default:
        fmt.Println(" email..", email)
    }
    return saved_password,err
}

func userId(w http.ResponseWriter, conn *sql.DB, email string) (user_id int, err error) {
    
    row := conn.QueryRow("SELECT user_id FROM users WHERE email=$1", email).Scan(&user_id)

    switch {
    case row == sql.ErrNoRows:
        fmt.Fprintf(w, "no user with user_id : %+v\n", user_id)
        fmt.Fprintf(w, "no user with user_id err : %+v\n", row)
        return
    case row != nil:
        break
    default:
        fmt.Println("user_id is : ", user_id)
    }
    return user_id,err
}


func IdUser(w http.ResponseWriter, conn *sql.DB, id int) (i AllUser, err error) {

    row := conn.QueryRow("SELECT user_id, username, email FROM users WHERE user_id=$1", id)

    err = row.Scan(
        &i.User_id,
        &i.Username,
        &i.Email,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "sql.ErrNoRows err..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "sql err..! : %+v\n", err)
        return
    }
    return i,err
}


func profilUser(w http.ResponseWriter, conn *sql.DB, cls *authtoken.Claims) (i *AllUser, err error) {

    i = &AllUser{}
    i = new(AllUser)
    user_id := cls.User_id

    row := conn.QueryRow("SELECT user_id, username, email, password, created_at, updated_at FROM users WHERE user_id=$1", user_id)

    err = row.Scan(
        &i.User_id,
        &i.Username,
        &i.Email,
        &i.Password,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "sql.ErrNoRows err..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "sql err..! : %+v\n", err)
        http.Error(w, http.StatusText(500), 500)
        return
    }
    return i,err
}


func emailSend(r *http.Request, token, email string) (err error) {

  from := os.Getenv("SMTP_FROM")
  password := os.Getenv("SMTP_PASSWORD")

  to := []string{ email }

  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  message := []byte("email message\r\n" + r.Host + "/verification?veri=" + token)

  auth := smtp.PlainAuth("", from, password, smtpHost)

  err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
  if err != nil {
    fmt.Println("err SendMail..!", err)
    return
  }

  fmt.Println("Email Sent Successfully!")

  return nil
}
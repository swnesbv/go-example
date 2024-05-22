package authtoken

import (
    "time"
    "fmt"
    "os"
    "errors"
    "net/http"
    "encoding/json"
    "encoding/base64"
    
    "github.com/golang-jwt/jwt/v5"
)


func TokenSearch(w http.ResponseWriter, r *http.Request) (cls *ListData, err error) {

    c,err := r.Cookie("Period")

    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Redirect(w,r, "/creat-period", http.StatusFound)
        case c == nil:
            fmt.Fprintf(w, "token nil..! : %+v\n", err)
        case err != nil:
            fmt.Fprintf(w, "err: ..! : %+v\n", err)
        }
        return
    }

    data,err := base64.StdEncoding.DecodeString(c.Value)
    if err != nil {
        fmt.Println("error:", err)
    }

    jerr := json.Unmarshal(data, &cls)
    if jerr != nil {
        fmt.Fprintf(w, "err: Unmarshal..! : %+v\n", jerr)
    }

    return cls,err
}


func RedirectToken(w http.ResponseWriter, r *http.Request) (cls *Claims,err error) {

    c,err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Redirect(w,r, "/login", http.StatusFound)
        }
        return
    }

    cls = &Claims{}

    token,err := jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        http.Redirect(w,r, "/login", http.StatusFound)
        return
    }
    if !token.Valid {
        http.Redirect(w,r, "/login", http.StatusFound)
        return
    }

    return cls,err
}


func OnToken(w http.ResponseWriter, r *http.Request) (cls *Claims, err error) {

    c,err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            fmt.Println("err r.Cookie().. ", err)
            http.Redirect(w, r, "/login", http.StatusUnauthorized)
        }
        return
    }

    cls = &Claims{}

    token,err := jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    switch {
    case token == nil:
        fmt.Fprintf(w, "token nil..! : %+v\n", err)
    case err != nil:
        fmt.Fprintf(w, "OnToken err: ParseWithClaims..! : %+v\n", err)
    }

    return cls,err
}


func SqlToken(w http.ResponseWriter, r *http.Request) (cls *Claims, err error) {

    c,err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            fmt.Println("err r.Cookie().. ", err)
            http.Redirect(w, r, "/login", http.StatusUnauthorized)
        }
        return
    }

    cls = &Claims{}

    token, err := jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            fmt.Fprintf(w, " Error: ErrSignatureInvalid..! : %+v\n", err)
            return
        }
        fmt.Fprintf(w, "SqlToken err: ParseWithClaims..! : %+v\n", err)
        return
    }
    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    return cls,err
}


func ListToken(w http.ResponseWriter, r *http.Request) (*Claims,error) {

    c,err := r.Cookie("Visitor")

    cls := &Claims{}

    switch {
    case errors.Is(err, http.ErrNoCookie):
        break
    default:
        _,err = jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
    }
    return cls,err
}


func WhoisWho(w http.ResponseWriter, r *http.Request) *Claims {

    c,err := r.Cookie("Visitor")

    cls := &Claims{}

    switch {
    case errors.Is(err, http.ErrNoCookie):
        break
    default:
        _,err = jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
    }
    return cls
}



func BuildSend(w http.ResponseWriter, email string) (string, error) {

    token := jwt.New(jwt.SigningMethodHS256)
    cls := token.Claims.(jwt.MapClaims)

    cls["email"] = email
    cls["exp"] = time.Now().Add(time.Minute * 60).Unix()

    tokenstr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, " Error: SignedString()..! : %+v\n", err)
    }
    return tokenstr,err
}


func VerifySendToken(w http.ResponseWriter, veri string) (*Claims,error) {

    cls := &Claims{}

    token, err := jwt.ParseWithClaims(veri, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if token == nil {
        fmt.Fprintf(w, "token nil..! : %+v\n", err)
    }
    if err != nil {
        fmt.Fprintf(w, "Verify err: jwt.ParseWithClaims()..! : %+v\n", err)
    }

    return cls,err
}
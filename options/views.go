package options

import (
    "fmt"
    "os"
    "strconv"
    "net/http"
    "path/filepath"
)

func OkBool(a, b int) bool {
    if a == b {
        return true
    }
    return false
}
func InSlice(i int, list []int) bool {
    for _, s := range list {
        if s == i {
            return true
        }
    }
    return false
}


func DelFolder(path string) (err error) {
    
    contents,err := filepath.Glob(path)
    if err != nil {
        return
    }
    for _, item := range contents {
        err = os.RemoveAll(item)
        if err != nil {
            return
        }
    }
    return
}


func ParseBool(str string) (bool, error) {

    switch str {
    case "1", "t", "T", "true", "TRUE", "True", "on", "yes", "ok":
        return true, nil
    case "", "0", "f", "F", "false", "FALSE", "False", "off", "no":
        return false, nil
    }
    return false, &strconv.NumError{Func: "ParseBool", Num: str, Err: strconv.ErrSyntax}
}


func IdUrl(w http.ResponseWriter, r *http.Request) (id int, err error) {

    id,err = strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        fmt.Println("id_err..", err)
        http.NotFound(w,r)
        return
    }
    return id,err
}


func PsFormString (
    w http.ResponseWriter, r *http.Request, str string) (list []string, err error) {

    pserr := r.ParseForm()
    if pserr != nil {
        fmt.Fprintf(w, "Error ParseForm..! : %+v\n", pserr)
        return
    }
    list = r.Form[str]
    ps := make([]string, len(list))

    for k, v := range list {
        ps[k] = v
    }
    return list,err
}

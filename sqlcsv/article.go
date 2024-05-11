package sqlcsv

import (
    "os"
    "fmt"
    "time"
    // "bytes"
    "net/http"
    "encoding/csv"
    "html/template"

    "go_authentication/connect"
    "go_authentication/authtoken"
)


/*func ImpArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        
        rows,err := conn.Query("SELECT * FROM article WHERE owner=$1", owner)

        if err != nil {
            switch {
                case true:
                fmt.Fprintf(w, "Error: rows..! : %+v\n", err)
                break
            }
            return
        }
        defer rows.Close()

        w.Header().Set("Content-type", "text/csv")
        w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")

        Write(w, rows)
        defer conn.Close()
    }
}*/


/*func ImpArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        
        rows,err := conn.Query("SELECT * FROM article WHERE owner=$1", owner)

        if err != nil {
            switch {
                case true:
                fmt.Fprintf(w, "Error: rows..! : %+v\n", err)
                break
            }
            return
        }
        defer rows.Close()
        
        col,err := rows.Columns()
        if err != nil {
            fmt.Println("Error getting column names:", err)
            return
        }

        file := bytes.NewBuffer(nil)

        wri := csv.NewWriter(file)
        header := []string{"id","title","description","img","owner","completed","created_at","updated_at"}
        wri.Write(header)
        defer wri.Flush()

        key := make([]interface{}, len(col))
        values := make([]interface{}, len(col))

        for i := range values {
            key[i] = &values[i]
        }

        for rows.Next() {
            err := rows.Scan(key...)
            if err != nil {
                fmt.Println("Error scanning row:", err)
                return
            }

            var row string
            for _, value := range values {
                if value != nil {
                    row += fmt.Sprintf("%v,", value)
                } else {
                    row += ","
                }
            }
            file.WriteString(fmt.Sprintf("%s\n", row[:len(row)-1]))
        }

        w.Header().Set("Content-type", "text/csv")
        w.Header().Set("Content-Disposition", "attachment; filename=\"report2.csv\"")
        w.Write(file.Bytes())
        
        // w.Header().Set("Location", "/allarticle")
        // w.WriteHeader(http.StatusFound)
        // fmt.Println("CSV successfully buf..!", file.String())
        defer conn.Close()
    }
}*/


func ImpArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        
        rows,err := conn.Query("SELECT * FROM article WHERE owner=$1", owner)

        if err != nil {
            switch {
                case true:
                fmt.Fprintf(w, "Error: rows..! : %+v\n", err)
                break
            }
            return
        }
        defer rows.Close()

        mkdirerr := os.MkdirAll("./sfl/static/csv/" + cls.Email, 0750)
        if mkdirerr != nil {
            fmt.Println("Error os.MkdirAll():", mkdirerr)
        }

        file,err := os.Create("./sfl/static/csv/" + cls.Email + "/article.csv")
        if err != nil {
            fmt.Println("Error os.Create():", err)
        }
        defer file.Close()

        wr := csv.NewWriter(file)
        header := []string{"id","title","description","img","owner","completed","created_at","updated_at"}

        wr.Write(header)
        defer wr.Flush()

        col,err := rows.Columns()
        if err != nil {
            fmt.Println("Error getting column names:", err)
            return
        }

        key := make([]interface{}, len(col))
        values := make([]interface{}, len(col))

        for i := range values {
            key[i] = &values[i]
        }

        for rows.Next() {
            err := rows.Scan(key...)
            if err != nil {
                fmt.Println("Error scanning row:", err)
                return
            }

            var row string
            for _, value := range values {
                if value != nil {
                    row += fmt.Sprintf("%v,", value)
                } else {
                    row += ","
                }
            }
            file.WriteString(fmt.Sprintf("%s\n", row[:len(row)-1]))
        }

        fmt.Println("CSV successfully..!", file)

        defer conn.Close()

        http.Redirect(w,r, "/static/csv/" + cls.Email + "/article.csv", http.StatusFound)

    }
}


func ExpArt(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/csv/export.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", nil)

    }


    if r.Method == "POST" {

        conn := connect.ConnSql()
        owner := cls.User_id

        file, handler, err := r.FormFile("file")
        if err != nil{
            fmt.Println("Error Data retrieving")
            fmt.Println(err)
            return
        }

        fileName := handler.Filename
        
        fmt.Printf("Uploaded File : %+v\n", fileName)
        fmt.Printf("File Size : %+v\n" , handler.Size)
        fmt.Printf("MIME Header : %+v\n" , handler.Header)

        reader := csv.NewReader(file)
        rows,err := reader.ReadAll()
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "err ReadAll()..! : %+v\n", err)
            return
        }
        fmt.Printf("rows ReadAll..! : %+v\n", rows)

        for _, row := range rows {

        i2 := row[1]
        i3 := row[2]

        sqlst := "INSERT INTO article (title,description,img,owner,completed,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"

        _, err := conn.Exec(sqlst, i2,i3,nil,owner,false,time.Now(),nil)

        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "err Exec..! : %+v\n", err)
            return
        }
        fmt.Println("OK..! Exec..", row)
        }

        defer conn.Close()

        http.Redirect(w,r, "/author-id-article", http.StatusFound)
    }
}


/*func ExpArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }
        owner := cls.User_id

        file,err := os.Open("report.csv")
        if err != nil {
            fmt.Println("os.Open..", err)
        }
        defer file.Close()

        reader := csv.NewReader(file)
        rows,err := reader.ReadAll()
        if err != nil {
            fmt.Println("reader.ReadAll..", err)
        }

        for _, row := range rows {

            i1 := row[0]
            i2 := row[1]
            i3 := row[2]
            i4 := row[3]

            _, err := conn.Exec("INSERT INTO article (id,title,description,img,owner,completed,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", i1,i2,i3,i4,owner,false,time.Now(),nil)
            if err != nil {
                fmt.Println("err Exec..", err)
            } else {
                fmt.Println("OK..! Exec..", row)
            }
        }

        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/export.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", nil)

    }
}*/

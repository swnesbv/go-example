package connect

import (
    "database/sql"
    "time"
    "fmt"
    "os"
    "runtime"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


func ConnSql() (*sql.DB) {

    start := time.Now()

    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found..", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    ch := make(chan *sql.DB)

    go func() {
        conn,err := sql.Open("postgres", connstr)
        if err != nil {
            fmt.Println("coon err: sql Open..", err)
        }

        ping_err := conn.Ping()
        if ping_err != nil {
            fmt.Println("coon err: Ping..", ping_err)
        }
        if ping_err == nil {
            fmt.Println("ConnSql OK..!")
        }

        ch <- conn
    }()

    
    fmt.Println(" ConnSql goroutine..", runtime.NumGoroutine())

    elapsed := time.Since(start)
    fmt.Printf(" sql conn time.. :  %s \n", elapsed)

    return <- ch
}


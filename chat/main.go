package chat

import (
	"database/sql"
	"time"
	"fmt"
    "os"
    "runtime"
	
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


var db *sql.DB

func connchat() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found chat", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db,err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() chat", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() chat", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected chat")
    }

}


func init() {

    start := time.Now()

	go connchat()
	go userCh()
	go groupCh()

    elapsed := time.Since(start)
    fmt.Printf(" sql chat time.. :  %s \n", elapsed)
    
    fmt.Println(" chat goroutine..", runtime.NumGoroutine())

    time.Sleep(100 * time.Millisecond)
}

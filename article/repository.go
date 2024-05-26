package article

import (
	"database/sql"
	"fmt"
	"net/http"
	"runtime"
)

func qArtCount(w http.ResponseWriter, conn *sql.DB) (int, error) {

	var count int
	err := conn.QueryRow("SELECT COUNT(*) FROM article").Scan(&count)

	if err != nil {
		fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
	}
	return count, err
}

func qArt(
	w http.ResponseWriter, conn *sql.DB, limit int, offset int) (rows *sql.Rows, err error) {

	rows, err = conn.Query("SELECT id, title, description, img, owner, completed, created_at, updated_at FROM article ORDER BY id LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		switch {
		case true:
			fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
			break
		}
		return
	}
	return rows, err
}

func qArtBool(w http.ResponseWriter, conn *sql.DB) (rows *sql.Rows, err error) {

	rows, err = conn.Query("SELECT id, title, description, img, owner, completed, created_at, updated_at FROM article WHERE completed=$1", true)

	if err != nil {
		switch {
		case true:
			fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
			break
		}
		return
	}
	fmt.Println(" qArt goroutine..", runtime.NumGoroutine())
	return rows, err
}

func qUsArt(w http.ResponseWriter, conn *sql.DB, owner int) (rows *sql.Rows, err error) {

	rows, err = conn.Query("SELECT * FROM article WHERE owner=$1", owner)

	if err != nil {
		switch {
		case true:
			fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
			break
		}
		return
	}
	return rows, err
}

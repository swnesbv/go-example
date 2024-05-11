package pagination

import (
	"fmt"
	"net/http"
	"strconv"
)


func InvalidArgument(w http.ResponseWriter) {
	http.Error(w, "Invalid argument(s)", http.StatusBadRequest)
}


func PageNumber(r *http.Request) (p int, err error) {
	queryValues := r.URL.Query()
	pageNumberStr := queryValues.Get("page")
	if pageNumberStr == "" {
		p = 1
		return
	}
	p, err = strconv.Atoi(pageNumberStr)
	if err != nil {
		fmt.Println("Invalid argument.. Page: ", pageNumberStr)
		fmt.Println("Invalid argument.. err..", err)
		return
	}
	return
}
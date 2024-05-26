package article

import (
	"time"
)

type Article struct {
	Id          int
	Title       string
	Description string
	Img         *string
	Owner       int
	Completed   bool
	Created_at  time.Time
	Updated_at  *time.Time
}

type Slider struct {
	Id            int
	Collection_id string
	Title         string
	Description   string
	Owner         int
	To_art     	  *int
	To_sch    	  *int
	To_prv_d 	  *int
	To_prv_h 	  *int
	Lt_t   	  	  []string
	Lt_d   	  	  []string
	Pfile   	  []string
	Completed     bool
	Created_at    time.Time
	Updated_at    *time.Time
}

type LSl struct {
	Id string
    T  string
    D  string
    F  string
}
type ListPageData struct {
	Users []Article
}

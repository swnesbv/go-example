package slider

import (
	"time"
)

type Collection struct {
	Id            int
	Collection_id string
	Owner         int
	Pfile     	  []string
	Completed     bool
	Created_at    time.Time
	Updated_at    *time.Time
}
type Slider struct {
	Id            int
	Collection_id string
	Title         string
	Description   string
	Owner         int
	To_art     	  int
	To_sch    	  int
	To_prv_d 	  int
	To_prv_h 	  int
	Pfile   	  []string
	Completed     bool
	Created_at    time.Time
	Updated_at    *time.Time
}


type Article struct {
	Id int
}
type Schedule struct {
	Id int
}
type Provision_d struct {
	Id int
}
type Provision_h struct {
	Id int
}
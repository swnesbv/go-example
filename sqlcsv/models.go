package sqlcsv

import (
	"time"
)

type Schedule struct {
	Id          int
	Title       string
	Description string
	Owner       int
	St_hour     time.Time
	En_hour     time.Time
	Hours       []string
	Occupied    []string
	Completed   bool
	Created_at  time.Time
	Updated_at  *time.Time
}
type Recording struct {
	Id          int
	Owner       int
	To_schedule int
	Record      *time.Time
	Completed   bool
	Created_at  time.Time
	Updated_at  *time.Time
}

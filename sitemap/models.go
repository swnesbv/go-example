package sitemap

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

package article

import (
	//"database/sql"
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
type CreatArticle struct {
	Title       string
	Description string
	Owner       int
	Created_at  time.Time
}

type ListPageData struct {
	Users []Article
}

type UpdateArticle struct {
	Title       string
	Description string
}

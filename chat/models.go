package chat

import (
    "time"
)


type ChUser struct {
    User_id  int
    Username string
    Email    string
}

type Subscription struct {
    Id int
    Owner int
    To_user int
    To_group int
    Completed bool
}


type ToUser struct {
    User_id  int
    Username string
    Email    string
    Img      *string
    Created_at time.Time
    Updated_at *time.Time
}

type MsgUser struct {
    Id int
    Coming string
    Img *string
    Owner int
    To_user int
    Completed bool
    Created_at time.Time
    Updated_at *time.Time
}

type Group struct {
    Id int
    Title string
    Description string
    Img *string
    Owner int
    Completed bool
    Created_at time.Time
    Updated_at *time.Time
}
type MsgGroup struct {
    Id int
    Coming string
    Img *string
    Owner int
    To_group int
    Completed bool
    Created_at time.Time
    Updated_at *time.Time
}
type CreatGroup struct {
    Title string
    Description string
    Owner int
    Created_at time.Time
}
type UpdateGroup struct {
    Title string
    Description string
}

type Message struct {
    Username string     `json:"username"`
    Message  string     `json:"message"`
    Creation string     `json:"time"`
}

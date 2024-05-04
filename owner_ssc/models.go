package owner_ssc

import (
    "time"
)


type Subscription struct {
    Id int
    Title string
    Description string
    Owner int
    To_user *int
    To_group *int
    Completed bool
    Created_at time.Time
    Updated_at *time.Time
}

type CreatSubscription struct {
    Title string
    Description string
    Owner int
    To_user int
    To_group int
    Created_at time.Time
}

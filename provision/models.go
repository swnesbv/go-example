package provision

import (
    "time"
)

type Provision_d struct {
    Id          int
    Title       string
    Description string
    Owner       int
    St_date     *time.Time
    En_date     *time.Time
    S_dates     []string
    E_dates     []string
    Dates       []string
    Completed   bool
    Created_at  time.Time 
    Updated_at  *time.Time 
}
type Provision_h struct {
    Id          int
    Title       string
    Description string
    Owner       int
    St_hour     *time.Time
    En_hour     *time.Time
    S_hours     []string
    E_hours     []string
    Hours       []string
    Completed   bool
    Created_at  time.Time 
    Updated_at  *time.Time 
}
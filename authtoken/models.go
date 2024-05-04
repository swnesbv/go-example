package authtoken

import (
    "time"
    
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    User_id int `json:"user_id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}
type ListData struct {
    D_start time.Time `json:"d_start"`
    D_end time.Time `json:"d_end"`
    Start time.Time `json:"start"`
    End time.Time `json:"end"`
}
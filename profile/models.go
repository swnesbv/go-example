package profile

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AllUser struct {
	User_id    int
	Username   string
	Email      string
	Password   string
	Img        *string
	Created_at time.Time
	Updated_at *time.Time
}
type User struct {
	Username string
	Email    string
	Password string
}
type UserList struct {
	User_id  int
	Username string
	Email    string
}

type Claims struct {
	User_id int    `json:"user_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

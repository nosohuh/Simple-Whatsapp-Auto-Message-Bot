package usermodels

import (
	"time"
)

type Users struct {
	ID         string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role       string    `json:"role"`
	Balance    float64   `json:"balance"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Level      int       `json:"level"`
	TOTPKey    string    `json:"top_otp"`
	DealerName string    `json:"dealer"`
	WebSite    string    `json:"web_site"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

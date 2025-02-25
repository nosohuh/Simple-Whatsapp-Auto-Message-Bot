package usermodels

import (
	"time"
)

type Servis struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Numara    string    `json:"numara"`
	Mesaj	 string    `json:"mesaj"`
	Durum     bool      `json:"durum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

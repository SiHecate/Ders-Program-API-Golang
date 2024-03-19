package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	UserID          int       `json:"user_id"`
	Baslik          string    `json:"başlık"`
	Plan            string    `json:"plan"`
	Gun             time.Time `json:"gün"`
	BaslangicZamani time.Time `json:"başlangıç_zamanı"`
	BitisZamani     time.Time `json:"bitiş_zamanı"`
	Durum           string    `json:"durum"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Başlık           string    `json:"başlık"`
	Plan             string    `json:"plan"`
	Gün              time.Time `json:"gün"`
	Başlangıç_zamanı time.Time `json:"başlangıç_zamanı"`
	Bitiş_zamanı     time.Time `json:"bitiş_zamanı"`
	State            string    `json:"durum"`
}

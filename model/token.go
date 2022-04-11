package model

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expiry       time.Time `json:"expiry"`
	TokenType    string    `json:"token_type"`
}

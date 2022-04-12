package model

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserId       string    `json:"user_id"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expiry       time.Time `json:"expiry"`
	TokenType    string    `json:"token_type"`
}

func CreateToken(db *gorm.DB, token *Token) (*Token, error) {
	err := db.Create(token).Error
	return token, err
}

func GetTokenByUserId(db *gorm.DB, userId string) (token Token, err error) {
	err = db.Where("user_id = ?", userId).First(&token).Error
	return token, err
}

func (t *Token) Save(db *gorm.DB) error {
	err := db.Save(t).Error
	return err
}

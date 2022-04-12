package model

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserId       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
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

func UpdateTokenByUserId(db *gorm.DB, userId string, accessToken, refreshToken, tokenType string, expiry time.Time) error {
	m := make(map[string]interface{})
	m["access_token"] = accessToken
	m["refresh_token"] = refreshToken
	m["token_type"] = tokenType
	m["expiry"] = expiry
	err := db.Model(&Token{}).Where("user_id = ?", userId).Updates(m).Error
	return err
}

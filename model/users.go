package model

import "time"

type (
	UserCredentials struct {
		Id           int       `json:"id" gorm:"primaryKey;autoIncrement"`
		Fullname     string    `json:"fullname,omitempty"`
		Password     string    `json:"password,omitempty"`
		DivisionCode string    `json:"divisionCode,omitempty"`
		IsReset      bool      `json:"isReset,omitempty" gorm:"default:0"`
		CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP,omitempty"`
		UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP,omitempty"`
	}

	ViewUsers struct {
		Id             int    `json:"id"`
		Fullname       string `json:"fullname"`
		Password       string `json:"password"`
		Name           string `json:"Name"`
		DivisionCode   string `json:"divisionCode"`
		IsReset        bool   `json:"isReset"`
		DateRegistered string `json:"dateRegistered"`
	}
)

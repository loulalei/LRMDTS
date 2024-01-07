package model

import "time"

var UserActivity = &UserActivities{}

type (
	UserCredentials struct {
		Id           int       `json:"id" gorm:"primaryKey;autoIncrement"`
		Fullname     string    `json:"fullname"`
		Password     string    `json:"password"`
		DivisionCode string    `json:"divisionCode"`
		IsReset      bool      `json:"isReset" gorm:"default:0"`
		CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

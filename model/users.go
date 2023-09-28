package model

type (
	UserCredentials struct {
		Id           int    `json:"id" gorm:"primary key"`
		Fullname     string `json:"fullname"`
		Password     string `json:"password"`
		DivisionCode string `json:"division_code"`
	}
)

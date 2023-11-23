package model

type (
	CommitteeList struct {
		ListId      int    `json:"ListId,omitempty" gorm:"primary key"`
		ItemNumber  string `json:"itemNumber"`
		CommitteeId int    `json:"committeeId"`
		UserId      int    `json:"userId"`
	}

	ViewCommittees struct {
		ListId     int    `json:"listId"`
		ItemNumber string `json:"itemNumber"`
		Name       string `json:"name"`
		Fullname   string `json:"fullname"`
	}
)

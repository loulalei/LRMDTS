package model

type (
	RequestCommittee struct {
		ItemNumber  string `json:"itemNumber"`
		CommitteeId int    `json:"committeeId"`
		UserId      int    `json:"userId"`
	}
	CommitteeList struct {
		ListId      int    `json:"ListId,omitempty" gorm:"primary key"`
		ItemNumber  string `json:"itemNumber"`
		CommitteeId int    `json:"committeeId"`
		UserId      int    `json:"userId"`
	}

	ViewCommittees struct {
		Id         int    `json:"committeeId,omitempty"`
		ListId     int    `json:"listId,omitempty"`
		ItemNumber string `json:"itemNumber,omitempty"`
		Name       string `json:"name,omitempty"`
		Fullname   string `json:"fullname,omitempty"`
	}
)

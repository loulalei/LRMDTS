package model

type (
	EmployeePerformances struct {
		Fullname         string `json:"fullname"`
		Date             string `json:"date"`
		RecordsCaptured  int    `json:"recordsCaptured"`
		RecordsRetrieved int    `json:"recordsRetrieved"`
	}
)

package model

type (
	EmployeePerformances struct {
		Date             string `json:"date"`
		RecordsCaptured  int    `json:"recordsCaptured"`
		RecordsRetrieved int    `json:"recordsRetrieved"`
	}
)

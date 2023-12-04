package model

import "time"

type (
	Receivings struct {
		ReceivingId      int `gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber   string
		ReceivedDate     string
		ReceivedTime     string
		Receiver         string
		Summary          string
		ReceivingTag     string
		ReceivingRemarks string
		ReceivedFile     string
		CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Agendas struct {
		AgendaId       int `gorm:"primaryKey;autoIncrement:true"`
		ItemNumber     string
		IsUrgent       bool
		DateCalendared string
		DateReported   string
		Source         string
		SourceResult   string
		AgendaTag      string
		AgendaRemarks  string
		CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Proponents struct {
		ProponentId int    `json:"proponentId" gorm:"primaryKey;autoIncrement:true"`
		Name        string `json:"name"`
		Term        string `json:"term"`
	}

	Committees struct {
		CommitteeId int    `json:"committeeId" gorm:"primaryKey;autoIncrement:true"`
		Name        string `json:"name"`
	}

	Departments struct {
		DepartmentId int    `json:"departmentId" gorm:"primaryKey;autoIncrement"`
		Name         string `json:"name"`
	}

	Divisions struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
)

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

	Routings struct {
		DocId       int       `json:"docId" gorm:"primaryKey;autoIncrement:true"`
		ReceivingId int       `json:"receivingId"`
		AgendaId    int       `json:"agendaId"`
		ApprovalId  int       `json:"approvalId"`
		ReleasingId int       `json:"releasingId"`
		FilingId    int       `json:"filingId"`
		ItemNumber  string    `json:"itemNumber"`
		DocumentTag string    `json:"documentTag"`
		Remarks     string    `json:"remarks"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Trackings struct {
		TrackingId    int       `json:"trackingId" gorm:"primaryKey;autoIncrement:true"`
		ItemNumber    string    `json:"itemNumber" gorm:"default:None"`
		LawNumber     string    `json:"lawNumber" gorm:"default:None"`
		Summary       string    `json:"summary" gorm:"default:None"`
		ReceivedDate  string    `json:"receivedDate" gorm:"default:None"`
		Calendared    string    `json:"calendared" gorm:"default:None"`
		EnactedDate   string    `json:"enactedDate" gorm:"default:None"`
		MayorDate     string    `json:"forwardedMayorDate" gorm:"default:None"`
		StaCruzDate   string    `json:"ForwardedStaCruzDate" gorm:"default:None"`
		ReleasedDate  string    `json:"releasedDate" gorm:"default:None"`
		PublishedDate string    `json:"publishedDate" gorm:"default:None"`
		FiledDate     string    `json:"filedDate" gorm:"default:None"`
		CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

	FilePaths struct {
		FileCode    int       `json:"fileCode"`
		FileName    string    `json:"filename"`
		Description string    `json:"description"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

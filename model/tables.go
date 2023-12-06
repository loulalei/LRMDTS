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
		Encoder          string
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
		Encoder        string
		CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Approves struct {
		ApproveId   int `gorm:"primaryKey;autoIncrement:true"`
		LawType     string
		LawNumber   string
		Series      string
		EnactedDate string
		MotionedBy  string
		Author      string
		ResOrdFile  string
		Encoder     string
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Routings struct {
		DocId       int       `json:"docId" gorm:"primaryKey;autoIncrement:true"`
		ReceivingId int       `json:"receivingId"`
		AgendaId    int       `json:"agendaId"`
		ApprovedId  int       `json:"approvedId"`
		ReleasingId int       `json:"releasingId"`
		FilingId    int       `json:"filingId"`
		ItemNumber  string    `json:"itemNumber"`
		DocumentTag string    `json:"documentTag"`
		Remarks     string    `json:"remarks"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Trackings struct {
		TrackingId     int       `json:"trackingId" gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber string    `json:"trackingNumber" `
		ItemNumber     string    `json:"itemNumber"`
		LawType        string    `json:"lawType" `
		LawNumber      string    `json:"lawNumber" `
		Summary        string    `json:"summary" `
		ReceivedDate   string    `json:"receivedDate"`
		Calendared     string    `json:"calendared" `
		EnactedDate    string    `json:"enactedDate" `
		MayorDate      string    `json:"forwardedMayorDate"`
		StaCruzDate    string    `json:"ForwardedStaCruzDate"`
		ReleasedDate   string    `json:"releasedDate" `
		PublishedDate  string    `json:"publishedDate"`
		FiledDate      string    `json:"filedDate"`
		CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	CommitteeLists struct {
		ListId      int       `json:"listId"`
		ItemNumber  string    `json:"itemNumber"`
		CommitteeId int       `json:"committeeId"`
		UserId      int       `json:"userId"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

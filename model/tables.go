package model

import "time"

type (
	Receivings struct {
		ReceivingId      int       `json:"receivingId" gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber   string    `json:"trackingNumber,omitempty"`
		ReceivedDate     string    `json:"receivedDate,omitempty"`
		ReceivedTime     string    `json:"receivedTime,omitempty"`
		Receiver         string    `json:"receiver,omitempty"`
		Summary          string    `json:"summary,omitempty"`
		ReceivingTag     string    `json:"receivingTag,omitempty"`
		ReceivingRemarks string    `json:"receivingRemarks,omitempty"`
		ReceivedFile     string    `json:"receivedFile,omitempty"`
		Encoder          string    `json:"encoder,omitempty"`
		ModifiedBy       string    `json:"modifedBy,omitempty"`
		CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Agendas struct {
		AgendaId       int       `json:"agendaId" gorm:"primaryKey;autoIncrement:true"`
		ItemNumber     string    `json:"itemNumber,omitempty"`
		IsUrgent       bool      `json:"isUrgent,omitempty"`
		DateCalendared string    `json:"dateCalendared,omitempty"`
		DateReported   string    `json:"dateReported,omitempty"`
		Source         string    `json:"source,omitempty"`
		SourceResult   string    `json:"sourceResult,omitempty"`
		AgendaTag      string    `json:"agendaTag,omitempty"`
		AgendaRemarks  string    `json:"agendaRemarks,omitempty"`
		Encoder        string    `json:"encoder,omitempty"`
		ModifiedBy     string    `json:"modifedBy,omitempty"`
		CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Approves struct {
		ApproveId   int       `json:"approvedId" gorm:"primaryKey;autoIncrement:true"`
		ItemNumber  int       `json:"itemNumber,omitempty"`
		LawType     string    `json:"lawType,omitempty"`
		LawNumber   string    `json:"lawNumber,omitempty"`
		Series      string    `json:"series,omitempty"`
		EnactedDate string    `json:"enactedDate"`
		MotionedBy  string    `json:"motionedBy,omitempty"`
		Author      string    `json:"author,omitempty"`
		ResOrdFile  string    `json:"attachedFile,omitempty"`
		TitleBody   string    `json:"titleBody,omitempty"`
		Encoder     string    `json:"encoder,omitempty"`
		ModifiedBy  string    `json:"modifedBy,omitempty"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Routings struct {
		DocId       int       `json:"docId" gorm:"primaryKey;autoIncrement:true"`
		ReceivingId int       `json:"receivingId,omitempty"`
		AgendaId    int       `json:"agendaId,omitempty"`
		ApprovedId  int       `json:"approvedId,omitempty"`
		ReleasingId int       `json:"releasingId,omitempty"`
		FilingId    int       `json:"filingId,omitempty"`
		ItemNumber  string    `json:"itemNumber,omitempty"`
		DocumentTag string    `json:"documentTag,omitempty"`
		Remarks     string    `json:"remarks,omitempty"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Trackings struct {
		TrackingId     int       `json:"trackingId" gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber string    `json:"trackingNumber,omitempty"`
		ItemNumber     string    `json:"itemNumber,omitempty"`
		LawType        string    `json:"lawType,omitempty"`
		LawNumber      string    `json:"lawNumber,omitempty"`
		Summary        string    `json:"summary,omitempty"`
		ReceivedDate   string    `json:"receivedDate,omitempty"`
		Calendared     string    `json:"calendared,omitempty"`
		EnactedDate    string    `json:"enactedDate,omitempty"`
		MayorDate      string    `json:"forwardedMayorDate,omitempty"`
		StaCruzDate    string    `json:"ForwardedStaCruzDate,omitempty"`
		ReleasedDate   string    `json:"releasedDate,omitempty"`
		PublishedDate  string    `json:"publishedDate,omitempty"`
		FiledDate      string    `json:"filedDate,omitempty"`
		CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	CommitteeLists struct {
		ListId      int       `json:"listId" gorm:"primaryKey;autoIncrement:true"`
		ItemNumber  string    `json:"itemNumber,omitempty"`
		CommitteeId int       `json:"committeeId,omitempty"`
		UserId      int       `json:"userId,omitempty"`
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

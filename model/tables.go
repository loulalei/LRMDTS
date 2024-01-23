package model

import "time"

type (
	Receivings struct {
		ReceivingId      int       `json:"receivingId" gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber   string    `json:"trackingNumber,omitempty"`
		ReceivedDate     string    `json:"receivedDate,omitempty"`
		ReceivedTime     string    `json:"receivedTime,omitempty"`
		Receiver         string    `json:"receiver,omitempty"`
		Description      string    `json:"description,omitempty"`
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
		ItemNumber  string    `json:"itemNumber,omitempty"`
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

	Releasings struct {
		DocId                 int       `json:"docId,omitempty"`
		ReleasingId           int       `json:"releasingId" gorm:"primaryKey;autoIncrement:true"`
		MayorDateForwarded    string    `json:"mayorDateForwarded,omitempty"`
		MayorDateApproved     string    `json:"mayorDateApproved,omitempty"`
		IsApprovedLachesMayor bool      `json:"isApprovedLachesMayor,omitempty"`
		SPDateForwarded       string    `json:"spDateForwarded,omitempty"`
		SPDateApproved        string    `json:"spDateApproved,omitempty"`
		IsApprovedLachesSP    bool      `json:"isApprovedLachesSP,omitempty"`
		SPResolutionNumber    string    `json:"spResolutionNumber,omitempty"`
		SPResolutionFile      string    `json:"spResolutionFile,omitempty"`
		LocalDateRelease      string    `json:"localDateRelease,omitempty"`
		LocalDatePublished    string    `json:"localDatePublished,omitempty"`
		EndorsementFile       string    `json:"endorsementFile,omitempty"`
		Encoder               string    `json:"encoder,omitempty"`
		ModifiedBy            string    `json:"modifiedBy,omitempty"`
		CreatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Filings struct {
		DocId         int       `json:"docId,omitempty"`
		FilingId      int       `json:"filingId" gorm:"primaryKey;autoIncrement:true"`
		CabinetNumber string    `json:"cabinetNumber,omitempty"`
		FolderName    string    `json:"folderName,omitempty"`
		DateFiled     string    `json:"dateFiled,omitempty"`
		IsBorrowed    bool      `json:"isBorrowed,omitempty"`
		DateBorrowed  string    `json:"dateBorrowed,omitempty"`
		Borrower      string    `json:"borrower,omitempty"`
		DateReturned  string    `json:"dateReturned,omitempty"`
		DatePublished string    `json:"datePublished,omitempty"`
		Publisher     string    `json:"publisher,omitempty"`
		Encoder       string    `json:"encoder,omitempty"`
		ModifiedBy    string    `json:"modifiedBy,omitempty"`
		CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Trackings struct {
		DocId          int    `json:"docId,omitempty"`
		TrackingId     int    `json:"trackingId" gorm:"primaryKey;autoIncrement:true"`
		TrackingNumber string `json:"trackingNumber,omitempty"`
		ItemNumber     string `json:"itemNumber,omitempty"`
		LawType        string `json:"lawType,omitempty"`
		LawNumber      string `json:"lawNumber,omitempty"`
		Series         string `json:"series,omitempty"`
		Description    string `json:"description,omitempty"`
		Summary        string `json:"summary,omitempty"`
		ReceivedDate   string `json:"receivedDate,omitempty"`
		Calendared     string `json:"calendared,omitempty"`
		EnactedDate    string `json:"enactedDate,omitempty"`
		MayorDate      string `json:"forwardedMayorDate,omitempty"`
		StaCruzDate    string `json:"forwardedStaCruzDate,omitempty"`
		ReleasedDate   string `json:"releasedDate,omitempty"`
		PublishedDate  string `json:"publishedDate,omitempty"`
		FiledDate      string `json:"filedDate,omitempty"`
		//---------
		Author     string    `json:"author,omitempty"`
		MotionedBy string    `json:"motionedBy,omitempty"`
		CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

	Folders struct {
		FolderId int    `json:"folderId" gorm:"primaryKey;autoIncrement"`
		Name     string `json:"name"`
	}

	Cabinet struct {
		CabinetNumber string `json:"cabinetNumber"`
		FolderId      int    `json:"folderId"`
		DocId         int    `json:"docId"`
	}

	Divisions struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	DocumentFolders struct {
		FolderId   int    `json:"folderId" gorm:"primaryKey;autoIncrement"`
		FolderName string `json:"folderName,omitempty"`
		Encoder    string `json:"encoder,omitempty"`
	}

	BorrowerHistories struct {
		DocId        int       `json:"docId,omitempty"`
		BorrowerId   int       `json:"borrowerId" gorm:"primaryKey;autoIncrement"`
		Borrower     string    `json:"borrower,omitempty"`
		DateBorrowed string    `json:"dateBorrowed,omitempty"`
		DateReturned string    `json:"dateReturned,omitempty"`
		CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	ActivityLogger struct {
		ActivityId int       `json:"activityId" gorm:"primaryKey;autoIncrement"`
		Activity   string    `json:"activity"`
		Event      string    `json:"event"`
		UserId     int       `json:"userId"`
		CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	EmployeePerformace struct {
		PerformanceId    int       `json:"performanceId" gorm:"primaryKey;autoIncrement"`
		RecordsCaptured  string    `json:"recordsCaptured"`
		RecordsRetrieved string    `json:"RecordsRetrieved"`
		UserId           int       `json:"userId"`
		CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

type (
	ViewBrrowerHistory struct {
		BorrowerId   int    `json:"borrowerId"`
		LawType      string `json:"lawType,omitempty"`
		ResOrdFile   string `json:"resOrdfile,omitempty"`
		Borrower     string `json:"borrower,omitempty"`
		DateBorrowed string `json:"dateBorrowed,omitempty"`
		DateReturned string `json:"dateReturned,omitempty"`
	}
	RequestCommittees struct {
		Name string `json:"name"`
	}
)

type (
	UserActivities struct {
		UAId      int       `json:"uaid" gorm:"primaryKey;autoIncrement"`
		UserId    int       `json:"userId"`
		Activity  string    `json:"activity"`
		IsLogged  bool      `json:"isLogged" gorm:"default:false"`
		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	EventCalendar struct {
		EventId int    `json:"-" gorm:"primaryKey;autoIncrement"`
		Title   string `json:"title"`
		Start   string `json:"start"`
		GroupId string `json:"groupId,omitempty"`
		End     string `json:"end,omitempty"`
	}
)

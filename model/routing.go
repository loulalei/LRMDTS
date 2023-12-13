package model

type (
	RequestCommittee struct {
		ItemNumber  string `json:"itemNumber"`
		CommitteeId int    `json:"committeeId"`
		UserId      int    `json:"userId"`
	}
	CommitteeList struct {
		ListId      uint   `json:"ListId,omitempty" gorm:"primaryKey;autoIncrement:true"`
		ItemNumber  string `json:"itemNumber"`
		CommitteeId int    `json:"committeeId"`
		UserId      int    `json:"userId"`
	}

	ViewCommittees struct {
		ListId      int    `json:"listId,omitempty"`
		CommitteeId int    `json:"committeeId,omitempty"`
		ItemNumber  string `json:"itemNumber,omitempty"`
		Name        string `json:"name,omitempty"`
		Fullname    string `json:"fullname,omitempty"`
	}
)

// NEW MODEL FOR ROUTING
type (
	ViewRoutings struct {
		DocId       int    `json:"docId"`
		ItemNumber  string `json:"itemNumber"`
		DocumentTag string `json:"documentTag"`
		// Receivings
		ReceivingId    int    `json:"receivingId"`
		TrackingNumber string `json:"trackingNumber"`
		ReceivedDate   string `json:"receivedDate"`
		ReceivedTime   string `json:"receivedTime"`
		Receiver       string `json:"receiver"`
		Sender         string `json:"sender"`
		Summary        string `json:"summary"`
		Remarks        string `json:"remarks"`
		ReceivedFile   string `json:"receivedFile"`
		Encoder        string `json:"encoder"`
		UpdatedAt      string `json:"updatedAt"`
	}

	ViewAgenda struct {
		AgendaId       int    `json:"agendaId,omitempty"`
		ItemNumber     string `json:"itemNumber,omitempty"`
		IsUrgent       bool   `json:"isUrgent,omitempty"`
		DateCalendared string `json:"dateCalendared,omitempty"`
		DateReported   string `json:"dateReported,omitempty"`
		Source         string `json:"source,omitempty"`
		SourceResult   string `json:"sourceResult,omitempty"`
		AgendaTag      string `json:"agendaTag,omitempty"`
		AgendaRemarks  string `json:"agendaRemarks,omitempty"`
		Encoder        string `json:"encoder,omitempty"`
		ModifiedBy     string `json:"modifiedBy,omitempty"`
	}
)

// RECEIVING MODEL
type (
	RequestReceiving struct {
		TrackingNumber   string `json:"trackingNumber,omitempty"`
		ReceivedDate     string `json:"receivedDate,omitempty"`
		ReceivedTime     string `json:"receivedTime,omitempty"`
		Receiver         string `json:"receiver,omitempty"`
		Summary          string `json:"summary,omitempty"`
		ReceivingTag     string `json:"receivingTag,omitempty"`
		ReceivingRemarks string `json:"receivingRemarks,omitempty"`
		ReceivedFile     string `json:"receivedFile,omitempty"`
	}

	ViewReceiving struct {
		ReceivedId int `json:"receivingId"`
	}
)

type (
	RequestRoutingIdentifications struct {
		ReceivingId int `json:"receivingId,omitempty"`
		AgendaId    int `json:"agendaId,omitempty"`
	}
)

// FOR AGENDA
type (
	RequestForAgenda struct {
		DocId          int    `json:"docId,omitempty"`
		TrackingNumber string `json:"trackingNumber,omitempty"`
		ItemNumber     string `json:"itemNumber,omitempty"`
		IsUrgent       bool   `json:"isUrgent,omitempty"`
		DateCalendared string `json:"dateCalendared,omitempty"`
		DateReported   string `json:"dateReported,omitempty"`
		Source         string `json:"source,omitempty"`
		SourceResult   string `json:"sourceResult,omitempty"`
		Department     string `json:"department,omitempty"`
		Proponent      string `json:"proponent,omitempty"`
		Other          string `json:"other,omitempty"`
		AgendaTag      string `json:"agendaTag,omitempty"`
		AgendaRemarks  string `json:"agendaRemarks,omitempty"`
		Encoder        string `json:"encodeer,omitempty"`
		ModifiedBy     string `json:"modifiedBy,omitempty"`
	}

	RequestApproved struct {
		DocId       int    `json:"docId,omitempty"`
		ApproveId   int    `json:"approvedId,omitempty"`
		ItemNumber  string `json:"itemNumber,omitempty"`
		LawType     string `json:"lawType,omitempty"`
		LawNumber   string `json:"lawNumber,omitempty"`
		Series      string `json:"series,omitempty"`
		EnactedDate string `json:"enactedDate"`
		MotionedBy  string `json:"motionedBy,omitempty"`
		Author      string `json:"author,omitempty"`
		ResOrdFile  string `json:"attachedFile,omitempty"`
		TitleBody   string `json:"titleBody,omitempty"`
		Encoder     string `json:"encoder,omitempty"`
		ModifiedBy  string `json:"modifedBy,omitempty"`
	}

	RequestReleasing struct {
		DocId                 int    `json:"docId,omitempty"`
		ItemNumber            string `json:"itemNumber,omitempty"`
		ReleasingId           int    `json:"releasingId,omitempty"`
		MayorDateForwarded    string `json:"mayorDateForwarded,omitempty"`
		MayorDateApproved     string `json:"mayorDateApproved,omitempty"`
		IsApprovedLachesMayor bool   `json:"isApprovedLachesMayor,omitempty"`
		SPDateForwarded       string `json:"spDateForwarded,omitempty"`
		SPDateApproved        string `json:"spDateApproved,omitempty"`
		IsApprovedLachesSP    bool   `json:"isApprovedLachesSP,omitempty"`
		SPResolutionNumber    string `json:"spResolutionNumber,omitempty"`
		SPResolutionFile      string `json:"spResolutionFile,omitempty"`
		LocalDateRelease      string `json:"localDateRelease,omitempty"`
		LocalDatePublished    string `json:"localDatePublished,omitempty"`
		EndorsementFile       string `json:"endorsementFile,omitempty"`
		Publisher             string `json:"publisher,omitempty"`
		Encoder               string `json:"encoder,omitempty"`
		ModifiedBy            string `json:"modifiedBy,omitempty"`
	}
)

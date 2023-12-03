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
	}

	ViewAgendas struct {
		AgendaId   int    `json:"agendaId"`
		ItemNumber string `json:"itemNumber"`
		IsUrgent   bool   `json:"isUrgent"`
	}
)

// RECEIVING MODEL
type (
	RequestReceiving struct {
		TrackingNumber   string `json:"trackingNumber"`
		ReceivedDate     string `json:"receivedDate"`
		ReceivedTime     string `json:"receivedTime"`
		Receiver         string `json:"receiver"`
		Sender           string `json:"sender"`
		Summary          string `json:"summary"`
		ReceivingTag     string `json:"receivingTag"`
		ReceivingRemarks string `json:"receivingRemarks"`
		ReceivedFile     string `json:"receivedFile"`
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
		DocId          int    `json:"docId"`
		ItemNumber     string `json:"itemNumber"`
		IsUrgent       bool   `json:"isUrgent"`
		DateCalendared string `json:"dateCalendared"`
		DateReported   string `json:"dateReported"`
		Source         string `json:"source"`
		SourceResult   string `json:"sourceResult"`
		Department     string `json:"department"`
		Proponent      string `json:"proponent"`
		AgendaTag      string `json:"agendaTag"`
		AgendaRemarks  string `json:"agendaRemarks"`
	}
)

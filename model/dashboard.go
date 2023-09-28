package model

type (
	Document struct {
		DocId          int    `json:"docId" gorm:"primary key"`
		TrackingNumber string `json:"trackingNumber"`
		ItemNumber     string `json:"itemNumber"`
		ReceiveDate    string `json:"receiveDate"`
		ReceiveTime    string `json:"receiveTime"`
		Receiver       string `json:"receiver"`
		Sender         string `json:"Sender"`
		Summary        string `json:"summary"`
		DocumentTag    string `json:"documentTag"`
		Remarks        string `json:"remarks"`

		IsUrgent       bool   `json:"isUrgent"`
		DateCalendared string `json:"dateCalendared"`
		DateReported   string `json:"dateReported"`
		Source         string `json:"source"`
		Department     string `json:"department"`
		Proponent      string `json:"proponent"`
		Committee      string `json:"committee"`

		LawType     string `json:"lawType"`
		LawNo       string `json:"lawNo"`
		Series      string `json:"series"`
		DateEnacted string `json:"dateEnacted"`
		Title       string `json:"title"`
		MotionedBy  string `json:"motionedBy"`
		Author      string `json:"author"`

		DateFMayor string `json:"dateFMayor"`
		DateAMayor string `json:"dateAMayor"`

		DateFSP string `json:"dateFSP"`

		DateASP string `json:"dateASP"`
		SPResNo string `json:"SPResNo"`

		DatePublished string `json:"datePublished"`
		Publisher     string `json:"publisher"`
		DateReleased  string `json:"dateReleased"`

		Cabinet      string `json:"cabinet"`
		DateFiled    string `json:"dateFiled"`
		Folder       string `json:"folder"`
		IsBorrowed   bool   `json:"isBorrowed"`
		DateBorrowed string `json:"dateBorrowed"`
		Borrower     string `json:"borrower"`
		DateDisposed string `json:"dateDisposed"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	}

	Receiving struct {
		DocId          int    `json:"docId"`
		TrackingNumber string `json:"trackingNumber"`

		ReceiveDate string `json:"receiveDate"`
		ReceiveTime string `json:"receiveTime"`
		Receiver    string `json:"receiver"`
		Sender      string `json:"Sender"`
		Summary     string `json:"summary"`
	}

	ForAgenda struct {
		DocId          int    `json:"docId"`
		ItemNumber     string `json:"itemNumber"`
		IsUrgent       bool   `json:"isUrgent"`
		DateCalendared string `json:"dateCalendared"`
		DateReported   string `json:"dateReported"`
		Source         string `json:"source"`
		Department     string `json:"department"`
		Proponent      string `json:"proponent"`
		Committee      string `json:"committee"`
	}

	ResOrd struct {
		DocId       int    `json:"docId"`
		LawType     string `json:"lawType"`
		LawNo       string `json:"lawNo"`
		Series      string `json:"series"`
		DateEnacted string `json:"dateEnacted"`
		Title       string `json:"title"`
		MotionedBy  string `json:"motionedBy"`
		Author      string `json:"author"`
	}

	Forward struct {
		DocId      int    `json:"docId"`
		DateFMayor string `json:"dateFMayor"`
		DateAMayor string `json:"dateAMayor"`

		DateFSP string `json:"dateFSP"`

		DateASP string `json:"dateASP"`
		SPResNo string `json:"SPResNo"`
	}

	Releasing struct {
		DocId         int    `json:"docId"`
		DatePublished string `json:"datePublished"`
		Publisher     string `json:"publisher"`
		DateReleased  string `json:"dateReleased"`
	}

	Filing struct {
		DocId        int    `json:"docId"`
		Cabinet      string `json:"cabinet"`
		DateFiled    string `json:"dateFiled"`
		Folder       string `json:"folder"`
		IsBorrowed   bool   `json:"isBorrowed"`
		DateBorrowed string `json:"dateBorrowed"`
		Borrower     string `json:"borrower"`
		DateDisposed string `json:"dateDisposed"`
	}

	DocumentFilepath struct {
		DocId         int    `json:"docId"`
		ReceivedFile  string `json:"receivedFile"`
		ApprovedFile  string `json:"approvedFile"`
		MReceiveFile  string `json:"MreceiveFile"`
		SPEndorseFile string `json:"SPendorseFile"`
		EndorseFile   string `json:"endorseFile"`
		SPResFile     string `json:"SPresFile"`
	}

	Notification struct {
		DocId int `json:"docId"`
	}
)

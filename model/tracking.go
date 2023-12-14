package model

type (
	RequestTracking struct {
		MotionedBy string `json:"motionedBy,omitempty"`
		Author     string `json:"author,omitempty"`
	}
)

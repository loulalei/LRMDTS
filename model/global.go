package model

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

// Page Parameters
var (
	YearNow        = time.Now().Year()
	Page           string
	Title          string
	LoginStatus    int
	Fullname       string
	UserCodeLogged string
	UserID         int
	Token          string
	Interface      []interface{}
	SessionManager *session.Session
)

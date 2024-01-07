package global

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	// "github.com/gofiber/session/v2"
	"github.com/gofiber/template/html/v2"
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
	DivisionCode   string
	Token          string
	Interface      []interface{}
)

var SessionConfig = session.Config{
	KeyLookup:    "cookie:sessionID",
	CookieSecure: true,
}

var App = fiber.New(fiber.Config{
	UnescapePath: true,
	Views:        html.New("./template", ".html"),
})

var Store = session.New(SessionConfig)

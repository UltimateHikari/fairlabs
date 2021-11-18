package api

import (
	"github.com/labstack/echo"
)

type ActionType int

const (
	Get ActionType = iota + 1
	Post
)

type Kind struct {
	Action   ActionType
	Endpoint string
}

type Controller interface {
	Handle(e echo.Context) error
}

func RegisterPath(e *echo.Echo, k Kind, c Controller) {
	switch k.Action {
	case Get:
		e.GET(k.Endpoint, c.Handle)
	case Post:
		e.POST(k.Endpoint, c.Handle)
	}
}

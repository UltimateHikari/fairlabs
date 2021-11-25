package api

import (
	"github.com/labstack/echo"
)

type ActionType int

const (
	Get ActionType = iota + 1
	Post
)

type Group struct {
	Prefix string
}

type Kind struct {
	Action   ActionType
	Endpoint string
	Group    Group
}

type Controller interface {
	Handle(e echo.Context) error
}

func RegisterPath(e *echo.Echo, k Kind, c Controller) {
	group := e.Group(k.Group.Prefix)
	switch k.Action {
	case Get:
		group.GET(k.Endpoint, c.Handle)
	case Post:
		group.POST(k.Endpoint, c.Handle)
	}
}

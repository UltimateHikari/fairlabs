package adapter

import (
	"fairlabs-server/api"
	"net/http"

	"github.com/labstack/echo"
)

var MiscGroup = api.Group{Prefix: "v1/misc"}

//Control-Misc-Server
var LookupGoalKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/get/goal",
	Group:    MiscGroup}
var LookupStatsKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/get/stats",
	Group:    MiscGroup}
var ProgressKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/get/progress",
	Group:    MiscGroup}
var PriorityKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/priority",
	Group:    MiscGroup}
var GoalKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/goal",
	Group:    MiscGroup}

type LookupGoalController struct{}
type LookupStatsController struct{}
type ProgressController struct{}
type PriorityController struct{}
type GoalController struct{}

func (k LookupGoalController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "-4")
}

func (k LookupStatsController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "-3")
}

func (k ProgressController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "-2")
}

func (k PriorityController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "-1")
}

func (k GoalController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "0")
}

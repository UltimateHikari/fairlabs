package adapter

import (
	"fairlabs-server/adapter/conv"
	"fairlabs-server/api"
	"fairlabs-server/logic"
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
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	goal, err := logic.LookupGoalService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToMiscGoal(goal))
}

func (k LookupStatsController) Handle(c echo.Context) error {
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	cstats, err := logic.LookupStatsService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToMiscCStats(cstats))
}

func (k ProgressController) Handle(c echo.Context) error {
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	progress, err := logic.ProgressService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToMiscProgress(progress))
}

func (k PriorityController) Handle(c echo.Context) error {
	priorityRequest := new(conv.MiscPriorityRequest)

	if err := c.Bind(priorityRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, priority := priorityRequest.ToPriority()

	if err := logic.PriorityService(context, priority); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.NoContent(http.StatusOK)
}

func (k GoalController) Handle(c echo.Context) error {
	goalRequest := new(conv.MiscGoalRequest)

	if err := c.Bind(goalRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, goal := goalRequest.ToGoal()

	if err := logic.GoalService(context, goal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.NoContent(http.StatusOK)
}

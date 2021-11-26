package adapter

import (
	"fairlabs-server/adapter/conv"
	"fairlabs-server/api"
	"fairlabs-server/logic"
	"net/http"

	"github.com/labstack/echo"
)

var TasksGroup = api.Group{Prefix: "v1/tasks"}

//Control-Tasks
var QueueKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/queue",
	Group:    TasksGroup}
var MyCoursesGetKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/mycourses",
	Group:    TasksGroup}
var AllCoursesGetKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/allcourses",
	Group:    TasksGroup}
var FollowKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/follow",
	Group:    TasksGroup}

//Tasks
type QueueController struct{}
type MyCoursesGetController struct{}
type AllCoursesGetController struct{}
type FollowController struct{}

func (k QueueController) Handle(c echo.Context) error {
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	queue, err := logic.QueueService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToQueue(queue))
}

func (k MyCoursesGetController) Handle(c echo.Context) error {
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	courses, err := logic.MyCoursesGetService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToCourses(courses))
}

func (k AllCoursesGetController) Handle(c echo.Context) error {
	contextRequest := new(conv.UserContext)

	if err := c.Bind(contextRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	courses, err := logic.AllCoursesGetService(contextRequest.ToContext())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToCourses(courses))
}

func (k FollowController) Handle(c echo.Context) error {
	followRequest := new(conv.TasksFollowRequest)

	if err := c.Bind(followRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, course := followRequest.ToCourse()

	if err := logic.FollowService(context, course); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.NoContent(http.StatusOK)
}

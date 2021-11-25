package adapter

import (
	"fairlabs-server/api"
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
	return c.JSON(http.StatusOK, "6")
}

func (k MyCoursesGetController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "7")
}

func (k AllCoursesGetController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "8")
}

func (k FollowController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "9")
}

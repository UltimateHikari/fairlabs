package adapter

import (
	"fairlabs-server/api"
	"net/http"

	"github.com/labstack/echo"
)

var AdminGroup = api.Group{Prefix: "v1/admin"}

//Control-Admin-Server
var AddKind = api.Kind{
	Action:   api.Post,
	Endpoint: "",
	Group:    AdminGroup}
var AlgoGetKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/get/algo",
	Group:    AdminGroup}
var AlgoPostKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/algo",
	Group:    AdminGroup}
var ConditionGetKind = api.Kind{
	Action:   api.Get,
	Endpoint: "/get/cond",
	Group:    AdminGroup}
var ConditionPostKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/cond",
	Group:    AdminGroup}

//Admin
type AddController struct{}
type AlgoGetController struct{}
type AlgoPostController struct{}
type ConditionGetController struct{}
type ConditionPostController struct{}

func (k AddController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "1")
}
func (k AlgoGetController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "2")
}
func (k AlgoPostController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "3")
}
func (k ConditionGetController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "4")
}
func (k ConditionPostController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "5")
}

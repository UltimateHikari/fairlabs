package adapter

import (
	"fairlabs-server/adapter/conv"
	"fairlabs-server/api"
	"fairlabs-server/logic"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var AdminGroup = api.Group{Prefix: "v1/admin"}

//Control-Admin-Server
var AddKind = api.Kind{
	Action:   api.Post,
	Endpoint: "/add",
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
	addRequest := new(conv.AdminAddRequest)

	if err := c.Bind(addRequest); err != nil {
		//TODO remove verbose message left for debug purpose
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, courseInfo := addRequest.ToCourseInfo()
	course, err := logic.AddService(context, courseInfo)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.JSON(http.StatusOK, conv.ToCourse(course))
}

func (k AlgoGetController) Handle(c echo.Context) error {
	algos, err := logic.AlgoGetService()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	algosResponse := conv.ToAlgoResponse(algos)
	return c.JSON(http.StatusOK, algosResponse)
}

func (k AlgoPostController) Handle(c echo.Context) error {
	algoRequest := new(conv.AdminAlgoPostRequest)

	if err := c.Bind(algoRequest); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, algo := algoRequest.ToAlgo()

	if err := logic.AlgoPostService(context, algo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.NoContent(http.StatusOK)
}

func (k ConditionGetController) Handle(c echo.Context) error {
	conds, err := logic.ConditionGetService()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	condsResponse := conv.ToConditionResponse(conds)
	return c.JSON(http.StatusOK, condsResponse)
}

func (k ConditionPostController) Handle(c echo.Context) error {
	condRequest := new(conv.AdminConditionPostRequest)

	if err := c.Bind(condRequest); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bind error")
	}

	context, cond := condRequest.ToCondition()

	if err := logic.ConditionPostService(context, cond); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Service error")
	}

	return c.NoContent(http.StatusOK)
}

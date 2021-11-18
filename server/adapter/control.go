package adapter

import (
	"net/http"

	"github.com/labstack/echo"
)

type LookupGoalsController struct {
}

type ProgressController struct {
}

func (k LookupGoalsController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "I have a dream")
}

func (k ProgressController) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "Progress? What is progress lmao")
}

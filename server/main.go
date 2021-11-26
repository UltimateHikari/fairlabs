package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	adr "fairlabs-server/adapter"
	"fairlabs-server/api"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const (
	serverPort = 8080
)

func registerAll(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,
			"I am api server and i have a dream")
	})
	api.RegisterPath(e, adr.AddKind, adr.AddController{})
	api.RegisterPath(e, adr.AlgoGetKind, adr.AlgoGetController{})
	api.RegisterPath(e, adr.AlgoPostKind, adr.AlgoPostController{})
	api.RegisterPath(e, adr.ConditionGetKind, adr.ConditionGetController{})
	api.RegisterPath(e, adr.ConditionPostKind, adr.ConditionPostController{})

	api.RegisterPath(e, adr.LookupGoalKind, adr.LookupGoalController{})
	api.RegisterPath(e, adr.LookupStatsKind, adr.LookupStatsController{})
	api.RegisterPath(e, adr.ProgressKind, adr.ProgressController{})
	api.RegisterPath(e, adr.PriorityKind, adr.PriorityController{})
	api.RegisterPath(e, adr.GoalKind, adr.GoalController{})

	api.RegisterPath(e, adr.QueueKind, adr.QueueController{})
	api.RegisterPath(e, adr.MyCoursesGetKind, adr.MyCoursesGetController{})
	api.RegisterPath(e, adr.AllCoursesGetKind, adr.AllCoursesGetController{})
	api.RegisterPath(e, adr.FollowKind, adr.FollowController{})
}

func main() {
	e := echo.New()

	registerAll(e)

	go func() {
		address := fmt.Sprintf(":%d", serverPort)
		fmt.Println(address)
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Info("Exiting")

	//graceful shutdown from docs
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

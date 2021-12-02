package dbcontrol

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type DBControl struct {
	pool *pgxpool.Pool
}

var once sync.Once
var controlInstance *DBControl = nil

func GetInstance() *DBControl {
	if controlInstance == nil {
		once.Do(
			func() {
				controlInstance = &DBControl{}
				controlInstance.init()
			})
	}
	return controlInstance
}

func (c *DBControl) init() {
	log.Info("connecting to: " + os.Getenv("DATABASE_URL"))
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	controlInstance.pool = dbpool
}

func CloseInstance() {
	controlInstance.pool.Close()
}

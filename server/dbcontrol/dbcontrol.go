package dbcontrol

import (
	"context"
	"fairlabs-server/logic/spec"
	"sync"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBControl struct {
	pool *pgxpool.Pool
}

var once sync.Once
var controlInstance *DBControl = nil

func getInstance() *DBControl {
	if controlInstance == nil {
		once.Do(
			func() {
				controlInstance = &DBControl{}
			})
	}
	return controlInstance
}

func (c *DBControl) connect(pool *pgxpool.Pool) {
	controlInstance.pool = pool
}

func (c *DBControl) getAlgorithms() ([]*spec.Algo, error) {
	var algos []*spec.Algo
	//TODO improve error handling with context
	ctx := context.Background()
	pgxscan.Select(ctx, c.pool, &algos, `SELECT * FROM algo`)
	return algos, nil
}

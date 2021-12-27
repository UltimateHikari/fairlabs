package dbcontrol

import (
	"context"
	"fairlabs-server/logic/spec"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/labstack/gommon/log"
)

const (
	lookup_goal_query = `SELECT user_goal FROM participants WHERE
	user_id = (SELECT user_id FROM users WHERE email = $1) AND
	course_id = $2;`
	lookup_stats_query    = ""
	lookup_progress_query = ""
	goal_query            = `UPDATE participants SET user_goal = $1 WHERE 
	user_id = (SELECT user_id FROM users WHERE email = $2) AND
	course_id = $3;`
	priority_query = `UPDATE participants SET user_priority = $1 WHERE 
	user_id = (SELECT user_id FROM users WHERE email = $2) AND
	course_id = $3;`
)

func (c *DBControl) LookupGoal(email string, course_id int) (*spec.Goal, error) {
	var err error
	ctx := context.Background()
	log.Info(course_id, email)
	rows, errr := c.pool.Query(ctx, lookup_goal_query, email, course_id)
	if errr != nil {
		log.Error(errr)
		return nil, errr
	}

	var goal spec.Goal
	if err = pgxscan.ScanOne(&goal, rows); err != nil {
		log.Error(err)
	}

	return &goal, err
}

func (c *DBControl) SetPriority(priority bool, email string, course_id int) error {
	var err error
	ctx := context.Background()
	if _, err = c.pool.Exec(
		ctx,
		priority_query,
		priority,
		email,
		course_id); err != nil {
		log.Error(err)
	}
	return err
}

func (c *DBControl) SetGoal(goal int, email string, course_id int) error {
	var err error
	ctx := context.Background()
	if _, err = c.pool.Exec(
		ctx,
		goal_query,
		goal,
		email,
		course_id); err != nil {
		log.Error(err)
	}
	return err
}

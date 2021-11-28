package dbcontrol

import (
	"context"
	"fairlabs-server/logic/spec"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/labstack/gommon/log"
)

const (
	is_admin_query       = "SELECT COUNT(*) FROM users WHERE users.email = $1 AND users.is_admin = TRUE;"
	get_algo_query       = `SELECT * FROM algos;`
	get_cond_query       = `SELECT * FROM conds;`
	create_query         = "INSERT INTO courses(course_name, university_group) VALUES($1,$2)"
	create_get_tag_query = "SELECT course_id FROM courses WHERE course_name = $1 AND university_group = $2"
)

func (c *DBControl) IsAdmin(email string) bool {
	var count int64
	if err := c.pool.QueryRow(
		context.Background(),
		is_admin_query,
		email).Scan(&count); err != nil {
		log.Error(err)
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func (c *DBControl) GetAlgorithms() ([]*spec.Algo, error) {
	var algos []*spec.Algo
	var err error
	ctx := context.Background()
	if err = pgxscan.Select(ctx, c.pool, &algos, get_algo_query); err != nil {
		log.Error(err)
	}
	return algos, err
}

func (c *DBControl) GetConditions() ([]*spec.Condition, error) {
	var conds []*spec.Condition
	var err error
	ctx := context.Background()
	if err := pgxscan.Select(ctx, c.pool, &conds, get_cond_query); err != nil {
		log.Error(err)
	}
	return conds, err
}

func (c *DBControl) CreateCourse(cinfo *spec.CourseInfo) (int, error) {
	var err error
	var count int
	if _, err = c.pool.Exec(
		context.Background(),
		create_query,
		cinfo.Name,
		cinfo.Group); err != nil {
		log.Error(err)
		return count, err
	}
	if err = c.pool.QueryRow(
		context.Background(),
		create_get_tag_query,
		cinfo.Name,
		cinfo.Group).Scan(&count); err != nil {
		log.Error(err)
	}

	return count, err
}

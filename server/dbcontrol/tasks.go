package dbcontrol

import (
	"context"
	"fairlabs-server/logic/spec"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/labstack/gommon/log"
)

const (
	//todo 1-> $1
	queue_query = `SELECT user_name, user_priority FROM 
	unnest((SELECT queue FROM courses WHERE course_id = $1)) queue_id 
	JOIN participants ON course_id = $1 AND queue_id = user_id
	JOIN users ON queue_id = users.user_id;`
	my_courses_query = `SELECT course_id, course_name, university_group FROM courses WHERE 
	course_id = (SELECT course_id FROM participants WHERE
	user_id = (SELECT user_id FROM users WHERE email = $1));`
	all_courses_query = "SELECT course_id, course_name, university_group FROM courses;"
	follow_query      = `INSERT INTO participants(user_id, course_id, $3) 
	SELECT user_id, $2 as course_id FROM users WHERE email = $1;`
	teacher_role = "TEACHER"
)

func (c *DBControl) GetQueue(course_id int) (*spec.Queue, error) {
	var queue spec.Queue
	var err error
	ctx := context.Background()
	if err := pgxscan.Select(ctx, c.pool, &(queue.Queue), queue_query, course_id); err != nil {
		log.Error(err)
	}
	return &queue, err
}

func (c *DBControl) GetMyCourses(email string) ([]*spec.Course, error) {
	var courses []*spec.Course
	var err error
	ctx := context.Background()
	if err := pgxscan.Select(ctx, c.pool, &courses, my_courses_query, email); err != nil {
		log.Error(err)
	}
	return courses, err
}

func (c *DBControl) GetAllCourses() ([]*spec.Course, error) {
	var courses []*spec.Course
	var err error
	ctx := context.Background()
	if err := pgxscan.Select(ctx, c.pool, &courses, all_courses_query); err != nil {
		log.Error(err)
	}
	return courses, err
}

func (c *DBControl) Follow(email string, id int) error {
	var err error
	if _, err = c.pool.Exec(
		context.Background(),
		follow_query,
		email,
		id,
		teacher_role); err != nil {
		log.Error(err)
	}
	return err
}

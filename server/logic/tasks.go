package logic

import (
	"errors"
	db "fairlabs-server/dbcontrol"
	"fairlabs-server/logic/spec"

	"github.com/labstack/gommon/log"
)

//TODO checks for course id

func QueueService(context *spec.Context) (*spec.Queue, error) {
	return db.GetInstance().GetQueue(context.CourseId)
}

func MyCoursesGetService(context *spec.Context) ([]*spec.Course, error) {
	if context.Email == "" {
		return nil, errors.New("Bad context: no email")
	}
	courses, err := db.GetInstance().GetMyCourses(context.Email)
	log.Info(courses)
	return courses, err
}

func AllCoursesGetService(context *spec.Context) ([]*spec.Course, error) {
	if err := checkPrivilieges(context); err != nil {
		log.Error(err)
		return nil, err
	}
	courses, err := db.GetInstance().GetAllCourses()
	return courses, err
}

func FollowService(context *spec.Context, course *spec.Course) error {
	if context.Email == "" {
		return errors.New("Bad context: no email")
	}
	return db.GetInstance().Follow(context.Email, course.Id)
}

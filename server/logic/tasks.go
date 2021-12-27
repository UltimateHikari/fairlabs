package logic

import (
	"errors"
	db "fairlabs-server/dbcontrol"
	"fairlabs-server/logic/spec"

	"github.com/labstack/gommon/log"
)

func QueueService(context *spec.Context) (*spec.Queue, error) {
	if err := VerifyEmail(context.Email); err != nil {
		return nil, err
	}
	if err := VerifyCourse(context.CourseId); err != nil {
		return nil, err
	}
	if err := VerifyCourseMembership(context.Email, context.CourseId); err != nil {
		return nil, err
	}
	return db.GetInstance().GetQueue(context.CourseId)
}

func MyCoursesGetService(context *spec.Context) ([]*spec.Course, error) {
	if err := VerifyEmail(context.Email); err != nil {
		return nil, err
	}
	courses, err := db.GetInstance().GetMyCourses(context.Email)
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
	if err := VerifyEmail(context.Email); err != nil {
		return err
	}
	if err := VerifyCourse(context.CourseId); err != nil {
		return err
	}
	return db.GetInstance().Follow(context.Email, course.Id)
}

func SubmitService(context *spec.Context, tasks *spec.Tasks) error {
	if err := VerifyEmail(context.Email); err != nil {
		return err
	}
	if err := VerifyCourse(context.CourseId); err != nil {
		return err
	}
	if err := VerifyCourseMembership(context.Email, context.CourseId); err != nil {
		return err
	}
	var err error
	switch tasks.Intent {
	case "PLANNED", "FINISHED":
		err = db.GetInstance().Submit(
			context.Email,
			context.CourseId,
			tasks)
	case "EMPTY":
		if err = checkPrivilieges(context); err != nil {
			log.Error(err)
		}
		err = db.GetInstance().Submit(
			context.Email,
			context.CourseId,
			tasks)
	default:
		err = errors.New("Wrong intent")
		log.Error(err)
	}
	return err
}

func QueryService(context *spec.Context, query *spec.Tasks) (*spec.Tasks, error) {
	// doin lil'
	if err := VerifyEmail(context.Email); err != nil {
		return nil, err
	}
	if err := VerifyCourse(context.CourseId); err != nil {
		return nil, err
	}
	if err := VerifyCourseMembership(context.Email, context.CourseId); err != nil {
		return nil, err
	}
	var err error
	var tasks *spec.Tasks
	switch query.Intent {
	case "PLANNED", "FINISHED", "EMPTY":
		tasks, err = db.GetInstance().QuerySubmitsWithStatus(
			context.Email,
			context.CourseId,
			query)
	default:
		err = errors.New("Wrong intent")
		log.Error(err)
	}
	return tasks, err
}

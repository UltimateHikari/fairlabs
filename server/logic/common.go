package logic

import (
	"errors"
	db "fairlabs-server/dbcontrol"

	"github.com/labstack/gommon/log"
)

func VerifyEmail(email string) error {
	if email == "" {
		return errors.New("Bad email")
	}
	return nil
}

func VerifyCourse(id int) error {
	if id < 1 {
		return errors.New("Bad course")
	}
	return nil
}

func VerifyCourseMembership(email string, course_id int) error {
	count, err := db.GetInstance().IsMember(email, course_id)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info(count)
	if count != 1 {
		log.Error("Not a course participant, count: ", count)
		return errors.New("Not a course participant")
	}
	return nil
}

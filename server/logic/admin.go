package logic

import (
	"errors"
	"fairlabs-server/logic/spec"

	db "fairlabs-server/dbcontrol"

	"github.com/labstack/gommon/log"
)

func checkPrivilieges(context *spec.Context) error {
	if context.Email == "" {
		//TODO look closely
		return errors.New("Bad context: no email")
	}
	if !db.GetInstance().IsAdmin(context.Email) {
		return errors.New(context.Email + " does't have privilieges")
	}

	return nil
}

func AddService(context *spec.Context, c *spec.CourseInfo) (*spec.Course, error) {
	var res spec.Course
	if err := checkPrivilieges(context); err != nil {
		log.Error(err)
		return &res, err
	}
	res.Name = c.Name
	res.Group = c.Group
	id, err := db.GetInstance().CreateCourse(c)
	if err != nil {
		log.Error(err)
	} else {
		res.Id = id
	}
	return &res, err
}

func AlgoGetService() ([]*spec.Algo, error) {
	algos, err := db.GetInstance().GetAlgorithms()
	return algos, err
}

func ConditionGetService() ([]*spec.Condition, error) {
	conds, err := db.GetInstance().GetConditions()
	return conds, err
}

func AlgoPostService(context *spec.Context, algo *spec.Algo) error {
	if err := checkPrivilieges(context); err != nil {
		log.Error(err)
		return err
	}
	//TODO:stub
	return nil
}

func ConditionPostService(context *spec.Context, algo *spec.Condition) error {
	if err := checkPrivilieges(context); err != nil {
		log.Error(err)
		return err
	}
	// TODO:stub
	return nil
}

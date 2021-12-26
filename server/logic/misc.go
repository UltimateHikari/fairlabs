package logic

import (
	"errors"
	db "fairlabs-server/dbcontrol"
	"fairlabs-server/logic/spec"
)

func LookupGoalService(context *spec.Context) (*spec.Goal, error) {
	if context.Email == "" {
		return nil, errors.New("Bad context: no email")
	}
	goal, err := db.GetInstance().LookupGoal(context.Email, context.CourseId)
	return goal, err
}

func LookupStatsService(context *spec.Context) (*spec.CStatsDTO, error) {
	return &spec.CStatsDTO{
			Progress: []spec.ProgressDTO{spec.ProgressDTO{
				Name:  "Andrey",
				Goal:  spec.GoalProgress{Current: 2, Max: 5},
				Conds: []spec.ConditionProgress{}}},
			Forecast: spec.Forecast{IsPositive: true}},
		nil
}

func ProgressService(context *spec.Context) (*spec.ProgressDTO, error) {
	return &spec.ProgressDTO{
			Name:  "Andrey",
			Goal:  spec.GoalProgress{Current: 2, Max: 5},
			Conds: []spec.ConditionProgress{}},
		nil
}

func PriorityService(context *spec.Context, priority bool) error {
	if context.Email == "" {
		return errors.New("Bad context: no email")
	}
	return db.GetInstance().SetPriority(priority, context.Email, context.CourseId)
}

func GoalService(context *spec.Context, goal *spec.Goal) error {
	if context.Email == "" {
		return errors.New("Bad context: no email")
	}
	return db.GetInstance().SetGoal(goal.Mark, context.Email, context.CourseId)
}

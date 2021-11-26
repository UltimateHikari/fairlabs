package logic

import "fairlabs-server/logic/spec"

func LookupGoalService(context *spec.Context) (*spec.Goal, error) {
	return &spec.Goal{Mark: 5}, nil
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
	return nil
}

func GoalService(context *spec.Context, goal *spec.Goal) error {
	return nil
}

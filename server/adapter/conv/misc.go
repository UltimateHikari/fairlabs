package conv

import "fairlabs-server/logic/spec"

type MiscGoal struct {
	Mark int `json:"mark"`
}

type MiscForecast struct {
}

type MiscProgress struct {
	Name  string                  `json:"name"`
	Goal  MiscGoalProgress        `json:"goal"`
	Conds []MiscConditionProgress `json:"conds"`
}

type MiscGoalRequest struct {
	UserContext
	Mark int `json:"mark"`
}

type MiscPriorityRequest struct {
	UserContext
	Priority bool `json:"priority"`
}

type MiscGoalProgress struct {
	Current int `json:"cur"`
	Max     int `json:"max"`
}

type MiscConditionProgress struct {
	Data []string `json:"data"`
}

func (r *MiscGoalRequest) ToGoal() (*spec.Context, *spec.Goal) {
	var goal spec.Goal
	goal.Mark = r.Mark
	return r.ToContext(), &goal
}

func (r *MiscPriorityRequest) ToPriority() (*spec.Context, bool) {
	return r.ToContext(), r.Priority
}

type MiscLookupStatsResponse struct {
	Progress []MiscProgress `json:"list"`
	Forecast MiscForecast   `json:"forecast"`
}

package conv

import "fairlabs-server/logic/spec"

type MiscGoal struct {
	Mark int `json:"mark"`
}

type MiscForecast struct {
	IsPositive bool
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

func ToMiscProgress(progress *spec.ProgressDTO) *MiscProgress {
	var res MiscProgress
	res.Name = progress.Name
	res.Goal = MiscGoalProgress{
		Current: progress.Goal.Current,
		Max:     progress.Goal.Max}
	res.Conds = make([]MiscConditionProgress, len(progress.Conds))
	for i, item := range progress.Conds {
		res.Conds[i].Data = item.Data
	}
	return &res
}

func ToMiscCStats(cstats *spec.CStatsDTO) *MiscLookupStatsResponse {
	var res MiscLookupStatsResponse
	res.Forecast = MiscForecast{IsPositive: cstats.Forecast.IsPositive}
	res.Progress = make([]MiscProgress, len(cstats.Progress))
	for i, item := range cstats.Progress {
		res.Progress[i] = *ToMiscProgress(&item)
	}
	return &res
}

func ToMiscGoal(goal *spec.Goal) *MiscGoal {
	var res MiscGoal
	res.Mark = goal.Mark
	return &res
}

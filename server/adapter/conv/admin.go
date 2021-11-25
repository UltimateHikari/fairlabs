package conv

import "fairlabs-server/logic/spec"

type AdminAddRequest struct {
	Context UserContext
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type AdminAlgoPostRequest struct {
	Context UserContext
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type AdminConditionPostRequest struct {
	Context UserContext
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Data    []string `json:"data"`
}

func (r *AdminAddRequest) toCourseInfo() *spec.CourseInfo {
	var courseInfo spec.CourseInfo
	courseInfo.Id = r.Id
	courseInfo.Name = r.Name
	return &courseInfo
}

func (r *AdminAlgoPostRequest) toAlgo() *spec.Algo {
	var algo spec.Algo
	algo.Id = r.Id
	algo.Name = r.Name
	return &algo
}

func (r *AdminConditionPostRequest) toCondition() *spec.Condition {
	var condition spec.Condition
	condition.Id = r.Id
	condition.Name = r.Name
	condition.Data = r.Data
	return &condition
}

package conv

import "fairlabs-server/logic/spec"

type AdminAddRequest struct {
	UserContext
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AdminAlgo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AdminAlgoPostRequest struct {
	UserContext
	AdminAlgo
}

type AdminCondition struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Data []string `json:"data"`
}

type AdminConditionPostRequest struct {
	UserContext
	AdminCondition
}

func (r *AdminAddRequest) ToCourseInfo() *spec.CourseInfo {
	var courseInfo spec.CourseInfo
	courseInfo.Id = r.Id
	courseInfo.Name = r.Name
	return &courseInfo
}

func (r *AdminAlgoPostRequest) ToAlgo() (*spec.Context, *spec.Algo) {
	var context spec.Context
	var algo spec.Algo
	context.Email = r.UserEmail
	context.Name = r.UserName
	context.Group = r.UserGroup
	context.CourseId = r.UserCourseId
	//TODO remove double context unmarshal
	algo.Id = r.Id
	algo.Name = r.Name
	return &context, &algo
}

func (r *AdminConditionPostRequest) ToCondition() (*spec.Context, *spec.Condition) {
	var context spec.Context
	var condition spec.Condition
	context.Email = r.UserEmail
	context.Name = r.UserName
	context.Group = r.UserGroup
	context.CourseId = r.UserCourseId

	condition.Id = r.Id
	condition.Name = r.Name
	condition.Data = r.Data
	return &context, &condition
}

type AdminAlgoGetResponse struct {
	Algos []AdminAlgo `json:"algos"`
}

type AdminConditionGetResponse struct {
	Conds []AdminCondition `json:"conditions"`
}

func ToAlgoResponse(algos []spec.Algo) *AdminAlgoGetResponse {
	var response AdminAlgoGetResponse
	response.Algos = make([]AdminAlgo, len(algos))
	for i, item := range algos {
		response.Algos[i].Id = item.Id
		response.Algos[i].Name = item.Name
	}
	return &response
}

func ToConditionResponse(conditions []spec.Condition) *AdminConditionGetResponse {
	var response AdminConditionGetResponse
	response.Conds = make([]AdminCondition, len(conditions))
	for i, item := range conditions {
		response.Conds[i].Id = item.Id
		response.Conds[i].Name = item.Name
	}
	return &response
}

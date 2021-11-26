package logic

import "fairlabs-server/logic/spec"

func AddService(context *spec.Context, c *spec.CourseInfo) error {
	//TODO:stub
	return nil
}

func AlgoGetService() ([]spec.Algo, error) {
	return []spec.Algo{}, nil
}

func ConditionGetService() ([]spec.Condition, error) {
	return []spec.Condition{}, nil
}

func AlgoPostService(context *spec.Context, algo *spec.Algo) error {
	// TODO:stub
	return nil
}

func ConditionPostService(context *spec.Context, algo *spec.Condition) error {
	// TODO:stub
	return nil
}

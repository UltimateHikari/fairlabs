package logic

import "fairlabs-server/logic/spec"

func QueueService(context *spec.Context) (*spec.Queue, error) {
	return &spec.Queue{
			Queue: []spec.QueueEntry{
				spec.QueueEntry{
					Name:          "Andrey",
					IsPrioritized: true}}},
		nil
}

func MyCoursesGetService(context *spec.Context) ([]spec.Course, error) {
	return []spec.Course{
			spec.Course{
				Id:    1,
				Name:  "WackoCourse",
				Group: 19201}},
		nil
}

func AllCoursesGetService(context *spec.Context) ([]spec.Course, error) {
	return []spec.Course{
			spec.Course{
				Id:    1,
				Name:  "WackoCourse",
				Group: 19201},
			spec.Course{
				Id:    2,
				Name:  "BestCourse",
				Group: 19212}},
		nil
}

func FollowService(context *spec.Context, course *spec.Course) error {
	return nil
}

package conv

import "fairlabs-server/logic/spec"

type TasksQueue struct {
	Queue []TasksQueueEntry `json:"queue"`
}

type TasksQueueEntry struct {
	Name          string `json:"name"`
	IsPrioritized bool   `json:"priority"`
}

type TasksCourse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CourseGroup int    `json:"cgroup"`
}

type TasksFollowRequest struct {
	UserContext
	TasksCourse
}

func (r *TasksFollowRequest) ToCourse() (*spec.Context, *spec.Course) {
	var course spec.Course
	course.Id = r.Id
	course.Name = r.Name
	course.Group = r.CourseGroup
	return r.ToContext(), &course
}

func ToCourses(courses []spec.Course) []TasksCourse {
	res := make([]TasksCourse, len(courses))
	for i, item := range courses {
		res[i].Id = item.Id
		res[i].Name = item.Name
		res[i].CourseGroup = item.Group
	}
	return res
}

func ToQueue(queue *spec.Queue) *TasksQueue {
	var taskQueue TasksQueue
	taskQueue.Queue = make([]TasksQueueEntry, len(queue.Queue))
	for i, item := range queue.Queue {
		taskQueue.Queue[i].Name = item.Name
		taskQueue.Queue[i].IsPrioritized = item.IsPrioritized
	}
	return &taskQueue
}

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

type CourseListResponse struct {
	Courses []TasksCourse `json:"courses"`
}

type Tasks struct {
	Intent    string  `json:"intent"`
	TasksList []int32 `json:"tasks"`
}

type TasksRequest struct {
	UserContext
	Tasks
}

type TasksListResponse struct {
	Tasks []int32 `json:"tasks"`
}

func (r *TasksRequest) ToTasks() (*spec.Context, *spec.Tasks) {
	var tasks spec.Tasks
	tasks.Intent = r.Intent
	tasks.Tasks = r.TasksList
	return r.ToContext(), &tasks
}

func (r *TasksFollowRequest) ToCourse() (*spec.Context, *spec.Course) {
	var course spec.Course
	course.Id = r.Id
	course.Name = r.Name
	course.Group = r.CourseGroup
	return r.ToContext(), &course
}

func ToTasks(tasks *spec.Tasks) *TasksListResponse {
	var response TasksListResponse
	response.Tasks = make([]int32, len(tasks.Tasks))
	for i, item := range tasks.Tasks {
		response.Tasks[i] = item
	}
	return &response
}

func ToCourses(courses []*spec.Course) *CourseListResponse {
	var response CourseListResponse
	response.Courses = make([]TasksCourse, len(courses))
	for i, item := range courses {
		response.Courses[i].Id = item.Id
		response.Courses[i].Name = item.Name
		response.Courses[i].CourseGroup = item.Group
	}
	return &response
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

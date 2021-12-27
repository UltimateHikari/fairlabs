package spec

//Admin's
type CourseInfo struct {
	Name   string `db:"course_name"`
	Group  int    `db:"university_group"`
	Amount int    `db:"tasks_amount"`
}

type Algo struct {
	Id   int    `db:"algo_id"`
	Name string `db:"algo_name"`
}

type Condition struct {
	Id   int    `db:"cond_id"`
	Name string `db:"cond_name"`
	// Condition-specific
	Data []string `db:"example_data"`
}

type Context struct {
	Email    string
	Name     string
	Group    int
	CourseId int
}

//Misc's
type CStatsDTO struct {
	Progress []ProgressDTO
	Forecast Forecast
}

type Forecast struct {
	IsPositive bool
}

type ProgressDTO struct {
	Name  string
	Goal  GoalProgress
	Conds []ConditionProgress
}

type GoalProgress struct {
	Current int
	Max     int
}

type ConditionProgress struct {
	Data []string // Condition-specific
}

type Goal struct {
	Mark int `db:"user_goal"`
}

type Priority struct {
	Priority bool `db:"user_priority"`
}

type Course struct {
	Id    int    `db:"course_id"`
	Name  string `db:"course_name"`
	Group int    `db:"university_group"`
}

//tasks
type QueueEntry struct {
	Name          string `db:"user_name"`
	IsPrioritized bool   `db:"user_priority"`
}

type Queue struct {
	Queue []QueueEntry
}

type Tasks struct {
	Intent string
	Tasks  []int32 `db:"task_id"`
}

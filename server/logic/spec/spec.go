package spec

//Admin's
type CourseInfo struct {
	Id   int
	Name string
	//TODO:extend
}

type Algo struct {
	Id   int
	Name string
}

type Condition struct {
	Id   int
	Name string
	Data []string // Condition-specific
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
	Mark int
}

type Course struct {
	Id    int
	Name  string
	Group int
}

//tasks
type QueueEntry struct {
	Name          string
	IsPrioritized bool
}

type Queue struct {
	Queue []QueueEntry
}

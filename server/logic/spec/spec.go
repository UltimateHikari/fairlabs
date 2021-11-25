package spec

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
	Data []string
}

type Context struct {
	Email    string
	Name     string
	Group    int
	CourseId int
}

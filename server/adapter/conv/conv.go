package conv

import "fairlabs-server/logic/spec"

// checks for admin or not are in logic
type UserContext struct {
	UserEmail    string `query:"email"`
	UserName     string `query:"username"`
	UserGroup    int    `query:"group"`
	UserCourseId int    `query:"course"`
}

func (c *UserContext) ToContext() *spec.Context {
	var context spec.Context
	context.Email = c.UserEmail
	context.Name = c.UserName
	context.Group = c.UserGroup
	context.CourseId = c.UserCourseId
	return &context
}

package conv

// checks for admin or not are in logic
type UserContext struct {
	UserEmail string `json:"email"`
	UserName  string `json:"username"`
	UserGroup string `json:"group"`
}

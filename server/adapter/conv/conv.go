package conv

// checks for admin or not are in logic
type UserContext struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

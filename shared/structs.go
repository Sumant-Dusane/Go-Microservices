package shared

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

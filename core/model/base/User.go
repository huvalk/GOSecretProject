package baseModels

type User struct {
	ID       int    `json:"ID"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Session  string `json:"session"`
}

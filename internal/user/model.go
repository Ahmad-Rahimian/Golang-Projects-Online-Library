package user

type User struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

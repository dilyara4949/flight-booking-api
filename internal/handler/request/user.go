package request

type UpdateUser struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

package request

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

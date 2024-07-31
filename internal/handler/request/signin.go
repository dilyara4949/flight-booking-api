package request

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

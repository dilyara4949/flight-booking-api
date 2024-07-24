package request

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	NewPassword string `json:"new_password"`
}

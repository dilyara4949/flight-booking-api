package request

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

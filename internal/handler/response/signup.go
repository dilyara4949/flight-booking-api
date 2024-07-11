package response

type Signup struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}

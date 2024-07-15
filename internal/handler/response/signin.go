package response

type Signin struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}

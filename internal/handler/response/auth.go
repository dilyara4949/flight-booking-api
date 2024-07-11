package response

type Error struct {
	Error string `json:"error"`
}

type Signup struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}

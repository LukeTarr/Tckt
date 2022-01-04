package api

type Register struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

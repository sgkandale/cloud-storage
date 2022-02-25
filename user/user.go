package user

type User struct {
	Username     string `json:"username"`
	PasswordText string `json:"-"`
	PasswordHash string `json:"-"`
}

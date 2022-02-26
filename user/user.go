package user

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordText string `json:"-"`
	PasswordHash string `json:"-"`
	CreatedAt    int64  `json:"-"`
	CreatedAtStr string `json:"created_at"`
}

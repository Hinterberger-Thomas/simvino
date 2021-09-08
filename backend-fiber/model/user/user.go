package user

type User struct {
	UserID string `json:"uuid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type UserTokens struct {
	Email  string      `json:"email"`
	Tokens []UserToken `json:"tokens"`
}

type UserToken struct {
	Token    string
	checksum string
}

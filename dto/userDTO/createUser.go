package userdto

type CreateUser struct {
	Email    string `json:"string"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

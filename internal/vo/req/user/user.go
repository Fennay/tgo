package user

type CreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

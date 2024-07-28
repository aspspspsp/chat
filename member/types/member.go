package types

type RegisterReq struct {
	Nickname string `form:"nickname" json:"nickname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginReq struct {
	Username string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type TokenData struct {
	Member       interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type InfoResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

package types

type MemberRegisterReq struct {
	Nickname string `form:"nickname" json:"nickname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

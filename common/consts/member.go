package consts

import "time"

const (
	UserDefaultAvatarOss     = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640" // OSS的默认头像
	MemberDefaultAvatarLocal = "avatar.JPG"                                   // OSS的默认头像
)

const (
	AccessTokenExpireDuration  = 24 * time.Hour
	RefreshTokenExpireDuration = 10 * 24 * time.Hour
)

const EncryptMoneyKeyLength = 6

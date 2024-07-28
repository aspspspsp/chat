package services

import (
	"common/configs"
	"common/consts"
	"context"
	"errors"
	"log"
	"member/repository/db/dao"
	"member/repository/db/models"
	"member/types"
	"sync"
)

var MemberSrvIns *MemberSrv
var MemberSrvOnce sync.Once

type MemberSrv struct {
}

func GetMemberSrv() *MemberSrv {
	MemberSrvOnce.Do(func() {
		MemberSrvIns = &MemberSrv{}
	})
	return MemberSrvIns
}

// Register 用户注册
func (s *MemberSrv) Register(ctx context.Context, req *types.MemberRegisterReq) (resp interface{}, err error) {
	memberDao := dao.NewMemberDao(ctx)
	_, exist, err := memberDao.IsExistsByUsername(req.Username)
	if err != nil {
		log.Fatal(err)
		return
	}
	if exist {
		err = errors.New("用户已经存在了")
		return
	}
	member := &models.Member{
		Nickname: req.Nickname,
		Username: req.Username,
		Status:   models.Active,
	}
	// 加密密码
	if err = member.SetPassword(req.Password); err != nil {
		log.Fatal(err)
		return
	}

	// 默认头像走的local
	member.Avatar = consts.MemberDefaultAvatarLocal
	if configs.GetConfig().System.UploadModel == consts.UploadModelOss {
		// 如果配置是走oss，则用url作为默认头像
		member.Avatar = consts.UserDefaultAvatarOss
	}

	// 创建用户
	err = memberDao.Create(member)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

package services

import (
	"common/configs"
	"common/consts"
	"common/repository/db/models"
	"context"
	"errors"
	"github.com/CocaineCong/gin-mall/pkg/utils/jwt"
	"log"
	"member/repository/db/dao"
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
func (s *MemberSrv) Register(ctx context.Context, req *types.RegisterReq) (resp interface{}, err error) {
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

// Login 用户登陆函数
func (s *MemberSrv) Login(ctx context.Context, req *types.LoginReq) (resp interface{}, err error) {
	var member *models.Member
	memberDao := dao.NewMemberDao(ctx)
	member, exist, err := memberDao.IsExistsByUsername(req.Username)
	if !exist { // 如果查询不到，返回相应的错误
		log.Fatal(err)
		return nil, errors.New("用户不存在")
	}

	if !member.CheckPassword(req.Password) {
		return nil, errors.New("账号/密码不正确")
	}

	accessToken, refreshToken, err := jwt.GenerateToken(member.ID, req.Username)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	memberInfoResp := &types.InfoResp{
		ID:       member.ID,
		Username: member.Username,
		Nickname: member.Nickname,
		Email:    member.Email,
		Status:   member.Status,
		Avatar:   member.AvatarURL(),
		CreateAt: member.CreatedAt.Unix(),
	}

	resp = &types.TokenData{
		Member:       memberInfoResp,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

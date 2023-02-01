package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/tool"
	"mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	//verify user exists
	userDB, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, xerr.NewErrCodeMsg(xerr.DataQueryError, err.Error())
		}
		return nil, err
	}

	//verify user password
	md5ByString, _ := tool.Md5ByString(in.Password)
	if md5ByString != userDB.Password {
		return nil, xerr.NewErrMsg("账号或密码错误")
	}

	//return sql
	var resp user.LoginResponse
	_ = copier.Copy(&resp, userDB)

	return &resp, nil
}

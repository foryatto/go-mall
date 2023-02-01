package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *UserInfoLogic) UserInfo(in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	userDB, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, xerr.NewErrCodeMsg(xerr.DataQueryError, err.Error())
		}
		return nil, err
	}

	var resp user.UserInfo
	_ = copier.Copy(&resp, userDB)
	resp.CreateTime = userDB.CreateTime.Unix()
	resp.UpdateTime = userDB.UpdateTime.Unix()
	return &user.UserInfoResponse{
		User: &resp,
	}, nil
}

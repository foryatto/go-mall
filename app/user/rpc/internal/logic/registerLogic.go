package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/tool"
	"mall/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	//verify user exists
	_, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil {
		return nil, xerr.NewErrCode(xerr.UserExistedError)
	}
	if err == model.ErrNotFound {
		pwd, err := tool.Md5ByString(in.Password)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.UserRegisterError)
		}
		newUser := model.User{
			Username: in.Username,
			Password: pwd,
			Phone:    in.Phone,
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, xerr.NewErrCodeMsg(xerr.DataInsertError, err.Error())
		}
		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, xerr.NewErrCodeMsg(xerr.DataInsertError, err.Error())
		}
		return &user.RegisterResponse{
			Id:       newUser.Id,
			Username: newUser.Username,
			Phone:    newUser.Phone,
		}, nil
	}
	return nil, xerr.NewErrCodeMsg(xerr.DataQueryError, err.Error())
}

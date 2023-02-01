package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAddressLogic {
	return &DelUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收货地址
func (l *DelUserAddressLogic) DelUserAddress(in *pb.UserAddressDelReq) (*pb.UserAddressDelResp, error) {
	_, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "收获地址不存在")
		}
		return nil, err
	}
	err = l.svcCtx.UserAddressModel.Delete(l.ctx, in.GetId())
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserReceiveAddress Database Exception : %+v , err: %v", in.Id, err)
	}
	return &user.UserAddressDelResp{}, nil
}

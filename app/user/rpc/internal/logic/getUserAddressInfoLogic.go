package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressInfoLogic {
	return &GetUserAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据主键id,查询收货地址
func (l *GetUserAddressInfoLogic) GetUserAddressInfo(in *pb.UserAddressInfoReq) (*pb.UserAddress, error) {
	address, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			//return nil, status.Error(100, "收获地址数据不存在")
			return nil, errors.Wrap(xerr.NewErrMsg("收获地址数据不存在"), "收获地址数据不存在")
		}
		return nil, err
	}
	var resp user.UserAddress
	_ = copier.Copy(&resp, address)
	resp.CreateTime = address.CreateTime.Unix()
	resp.UpdateTime = address.UpdateTime.Unix()
	return &resp, nil
}

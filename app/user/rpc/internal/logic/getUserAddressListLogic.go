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

type GetUserAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressListLogic {
	return &GetUserAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取收货地址列表
func (l *GetUserAddressListLogic) GetUserAddressList(in *pb.UserAddressListReq) (*pb.UserAddressListResp, error) {
	addressesList, err := l.svcCtx.UserAddressModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed  get user's receive address list err : %v , in :%+v", err, in)
	}
	var resp []*user.UserAddress
	for _, receiveAddresses := range addressesList {
		var pbAddress user.UserAddress
		_ = copier.Copy(&pbAddress, receiveAddresses)
		resp = append(resp, &pbAddress)
	}
	return &user.UserAddressListResp{
		List: resp,
	}, nil
}
